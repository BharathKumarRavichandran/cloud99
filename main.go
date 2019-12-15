package main

import (
	"github.com/BharathKumarRavichandran/cloud99/db"
	"github.com/BharathKumarRavichandran/cloud99/utils"

	indexRouter "github.com/BharathKumarRavichandran/cloud99/routers"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func initDB(config *utils.Config) {
	// Connect to database
	if err := db.Open(config); err != nil {
		panic("Could not connect to database")
	}
	//defer db.Close()
}

func RealMain() {
	config := utils.GetConfiguration()
	utils.Init(config)
	initDB(config)

	// Read env file
	err := godotenv.Load()
	if err != nil {
		utils.Logger.Fatal("Error loading .env file")
	}

	// Configure router
	router := gin.Default()

	// Configure session store
	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{Path: "/", HttpOnly: true, MaxAge: 7 * 86400})
	router.Use(sessions.Sessions("cloud99", store))

	// Configure routes; Redirect all routes to indexRouter
	indexRouter.Routes(router)

	// Start router and serve application
	utils.Logger.Infof("Listening and serving HTTP on %s", string(config.ServerPort))
	utils.Logger.Fatal(router.Run(config.ServerPort))
}

func main() {
	for {
		RealMain()
	}
}
