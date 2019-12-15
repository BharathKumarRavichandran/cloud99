package auth

import (
	"fmt"
	"net/http"

	authForms "github.com/BharathKumarRavichandran/cloud99/forms/auth"
	"github.com/BharathKumarRavichandran/cloud99/models"
	"github.com/BharathKumarRavichandran/cloud99/utils"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	_ "github.com/sirupsen/logrus"
)

type UserController struct{}

var userModel = new(models.UserModel)

func (ctrl UserController) LoginUser(c *gin.Context) {
	var loginForm authForms.LoginForm

	if err := c.Bind(&loginForm); err != nil {
		utils.Logger.Infof("UserEmail %s : %s", loginForm.Email, err.Error())
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status_code": http.StatusNotAcceptable,
			"message":     err.Error(),
		})
		c.Abort()
		return
	}

	user, err := userModel.Login(loginForm)
	if err != nil {
		utils.Logger.Infof("UserEmail %s : %s", loginForm.Email, err.Error())
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status_code": http.StatusNotAcceptable,
			"message":     err.Error(),
		})
		c.Abort()
		return
	}

	// Save session for registered user
	session := sessions.Default(c)
	logged_user := struct {
		UserID uint
		Name   string
		Email  string
	}{
		UserID: user.UserID,
		Name:   user.Name,
		Email:  user.Email,
	}
	session.Set("user", logged_user)
	session.Save()

	fmt.Println(sessions.Default(c).Get("user"))

	utils.Logger.Infof("UserID %d : Successfully logged in", user.UserID)
	c.JSON(http.StatusOK, gin.H{
		"status_code": http.StatusOK,
		"message":     "Successfully logged in",
	})
	return
}

func (ctrl UserController) RegisterUser(c *gin.Context) {
	var registerForm authForms.RegisterForm

	if err := c.Bind(&registerForm); err != nil {
		utils.Logger.Infof("UserEmail %s : %s", registerForm.Email, err.Error())
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status_code": http.StatusNotAcceptable,
			"message":     err.Error(),
		})
		c.Abort()
		return
	}

	user, err := userModel.Register(registerForm)
	if err != nil {
		utils.Logger.Infof("UserEmail %s : %s", registerForm.Email, err.Error())
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status_code": http.StatusNotAcceptable,
			"message":     err.Error(),
		})
		c.Abort()
		return
	}

	if user.UserID == 0 {
		utils.Logger.Infof("UserID %d : Could not register user", user.UserID)
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status_code": http.StatusNotAcceptable,
			"message":     "Could not register user",
		})
		c.Abort()
		return
	}

	// Save session for registered user
	session := sessions.Default(c)
	logged_user := struct {
		UserID uint
		Name   string
		Email  string
	}{
		UserID: user.UserID,
		Name:   user.Name,
		Email:  user.Email,
	}
	session.Set("user", logged_user)
	session.Save()

	utils.Logger.Infof("UserID %d : Successfully registered user", user.UserID)
	c.JSON(http.StatusOK, gin.H{
		"status_code": http.StatusOK,
		"message":     "Successfully registered user",
	})
	return
}

func (ctrl UserController) LogoutUser(c *gin.Context) {
	session := sessions.Default(c)

	userEmail := session.Get("user")

	// Destroy user session
	session.Delete("user")
	session.Save()

	utils.Logger.Infof("UserID %s : Successfully logged out", userEmail)
	c.JSON(http.StatusOK, gin.H{
		"status_code": http.StatusOK,
		"message":     "Successfully logged out",
	})
	return
}
