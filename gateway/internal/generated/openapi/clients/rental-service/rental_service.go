// Package rental_service provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package rental_service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/oapi-codegen/runtime"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// Defines values for RentalResponseStatus.
const (
	CANCELED   RentalResponseStatus = "CANCELED"
	FINISHED   RentalResponseStatus = "FINISHED"
	INPROGRESS RentalResponseStatus = "IN_PROGRESS"
)

// CreateRentalRequest defines model for CreateRentalRequest.
type CreateRentalRequest struct {
	// CarUid UUID автомобиля
	CarUid openapi_types.UUID `json:"carUid"`

	// DateFrom Дата начала аренды
	DateFrom string `json:"dateFrom"`

	// DateTo Дата окончания аренды
	DateTo string `json:"dateTo"`

	// PaymentUid UUID платежа
	PaymentUid openapi_types.UUID `json:"paymentUid"`
}

// ErrorDescription defines model for ErrorDescription.
type ErrorDescription struct {
	Error string `json:"error"`
	Field string `json:"field"`
}

// ErrorResponse defines model for ErrorResponse.
type ErrorResponse struct {
	// Message Информация об ошибке
	Message string `json:"message"`
}

// RentalResponse defines model for RentalResponse.
type RentalResponse struct {
	// CarUid UUID автомобиля
	CarUid openapi_types.UUID `json:"carUid"`

	// DateFrom Дата начала аренды
	DateFrom string `json:"dateFrom"`

	// DateTo Дата окончания аренды
	DateTo string `json:"dateTo"`

	// PaymentUid UUID платежа
	PaymentUid openapi_types.UUID `json:"paymentUid"`

	// RentalUid UUID аренды
	RentalUid openapi_types.UUID `json:"rentalUid"`

	// Status Статус аренды
	Status RentalResponseStatus `json:"status"`
}

// RentalResponseStatus Статус аренды
type RentalResponseStatus string

// ValidationErrorResponse defines model for ValidationErrorResponse.
type ValidationErrorResponse struct {
	// Errors Массив полей с описанием ошибки
	Errors []ErrorDescription `json:"errors"`

	// Message Информация об ошибке
	Message string `json:"message"`
}

// GetUserRentalsParams defines parameters for GetUserRentals.
type GetUserRentalsParams struct {
	// UserName Имя пользователя
	UserName string `form:"user_name" json:"user_name"`
}

// CreateParams defines parameters for Create.
type CreateParams struct {
	// UserName Имя пользователя
	UserName string `form:"user_name" json:"user_name"`
}

// CancelParams defines parameters for Cancel.
type CancelParams struct {
	// UserName Имя пользователя
	UserName string `form:"user_name" json:"user_name"`
}

// GetParams defines parameters for Get.
type GetParams struct {
	// UserName Имя пользователя
	UserName string `form:"user_name" json:"user_name"`
}

// FinishParams defines parameters for Finish.
type FinishParams struct {
	// UserName Имя пользователя
	UserName string `form:"user_name" json:"user_name"`
}

