package main

import (
	"log"
	"os"

	"github.com/YoniMk/ecommerce-yt/controllers"
	"github.com/YoniMk/ecommerce-yt/database"
	"github.com/YoniMk/ecommerce-yt/middleware"
	"github.com/YoniMk/ecommerce-yt/routes"
	"github.com/gin-gonic/gin"
)

func main(){
	port := os.Getenv("PORT")
	if port ==  ""{
		port = "8000"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/addtocart", app.AddToCart())
	router.GET("removeitem", app.RemoveItem())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instanbuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))
}