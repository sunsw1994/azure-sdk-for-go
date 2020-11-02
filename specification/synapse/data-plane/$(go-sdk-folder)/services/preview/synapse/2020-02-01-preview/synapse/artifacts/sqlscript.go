package synapse

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
    "github.com/Azure/go-autorest/autorest"
    "github.com/Azure/go-autorest/autorest/azure"
    "net/http"
    "context"
    "github.com/Azure/go-autorest/tracing"
    "github.com/Azure/go-autorest/autorest/validation"
)

// SQLScriptClient is the client for the SQLScript methods of the Synapse service.
type SQLScriptClient struct {
    BaseClient
}
// NewSQLScriptClient creates an instance of the SQLScriptClient client.
func NewSQLScriptClient(endpoint string) SQLScriptClient {
    return SQLScriptClient{ New(endpoint)}
}

// CreateOrUpdateSQLScript creates or updates a Sql Script.
    // Parameters:
        // SQLScriptName - the sql script name.
        // SQLScript - sql Script resource definition.
        // ifMatch - eTag of the SQL script entity.  Should only be specified for update, for which it should match
        // existing entity or can be * for unconditional update.
func (client SQLScriptClient) CreateOrUpdateSQLScript(ctx context.Context, SQLScriptName string, SQLScript SQLScriptResource, ifMatch string) (result SQLScriptResource, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/SQLScriptClient.CreateOrUpdateSQLScript")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        if err := validation.Validate([]validation.Validation{
        { TargetValue: SQLScript,
         Constraints: []validation.Constraint{	{Target: "SQLScript.Name", Name: validation.Null, Rule: true, Chain: nil },
        	{Target: "SQLScript.Properties", Name: validation.Null, Rule: true ,
        Chain: []validation.Constraint{	{Target: "SQLScript.Properties.Content", Name: validation.Null, Rule: true ,
        Chain: []validation.Constraint{	{Target: "SQLScript.Properties.Content.Query", Name: validation.Null, Rule: true, Chain: nil },
        	{Target: "SQLScript.Properties.Content.CurrentConnection", Name: validation.Null, Rule: true ,
        Chain: []validation.Constraint{	{Target: "SQLScript.Properties.Content.CurrentConnection.Name", Name: validation.Null, Rule: true, Chain: nil },
        }},
        }},
        }}}}}); err != nil {
        return result, validation.NewError("synapse.SQLScriptClient", "CreateOrUpdateSQLScript", err.Error())
        }

        req, err := client.CreateOrUpdateSQLScriptPreparer(ctx, SQLScriptName, SQLScript, ifMatch)
    if err != nil {
    err = autorest.NewErrorWithError(err, "synapse.SQLScriptClient", "CreateOrUpdateSQLScript", nil , "Failure preparing request")
    return
    }

        resp, err := client.CreateOrUpdateSQLScriptSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "synapse.SQLScriptClient", "CreateOrUpdateSQLScript", resp, "Failure sending request")
        return
        }

        result, err = client.CreateOrUpdateSQLScriptResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "synapse.SQLScriptClient", "CreateOrUpdateSQLScript", resp, "Failure responding to request")
        }

    return
}

    // CreateOrUpdateSQLScriptPreparer prepares the CreateOrUpdateSQLScript request.
    func (client SQLScriptClient) CreateOrUpdateSQLScriptPreparer(ctx context.Context, SQLScriptName string, SQLScript SQLScriptResource, ifMatch string) (*http.Request, error) {
        urlParameters := map[string]interface{} {
        "endpoint": client.Endpoint,
        }

        pathParameters := map[string]interface{} {
        "sqlScriptName": autorest.Encode("path",SQLScriptName),
        }

            const APIVersion = "2019-06-01-preview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

            SQLScript.ID = nil
            SQLScript.Type = nil
            SQLScript.Etag = nil
    preparer := autorest.CreatePreparer(
autorest.AsContentType("application/json; charset=utf-8"),
autorest.AsPut(),
autorest.WithCustomBaseURL("{endpoint}", urlParameters),
autorest.WithPathParameters("/sqlScripts/{sqlScriptName}",pathParameters),
autorest.WithJSON(SQLScript),
autorest.WithQueryParameters(queryParameters))
        if len(ifMatch) > 0 {
        preparer = autorest.DecoratePreparer(preparer,
        autorest.WithHeader("If-Match",autorest.String(ifMatch)))
        }
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CreateOrUpdateSQLScriptSender sends the CreateOrUpdateSQLScript request. The method will close the
    // http.Response Body if it receives an error.
    func (client SQLScriptClient) CreateOrUpdateSQLScriptSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            }

    // CreateOrUpdateSQLScriptResponder handles the response to the CreateOrUpdateSQLScript request. The method always
    // closes the http.Response Body.
    func (client SQLScriptClient) CreateOrUpdateSQLScriptResponder(resp *http.Response) (result SQLScriptResource, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// DeleteSQLScript deletes a Sql Script.
    // Parameters:
        // SQLScriptName - the sql script name.
func (client SQLScriptClient) DeleteSQLScript(ctx context.Context, SQLScriptName string) (result autorest.Response, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/SQLScriptClient.DeleteSQLScript")
        defer func() {
            sc := -1
        if result.Response != nil {
        sc = result.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.DeleteSQLScriptPreparer(ctx, SQLScriptName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "synapse.SQLScriptClient", "DeleteSQLScript", nil , "Failure preparing request")
    return
    }

        resp, err := client.DeleteSQLScriptSender(req)
        if err != nil {
        result.Response = resp
        err = autorest.NewErrorWithError(err, "synapse.SQLScriptClient", "DeleteSQLScript", resp, "Failure sending request")
        return
        }

        result, err = client.DeleteSQLScriptResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "synapse.SQLScriptClient", "DeleteSQLScript", resp, "Failure responding to request")
        }

    return
}

    // DeleteSQLScriptPreparer prepares the DeleteSQLScript request.
    func (client SQLScriptClient) DeleteSQLScriptPreparer(ctx context.Context, SQLScriptName string) (*http.Request, error) {
        urlParameters := map[string]interface{} {
        "endpoint": client.Endpoint,
        }

        pathParameters := map[string]interface{} {
        "sqlScriptName": autorest.Encode("path",SQLScriptName),
        }

            const APIVersion = "2019-06-01-preview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsDelete(),
autorest.WithCustomBaseURL("{endpoint}", urlParameters),
autorest.WithPathParameters("/sqlScripts/{sqlScriptName}",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // DeleteSQLScriptSender sends the DeleteSQLScript request. The method will close the
    // http.Response Body if it receives an error.
    func (client SQLScriptClient) DeleteSQLScriptSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            }

    // DeleteSQLScriptResponder handles the response to the DeleteSQLScript request. The method always
    // closes the http.Response Body.
    func (client SQLScriptClient) DeleteSQLScriptResponder(resp *http.Response) (result autorest.Response, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusNoContent),
            autorest.ByClosing())
            result.Response = resp
            return
    }

