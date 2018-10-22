package main

import (
	"context"
	"errors"
	"log"

	micro "github.com/micro/go-micro"
	pb "github.com/whitesyn/shippy/vessel-service/proto/vessel"
)

// Repository Interface for repository to store vessels
type Repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
}

// VesselsRepository - Dummy repository, this simulates the use of a datastore
// of some kind. We'll replace this with a real implementation later on.
type VesselsRepository struct {
	vessels []*pb.Vessel
}

// FindAvailable - checks a specification against a map of vessels,
// if capacity and max weight are below a vessels capacity and max weight,
// then return that vessel.
func (repo *VesselsRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	for _, vessel := range repo.vessels {
		if spec.Capacity <= vessel.Capacity && spec.MaxWeight <= vessel.MaxWeight {
			return vessel, nil
		}
	}
	return nil, errors.New("No vessel found by that spec")
}

// Service definition
type service struct {
	repo Repository
}

// FindAvailable vessels handler
func (s *service) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.Response) error {
	vessel, err := s.repo.FindAvailable(req)
	if err != nil {
		return err
	}

	res.Vessel = vessel
	return nil
}

func main() {
	vessels := []*pb.Vessel{
		&pb.Vessel{Id: "vessel001", Name: "Boaty McBoatface", MaxWeight: 200000, Capacity: 500},
		&pb.Vessel{Id: "vessel002", Name: "Boaty McBoaty", MaxWeight: 250000, Capacity: 550},
	}
	repo := &VesselsRepository{vessels}

	// Create a new service
	srv := micro.NewService(
		// This name must match the package name given in protobuf definition
		micro.Name("go.micro.srv.vessel"),
		micro.Version("latest"),
	)

	// Init will parse the command line flags
	srv.Init()

	// Register handler
	pb.RegisterVesselServiceHandler(srv.Server(), &service{repo})

	if err := srv.Run(); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
