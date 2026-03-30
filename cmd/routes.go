package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func setUpRoutes(router *gin.Engine, handler *Handler, store sessions.Store) {

	router.Use(sessions.Sessions("service-tracker-session", store))

	router.GET("/", handler.ServeNewVehicleForm)
	router.POST("/new-vehicle", handler.HandleNewVehiclePost)
	router.GET("/customers/:id", handler.serveCustomer)
	router.GET("/notifications", handler.NotificationHandler)

	router.GET("/login", handler.HandleLoginGet)
	router.POST("/login", handler.HandleLoginPost)
	router.GET("/logout", handler.HandleLogout)

	admin := router.Group("/admin")
	admin.Use(handler.AuthMiddleware())
	{
		admin.GET("", handler.ServeAdminDashboard)
		admin.POST("/vehicles/:id/status", handler.HandleOrderPut)
		admin.POST("/vehicles/:id", handler.HandleVehicleDelete)
		admin.GET("/notifications", handler.adminNotificationHandler)
	}

	router.Static("/static", "./templates/static")
}
