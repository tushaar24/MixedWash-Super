package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tushaar24/mixedWash-backend/orders/services"
	"github.com/tushaar24/mixedWash-backend/orders/services/models"
)

func FetchAllOrders(context *gin.Context){
	orders, err := services.FetchAllOrders()
	var orderDashboard []models.OrderDashboardModel
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
	}
	for _, order := range orders {
		orderDashboard = append(orderDashboard, order.ConvertToOrderDashboardModel())
	}
	context.JSON(http.StatusOK, orderDashboard)
}

func GetOrdersByUserId(context *gin.Context, userId uuid.UUID){
	orders, err := services.GetAllOrderOfUser(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
	}
	context.JSON(http.StatusOK, orders)
}

func CreateCustomer(context *gin.Context, customer models.CustomerCreationDTO){
	services.CreateCustomer(customer)
	context.JSON(http.StatusOK, "")
}