// GetSQLScript gets a sql script.
    // Parameters:
        // SQLScriptName - the sql script name.
        // ifNoneMatch - eTag of the sql compute entity. Should only be specified for get. If the ETag matches the
        // existing entity tag, or if * was provided, then no content will be returned.
func (client SQLScriptClient) GetSQLScript(ctx context.Context, SQLScriptName string, ifNoneMatch string) (result SQLScriptResource, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/SQLScriptClient.GetSQLScript")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.GetSQLScriptPreparer(ctx, SQLScriptName, ifNoneMatch)
    if err != nil {
    err = autorest.NewErrorWithError(err, "synapse.SQLScriptClient", "GetSQLScript", nil , "Failure preparing request")
    return
    }

        resp, err := client.GetSQLScriptSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "synapse.SQLScriptClient", "GetSQLScript", resp, "Failure sending request")
        return
        }

        result, err = client.GetSQLScriptResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "synapse.SQLScriptClient", "GetSQLScript", resp, "Failure responding to request")
        }

    return
}

    // GetSQLScriptPreparer prepares the GetSQLScript request.
    func (client SQLScriptClient) GetSQLScriptPreparer(ctx context.Context, SQLScriptName string, ifNoneMatch string) (*http.Request, error) {
        urlParameters := map[string]interface{} {
        "endpoint": client.Endpoint,
        }

        pathParameters := map[string]interface{} {
        "sqlScriptName": autorest.Encode("path",SQLScriptName),
        }

            const APIVersion = "2019-06-01-preview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithCustomBaseURL("{endpoint}", urlParameters),
autorest.WithPathParameters("/sqlScripts/{sqlScriptName}",pathParameters),
autorest.WithQueryParameters(queryParameters))
        if len(ifNoneMatch) > 0 {
        preparer = autorest.DecoratePreparer(preparer,
        autorest.WithHeader("If-None-Match",autorest.String(ifNoneMatch)))
        }
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSQLScriptSender sends the GetSQLScript request. The method will close the
    // http.Response Body if it receives an error.
    func (client SQLScriptClient) GetSQLScriptSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            }

    // GetSQLScriptResponder handles the response to the GetSQLScript request. The method always
    // closes the http.Response Body.
    func (client SQLScriptClient) GetSQLScriptResponder(resp *http.Response) (result SQLScriptResource, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusNotModified),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// GetSQLScriptsByWorkspace lists sql scripts.
