package main

import (
	kafka "ZCOM/src/client/kafka"
	//"time"
	//"context"

	"context"
	"fmt"
	"time"

	"ZCOM/src/client"
	handlers "ZCOM/src/handler"

	"github.com/bugsnag/bugsnag-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Errorf("Error loading .env file")
	}

	//Context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	//Connection to Mongo
	if err := client.GetInstance().Initialize(ctx); err != nil {
		fmt.Errorf("mongo off")
		bugsnag.Notify(fmt.Errorf("[MONGO DB - ZCOM] Could not resolve Data access layer. Error:"))
	}

	if err := kafka.GetInstanceKafka().Initialize(); err != nil {
		fmt.Println(err)
		fmt.Errorf("Error initialize kafka Producer")
	}

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"*"}
	config.AllowMethods = []string{"*"}
	config.AllowCredentials = true

	router.Use(cors.New(config))

	router.POST("/login", handlers.ValidateUser)
	router.POST("/cadastro", handlers.CreateUser)
	router.PUT("/user/edit", handlers.EditUser)
	router.DELETE("/user/delete", handlers.DeleteUser)
	router.GET("/user/getByUserId", handlers.GetInformationByUserId)
	router.GET("/user/getByName", handlers.GetUserByName)
	router.GET("/user/getByAcess", handlers.GetUsersByAcess)
	router.GET("/users", handlers.GetUsers)
	router.GET("/user/save", handlers.SaveUser)

	router.POST("/product/create", handlers.CreateProduct)
	router.GET("/product/getByName", handlers.GetProductByName)
	router.GET("/product/getByGTIN", handlers.GetProduct)
	router.GET("/product/getAll", handlers.GetProducts)
	router.PUT("/product/edit", handlers.EditProduct)
	router.DELETE("/product/delete", handlers.DeleteProduct)
	router.GET("/product/save", handlers.SaveProduct) //rota para salvar produtos

	router.GET("/marca", handlers.GetMarca)
	router.GET("/marcas", handlers.GetMarcas)

	router.GET("/grupo", handlers.GetGrupo)
	router.GET("/grupos", handlers.GetGrupos)

	router.GET("/unidade", handlers.GetUnidade)
	router.GET("/unidades", handlers.GetUnidades)

	router.POST("/saleXml", handlers.CreateSaleXml)
	router.POST("/saleSend", handlers.SendSale) //mandar os dados da venda aqui
	//router.POST("/getUpdates", handlers.GetUpdated) //pegar dados produto aqui que deu updated.

	router.Run(":1323")
}
