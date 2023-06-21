// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by aliasgen. DO NOT EDIT.

// Package gsuiteaddons aliases all exported identifiers in package
// "cloud.google.com/go/gsuiteaddons/apiv1/gsuiteaddonspb".
//
// Deprecated: Please use types in: cloud.google.com/go/gsuiteaddons/apiv1/gsuiteaddonspb.
// Please read https://github.com/googleapis/google-cloud-go/blob/main/migration.md
// for more details.
package gsuiteaddons

import (
	src "cloud.google.com/go/gsuiteaddons/apiv1/gsuiteaddonspb"
	grpc "google.golang.org/grpc"
)

// Deprecated: Please use vars in: cloud.google.com/go/gsuiteaddons/apiv1/gsuiteaddonspb
var (
	File_google_cloud_gsuiteaddons_v1_gsuiteaddons_proto = src.File_google_cloud_gsuiteaddons_v1_gsuiteaddons_proto
)

// A Google Workspace Add-on configuration.
//
// Deprecated: Please use types in: cloud.google.com/go/gsuiteaddons/apiv1/gsuiteaddonspb
type AddOns = src.AddOns

// The authorization information used when invoking deployment endpoints.
//
// Deprecated: Please use types in: cloud.google.com/go/gsuiteaddons/apiv1/gsuiteaddonspb
type Authorization = src.Authorization

// Request message to create a deployment.
//
// Deprecated: Please use types in: cloud.google.com/go/gsuiteaddons/apiv1/gsuiteaddonspb
type CreateDeploymentRequest = src.CreateDeploymentRequest

// Request message to delete a deployment.
//
// Deprecated: Please use types in: cloud.google.com/go/gsuiteaddons/apiv1/gsuiteaddonspb
type DeleteDeploymentRequest = src.DeleteDeploymentRequest

// A Google Workspace Add-on deployment
//
// Deprecated: Please use types in: cloud.google.com/go/gsuiteaddons/apiv1/gsuiteaddonspb
type Deployment = src.Deployment

// GSuiteAddOnsClient is the client API for GSuiteAddOns service. For
// semantics around ctx use and closing/ending streaming RPCs, please refer to
// https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
//
// Deprecated: Please use types in: cloud.google.com/go/gsuiteaddons/apiv1/gsuiteaddonspb
type GSuiteAddOnsClient = src.GSuiteAddOnsClient

// GSuiteAddOnsServer is the server API for GSuiteAddOns service.
//
// Deprecated: Please use types in: cloud.google.com/go/gsuiteaddons/apiv1/gsuiteaddonspb
type GSuiteAddOnsServer = src.GSuiteAddOnsServer

// Request message to get Google Workspace Add-ons authorization information.
//
// Deprecated: Please use types in: cloud.google.com/go/gsuiteaddons/apiv1/gsuiteaddonspb
type GetAuthorizationRequest = src.GetAuthorizationRequest

// Request message to get a deployment.
//
// Deprecated: Please use types in: cloud.google.com/go/gsuiteaddons/apiv1/gsuiteaddonspb
type GetDeploymentRequest = src.GetDeploymentRequest

// Request message to get the install status of a developer mode deployment.
//
// Deprecated: Please use types in: cloud.google.com/go/gsuiteaddons/apiv1/gsuiteaddonspb
type GetInstallStatusRequest = src.GetInstallStatusRequest

// Request message to install a developer mode deployment.
//
// Deprecated: Please use types in: cloud.google.com/go/gsuiteaddons/apiv1/gsuiteaddonspb
type InstallDeploymentRequest = src.InstallDeploymentRequest

// Developer mode install status of a deployment
//
// Deprecated: Please use types in: cloud.google.com/go/gsuiteaddons/apiv1/gsuiteaddonspb
type InstallStatus = src.InstallStatus

// Request message to list deployments for a project.
//
// Deprecated: Please use types in: cloud.google.com/go/gsuiteaddons/apiv1/gsuiteaddonspb
type ListDeploymentsRequest = src.ListDeploymentsRequest

// Response message to list deployments.
//
// Deprecated: Please use types in: cloud.google.com/go/gsuiteaddons/apiv1/gsuiteaddonspb
type ListDeploymentsResponse = src.ListDeploymentsResponse

// Request message to create or replace a deployment.
//
// Deprecated: Please use types in: cloud.google.com/go/gsuiteaddons/apiv1/gsuiteaddonspb
type ReplaceDeploymentRequest = src.ReplaceDeploymentRequest

// UnimplementedGSuiteAddOnsServer can be embedded to have forward compatible
// implementations.
//
// Deprecated: Please use types in: cloud.google.com/go/gsuiteaddons/apiv1/gsuiteaddonspb
type UnimplementedGSuiteAddOnsServer = src.UnimplementedGSuiteAddOnsServer

// Request message to uninstall a developer mode deployment.
//
// Deprecated: Please use types in: cloud.google.com/go/gsuiteaddons/apiv1/gsuiteaddonspb
type UninstallDeploymentRequest = src.UninstallDeploymentRequest

// Deprecated: Please use funcs in: cloud.google.com/go/gsuiteaddons/apiv1/gsuiteaddonspb
func NewGSuiteAddOnsClient(cc grpc.ClientConnInterface) GSuiteAddOnsClient {
	return src.NewGSuiteAddOnsClient(cc)
}

// Deprecated: Please use funcs in: cloud.google.com/go/gsuiteaddons/apiv1/gsuiteaddonspb
func RegisterGSuiteAddOnsServer(s *grpc.Server, srv GSuiteAddOnsServer) {
	src.RegisterGSuiteAddOnsServer(s, srv)
}