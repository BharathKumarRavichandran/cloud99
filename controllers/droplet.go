package controllers

import (
	"net/http"

	"github.com/BharathKumarRavichandran/cloud99/db"
	"github.com/BharathKumarRavichandran/cloud99/models"

	"github.com/google/uuid"

	"github.com/BharathKumarRavichandran/cloud99/helpers"
	"github.com/BharathKumarRavichandran/cloud99/utils"
	"github.com/gin-gonic/gin"
)

type DropletController struct{}

func (ctrl DropletController) CreateDroplet(c *gin.Context) {

	userID := 1
	image := "ubuntu:19.04"
	var user models.User
	db.DB.First(&user, userID)

	containerID, err := helpers.CreateContainer(image)
	if err != nil {
		responseMessage := "Unable to create container"
		utils.Logger.Infof("UserID %d: %s: %s", userID, responseMessage, containerID)
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status_code": http.StatusNotAcceptable,
			"message":     responseMessage,
		})
		c.Abort()
		return
	}

	// Save droplet details in database
	droplet := models.Droplet{
		DropletUID:  uuid.New().String(),
		UserID:      user,
		Ram:         "1GB",
		DiskStorage: "1GB",
		OS:          "ubuntu:19.04",
		Status:      "1",
	}
	db.DB.Create(&droplet)

	// Save created container details in database
	container := models.Container{
		ContainerID: containerID,
		Image:       image,
		DropletID:   droplet,
	}
	db.DB.Create(&container)

	_, err = helpers.StartContainer(containerID)
	if err != nil {
		responseMessage := "Unable to start container"
		utils.Logger.Infof("UserID %d: %s: %s", userID, responseMessage, containerID)
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status_code": http.StatusNotAcceptable,
			"message":     responseMessage,
		})
		c.Abort()
		return
	}

	responseMessage := "Successfully created and started droplet"
	utils.Logger.Infof("UserID %d : %s : %s", 1, responseMessage, droplet.DropletUID)
	c.JSON(http.StatusOK, gin.H{
		"status_code": http.StatusOK,
		"message":     responseMessage,
	})
	return
}

func (ctrl DropletController) StartDroplet(c *gin.Context) {

	userID := 1
	containerID := "8f58355077f3"

	_, err := helpers.StartContainer(containerID)
	if err != nil {
		responseMessage := "Unable to start container"
		utils.Logger.Infof("UserID %d: %s: %s", userID, responseMessage, containerID)
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status_code": http.StatusNotAcceptable,
			"message":     responseMessage,
		})
		c.Abort()
		return
	}

	responseMessage := "Successfully started droplet"
	utils.Logger.Infof("UserID %d : %s : %s", 1, responseMessage, containerID)
	c.JSON(http.StatusOK, gin.H{
		"status_code": http.StatusOK,
		"message":     responseMessage,
	})
	return
}

func (ctrl DropletController) StopDroplet(c *gin.Context) {

	responseMessage := "Successfully stopped droplet"
	utils.Logger.Infof("UserID %d : %s : %s", 1, responseMessage, "1")
	c.JSON(http.StatusOK, gin.H{
		"status_code": http.StatusOK,
		"message":     responseMessage,
	})
	return
}

func (ctrl DropletController) DeleteDroplet(c *gin.Context) {

	responseMessage := "Successfully deleted droplet"
	utils.Logger.Infof("UserID %d : %s : %s", 1, responseMessage, "1")
	c.JSON(http.StatusOK, gin.H{
		"status_code": http.StatusOK,
		"message":     responseMessage,
	})
	return
}

func (ctrl DropletController) GetDropletStatus(c *gin.Context) {

	responseMessage := "Successfully retrieved status for droplet"
	utils.Logger.Infof("UserID %d : %s : %s", 1, responseMessage, "1")
	c.JSON(http.StatusOK, gin.H{
		"status_code": http.StatusOK,
		"message":     responseMessage,
	})
	return
}

func (ctrl DropletController) GetDropletSettings(c *gin.Context) {

	responseMessage := "Successfully retrieved settings for droplet"
	utils.Logger.Infof("UserID %d : %s : %s", 1, responseMessage, "1")
	c.JSON(http.StatusOK, gin.H{
		"status_code": http.StatusOK,
		"message":     responseMessage,
	})
	return
}

func (ctrl DropletController) UpdateDropletSettings(c *gin.Context) {

	responseMessage := "Successfully updated settings for droplet"
	utils.Logger.Infof("UserID %d : %s : %s", 1, responseMessage, "1")
	c.JSON(http.StatusOK, gin.H{
		"status_code": http.StatusOK,
		"message":     responseMessage,
	})
	return
}
