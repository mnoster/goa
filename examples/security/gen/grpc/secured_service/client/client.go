// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// secured_service GRPC client
//
// Command:
// $ goa gen goa.design/goa/examples/security/design -o
// $(GOPATH)/src/goa.design/goa/examples/security

package client

import (
	"context"

	goa "goa.design/goa"
	secured_servicepb "goa.design/goa/examples/security/gen/grpc/secured_service"
	securedservice "goa.design/goa/examples/security/gen/secured_service"
	goagrpc "goa.design/goa/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// Client lists the service endpoint gRPC clients.
type Client struct {
	grpccli secured_servicepb.SecuredServiceClient
	opts    []grpc.CallOption
}

// NewClient instantiates gRPC client for all the secured_service service
// servers.
func NewClient(cc *grpc.ClientConn, opts ...grpc.CallOption) *Client {
	return &Client{
		grpccli: secured_servicepb.NewSecuredServiceClient(cc),
		opts:    opts,
	}
}

// Signin calls the "Signin" function in secured_servicepb.SecuredServiceClient
// interface.
func (c *Client) Signin() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		er, err := EncodeSigninRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		req := er.(*secured_servicepb.SigninRequest)
		p, ok := v.(*securedservice.SigninPayload)
		if !ok {
			return nil, goagrpc.ErrInvalidType("secured_service", "signin", "*securedservice.SigninPayload", v)
		}
		ctx = metadata.AppendToOutgoingContext(ctx, "username", p.Username)
		ctx = metadata.AppendToOutgoingContext(ctx, "password", p.Password)
		resp, err := c.grpccli.Signin(ctx, req, c.opts...)
		if err != nil {
			return nil, err
		}
		r, err := DecodeSigninResponse(ctx, resp)
		if err != nil {
			return nil, err
		}
		res := r.(*securedservice.Creds)
		return res, nil
	}
}

// Secure calls the "Secure" function in secured_servicepb.SecuredServiceClient
// interface.
func (c *Client) Secure() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		er, err := EncodeSecureRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		req := er.(*secured_servicepb.SecureRequest)
		p, ok := v.(*securedservice.SecurePayload)
		if !ok {
			return nil, goagrpc.ErrInvalidType("secured_service", "secure", "*securedservice.SecurePayload", v)
		}
		ctx = metadata.AppendToOutgoingContext(ctx, "authorization", p.Token)
		resp, err := c.grpccli.Secure(ctx, req, c.opts...)
		if err != nil {
			return nil, err
		}
		r, err := DecodeSecureResponse(ctx, resp)
		if err != nil {
			return nil, err
		}
		res := r.(string)
		return res, nil
	}
}

// DoublySecure calls the "DoublySecure" function in
// secured_servicepb.SecuredServiceClient interface.
func (c *Client) DoublySecure() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		er, err := EncodeDoublySecureRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		req := er.(*secured_servicepb.DoublySecureRequest)
		p, ok := v.(*securedservice.DoublySecurePayload)
		if !ok {
			return nil, goagrpc.ErrInvalidType("secured_service", "doubly_secure", "*securedservice.DoublySecurePayload", v)
		}
		ctx = metadata.AppendToOutgoingContext(ctx, "authorization", p.Token)
		resp, err := c.grpccli.DoublySecure(ctx, req, c.opts...)
		if err != nil {
			return nil, err
		}
		r, err := DecodeDoublySecureResponse(ctx, resp)
		if err != nil {
			return nil, err
		}
		res := r.(string)
		return res, nil
	}
}

// AlsoDoublySecure calls the "AlsoDoublySecure" function in
// secured_servicepb.SecuredServiceClient interface.
func (c *Client) AlsoDoublySecure() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		er, err := EncodeAlsoDoublySecureRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		req := er.(*secured_servicepb.AlsoDoublySecureRequest)
		p, ok := v.(*securedservice.AlsoDoublySecurePayload)
		if !ok {
			return nil, goagrpc.ErrInvalidType("secured_service", "also_doubly_secure", "*securedservice.AlsoDoublySecurePayload", v)
		}
		if p.OauthToken != nil {
			ctx = metadata.AppendToOutgoingContext(ctx, "oauth", *p.OauthToken)
		}
		if p.Token != nil {
			ctx = metadata.AppendToOutgoingContext(ctx, "authorization", *p.Token)
		}
		resp, err := c.grpccli.AlsoDoublySecure(ctx, req, c.opts...)
		if err != nil {
			return nil, err
		}
		r, err := DecodeAlsoDoublySecureResponse(ctx, resp)
		if err != nil {
			return nil, err
		}
		res := r.(string)
		return res, nil
	}
}
