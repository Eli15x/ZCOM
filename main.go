package main

import (
	//"time"
	//"context"

	"context"
	"fmt"
	"time"

	"github.com/Eli15x/ZCOM/src/client"
	"github.com/Eli15x/ZCOM/src/handler"
	"github.com/bugsnag/bugsnag-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	//bugsnag configure
	bugsnag.Configure(bugsnag.Configuration{
		APIKey:       "3ecac0ed23b7b1f4b863073135c602b8",
		ReleaseStage: "production",
		// The import paths for the Go packages containing your source files
		ProjectPackages: []string{"main", "github.com/org/myapp"},
		// more configuration options
	})

	bugsnag.Notify(fmt.Errorf("Test error"))

	//Context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	//Connection to Mongo
	if err := client.GetInstance().Initialize(ctx); err != nil {
		bugsnag.Notify(fmt.Errorf("[MONGO DB - MovieWorkNow] Could not resolve Data access layer. Error:"))
	}

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"*"}
	config.AllowMethods = []string{"*"}
	config.AllowCredentials = true

	router.Use(cors.New(config))

	/*router.POST("/login", handlers.ValidateUser)
	router.POST("/cadastro", handlers.CreateUser)
	router.POST("/user/edit", handlers.EditUser)
	router.POST("/user/delete", handlers.DeleteUser)
	router.POST("/user/getByUserId", handlers.DeleteUser)
	router.POST("/user/getByName", handlers.DeleteUser)
	router.POST("/user/getByFunction", handlers.DeleteUser)
	router.POST("/user/getByAcess", handlers.DeleteUser)

	router.POST("/product/create", handlers.CreateWriter)
	router.GET("/product/getByName", handlers.GetInformationWriter)
	router.GET("/product/getByCode", handlers.DeleteWriter)
	router.GET("/product/getByBarCode", handlers.DeleteWriter)
	router.GET("/product/getAll", handlers.GetInformationWriters)
	router.POST("/product/edit", handlers.EditWriter)
	router.DELETE("/product/delete", handlers.DeleteWriter)
	
	//router.POST("/sale/send", handlers.CreateBook)
	//.POST("/sales/send", handlers.CreateBook)*/


	router.Run(":1323")
}