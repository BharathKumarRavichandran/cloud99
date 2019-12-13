package controllers

import (
	"net/http"

	"github.com/BharathKumarRavichandran/cloud99/utils"
	"github.com/gin-gonic/gin"
)

type DropletController struct{}

func (ctrl DropletController) CreateDroplet(c *gin.Context) {

	responseMessage := "Successfully created droplet"
	utils.Logger.Infof("UserID %d : %s : %d", 1, responseMessage, 1)
	c.JSON(http.StatusOK, gin.H{
		"status_code": http.StatusOK,
		"message":     responseMessage,
	})
	return
}

func (ctrl DropletController) StartDroplet(c *gin.Context) {

	responseMessage := "Successfully started droplet"
	utils.Logger.Infof("UserID %d : %s : %d", 1, responseMessage, 1)
	c.JSON(http.StatusOK, gin.H{
		"status_code": http.StatusOK,
		"message":     responseMessage,
	})
	return
}

func (ctrl DropletController) StopDroplet(c *gin.Context) {

	responseMessage := "Successfully stopped droplet"
	utils.Logger.Infof("UserID %d : %s : %d", 1, responseMessage, 1)
	c.JSON(http.StatusOK, gin.H{
		"status_code": http.StatusOK,
		"message":     responseMessage,
	})
	return
}

func (ctrl DropletController) DeleteDroplet(c *gin.Context) {

	responseMessage := "Successfully deleted droplet"
	utils.Logger.Infof("UserID %d : %s : %d", 1, responseMessage, 1)
	c.JSON(http.StatusOK, gin.H{
		"status_code": http.StatusOK,
		"message":     responseMessage,
	})
	return
}

func (ctrl DropletController) GetDropletStatus(c *gin.Context) {

	responseMessage := "Successfully retrieved status for droplet"
	utils.Logger.Infof("UserID %d : %s : %d", 1, responseMessage, 1)
	c.JSON(http.StatusOK, gin.H{
		"status_code": http.StatusOK,
		"message":     responseMessage,
	})
	return
}

func (ctrl DropletController) GetDropletSettings(c *gin.Context) {

	responseMessage := "Successfully retrieved settings for droplet"
	utils.Logger.Infof("UserID %d : %s : %d", 1, responseMessage, 1)
	c.JSON(http.StatusOK, gin.H{
		"status_code": http.StatusOK,
		"message":     responseMessage,
	})
	return
}

func (ctrl DropletController) UpdateDropletSettings(c *gin.Context) {

	responseMessage := "Successfully updated settings for droplet"
	utils.Logger.Infof("UserID %d : %s : %d", 1, responseMessage, 1)
	c.JSON(http.StatusOK, gin.H{
		"status_code": http.StatusOK,
		"message":     responseMessage,
	})
	return
}
