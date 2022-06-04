package handlers

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	pb "maples"
	"maples/maples-service/module/entity"
	"maples/maples-service/module/resposiitory"
)

// NewService returns a nave, stateless implementation of Service.
func NewService() pb.MaplesServer {
	return maplesService{
		userModel: resposiitory.NewUserModel(),
	}
}

type maplesService struct {
	userModel resposiitory.UserModel
}

func (s maplesService) Hello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	helloMessage := fmt.Sprintf("Hello, %s", in.Name)
	resp := pb.HelloResponse{
		Msg:  "SUCCESS",
		Code: 200,
		Data: &pb.HelloData{Message: helloMessage},
	}
	return &resp, nil
}

func (s maplesService) UpdateUserMessage(ctx context.Context, in *pb.UserMessageRequest) (*pb.UserMessageResponse, error) {
	if len(in.Phone) == 0 {
		return nil, errors.New("update user fata")
	}
	user := entity.User{
		Name:     in.Name,
		Sex:      int(in.Sex),
		Phone:    in.Phone,
		Birthday: in.Birthday,
	}
	err := s.userModel.Update(ctx, &user)
	if err != nil {
		return nil, errors.New("update user fata")
	}
	resp := &pb.UserMessageResponse{
		Msg:  "SUCCESS",
		Code: 200,
		Data: nil,
	}
	return resp, nil
}

func (s maplesService) AddUser(ctx context.Context, in *pb.UserMessageRequest) (*pb.UserMessageResponse, error) {
	user := entity.User{
		Name:     in.Name,
		Sex:      int(in.Sex),
		Phone:    in.Phone,
		Birthday: in.Birthday,
	}
	err := s.userModel.Create(ctx, &user)
	if err != nil {
		return nil, errors.New("create user fata")
	}
	resp := &pb.UserMessageResponse{
		Msg:  "SUCCESS",
		Code: 200,
		Data: nil,
	}
	return resp, nil
}
