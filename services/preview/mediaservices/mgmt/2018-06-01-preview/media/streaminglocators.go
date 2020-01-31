package media

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

// StreamingLocatorsClient is the client for the StreamingLocators methods of the Media service.
type StreamingLocatorsClient struct {
	BaseClient
}

// NewStreamingLocatorsClient creates an instance of the StreamingLocatorsClient client.
func NewStreamingLocatorsClient(subscriptionID string) StreamingLocatorsClient {
	return NewStreamingLocatorsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewStreamingLocatorsClientWithBaseURI creates an instance of the StreamingLocatorsClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewStreamingLocatorsClientWithBaseURI(baseURI string, subscriptionID string) StreamingLocatorsClient {
	return StreamingLocatorsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create create a Streaming Locator in the Media Services account
// Parameters:
// resourceGroupName - the name of the resource group within the Azure subscription.
// accountName - the Media Services account name.
// streamingLocatorName - the Streaming Locator name.
// parameters - the request parameters
func (client StreamingLocatorsClient) Create(ctx context.Context, resourceGroupName string, accountName string, streamingLocatorName string, parameters StreamingLocator) (result StreamingLocator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StreamingLocatorsClient.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.StreamingLocatorProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.StreamingLocatorProperties.AssetName", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.StreamingLocatorProperties.StreamingPolicyName", Name: validation.Null, Rule: true, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("media.StreamingLocatorsClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, accountName, streamingLocatorName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.StreamingLocatorsClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "media.StreamingLocatorsClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.StreamingLocatorsClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client StreamingLocatorsClient) CreatePreparer(ctx context.Context, resourceGroupName string, accountName string, streamingLocatorName string, parameters StreamingLocator) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":          autorest.Encode("path", accountName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"streamingLocatorName": autorest.Encode("path", streamingLocatorName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/streamingLocators/{streamingLocatorName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client StreamingLocatorsClient) CreateSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client StreamingLocatorsClient) CreateResponder(resp *http.Response) (result StreamingLocator, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a Streaming Locator in the Media Services account
// Parameters:
// resourceGroupName - the name of the resource group within the Azure subscription.
// accountName - the Media Services account name.
// streamingLocatorName - the Streaming Locator name.
func (client StreamingLocatorsClient) Delete(ctx context.Context, resourceGroupName string, accountName string, streamingLocatorName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StreamingLocatorsClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, accountName, streamingLocatorName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.StreamingLocatorsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "media.StreamingLocatorsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.StreamingLocatorsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client StreamingLocatorsClient) DeletePreparer(ctx context.Context, resourceGroupName string, accountName string, streamingLocatorName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":          autorest.Encode("path", accountName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"streamingLocatorName": autorest.Encode("path", streamingLocatorName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/streamingLocators/{streamingLocatorName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client StreamingLocatorsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client StreamingLocatorsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get the details of a Streaming Locator in the Media Services account
// Parameters:
// resourceGroupName - the name of the resource group within the Azure subscription.
// accountName - the Media Services account name.
// streamingLocatorName - the Streaming Locator name.
func (client StreamingLocatorsClient) Get(ctx context.Context, resourceGroupName string, accountName string, streamingLocatorName string) (result StreamingLocator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StreamingLocatorsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, accountName, streamingLocatorName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.StreamingLocatorsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "media.StreamingLocatorsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.StreamingLocatorsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client StreamingLocatorsClient) GetPreparer(ctx context.Context, resourceGroupName string, accountName string, streamingLocatorName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":          autorest.Encode("path", accountName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"streamingLocatorName": autorest.Encode("path", streamingLocatorName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/streamingLocators/{streamingLocatorName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client StreamingLocatorsClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client StreamingLocatorsClient) GetResponder(resp *http.Response) (result StreamingLocator, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNotFound),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists the Streaming Locators in the account
// Parameters:
// resourceGroupName - the name of the resource group within the Azure subscription.
// accountName - the Media Services account name.
// filter - restricts the set of items returned.
// top - specifies a non-negative integer n that limits the number of items returned from a collection. The
// service returns the number of available items up to but not greater than the specified value n.
// orderby - specifies the key by which the result collection should be ordered.
func (client StreamingLocatorsClient) List(ctx context.Context, resourceGroupName string, accountName string, filter string, top *int32, orderby string) (result StreamingLocatorCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StreamingLocatorsClient.List")
		defer func() {
			sc := -1
			if result.slc.Response.Response != nil {
				sc = result.slc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, accountName, filter, top, orderby)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.StreamingLocatorsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.slc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "media.StreamingLocatorsClient", "List", resp, "Failure sending request")
		return
	}

	result.slc, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.StreamingLocatorsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client StreamingLocatorsClient) ListPreparer(ctx context.Context, resourceGroupName string, accountName string, filter string, top *int32, orderby string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(orderby) > 0 {
		queryParameters["$orderby"] = autorest.Encode("query", orderby)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/streamingLocators", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client StreamingLocatorsClient) ListSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client StreamingLocatorsClient) ListResponder(resp *http.Response) (result StreamingLocatorCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client StreamingLocatorsClient) listNextResults(ctx context.Context, lastResults StreamingLocatorCollection) (result StreamingLocatorCollection, err error) {
	req, err := lastResults.streamingLocatorCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "media.StreamingLocatorsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "media.StreamingLocatorsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.StreamingLocatorsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client StreamingLocatorsClient) ListComplete(ctx context.Context, resourceGroupName string, accountName string, filter string, top *int32, orderby string) (result StreamingLocatorCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StreamingLocatorsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, accountName, filter, top, orderby)
	return
}

// ListContentKeys list Content Keys used by this Streaming Locator
// Parameters:
// resourceGroupName - the name of the resource group within the Azure subscription.
// accountName - the Media Services account name.
// streamingLocatorName - the Streaming Locator name.
func (client StreamingLocatorsClient) ListContentKeys(ctx context.Context, resourceGroupName string, accountName string, streamingLocatorName string) (result ListContentKeysResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StreamingLocatorsClient.ListContentKeys")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListContentKeysPreparer(ctx, resourceGroupName, accountName, streamingLocatorName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.StreamingLocatorsClient", "ListContentKeys", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListContentKeysSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "media.StreamingLocatorsClient", "ListContentKeys", resp, "Failure sending request")
		return
	}

	result, err = client.ListContentKeysResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.StreamingLocatorsClient", "ListContentKeys", resp, "Failure responding to request")
	}

	return
}

// ListContentKeysPreparer prepares the ListContentKeys request.
func (client StreamingLocatorsClient) ListContentKeysPreparer(ctx context.Context, resourceGroupName string, accountName string, streamingLocatorName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":          autorest.Encode("path", accountName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"streamingLocatorName": autorest.Encode("path", streamingLocatorName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/streamingLocators/{streamingLocatorName}/listContentKeys", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListContentKeysSender sends the ListContentKeys request. The method will close the
// http.Response Body if it receives an error.
func (client StreamingLocatorsClient) ListContentKeysSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListContentKeysResponder handles the response to the ListContentKeys request. The method always
// closes the http.Response Body.
func (client StreamingLocatorsClient) ListContentKeysResponder(resp *http.Response) (result ListContentKeysResponse, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListPaths list Paths supported by this Streaming Locator
// Parameters:
// resourceGroupName - the name of the resource group within the Azure subscription.
// accountName - the Media Services account name.
// streamingLocatorName - the Streaming Locator name.
func (client StreamingLocatorsClient) ListPaths(ctx context.Context, resourceGroupName string, accountName string, streamingLocatorName string) (result ListPathsResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StreamingLocatorsClient.ListPaths")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListPathsPreparer(ctx, resourceGroupName, accountName, streamingLocatorName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.StreamingLocatorsClient", "ListPaths", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListPathsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "media.StreamingLocatorsClient", "ListPaths", resp, "Failure sending request")
		return
	}

	result, err = client.ListPathsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.StreamingLocatorsClient", "ListPaths", resp, "Failure responding to request")
	}

	return
}

// ListPathsPreparer prepares the ListPaths request.
func (client StreamingLocatorsClient) ListPathsPreparer(ctx context.Context, resourceGroupName string, accountName string, streamingLocatorName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":          autorest.Encode("path", accountName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"streamingLocatorName": autorest.Encode("path", streamingLocatorName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/streamingLocators/{streamingLocatorName}/listPaths", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListPathsSender sends the ListPaths request. The method will close the
// http.Response Body if it receives an error.
func (client StreamingLocatorsClient) ListPathsSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListPathsResponder handles the response to the ListPaths request. The method always
// closes the http.Response Body.
func (client StreamingLocatorsClient) ListPathsResponder(resp *http.Response) (result ListPathsResponse, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
