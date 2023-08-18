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

package mapsplatformdatasets

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"time"

	mapsplatformdatasetspb "cloud.google.com/go/maps/mapsplatformdatasets/apiv1alpha/mapsplatformdatasetspb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httptransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var newMapsPlatformDatasetsV1AlphaClientHook clientHook

// MapsPlatformDatasetsV1AlphaCallOptions contains the retry settings for each method of MapsPlatformDatasetsV1AlphaClient.
type MapsPlatformDatasetsV1AlphaCallOptions struct {
	CreateDataset         []gax.CallOption
	UpdateDatasetMetadata []gax.CallOption
	GetDataset            []gax.CallOption
	ListDatasetVersions   []gax.CallOption
	ListDatasets          []gax.CallOption
	DeleteDataset         []gax.CallOption
	DeleteDatasetVersion  []gax.CallOption
}

func defaultMapsPlatformDatasetsV1AlphaGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("mapsplatformdatasets.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("mapsplatformdatasets.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://mapsplatformdatasets.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultMapsPlatformDatasetsV1AlphaCallOptions() *MapsPlatformDatasetsV1AlphaCallOptions {
	return &MapsPlatformDatasetsV1AlphaCallOptions{
		CreateDataset: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		UpdateDatasetMetadata: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		GetDataset: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		ListDatasetVersions: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		ListDatasets: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		DeleteDataset: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		DeleteDatasetVersion: []gax.CallOption{},
	}
}

func defaultMapsPlatformDatasetsV1AlphaRESTCallOptions() *MapsPlatformDatasetsV1AlphaCallOptions {
	return &MapsPlatformDatasetsV1AlphaCallOptions{
		CreateDataset: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		UpdateDatasetMetadata: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		GetDataset: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable)
			}),
		},
		ListDatasetVersions: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable)
			}),
		},
		ListDatasets: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable)
			}),
		},
		DeleteDataset: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		DeleteDatasetVersion: []gax.CallOption{},
	}
}

// internalMapsPlatformDatasetsV1AlphaClient is an interface that defines the methods available from Maps Platform Datasets API.
type internalMapsPlatformDatasetsV1AlphaClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	CreateDataset(context.Context, *mapsplatformdatasetspb.CreateDatasetRequest, ...gax.CallOption) (*mapsplatformdatasetspb.Dataset, error)
	UpdateDatasetMetadata(context.Context, *mapsplatformdatasetspb.UpdateDatasetMetadataRequest, ...gax.CallOption) (*mapsplatformdatasetspb.Dataset, error)
	GetDataset(context.Context, *mapsplatformdatasetspb.GetDatasetRequest, ...gax.CallOption) (*mapsplatformdatasetspb.Dataset, error)
	ListDatasetVersions(context.Context, *mapsplatformdatasetspb.ListDatasetVersionsRequest, ...gax.CallOption) *DatasetIterator
	ListDatasets(context.Context, *mapsplatformdatasetspb.ListDatasetsRequest, ...gax.CallOption) *DatasetIterator
	DeleteDataset(context.Context, *mapsplatformdatasetspb.DeleteDatasetRequest, ...gax.CallOption) error
	DeleteDatasetVersion(context.Context, *mapsplatformdatasetspb.DeleteDatasetVersionRequest, ...gax.CallOption) error
}

