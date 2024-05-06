// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: ssoready/v1/ssoready.proto

package ssoreadyv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/ssoready/ssoready/internal/gen/ssoready/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion0_1_0

const (
	// SSOReadyServiceName is the fully-qualified name of the SSOReadyService service.
	SSOReadyServiceName = "ssoready.v1.SSOReadyService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// SSOReadyServiceGetSAMLRedirectURLProcedure is the fully-qualified name of the SSOReadyService's
	// GetSAMLRedirectURL RPC.
	SSOReadyServiceGetSAMLRedirectURLProcedure = "/ssoready.v1.SSOReadyService/GetSAMLRedirectURL"
	// SSOReadyServiceRedeemSAMLAccessCodeProcedure is the fully-qualified name of the SSOReadyService's
	// RedeemSAMLAccessCode RPC.
	SSOReadyServiceRedeemSAMLAccessCodeProcedure = "/ssoready.v1.SSOReadyService/RedeemSAMLAccessCode"
	// SSOReadyServiceSignInProcedure is the fully-qualified name of the SSOReadyService's SignIn RPC.
	SSOReadyServiceSignInProcedure = "/ssoready.v1.SSOReadyService/SignIn"
	// SSOReadyServiceWhoamiProcedure is the fully-qualified name of the SSOReadyService's Whoami RPC.
	SSOReadyServiceWhoamiProcedure = "/ssoready.v1.SSOReadyService/Whoami"
	// SSOReadyServiceListEnvironmentsProcedure is the fully-qualified name of the SSOReadyService's
	// ListEnvironments RPC.
	SSOReadyServiceListEnvironmentsProcedure = "/ssoready.v1.SSOReadyService/ListEnvironments"
	// SSOReadyServiceGetEnvironmentProcedure is the fully-qualified name of the SSOReadyService's
	// GetEnvironment RPC.
	SSOReadyServiceGetEnvironmentProcedure = "/ssoready.v1.SSOReadyService/GetEnvironment"
	// SSOReadyServiceUpdateEnvironmentProcedure is the fully-qualified name of the SSOReadyService's
	// UpdateEnvironment RPC.
	SSOReadyServiceUpdateEnvironmentProcedure = "/ssoready.v1.SSOReadyService/UpdateEnvironment"
	// SSOReadyServiceListOrganizationsProcedure is the fully-qualified name of the SSOReadyService's
	// ListOrganizations RPC.
	SSOReadyServiceListOrganizationsProcedure = "/ssoready.v1.SSOReadyService/ListOrganizations"
	// SSOReadyServiceGetOrganizationProcedure is the fully-qualified name of the SSOReadyService's
	// GetOrganization RPC.
	SSOReadyServiceGetOrganizationProcedure = "/ssoready.v1.SSOReadyService/GetOrganization"
	// SSOReadyServiceCreateOrganizationProcedure is the fully-qualified name of the SSOReadyService's
	// CreateOrganization RPC.
	SSOReadyServiceCreateOrganizationProcedure = "/ssoready.v1.SSOReadyService/CreateOrganization"
	// SSOReadyServiceListSAMLConnectionsProcedure is the fully-qualified name of the SSOReadyService's
	// ListSAMLConnections RPC.
	SSOReadyServiceListSAMLConnectionsProcedure = "/ssoready.v1.SSOReadyService/ListSAMLConnections"
	// SSOReadyServiceGetSAMLConnectionProcedure is the fully-qualified name of the SSOReadyService's
	// GetSAMLConnection RPC.
	SSOReadyServiceGetSAMLConnectionProcedure = "/ssoready.v1.SSOReadyService/GetSAMLConnection"
	// SSOReadyServiceCreateSAMLConnectionProcedure is the fully-qualified name of the SSOReadyService's
	// CreateSAMLConnection RPC.
	SSOReadyServiceCreateSAMLConnectionProcedure = "/ssoready.v1.SSOReadyService/CreateSAMLConnection"
	// SSOReadyServiceUpdateSAMLConnectionProcedure is the fully-qualified name of the SSOReadyService's
	// UpdateSAMLConnection RPC.
	SSOReadyServiceUpdateSAMLConnectionProcedure = "/ssoready.v1.SSOReadyService/UpdateSAMLConnection"
	// SSOReadyServiceListSAMLFlowsProcedure is the fully-qualified name of the SSOReadyService's
	// ListSAMLFlows RPC.
	SSOReadyServiceListSAMLFlowsProcedure = "/ssoready.v1.SSOReadyService/ListSAMLFlows"
	// SSOReadyServiceGetSAMLFlowProcedure is the fully-qualified name of the SSOReadyService's
	// GetSAMLFlow RPC.
	SSOReadyServiceGetSAMLFlowProcedure = "/ssoready.v1.SSOReadyService/GetSAMLFlow"
)

