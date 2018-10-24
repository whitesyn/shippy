package main

import (
	"log"

	micro "github.com/micro/go-micro"
	pb "github.com/whitesyn/shippy/user-service/proto/user"
)

func main() {
	db, err := CreateConnection()
	if err != nil {
		log.Panicf("Could not connect to database - %v", err)
	}
	defer db.Close()

	repo := &UsersRepository{db}

	// Create a new service
	srv := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
	)

	// Init will parse the command line flags
	srv.Init()

	// Register handler
	pb.RegisterUserServiceHandler(srv.Server(), &handler{repo})

	if err := srv.Run(); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}