package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tushaar24/mixedWash-backend/orders/services"
	"github.com/tushaar24/mixedWash-backend/orders/services/models"
	"net/http"
)

func FetchAllOrders(context *gin.Context) {

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

func GetOrdersByUserId(context *gin.Context, userId uuid.UUID) {
	orders, err := services.GetAllOrderOfUser(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
	}

	context.JSON(http.StatusOK, orders)
}

func CreateCustomer(context *gin.Context, customer models.CustomerCreationDTO) {
	id, err := services.CreateCustomer(customer)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
	}

	context.JSON(http.StatusOK, id)
}

func CreateOrderAdmin(context *gin.Context, order models.OrderCreationDTO) {
	err := services.CreateOrderAdmin(order)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
	}

	context.JSON(http.StatusOK, "")
}

func GetCustomerByPhone(context *gin.Context, phoneNumber string) {

	customer, tempCustomer, err := services.GetCustomerByPhoneNo(phoneNumber)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
	}

	if customer == nil {
		context.JSON(http.StatusOK, tempCustomer)
		return
	}

	context.JSON(http.StatusOK, customer)
}

func GetCustomerAddresses(context *gin.Context, userId string) {

	addresses, err := services.GetCustomerAddresses(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
	}

	var consiseAddresses []models.CustomerAddressesByUserIdModel

	for _, address := range addresses {
		consiseAddresses = append(consiseAddresses, address.ToModel())
	}

	context.JSON(http.StatusOK, consiseAddresses)
}

func AddAddressAdmin(context *gin.Context, address models.AddAddressAdminDTO) {

	err := services.AddAddressAdmin(address)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
	}

	context.JSON(http.StatusOK, "")

}

func GetAdminOrderCreationScreen (context *gin.Context) {

	screenDTO, err := services.GetAdminOrderCreationScreen()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
	}

	context.JSON(http.StatusOK, screenDTO)
}
