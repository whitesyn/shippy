package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	"github.com/micro/go-micro/broker"
	pb "github.com/whitesyn/shippy/user-service/proto/user"
	"golang.org/x/crypto/bcrypt"
)

const topic = "user.created"

type handler struct {
	repo         Repository
	tokenService Authable
	pubsub       broker.Broker
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

	if err := h.PublishEvent(req); err != nil {
		return err
	}

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

func (h *handler) PublishEvent(user *pb.User) error {
	body, err := json.Marshal(user)
	if err != nil {
		return err
	}

	msg := &broker.Message{
		Header: map[string]string{
			"id": user.Id,
		},
		Body: body,
	}

	if err := h.pubsub.Publish(topic, msg); err != nil {
		log.Printf("[pub] failed: %v", err)
		return err
	}

	return nil
}
