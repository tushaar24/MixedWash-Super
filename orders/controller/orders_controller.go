package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tushaar24/mixedWash-backend/orders/services"
)


func FetchAllOrders(context *gin.Context){
	orders, err := services.FetchAllOrders()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
	}
	context.JSON(http.StatusOK, orders)
}

func GetOrdersByUserId(context *gin.Context, userId uuid.UUID){
	orders, err := services.GetAllOrderOfUser(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
	}
	context.JSON(http.StatusOK, orders)
}





