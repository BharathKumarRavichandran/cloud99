package router

import (
	controllers "github.com/BharathKumarRavichandran/cloud99/controllers"
	authControllers "github.com/BharathKumarRavichandran/cloud99/controllers/auth"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {

	userAuthController := new(authControllers.UserController)
	dropletController := new(controllers.DropletController)

	auth := router.Group("/auth")
	userAuth := auth.Group("user")
	{
		userAuth.POST("/login", userAuthController.LoginUser)
		userAuth.POST("/register", userAuthController.RegisterUser)
		userAuth.POST("/logout", userAuthController.LogoutUser)
	}

	droplet := router.Group("/droplet")
	droplet.Use()
	{
		droplet.POST("/create", dropletController.CreateDroplet)
		droplet.POST("/delete", dropletController.DeleteDroplet)
		droplet.POST("/start", dropletController.StartDroplet)
		droplet.POST("/stop", dropletController.StopDroplet)
		droplet.GET("/status/get", dropletController.GetDropletStatus)

		dropletSettings := droplet.Group("settings")
		{
			dropletSettings.GET("/get", dropletController.GetDropletSettings)
			dropletSettings.POST("/update", dropletController.UpdateDropletSettings)
		}
	}

}
