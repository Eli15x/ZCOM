package main

import (
	"github.com/joho/godotenv"
	"ZCOM/src/model"
	"ZCOM/src/service"
	"encoding/json"
	"context"
	"fmt"
	"time"

	"ZCOM/src/client"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {

	err := godotenv.Load("../.env")
    if err != nil {
        fmt.Errorf("Error loading .env file")
    }

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{"createUser", "editUser", "deleteUser", "createProduct","editProduct", "deleteProduct" ,"^aRegex.*[Tt]opic", "createSaleXML"}, nil)
	run := true

	for run {
		if err := client.GetInstance().Initialize(context.Background()); err != nil {
			fmt.Println("MongoDB is offline")
		} else {
			msg, err := c.ReadMessage(time.Second)
			fmt.Println(msg)
			
			if err == nil {
				fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
				switch *msg.TopicPartition.Topic  {
				case "createUser":	
					var result model.User
					json.Unmarshal([]byte(msg.Value), &result)
					err := service.GetInstanceUser().CreateUser(context.Background(), result)
					if err != nil {
						fmt.Println("Create user: problem to insert into MongoDB")
					}
				case "editUser":
					var result model.User
					json.Unmarshal([]byte(msg.Value), &result)
					err := service.GetInstanceUser().EditUser(context.Background(), result)
					if err != nil {
						fmt.Println("Create user: problem to insert into MongoDB")
					}
				case "deleteUser":
					var result model.User
					json.Unmarshal([]byte(msg.Value), &result)
					err := service.GetInstanceUser().DeleteUser(context.Background(), result)
					if err != nil {
						fmt.Println("Create user: problem to insert into MongoDB")
					}
				case "createProduct":
					var result model.Product
					json.Unmarshal([]byte(msg.Value), &result)
					err := service.GetInstanceProduct().CreateProduct(context.Background(), result)
					if err != nil {
						fmt.Println("Create Product: problem to insert into MongoDB")
					}
				case "editProduct":
					var result model.Product
					json.Unmarshal([]byte(msg.Value), &result)
					err := service.GetInstanceProduct().EditProduct(context.Background(), result)
					if err != nil {
						fmt.Println("Create Product: problem to insert into MongoDB")
					}
				case "deleteProduct":
					var result model.Product
					json.Unmarshal([]byte(msg.Value), &result)
					err := service.GetInstanceProduct().DeleteProduct(context.Background(), result)
					if err != nil {
						fmt.Println("Create Product: problem to insert into MongoDB")
					}
				case "createSaleXML":
					var result model.SaleXML
					json.Unmarshal([]byte(msg.Value), &result)
					err := service.GetInstanceSale().SaveSaleXML(context.Background(), result)
					if err != nil {
						fmt.Println("Create saleXML: problem to put on file")
					}
				default:
					// freebsd, openbsd,
					// plan9, windows...
					fmt.Printf("%s.\n", "not exists this topic")
				}

			} else if !err.(kafka.Error).IsTimeout() {
				fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			}
		}
	}

	c.Close()
}

