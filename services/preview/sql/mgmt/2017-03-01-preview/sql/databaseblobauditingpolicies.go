package sql

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
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// DatabaseBlobAuditingPoliciesClient is the the Azure SQL Database management API provides a RESTful set of web
// services that interact with Azure SQL Database services to manage your databases. The API enables you to create,
// retrieve, update, and delete databases.
type DatabaseBlobAuditingPoliciesClient struct {
	BaseClient
}

// NewDatabaseBlobAuditingPoliciesClient creates an instance of the DatabaseBlobAuditingPoliciesClient client.
func NewDatabaseBlobAuditingPoliciesClient(subscriptionID string) DatabaseBlobAuditingPoliciesClient {
	return NewDatabaseBlobAuditingPoliciesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewDatabaseBlobAuditingPoliciesClientWithBaseURI creates an instance of the DatabaseBlobAuditingPoliciesClient
// client.
func NewDatabaseBlobAuditingPoliciesClientWithBaseURI(baseURI string, subscriptionID string) DatabaseBlobAuditingPoliciesClient {
	return DatabaseBlobAuditingPoliciesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a database's blob auditing policy.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
// databaseName - the name of the database.
// parameters - the database blob auditing policy.
func (client DatabaseBlobAuditingPoliciesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, databaseName string, parameters DatabaseBlobAuditingPolicy) (result DatabaseBlobAuditingPolicy, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DatabaseBlobAuditingPoliciesClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, serverName, databaseName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.DatabaseBlobAuditingPoliciesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.DatabaseBlobAuditingPoliciesClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.DatabaseBlobAuditingPoliciesClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client DatabaseBlobAuditingPoliciesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, serverName string, databaseName string, parameters DatabaseBlobAuditingPolicy) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"blobAuditingPolicyName": autorest.Encode("path", "default"),
		"databaseName":           autorest.Encode("path", databaseName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"serverName":             autorest.Encode("path", serverName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	parameters.Kind = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/auditingSettings/{blobAuditingPolicyName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client DatabaseBlobAuditingPoliciesClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client DatabaseBlobAuditingPoliciesClient) CreateOrUpdateResponder(resp *http.Response) (result DatabaseBlobAuditingPolicy, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get gets a database's blob auditing policy.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
// databaseName - the name of the database.
func (client DatabaseBlobAuditingPoliciesClient) Get(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result DatabaseBlobAuditingPolicy, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DatabaseBlobAuditingPoliciesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, serverName, databaseName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.DatabaseBlobAuditingPoliciesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.DatabaseBlobAuditingPoliciesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.DatabaseBlobAuditingPoliciesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client DatabaseBlobAuditingPoliciesClient) GetPreparer(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"blobAuditingPolicyName": autorest.Encode("path", "default"),
		"databaseName":           autorest.Encode("path", databaseName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"serverName":             autorest.Encode("path", serverName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/auditingSettings/{blobAuditingPolicyName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client DatabaseBlobAuditingPoliciesClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client DatabaseBlobAuditingPoliciesClient) GetResponder(resp *http.Response) (result DatabaseBlobAuditingPolicy, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}