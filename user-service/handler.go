package main

import (
	"context"

	pb "github.com/whitesyn/shippy/user-service/proto/user"
)

type handler struct {
	repo Repository
	// tokenService Authable
}

func (h *handler) Create(ctx context.Context, req *pb.User, res *pb.CreateResponse) error {
	if err := h.repo.Create(req); err != nil {
		return err
	}

	res.Created = true
	res.User = req
	return nil
}

func (h *handler) Get(ctx context.Context, req *pb.User, res *pb.UserResponse) error {
	user, err := h.repo.Get(req.Id)
	if err != nil {
		return err
	}

	res.User = user
	return nil
}

func (h *handler) GetAll(ctx context.Context, req *pb.Request, res *pb.UsersResponse) error {
	users, err := h.repo.GetAll()
	if err != nil {
		return err
	}

	res.Users = users
	return nil
}

func (h *handler) Auth(ctx context.Context, req *pb.User, res *pb.Token) error {
	user, err := h.repo.GetByEmailAndPassword(req.Email, req.Password)
	if err != nil {
		return nil
	}

	res.Token = user.Id + "_" + user.Email
	res.Valid = true
	return nil
}

func (h *handler) ValidateToken(ctx context.Context, req *pb.Token, res *pb.Token) error {
	return nil
}
