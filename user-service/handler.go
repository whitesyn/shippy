package main

import (
	"context"
	"errors"

	pb "github.com/whitesyn/shippy/user-service/proto/user"
	"golang.org/x/crypto/bcrypt"
)

type handler struct {
	repo         Repository
	tokenService Authable
}

func (h *handler) Create(ctx context.Context, req *pb.User, res *pb.CreateResponse) error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	req.Password = string(hashedPass)

	user, err := h.repo.GetByEmail(req.Email)
	if err != nil {
		return err
	}

	if user != nil {
		return errors.New("User with this e-mail already exists")
	}

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
	user, err := h.repo.GetByEmail(req.Email)
	if err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return err
	}
	token, err := h.tokenService.Encode(user)
	if err != nil {
		return err
	}
	res.Token = token
	res.Valid = true
	return nil
}

func (h *handler) ValidateToken(ctx context.Context, req *pb.Token, res *pb.Token) error {
	claims, err := h.tokenService.Decode(req.Token)
	if err != nil {
		return err
	}

	if claims.User.Id == "" {
		return errors.New("invalid user")
	}

	res.Valid = true
	return nil
}
