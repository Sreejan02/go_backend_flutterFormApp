package controllers

import (
	"net/http"

	"backend/database"

	"github.com/gin-gonic/gin"
)

// ===============================
// RESPONSE DTO
// ===============================
type AdminUserResponse struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	PostAppliedFor string `json:"postAppliedFor"`
	InviteUid      string `json:"inviteUid"`
	Submitted      bool   `json:"submitted"`
}


func GetAdminUsers(c *gin.Context) {
	var users []AdminUserResponse

	query := `
		SELECT
				u.id,
				u.name,
				u.email,
				u.post_applied_for,
				u.invite_uid,
				EXISTS (
					SELECT 1
					FROM personal_particulars p
					WHERE p.invite_uid = u.invite_uid
				) AS submitted
			FROM users u
			ORDER BY u.created_at DESC

	`

	if err := database.DB.Raw(query).Scan(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, users)
}
