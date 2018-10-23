package main

import (
	"context"

	pb "github.com/whitesyn/shippy/vessel-service/proto/vessel"
	mgo "gopkg.in/mgo.v2"
)

// Service should implement all of the methods to satisfy the service
// we defined in our protobuf definition. You can check the interface
// in the generated code itself for the exact method signatures etc
// to give you a better idea.
type handler struct {
	session *mgo.Session
}

func (s *handler) GetRepo() Repository {
	return &VesselsRepository{s.session.Clone()}
}

// FindAvailable vessels handler
func (s *handler) CreateVessel(ctx context.Context, req *pb.Vessel, res *pb.CreateResponse) error {
	repo := s.GetRepo()
	defer repo.Close()

	if err := repo.Create(req); err != nil {
		return err
	}

	res.Created = true
	res.Vessel = req
	return nil
}

// FindAvailable vessels handler
func (s *handler) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.VesselResponse) error {
	repo := s.GetRepo()
	defer repo.Close()

	vessel, err := repo.FindAvailable(req)
	if err != nil {
		return err
	}

	res.Vessel = vessel
	return nil
}