// SSOReadyServiceClient is a client for the ssoready.v1.SSOReadyService service.
type SSOReadyServiceClient interface {
	GetSAMLRedirectURL(context.Context, *connect.Request[v1.GetSAMLRedirectURLRequest]) (*connect.Response[v1.GetSAMLRedirectURLResponse], error)
	RedeemSAMLAccessCode(context.Context, *connect.Request[v1.RedeemSAMLAccessCodeRequest]) (*connect.Response[v1.RedeemSAMLAccessCodeResponse], error)
	SignIn(context.Context, *connect.Request[v1.SignInRequest]) (*connect.Response[v1.SignInResponse], error)
	Whoami(context.Context, *connect.Request[v1.WhoamiRequest]) (*connect.Response[v1.WhoamiResponse], error)
	ListEnvironments(context.Context, *connect.Request[v1.ListEnvironmentsRequest]) (*connect.Response[v1.ListEnvironmentsResponse], error)
	GetEnvironment(context.Context, *connect.Request[v1.GetEnvironmentRequest]) (*connect.Response[v1.Environment], error)
	UpdateEnvironment(context.Context, *connect.Request[v1.UpdateEnvironmentRequest]) (*connect.Response[v1.Environment], error)
	ListOrganizations(context.Context, *connect.Request[v1.ListOrganizationsRequest]) (*connect.Response[v1.ListOrganizationsResponse], error)
	GetOrganization(context.Context, *connect.Request[v1.GetOrganizationRequest]) (*connect.Response[v1.Organization], error)
	CreateOrganization(context.Context, *connect.Request[v1.CreateOrganizationRequest]) (*connect.Response[v1.Organization], error)
	ListSAMLConnections(context.Context, *connect.Request[v1.ListSAMLConnectionsRequest]) (*connect.Response[v1.ListSAMLConnectionsResponse], error)
	GetSAMLConnection(context.Context, *connect.Request[v1.GetSAMLConnectionRequest]) (*connect.Response[v1.SAMLConnection], error)
	CreateSAMLConnection(context.Context, *connect.Request[v1.CreateSAMLConnectionRequest]) (*connect.Response[v1.SAMLConnection], error)
	UpdateSAMLConnection(context.Context, *connect.Request[v1.UpdateSAMLConnectionRequest]) (*connect.Response[v1.SAMLConnection], error)
	ListSAMLFlows(context.Context, *connect.Request[v1.ListSAMLFlowsRequest]) (*connect.Response[v1.ListSAMLFlowsResponse], error)
	GetSAMLFlow(context.Context, *connect.Request[v1.GetSAMLFlowRequest]) (*connect.Response[v1.SAMLFlow], error)
}

