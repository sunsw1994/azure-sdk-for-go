package security

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

// AdaptiveApplicationControlsClient is the API spec for Microsoft.Security (Azure Security Center) resource provider
type AdaptiveApplicationControlsClient struct {
	BaseClient
}

// NewAdaptiveApplicationControlsClient creates an instance of the AdaptiveApplicationControlsClient client.
func NewAdaptiveApplicationControlsClient(subscriptionID string, ascLocation string) AdaptiveApplicationControlsClient {
	return NewAdaptiveApplicationControlsClientWithBaseURI(DefaultBaseURI, subscriptionID, ascLocation)
}

// NewAdaptiveApplicationControlsClientWithBaseURI creates an instance of the AdaptiveApplicationControlsClient client
// using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewAdaptiveApplicationControlsClientWithBaseURI(baseURI string, subscriptionID string, ascLocation string) AdaptiveApplicationControlsClient {
	return AdaptiveApplicationControlsClient{NewWithBaseURI(baseURI, subscriptionID, ascLocation)}
}

// Get gets an application control VM/server group.
// Parameters:
// groupName - name of an application control VM/server group
func (client AdaptiveApplicationControlsClient) Get(ctx context.Context, groupName string) (result AppWhitelistingGroup, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AdaptiveApplicationControlsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("security.AdaptiveApplicationControlsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, groupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.AdaptiveApplicationControlsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "security.AdaptiveApplicationControlsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.AdaptiveApplicationControlsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client AdaptiveApplicationControlsClient) GetPreparer(ctx context.Context, groupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"ascLocation":    autorest.Encode("path", client.AscLocation),
		"groupName":      autorest.Encode("path", groupName),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Security/locations/{ascLocation}/applicationWhitelistings/{groupName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client AdaptiveApplicationControlsClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client AdaptiveApplicationControlsClient) GetResponder(resp *http.Response) (result AppWhitelistingGroup, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets a list of application control VM/server groups for the subscription.
// Parameters:
// includePathRecommendations - include the policy rules
// summary - return output in a summarized form
func (client AdaptiveApplicationControlsClient) List(ctx context.Context, includePathRecommendations *bool, summary *bool) (result AppWhitelistingGroups, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AdaptiveApplicationControlsClient.List")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("security.AdaptiveApplicationControlsClient", "List", err.Error())
	}

	req, err := client.ListPreparer(ctx, includePathRecommendations, summary)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.AdaptiveApplicationControlsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "security.AdaptiveApplicationControlsClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.AdaptiveApplicationControlsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client AdaptiveApplicationControlsClient) ListPreparer(ctx context.Context, includePathRecommendations *bool, summary *bool) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if includePathRecommendations != nil {
		queryParameters["includePathRecommendations"] = autorest.Encode("query", *includePathRecommendations)
	}
	if summary != nil {
		queryParameters["summary"] = autorest.Encode("query", *summary)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Security/applicationWhitelistings", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client AdaptiveApplicationControlsClient) ListSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client AdaptiveApplicationControlsClient) ListResponder(resp *http.Response) (result AppWhitelistingGroups, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Put update an application control VM/server group
// Parameters:
// groupName - name of an application control VM/server group
// body - the updated VM/server group data
func (client AdaptiveApplicationControlsClient) Put(ctx context.Context, groupName string, body AppWhitelistingPutGroupData) (result AppWhitelistingGroup, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AdaptiveApplicationControlsClient.Put")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("security.AdaptiveApplicationControlsClient", "Put", err.Error())
	}

	req, err := client.PutPreparer(ctx, groupName, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.AdaptiveApplicationControlsClient", "Put", nil, "Failure preparing request")
		return
	}

	resp, err := client.PutSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "security.AdaptiveApplicationControlsClient", "Put", resp, "Failure sending request")
		return
	}

	result, err = client.PutResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.AdaptiveApplicationControlsClient", "Put", resp, "Failure responding to request")
	}

	return
}

// PutPreparer prepares the Put request.
func (client AdaptiveApplicationControlsClient) PutPreparer(ctx context.Context, groupName string, body AppWhitelistingPutGroupData) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"ascLocation":    autorest.Encode("path", client.AscLocation),
		"groupName":      autorest.Encode("path", groupName),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Security/locations/{ascLocation}/applicationWhitelistings/{groupName}", pathParameters),
		autorest.WithJSON(body),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PutSender sends the Put request. The method will close the
// http.Response Body if it receives an error.
func (client AdaptiveApplicationControlsClient) PutSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// PutResponder handles the response to the Put request. The method always
// closes the http.Response Body.
func (client AdaptiveApplicationControlsClient) PutResponder(resp *http.Response) (result AppWhitelistingGroup, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
