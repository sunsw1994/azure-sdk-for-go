package siterecovery

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

// ReplicationStorageClassificationsClient is the client for the ReplicationStorageClassifications methods of the
// Siterecovery service.
type ReplicationStorageClassificationsClient struct {
	BaseClient
}

// NewReplicationStorageClassificationsClient creates an instance of the ReplicationStorageClassificationsClient
// client.
func NewReplicationStorageClassificationsClient(subscriptionID string, resourceGroupName string, resourceName string) ReplicationStorageClassificationsClient {
	return NewReplicationStorageClassificationsClientWithBaseURI(DefaultBaseURI, subscriptionID, resourceGroupName, resourceName)
}

// NewReplicationStorageClassificationsClientWithBaseURI creates an instance of the
// ReplicationStorageClassificationsClient client using a custom endpoint.  Use this when interacting with an Azure
// cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewReplicationStorageClassificationsClientWithBaseURI(baseURI string, subscriptionID string, resourceGroupName string, resourceName string) ReplicationStorageClassificationsClient {
	return ReplicationStorageClassificationsClient{NewWithBaseURI(baseURI, subscriptionID, resourceGroupName, resourceName)}
}

// Get gets the details of the specified storage classification.
// Parameters:
// fabricName - fabric name.
// storageClassificationName - storage classification name.
func (client ReplicationStorageClassificationsClient) Get(ctx context.Context, fabricName string, storageClassificationName string) (result StorageClassification, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReplicationStorageClassificationsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, fabricName, storageClassificationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationStorageClassificationsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationStorageClassificationsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationStorageClassificationsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ReplicationStorageClassificationsClient) GetPreparer(ctx context.Context, fabricName string, storageClassificationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fabricName":                autorest.Encode("path", fabricName),
		"resourceGroupName":         autorest.Encode("path", client.ResourceGroupName),
		"resourceName":              autorest.Encode("path", client.ResourceName),
		"storageClassificationName": autorest.Encode("path", storageClassificationName),
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-08-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationStorageClassifications/{storageClassificationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationStorageClassificationsClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ReplicationStorageClassificationsClient) GetResponder(resp *http.Response) (result StorageClassification, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists the storage classifications in the vault.
func (client ReplicationStorageClassificationsClient) List(ctx context.Context) (result StorageClassificationCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReplicationStorageClassificationsClient.List")
		defer func() {
			sc := -1
			if result.scc.Response.Response != nil {
				sc = result.scc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationStorageClassificationsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.scc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationStorageClassificationsClient", "List", resp, "Failure sending request")
		return
	}

	result.scc, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationStorageClassificationsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ReplicationStorageClassificationsClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", client.ResourceGroupName),
		"resourceName":      autorest.Encode("path", client.ResourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-08-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationStorageClassifications", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationStorageClassificationsClient) ListSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ReplicationStorageClassificationsClient) ListResponder(resp *http.Response) (result StorageClassificationCollection, err error) {
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
func (client ReplicationStorageClassificationsClient) listNextResults(ctx context.Context, lastResults StorageClassificationCollection) (result StorageClassificationCollection, err error) {
	req, err := lastResults.storageClassificationCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "siterecovery.ReplicationStorageClassificationsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "siterecovery.ReplicationStorageClassificationsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationStorageClassificationsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client ReplicationStorageClassificationsClient) ListComplete(ctx context.Context) (result StorageClassificationCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReplicationStorageClassificationsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx)
	return
}

// ListByReplicationFabrics lists the storage classifications available in the specified fabric.
// Parameters:
// fabricName - site name of interest.
func (client ReplicationStorageClassificationsClient) ListByReplicationFabrics(ctx context.Context, fabricName string) (result StorageClassificationCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReplicationStorageClassificationsClient.ListByReplicationFabrics")
		defer func() {
			sc := -1
			if result.scc.Response.Response != nil {
				sc = result.scc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByReplicationFabricsNextResults
	req, err := client.ListByReplicationFabricsPreparer(ctx, fabricName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationStorageClassificationsClient", "ListByReplicationFabrics", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByReplicationFabricsSender(req)
	if err != nil {
		result.scc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationStorageClassificationsClient", "ListByReplicationFabrics", resp, "Failure sending request")
		return
	}

	result.scc, err = client.ListByReplicationFabricsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationStorageClassificationsClient", "ListByReplicationFabrics", resp, "Failure responding to request")
	}

	return
}

// ListByReplicationFabricsPreparer prepares the ListByReplicationFabrics request.
func (client ReplicationStorageClassificationsClient) ListByReplicationFabricsPreparer(ctx context.Context, fabricName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fabricName":        autorest.Encode("path", fabricName),
		"resourceGroupName": autorest.Encode("path", client.ResourceGroupName),
		"resourceName":      autorest.Encode("path", client.ResourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-08-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationStorageClassifications", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByReplicationFabricsSender sends the ListByReplicationFabrics request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationStorageClassificationsClient) ListByReplicationFabricsSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListByReplicationFabricsResponder handles the response to the ListByReplicationFabrics request. The method always
// closes the http.Response Body.
func (client ReplicationStorageClassificationsClient) ListByReplicationFabricsResponder(resp *http.Response) (result StorageClassificationCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByReplicationFabricsNextResults retrieves the next set of results, if any.
func (client ReplicationStorageClassificationsClient) listByReplicationFabricsNextResults(ctx context.Context, lastResults StorageClassificationCollection) (result StorageClassificationCollection, err error) {
	req, err := lastResults.storageClassificationCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "siterecovery.ReplicationStorageClassificationsClient", "listByReplicationFabricsNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByReplicationFabricsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "siterecovery.ReplicationStorageClassificationsClient", "listByReplicationFabricsNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByReplicationFabricsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationStorageClassificationsClient", "listByReplicationFabricsNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByReplicationFabricsComplete enumerates all values, automatically crossing page boundaries as required.
func (client ReplicationStorageClassificationsClient) ListByReplicationFabricsComplete(ctx context.Context, fabricName string) (result StorageClassificationCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReplicationStorageClassificationsClient.ListByReplicationFabrics")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByReplicationFabrics(ctx, fabricName)
	return
}
