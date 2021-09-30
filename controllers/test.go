package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type TestController struct{}

func (h TestController) Status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "alive"})
}
