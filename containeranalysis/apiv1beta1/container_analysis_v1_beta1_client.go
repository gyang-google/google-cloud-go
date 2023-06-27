// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package containeranalysis

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"time"

	containeranalysispb "cloud.google.com/go/containeranalysis/apiv1beta1/containeranalysispb"
	iampb "cloud.google.com/go/iam/apiv1/iampb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httptransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
)

var newContainerAnalysisV1Beta1ClientHook clientHook

// ContainerAnalysisV1Beta1CallOptions contains the retry settings for each method of ContainerAnalysisV1Beta1Client.
type ContainerAnalysisV1Beta1CallOptions struct {
	SetIamPolicy            []gax.CallOption
	GetIamPolicy            []gax.CallOption
	TestIamPermissions      []gax.CallOption
	GeneratePackagesSummary []gax.CallOption
}

func defaultContainerAnalysisV1Beta1GRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("containeranalysis.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("containeranalysis.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://containeranalysis.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultContainerAnalysisV1Beta1CallOptions() *ContainerAnalysisV1Beta1CallOptions {
	return &ContainerAnalysisV1Beta1CallOptions{
		SetIamPolicy: []gax.CallOption{
			gax.WithTimeout(30000 * time.Millisecond),
		},
		GetIamPolicy: []gax.CallOption{
			gax.WithTimeout(30000 * time.Millisecond),
		},
		TestIamPermissions: []gax.CallOption{
			gax.WithTimeout(30000 * time.Millisecond),
		},
		GeneratePackagesSummary: []gax.CallOption{},
	}
}

func defaultContainerAnalysisV1Beta1RESTCallOptions() *ContainerAnalysisV1Beta1CallOptions {
	return &ContainerAnalysisV1Beta1CallOptions{
		SetIamPolicy: []gax.CallOption{
			gax.WithTimeout(30000 * time.Millisecond),
		},
		GetIamPolicy: []gax.CallOption{
			gax.WithTimeout(30000 * time.Millisecond),
		},
		TestIamPermissions: []gax.CallOption{
			gax.WithTimeout(30000 * time.Millisecond),
		},
		GeneratePackagesSummary: []gax.CallOption{},
	}
}

// internalContainerAnalysisV1Beta1Client is an interface that defines the methods available from Container Analysis API.
type internalContainerAnalysisV1Beta1Client interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	SetIamPolicy(context.Context, *iampb.SetIamPolicyRequest, ...gax.CallOption) (*iampb.Policy, error)
	GetIamPolicy(context.Context, *iampb.GetIamPolicyRequest, ...gax.CallOption) (*iampb.Policy, error)
	TestIamPermissions(context.Context, *iampb.TestIamPermissionsRequest, ...gax.CallOption) (*iampb.TestIamPermissionsResponse, error)
	GeneratePackagesSummary(context.Context, *containeranalysispb.GeneratePackagesSummaryRequest, ...gax.CallOption) (*containeranalysispb.PackagesSummaryResponse, error)
}

// ContainerAnalysisV1Beta1Client is a client for interacting with Container Analysis API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Retrieves analysis results of Cloud components such as Docker container
// images. The Container Analysis API is an implementation of the
// Grafeas (at https://grafeas.io) API.
//
// Analysis results are stored as a series of occurrences. An Occurrence
// contains information about a specific analysis instance on a resource. An
// occurrence refers to a Note. A note contains details describing the
// analysis and is generally stored in a separate project, called a Provider.
// Multiple occurrences can refer to the same note.
//
// For example, an SSL vulnerability could affect multiple images. In this case,
// there would be one note for the vulnerability and an occurrence for each
// image with the vulnerability referring to that note.
type ContainerAnalysisV1Beta1Client struct {
	// The internal transport-dependent client.
	internalClient internalContainerAnalysisV1Beta1Client

	// The call options for this service.
	CallOptions *ContainerAnalysisV1Beta1CallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *ContainerAnalysisV1Beta1Client) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *ContainerAnalysisV1Beta1Client) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *ContainerAnalysisV1Beta1Client) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// SetIamPolicy sets the access control policy on the specified note or occurrence.
// Requires containeranalysis.notes.setIamPolicy or
// containeranalysis.occurrences.setIamPolicy permission if the resource is
// a note or an occurrence, respectively.
//
// The resource takes the format projects/[PROJECT_ID]/notes/[NOTE_ID] for
// notes and projects/[PROJECT_ID]/occurrences/[OCCURRENCE_ID] for
// occurrences.
func (c *ContainerAnalysisV1Beta1Client) SetIamPolicy(ctx context.Context, req *iampb.SetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	return c.internalClient.SetIamPolicy(ctx, req, opts...)
}

// GetIamPolicy gets the access control policy for a note or an occurrence resource.
// Requires containeranalysis.notes.setIamPolicy or
// containeranalysis.occurrences.setIamPolicy permission if the resource is
// a note or occurrence, respectively.
//
// The resource takes the format projects/[PROJECT_ID]/notes/[NOTE_ID] for
// notes and projects/[PROJECT_ID]/occurrences/[OCCURRENCE_ID] for
// occurrences.
func (c *ContainerAnalysisV1Beta1Client) GetIamPolicy(ctx context.Context, req *iampb.GetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	return c.internalClient.GetIamPolicy(ctx, req, opts...)
}

