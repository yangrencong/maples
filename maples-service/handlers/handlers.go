package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"maples/maples-service/infra/mysql"
	"maples/maples-service/module/entity"
	"maples/maples-service/module/resposiitory"
	pb "maples/pb"
	"time"
)

// NewService returns a nave, stateless implementation of Service.
func NewService() pb.MaplesServer {
	return maplesService{}
}

type maplesService struct{}

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
	msg, _ := json.Marshal(in)
	log.Printf("request message is %v\n", string(msg))
	resp := &pb.UserMessageResponse{
		Msg:  "SUCCESS",
		Code: 200,
		Data: nil,
	}
	return resp, nil
}

func (s maplesService) AddUser(ctx context.Context, in *pb.UserMessageRequest) (*pb.UserMessageResponse, error) {
	db := mysql.GetMysqlClient()
	if db == nil {
		return nil, errors.New("db error")
	}
	user := entity.User{
		Name:       in.Name,
		Sex:        in.Sex,
		Phone:      in.Phone,
		Birthday:   in.Birthday,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}

	userModel := resposiitory.NewUserModel()
	err := userModel.Create(ctx, db, &user)
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

func (s maplesService) GetUserMessage(ctx context.Context, in *pb.GetUserMessageRequest) (*pb.GetUserMessageResponse, error) {

	db := mysql.GetMysqlClient()
	if db == nil {
		return nil, errors.New("db error")
	}

	userModel := resposiitory.NewUserModel()
	userData, err := userModel.Find(ctx, db, &entity.User{
		Name:  in.Name,
		Phone: in.Phone,
	})
	if err != nil {
		return nil, err
	}

	resp := &pb.GetUserMessageResponse{
		Msg:  "SUCCESS",
		Code: 200,
		Data: &pb.UserData{
			Name:     userData.Name,
			Sex:      userData.Sex,
			Birthday: userData.Birthday,
			Phone:    userData.Phone,
		},
	}
	return resp, nil
}
