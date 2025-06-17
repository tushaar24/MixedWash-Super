package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tushaar24/mixedWash-backend/orders/controller"
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
}