// CreateJSONRequestBody defines body for Create for application/json ContentType.
type CreateJSONRequestBody = CreateRentalRequest

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
	// GetUserRentals request
	GetUserRentals(ctx context.Context, params *GetUserRentalsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateWithBody request with any body
	CreateWithBody(ctx context.Context, params *CreateParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	Create(ctx context.Context, params *CreateParams, body CreateJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// Cancel request
	Cancel(ctx context.Context, rentalUid openapi_types.UUID, params *CancelParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// Get request
	Get(ctx context.Context, rentalUid openapi_types.UUID, params *GetParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// Finish request
	Finish(ctx context.Context, rentalUid openapi_types.UUID, params *FinishParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// Live request
	Live(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetUserRentals(ctx context.Context, params *GetUserRentalsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetUserRentalsRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateWithBody(ctx context.Context, params *CreateParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateRequestWithBody(c.Server, params, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) Create(ctx context.Context, params *CreateParams, body CreateJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateRequest(c.Server, params, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) Cancel(ctx context.Context, rentalUid openapi_types.UUID, params *CancelParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCancelRequest(c.Server, rentalUid, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) Get(ctx context.Context, rentalUid openapi_types.UUID, params *GetParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetRequest(c.Server, rentalUid, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) Finish(ctx context.Context, rentalUid openapi_types.UUID, params *FinishParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewFinishRequest(c.Server, rentalUid, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) Live(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewLiveRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetUserRentalsRequest generates requests for GetUserRentals
func NewGetUserRentalsRequest(server string, params *GetUserRentalsParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v1/rental")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "user_name", runtime.ParamLocationQuery, params.UserName); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

		queryURL.RawQuery = queryValues.Encode()
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewCreateRequest calls the generic Create builder with application/json body
func NewCreateRequest(server string, params *CreateParams, body CreateJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCreateRequestWithBody(server, params, "application/json", bodyReader)
}

// NewCreateRequestWithBody generates requests for Create with any type of body
func NewCreateRequestWithBody(server string, params *CreateParams, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v1/rental")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "user_name", runtime.ParamLocationQuery, params.UserName); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

		queryURL.RawQuery = queryValues.Encode()
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewCancelRequest generates requests for Cancel
func NewCancelRequest(server string, rentalUid openapi_types.UUID, params *CancelParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "rentalUid", runtime.ParamLocationPath, rentalUid)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v1/rental/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "user_name", runtime.ParamLocationQuery, params.UserName); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

		queryURL.RawQuery = queryValues.Encode()
	}

	req, err := http.NewRequest("DELETE", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetRequest generates requests for Get
func NewGetRequest(server string, rentalUid openapi_types.UUID, params *GetParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "rentalUid", runtime.ParamLocationPath, rentalUid)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v1/rental/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "user_name", runtime.ParamLocationQuery, params.UserName); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

		queryURL.RawQuery = queryValues.Encode()
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewFinishRequest generates requests for Finish
func NewFinishRequest(server string, rentalUid openapi_types.UUID, params *FinishParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "rentalUid", runtime.ParamLocationPath, rentalUid)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v1/rental/%s/finish", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "user_name", runtime.ParamLocationQuery, params.UserName); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

		queryURL.RawQuery = queryValues.Encode()
	}

	req, err := http.NewRequest("POST", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewLiveRequest generates requests for Live
func NewLiveRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/manage/health")
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
	// GetUserRentalsWithResponse request
	GetUserRentalsWithResponse(ctx context.Context, params *GetUserRentalsParams, reqEditors ...RequestEditorFn) (*GetUserRentalsResponse, error)

	// CreateWithBodyWithResponse request with any body
	CreateWithBodyWithResponse(ctx context.Context, params *CreateParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateResponse, error)

	CreateWithResponse(ctx context.Context, params *CreateParams, body CreateJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateResponse, error)

	// CancelWithResponse request
	CancelWithResponse(ctx context.Context, rentalUid openapi_types.UUID, params *CancelParams, reqEditors ...RequestEditorFn) (*CancelResponse, error)

	// GetWithResponse request
	GetWithResponse(ctx context.Context, rentalUid openapi_types.UUID, params *GetParams, reqEditors ...RequestEditorFn) (*GetResponse, error)

	// FinishWithResponse request
	FinishWithResponse(ctx context.Context, rentalUid openapi_types.UUID, params *FinishParams, reqEditors ...RequestEditorFn) (*FinishResponse, error)

	// LiveWithResponse request
	LiveWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*LiveResponse, error)
}

type GetUserRentalsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]RentalResponse
}

// Status returns HTTPResponse.Status
func (r GetUserRentalsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetUserRentalsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type CreateResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *RentalResponse
	JSON400      *ValidationErrorResponse
}

// Status returns HTTPResponse.Status
func (r CreateResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreateResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type CancelResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON403      *ErrorResponse
	JSON404      *ErrorResponse
}

// Status returns HTTPResponse.Status
func (r CancelResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CancelResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *RentalResponse
	JSON403      *ErrorResponse
	JSON404      *ErrorResponse
}

// Status returns HTTPResponse.Status
func (r GetResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type FinishResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON404      *ErrorResponse
}

// Status returns HTTPResponse.Status
func (r FinishResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r FinishResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type LiveResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r LiveResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r LiveResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetUserRentalsWithResponse request returning *GetUserRentalsResponse
func (c *ClientWithResponses) GetUserRentalsWithResponse(ctx context.Context, params *GetUserRentalsParams, reqEditors ...RequestEditorFn) (*GetUserRentalsResponse, error) {
	rsp, err := c.GetUserRentals(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetUserRentalsResponse(rsp)
}

// CreateWithBodyWithResponse request with arbitrary body returning *CreateResponse
func (c *ClientWithResponses) CreateWithBodyWithResponse(ctx context.Context, params *CreateParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateResponse, error) {
	rsp, err := c.CreateWithBody(ctx, params, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateResponse(rsp)
}

func (c *ClientWithResponses) CreateWithResponse(ctx context.Context, params *CreateParams, body CreateJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateResponse, error) {
	rsp, err := c.Create(ctx, params, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateResponse(rsp)
}

// CancelWithResponse request returning *CancelResponse
func (c *ClientWithResponses) CancelWithResponse(ctx context.Context, rentalUid openapi_types.UUID, params *CancelParams, reqEditors ...RequestEditorFn) (*CancelResponse, error) {
	rsp, err := c.Cancel(ctx, rentalUid, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCancelResponse(rsp)
}

// GetWithResponse request returning *GetResponse
func (c *ClientWithResponses) GetWithResponse(ctx context.Context, rentalUid openapi_types.UUID, params *GetParams, reqEditors ...RequestEditorFn) (*GetResponse, error) {
	rsp, err := c.Get(ctx, rentalUid, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetResponse(rsp)
}

// FinishWithResponse request returning *FinishResponse
func (c *ClientWithResponses) FinishWithResponse(ctx context.Context, rentalUid openapi_types.UUID, params *FinishParams, reqEditors ...RequestEditorFn) (*FinishResponse, error) {
	rsp, err := c.Finish(ctx, rentalUid, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseFinishResponse(rsp)
}

// LiveWithResponse request returning *LiveResponse
func (c *ClientWithResponses) LiveWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*LiveResponse, error) {
	rsp, err := c.Live(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseLiveResponse(rsp)
}

// ParseGetUserRentalsResponse parses an HTTP response from a GetUserRentalsWithResponse call
func ParseGetUserRentalsResponse(rsp *http.Response) (*GetUserRentalsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetUserRentalsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []RentalResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseCreateResponse parses an HTTP response from a CreateWithResponse call
func ParseCreateResponse(rsp *http.Response) (*CreateResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CreateResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest RentalResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest ValidationErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	}

	return response, nil
}

// ParseCancelResponse parses an HTTP response from a CancelWithResponse call
func ParseCancelResponse(rsp *http.Response) (*CancelResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CancelResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	}

	return response, nil
}

// ParseGetResponse parses an HTTP response from a GetWithResponse call
func ParseGetResponse(rsp *http.Response) (*GetResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest RentalResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	}

	return response, nil
}

// ParseFinishResponse parses an HTTP response from a FinishWithResponse call
func ParseFinishResponse(rsp *http.Response) (*FinishResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &FinishResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	}

	return response, nil
}

// ParseLiveResponse parses an HTTP response from a LiveWithResponse call
func ParseLiveResponse(rsp *http.Response) (*LiveResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &LiveResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}