// TestIamPermissions returns the permissions that a caller has on the specified note or
// occurrence. Requires list permission on the project (for example,
// containeranalysis.notes.list).
//
// The resource takes the format projects/[PROJECT_ID]/notes/[NOTE_ID] for
// notes and projects/[PROJECT_ID]/occurrences/[OCCURRENCE_ID] for
// occurrences.
func (c *ContainerAnalysisV1Beta1Client) TestIamPermissions(ctx context.Context, req *iampb.TestIamPermissionsRequest, opts ...gax.CallOption) (*iampb.TestIamPermissionsResponse, error) {
	return c.internalClient.TestIamPermissions(ctx, req, opts...)
}

// GeneratePackagesSummary gets a summary of the packages within a given resource.
func (c *ContainerAnalysisV1Beta1Client) GeneratePackagesSummary(ctx context.Context, req *containeranalysispb.GeneratePackagesSummaryRequest, opts ...gax.CallOption) (*containeranalysispb.PackagesSummaryResponse, error) {
	return c.internalClient.GeneratePackagesSummary(ctx, req, opts...)
}

// containerAnalysisV1Beta1GRPCClient is a client for interacting with Container Analysis API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type containerAnalysisV1Beta1GRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing ContainerAnalysisV1Beta1Client
	CallOptions **ContainerAnalysisV1Beta1CallOptions

	// The gRPC API client.
	containerAnalysisV1Beta1Client containeranalysispb.ContainerAnalysisV1Beta1Client

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewContainerAnalysisV1Beta1Client creates a new container analysis v1 beta1 client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Retrieves analysis results of Cloud components such as Docker container
// images. The Container Analysis API is an implementation of the
// Grafeas (at https://grafeas.io) API.
//
// Analysis results are stored as a series of occurrences. An Occurrence
// contains information about a specific analysis instance on a resource. An
// occurrence refers to a Note. A note contains details describing the
// analysis and is generally stored in a separate project, called a Provider.
// Multiple occurrences can refer to the same note.
//
// For example, an SSL vulnerability could affect multiple images. In this case,
// there would be one note for the vulnerability and an occurrence for each
// image with the vulnerability referring to that note.
func NewContainerAnalysisV1Beta1Client(ctx context.Context, opts ...option.ClientOption) (*ContainerAnalysisV1Beta1Client, error) {
	clientOpts := defaultContainerAnalysisV1Beta1GRPCClientOptions()
	if newContainerAnalysisV1Beta1ClientHook != nil {
		hookOpts, err := newContainerAnalysisV1Beta1ClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := ContainerAnalysisV1Beta1Client{CallOptions: defaultContainerAnalysisV1Beta1CallOptions()}

	c := &containerAnalysisV1Beta1GRPCClient{
		connPool:                       connPool,
		containerAnalysisV1Beta1Client: containeranalysispb.NewContainerAnalysisV1Beta1Client(connPool),
		CallOptions:                    &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *containerAnalysisV1Beta1GRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *containerAnalysisV1Beta1GRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *containerAnalysisV1Beta1GRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type containerAnalysisV1Beta1RESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD

	// Points back to the CallOptions field of the containing ContainerAnalysisV1Beta1Client
	CallOptions **ContainerAnalysisV1Beta1CallOptions
}

// NewContainerAnalysisV1Beta1RESTClient creates a new container analysis v1 beta1 rest client.
//
// Retrieves analysis results of Cloud components such as Docker container
// images. The Container Analysis API is an implementation of the
// Grafeas (at https://grafeas.io) API.
//
// Analysis results are stored as a series of occurrences. An Occurrence
// contains information about a specific analysis instance on a resource. An
// occurrence refers to a Note. A note contains details describing the
// analysis and is generally stored in a separate project, called a Provider.
// Multiple occurrences can refer to the same note.
//
// For example, an SSL vulnerability could affect multiple images. In this case,
// there would be one note for the vulnerability and an occurrence for each
// image with the vulnerability referring to that note.
func NewContainerAnalysisV1Beta1RESTClient(ctx context.Context, opts ...option.ClientOption) (*ContainerAnalysisV1Beta1Client, error) {
	clientOpts := append(defaultContainerAnalysisV1Beta1RESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultContainerAnalysisV1Beta1RESTCallOptions()
	c := &containerAnalysisV1Beta1RESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
	}
	c.setGoogleClientInfo()

	return &ContainerAnalysisV1Beta1Client{internalClient: c, CallOptions: callOpts}, nil
}

func defaultContainerAnalysisV1Beta1RESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://containeranalysis.googleapis.com"),
		internaloption.WithDefaultMTLSEndpoint("https://containeranalysis.mtls.googleapis.com"),
		internaloption.WithDefaultAudience("https://containeranalysis.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *containerAnalysisV1Beta1RESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *containerAnalysisV1Beta1RESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *containerAnalysisV1Beta1RESTClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *containerAnalysisV1Beta1GRPCClient) SetIamPolicy(ctx context.Context, req *iampb.SetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource", url.QueryEscape(req.GetResource())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).SetIamPolicy[0:len((*c.CallOptions).SetIamPolicy):len((*c.CallOptions).SetIamPolicy)], opts...)
	var resp *iampb.Policy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.containerAnalysisV1Beta1Client.SetIamPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *containerAnalysisV1Beta1GRPCClient) GetIamPolicy(ctx context.Context, req *iampb.GetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource", url.QueryEscape(req.GetResource())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetIamPolicy[0:len((*c.CallOptions).GetIamPolicy):len((*c.CallOptions).GetIamPolicy)], opts...)
	var resp *iampb.Policy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.containerAnalysisV1Beta1Client.GetIamPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *containerAnalysisV1Beta1GRPCClient) TestIamPermissions(ctx context.Context, req *iampb.TestIamPermissionsRequest, opts ...gax.CallOption) (*iampb.TestIamPermissionsResponse, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource", url.QueryEscape(req.GetResource())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).TestIamPermissions[0:len((*c.CallOptions).TestIamPermissions):len((*c.CallOptions).TestIamPermissions)], opts...)
	var resp *iampb.TestIamPermissionsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.containerAnalysisV1Beta1Client.TestIamPermissions(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *containerAnalysisV1Beta1GRPCClient) GeneratePackagesSummary(ctx context.Context, req *containeranalysispb.GeneratePackagesSummaryRequest, opts ...gax.CallOption) (*containeranalysispb.PackagesSummaryResponse, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GeneratePackagesSummary[0:len((*c.CallOptions).GeneratePackagesSummary):len((*c.CallOptions).GeneratePackagesSummary)], opts...)
	var resp *containeranalysispb.PackagesSummaryResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.containerAnalysisV1Beta1Client.GeneratePackagesSummary(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SetIamPolicy sets the access control policy on the specified note or occurrence.
// Requires containeranalysis.notes.setIamPolicy or
// containeranalysis.occurrences.setIamPolicy permission if the resource is
// a note or an occurrence, respectively.
//
// The resource takes the format projects/[PROJECT_ID]/notes/[NOTE_ID] for
// notes and projects/[PROJECT_ID]/occurrences/[OCCURRENCE_ID] for
// occurrences.
func (c *containerAnalysisV1Beta1RESTClient) SetIamPolicy(ctx context.Context, req *iampb.SetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1beta1/%v:setIamPolicy", req.GetResource())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource", url.QueryEscape(req.GetResource())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).SetIamPolicy[0:len((*c.CallOptions).SetIamPolicy):len((*c.CallOptions).SetIamPolicy)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &iampb.Policy{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := io.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// GetIamPolicy gets the access control policy for a note or an occurrence resource.
// Requires containeranalysis.notes.setIamPolicy or
// containeranalysis.occurrences.setIamPolicy permission if the resource is
// a note or occurrence, respectively.
//
// The resource takes the format projects/[PROJECT_ID]/notes/[NOTE_ID] for
// notes and projects/[PROJECT_ID]/occurrences/[OCCURRENCE_ID] for
// occurrences.
func (c *containerAnalysisV1Beta1RESTClient) GetIamPolicy(ctx context.Context, req *iampb.GetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1beta1/%v:getIamPolicy", req.GetResource())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource", url.QueryEscape(req.GetResource())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).GetIamPolicy[0:len((*c.CallOptions).GetIamPolicy):len((*c.CallOptions).GetIamPolicy)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &iampb.Policy{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := io.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// TestIamPermissions returns the permissions that a caller has on the specified note or
// occurrence. Requires list permission on the project (for example,
// containeranalysis.notes.list).
//
// The resource takes the format projects/[PROJECT_ID]/notes/[NOTE_ID] for
// notes and projects/[PROJECT_ID]/occurrences/[OCCURRENCE_ID] for
// occurrences.
func (c *containerAnalysisV1Beta1RESTClient) TestIamPermissions(ctx context.Context, req *iampb.TestIamPermissionsRequest, opts ...gax.CallOption) (*iampb.TestIamPermissionsResponse, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1beta1/%v:testIamPermissions", req.GetResource())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource", url.QueryEscape(req.GetResource())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).TestIamPermissions[0:len((*c.CallOptions).TestIamPermissions):len((*c.CallOptions).TestIamPermissions)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &iampb.TestIamPermissionsResponse{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := io.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// GeneratePackagesSummary gets a summary of the packages within a given resource.
func (c *containerAnalysisV1Beta1RESTClient) GeneratePackagesSummary(ctx context.Context, req *containeranalysispb.GeneratePackagesSummaryRequest, opts ...gax.CallOption) (*containeranalysispb.PackagesSummaryResponse, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1beta1/%v:generatePackagesSummary", req.GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).GeneratePackagesSummary[0:len((*c.CallOptions).GeneratePackagesSummary):len((*c.CallOptions).GeneratePackagesSummary)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &containeranalysispb.PackagesSummaryResponse{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := io.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}
