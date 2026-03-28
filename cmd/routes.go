package main

import "github.com/gin-gonic/gin"

func setUpRoutes(router *gin.Engine, handler *Handler) {
	router.GET("/", handler.ServeNewOrderForm)
	router.POST("/new-service", handler.HandleNewVehiclePost)
	router.GET("/customers", handler.HandleGetCustomers)

	router.Static("/static", "./template/static")
}
