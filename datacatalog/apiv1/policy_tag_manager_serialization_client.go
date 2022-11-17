// Copyright 2022 Google LLC
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

package datacatalog

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	datacatalogpb "cloud.google.com/go/datacatalog/apiv1/datacatalogpb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var newPolicyTagManagerSerializationClientHook clientHook

// PolicyTagManagerSerializationCallOptions contains the retry settings for each method of PolicyTagManagerSerializationClient.
type PolicyTagManagerSerializationCallOptions struct {
	ReplaceTaxonomy  []gax.CallOption
	ImportTaxonomies []gax.CallOption
	ExportTaxonomies []gax.CallOption
}

func defaultPolicyTagManagerSerializationGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("datacatalog.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("datacatalog.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://datacatalog.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultPolicyTagManagerSerializationCallOptions() *PolicyTagManagerSerializationCallOptions {
	return &PolicyTagManagerSerializationCallOptions{
		ReplaceTaxonomy:  []gax.CallOption{},
		ImportTaxonomies: []gax.CallOption{},
		ExportTaxonomies: []gax.CallOption{},
	}
}

// internalPolicyTagManagerSerializationClient is an interface that defines the methods available from Google Cloud Data Catalog API.
type internalPolicyTagManagerSerializationClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	ReplaceTaxonomy(context.Context, *datacatalogpb.ReplaceTaxonomyRequest, ...gax.CallOption) (*datacatalogpb.Taxonomy, error)
	ImportTaxonomies(context.Context, *datacatalogpb.ImportTaxonomiesRequest, ...gax.CallOption) (*datacatalogpb.ImportTaxonomiesResponse, error)
	ExportTaxonomies(context.Context, *datacatalogpb.ExportTaxonomiesRequest, ...gax.CallOption) (*datacatalogpb.ExportTaxonomiesResponse, error)
}

// PolicyTagManagerSerializationClient is a client for interacting with Google Cloud Data Catalog API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Policy Tag Manager Serialization API service allows you to manipulate
// your policy tags and taxonomies in a serialized format.
//
// Taxonomy is a hierarchical group of policy tags.
type PolicyTagManagerSerializationClient struct {
	// The internal transport-dependent client.
	internalClient internalPolicyTagManagerSerializationClient

	// The call options for this service.
	CallOptions *PolicyTagManagerSerializationCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *PolicyTagManagerSerializationClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *PolicyTagManagerSerializationClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *PolicyTagManagerSerializationClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// ReplaceTaxonomy replaces (updates) a taxonomy and all its policy tags.
//
// The taxonomy and its entire hierarchy of policy tags must be
// represented literally by SerializedTaxonomy and the nested
// SerializedPolicyTag messages.
//
// This operation automatically does the following:
//
//	Deletes the existing policy tags that are missing from the
//	SerializedPolicyTag.
//
//	Creates policy tags that don’t have resource names. They are considered
//	new.
//
//	Updates policy tags with valid resources names accordingly.
func (c *PolicyTagManagerSerializationClient) ReplaceTaxonomy(ctx context.Context, req *datacatalogpb.ReplaceTaxonomyRequest, opts ...gax.CallOption) (*datacatalogpb.Taxonomy, error) {
	return c.internalClient.ReplaceTaxonomy(ctx, req, opts...)
}

// ImportTaxonomies creates new taxonomies (including their policy tags) in a given project
// by importing from inlined or cross-regional sources.
//
// For a cross-regional source, new taxonomies are created by copying
// from a source in another region.
//
// For an inlined source, taxonomies and policy tags are created in bulk using
// nested protocol buffer structures.
func (c *PolicyTagManagerSerializationClient) ImportTaxonomies(ctx context.Context, req *datacatalogpb.ImportTaxonomiesRequest, opts ...gax.CallOption) (*datacatalogpb.ImportTaxonomiesResponse, error) {
	return c.internalClient.ImportTaxonomies(ctx, req, opts...)
}

// ExportTaxonomies exports taxonomies in the requested type and returns them,
// including their policy tags. The requested taxonomies must belong to the
// same project.
//
// This method generates SerializedTaxonomy protocol buffers with nested
// policy tags that can be used as input for ImportTaxonomies calls.
func (c *PolicyTagManagerSerializationClient) ExportTaxonomies(ctx context.Context, req *datacatalogpb.ExportTaxonomiesRequest, opts ...gax.CallOption) (*datacatalogpb.ExportTaxonomiesResponse, error) {
	return c.internalClient.ExportTaxonomies(ctx, req, opts...)
}

// policyTagManagerSerializationGRPCClient is a client for interacting with Google Cloud Data Catalog API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type policyTagManagerSerializationGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing PolicyTagManagerSerializationClient
	CallOptions **PolicyTagManagerSerializationCallOptions

	// The gRPC API client.
	policyTagManagerSerializationClient datacatalogpb.PolicyTagManagerSerializationClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewPolicyTagManagerSerializationClient creates a new policy tag manager serialization client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Policy Tag Manager Serialization API service allows you to manipulate
// your policy tags and taxonomies in a serialized format.
//
// Taxonomy is a hierarchical group of policy tags.
func NewPolicyTagManagerSerializationClient(ctx context.Context, opts ...option.ClientOption) (*PolicyTagManagerSerializationClient, error) {
	clientOpts := defaultPolicyTagManagerSerializationGRPCClientOptions()
	if newPolicyTagManagerSerializationClientHook != nil {
		hookOpts, err := newPolicyTagManagerSerializationClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := PolicyTagManagerSerializationClient{CallOptions: defaultPolicyTagManagerSerializationCallOptions()}

	c := &policyTagManagerSerializationGRPCClient{
		connPool:                            connPool,
		disableDeadlines:                    disableDeadlines,
		policyTagManagerSerializationClient: datacatalogpb.NewPolicyTagManagerSerializationClient(connPool),
		CallOptions:                         &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *policyTagManagerSerializationGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *policyTagManagerSerializationGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *policyTagManagerSerializationGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *policyTagManagerSerializationGRPCClient) ReplaceTaxonomy(ctx context.Context, req *datacatalogpb.ReplaceTaxonomyRequest, opts ...gax.CallOption) (*datacatalogpb.Taxonomy, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ReplaceTaxonomy[0:len((*c.CallOptions).ReplaceTaxonomy):len((*c.CallOptions).ReplaceTaxonomy)], opts...)
	var resp *datacatalogpb.Taxonomy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.policyTagManagerSerializationClient.ReplaceTaxonomy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *policyTagManagerSerializationGRPCClient) ImportTaxonomies(ctx context.Context, req *datacatalogpb.ImportTaxonomiesRequest, opts ...gax.CallOption) (*datacatalogpb.ImportTaxonomiesResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ImportTaxonomies[0:len((*c.CallOptions).ImportTaxonomies):len((*c.CallOptions).ImportTaxonomies)], opts...)
	var resp *datacatalogpb.ImportTaxonomiesResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.policyTagManagerSerializationClient.ImportTaxonomies(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *policyTagManagerSerializationGRPCClient) ExportTaxonomies(ctx context.Context, req *datacatalogpb.ExportTaxonomiesRequest, opts ...gax.CallOption) (*datacatalogpb.ExportTaxonomiesResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ExportTaxonomies[0:len((*c.CallOptions).ExportTaxonomies):len((*c.CallOptions).ExportTaxonomies)], opts...)
	var resp *datacatalogpb.ExportTaxonomiesResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.policyTagManagerSerializationClient.ExportTaxonomies(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