// NewSSOReadyServiceClient constructs a client for the ssoready.v1.SSOReadyService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewSSOReadyServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) SSOReadyServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &sSOReadyServiceClient{
		getSAMLRedirectURL: connect.NewClient[v1.GetSAMLRedirectURLRequest, v1.GetSAMLRedirectURLResponse](
			httpClient,
			baseURL+SSOReadyServiceGetSAMLRedirectURLProcedure,
			opts...,
		),
		redeemSAMLAccessCode: connect.NewClient[v1.RedeemSAMLAccessCodeRequest, v1.RedeemSAMLAccessCodeResponse](
			httpClient,
			baseURL+SSOReadyServiceRedeemSAMLAccessCodeProcedure,
			opts...,
		),
		signIn: connect.NewClient[v1.SignInRequest, v1.SignInResponse](
			httpClient,
			baseURL+SSOReadyServiceSignInProcedure,
			opts...,
		),
		whoami: connect.NewClient[v1.WhoamiRequest, v1.WhoamiResponse](
			httpClient,
			baseURL+SSOReadyServiceWhoamiProcedure,
			opts...,
		),
		listEnvironments: connect.NewClient[v1.ListEnvironmentsRequest, v1.ListEnvironmentsResponse](
			httpClient,
			baseURL+SSOReadyServiceListEnvironmentsProcedure,
			opts...,
		),
		getEnvironment: connect.NewClient[v1.GetEnvironmentRequest, v1.Environment](
			httpClient,
			baseURL+SSOReadyServiceGetEnvironmentProcedure,
			opts...,
		),
		updateEnvironment: connect.NewClient[v1.UpdateEnvironmentRequest, v1.Environment](
			httpClient,
			baseURL+SSOReadyServiceUpdateEnvironmentProcedure,
			opts...,
		),
		listOrganizations: connect.NewClient[v1.ListOrganizationsRequest, v1.ListOrganizationsResponse](
			httpClient,
			baseURL+SSOReadyServiceListOrganizationsProcedure,
			opts...,
		),
		getOrganization: connect.NewClient[v1.GetOrganizationRequest, v1.Organization](
			httpClient,
			baseURL+SSOReadyServiceGetOrganizationProcedure,
			opts...,
		),
		createOrganization: connect.NewClient[v1.CreateOrganizationRequest, v1.Organization](
			httpClient,
			baseURL+SSOReadyServiceCreateOrganizationProcedure,
			opts...,
		),
		listSAMLConnections: connect.NewClient[v1.ListSAMLConnectionsRequest, v1.ListSAMLConnectionsResponse](
			httpClient,
			baseURL+SSOReadyServiceListSAMLConnectionsProcedure,
			opts...,
		),
		getSAMLConnection: connect.NewClient[v1.GetSAMLConnectionRequest, v1.SAMLConnection](
			httpClient,
			baseURL+SSOReadyServiceGetSAMLConnectionProcedure,
			opts...,
		),
		createSAMLConnection: connect.NewClient[v1.CreateSAMLConnectionRequest, v1.SAMLConnection](
			httpClient,
			baseURL+SSOReadyServiceCreateSAMLConnectionProcedure,
			opts...,
		),
		updateSAMLConnection: connect.NewClient[v1.UpdateSAMLConnectionRequest, v1.SAMLConnection](
			httpClient,
			baseURL+SSOReadyServiceUpdateSAMLConnectionProcedure,
			opts...,
		),
		listSAMLFlows: connect.NewClient[v1.ListSAMLFlowsRequest, v1.ListSAMLFlowsResponse](
			httpClient,
			baseURL+SSOReadyServiceListSAMLFlowsProcedure,
			opts...,
		),
		getSAMLFlow: connect.NewClient[v1.GetSAMLFlowRequest, v1.SAMLFlow](
			httpClient,
			baseURL+SSOReadyServiceGetSAMLFlowProcedure,
			opts...,
		),
	}
}

