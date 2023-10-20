package main

import (
	"context"
	"github.com/pkg/errors"
	"fmt"
	"time"

	"github.com/Eli15x/ZCOM/src/client"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/fatih/structs"
	//"go.mongodb.org/mongo-driver/bson"
)

func main() {

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{"myTopic", "^aRegex.*[Tt]opic"}, nil)

	// A signal handler or similar could be used to set this to false to break the loop.
	run := true

	for run {
		msg, err := c.ReadMessage(time.Second)
		fmt.Println(msg)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
			//switch msg.TopicPartition  {
			/*case "createUser":
				err := createUser(ctx, msg.Value)
				if err!= nil {
					
				}
			case "editUser":
				fmt.Println("Linux.")
			default:
				// freebsd, openbsd,
				// plan9, windows...
				fmt.Printf("%s.\n", os)
			}*/
			fmt.Println(msg)



			//case topic is x,y,z...
		} else if !err.(kafka.Error).IsTimeout() {
			// The client will automatically try to recover from all errors.
			// Timeout is not considered an error because it is raised by
			// ReadMessage in absence of messages.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}

	c.Close()
}


func CreateUser(ctx context.Context,userJson []byte) error{
	userInsert := structs.Map(userJson)

	_, err := client.GetInstance().Insert(ctx, "user", userInsert)
	if err != nil {
		return errors.New("Create user: problem to insert into MongoDB")
	}
	return nil
}