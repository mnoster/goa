// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// secured_service GRPC server
//
// Command:
// $ goa gen goa.design/goa/examples/security/design -o
// $(GOPATH)/src/goa.design/goa/examples/security

package server

import (
	"context"

	secured_servicepb "goa.design/goa/examples/security/gen/grpc/secured_service"
	securedservice "goa.design/goa/examples/security/gen/secured_service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Server implements the secured_servicepb.SecuredServiceServer interface.
type Server struct {
	endpoints *securedservice.Endpoints
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the expr.
type ErrorNamer interface {
	ErrorName() string
}

// New instantiates the server struct with the secured_service service
// endpoints.
func New(e *securedservice.Endpoints) *Server {
	return &Server{e}
}

// Signin implements the "Signin" method in
// secured_servicepb.SecuredServiceServer interface.
func (s *Server) Signin(ctx context.Context, message *secured_servicepb.SigninRequest) (*secured_servicepb.SigninResponse, error) {
	p, err := DecodeSigninRequest(ctx, message)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	payload := p.(*securedservice.SigninPayload)
	v, err := s.endpoints.Signin(ctx, payload)
	if err != nil {
		en, ok := err.(ErrorNamer)
		if !ok {
			return nil, err
		}
		switch en.ErrorName() {
		case "unauthorized":
			return nil, status.Error(codes.Unauthenticated, err.Error())
		}
	}
	r, err := EncodeSigninResponse(ctx, v)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return r.(*secured_servicepb.SigninResponse), nil
}

// Secure implements the "Secure" method in
// secured_servicepb.SecuredServiceServer interface.
func (s *Server) Secure(ctx context.Context, message *secured_servicepb.SecureRequest) (*secured_servicepb.SecureResponse, error) {
	p, err := DecodeSecureRequest(ctx, message)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	payload := p.(*securedservice.SecurePayload)
	v, err := s.endpoints.Secure(ctx, payload)
	if err != nil {
		en, ok := err.(ErrorNamer)
		if !ok {
			return nil, err
		}
		switch en.ErrorName() {
		case "invalid-scopes":
			return nil, status.Error(codes.Unauthenticated, err.Error())
		case "unauthorized":
			return nil, status.Error(codes.Unauthenticated, err.Error())
		}
	}
	r, err := EncodeSecureResponse(ctx, v)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return r.(*secured_servicepb.SecureResponse), nil
}

// DoublySecure implements the "DoublySecure" method in
// secured_servicepb.SecuredServiceServer interface.
func (s *Server) DoublySecure(ctx context.Context, message *secured_servicepb.DoublySecureRequest) (*secured_servicepb.DoublySecureResponse, error) {
	p, err := DecodeDoublySecureRequest(ctx, message)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	payload := p.(*securedservice.DoublySecurePayload)
	v, err := s.endpoints.DoublySecure(ctx, payload)
	if err != nil {
		en, ok := err.(ErrorNamer)
		if !ok {
			return nil, err
		}
		switch en.ErrorName() {
		case "invalid-scopes":
			return nil, status.Error(codes.Unauthenticated, err.Error())
		case "unauthorized":
			return nil, status.Error(codes.Unauthenticated, err.Error())
		}
	}
	r, err := EncodeDoublySecureResponse(ctx, v)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return r.(*secured_servicepb.DoublySecureResponse), nil
}

// AlsoDoublySecure implements the "AlsoDoublySecure" method in
// secured_servicepb.SecuredServiceServer interface.
func (s *Server) AlsoDoublySecure(ctx context.Context, message *secured_servicepb.AlsoDoublySecureRequest) (*secured_servicepb.AlsoDoublySecureResponse, error) {
	p, err := DecodeAlsoDoublySecureRequest(ctx, message)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	payload := p.(*securedservice.AlsoDoublySecurePayload)
	v, err := s.endpoints.AlsoDoublySecure(ctx, payload)
	if err != nil {
		en, ok := err.(ErrorNamer)
		if !ok {
			return nil, err
		}
		switch en.ErrorName() {
		case "invalid-scopes":
			return nil, status.Error(codes.Unauthenticated, err.Error())
		case "unauthorized":
			return nil, status.Error(codes.Unauthenticated, err.Error())
		}
	}
	r, err := EncodeAlsoDoublySecureResponse(ctx, v)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return r.(*secured_servicepb.AlsoDoublySecureResponse), nil
}