// sSOReadyServiceClient implements SSOReadyServiceClient.
type sSOReadyServiceClient struct {
	getSAMLRedirectURL   *connect.Client[v1.GetSAMLRedirectURLRequest, v1.GetSAMLRedirectURLResponse]
	redeemSAMLAccessCode *connect.Client[v1.RedeemSAMLAccessCodeRequest, v1.RedeemSAMLAccessCodeResponse]
	signIn               *connect.Client[v1.SignInRequest, v1.SignInResponse]
	whoami               *connect.Client[v1.WhoamiRequest, v1.WhoamiResponse]
	listEnvironments     *connect.Client[v1.ListEnvironmentsRequest, v1.ListEnvironmentsResponse]
	getEnvironment       *connect.Client[v1.GetEnvironmentRequest, v1.Environment]
	updateEnvironment    *connect.Client[v1.UpdateEnvironmentRequest, v1.Environment]
	listOrganizations    *connect.Client[v1.ListOrganizationsRequest, v1.ListOrganizationsResponse]
	getOrganization      *connect.Client[v1.GetOrganizationRequest, v1.Organization]
	createOrganization   *connect.Client[v1.CreateOrganizationRequest, v1.Organization]
	listSAMLConnections  *connect.Client[v1.ListSAMLConnectionsRequest, v1.ListSAMLConnectionsResponse]
	getSAMLConnection    *connect.Client[v1.GetSAMLConnectionRequest, v1.SAMLConnection]
	createSAMLConnection *connect.Client[v1.CreateSAMLConnectionRequest, v1.SAMLConnection]
	updateSAMLConnection *connect.Client[v1.UpdateSAMLConnectionRequest, v1.SAMLConnection]
	listSAMLFlows        *connect.Client[v1.ListSAMLFlowsRequest, v1.ListSAMLFlowsResponse]
	getSAMLFlow          *connect.Client[v1.GetSAMLFlowRequest, v1.SAMLFlow]
}

// GetSAMLRedirectURL calls ssoready.v1.SSOReadyService.GetSAMLRedirectURL.
func (c *sSOReadyServiceClient) GetSAMLRedirectURL(ctx context.Context, req *connect.Request[v1.GetSAMLRedirectURLRequest]) (*connect.Response[v1.GetSAMLRedirectURLResponse], error) {
	return c.getSAMLRedirectURL.CallUnary(ctx, req)
}

// RedeemSAMLAccessCode calls ssoready.v1.SSOReadyService.RedeemSAMLAccessCode.
func (c *sSOReadyServiceClient) RedeemSAMLAccessCode(ctx context.Context, req *connect.Request[v1.RedeemSAMLAccessCodeRequest]) (*connect.Response[v1.RedeemSAMLAccessCodeResponse], error) {
	return c.redeemSAMLAccessCode.CallUnary(ctx, req)
}

// SignIn calls ssoready.v1.SSOReadyService.SignIn.
func (c *sSOReadyServiceClient) SignIn(ctx context.Context, req *connect.Request[v1.SignInRequest]) (*connect.Response[v1.SignInResponse], error) {
	return c.signIn.CallUnary(ctx, req)
}

// Whoami calls ssoready.v1.SSOReadyService.Whoami.
func (c *sSOReadyServiceClient) Whoami(ctx context.Context, req *connect.Request[v1.WhoamiRequest]) (*connect.Response[v1.WhoamiResponse], error) {
	return c.whoami.CallUnary(ctx, req)
}

// ListEnvironments calls ssoready.v1.SSOReadyService.ListEnvironments.
func (c *sSOReadyServiceClient) ListEnvironments(ctx context.Context, req *connect.Request[v1.ListEnvironmentsRequest]) (*connect.Response[v1.ListEnvironmentsResponse], error) {
	return c.listEnvironments.CallUnary(ctx, req)
}

// GetEnvironment calls ssoready.v1.SSOReadyService.GetEnvironment.
func (c *sSOReadyServiceClient) GetEnvironment(ctx context.Context, req *connect.Request[v1.GetEnvironmentRequest]) (*connect.Response[v1.Environment], error) {
	return c.getEnvironment.CallUnary(ctx, req)
}

// UpdateEnvironment calls ssoready.v1.SSOReadyService.UpdateEnvironment.
func (c *sSOReadyServiceClient) UpdateEnvironment(ctx context.Context, req *connect.Request[v1.UpdateEnvironmentRequest]) (*connect.Response[v1.Environment], error) {
	return c.updateEnvironment.CallUnary(ctx, req)
}

// ListOrganizations calls ssoready.v1.SSOReadyService.ListOrganizations.
func (c *sSOReadyServiceClient) ListOrganizations(ctx context.Context, req *connect.Request[v1.ListOrganizationsRequest]) (*connect.Response[v1.ListOrganizationsResponse], error) {
	return c.listOrganizations.CallUnary(ctx, req)
}

