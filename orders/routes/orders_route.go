package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tushaar24/mixedWash-backend/orders/controller"
	"github.com/tushaar24/mixedWash-backend/orders/services/models"
)

func RegisterRoute(router *gin.Engine) {

	router.GET("/orders", controller.FetchAllOrders)

	router.GET("/users/:user_id/orders", func(ctx *gin.Context) {

		userIdStr := ctx.Param("user_id")
		userId, err := uuid.Parse(userIdStr)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error":"Invalid UUID format"})
			return
		}

		controller.GetOrdersByUserId(ctx, userId)
	})

	router.POST("/customer/create", func(ctx *gin.Context) {

		var customer models.CustomerCreationDTO

		if err := ctx.ShouldBindJSON(&customer); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		controller.CreateCustomer(ctx,customer)
	})

	router.POST("/order/create", func(ctx *gin.Context) {

		var order models.OrderCreationDTO

		if err := ctx.ShouldBindJSON(&order); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		controller.CreateOrderAdmin(ctx, order)
	})

	router.GET("/user/getUser/:phone_number", func(ctx *gin.Context) {
		phoneNumber := ctx.Param("phone_number")
		controller.GetCustomerByPhone(ctx,phoneNumber)
	})

	router.GET("user/addresses/getAddress", func(ctx *gin.Context) {
		userId := ctx.Query("user_id")
		controller.GetCustomerAddresses(ctx, userId)
	})


	router.POST("addresses/admin/add", func(ctx *gin.Context) {

		var address models.AddAddressAdminDTO

		if err := ctx.ShouldBindJSON(&address); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		controller.AddAddressAdmin(ctx, address)

	})

	router.GET("admin/order/screenResponse", controller.GetAdminOrderCreationScreen)
}



