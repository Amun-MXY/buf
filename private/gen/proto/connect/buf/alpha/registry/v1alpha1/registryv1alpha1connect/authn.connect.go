// Copyright 2020-2023 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: buf/alpha/registry/v1alpha1/authn.proto

package registryv1alpha1connect

import (
	context "context"
	errors "errors"
	v1alpha1 "github.com/bufbuild/buf/private/gen/proto/go/buf/alpha/registry/v1alpha1"
	connect_go "github.com/bufbuild/connect-go"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// AuthnServiceName is the fully-qualified name of the AuthnService service.
	AuthnServiceName = "buf.alpha.registry.v1alpha1.AuthnService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// AuthnServiceGetCurrentUserProcedure is the fully-qualified name of the AuthnService's
	// GetCurrentUser RPC.
	AuthnServiceGetCurrentUserProcedure = "/buf.alpha.registry.v1alpha1.AuthnService/GetCurrentUser"
	// AuthnServiceGetCurrentUserSubjectProcedure is the fully-qualified name of the AuthnService's
	// GetCurrentUserSubject RPC.
	AuthnServiceGetCurrentUserSubjectProcedure = "/buf.alpha.registry.v1alpha1.AuthnService/GetCurrentUserSubject"
)

// AuthnServiceClient is a client for the buf.alpha.registry.v1alpha1.AuthnService service.
type AuthnServiceClient interface {
	// GetCurrentUser gets information associated with the current user.
	//
	// The user's ID is retrieved from the request's authentication header.
	GetCurrentUser(context.Context, *connect_go.Request[v1alpha1.GetCurrentUserRequest]) (*connect_go.Response[v1alpha1.GetCurrentUserResponse], error)
	// GetCurrentUserSubject gets the currently logged in users subject.
	//
	// The user's ID is retrieved from the request's authentication header.
	GetCurrentUserSubject(context.Context, *connect_go.Request[v1alpha1.GetCurrentUserSubjectRequest]) (*connect_go.Response[v1alpha1.GetCurrentUserSubjectResponse], error)
}

// NewAuthnServiceClient constructs a client for the buf.alpha.registry.v1alpha1.AuthnService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewAuthnServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) AuthnServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &authnServiceClient{
		getCurrentUser: connect_go.NewClient[v1alpha1.GetCurrentUserRequest, v1alpha1.GetCurrentUserResponse](
			httpClient,
			baseURL+AuthnServiceGetCurrentUserProcedure,
			opts...,
		),
		getCurrentUserSubject: connect_go.NewClient[v1alpha1.GetCurrentUserSubjectRequest, v1alpha1.GetCurrentUserSubjectResponse](
			httpClient,
			baseURL+AuthnServiceGetCurrentUserSubjectProcedure,
			opts...,
		),
	}
}

// authnServiceClient implements AuthnServiceClient.
type authnServiceClient struct {
	getCurrentUser        *connect_go.Client[v1alpha1.GetCurrentUserRequest, v1alpha1.GetCurrentUserResponse]
	getCurrentUserSubject *connect_go.Client[v1alpha1.GetCurrentUserSubjectRequest, v1alpha1.GetCurrentUserSubjectResponse]
}

// GetCurrentUser calls buf.alpha.registry.v1alpha1.AuthnService.GetCurrentUser.
func (c *authnServiceClient) GetCurrentUser(ctx context.Context, req *connect_go.Request[v1alpha1.GetCurrentUserRequest]) (*connect_go.Response[v1alpha1.GetCurrentUserResponse], error) {
	return c.getCurrentUser.CallUnary(ctx, req)
}

// GetCurrentUserSubject calls buf.alpha.registry.v1alpha1.AuthnService.GetCurrentUserSubject.
func (c *authnServiceClient) GetCurrentUserSubject(ctx context.Context, req *connect_go.Request[v1alpha1.GetCurrentUserSubjectRequest]) (*connect_go.Response[v1alpha1.GetCurrentUserSubjectResponse], error) {
	return c.getCurrentUserSubject.CallUnary(ctx, req)
}

// AuthnServiceHandler is an implementation of the buf.alpha.registry.v1alpha1.AuthnService service.
type AuthnServiceHandler interface {
	// GetCurrentUser gets information associated with the current user.
	//
	// The user's ID is retrieved from the request's authentication header.
	GetCurrentUser(context.Context, *connect_go.Request[v1alpha1.GetCurrentUserRequest]) (*connect_go.Response[v1alpha1.GetCurrentUserResponse], error)
	// GetCurrentUserSubject gets the currently logged in users subject.
	//
	// The user's ID is retrieved from the request's authentication header.
	GetCurrentUserSubject(context.Context, *connect_go.Request[v1alpha1.GetCurrentUserSubjectRequest]) (*connect_go.Response[v1alpha1.GetCurrentUserSubjectResponse], error)
}

// NewAuthnServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewAuthnServiceHandler(svc AuthnServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle(AuthnServiceGetCurrentUserProcedure, connect_go.NewUnaryHandler(
		AuthnServiceGetCurrentUserProcedure,
		svc.GetCurrentUser,
		opts...,
	))
	mux.Handle(AuthnServiceGetCurrentUserSubjectProcedure, connect_go.NewUnaryHandler(
		AuthnServiceGetCurrentUserSubjectProcedure,
		svc.GetCurrentUserSubject,
		opts...,
	))
	return "/buf.alpha.registry.v1alpha1.AuthnService/", mux
}

// UnimplementedAuthnServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedAuthnServiceHandler struct{}

func (UnimplementedAuthnServiceHandler) GetCurrentUser(context.Context, *connect_go.Request[v1alpha1.GetCurrentUserRequest]) (*connect_go.Response[v1alpha1.GetCurrentUserResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("buf.alpha.registry.v1alpha1.AuthnService.GetCurrentUser is not implemented"))
}

func (UnimplementedAuthnServiceHandler) GetCurrentUserSubject(context.Context, *connect_go.Request[v1alpha1.GetCurrentUserSubjectRequest]) (*connect_go.Response[v1alpha1.GetCurrentUserSubjectResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("buf.alpha.registry.v1alpha1.AuthnService.GetCurrentUserSubject is not implemented"))
}
