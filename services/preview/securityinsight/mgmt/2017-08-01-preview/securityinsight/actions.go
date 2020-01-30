package securityinsight

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

// ActionsClient is the API spec for Microsoft.SecurityInsights (Azure Security Insights) resource provider
type ActionsClient struct {
	BaseClient
}

// NewActionsClient creates an instance of the ActionsClient client.
func NewActionsClient(subscriptionID string) ActionsClient {
	return NewActionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewActionsClientWithBaseURI creates an instance of the ActionsClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewActionsClientWithBaseURI(baseURI string, subscriptionID string) ActionsClient {
	return ActionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// ListByAlertRule gets all actions of alert rule.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// operationalInsightsResourceProvider - the namespace of workspaces resource provider-
// Microsoft.OperationalInsights.
// workspaceName - the name of the workspace.
// ruleID - alert rule ID
func (client ActionsClient) ListByAlertRule(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, ruleID string) (result ActionsListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ActionsClient.ListByAlertRule")
		defer func() {
			sc := -1
			if result.al.Response.Response != nil {
				sc = result.al.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("securityinsight.ActionsClient", "ListByAlertRule", err.Error())
	}

	result.fn = client.listByAlertRuleNextResults
	req, err := client.ListByAlertRulePreparer(ctx, resourceGroupName, operationalInsightsResourceProvider, workspaceName, ruleID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.ActionsClient", "ListByAlertRule", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByAlertRuleSender(req)
	if err != nil {
		result.al.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "securityinsight.ActionsClient", "ListByAlertRule", resp, "Failure sending request")
		return
	}

	result.al, err = client.ListByAlertRuleResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.ActionsClient", "ListByAlertRule", resp, "Failure responding to request")
	}

	return
}

// ListByAlertRulePreparer prepares the ListByAlertRule request.
func (client ActionsClient) ListByAlertRulePreparer(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, ruleID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"operationalInsightsResourceProvider": autorest.Encode("path", operationalInsightsResourceProvider),
		"resourceGroupName":                   autorest.Encode("path", resourceGroupName),
		"ruleId":                              autorest.Encode("path", ruleID),
		"subscriptionId":                      autorest.Encode("path", client.SubscriptionID),
		"workspaceName":                       autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2019-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{operationalInsightsResourceProvider}/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/alertRules/{ruleId}/actions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByAlertRuleSender sends the ListByAlertRule request. The method will close the
// http.Response Body if it receives an error.
func (client ActionsClient) ListByAlertRuleSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListByAlertRuleResponder handles the response to the ListByAlertRule request. The method always
// closes the http.Response Body.
func (client ActionsClient) ListByAlertRuleResponder(resp *http.Response) (result ActionsList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByAlertRuleNextResults retrieves the next set of results, if any.
func (client ActionsClient) listByAlertRuleNextResults(ctx context.Context, lastResults ActionsList) (result ActionsList, err error) {
	req, err := lastResults.actionsListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "securityinsight.ActionsClient", "listByAlertRuleNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByAlertRuleSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "securityinsight.ActionsClient", "listByAlertRuleNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByAlertRuleResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.ActionsClient", "listByAlertRuleNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByAlertRuleComplete enumerates all values, automatically crossing page boundaries as required.
func (client ActionsClient) ListByAlertRuleComplete(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, ruleID string) (result ActionsListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ActionsClient.ListByAlertRule")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByAlertRule(ctx, resourceGroupName, operationalInsightsResourceProvider, workspaceName, ruleID)
	return
}
