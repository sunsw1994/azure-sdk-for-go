package hdinsight

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

// ExtensionsClient is the hDInsight Management Client
type ExtensionsClient struct {
	BaseClient
}

// NewExtensionsClient creates an instance of the ExtensionsClient client.
func NewExtensionsClient(subscriptionID string) ExtensionsClient {
	return NewExtensionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewExtensionsClientWithBaseURI creates an instance of the ExtensionsClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewExtensionsClientWithBaseURI(baseURI string, subscriptionID string) ExtensionsClient {
	return ExtensionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create creates an HDInsight cluster extension.
// Parameters:
// resourceGroupName - the name of the resource group.
// clusterName - the name of the cluster.
// extensionName - the name of the cluster extension.
// parameters - the cluster extensions create request.
func (client ExtensionsClient) Create(ctx context.Context, resourceGroupName string, clusterName string, extensionName string, parameters Extension) (result ExtensionsCreateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExtensionsClient.Create")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreatePreparer(ctx, resourceGroupName, clusterName, extensionName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hdinsight.ExtensionsClient", "Create", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hdinsight.ExtensionsClient", "Create", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client ExtensionsClient) CreatePreparer(ctx context.Context, resourceGroupName string, clusterName string, extensionName string, parameters Extension) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"extensionName":     autorest.Encode("path", extensionName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/extensions/{extensionName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client ExtensionsClient) CreateSender(req *http.Request) (future ExtensionsCreateFuture, err error) {
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
func (client ExtensionsClient) CreateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Delete deletes the specified extension for HDInsight cluster.
// Parameters:
// resourceGroupName - the name of the resource group.
// clusterName - the name of the cluster.
// extensionName - the name of the cluster extension.
func (client ExtensionsClient) Delete(ctx context.Context, resourceGroupName string, clusterName string, extensionName string) (result ExtensionsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExtensionsClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, clusterName, extensionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hdinsight.ExtensionsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hdinsight.ExtensionsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ExtensionsClient) DeletePreparer(ctx context.Context, resourceGroupName string, clusterName string, extensionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"extensionName":     autorest.Encode("path", extensionName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/extensions/{extensionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ExtensionsClient) DeleteSender(req *http.Request) (future ExtensionsDeleteFuture, err error) {
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
func (client ExtensionsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// DisableMonitoring disables the Operations Management Suite (OMS) on the HDInsight cluster.
// Parameters:
// resourceGroupName - the name of the resource group.
// clusterName - the name of the cluster.
func (client ExtensionsClient) DisableMonitoring(ctx context.Context, resourceGroupName string, clusterName string) (result ExtensionsDisableMonitoringFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExtensionsClient.DisableMonitoring")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DisableMonitoringPreparer(ctx, resourceGroupName, clusterName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hdinsight.ExtensionsClient", "DisableMonitoring", nil, "Failure preparing request")
		return
	}

	result, err = client.DisableMonitoringSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hdinsight.ExtensionsClient", "DisableMonitoring", result.Response(), "Failure sending request")
		return
	}

	return
}

// DisableMonitoringPreparer prepares the DisableMonitoring request.
func (client ExtensionsClient) DisableMonitoringPreparer(ctx context.Context, resourceGroupName string, clusterName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/extensions/clustermonitoring", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DisableMonitoringSender sends the DisableMonitoring request. The method will close the
// http.Response Body if it receives an error.
func (client ExtensionsClient) DisableMonitoringSender(req *http.Request) (future ExtensionsDisableMonitoringFuture, err error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req, sd...)
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// DisableMonitoringResponder handles the response to the DisableMonitoring request. The method always
// closes the http.Response Body.
func (client ExtensionsClient) DisableMonitoringResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// EnableMonitoring enables the Operations Management Suite (OMS) on the HDInsight cluster.
// Parameters:
// resourceGroupName - the name of the resource group.
// clusterName - the name of the cluster.
// parameters - the Operations Management Suite (OMS) workspace parameters.
func (client ExtensionsClient) EnableMonitoring(ctx context.Context, resourceGroupName string, clusterName string, parameters ClusterMonitoringRequest) (result ExtensionsEnableMonitoringFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExtensionsClient.EnableMonitoring")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.EnableMonitoringPreparer(ctx, resourceGroupName, clusterName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hdinsight.ExtensionsClient", "EnableMonitoring", nil, "Failure preparing request")
		return
	}

	result, err = client.EnableMonitoringSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hdinsight.ExtensionsClient", "EnableMonitoring", result.Response(), "Failure sending request")
		return
	}

	return
}

// EnableMonitoringPreparer prepares the EnableMonitoring request.
func (client ExtensionsClient) EnableMonitoringPreparer(ctx context.Context, resourceGroupName string, clusterName string, parameters ClusterMonitoringRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/extensions/clustermonitoring", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// EnableMonitoringSender sends the EnableMonitoring request. The method will close the
// http.Response Body if it receives an error.
func (client ExtensionsClient) EnableMonitoringSender(req *http.Request) (future ExtensionsEnableMonitoringFuture, err error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req, sd...)
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// EnableMonitoringResponder handles the response to the EnableMonitoring request. The method always
// closes the http.Response Body.
func (client ExtensionsClient) EnableMonitoringResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the extension properties for the specified HDInsight cluster extension.
// Parameters:
// resourceGroupName - the name of the resource group.
// clusterName - the name of the cluster.
// extensionName - the name of the cluster extension.
func (client ExtensionsClient) Get(ctx context.Context, resourceGroupName string, clusterName string, extensionName string) (result Extension, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExtensionsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, clusterName, extensionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hdinsight.ExtensionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "hdinsight.ExtensionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hdinsight.ExtensionsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ExtensionsClient) GetPreparer(ctx context.Context, resourceGroupName string, clusterName string, extensionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"extensionName":     autorest.Encode("path", extensionName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/extensions/{extensionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ExtensionsClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ExtensionsClient) GetResponder(resp *http.Response) (result Extension, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetMonitoringStatus gets the status of Operations Management Suite (OMS) on the HDInsight cluster.
// Parameters:
// resourceGroupName - the name of the resource group.
// clusterName - the name of the cluster.
func (client ExtensionsClient) GetMonitoringStatus(ctx context.Context, resourceGroupName string, clusterName string) (result ClusterMonitoringResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExtensionsClient.GetMonitoringStatus")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetMonitoringStatusPreparer(ctx, resourceGroupName, clusterName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hdinsight.ExtensionsClient", "GetMonitoringStatus", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetMonitoringStatusSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "hdinsight.ExtensionsClient", "GetMonitoringStatus", resp, "Failure sending request")
		return
	}

	result, err = client.GetMonitoringStatusResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hdinsight.ExtensionsClient", "GetMonitoringStatus", resp, "Failure responding to request")
	}

	return
}

// GetMonitoringStatusPreparer prepares the GetMonitoringStatus request.
func (client ExtensionsClient) GetMonitoringStatusPreparer(ctx context.Context, resourceGroupName string, clusterName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/extensions/clustermonitoring", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetMonitoringStatusSender sends the GetMonitoringStatus request. The method will close the
// http.Response Body if it receives an error.
func (client ExtensionsClient) GetMonitoringStatusSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetMonitoringStatusResponder handles the response to the GetMonitoringStatus request. The method always
// closes the http.Response Body.
func (client ExtensionsClient) GetMonitoringStatusResponder(resp *http.Response) (result ClusterMonitoringResponse, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
