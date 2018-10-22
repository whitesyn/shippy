package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"

	"golang.org/x/net/context"

	pb "github.com/whitesyn/shippy/consignment-service/proto/consignment"
)

const (
	defaultFilename = "consignment.json"
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

	client := pb.NewShippingServiceClient("go.micro.srv.consignment", microclient.DefaultClient)

	// Parse JSON-file to Protobuf structure
	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}
	consignment, err := parseFile(file)
	if err != nil {
		log.Fatalf("Can not parse consignment file: %v", err)
	}

	// Contact the server and print out its response.
	r, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("Can not create consignment: %v", err)
	}
	log.Printf("Created: %t", r.Created)

	// Get all created consignments
	getAll, err := client.GetConsignments(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list consignments: %v", err)
	}
	for _, v := range getAll.Consignments {
		log.Println(v)
	}
}
