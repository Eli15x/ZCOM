package main

import (
	"github.com/joho/godotenv"
	"ZCOM/src/model"
	"encoding/json"
	"context"
	"github.com/pkg/errors"
	"fmt"
	"time"

	"ZCOM/src/client"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/fatih/structs"
	"go.mongodb.org/mongo-driver/bson"
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

	c.SubscribeTopics([]string{"createUser", "editUser", "deleteUser", "createProduct","editProduct", "deleteProduct" ,"^aRegex.*[Tt]opic"}, nil)
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
					err := createUser(context.Background(), result)
					if err != nil {
						fmt.Println("Create user: problem to insert into MongoDB")
					}
				case "editUser":
					var result model.User
					json.Unmarshal([]byte(msg.Value), &result)
					err := editUser(context.Background(), result)
					if err != nil {
						fmt.Println("Create user: problem to insert into MongoDB")
					}
				case "deleteUser":
					var result model.User
					json.Unmarshal([]byte(msg.Value), &result)
					err := deleteUser(context.Background(), result)
					if err != nil {
						fmt.Println("Create user: problem to insert into MongoDB")
					}
				case "createProduct":
					var result model.Product
					json.Unmarshal([]byte(msg.Value), &result)
					err := createProduct(context.Background(), result)
					if err != nil {
						fmt.Println("Create user: problem to insert into MongoDB")
					}
				case "editProduct":
					var result model.Product
					json.Unmarshal([]byte(msg.Value), &result)
					err := editProduct(context.Background(), result)
					if err != nil {
						fmt.Println("Create user: problem to insert into MongoDB")
					}
				case "deleteProduct":
					var result model.Product
					json.Unmarshal([]byte(msg.Value), &result)
					err := deleteProduct(context.Background(), result)
					if err != nil {
						fmt.Println("Create user: problem to insert into MongoDB")
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


func createUser(ctx context.Context,user model.User) error{
	userInsert:= structs.Map(user)
	_, err := client.GetInstance().Insert(ctx, "user", userInsert)
	if err != nil {
		return errors.New("Create user: problem to insert into MongoDB")
	}
	return nil
}

func editUser(ctx context.Context,user model.User) error {
	userUpdate:= structs.Map(user)
	userId := map[string]interface{}{"UserId": user.UserId}
	change := bson.M{"$set": userUpdate}

	_, err := client.GetInstance().UpdateOne(ctx, "user", userId, change)
	if err != nil {
		return errors.New("Edit User: problem to update into MongoDB")
	}
	return nil
}

func deleteUser(ctx context.Context,user model.User) error {
	//casos de duplicidade ajustar
	userId := map[string]interface{}{"UserId": user.UserId}

	err := client.GetInstance().Remove(ctx, "user", userId)
	if err != nil {
		return errors.New("Delete User: problem to delete into MongoDB")
	}

	return nil
}

func createProduct(ctx context.Context,product model.Product) error {
	//casos de duplicidade ajustar
	productInsert := structs.Map(product)

	_, err := client.GetInstance().Insert(ctx, "product", productInsert)
	if err != nil {
		return errors.New("Product: problem to insert into MongoDB")
	}

	return nil
}

func editProduct(ctx context.Context,product model.Product) error {
	//casos de duplicidade ajustar
	productUpdate:= structs.Map(product)
	CODIGO_CEST := map[string]interface{}{"CODIGO_CEST": product.CODIGO_CEST}
	change := bson.M{"$set": productUpdate}

	_, err := client.GetInstance().UpdateOne(ctx, "product", CODIGO_CEST, change)
	if err != nil {
		return errors.New("Edit product: problem to update into MongoDB")
	}
	return nil
}

func deleteProduct(ctx context.Context,product model.Product) error {
	//casos de duplicidade ajustar
	CODIGO_CEST := map[string]interface{}{"CODIGO_CEST": product.CODIGO_CEST}

	err := client.GetInstance().Remove(ctx, "product", CODIGO_CEST)
	if err != nil {
		return errors.New("Delete Product: problem to delete into MongoDB")
	}

	return nil
}