package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tushaar24/mixedWash-backend/drivers/services"
)

func GetTodaysTasks(context *gin.Context) {

	todayTasks, err := services.GetTodayTask()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
	}

	context.JSON(http.StatusOK, todayTasks)

}
