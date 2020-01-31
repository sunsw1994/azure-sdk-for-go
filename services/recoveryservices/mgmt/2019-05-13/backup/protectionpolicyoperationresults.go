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

// ProtectionPolicyOperationResultsClient is the open API 2.0 Specs for Azure RecoveryServices Backup service
type ProtectionPolicyOperationResultsClient struct {
	BaseClient
}

// NewProtectionPolicyOperationResultsClient creates an instance of the ProtectionPolicyOperationResultsClient client.
func NewProtectionPolicyOperationResultsClient(subscriptionID string) ProtectionPolicyOperationResultsClient {
	return NewProtectionPolicyOperationResultsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewProtectionPolicyOperationResultsClientWithBaseURI creates an instance of the
// ProtectionPolicyOperationResultsClient client using a custom endpoint.  Use this when interacting with an Azure
// cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewProtectionPolicyOperationResultsClientWithBaseURI(baseURI string, subscriptionID string) ProtectionPolicyOperationResultsClient {
	return ProtectionPolicyOperationResultsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get provides the result of an operation.
// Parameters:
// vaultName - the name of the recovery services vault.
// resourceGroupName - the name of the resource group where the recovery services vault is present.
// policyName - backup policy name whose operation's result needs to be fetched.
// operationID - operation ID which represents the operation whose result needs to be fetched.
func (client ProtectionPolicyOperationResultsClient) Get(ctx context.Context, vaultName string, resourceGroupName string, policyName string, operationID string) (result ProtectionPolicyResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProtectionPolicyOperationResultsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, vaultName, resourceGroupName, policyName, operationID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backup.ProtectionPolicyOperationResultsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "backup.ProtectionPolicyOperationResultsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backup.ProtectionPolicyOperationResultsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ProtectionPolicyOperationResultsClient) GetPreparer(ctx context.Context, vaultName string, resourceGroupName string, policyName string, operationID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"operationId":       autorest.Encode("path", operationID),
		"policyName":        autorest.Encode("path", policyName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vaultName":         autorest.Encode("path", vaultName),
	}

	const APIVersion = "2019-05-13"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupPolicies/{policyName}/operationResults/{operationId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ProtectionPolicyOperationResultsClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ProtectionPolicyOperationResultsClient) GetResponder(resp *http.Response) (result ProtectionPolicyResource, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