// GetOrganization calls ssoready.v1.SSOReadyService.GetOrganization.
func (c *sSOReadyServiceClient) GetOrganization(ctx context.Context, req *connect.Request[v1.GetOrganizationRequest]) (*connect.Response[v1.Organization], error) {
	return c.getOrganization.CallUnary(ctx, req)
}

// CreateOrganization calls ssoready.v1.SSOReadyService.CreateOrganization.
func (c *sSOReadyServiceClient) CreateOrganization(ctx context.Context, req *connect.Request[v1.CreateOrganizationRequest]) (*connect.Response[v1.Organization], error) {
	return c.createOrganization.CallUnary(ctx, req)
}

// ListSAMLConnections calls ssoready.v1.SSOReadyService.ListSAMLConnections.
func (c *sSOReadyServiceClient) ListSAMLConnections(ctx context.Context, req *connect.Request[v1.ListSAMLConnectionsRequest]) (*connect.Response[v1.ListSAMLConnectionsResponse], error) {
	return c.listSAMLConnections.CallUnary(ctx, req)
}

// GetSAMLConnection calls ssoready.v1.SSOReadyService.GetSAMLConnection.
func (c *sSOReadyServiceClient) GetSAMLConnection(ctx context.Context, req *connect.Request[v1.GetSAMLConnectionRequest]) (*connect.Response[v1.SAMLConnection], error) {
	return c.getSAMLConnection.CallUnary(ctx, req)
}

// CreateSAMLConnection calls ssoready.v1.SSOReadyService.CreateSAMLConnection.
func (c *sSOReadyServiceClient) CreateSAMLConnection(ctx context.Context, req *connect.Request[v1.CreateSAMLConnectionRequest]) (*connect.Response[v1.SAMLConnection], error) {
	return c.createSAMLConnection.CallUnary(ctx, req)
}

// UpdateSAMLConnection calls ssoready.v1.SSOReadyService.UpdateSAMLConnection.
func (c *sSOReadyServiceClient) UpdateSAMLConnection(ctx context.Context, req *connect.Request[v1.UpdateSAMLConnectionRequest]) (*connect.Response[v1.SAMLConnection], error) {
	return c.updateSAMLConnection.CallUnary(ctx, req)
}

// ListSAMLFlows calls ssoready.v1.SSOReadyService.ListSAMLFlows.
func (c *sSOReadyServiceClient) ListSAMLFlows(ctx context.Context, req *connect.Request[v1.ListSAMLFlowsRequest]) (*connect.Response[v1.ListSAMLFlowsResponse], error) {
	return c.listSAMLFlows.CallUnary(ctx, req)
}

// GetSAMLFlow calls ssoready.v1.SSOReadyService.GetSAMLFlow.
func (c *sSOReadyServiceClient) GetSAMLFlow(ctx context.Context, req *connect.Request[v1.GetSAMLFlowRequest]) (*connect.Response[v1.SAMLFlow], error) {
	return c.getSAMLFlow.CallUnary(ctx, req)
}

