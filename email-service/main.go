package main

import (
	"encoding/json"
	"log"

	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	pb "github.com/whitesyn/shippy/user-service/proto/user"
)

const topic = "user.created"

func main() {
	// Create a new service
	srv := micro.NewService(
		micro.Name("go.micro.srv.email"),
		micro.Version("latest"),
	)

	// Init will parse the command line flags
	srv.Init()

	// Get instance of the broker using our defaults
	pubsub := srv.Server().Options().Broker
	if err := pubsub.Connect(); err != nil {
		log.Fatal(err)
	}

	_, err := pubsub.Subscribe(topic, func(p broker.Publication) error {
		var user *pb.User
		if err := json.Unmarshal(p.Message().Body, &user); err != nil {
			return err
		}
		log.Printf("User created: %v", user)
		go sendEmail(user)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	if err := srv.Run(); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func sendEmail(user *pb.User) error {
	log.Println("Sending email to:", user.Name)
	return nil
}
