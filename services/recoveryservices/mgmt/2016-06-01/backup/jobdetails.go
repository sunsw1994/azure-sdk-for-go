package backup

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

// JobDetailsClient is the open API 2.0 Specs for Azure RecoveryServices Backup service
type JobDetailsClient struct {
	BaseClient
}

// NewJobDetailsClient creates an instance of the JobDetailsClient client.
func NewJobDetailsClient(subscriptionID string) JobDetailsClient {
	return NewJobDetailsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewJobDetailsClientWithBaseURI creates an instance of the JobDetailsClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewJobDetailsClientWithBaseURI(baseURI string, subscriptionID string) JobDetailsClient {
	return JobDetailsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get gets extended information associated with the job.
// Parameters:
// vaultName - the name of the Recovery Services vault.
// resourceGroupName - the name of the resource group associated with the Recovery Services vault.
// jobName - name of the job associated with this GET operation.
func (client JobDetailsClient) Get(ctx context.Context, vaultName string, resourceGroupName string, jobName string) (result JobResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobDetailsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, vaultName, resourceGroupName, jobName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backup.JobDetailsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "backup.JobDetailsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backup.JobDetailsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client JobDetailsClient) GetPreparer(ctx context.Context, vaultName string, resourceGroupName string, jobName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobName":           autorest.Encode("path", jobName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vaultName":         autorest.Encode("path", vaultName),
	}

	const APIVersion = "2016-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupJobs/{jobName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client JobDetailsClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client JobDetailsClient) GetResponder(resp *http.Response) (result JobResource, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
