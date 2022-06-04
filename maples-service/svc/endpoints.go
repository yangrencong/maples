// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: cda8b667e3
// Version Date: 2021-09-27T00:46:54Z

package svc

// This file contains methods to make individual endpoints from services,
// request and response types to serve those endpoints, as well as encoders and
// decoders for those types, for all of our supported transport serialization
// formats.

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/endpoint"

	pb "maples"
)

// Endpoints collects all of the endpoints that compose an add service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
//
// In a server, it's useful for functions that need to operate on a per-endpoint
// basis. For example, you might pass an Endpoints to a function that produces
// an http.Handler, with each method (endpoint) wired up to a specific path. (It
// is probably a mistake in design to invoke the Service methods on the
// Endpoints struct in a server.)
//
// In a client, it's useful to collect individually constructed endpoints into a
// single type that implements the Service interface. For example, you might
// construct individual endpoints using transport/http.NewClient, combine them into an Endpoints, and return it to the caller as a Service.
type Endpoints struct {
	HelloEndpoint             endpoint.Endpoint
	AddUserEndpoint           endpoint.Endpoint
	UpdateUserMessageEndpoint endpoint.Endpoint
}

// Endpoints

func (e Endpoints) Hello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	response, err := e.HelloEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.HelloResponse), nil
}

func (e Endpoints) AddUser(ctx context.Context, in *pb.UserMessageRequest) (*pb.UserMessageResponse, error) {
	response, err := e.AddUserEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.UserMessageResponse), nil
}

func (e Endpoints) UpdateUserMessage(ctx context.Context, in *pb.UserMessageRequest) (*pb.UserMessageResponse, error) {
	response, err := e.UpdateUserMessageEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.UserMessageResponse), nil
}

// Make Endpoints

func MakeHelloEndpoint(s pb.MaplesServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.HelloRequest)
		v, err := s.Hello(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func MakeAddUserEndpoint(s pb.MaplesServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.UserMessageRequest)
		v, err := s.AddUser(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func MakeUpdateUserMessageEndpoint(s pb.MaplesServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.UserMessageRequest)
		v, err := s.UpdateUserMessage(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

// WrapAllExcept wraps each Endpoint field of struct Endpoints with a
// go-kit/kit/endpoint.Middleware.
// Use this for applying a set of middlewares to every endpoint in the service.
// Optionally, endpoints can be passed in by name to be excluded from being wrapped.
// WrapAllExcept(middleware, "Status", "Ping")
func (e *Endpoints) WrapAllExcept(middleware endpoint.Middleware, excluded ...string) {
	included := map[string]struct{}{
		"Hello":             {},
		"AddUser":           {},
		"UpdateUserMessage": {},
	}

	for _, ex := range excluded {
		if _, ok := included[ex]; !ok {
			panic(fmt.Sprintf("Excluded endpoint '%s' does not exist; see middlewares/endpoints.go", ex))
		}
		delete(included, ex)
	}

	for inc := range included {
		if inc == "Hello" {
			e.HelloEndpoint = middleware(e.HelloEndpoint)
		}
		if inc == "AddUser" {
			e.AddUserEndpoint = middleware(e.AddUserEndpoint)
		}
		if inc == "UpdateUserMessage" {
			e.UpdateUserMessageEndpoint = middleware(e.UpdateUserMessageEndpoint)
		}
	}
}

// LabeledMiddleware will get passed the endpoint name when passed to
// WrapAllLabeledExcept, this can be used to write a generic metrics
// middleware which can send the endpoint name to the metrics collector.
type LabeledMiddleware func(string, endpoint.Endpoint) endpoint.Endpoint

// WrapAllLabeledExcept wraps each Endpoint field of struct Endpoints with a
// LabeledMiddleware, which will receive the name of the endpoint. See
// LabeldMiddleware. See method WrapAllExept for details on excluded
// functionality.
func (e *Endpoints) WrapAllLabeledExcept(middleware func(string, endpoint.Endpoint) endpoint.Endpoint, excluded ...string) {
	included := map[string]struct{}{
		"Hello":             {},
		"AddUser":           {},
		"UpdateUserMessage": {},
	}

	for _, ex := range excluded {
		if _, ok := included[ex]; !ok {
			panic(fmt.Sprintf("Excluded endpoint '%s' does not exist; see middlewares/endpoints.go", ex))
		}
		delete(included, ex)
	}

	for inc := range included {
		if inc == "Hello" {
			e.HelloEndpoint = middleware("Hello", e.HelloEndpoint)
		}
		if inc == "AddUser" {
			e.AddUserEndpoint = middleware("AddUser", e.AddUserEndpoint)
		}
		if inc == "UpdateUserMessage" {
			e.UpdateUserMessageEndpoint = middleware("UpdateUserMessage", e.UpdateUserMessageEndpoint)
		}
	}
}