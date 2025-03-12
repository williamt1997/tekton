package router

import (
	"homers-backend/controllers"
	"homers-backend/database"
	"homers-backend/middleware"
	"log"

	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()

	router.Use(middleware.CORSMiddleware())

	addView := controllers.AddController{Database: database.DB}
	router.PUT("/add", addView.AddPageView)

	getView := controllers.GetController{Database: database.DB}
	router.GET("/get", getView.GetPageView)

	if err := router.Run("0.0.0.0:443"); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
