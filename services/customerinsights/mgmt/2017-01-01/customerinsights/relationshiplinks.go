package customerinsights

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
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// RelationshipLinksClient is the the Azure Customer Insights management API provides a RESTful set of web services
// that interact with Azure Customer Insights service to manage your resources. The API has entities that capture the
// relationship between an end user and the Azure Customer Insights service.
type RelationshipLinksClient struct {
	BaseClient
}

// NewRelationshipLinksClient creates an instance of the RelationshipLinksClient client.
func NewRelationshipLinksClient(subscriptionID string) RelationshipLinksClient {
	return NewRelationshipLinksClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewRelationshipLinksClientWithBaseURI creates an instance of the RelationshipLinksClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewRelationshipLinksClientWithBaseURI(baseURI string, subscriptionID string) RelationshipLinksClient {
	return RelationshipLinksClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates a relationship link or updates an existing relationship link within a hub.
// Parameters:
// resourceGroupName - the name of the resource group.
// hubName - the name of the hub.
// relationshipLinkName - the name of the relationship link.
// parameters - parameters supplied to the CreateOrUpdate relationship link operation.
func (client RelationshipLinksClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, hubName string, relationshipLinkName string, parameters RelationshipLinkResourceFormat) (result RelationshipLinksCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RelationshipLinksClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: relationshipLinkName,
			Constraints: []validation.Constraint{{Target: "relationshipLinkName", Name: validation.MaxLength, Rule: 512, Chain: nil},
				{Target: "relationshipLinkName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "relationshipLinkName", Name: validation.Pattern, Rule: `^[a-zA-Z][a-zA-Z0-9_]+$`, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.RelationshipLinkDefinition", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.RelationshipLinkDefinition.InteractionType", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.RelationshipLinkDefinition.ProfilePropertyReferences", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.RelationshipLinkDefinition.RelatedProfilePropertyReferences", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.RelationshipLinkDefinition.RelationshipName", Name: validation.Null, Rule: true, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("customerinsights.RelationshipLinksClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, hubName, relationshipLinkName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.RelationshipLinksClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.RelationshipLinksClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client RelationshipLinksClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, hubName string, relationshipLinkName string, parameters RelationshipLinkResourceFormat) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hubName":              autorest.Encode("path", hubName),
		"relationshipLinkName": autorest.Encode("path", relationshipLinkName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/relationshipLinks/{relationshipLinkName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client RelationshipLinksClient) CreateOrUpdateSender(req *http.Request) (future RelationshipLinksCreateOrUpdateFuture, err error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req, sd...)
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client RelationshipLinksClient) CreateOrUpdateResponder(resp *http.Response) (result RelationshipLinkResourceFormat, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a relationship link within a hub.
// Parameters:
// resourceGroupName - the name of the resource group.
// hubName - the name of the hub.
// relationshipLinkName - the name of the relationship.
func (client RelationshipLinksClient) Delete(ctx context.Context, resourceGroupName string, hubName string, relationshipLinkName string) (result RelationshipLinksDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RelationshipLinksClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, hubName, relationshipLinkName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.RelationshipLinksClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.RelationshipLinksClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client RelationshipLinksClient) DeletePreparer(ctx context.Context, resourceGroupName string, hubName string, relationshipLinkName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hubName":              autorest.Encode("path", hubName),
		"relationshipLinkName": autorest.Encode("path", relationshipLinkName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/relationshipLinks/{relationshipLinkName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client RelationshipLinksClient) DeleteSender(req *http.Request) (future RelationshipLinksDeleteFuture, err error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req, sd...)
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client RelationshipLinksClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets information about the specified relationship Link.
// Parameters:
// resourceGroupName - the name of the resource group.
// hubName - the name of the hub.
// relationshipLinkName - the name of the relationship link.
func (client RelationshipLinksClient) Get(ctx context.Context, resourceGroupName string, hubName string, relationshipLinkName string) (result RelationshipLinkResourceFormat, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RelationshipLinksClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, hubName, relationshipLinkName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.RelationshipLinksClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "customerinsights.RelationshipLinksClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.RelationshipLinksClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client RelationshipLinksClient) GetPreparer(ctx context.Context, resourceGroupName string, hubName string, relationshipLinkName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hubName":              autorest.Encode("path", hubName),
		"relationshipLinkName": autorest.Encode("path", relationshipLinkName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/relationshipLinks/{relationshipLinkName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client RelationshipLinksClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client RelationshipLinksClient) GetResponder(resp *http.Response) (result RelationshipLinkResourceFormat, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByHub gets all relationship links in the hub.
// Parameters:
// resourceGroupName - the name of the resource group.
// hubName - the name of the hub.
func (client RelationshipLinksClient) ListByHub(ctx context.Context, resourceGroupName string, hubName string) (result RelationshipLinkListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RelationshipLinksClient.ListByHub")
		defer func() {
			sc := -1
			if result.rllr.Response.Response != nil {
				sc = result.rllr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByHubNextResults
	req, err := client.ListByHubPreparer(ctx, resourceGroupName, hubName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.RelationshipLinksClient", "ListByHub", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByHubSender(req)
	if err != nil {
		result.rllr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "customerinsights.RelationshipLinksClient", "ListByHub", resp, "Failure sending request")
		return
	}

	result.rllr, err = client.ListByHubResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.RelationshipLinksClient", "ListByHub", resp, "Failure responding to request")
	}

	return
}

// ListByHubPreparer prepares the ListByHub request.
func (client RelationshipLinksClient) ListByHubPreparer(ctx context.Context, resourceGroupName string, hubName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hubName":           autorest.Encode("path", hubName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/relationshipLinks", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByHubSender sends the ListByHub request. The method will close the
// http.Response Body if it receives an error.
func (client RelationshipLinksClient) ListByHubSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListByHubResponder handles the response to the ListByHub request. The method always
// closes the http.Response Body.
func (client RelationshipLinksClient) ListByHubResponder(resp *http.Response) (result RelationshipLinkListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByHubNextResults retrieves the next set of results, if any.
func (client RelationshipLinksClient) listByHubNextResults(ctx context.Context, lastResults RelationshipLinkListResult) (result RelationshipLinkListResult, err error) {
	req, err := lastResults.relationshipLinkListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "customerinsights.RelationshipLinksClient", "listByHubNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByHubSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "customerinsights.RelationshipLinksClient", "listByHubNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByHubResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.RelationshipLinksClient", "listByHubNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByHubComplete enumerates all values, automatically crossing page boundaries as required.
func (client RelationshipLinksClient) ListByHubComplete(ctx context.Context, resourceGroupName string, hubName string) (result RelationshipLinkListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RelationshipLinksClient.ListByHub")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByHub(ctx, resourceGroupName, hubName)
	return
}
