package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheck(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "Printing Management System API is running",
	})
}
