// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: cda8b667e3
// Version Date: 2021-09-27T00:46:54Z

// Package grpc provides a gRPC client for the Maples service.
package grpc

import (
	"context"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"

	// This Service
	"maples/maples-service/svc"
	pb "maples/pb"
)

// New returns an service backed by a gRPC client connection. It is the
// responsibility of the caller to dial, and later close, the connection.
func New(conn *grpc.ClientConn, options ...ClientOption) (pb.MaplesServer, error) {
	var cc clientConfig

	for _, f := range options {
		err := f(&cc)
		if err != nil {
			return nil, errors.Wrap(err, "cannot apply option")
		}
	}

	clientOptions := []grpctransport.ClientOption{
		grpctransport.ClientBefore(
			contextValuesToGRPCMetadata(cc.headers)),
	}
	var helloEndpoint endpoint.Endpoint
	{
		helloEndpoint = grpctransport.NewClient(
			conn,
			"maples.Maples",
			"Hello",
			EncodeGRPCHelloRequest,
			DecodeGRPCHelloResponse,
			pb.HelloResponse{},
			clientOptions...,
		).Endpoint()
	}

	var adduserEndpoint endpoint.Endpoint
	{
		adduserEndpoint = grpctransport.NewClient(
			conn,
			"maples.Maples",
			"AddUser",
			EncodeGRPCAddUserRequest,
			DecodeGRPCAddUserResponse,
			pb.UserMessageResponse{},
			clientOptions...,
		).Endpoint()
	}

	var updateusermessageEndpoint endpoint.Endpoint
	{
		updateusermessageEndpoint = grpctransport.NewClient(
			conn,
			"maples.Maples",
			"UpdateUserMessage",
			EncodeGRPCUpdateUserMessageRequest,
			DecodeGRPCUpdateUserMessageResponse,
			pb.UserMessageResponse{},
			clientOptions...,
		).Endpoint()
	}

	var getusermessageEndpoint endpoint.Endpoint
	{
		getusermessageEndpoint = grpctransport.NewClient(
			conn,
			"maples.Maples",
			"GetUserMessage",
			EncodeGRPCGetUserMessageRequest,
			DecodeGRPCGetUserMessageResponse,
			pb.GetUserMessageResponse{},
			clientOptions...,
		).Endpoint()
	}

	return svc.Endpoints{
		HelloEndpoint:             helloEndpoint,
		AddUserEndpoint:           adduserEndpoint,
		UpdateUserMessageEndpoint: updateusermessageEndpoint,
		GetUserMessageEndpoint:    getusermessageEndpoint,
	}, nil
}

// GRPC Client Decode

// DecodeGRPCHelloResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC hello reply to a user-domain hello response. Primarily useful in a client.
func DecodeGRPCHelloResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.HelloResponse)
	return reply, nil
}

// DecodeGRPCAddUserResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC adduser reply to a user-domain adduser response. Primarily useful in a client.
func DecodeGRPCAddUserResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.UserMessageResponse)
	return reply, nil
}

// DecodeGRPCUpdateUserMessageResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC updateusermessage reply to a user-domain updateusermessage response. Primarily useful in a client.
func DecodeGRPCUpdateUserMessageResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.UserMessageResponse)
	return reply, nil
}

// DecodeGRPCGetUserMessageResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC getusermessage reply to a user-domain getusermessage response. Primarily useful in a client.
func DecodeGRPCGetUserMessageResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.GetUserMessageResponse)
	return reply, nil
}

// GRPC Client Encode

// EncodeGRPCHelloRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain hello request to a gRPC hello request. Primarily useful in a client.
func EncodeGRPCHelloRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.HelloRequest)
	return req, nil
}

// EncodeGRPCAddUserRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain adduser request to a gRPC adduser request. Primarily useful in a client.
func EncodeGRPCAddUserRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UserMessageRequest)
	return req, nil
}

// EncodeGRPCUpdateUserMessageRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain updateusermessage request to a gRPC updateusermessage request. Primarily useful in a client.
func EncodeGRPCUpdateUserMessageRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UserMessageRequest)
	return req, nil
}

// EncodeGRPCGetUserMessageRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain getusermessage request to a gRPC getusermessage request. Primarily useful in a client.
func EncodeGRPCGetUserMessageRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetUserMessageRequest)
	return req, nil
}

type clientConfig struct {
	headers []string
}

// ClientOption is a function that modifies the client config
type ClientOption func(*clientConfig) error

func CtxValuesToSend(keys ...string) ClientOption {
	return func(o *clientConfig) error {
		o.headers = keys
		return nil
	}
}

func contextValuesToGRPCMetadata(keys []string) grpctransport.ClientRequestFunc {
	return func(ctx context.Context, md *metadata.MD) context.Context {
		var pairs []string
		for _, k := range keys {
			if v, ok := ctx.Value(k).(string); ok {
				pairs = append(pairs, k, v)
			}
		}

		if pairs != nil {
			*md = metadata.Join(*md, metadata.Pairs(pairs...))
		}

		return ctx
	}
}
