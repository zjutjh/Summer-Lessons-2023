package router

import (
	"CMS/app/controllers/contactController"
	usercontroller "CMS/app/controllers/userController"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	const pre = "/api"

	api := r.Group(pre)
	{
		api.POST("/login", usercontroller.Login)
		api.POST("/register", usercontroller.Register)

		contact := api.Group("/contact")
		{
			contact.GET("", contactController.GetContact)
			contact.POST("", contactController.CreateContact)
			contact.PUT("", contactController.UpdateContact)
			contact.DELETE("", contactController.DeleteContact)
		}
	}
}
