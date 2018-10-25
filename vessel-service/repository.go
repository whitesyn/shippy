package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	pb "github.com/whitesyn/shippy/vessel-service/proto/vessel"
)

const (
	dbName           = "shippy"
	vesselCollection = "vessels"
)

// Repository Interface for repository to store vessels
type Repository interface {
	Create(*pb.Vessel) error
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
	Close()
}

// VesselsRepository structure
type VesselsRepository struct {
	session *mgo.Session
}

// Create new vessel
func (repo *VesselsRepository) Create(vessel *pb.Vessel) error {
	return repo.collection().Insert(vessel)
}

// FindAvailable - checks a specification against a map of vessels,
// if capacity and max weight are below a vessels capacity and max weight,
// then return that vessel.
func (repo *VesselsRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	var vessel *pb.Vessel
	err := repo.collection().Find(bson.M{
		"capacity":  bson.M{"$gte": spec.Capacity},
		"maxweight": bson.M{"$gte": spec.MaxWeight},
	}).One(&vessel)
	if err != nil {
		return nil, err
	}
	return vessel, nil
}

// Close closes the database session after each query has ran.
func (repo *VesselsRepository) Close() {
	repo.session.Close()
}

// Get vessels connection
func (repo *VesselsRepository) collection() *mgo.Collection {
	return repo.session.DB(dbName).C(vesselCollection)
}
