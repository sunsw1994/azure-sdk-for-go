package synapseapi

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
    "context"
    "$(go-sdk-folder)/services/preview/synapse/2020-02-01-preview/synapse/accesscontrol"
    "github.com/Azure/go-autorest/autorest"
)

        // BaseClientAPI contains the set of methods on the BaseClient type.
        type BaseClientAPI interface {
            CreateRoleAssignment(ctx context.Context, createRoleAssignmentOptions synapse.RoleAssignmentOptions) (result synapse.RoleAssignmentDetails, err error)
            DeleteRoleAssignmentByID(ctx context.Context, roleAssignmentID string) (result autorest.Response, err error)
            GetCallerRoleAssignments(ctx context.Context) (result synapse.ListString, err error)
            GetRoleAssignmentByID(ctx context.Context, roleAssignmentID string) (result synapse.RoleAssignmentDetails, err error)
            GetRoleAssignments(ctx context.Context, roleID string, principalID string, continuationToken string) (result synapse.ListRoleAssignmentDetails, err error)
            GetRoleDefinitionByID(ctx context.Context, roleID string) (result synapse.Role, err error)
            GetRoleDefinitions(ctx context.Context) (result synapse.RolesListResponsePage, err error)
                GetRoleDefinitionsComplete(ctx context.Context) (result synapse.RolesListResponseIterator, err error)
        }

        var _ BaseClientAPI = (*synapse.BaseClient)(nil)
