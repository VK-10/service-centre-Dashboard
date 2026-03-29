package main

import "github.com/gin-gonic/gin"

func setUpRoutes(router *gin.Engine, handler *Handler) {
	router.GET("/", handler.ServeNewVehicleForm)
	router.POST("/new-vehicle", handler.HandleNewVehiclePost)
	router.GET("/customers/:id", handler.serveCustomer)

	router.Static("/static", "./templates/static")
}