// SSOReadyServiceHandler is an implementation of the ssoready.v1.SSOReadyService service.
type SSOReadyServiceHandler interface {
	GetSAMLRedirectURL(context.Context, *connect.Request[v1.GetSAMLRedirectURLRequest]) (*connect.Response[v1.GetSAMLRedirectURLResponse], error)
	RedeemSAMLAccessCode(context.Context, *connect.Request[v1.RedeemSAMLAccessCodeRequest]) (*connect.Response[v1.RedeemSAMLAccessCodeResponse], error)
	SignIn(context.Context, *connect.Request[v1.SignInRequest]) (*connect.Response[v1.SignInResponse], error)
	Whoami(context.Context, *connect.Request[v1.WhoamiRequest]) (*connect.Response[v1.WhoamiResponse], error)
	ListEnvironments(context.Context, *connect.Request[v1.ListEnvironmentsRequest]) (*connect.Response[v1.ListEnvironmentsResponse], error)
	GetEnvironment(context.Context, *connect.Request[v1.GetEnvironmentRequest]) (*connect.Response[v1.Environment], error)
	UpdateEnvironment(context.Context, *connect.Request[v1.UpdateEnvironmentRequest]) (*connect.Response[v1.Environment], error)
	ListOrganizations(context.Context, *connect.Request[v1.ListOrganizationsRequest]) (*connect.Response[v1.ListOrganizationsResponse], error)
	GetOrganization(context.Context, *connect.Request[v1.GetOrganizationRequest]) (*connect.Response[v1.Organization], error)
	CreateOrganization(context.Context, *connect.Request[v1.CreateOrganizationRequest]) (*connect.Response[v1.Organization], error)
	ListSAMLConnections(context.Context, *connect.Request[v1.ListSAMLConnectionsRequest]) (*connect.Response[v1.ListSAMLConnectionsResponse], error)
	GetSAMLConnection(context.Context, *connect.Request[v1.GetSAMLConnectionRequest]) (*connect.Response[v1.SAMLConnection], error)
	CreateSAMLConnection(context.Context, *connect.Request[v1.CreateSAMLConnectionRequest]) (*connect.Response[v1.SAMLConnection], error)
	UpdateSAMLConnection(context.Context, *connect.Request[v1.UpdateSAMLConnectionRequest]) (*connect.Response[v1.SAMLConnection], error)
	ListSAMLFlows(context.Context, *connect.Request[v1.ListSAMLFlowsRequest]) (*connect.Response[v1.ListSAMLFlowsResponse], error)
	GetSAMLFlow(context.Context, *connect.Request[v1.GetSAMLFlowRequest]) (*connect.Response[v1.SAMLFlow], error)
}

// NewSSOReadyServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewSSOReadyServiceHandler(svc SSOReadyServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	sSOReadyServiceGetSAMLRedirectURLHandler := connect.NewUnaryHandler(
		SSOReadyServiceGetSAMLRedirectURLProcedure,
		svc.GetSAMLRedirectURL,
		opts...,
	)
	sSOReadyServiceRedeemSAMLAccessCodeHandler := connect.NewUnaryHandler(
		SSOReadyServiceRedeemSAMLAccessCodeProcedure,
		svc.RedeemSAMLAccessCode,
		opts...,
	)
	sSOReadyServiceSignInHandler := connect.NewUnaryHandler(
		SSOReadyServiceSignInProcedure,
		svc.SignIn,
		opts...,
	)
	sSOReadyServiceWhoamiHandler := connect.NewUnaryHandler(
		SSOReadyServiceWhoamiProcedure,
		svc.Whoami,
		opts...,
	)
	sSOReadyServiceListEnvironmentsHandler := connect.NewUnaryHandler(
		SSOReadyServiceListEnvironmentsProcedure,
		svc.ListEnvironments,
		opts...,
	)
	sSOReadyServiceGetEnvironmentHandler := connect.NewUnaryHandler(
		SSOReadyServiceGetEnvironmentProcedure,
		svc.GetEnvironment,
		opts...,
	)
	sSOReadyServiceUpdateEnvironmentHandler := connect.NewUnaryHandler(
		SSOReadyServiceUpdateEnvironmentProcedure,
		svc.UpdateEnvironment,
		opts...,
	)
	sSOReadyServiceListOrganizationsHandler := connect.NewUnaryHandler(
		SSOReadyServiceListOrganizationsProcedure,
		svc.ListOrganizations,
		opts...,
	)
	sSOReadyServiceGetOrganizationHandler := connect.NewUnaryHandler(
		SSOReadyServiceGetOrganizationProcedure,
		svc.GetOrganization,
		opts...,
	)
	sSOReadyServiceCreateOrganizationHandler := connect.NewUnaryHandler(
		SSOReadyServiceCreateOrganizationProcedure,
		svc.CreateOrganization,
		opts...,
	)
	sSOReadyServiceListSAMLConnectionsHandler := connect.NewUnaryHandler(
		SSOReadyServiceListSAMLConnectionsProcedure,
		svc.ListSAMLConnections,
		opts...,
	)
	sSOReadyServiceGetSAMLConnectionHandler := connect.NewUnaryHandler(
		SSOReadyServiceGetSAMLConnectionProcedure,
		svc.GetSAMLConnection,
		opts...,
	)
	sSOReadyServiceCreateSAMLConnectionHandler := connect.NewUnaryHandler(
		SSOReadyServiceCreateSAMLConnectionProcedure,
		svc.CreateSAMLConnection,
		opts...,
	)
	sSOReadyServiceUpdateSAMLConnectionHandler := connect.NewUnaryHandler(
		SSOReadyServiceUpdateSAMLConnectionProcedure,
		svc.UpdateSAMLConnection,
		opts...,
	)
	sSOReadyServiceListSAMLFlowsHandler := connect.NewUnaryHandler(
		SSOReadyServiceListSAMLFlowsProcedure,
		svc.ListSAMLFlows,
		opts...,
	)
	sSOReadyServiceGetSAMLFlowHandler := connect.NewUnaryHandler(
		SSOReadyServiceGetSAMLFlowProcedure,
		svc.GetSAMLFlow,
		opts...,
	)
	return "/ssoready.v1.SSOReadyService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case SSOReadyServiceGetSAMLRedirectURLProcedure:
			sSOReadyServiceGetSAMLRedirectURLHandler.ServeHTTP(w, r)
		case SSOReadyServiceRedeemSAMLAccessCodeProcedure:
			sSOReadyServiceRedeemSAMLAccessCodeHandler.ServeHTTP(w, r)
		case SSOReadyServiceSignInProcedure:
			sSOReadyServiceSignInHandler.ServeHTTP(w, r)
		case SSOReadyServiceWhoamiProcedure:
			sSOReadyServiceWhoamiHandler.ServeHTTP(w, r)
		case SSOReadyServiceListEnvironmentsProcedure:
			sSOReadyServiceListEnvironmentsHandler.ServeHTTP(w, r)
		case SSOReadyServiceGetEnvironmentProcedure:
			sSOReadyServiceGetEnvironmentHandler.ServeHTTP(w, r)
		case SSOReadyServiceUpdateEnvironmentProcedure:
			sSOReadyServiceUpdateEnvironmentHandler.ServeHTTP(w, r)
		case SSOReadyServiceListOrganizationsProcedure:
			sSOReadyServiceListOrganizationsHandler.ServeHTTP(w, r)
		case SSOReadyServiceGetOrganizationProcedure:
			sSOReadyServiceGetOrganizationHandler.ServeHTTP(w, r)
		case SSOReadyServiceCreateOrganizationProcedure:
			sSOReadyServiceCreateOrganizationHandler.ServeHTTP(w, r)
		case SSOReadyServiceListSAMLConnectionsProcedure:
			sSOReadyServiceListSAMLConnectionsHandler.ServeHTTP(w, r)
		case SSOReadyServiceGetSAMLConnectionProcedure:
			sSOReadyServiceGetSAMLConnectionHandler.ServeHTTP(w, r)
		case SSOReadyServiceCreateSAMLConnectionProcedure:
			sSOReadyServiceCreateSAMLConnectionHandler.ServeHTTP(w, r)
		case SSOReadyServiceUpdateSAMLConnectionProcedure:
			sSOReadyServiceUpdateSAMLConnectionHandler.ServeHTTP(w, r)
		case SSOReadyServiceListSAMLFlowsProcedure:
			sSOReadyServiceListSAMLFlowsHandler.ServeHTTP(w, r)
		case SSOReadyServiceGetSAMLFlowProcedure:
			sSOReadyServiceGetSAMLFlowHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedSSOReadyServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedSSOReadyServiceHandler struct{}

func (UnimplementedSSOReadyServiceHandler) GetSAMLRedirectURL(context.Context, *connect.Request[v1.GetSAMLRedirectURLRequest]) (*connect.Response[v1.GetSAMLRedirectURLResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("ssoready.v1.SSOReadyService.GetSAMLRedirectURL is not implemented"))
}

func (UnimplementedSSOReadyServiceHandler) RedeemSAMLAccessCode(context.Context, *connect.Request[v1.RedeemSAMLAccessCodeRequest]) (*connect.Response[v1.RedeemSAMLAccessCodeResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("ssoready.v1.SSOReadyService.RedeemSAMLAccessCode is not implemented"))
}

