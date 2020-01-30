package batchai

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

// ClustersClient is the the Azure BatchAI Management API.
type ClustersClient struct {
	BaseClient
}

// NewClustersClient creates an instance of the ClustersClient client.
func NewClustersClient(subscriptionID string) ClustersClient {
	return NewClustersClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewClustersClientWithBaseURI creates an instance of the ClustersClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewClustersClientWithBaseURI(baseURI string, subscriptionID string) ClustersClient {
	return ClustersClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create adds a cluster. A cluster is a collection of compute nodes. Multiple jobs can be run on the same cluster.
// Parameters:
// resourceGroupName - name of the resource group to which the resource belongs.
// clusterName - the name of the cluster within the specified resource group. Cluster names can only contain a
// combination of alphanumeric characters along with dash (-) and underscore (_). The name must be from 1
// through 64 characters long.
// parameters - the parameters to provide for cluster creation.
func (client ClustersClient) Create(ctx context.Context, resourceGroupName string, clusterName string, parameters ClusterCreateParameters) (result ClustersCreateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ClustersClient.Create")
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
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: clusterName,
			Constraints: []validation.Constraint{{Target: "clusterName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "clusterName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "clusterName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.Location", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.ClusterBaseProperties", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "parameters.ClusterBaseProperties.VMSize", Name: validation.Null, Rule: true, Chain: nil},
						{Target: "parameters.ClusterBaseProperties.ScaleSettings", Name: validation.Null, Rule: false,
							Chain: []validation.Constraint{{Target: "parameters.ClusterBaseProperties.ScaleSettings.Manual", Name: validation.Null, Rule: false,
								Chain: []validation.Constraint{{Target: "parameters.ClusterBaseProperties.ScaleSettings.Manual.TargetNodeCount", Name: validation.Null, Rule: true, Chain: nil}}},
								{Target: "parameters.ClusterBaseProperties.ScaleSettings.AutoScale", Name: validation.Null, Rule: false,
									Chain: []validation.Constraint{{Target: "parameters.ClusterBaseProperties.ScaleSettings.AutoScale.MinimumNodeCount", Name: validation.Null, Rule: true, Chain: nil},
										{Target: "parameters.ClusterBaseProperties.ScaleSettings.AutoScale.MaximumNodeCount", Name: validation.Null, Rule: true, Chain: nil},
									}},
							}},
						{Target: "parameters.ClusterBaseProperties.VirtualMachineConfiguration", Name: validation.Null, Rule: false,
							Chain: []validation.Constraint{{Target: "parameters.ClusterBaseProperties.VirtualMachineConfiguration.ImageReference", Name: validation.Null, Rule: false,
								Chain: []validation.Constraint{{Target: "parameters.ClusterBaseProperties.VirtualMachineConfiguration.ImageReference.Publisher", Name: validation.Null, Rule: true, Chain: nil},
									{Target: "parameters.ClusterBaseProperties.VirtualMachineConfiguration.ImageReference.Offer", Name: validation.Null, Rule: true, Chain: nil},
									{Target: "parameters.ClusterBaseProperties.VirtualMachineConfiguration.ImageReference.Sku", Name: validation.Null, Rule: true, Chain: nil},
								}},
							}},
						{Target: "parameters.ClusterBaseProperties.NodeSetup", Name: validation.Null, Rule: false,
							Chain: []validation.Constraint{{Target: "parameters.ClusterBaseProperties.NodeSetup.SetupTask", Name: validation.Null, Rule: false,
								Chain: []validation.Constraint{{Target: "parameters.ClusterBaseProperties.NodeSetup.SetupTask.CommandLine", Name: validation.Null, Rule: true, Chain: nil},
									{Target: "parameters.ClusterBaseProperties.NodeSetup.SetupTask.StdOutErrPathPrefix", Name: validation.Null, Rule: true, Chain: nil},
								}},
								{Target: "parameters.ClusterBaseProperties.NodeSetup.PerformanceCountersSettings", Name: validation.Null, Rule: false,
									Chain: []validation.Constraint{{Target: "parameters.ClusterBaseProperties.NodeSetup.PerformanceCountersSettings.AppInsightsReference", Name: validation.Null, Rule: true,
										Chain: []validation.Constraint{{Target: "parameters.ClusterBaseProperties.NodeSetup.PerformanceCountersSettings.AppInsightsReference.Component", Name: validation.Null, Rule: true,
											Chain: []validation.Constraint{{Target: "parameters.ClusterBaseProperties.NodeSetup.PerformanceCountersSettings.AppInsightsReference.Component.ID", Name: validation.Null, Rule: true, Chain: nil}}},
											{Target: "parameters.ClusterBaseProperties.NodeSetup.PerformanceCountersSettings.AppInsightsReference.InstrumentationKeySecretReference", Name: validation.Null, Rule: false,
												Chain: []validation.Constraint{{Target: "parameters.ClusterBaseProperties.NodeSetup.PerformanceCountersSettings.AppInsightsReference.InstrumentationKeySecretReference.SourceVault", Name: validation.Null, Rule: true,
													Chain: []validation.Constraint{{Target: "parameters.ClusterBaseProperties.NodeSetup.PerformanceCountersSettings.AppInsightsReference.InstrumentationKeySecretReference.SourceVault.ID", Name: validation.Null, Rule: true, Chain: nil}}},
													{Target: "parameters.ClusterBaseProperties.NodeSetup.PerformanceCountersSettings.AppInsightsReference.InstrumentationKeySecretReference.SecretURL", Name: validation.Null, Rule: true, Chain: nil},
												}},
										}},
									}},
							}},
						{Target: "parameters.ClusterBaseProperties.UserAccountSettings", Name: validation.Null, Rule: true,
							Chain: []validation.Constraint{{Target: "parameters.ClusterBaseProperties.UserAccountSettings.AdminUserName", Name: validation.Null, Rule: true, Chain: nil}}},
						{Target: "parameters.ClusterBaseProperties.Subnet", Name: validation.Null, Rule: false,
							Chain: []validation.Constraint{{Target: "parameters.ClusterBaseProperties.Subnet.ID", Name: validation.Null, Rule: true, Chain: nil}}},
					}}}}}); err != nil {
		return result, validation.NewError("batchai.ClustersClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, clusterName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.ClustersClient", "Create", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.ClustersClient", "Create", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client ClustersClient) CreatePreparer(ctx context.Context, resourceGroupName string, clusterName string, parameters ClusterCreateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/clusters/{clusterName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client ClustersClient) CreateSender(req *http.Request) (future ClustersCreateFuture, err error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req, sd...)
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client ClustersClient) CreateResponder(resp *http.Response) (result Cluster, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a Cluster.
// Parameters:
// resourceGroupName - name of the resource group to which the resource belongs.
// clusterName - the name of the cluster within the specified resource group. Cluster names can only contain a
// combination of alphanumeric characters along with dash (-) and underscore (_). The name must be from 1
// through 64 characters long.
func (client ClustersClient) Delete(ctx context.Context, resourceGroupName string, clusterName string) (result ClustersDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ClustersClient.Delete")
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
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: clusterName,
			Constraints: []validation.Constraint{{Target: "clusterName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "clusterName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "clusterName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("batchai.ClustersClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, clusterName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.ClustersClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.ClustersClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ClustersClient) DeletePreparer(ctx context.Context, resourceGroupName string, clusterName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/clusters/{clusterName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ClustersClient) DeleteSender(req *http.Request) (future ClustersDeleteFuture, err error) {
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
func (client ClustersClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets information about the specified Cluster.
// Parameters:
// resourceGroupName - name of the resource group to which the resource belongs.
// clusterName - the name of the cluster within the specified resource group. Cluster names can only contain a
// combination of alphanumeric characters along with dash (-) and underscore (_). The name must be from 1
// through 64 characters long.
func (client ClustersClient) Get(ctx context.Context, resourceGroupName string, clusterName string) (result Cluster, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ClustersClient.Get")
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
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: clusterName,
			Constraints: []validation.Constraint{{Target: "clusterName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "clusterName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "clusterName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("batchai.ClustersClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, clusterName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.ClustersClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "batchai.ClustersClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.ClustersClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ClustersClient) GetPreparer(ctx context.Context, resourceGroupName string, clusterName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/clusters/{clusterName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ClustersClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ClustersClient) GetResponder(resp *http.Response) (result Cluster, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets information about the Clusters associated with the subscription.
// Parameters:
// filter - an OData $filter clause. Used to filter results that are returned in the GET response.
// selectParameter - an OData $select clause. Used to select the properties to be returned in the GET response.
// maxResults - the maximum number of items to return in the response. A maximum of 1000 files can be returned.
func (client ClustersClient) List(ctx context.Context, filter string, selectParameter string, maxResults *int32) (result ClusterListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ClustersClient.List")
		defer func() {
			sc := -1
			if result.clr.Response.Response != nil {
				sc = result.clr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: maxResults,
			Constraints: []validation.Constraint{{Target: "maxResults", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "maxResults", Name: validation.InclusiveMaximum, Rule: int64(1000), Chain: nil},
					{Target: "maxResults", Name: validation.InclusiveMinimum, Rule: int64(1), Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("batchai.ClustersClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, filter, selectParameter, maxResults)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.ClustersClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.clr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "batchai.ClustersClient", "List", resp, "Failure sending request")
		return
	}

	result.clr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.ClustersClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ClustersClient) ListPreparer(ctx context.Context, filter string, selectParameter string, maxResults *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if len(selectParameter) > 0 {
		queryParameters["$select"] = autorest.Encode("query", selectParameter)
	}
	if maxResults != nil {
		queryParameters["maxresults"] = autorest.Encode("query", *maxResults)
	} else {
		queryParameters["maxresults"] = autorest.Encode("query", 1000)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.BatchAI/clusters", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ClustersClient) ListSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ClustersClient) ListResponder(resp *http.Response) (result ClusterListResult, err error) {
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
func (client ClustersClient) listNextResults(ctx context.Context, lastResults ClusterListResult) (result ClusterListResult, err error) {
	req, err := lastResults.clusterListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "batchai.ClustersClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "batchai.ClustersClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.ClustersClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client ClustersClient) ListComplete(ctx context.Context, filter string, selectParameter string, maxResults *int32) (result ClusterListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ClustersClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, filter, selectParameter, maxResults)
	return
}

// ListByResourceGroup gets information about the Clusters associated within the specified resource group.
// Parameters:
// resourceGroupName - name of the resource group to which the resource belongs.
// filter - an OData $filter clause. Used to filter results that are returned in the GET response.
// selectParameter - an OData $select clause. Used to select the properties to be returned in the GET response.
// maxResults - the maximum number of items to return in the response. A maximum of 1000 files can be returned.
func (client ClustersClient) ListByResourceGroup(ctx context.Context, resourceGroupName string, filter string, selectParameter string, maxResults *int32) (result ClusterListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ClustersClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.clr.Response.Response != nil {
				sc = result.clr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: maxResults,
			Constraints: []validation.Constraint{{Target: "maxResults", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "maxResults", Name: validation.InclusiveMaximum, Rule: int64(1000), Chain: nil},
					{Target: "maxResults", Name: validation.InclusiveMinimum, Rule: int64(1), Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("batchai.ClustersClient", "ListByResourceGroup", err.Error())
	}

	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName, filter, selectParameter, maxResults)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.ClustersClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.clr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "batchai.ClustersClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.clr, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.ClustersClient", "ListByResourceGroup", resp, "Failure responding to request")
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client ClustersClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string, filter string, selectParameter string, maxResults *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if len(selectParameter) > 0 {
		queryParameters["$select"] = autorest.Encode("query", selectParameter)
	}
	if maxResults != nil {
		queryParameters["maxresults"] = autorest.Encode("query", *maxResults)
	} else {
		queryParameters["maxresults"] = autorest.Encode("query", 1000)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/clusters", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client ClustersClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client ClustersClient) ListByResourceGroupResponder(resp *http.Response) (result ClusterListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client ClustersClient) listByResourceGroupNextResults(ctx context.Context, lastResults ClusterListResult) (result ClusterListResult, err error) {
	req, err := lastResults.clusterListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "batchai.ClustersClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "batchai.ClustersClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.ClustersClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client ClustersClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, filter string, selectParameter string, maxResults *int32) (result ClusterListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ClustersClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName, filter, selectParameter, maxResults)
	return
}

// ListRemoteLoginInformation get the IP address, port of all the compute nodes in the cluster.
// Parameters:
// resourceGroupName - name of the resource group to which the resource belongs.
// clusterName - the name of the cluster within the specified resource group. Cluster names can only contain a
// combination of alphanumeric characters along with dash (-) and underscore (_). The name must be from 1
// through 64 characters long.
func (client ClustersClient) ListRemoteLoginInformation(ctx context.Context, resourceGroupName string, clusterName string) (result RemoteLoginInformationListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ClustersClient.ListRemoteLoginInformation")
		defer func() {
			sc := -1
			if result.rlilr.Response.Response != nil {
				sc = result.rlilr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: clusterName,
			Constraints: []validation.Constraint{{Target: "clusterName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "clusterName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "clusterName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("batchai.ClustersClient", "ListRemoteLoginInformation", err.Error())
	}

	result.fn = client.listRemoteLoginInformationNextResults
	req, err := client.ListRemoteLoginInformationPreparer(ctx, resourceGroupName, clusterName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.ClustersClient", "ListRemoteLoginInformation", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListRemoteLoginInformationSender(req)
	if err != nil {
		result.rlilr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "batchai.ClustersClient", "ListRemoteLoginInformation", resp, "Failure sending request")
		return
	}

	result.rlilr, err = client.ListRemoteLoginInformationResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.ClustersClient", "ListRemoteLoginInformation", resp, "Failure responding to request")
	}

	return
}

// ListRemoteLoginInformationPreparer prepares the ListRemoteLoginInformation request.
func (client ClustersClient) ListRemoteLoginInformationPreparer(ctx context.Context, resourceGroupName string, clusterName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/clusters/{clusterName}/listRemoteLoginInformation", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListRemoteLoginInformationSender sends the ListRemoteLoginInformation request. The method will close the
// http.Response Body if it receives an error.
func (client ClustersClient) ListRemoteLoginInformationSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListRemoteLoginInformationResponder handles the response to the ListRemoteLoginInformation request. The method always
// closes the http.Response Body.
func (client ClustersClient) ListRemoteLoginInformationResponder(resp *http.Response) (result RemoteLoginInformationListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listRemoteLoginInformationNextResults retrieves the next set of results, if any.
func (client ClustersClient) listRemoteLoginInformationNextResults(ctx context.Context, lastResults RemoteLoginInformationListResult) (result RemoteLoginInformationListResult, err error) {
	req, err := lastResults.remoteLoginInformationListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "batchai.ClustersClient", "listRemoteLoginInformationNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListRemoteLoginInformationSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "batchai.ClustersClient", "listRemoteLoginInformationNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListRemoteLoginInformationResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.ClustersClient", "listRemoteLoginInformationNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListRemoteLoginInformationComplete enumerates all values, automatically crossing page boundaries as required.
func (client ClustersClient) ListRemoteLoginInformationComplete(ctx context.Context, resourceGroupName string, clusterName string) (result RemoteLoginInformationListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ClustersClient.ListRemoteLoginInformation")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListRemoteLoginInformation(ctx, resourceGroupName, clusterName)
	return
}

// Update update the properties of a given cluster.
// Parameters:
// resourceGroupName - name of the resource group to which the resource belongs.
// clusterName - the name of the cluster within the specified resource group. Cluster names can only contain a
// combination of alphanumeric characters along with dash (-) and underscore (_). The name must be from 1
// through 64 characters long.
// parameters - additional parameters for cluster update.
func (client ClustersClient) Update(ctx context.Context, resourceGroupName string, clusterName string, parameters ClusterUpdateParameters) (result Cluster, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ClustersClient.Update")
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
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: clusterName,
			Constraints: []validation.Constraint{{Target: "clusterName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "clusterName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "clusterName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("batchai.ClustersClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, resourceGroupName, clusterName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.ClustersClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "batchai.ClustersClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.ClustersClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client ClustersClient) UpdatePreparer(ctx context.Context, resourceGroupName string, clusterName string, parameters ClusterUpdateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/clusters/{clusterName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client ClustersClient) UpdateSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client ClustersClient) UpdateResponder(resp *http.Response) (result Cluster, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
