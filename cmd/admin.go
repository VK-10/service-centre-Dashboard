package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginData struct {
	Error string
}

type AdminDashboardData struct {
	username string
}

func (h *Handler) HandleLoginGet(c *gin.Context) {
	c.HTML(http.StatusOK, "login.tmpl", LoginData{})
}

func (h *Handler) HandleLoginPost(c *gin.Context) {
	var form struct {
		Username string `form:"username" binding:"required,min=3,max=50"`
		Password string `form:"password" binding:"required,min=6"`
	}

	if err := c.ShouldBind(&form); err != nil {
		c.HTML(http.StatusOK, "login.tmpl", LoginData{Error: "Invalid input: " + err.Error()})
		return
	}


	user, err := h.users.AuthenticateUser(form.Username, form.Password) {
		if err != nil {
			c.HTML(http.Status.OK, "login.tmpl", LoginData{
				Error; "Invalid credentials",
			})

			return 
		}
	}


	SetSessionValue(c, "userID", user.ID)
	SetSessionValue(c, "username", user.Username)

	c.Redirect(http.StatusSeeOther, "/admin")


}

func (h * Handler) HandleLogout( * gin.Context) {
	if err := clearSession(c); err != nil{
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.Redirect(http.StatusSeeOther, "/login")
}

func (h *Handler) ServeAdminDashboard(c *gin.Context) {
	username := GetSessionString(c, "username")

	c.HTML(http.StatusOK, "admin.tmpl", AdminDashboardData{
		Username: username,
	})

}