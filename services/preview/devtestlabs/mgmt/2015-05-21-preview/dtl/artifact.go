package dtl

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

// ArtifactClient is the azure DevTest Labs REST API version 2015-05-21-preview.
type ArtifactClient struct {
	BaseClient
}

// NewArtifactClient creates an instance of the ArtifactClient client.
func NewArtifactClient(subscriptionID string) ArtifactClient {
	return NewArtifactClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewArtifactClientWithBaseURI creates an instance of the ArtifactClient client.
func NewArtifactClientWithBaseURI(baseURI string, subscriptionID string) ArtifactClient {
	return ArtifactClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// GenerateArmTemplate generates an ARM template for the given artifact, uploads the required files to a storage
// account, and validates the generated artifact.
// Parameters:
// resourceGroupName - the name of the resource group.
// labName - the name of the lab.
// artifactSourceName - the name of the artifact source.
// name - the name of the artifact.
func (client ArtifactClient) GenerateArmTemplate(ctx context.Context, resourceGroupName string, labName string, artifactSourceName string, name string, generateArmTemplateRequest GenerateArmTemplateRequest) (result ArmTemplateInfo, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ArtifactClient.GenerateArmTemplate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GenerateArmTemplatePreparer(ctx, resourceGroupName, labName, artifactSourceName, name, generateArmTemplateRequest)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.ArtifactClient", "GenerateArmTemplate", nil, "Failure preparing request")
		return
	}

	resp, err := client.GenerateArmTemplateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "dtl.ArtifactClient", "GenerateArmTemplate", resp, "Failure sending request")
		return
	}

	result, err = client.GenerateArmTemplateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.ArtifactClient", "GenerateArmTemplate", resp, "Failure responding to request")
	}

	return
}

// GenerateArmTemplatePreparer prepares the GenerateArmTemplate request.
func (client ArtifactClient) GenerateArmTemplatePreparer(ctx context.Context, resourceGroupName string, labName string, artifactSourceName string, name string, generateArmTemplateRequest GenerateArmTemplateRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"artifactSourceName": autorest.Encode("path", artifactSourceName),
		"labName":            autorest.Encode("path", labName),
		"name":               autorest.Encode("path", name),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-05-21-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/artifactsources/{artifactSourceName}/artifacts/{name}/generateArmTemplate", pathParameters),
		autorest.WithJSON(generateArmTemplateRequest),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GenerateArmTemplateSender sends the GenerateArmTemplate request. The method will close the
// http.Response Body if it receives an error.
func (client ArtifactClient) GenerateArmTemplateSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GenerateArmTemplateResponder handles the response to the GenerateArmTemplate request. The method always
// closes the http.Response Body.
func (client ArtifactClient) GenerateArmTemplateResponder(resp *http.Response) (result ArmTemplateInfo, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetResource get artifact.
// Parameters:
// resourceGroupName - the name of the resource group.
// labName - the name of the lab.
// artifactSourceName - the name of the artifact source.
// name - the name of the artifact.
func (client ArtifactClient) GetResource(ctx context.Context, resourceGroupName string, labName string, artifactSourceName string, name string) (result Artifact, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ArtifactClient.GetResource")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetResourcePreparer(ctx, resourceGroupName, labName, artifactSourceName, name)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.ArtifactClient", "GetResource", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetResourceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "dtl.ArtifactClient", "GetResource", resp, "Failure sending request")
		return
	}

	result, err = client.GetResourceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.ArtifactClient", "GetResource", resp, "Failure responding to request")
	}

	return
}

// GetResourcePreparer prepares the GetResource request.
func (client ArtifactClient) GetResourcePreparer(ctx context.Context, resourceGroupName string, labName string, artifactSourceName string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"artifactSourceName": autorest.Encode("path", artifactSourceName),
		"labName":            autorest.Encode("path", labName),
		"name":               autorest.Encode("path", name),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-05-21-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/artifactsources/{artifactSourceName}/artifacts/{name}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetResourceSender sends the GetResource request. The method will close the
// http.Response Body if it receives an error.
func (client ArtifactClient) GetResourceSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResourceResponder handles the response to the GetResource request. The method always
// closes the http.Response Body.
func (client ArtifactClient) GetResourceResponder(resp *http.Response) (result Artifact, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list artifacts.
// Parameters:
// resourceGroupName - the name of the resource group.
// labName - the name of the lab.
// artifactSourceName - the name of the artifact source.
// filter - the filter to apply on the operation.
func (client ArtifactClient) List(ctx context.Context, resourceGroupName string, labName string, artifactSourceName string, filter string, top *int32, orderBy string) (result ResponseWithContinuationArtifactPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ArtifactClient.List")
		defer func() {
			sc := -1
			if result.rwca.Response.Response != nil {
				sc = result.rwca.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, labName, artifactSourceName, filter, top, orderBy)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.ArtifactClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.rwca.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "dtl.ArtifactClient", "List", resp, "Failure sending request")
		return
	}

	result.rwca, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.ArtifactClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ArtifactClient) ListPreparer(ctx context.Context, resourceGroupName string, labName string, artifactSourceName string, filter string, top *int32, orderBy string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"artifactSourceName": autorest.Encode("path", artifactSourceName),
		"labName":            autorest.Encode("path", labName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-05-21-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(orderBy) > 0 {
		queryParameters["$orderBy"] = autorest.Encode("query", orderBy)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/artifactsources/{artifactSourceName}/artifacts", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ArtifactClient) ListSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ArtifactClient) ListResponder(resp *http.Response) (result ResponseWithContinuationArtifact, err error) {
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
func (client ArtifactClient) listNextResults(ctx context.Context, lastResults ResponseWithContinuationArtifact) (result ResponseWithContinuationArtifact, err error) {
	req, err := lastResults.responseWithContinuationArtifactPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "dtl.ArtifactClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "dtl.ArtifactClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.ArtifactClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client ArtifactClient) ListComplete(ctx context.Context, resourceGroupName string, labName string, artifactSourceName string, filter string, top *int32, orderBy string) (result ResponseWithContinuationArtifactIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ArtifactClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, labName, artifactSourceName, filter, top, orderBy)
	return
}