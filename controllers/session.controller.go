package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateSession(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "CreateSession not implemented"})
}

func GetSessionByID(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "GetSessionByID not implemented"})
}

func ListSessionsByUser(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "ListSessionsByUser not implemented"})
}

func DeactivateSession(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "DeactivateSession not implemented"})
}
