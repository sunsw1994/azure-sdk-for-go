package deploymentmanager

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

// RolloutsClient is the REST APIs for orchestrating deployments using the Azure Deployment Manager (ADM). See
// https://docs.microsoft.com/en-us/azure/azure-resource-manager/deployment-manager-overview for more information.
type RolloutsClient struct {
	BaseClient
}

// NewRolloutsClient creates an instance of the RolloutsClient client.
func NewRolloutsClient(subscriptionID string) RolloutsClient {
	return NewRolloutsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewRolloutsClientWithBaseURI creates an instance of the RolloutsClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewRolloutsClientWithBaseURI(baseURI string, subscriptionID string) RolloutsClient {
	return RolloutsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Cancel only running rollouts can be canceled.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// rolloutName - the rollout name.
func (client RolloutsClient) Cancel(ctx context.Context, resourceGroupName string, rolloutName string) (result Rollout, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RolloutsClient.Cancel")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("deploymentmanager.RolloutsClient", "Cancel", err.Error())
	}

	req, err := client.CancelPreparer(ctx, resourceGroupName, rolloutName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deploymentmanager.RolloutsClient", "Cancel", nil, "Failure preparing request")
		return
	}

	resp, err := client.CancelSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "deploymentmanager.RolloutsClient", "Cancel", resp, "Failure sending request")
		return
	}

	result, err = client.CancelResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deploymentmanager.RolloutsClient", "Cancel", resp, "Failure responding to request")
	}

	return
}

// CancelPreparer prepares the Cancel request.
func (client RolloutsClient) CancelPreparer(ctx context.Context, resourceGroupName string, rolloutName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"rolloutName":       autorest.Encode("path", rolloutName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeploymentManager/rollouts/{rolloutName}/cancel", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CancelSender sends the Cancel request. The method will close the
// http.Response Body if it receives an error.
func (client RolloutsClient) CancelSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// CancelResponder handles the response to the Cancel request. The method always
// closes the http.Response Body.
func (client RolloutsClient) CancelResponder(resp *http.Response) (result Rollout, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CreateOrUpdate this is an asynchronous operation and can be polled to completion using the location header returned
// by this operation.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// rolloutName - the rollout name.
// rolloutRequest - source rollout request object that defines the rollout.
func (client RolloutsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, rolloutName string, rolloutRequest *RolloutRequest) (result RolloutsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RolloutsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: rolloutRequest,
			Constraints: []validation.Constraint{{Target: "rolloutRequest", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "rolloutRequest.Identity", Name: validation.Null, Rule: true,
					Chain: []validation.Constraint{{Target: "rolloutRequest.Identity.Type", Name: validation.Null, Rule: true, Chain: nil},
						{Target: "rolloutRequest.Identity.IdentityIds", Name: validation.Null, Rule: true, Chain: nil},
					}},
					{Target: "rolloutRequest.RolloutRequestProperties", Name: validation.Null, Rule: true,
						Chain: []validation.Constraint{{Target: "rolloutRequest.RolloutRequestProperties.BuildVersion", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "rolloutRequest.RolloutRequestProperties.TargetServiceTopologyID", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "rolloutRequest.RolloutRequestProperties.StepGroups", Name: validation.Null, Rule: true, Chain: nil},
						}},
				}}}}}); err != nil {
		return result, validation.NewError("deploymentmanager.RolloutsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, rolloutName, rolloutRequest)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deploymentmanager.RolloutsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deploymentmanager.RolloutsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client RolloutsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, rolloutName string, rolloutRequest *RolloutRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"rolloutName":       autorest.Encode("path", rolloutName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeploymentManager/rollouts/{rolloutName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if rolloutRequest != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(rolloutRequest))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client RolloutsClient) CreateOrUpdateSender(req *http.Request) (future RolloutsCreateOrUpdateFuture, err error) {
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
func (client RolloutsClient) CreateOrUpdateResponder(resp *http.Response) (result RolloutRequest, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete only rollouts in terminal state can be deleted.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// rolloutName - the rollout name.
func (client RolloutsClient) Delete(ctx context.Context, resourceGroupName string, rolloutName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RolloutsClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("deploymentmanager.RolloutsClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, rolloutName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deploymentmanager.RolloutsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "deploymentmanager.RolloutsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deploymentmanager.RolloutsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client RolloutsClient) DeletePreparer(ctx context.Context, resourceGroupName string, rolloutName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"rolloutName":       autorest.Encode("path", rolloutName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeploymentManager/rollouts/{rolloutName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client RolloutsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client RolloutsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get sends the get request.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// rolloutName - the rollout name.
// retryAttempt - rollout retry attempt ordinal to get the result of. If not specified, result of the latest
// attempt will be returned.
func (client RolloutsClient) Get(ctx context.Context, resourceGroupName string, rolloutName string, retryAttempt *int32) (result Rollout, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RolloutsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("deploymentmanager.RolloutsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, rolloutName, retryAttempt)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deploymentmanager.RolloutsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "deploymentmanager.RolloutsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deploymentmanager.RolloutsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client RolloutsClient) GetPreparer(ctx context.Context, resourceGroupName string, rolloutName string, retryAttempt *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"rolloutName":       autorest.Encode("path", rolloutName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if retryAttempt != nil {
		queryParameters["retryAttempt"] = autorest.Encode("query", *retryAttempt)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeploymentManager/rollouts/{rolloutName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client RolloutsClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client RolloutsClient) GetResponder(resp *http.Response) (result Rollout, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Restart only failed rollouts can be restarted.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// rolloutName - the rollout name.
// skipSucceeded - if true, will skip all succeeded steps so far in the rollout. If false, will execute the
// entire rollout again regardless of the current state of individual resources. Defaults to false if not
// specified.
func (client RolloutsClient) Restart(ctx context.Context, resourceGroupName string, rolloutName string, skipSucceeded *bool) (result Rollout, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RolloutsClient.Restart")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("deploymentmanager.RolloutsClient", "Restart", err.Error())
	}

	req, err := client.RestartPreparer(ctx, resourceGroupName, rolloutName, skipSucceeded)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deploymentmanager.RolloutsClient", "Restart", nil, "Failure preparing request")
		return
	}

	resp, err := client.RestartSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "deploymentmanager.RolloutsClient", "Restart", resp, "Failure sending request")
		return
	}

	result, err = client.RestartResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deploymentmanager.RolloutsClient", "Restart", resp, "Failure responding to request")
	}

	return
}

// RestartPreparer prepares the Restart request.
func (client RolloutsClient) RestartPreparer(ctx context.Context, resourceGroupName string, rolloutName string, skipSucceeded *bool) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"rolloutName":       autorest.Encode("path", rolloutName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if skipSucceeded != nil {
		queryParameters["skipSucceeded"] = autorest.Encode("query", *skipSucceeded)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeploymentManager/rollouts/{rolloutName}/restart", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RestartSender sends the Restart request. The method will close the
// http.Response Body if it receives an error.
func (client RolloutsClient) RestartSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// RestartResponder handles the response to the Restart request. The method always
// closes the http.Response Body.
func (client RolloutsClient) RestartResponder(resp *http.Response) (result Rollout, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
