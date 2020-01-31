package datafactory

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

// TriggerRunsClient is the the Azure Data Factory V2 management API provides a RESTful set of web services that
// interact with Azure Data Factory V2 services.
type TriggerRunsClient struct {
	BaseClient
}

// NewTriggerRunsClient creates an instance of the TriggerRunsClient client.
func NewTriggerRunsClient(subscriptionID string) TriggerRunsClient {
	return NewTriggerRunsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewTriggerRunsClientWithBaseURI creates an instance of the TriggerRunsClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewTriggerRunsClientWithBaseURI(baseURI string, subscriptionID string) TriggerRunsClient {
	return TriggerRunsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// QueryByFactory query trigger runs.
// Parameters:
// resourceGroupName - the resource group name.
// factoryName - the factory name.
// filterParameters - parameters to filter the pipeline run.
func (client TriggerRunsClient) QueryByFactory(ctx context.Context, resourceGroupName string, factoryName string, filterParameters RunFilterParameters) (result TriggerRunsQueryResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TriggerRunsClient.QueryByFactory")
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
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: factoryName,
			Constraints: []validation.Constraint{{Target: "factoryName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "factoryName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "factoryName", Name: validation.Pattern, Rule: `^[A-Za-z0-9]+(?:-[A-Za-z0-9]+)*$`, Chain: nil}}},
		{TargetValue: filterParameters,
			Constraints: []validation.Constraint{{Target: "filterParameters.LastUpdatedAfter", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "filterParameters.LastUpdatedBefore", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("datafactory.TriggerRunsClient", "QueryByFactory", err.Error())
	}

	req, err := client.QueryByFactoryPreparer(ctx, resourceGroupName, factoryName, filterParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.TriggerRunsClient", "QueryByFactory", nil, "Failure preparing request")
		return
	}

	resp, err := client.QueryByFactorySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datafactory.TriggerRunsClient", "QueryByFactory", resp, "Failure sending request")
		return
	}

	result, err = client.QueryByFactoryResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.TriggerRunsClient", "QueryByFactory", resp, "Failure responding to request")
	}

	return
}

// QueryByFactoryPreparer prepares the QueryByFactory request.
func (client TriggerRunsClient) QueryByFactoryPreparer(ctx context.Context, resourceGroupName string, factoryName string, filterParameters RunFilterParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"factoryName":       autorest.Encode("path", factoryName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/queryTriggerRuns", pathParameters),
		autorest.WithJSON(filterParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// QueryByFactorySender sends the QueryByFactory request. The method will close the
// http.Response Body if it receives an error.
func (client TriggerRunsClient) QueryByFactorySender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// QueryByFactoryResponder handles the response to the QueryByFactory request. The method always
// closes the http.Response Body.
func (client TriggerRunsClient) QueryByFactoryResponder(resp *http.Response) (result TriggerRunsQueryResponse, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Rerun rerun single trigger instance by runId.
// Parameters:
// resourceGroupName - the resource group name.
// factoryName - the factory name.
// triggerName - the trigger name.
// runID - the pipeline run identifier.
func (client TriggerRunsClient) Rerun(ctx context.Context, resourceGroupName string, factoryName string, triggerName string, runID string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TriggerRunsClient.Rerun")
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
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: factoryName,
			Constraints: []validation.Constraint{{Target: "factoryName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "factoryName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "factoryName", Name: validation.Pattern, Rule: `^[A-Za-z0-9]+(?:-[A-Za-z0-9]+)*$`, Chain: nil}}},
		{TargetValue: triggerName,
			Constraints: []validation.Constraint{{Target: "triggerName", Name: validation.MaxLength, Rule: 260, Chain: nil},
				{Target: "triggerName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "triggerName", Name: validation.Pattern, Rule: `^[A-Za-z0-9_][^<>*#.%&:\\+?/]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("datafactory.TriggerRunsClient", "Rerun", err.Error())
	}

	req, err := client.RerunPreparer(ctx, resourceGroupName, factoryName, triggerName, runID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.TriggerRunsClient", "Rerun", nil, "Failure preparing request")
		return
	}

	resp, err := client.RerunSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "datafactory.TriggerRunsClient", "Rerun", resp, "Failure sending request")
		return
	}

	result, err = client.RerunResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.TriggerRunsClient", "Rerun", resp, "Failure responding to request")
	}

	return
}

// RerunPreparer prepares the Rerun request.
func (client TriggerRunsClient) RerunPreparer(ctx context.Context, resourceGroupName string, factoryName string, triggerName string, runID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"factoryName":       autorest.Encode("path", factoryName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"runId":             autorest.Encode("path", runID),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"triggerName":       autorest.Encode("path", triggerName),
	}

	const APIVersion = "2018-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/triggers/{triggerName}/triggerRuns/{runId}/rerun", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RerunSender sends the Rerun request. The method will close the
// http.Response Body if it receives an error.
func (client TriggerRunsClient) RerunSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// RerunResponder handles the response to the Rerun request. The method always
// closes the http.Response Body.
func (client TriggerRunsClient) RerunResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}