func (UnimplementedSSOReadyServiceHandler) SignIn(context.Context, *connect.Request[v1.SignInRequest]) (*connect.Response[v1.SignInResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("ssoready.v1.SSOReadyService.SignIn is not implemented"))
}

func (UnimplementedSSOReadyServiceHandler) Whoami(context.Context, *connect.Request[v1.WhoamiRequest]) (*connect.Response[v1.WhoamiResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("ssoready.v1.SSOReadyService.Whoami is not implemented"))
}

func (UnimplementedSSOReadyServiceHandler) ListEnvironments(context.Context, *connect.Request[v1.ListEnvironmentsRequest]) (*connect.Response[v1.ListEnvironmentsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("ssoready.v1.SSOReadyService.ListEnvironments is not implemented"))
}

func (UnimplementedSSOReadyServiceHandler) GetEnvironment(context.Context, *connect.Request[v1.GetEnvironmentRequest]) (*connect.Response[v1.Environment], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("ssoready.v1.SSOReadyService.GetEnvironment is not implemented"))
}

func (UnimplementedSSOReadyServiceHandler) UpdateEnvironment(context.Context, *connect.Request[v1.UpdateEnvironmentRequest]) (*connect.Response[v1.Environment], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("ssoready.v1.SSOReadyService.UpdateEnvironment is not implemented"))
}

func (UnimplementedSSOReadyServiceHandler) ListOrganizations(context.Context, *connect.Request[v1.ListOrganizationsRequest]) (*connect.Response[v1.ListOrganizationsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("ssoready.v1.SSOReadyService.ListOrganizations is not implemented"))
}

func (UnimplementedSSOReadyServiceHandler) GetOrganization(context.Context, *connect.Request[v1.GetOrganizationRequest]) (*connect.Response[v1.Organization], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("ssoready.v1.SSOReadyService.GetOrganization is not implemented"))
}

func (UnimplementedSSOReadyServiceHandler) CreateOrganization(context.Context, *connect.Request[v1.CreateOrganizationRequest]) (*connect.Response[v1.Organization], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("ssoready.v1.SSOReadyService.CreateOrganization is not implemented"))
}

func (UnimplementedSSOReadyServiceHandler) ListSAMLConnections(context.Context, *connect.Request[v1.ListSAMLConnectionsRequest]) (*connect.Response[v1.ListSAMLConnectionsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("ssoready.v1.SSOReadyService.ListSAMLConnections is not implemented"))
}

func (UnimplementedSSOReadyServiceHandler) GetSAMLConnection(context.Context, *connect.Request[v1.GetSAMLConnectionRequest]) (*connect.Response[v1.SAMLConnection], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("ssoready.v1.SSOReadyService.GetSAMLConnection is not implemented"))
}

func (UnimplementedSSOReadyServiceHandler) CreateSAMLConnection(context.Context, *connect.Request[v1.CreateSAMLConnectionRequest]) (*connect.Response[v1.SAMLConnection], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("ssoready.v1.SSOReadyService.CreateSAMLConnection is not implemented"))
}

func (UnimplementedSSOReadyServiceHandler) UpdateSAMLConnection(context.Context, *connect.Request[v1.UpdateSAMLConnectionRequest]) (*connect.Response[v1.SAMLConnection], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("ssoready.v1.SSOReadyService.UpdateSAMLConnection is not implemented"))
}

func (UnimplementedSSOReadyServiceHandler) ListSAMLFlows(context.Context, *connect.Request[v1.ListSAMLFlowsRequest]) (*connect.Response[v1.ListSAMLFlowsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("ssoready.v1.SSOReadyService.ListSAMLFlows is not implemented"))
}

func (UnimplementedSSOReadyServiceHandler) GetSAMLFlow(context.Context, *connect.Request[v1.GetSAMLFlowRequest]) (*connect.Response[v1.SAMLFlow], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("ssoready.v1.SSOReadyService.GetSAMLFlow is not implemented"))
}
