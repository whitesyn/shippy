package main

import (
	"github.com/jinzhu/gorm"
	pb "github.com/whitesyn/shippy/user-service/proto/user"
)

const (
	dbName     = "shippy"
	usersTable = "users"
)

// Repository Interface for repository to store vessels
type Repository interface {
	Create(user *pb.User) error
	Get(id string) (*pb.User, error)
	GetAll() ([]*pb.User, error)
	GetByEmailAndPassword(email string, password string) (*pb.User, error)
}

// UsersRepository structure
type UsersRepository struct {
	db *gorm.DB
}

// Get user by ID
func (repo *UsersRepository) Get(id string) (*pb.User, error) {
	var user *pb.User
	user.Id = id

	if err := repo.db.First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// Create new user
func (repo *UsersRepository) Create(user *pb.User) error {
	return repo.db.Create(user).Error
}

// GetAll users from repository
func (repo *UsersRepository) GetAll() ([]*pb.User, error) {
	var users []*pb.User
	if err := repo.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// GetByEmailAndPassword returns user by email and password
func (repo *UsersRepository) GetByEmailAndPassword(email string, password string) (*pb.User, error) {
	var user *pb.User
	user.Email = email
	user.Password = password

	if err := repo.db.First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
