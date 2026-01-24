package controllers

import (
	"net/http"

	"backend/database"
	"backend/models"
	"backend/utils"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// ===============================
// CREATE USER (ADMIN)
// ===============================
func CreateUser(c *gin.Context) {
	var input struct {
		Name           string `json:"name" binding:"required"`
		Email          string `json:"email" binding:"required,email"`
		Phone          string `json:"phone" binding:"required"`
		PostAppliedFor string `json:"postAppliedFor"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	inviteUID := uuid.New().String()

	user := models.User{
		Name:           input.Name,
		Email:          input.Email,
		Phone:          input.Phone,
		PostAppliedFor: input.PostAppliedFor,
		InviteUID:      &inviteUID,
		Role:           "user",
	}

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 🔗 Invite link
	frontendBaseURL := os.Getenv("FRONTEND_BASE_URL")
	if frontendBaseURL == "" {
		frontendBaseURL = "http://localhost:57226"
	}
	inviteLink := frontendBaseURL + "/form?uid=" + inviteUID

	// ✉️ Email body
	emailBody := "Hello " + user.Name + ",\n\n" +
		"You have been invited to complete your job application.\n\n" +
		"Click the link below to continue:\n\n" +
		inviteLink + "\n\n" +
		"Regards,\nTalentDesk Team"

	// ✅ Send email asynchronously
	go func() {
		if err := utils.SendEmail(
			user.Email,
			"Complete Your Application",
			emailBody,
		); err != nil {
			log.Println("Email send failed:", err)
		}
	}()

	// ✅ RESPONSE (THIS WAS MISSING)
	c.JSON(http.StatusCreated, gin.H{
		"user":       user,
		"inviteLink": inviteLink,
	})
}



// ===============================
// GET /users/:id
// ===============================
func GetUserByID(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}