// MapsPlatformDatasetsV1AlphaClient is a client for interacting with Maps Platform Datasets API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service definition for the Maps Platform Datasets API.
type MapsPlatformDatasetsV1AlphaClient struct {
	// The internal transport-dependent client.
	internalClient internalMapsPlatformDatasetsV1AlphaClient

	// The call options for this service.
	CallOptions *MapsPlatformDatasetsV1AlphaCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *MapsPlatformDatasetsV1AlphaClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *MapsPlatformDatasetsV1AlphaClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *MapsPlatformDatasetsV1AlphaClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// CreateDataset create a new dataset for the specified project.
func (c *MapsPlatformDatasetsV1AlphaClient) CreateDataset(ctx context.Context, req *mapsplatformdatasetspb.CreateDatasetRequest, opts ...gax.CallOption) (*mapsplatformdatasetspb.Dataset, error) {
	return c.internalClient.CreateDataset(ctx, req, opts...)
}

// UpdateDatasetMetadata update the metadata for the dataset. To update the data use: UploadDataset.
func (c *MapsPlatformDatasetsV1AlphaClient) UpdateDatasetMetadata(ctx context.Context, req *mapsplatformdatasetspb.UpdateDatasetMetadataRequest, opts ...gax.CallOption) (*mapsplatformdatasetspb.Dataset, error) {
	return c.internalClient.UpdateDatasetMetadata(ctx, req, opts...)
}

// GetDataset get the published or latest version of the dataset.
func (c *MapsPlatformDatasetsV1AlphaClient) GetDataset(ctx context.Context, req *mapsplatformdatasetspb.GetDatasetRequest, opts ...gax.CallOption) (*mapsplatformdatasetspb.Dataset, error) {
	return c.internalClient.GetDataset(ctx, req, opts...)
}

// ListDatasetVersions list all the versions of a dataset.
func (c *MapsPlatformDatasetsV1AlphaClient) ListDatasetVersions(ctx context.Context, req *mapsplatformdatasetspb.ListDatasetVersionsRequest, opts ...gax.CallOption) *DatasetIterator {
	return c.internalClient.ListDatasetVersions(ctx, req, opts...)
}

// ListDatasets list all the datasets for the specified project.
func (c *MapsPlatformDatasetsV1AlphaClient) ListDatasets(ctx context.Context, req *mapsplatformdatasetspb.ListDatasetsRequest, opts ...gax.CallOption) *DatasetIterator {
	return c.internalClient.ListDatasets(ctx, req, opts...)
}

// DeleteDataset delete the specified dataset and optionally all its corresponding
// versions.
func (c *MapsPlatformDatasetsV1AlphaClient) DeleteDataset(ctx context.Context, req *mapsplatformdatasetspb.DeleteDatasetRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteDataset(ctx, req, opts...)
}

// DeleteDatasetVersion delete a specific version of the dataset.
func (c *MapsPlatformDatasetsV1AlphaClient) DeleteDatasetVersion(ctx context.Context, req *mapsplatformdatasetspb.DeleteDatasetVersionRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteDatasetVersion(ctx, req, opts...)
}

// mapsPlatformDatasetsV1AlphaGRPCClient is a client for interacting with Maps Platform Datasets API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type mapsPlatformDatasetsV1AlphaGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing MapsPlatformDatasetsV1AlphaClient
	CallOptions **MapsPlatformDatasetsV1AlphaCallOptions

	// The gRPC API client.
	mapsPlatformDatasetsV1AlphaClient mapsplatformdatasetspb.MapsPlatformDatasetsV1AlphaClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewMapsPlatformDatasetsV1AlphaClient creates a new maps platform datasets v1 alpha client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service definition for the Maps Platform Datasets API.
func NewMapsPlatformDatasetsV1AlphaClient(ctx context.Context, opts ...option.ClientOption) (*MapsPlatformDatasetsV1AlphaClient, error) {
	clientOpts := defaultMapsPlatformDatasetsV1AlphaGRPCClientOptions()
	if newMapsPlatformDatasetsV1AlphaClientHook != nil {
		hookOpts, err := newMapsPlatformDatasetsV1AlphaClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := MapsPlatformDatasetsV1AlphaClient{CallOptions: defaultMapsPlatformDatasetsV1AlphaCallOptions()}

	c := &mapsPlatformDatasetsV1AlphaGRPCClient{
		connPool:                          connPool,
		mapsPlatformDatasetsV1AlphaClient: mapsplatformdatasetspb.NewMapsPlatformDatasetsV1AlphaClient(connPool),
		CallOptions:                       &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *mapsPlatformDatasetsV1AlphaGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *mapsPlatformDatasetsV1AlphaGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{"x-goog-api-client", gax.XGoogHeader(kv...)}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *mapsPlatformDatasetsV1AlphaGRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type mapsPlatformDatasetsV1AlphaRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* headers to be sent with each request.
	xGoogHeaders []string

	// Points back to the CallOptions field of the containing MapsPlatformDatasetsV1AlphaClient
	CallOptions **MapsPlatformDatasetsV1AlphaCallOptions
}

// NewMapsPlatformDatasetsV1AlphaRESTClient creates a new maps platform datasets v1 alpha rest client.
//
// Service definition for the Maps Platform Datasets API.
func NewMapsPlatformDatasetsV1AlphaRESTClient(ctx context.Context, opts ...option.ClientOption) (*MapsPlatformDatasetsV1AlphaClient, error) {
	clientOpts := append(defaultMapsPlatformDatasetsV1AlphaRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultMapsPlatformDatasetsV1AlphaRESTCallOptions()
	c := &mapsPlatformDatasetsV1AlphaRESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
	}
	c.setGoogleClientInfo()

	return &MapsPlatformDatasetsV1AlphaClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultMapsPlatformDatasetsV1AlphaRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://mapsplatformdatasets.googleapis.com"),
		internaloption.WithDefaultMTLSEndpoint("https://mapsplatformdatasets.mtls.googleapis.com"),
		internaloption.WithDefaultAudience("https://mapsplatformdatasets.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *mapsPlatformDatasetsV1AlphaRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogHeaders = []string{"x-goog-api-client", gax.XGoogHeader(kv...)}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *mapsPlatformDatasetsV1AlphaRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *mapsPlatformDatasetsV1AlphaRESTClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *mapsPlatformDatasetsV1AlphaGRPCClient) CreateDataset(ctx context.Context, req *mapsplatformdatasetspb.CreateDatasetRequest, opts ...gax.CallOption) (*mapsplatformdatasetspb.Dataset, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).CreateDataset[0:len((*c.CallOptions).CreateDataset):len((*c.CallOptions).CreateDataset)], opts...)
	var resp *mapsplatformdatasetspb.Dataset
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.mapsPlatformDatasetsV1AlphaClient.CreateDataset(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *mapsPlatformDatasetsV1AlphaGRPCClient) UpdateDatasetMetadata(ctx context.Context, req *mapsplatformdatasetspb.UpdateDatasetMetadataRequest, opts ...gax.CallOption) (*mapsplatformdatasetspb.Dataset, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "dataset.name", url.QueryEscape(req.GetDataset().GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).UpdateDatasetMetadata[0:len((*c.CallOptions).UpdateDatasetMetadata):len((*c.CallOptions).UpdateDatasetMetadata)], opts...)
	var resp *mapsplatformdatasetspb.Dataset
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.mapsPlatformDatasetsV1AlphaClient.UpdateDatasetMetadata(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *mapsPlatformDatasetsV1AlphaGRPCClient) GetDataset(ctx context.Context, req *mapsplatformdatasetspb.GetDatasetRequest, opts ...gax.CallOption) (*mapsplatformdatasetspb.Dataset, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetDataset[0:len((*c.CallOptions).GetDataset):len((*c.CallOptions).GetDataset)], opts...)
	var resp *mapsplatformdatasetspb.Dataset
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.mapsPlatformDatasetsV1AlphaClient.GetDataset(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *mapsPlatformDatasetsV1AlphaGRPCClient) ListDatasetVersions(ctx context.Context, req *mapsplatformdatasetspb.ListDatasetVersionsRequest, opts ...gax.CallOption) *DatasetIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ListDatasetVersions[0:len((*c.CallOptions).ListDatasetVersions):len((*c.CallOptions).ListDatasetVersions)], opts...)
	it := &DatasetIterator{}
	req = proto.Clone(req).(*mapsplatformdatasetspb.ListDatasetVersionsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*mapsplatformdatasetspb.Dataset, string, error) {
		resp := &mapsplatformdatasetspb.ListDatasetVersionsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.mapsPlatformDatasetsV1AlphaClient.ListDatasetVersions(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetDatasets(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

func (c *mapsPlatformDatasetsV1AlphaGRPCClient) ListDatasets(ctx context.Context, req *mapsplatformdatasetspb.ListDatasetsRequest, opts ...gax.CallOption) *DatasetIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ListDatasets[0:len((*c.CallOptions).ListDatasets):len((*c.CallOptions).ListDatasets)], opts...)
	it := &DatasetIterator{}
	req = proto.Clone(req).(*mapsplatformdatasetspb.ListDatasetsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*mapsplatformdatasetspb.Dataset, string, error) {
		resp := &mapsplatformdatasetspb.ListDatasetsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.mapsPlatformDatasetsV1AlphaClient.ListDatasets(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetDatasets(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

func (c *mapsPlatformDatasetsV1AlphaGRPCClient) DeleteDataset(ctx context.Context, req *mapsplatformdatasetspb.DeleteDatasetRequest, opts ...gax.CallOption) error {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).DeleteDataset[0:len((*c.CallOptions).DeleteDataset):len((*c.CallOptions).DeleteDataset)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.mapsPlatformDatasetsV1AlphaClient.DeleteDataset(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *mapsPlatformDatasetsV1AlphaGRPCClient) DeleteDatasetVersion(ctx context.Context, req *mapsplatformdatasetspb.DeleteDatasetVersionRequest, opts ...gax.CallOption) error {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).DeleteDatasetVersion[0:len((*c.CallOptions).DeleteDatasetVersion):len((*c.CallOptions).DeleteDatasetVersion)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.mapsPlatformDatasetsV1AlphaClient.DeleteDatasetVersion(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

// CreateDataset create a new dataset for the specified project.
func (c *mapsPlatformDatasetsV1AlphaRESTClient) CreateDataset(ctx context.Context, req *mapsplatformdatasetspb.CreateDatasetRequest, opts ...gax.CallOption) (*mapsplatformdatasetspb.Dataset, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	body := req.GetDataset()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1alpha/%v/datasets", req.GetParent())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).CreateDataset[0:len((*c.CallOptions).CreateDataset):len((*c.CallOptions).CreateDataset)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &mapsplatformdatasetspb.Dataset{}
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

// UpdateDatasetMetadata update the metadata for the dataset. To update the data use: UploadDataset.
func (c *mapsPlatformDatasetsV1AlphaRESTClient) UpdateDatasetMetadata(ctx context.Context, req *mapsplatformdatasetspb.UpdateDatasetMetadataRequest, opts ...gax.CallOption) (*mapsplatformdatasetspb.Dataset, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	body := req.GetDataset()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1alpha/%v", req.GetDataset().GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")
	if req.GetUpdateMask() != nil {
		updateMask, err := protojson.Marshal(req.GetUpdateMask())
		if err != nil {
			return nil, err
		}
		params.Add("updateMask", string(updateMask[1:len(updateMask)-1]))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "dataset.name", url.QueryEscape(req.GetDataset().GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).UpdateDatasetMetadata[0:len((*c.CallOptions).UpdateDatasetMetadata):len((*c.CallOptions).UpdateDatasetMetadata)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &mapsplatformdatasetspb.Dataset{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("PATCH", baseUrl.String(), bytes.NewReader(jsonReq))
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

// GetDataset get the published or latest version of the dataset.
func (c *mapsPlatformDatasetsV1AlphaRESTClient) GetDataset(ctx context.Context, req *mapsplatformdatasetspb.GetDatasetRequest, opts ...gax.CallOption) (*mapsplatformdatasetspb.Dataset, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1alpha/%v", req.GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")
	if req.GetPublishedUsage() != 0 {
		params.Add("publishedUsage", fmt.Sprintf("%v", req.GetPublishedUsage()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).GetDataset[0:len((*c.CallOptions).GetDataset):len((*c.CallOptions).GetDataset)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &mapsplatformdatasetspb.Dataset{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
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

// ListDatasetVersions list all the versions of a dataset.
func (c *mapsPlatformDatasetsV1AlphaRESTClient) ListDatasetVersions(ctx context.Context, req *mapsplatformdatasetspb.ListDatasetVersionsRequest, opts ...gax.CallOption) *DatasetIterator {
	it := &DatasetIterator{}
	req = proto.Clone(req).(*mapsplatformdatasetspb.ListDatasetVersionsRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*mapsplatformdatasetspb.Dataset, string, error) {
		resp := &mapsplatformdatasetspb.ListDatasetVersionsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		baseUrl, err := url.Parse(c.endpoint)
		if err != nil {
			return nil, "", err
		}
		baseUrl.Path += fmt.Sprintf("/v1alpha/%v:listVersions", req.GetName())

		params := url.Values{}
		params.Add("$alt", "json;enum-encoding=int")
		if req.GetPageSize() != 0 {
			params.Add("pageSize", fmt.Sprintf("%v", req.GetPageSize()))
		}
		if req.GetPageToken() != "" {
			params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
		}

		baseUrl.RawQuery = params.Encode()

		// Build HTTP headers from client and context metadata.
		hds := append(c.xGoogHeaders, "Content-Type", "application/json")
		headers := gax.BuildHeaders(ctx, hds...)
		e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			if settings.Path != "" {
				baseUrl.Path = settings.Path
			}
			httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
			if err != nil {
				return err
			}
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
			return nil, "", e
		}
		it.Response = resp
		return resp.GetDatasets(), resp.GetNextPageToken(), nil
	}

	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

// ListDatasets list all the datasets for the specified project.
func (c *mapsPlatformDatasetsV1AlphaRESTClient) ListDatasets(ctx context.Context, req *mapsplatformdatasetspb.ListDatasetsRequest, opts ...gax.CallOption) *DatasetIterator {
	it := &DatasetIterator{}
	req = proto.Clone(req).(*mapsplatformdatasetspb.ListDatasetsRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*mapsplatformdatasetspb.Dataset, string, error) {
		resp := &mapsplatformdatasetspb.ListDatasetsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		baseUrl, err := url.Parse(c.endpoint)
		if err != nil {
			return nil, "", err
		}
		baseUrl.Path += fmt.Sprintf("/v1alpha/%v/datasets", req.GetParent())

		params := url.Values{}
		params.Add("$alt", "json;enum-encoding=int")
		if req.GetPageSize() != 0 {
			params.Add("pageSize", fmt.Sprintf("%v", req.GetPageSize()))
		}
		if req.GetPageToken() != "" {
			params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
		}

		baseUrl.RawQuery = params.Encode()

		// Build HTTP headers from client and context metadata.
		hds := append(c.xGoogHeaders, "Content-Type", "application/json")
		headers := gax.BuildHeaders(ctx, hds...)
		e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			if settings.Path != "" {
				baseUrl.Path = settings.Path
			}
			httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
			if err != nil {
				return err
			}
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
			return nil, "", e
		}
		it.Response = resp
		return resp.GetDatasets(), resp.GetNextPageToken(), nil
	}

	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

// DeleteDataset delete the specified dataset and optionally all its corresponding
// versions.
func (c *mapsPlatformDatasetsV1AlphaRESTClient) DeleteDataset(ctx context.Context, req *mapsplatformdatasetspb.DeleteDatasetRequest, opts ...gax.CallOption) error {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return err
	}
	baseUrl.Path += fmt.Sprintf("/v1alpha/%v", req.GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")
	if req.GetForce() {
		params.Add("force", fmt.Sprintf("%v", req.GetForce()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	return gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("DELETE", baseUrl.String(), nil)
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

		// Returns nil if there is no error, otherwise wraps
		// the response code and body into a non-nil error
		return googleapi.CheckResponse(httpRsp)
	}, opts...)
}

// DeleteDatasetVersion delete a specific version of the dataset.
func (c *mapsPlatformDatasetsV1AlphaRESTClient) DeleteDatasetVersion(ctx context.Context, req *mapsplatformdatasetspb.DeleteDatasetVersionRequest, opts ...gax.CallOption) error {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return err
	}
	baseUrl.Path += fmt.Sprintf("/v1alpha/%v:deleteVersion", req.GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	return gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("DELETE", baseUrl.String(), nil)
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

		// Returns nil if there is no error, otherwise wraps
		// the response code and body into a non-nil error
		return googleapi.CheckResponse(httpRsp)
	}, opts...)
}

// DatasetIterator manages a stream of *mapsplatformdatasetspb.Dataset.
type DatasetIterator struct {
	items    []*mapsplatformdatasetspb.Dataset
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*mapsplatformdatasetspb.Dataset, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *DatasetIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *DatasetIterator) Next() (*mapsplatformdatasetspb.Dataset, error) {
	var item *mapsplatformdatasetspb.Dataset
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *DatasetIterator) bufLen() int {
	return len(it.items)
}

func (it *DatasetIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