func (client SQLScriptClient) GetSQLScriptsByWorkspace(ctx context.Context) (result SQLScriptsListResponsePage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/SQLScriptClient.GetSQLScriptsByWorkspace")
        defer func() {
            sc := -1
        if result.sslr.Response.Response != nil {
        sc = result.sslr.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        result.fn = client.getSQLScriptsByWorkspaceNextResults
    req, err := client.GetSQLScriptsByWorkspacePreparer(ctx)
    if err != nil {
    err = autorest.NewErrorWithError(err, "synapse.SQLScriptClient", "GetSQLScriptsByWorkspace", nil , "Failure preparing request")
    return
    }

        resp, err := client.GetSQLScriptsByWorkspaceSender(req)
        if err != nil {
        result.sslr.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "synapse.SQLScriptClient", "GetSQLScriptsByWorkspace", resp, "Failure sending request")
        return
        }

        result.sslr, err = client.GetSQLScriptsByWorkspaceResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "synapse.SQLScriptClient", "GetSQLScriptsByWorkspace", resp, "Failure responding to request")
        }
            if result.sslr.hasNextLink() && result.sslr.IsEmpty() {
            err = result.NextWithContext(ctx)
            }

    return
}

    // GetSQLScriptsByWorkspacePreparer prepares the GetSQLScriptsByWorkspace request.
    func (client SQLScriptClient) GetSQLScriptsByWorkspacePreparer(ctx context.Context) (*http.Request, error) {
        urlParameters := map[string]interface{} {
        "endpoint": client.Endpoint,
        }

            const APIVersion = "2019-06-01-preview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithCustomBaseURL("{endpoint}", urlParameters),
autorest.WithPath("/sqlScripts"),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSQLScriptsByWorkspaceSender sends the GetSQLScriptsByWorkspace request. The method will close the
    // http.Response Body if it receives an error.
    func (client SQLScriptClient) GetSQLScriptsByWorkspaceSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            }

    // GetSQLScriptsByWorkspaceResponder handles the response to the GetSQLScriptsByWorkspace request. The method always
    // closes the http.Response Body.
    func (client SQLScriptClient) GetSQLScriptsByWorkspaceResponder(resp *http.Response) (result SQLScriptsListResponse, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

            // getSQLScriptsByWorkspaceNextResults retrieves the next set of results, if any.
            func (client SQLScriptClient) getSQLScriptsByWorkspaceNextResults(ctx context.Context, lastResults SQLScriptsListResponse) (result SQLScriptsListResponse, err error) {
            req, err := lastResults.sQLScriptsListResponsePreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "synapse.SQLScriptClient", "getSQLScriptsByWorkspaceNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.GetSQLScriptsByWorkspaceSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "synapse.SQLScriptClient", "getSQLScriptsByWorkspaceNextResults", resp, "Failure sending next results request")
            }
            result, err = client.GetSQLScriptsByWorkspaceResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "synapse.SQLScriptClient", "getSQLScriptsByWorkspaceNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

            // GetSQLScriptsByWorkspaceComplete enumerates all values, automatically crossing page boundaries as required.
            func (client SQLScriptClient) GetSQLScriptsByWorkspaceComplete(ctx context.Context) (result SQLScriptsListResponseIterator, err error) {
            if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/SQLScriptClient.GetSQLScriptsByWorkspace")
            defer func() {
            sc := -1
            if result.Response().Response.Response != nil {
            sc = result.page.Response().Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
            }()
            }
                    result.page, err = client.GetSQLScriptsByWorkspace(ctx)
                            return
            }

