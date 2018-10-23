package main

import (
	"log"
	"os"

	micro "github.com/micro/go-micro"
	pb "github.com/whitesyn/shippy/vessel-service/proto/vessel"
)

const (
	defaultHost = "localhost:27017"
)

func creatDeummyData(repo Repository) {
	defer repo.Close()
	vessels := []*pb.Vessel{
		&pb.Vessel{Id: "vessel001", Name: "Kane's Salty Secret", MaxWeight: 200000, Capacity: 500},
		&pb.Vessel{Id: "vessel002", Name: "Boaty McBoatface", MaxWeight: 500000, Capacity: 1000},
		&pb.Vessel{Id: "vessel003", Name: "Boaty McBoaty", MaxWeight: 250000, Capacity: 550},
	}
	for _, vessel := range vessels {
		repo.Create(vessel)
	}
}

func main() {
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = defaultHost
	}

	session, err := CreateSession(host)

	// Mgo creates a 'master' session, we need to end that session before the main function closes.
	defer session.Close()
	if err != nil {
		// We're wrapping the error returned from our CreateSession
		// here to add some context to the error.
		log.Panicf("Could not connect to datastore with host %s - %v", host, err)
	}

	repo := &VesselsRepository{session.Copy()}
	creatDeummyData(repo)

	// Create a new service
	srv := micro.NewService(
		// This name must match the package name given in protobuf definition
		micro.Name("go.micro.srv.vessel"),
		micro.Version("latest"),
	)

	// Init will parse the command line flags
	srv.Init()

	// Register handler
	pb.RegisterVesselServiceHandler(srv.Server(), &handler{session})

	if err := srv.Run(); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
