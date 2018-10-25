package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/metadata"

	"golang.org/x/net/context"

	pb "github.com/whitesyn/shippy/consignment-service/proto/consignment"
)

func parseFile(file string) (*pb.Consignment, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	var consignment *pb.Consignment

	err = json.Unmarshal(data, &consignment)
	if err != nil {
		return nil, err
	}

	return consignment, nil
}

func main() {
	cmd.Init()

	if len(os.Args) < 3 {
		log.Fatal("Not enough arguments, expecing file and token.")
	}

	// Parse JSON-file to Protobuf structure
	file := os.Args[1]
	consignment, err := parseFile(file)
	if err != nil {
		log.Fatalf("Can not parse consignment file: %v", err)
	}

	token := os.Args[2]

	// Create a new context which contains our given token.
	// This same context will be passed into both the calls we make
	// to our consignment-service.
	ctx := metadata.NewContext(context.Background(), map[string]string{
		"token": token,
	})

	client := pb.NewShippingServiceClient("go.micro.srv.consignment", microclient.DefaultClient)

	// Contact the server and print out its response.
	r, err := client.CreateConsignment(ctx, consignment)
	if err != nil {
		log.Fatalf("Can not create consignment: %v", err)
	}
	log.Printf("Created: %t", r.Created)

	// Get all created consignments
	getAll, err := client.GetConsignments(ctx, &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list consignments: %v", err)
	}
	for _, v := range getAll.Consignments {
		log.Println(v)
	}
}
