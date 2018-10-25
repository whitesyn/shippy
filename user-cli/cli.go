package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"

	pb "github.com/whitesyn/shippy/user-service/proto/user"
)

func parseFile(file string) (*pb.User, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	var user *pb.User

	err = json.Unmarshal(data, &user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func main() {
	cmd.Init()

	if len(os.Args) < 2 {
		log.Fatal("Not enough arguments, expecing file")
	}
	// Parse JSON-file to Protobuf structure
	file := os.Args[1]
	user, err := parseFile(file)
	if err != nil {
		log.Fatalf("Can not parse user file: %v", err)
	}

	client := pb.NewUserServiceClient("go.micro.srv.user", microclient.DefaultClient)

	// Call our user service
	log.Printf("User: %v\n", user)
	r, err := client.Create(context.TODO(), user)
	if err != nil {
		log.Fatalf("Could not create: %v", err)
	}
	log.Printf("Created user with ID and email: %s %s\n", r.User.Id, r.User.Email)

	authResponse, err := client.Auth(context.TODO(), &pb.User{
		Email:    user.Email,
		Password: user.Password,
	})
	if err != nil {
		log.Panicf("Could not auth user: %v", err)
	}
	log.Printf("Your access token is: %s", authResponse.Token)
}
