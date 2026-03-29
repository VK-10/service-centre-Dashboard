package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := GetSessionString(c, "userID")

		if userId == "" {
			c.Redirect(http.StatusSeeOther, "/login")
			c.Abort()
			return

		}

		user, err := h.users.GetUserByID(userId)
		if err != nil {
			ClearSession(c)
			c.Redirect(http.StatusSeeOther, "/login")
			c.Abort()
			return
		}

		c.Next()
	}
}
