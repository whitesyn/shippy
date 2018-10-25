package main

import (
	"context"
	"log"

	micro "github.com/micro/go-micro"
	pb "github.com/whitesyn/shippy/user-service/proto/user"
)

const topic = "user.created"

// Subscriber structure
type Subscriber struct{}

// Process topic message
func (sub *Subscriber) Process(ctx context.Context, user *pb.User) error {
	log.Println("Picked up a new message")
	log.Println("Sending email to:", user.Name)
	return nil
}

func main() {
	// Create a new service
	srv := micro.NewService(
		micro.Name("go.micro.srv.email"),
		micro.Version("latest"),
	)

	// Init will parse the command line flags
	srv.Init()

	// Subscribe on "user.created" topic
	micro.RegisterSubscriber(topic, srv.Server(), new(Subscriber))

	if err := srv.Run(); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func sendEmail(user *pb.User) error {
	log.Println("Sending email to:", user.Name)
	return nil
}
