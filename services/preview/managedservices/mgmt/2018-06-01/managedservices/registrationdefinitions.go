package managedservices

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

// RegistrationDefinitionsClient is the specification for ManagedServices.
type RegistrationDefinitionsClient struct {
	BaseClient
}

// NewRegistrationDefinitionsClient creates an instance of the RegistrationDefinitionsClient client.
func NewRegistrationDefinitionsClient() RegistrationDefinitionsClient {
	return NewRegistrationDefinitionsClientWithBaseURI(DefaultBaseURI)
}

// NewRegistrationDefinitionsClientWithBaseURI creates an instance of the RegistrationDefinitionsClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewRegistrationDefinitionsClientWithBaseURI(baseURI string) RegistrationDefinitionsClient {
	return RegistrationDefinitionsClient{NewWithBaseURI(baseURI)}
}

// CreateOrUpdate creates or updates a registration definition.
// Parameters:
// registrationDefinitionID - guid of the registration definition.
// APIVersion - the API version to use for this operation.
// scope - scope of the resource.
// requestBody - the parameters required to create new registration definition.
func (client RegistrationDefinitionsClient) CreateOrUpdate(ctx context.Context, registrationDefinitionID string, APIVersion string, scope string, requestBody RegistrationDefinition) (result RegistrationDefinition, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RegistrationDefinitionsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: requestBody,
			Constraints: []validation.Constraint{{Target: "requestBody.Properties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "requestBody.Properties.Authorizations", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "requestBody.Properties.ManagedByTenantID", Name: validation.Null, Rule: true, Chain: nil},
				}},
				{Target: "requestBody.Plan", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "requestBody.Plan.Name", Name: validation.Null, Rule: true, Chain: nil},
						{Target: "requestBody.Plan.Publisher", Name: validation.Null, Rule: true, Chain: nil},
						{Target: "requestBody.Plan.Product", Name: validation.Null, Rule: true, Chain: nil},
						{Target: "requestBody.Plan.Version", Name: validation.Null, Rule: true, Chain: nil},
					}}}}}); err != nil {
		return result, validation.NewError("managedservices.RegistrationDefinitionsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, registrationDefinitionID, APIVersion, scope, requestBody)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedservices.RegistrationDefinitionsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "managedservices.RegistrationDefinitionsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedservices.RegistrationDefinitionsClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client RegistrationDefinitionsClient) CreateOrUpdatePreparer(ctx context.Context, registrationDefinitionID string, APIVersion string, scope string, requestBody RegistrationDefinition) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"registrationDefinitionId": autorest.Encode("path", registrationDefinitionID),
		"scope":                    scope,
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	requestBody.ID = nil
	requestBody.Type = nil
	requestBody.Name = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.ManagedServices/registrationDefinitions/{registrationDefinitionId}", pathParameters),
		autorest.WithJSON(requestBody),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client RegistrationDefinitionsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client RegistrationDefinitionsClient) CreateOrUpdateResponder(resp *http.Response) (result RegistrationDefinition, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the registration definition.
// Parameters:
// registrationDefinitionID - guid of the registration definition.
// APIVersion - the API version to use for this operation.
// scope - scope of the resource.
func (client RegistrationDefinitionsClient) Delete(ctx context.Context, registrationDefinitionID string, APIVersion string, scope string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RegistrationDefinitionsClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, registrationDefinitionID, APIVersion, scope)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedservices.RegistrationDefinitionsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "managedservices.RegistrationDefinitionsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedservices.RegistrationDefinitionsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client RegistrationDefinitionsClient) DeletePreparer(ctx context.Context, registrationDefinitionID string, APIVersion string, scope string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"registrationDefinitionId": autorest.Encode("path", registrationDefinitionID),
		"scope":                    scope,
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.ManagedServices/registrationDefinitions/{registrationDefinitionId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client RegistrationDefinitionsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client RegistrationDefinitionsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the registration definition details.
// Parameters:
// scope - scope of the resource.
// registrationDefinitionID - guid of the registration definition.
// APIVersion - the API version to use for this operation.
func (client RegistrationDefinitionsClient) Get(ctx context.Context, scope string, registrationDefinitionID string, APIVersion string) (result RegistrationDefinition, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RegistrationDefinitionsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, scope, registrationDefinitionID, APIVersion)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedservices.RegistrationDefinitionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "managedservices.RegistrationDefinitionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedservices.RegistrationDefinitionsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client RegistrationDefinitionsClient) GetPreparer(ctx context.Context, scope string, registrationDefinitionID string, APIVersion string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"registrationDefinitionId": autorest.Encode("path", registrationDefinitionID),
		"scope":                    scope,
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.ManagedServices/registrationDefinitions/{registrationDefinitionId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client RegistrationDefinitionsClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client RegistrationDefinitionsClient) GetResponder(resp *http.Response) (result RegistrationDefinition, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets a list of the registration definitions.
// Parameters:
// scope - scope of the resource.
// APIVersion - the API version to use for this operation.
func (client RegistrationDefinitionsClient) List(ctx context.Context, scope string, APIVersion string) (result RegistrationDefinitionListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RegistrationDefinitionsClient.List")
		defer func() {
			sc := -1
			if result.rdl.Response.Response != nil {
				sc = result.rdl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, scope, APIVersion)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedservices.RegistrationDefinitionsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.rdl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "managedservices.RegistrationDefinitionsClient", "List", resp, "Failure sending request")
		return
	}

	result.rdl, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedservices.RegistrationDefinitionsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client RegistrationDefinitionsClient) ListPreparer(ctx context.Context, scope string, APIVersion string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"scope": scope,
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.ManagedServices/registrationDefinitions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client RegistrationDefinitionsClient) ListSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client RegistrationDefinitionsClient) ListResponder(resp *http.Response) (result RegistrationDefinitionList, err error) {
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
func (client RegistrationDefinitionsClient) listNextResults(ctx context.Context, lastResults RegistrationDefinitionList) (result RegistrationDefinitionList, err error) {
	req, err := lastResults.registrationDefinitionListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "managedservices.RegistrationDefinitionsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "managedservices.RegistrationDefinitionsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedservices.RegistrationDefinitionsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client RegistrationDefinitionsClient) ListComplete(ctx context.Context, scope string, APIVersion string) (result RegistrationDefinitionListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RegistrationDefinitionsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, scope, APIVersion)
	return
}
