package controllers

import (
	"net/http"

	"backend/utils"
	"github.com/gin-gonic/gin"
)

func GeneratePersonalParticularsPDF(c *gin.Context) {
	inviteUid := c.Param("inviteUid") // ✅ get from URL

	if inviteUid == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "inviteUid is required",
		})
		return
	}

	pdfBytes, err := utils.GeneratePersonalParticularsPDF(inviteUid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Disposition", "attachment; filename=personal_particulars.pdf")
	c.Data(http.StatusOK, "application/pdf", pdfBytes)
}
