package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateIDCard(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "CreateIDCard not implemented"})
}

func GetIDCardByID(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "GetIDCardByID not implemented"})
}

func GetIDCardByUser(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "GetIDCardByUser not implemented"})
}

func UpdateIDCard(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "UpdateIDCard not implemented"})
}
