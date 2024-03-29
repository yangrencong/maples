// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: cda8b667e3
// Version Date: 2021-09-27T00:46:54Z

package svc

// This file provides server-side bindings for the gRPC transport.
// It utilizes the transport/grpc.Server.

import (
	"context"
	"net/http"

	"google.golang.org/grpc/metadata"

	grpctransport "github.com/go-kit/kit/transport/grpc"

	// This Service
	pb "maples/pb"
)

// MakeGRPCServer makes a set of endpoints available as a gRPC MaplesServer.
func MakeGRPCServer(endpoints Endpoints, options ...grpctransport.ServerOption) pb.MaplesServer {
	serverOptions := []grpctransport.ServerOption{
		grpctransport.ServerBefore(metadataToContext),
	}
	serverOptions = append(serverOptions, options...)
	return &grpcServer{
		// maples

		hello: grpctransport.NewServer(
			endpoints.HelloEndpoint,
			DecodeGRPCHelloRequest,
			EncodeGRPCHelloResponse,
			serverOptions...,
		),
		adduser: grpctransport.NewServer(
			endpoints.AddUserEndpoint,
			DecodeGRPCAddUserRequest,
			EncodeGRPCAddUserResponse,
			serverOptions...,
		),
		updateusermessage: grpctransport.NewServer(
			endpoints.UpdateUserMessageEndpoint,
			DecodeGRPCUpdateUserMessageRequest,
			EncodeGRPCUpdateUserMessageResponse,
			serverOptions...,
		),
		getusermessage: grpctransport.NewServer(
			endpoints.GetUserMessageEndpoint,
			DecodeGRPCGetUserMessageRequest,
			EncodeGRPCGetUserMessageResponse,
			serverOptions...,
		),
	}
}

// grpcServer implements the MaplesServer interface
type grpcServer struct {
	hello             grpctransport.Handler
	adduser           grpctransport.Handler
	updateusermessage grpctransport.Handler
	getusermessage    grpctransport.Handler
}

// Methods for grpcServer to implement MaplesServer interface

func (s *grpcServer) Hello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	_, rep, err := s.hello.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.HelloResponse), nil
}

func (s *grpcServer) AddUser(ctx context.Context, req *pb.UserMessageRequest) (*pb.UserMessageResponse, error) {
	_, rep, err := s.adduser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.UserMessageResponse), nil
}

func (s *grpcServer) UpdateUserMessage(ctx context.Context, req *pb.UserMessageRequest) (*pb.UserMessageResponse, error) {
	_, rep, err := s.updateusermessage.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.UserMessageResponse), nil
}

func (s *grpcServer) GetUserMessage(ctx context.Context, req *pb.GetUserMessageRequest) (*pb.GetUserMessageResponse, error) {
	_, rep, err := s.getusermessage.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetUserMessageResponse), nil
}

// Server Decode

// DecodeGRPCHelloRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC hello request to a user-domain hello request. Primarily useful in a server.
func DecodeGRPCHelloRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.HelloRequest)
	return req, nil
}

// DecodeGRPCAddUserRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC adduser request to a user-domain adduser request. Primarily useful in a server.
func DecodeGRPCAddUserRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.UserMessageRequest)
	return req, nil
}

// DecodeGRPCUpdateUserMessageRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC updateusermessage request to a user-domain updateusermessage request. Primarily useful in a server.
func DecodeGRPCUpdateUserMessageRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.UserMessageRequest)
	return req, nil
}

// DecodeGRPCGetUserMessageRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC getusermessage request to a user-domain getusermessage request. Primarily useful in a server.
func DecodeGRPCGetUserMessageRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.GetUserMessageRequest)
	return req, nil
}

// Server Encode

// EncodeGRPCHelloResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain hello response to a gRPC hello reply. Primarily useful in a server.
func EncodeGRPCHelloResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.HelloResponse)
	return resp, nil
}

// EncodeGRPCAddUserResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain adduser response to a gRPC adduser reply. Primarily useful in a server.
func EncodeGRPCAddUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.UserMessageResponse)
	return resp, nil
}

// EncodeGRPCUpdateUserMessageResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain updateusermessage response to a gRPC updateusermessage reply. Primarily useful in a server.
func EncodeGRPCUpdateUserMessageResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.UserMessageResponse)
	return resp, nil
}

// EncodeGRPCGetUserMessageResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain getusermessage response to a gRPC getusermessage reply. Primarily useful in a server.
func EncodeGRPCGetUserMessageResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.GetUserMessageResponse)
	return resp, nil
}

// Helpers

func metadataToContext(ctx context.Context, md metadata.MD) context.Context {
	for k, v := range md {
		if v != nil {
			// The key is added both in metadata format (k) which is all lower
			// and the http.CanonicalHeaderKey of the key so that it can be
			// accessed in either format
			ctx = context.WithValue(ctx, k, v[0])
			ctx = context.WithValue(ctx, http.CanonicalHeaderKey(k), v[0])
		}
	}

	return ctx
}
