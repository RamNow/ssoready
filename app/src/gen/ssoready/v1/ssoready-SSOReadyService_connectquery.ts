// @generated by protoc-gen-connect-query v1.3.1 with parameter "target=ts"
// @generated from file ssoready/v1/ssoready.proto (package ssoready.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { MethodKind } from "@bufbuild/protobuf";
import { ListEnvironmentsRequest, ListEnvironmentsResponse, ListOrganizationsRequest, ListOrganizationsResponse, ListSAMLConnectionsRequest, ListSAMLConnectionsResponse, RedeemSAMLAccessTokenRequest, RedeemSAMLAccessTokenResponse, SignInRequest, SignInResponse, WhoamiRequest, WhoamiResponse } from "./ssoready_pb.js";

/**
 * @generated from rpc ssoready.v1.SSOReadyService.RedeemSAMLAccessToken
 */
export const redeemSAMLAccessToken = {
  localName: "redeemSAMLAccessToken",
  name: "RedeemSAMLAccessToken",
  kind: MethodKind.Unary,
  I: RedeemSAMLAccessTokenRequest,
  O: RedeemSAMLAccessTokenResponse,
  service: {
    typeName: "ssoready.v1.SSOReadyService"
  }
} as const;

/**
 * @generated from rpc ssoready.v1.SSOReadyService.SignIn
 */
export const signIn = {
  localName: "signIn",
  name: "SignIn",
  kind: MethodKind.Unary,
  I: SignInRequest,
  O: SignInResponse,
  service: {
    typeName: "ssoready.v1.SSOReadyService"
  }
} as const;

/**
 * @generated from rpc ssoready.v1.SSOReadyService.Whoami
 */
export const whoami = {
  localName: "whoami",
  name: "Whoami",
  kind: MethodKind.Unary,
  I: WhoamiRequest,
  O: WhoamiResponse,
  service: {
    typeName: "ssoready.v1.SSOReadyService"
  }
} as const;

/**
 * @generated from rpc ssoready.v1.SSOReadyService.ListEnvironments
 */
export const listEnvironments = {
  localName: "listEnvironments",
  name: "ListEnvironments",
  kind: MethodKind.Unary,
  I: ListEnvironmentsRequest,
  O: ListEnvironmentsResponse,
  service: {
    typeName: "ssoready.v1.SSOReadyService"
  }
} as const;

/**
 * @generated from rpc ssoready.v1.SSOReadyService.ListOrganizations
 */
export const listOrganizations = {
  localName: "listOrganizations",
  name: "ListOrganizations",
  kind: MethodKind.Unary,
  I: ListOrganizationsRequest,
  O: ListOrganizationsResponse,
  service: {
    typeName: "ssoready.v1.SSOReadyService"
  }
} as const;

/**
 * @generated from rpc ssoready.v1.SSOReadyService.ListSAMLConnections
 */
export const listSAMLConnections = {
  localName: "listSAMLConnections",
  name: "ListSAMLConnections",
  kind: MethodKind.Unary,
  I: ListSAMLConnectionsRequest,
  O: ListSAMLConnectionsResponse,
  service: {
    typeName: "ssoready.v1.SSOReadyService"
  }
} as const;
