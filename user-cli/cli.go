package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/micro/cli"
	micro "github.com/micro/go-micro"
	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"

	pb "github.com/whitesyn/shippy/user-service/proto/user"
)

const (
	defaultFilename = "user.json"
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

	client := pb.NewUserServiceClient("go.micro.srv.user", microclient.DefaultClient)

	// Define our flags
	service := micro.NewService(
		micro.Flags(
			cli.StringFlag{
				Name:  "name",
				Usage: "You full name",
			},
			cli.StringFlag{
				Name:  "email",
				Usage: "Your email",
			},
			cli.StringFlag{
				Name:  "password",
				Usage: "Your password",
			},
			cli.StringFlag{
				Name:  "company",
				Usage: "Your company",
			},
		),
	)

	// Start as service
	service.Init(
		micro.Action(func(c *cli.Context) {
			name := c.String("name")
			email := c.String("email")
			password := c.String("password")
			company := c.String("company")

			// Call our user service
			r, err := client.Create(context.TODO(), &pb.User{
				Name:     name,
				Email:    email,
				Password: password,
				Company:  company,
			})
			if err != nil {
				log.Fatalf("Could not create: %v", err)
			}
			log.Printf("Created: %s", r.User.Id)

			authResponse, err := client.Auth(context.TODO(), &pb.User{
				Email:    email,
				Password: password,
			})
			if err != nil {
				log.Panicf("Could not auth user: %v", err)
			}
			log.Printf("Your access token is: %s", authResponse.Token)

			os.Exit(0)
		}),
	)

	// Run the server
	if err := service.Run(); err != nil {
		log.Println(err)
	}
}
