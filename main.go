package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tushaar24/mixedWash-backend/orders/routes"
)

func main(){
	router := gin.Default()
	routes.RegisterRoute(router)
	router.Run(":8008")
}
