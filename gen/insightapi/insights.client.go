// Package insightapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.10.1 DO NOT EDIT.
package insightapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
)

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// GetHealthCheckStatus request
	GetHealthCheckStatus(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetPackageVersionInsight request
	GetPackageVersionInsight(ctx context.Context, ecosystem string, name string, version string, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetHealthCheckStatus(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetHealthCheckStatusRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetPackageVersionInsight(ctx context.Context, ecosystem string, name string, version string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetPackageVersionInsightRequest(c.Server, ecosystem, name, version)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetHealthCheckStatusRequest generates requests for GetHealthCheckStatus
func NewGetHealthCheckStatusRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/healthz")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetPackageVersionInsightRequest generates requests for GetPackageVersionInsight
func NewGetPackageVersionInsightRequest(server string, ecosystem string, name string, version string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "ecosystem", runtime.ParamLocationPath, ecosystem)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "name", runtime.ParamLocationPath, name)
	if err != nil {
		return nil, err
	}

	var pathParam2 string

	pathParam2, err = runtime.StyleParamWithLocation("simple", false, "version", runtime.ParamLocationPath, version)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/%s/packages/%s/versions/%s", pathParam0, pathParam1, pathParam2)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// GetHealthCheckStatus request
	GetHealthCheckStatusWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetHealthCheckStatusResponse, error)

	// GetPackageVersionInsight request
	GetPackageVersionInsightWithResponse(ctx context.Context, ecosystem string, name string, version string, reqEditors ...RequestEditorFn) (*GetPackageVersionInsightResponse, error)
}

type GetHealthCheckStatusResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r GetHealthCheckStatusResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetHealthCheckStatusResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetPackageVersionInsightResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *PackageVersionInsight
	JSON403      *ApiError
	JSON404      *ApiError
	JSON429      *ApiError
	JSON500      *ApiError
}

// Status returns HTTPResponse.Status
func (r GetPackageVersionInsightResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetPackageVersionInsightResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetHealthCheckStatusWithResponse request returning *GetHealthCheckStatusResponse
func (c *ClientWithResponses) GetHealthCheckStatusWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetHealthCheckStatusResponse, error) {
	rsp, err := c.GetHealthCheckStatus(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetHealthCheckStatusResponse(rsp)
}

// GetPackageVersionInsightWithResponse request returning *GetPackageVersionInsightResponse
func (c *ClientWithResponses) GetPackageVersionInsightWithResponse(ctx context.Context, ecosystem string, name string, version string, reqEditors ...RequestEditorFn) (*GetPackageVersionInsightResponse, error) {
	rsp, err := c.GetPackageVersionInsight(ctx, ecosystem, name, version, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetPackageVersionInsightResponse(rsp)
}

// ParseGetHealthCheckStatusResponse parses an HTTP response from a GetHealthCheckStatusWithResponse call
func ParseGetHealthCheckStatusResponse(rsp *http.Response) (*GetHealthCheckStatusResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetHealthCheckStatusResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// ParseGetPackageVersionInsightResponse parses an HTTP response from a GetPackageVersionInsightWithResponse call
func ParseGetPackageVersionInsightResponse(rsp *http.Response) (*GetPackageVersionInsightResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetPackageVersionInsightResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest PackageVersionInsight
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest ApiError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest ApiError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest ApiError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest ApiError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}