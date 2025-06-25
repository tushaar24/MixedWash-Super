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

func UpdateDriver(context *gin.Context, taskId string, driverId string) {

	err := services.UpdateDriver(driverId, taskId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
	}
	context.JSON(http.StatusOK, "")
}

func GetDrivers(context *gin.Context) {

	drivers, err := services.GetDrivers()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
	}
	context.JSON(http.StatusOK, drivers)
}

func UpdateDriverTaskStatus(status string, taskId string, context *gin.Context) {

	err := services.UpdateDriverTaskStatus(taskId, status)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
	}

	context.JSON(http.StatusOK, "")
}
