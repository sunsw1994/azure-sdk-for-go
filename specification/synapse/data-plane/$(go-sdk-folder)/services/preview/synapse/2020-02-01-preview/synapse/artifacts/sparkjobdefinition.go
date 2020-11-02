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

// SparkJobDefinitionClient is the client for the SparkJobDefinition methods of the Synapse service.
type SparkJobDefinitionClient struct {
    BaseClient
}
// NewSparkJobDefinitionClient creates an instance of the SparkJobDefinitionClient client.
func NewSparkJobDefinitionClient(endpoint string) SparkJobDefinitionClient {
    return SparkJobDefinitionClient{ New(endpoint)}
}

// CreateOrUpdateSparkJobDefinition creates or updates a Spark Job Definition.
    // Parameters:
        // sparkJobDefinitionName - the spark job definition name.
        // sparkJobDefinition - spark Job Definition resource definition.
        // ifMatch - eTag of the Spark Job Definition entity.  Should only be specified for update, for which it should
        // match existing entity or can be * for unconditional update.
func (client SparkJobDefinitionClient) CreateOrUpdateSparkJobDefinition(ctx context.Context, sparkJobDefinitionName string, sparkJobDefinition SparkJobDefinitionResource, ifMatch string) (result SparkJobDefinitionResource, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/SparkJobDefinitionClient.CreateOrUpdateSparkJobDefinition")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        if err := validation.Validate([]validation.Validation{
        { TargetValue: sparkJobDefinition,
         Constraints: []validation.Constraint{	{Target: "sparkJobDefinition.Properties", Name: validation.Null, Rule: true ,
        Chain: []validation.Constraint{	{Target: "sparkJobDefinition.Properties.TargetBigDataPool", Name: validation.Null, Rule: true ,
        Chain: []validation.Constraint{	{Target: "sparkJobDefinition.Properties.TargetBigDataPool.Type", Name: validation.Null, Rule: true, Chain: nil },
        	{Target: "sparkJobDefinition.Properties.TargetBigDataPool.ReferenceName", Name: validation.Null, Rule: true, Chain: nil },
        }},
        	{Target: "sparkJobDefinition.Properties.JobProperties", Name: validation.Null, Rule: true ,
        Chain: []validation.Constraint{	{Target: "sparkJobDefinition.Properties.JobProperties.File", Name: validation.Null, Rule: true, Chain: nil },
        	{Target: "sparkJobDefinition.Properties.JobProperties.DriverMemory", Name: validation.Null, Rule: true, Chain: nil },
        	{Target: "sparkJobDefinition.Properties.JobProperties.DriverCores", Name: validation.Null, Rule: true, Chain: nil },
        	{Target: "sparkJobDefinition.Properties.JobProperties.ExecutorMemory", Name: validation.Null, Rule: true, Chain: nil },
        	{Target: "sparkJobDefinition.Properties.JobProperties.ExecutorCores", Name: validation.Null, Rule: true, Chain: nil },
        	{Target: "sparkJobDefinition.Properties.JobProperties.NumExecutors", Name: validation.Null, Rule: true, Chain: nil },
        }},
        }}}}}); err != nil {
        return result, validation.NewError("synapse.SparkJobDefinitionClient", "CreateOrUpdateSparkJobDefinition", err.Error())
        }

        req, err := client.CreateOrUpdateSparkJobDefinitionPreparer(ctx, sparkJobDefinitionName, sparkJobDefinition, ifMatch)
    if err != nil {
    err = autorest.NewErrorWithError(err, "synapse.SparkJobDefinitionClient", "CreateOrUpdateSparkJobDefinition", nil , "Failure preparing request")
    return
    }

        resp, err := client.CreateOrUpdateSparkJobDefinitionSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "synapse.SparkJobDefinitionClient", "CreateOrUpdateSparkJobDefinition", resp, "Failure sending request")
        return
        }

        result, err = client.CreateOrUpdateSparkJobDefinitionResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "synapse.SparkJobDefinitionClient", "CreateOrUpdateSparkJobDefinition", resp, "Failure responding to request")
        }

    return
}

    // CreateOrUpdateSparkJobDefinitionPreparer prepares the CreateOrUpdateSparkJobDefinition request.
    func (client SparkJobDefinitionClient) CreateOrUpdateSparkJobDefinitionPreparer(ctx context.Context, sparkJobDefinitionName string, sparkJobDefinition SparkJobDefinitionResource, ifMatch string) (*http.Request, error) {
        urlParameters := map[string]interface{} {
        "endpoint": client.Endpoint,
        }

        pathParameters := map[string]interface{} {
        "sparkJobDefinitionName": autorest.Encode("path",sparkJobDefinitionName),
        }

            const APIVersion = "2019-06-01-preview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsContentType("application/json; charset=utf-8"),
autorest.AsPut(),
autorest.WithCustomBaseURL("{endpoint}", urlParameters),
autorest.WithPathParameters("/sparkJobDefinitions/{sparkJobDefinitionName}",pathParameters),
autorest.WithJSON(sparkJobDefinition),
autorest.WithQueryParameters(queryParameters))
        if len(ifMatch) > 0 {
        preparer = autorest.DecoratePreparer(preparer,
        autorest.WithHeader("If-Match",autorest.String(ifMatch)))
        }
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CreateOrUpdateSparkJobDefinitionSender sends the CreateOrUpdateSparkJobDefinition request. The method will close the
    // http.Response Body if it receives an error.
    func (client SparkJobDefinitionClient) CreateOrUpdateSparkJobDefinitionSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            }

    // CreateOrUpdateSparkJobDefinitionResponder handles the response to the CreateOrUpdateSparkJobDefinition request. The method always
    // closes the http.Response Body.
    func (client SparkJobDefinitionClient) CreateOrUpdateSparkJobDefinitionResponder(resp *http.Response) (result SparkJobDefinitionResource, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// DebugSparkJobDefinition debug the spark job definition.
    // Parameters:
        // sparkJobDefinitionAzureResource - spark Job Definition resource definition.
func (client SparkJobDefinitionClient) DebugSparkJobDefinition(ctx context.Context, sparkJobDefinitionAzureResource SparkJobDefinitionResource) (result SparkJobDefinitionDebugSparkJobDefinitionFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/SparkJobDefinitionClient.DebugSparkJobDefinition")
        defer func() {
            sc := -1
        if result.Response() != nil {
        sc = result.Response().StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        if err := validation.Validate([]validation.Validation{
        { TargetValue: sparkJobDefinitionAzureResource,
         Constraints: []validation.Constraint{	{Target: "sparkJobDefinitionAzureResource.Properties", Name: validation.Null, Rule: true ,
        Chain: []validation.Constraint{	{Target: "sparkJobDefinitionAzureResource.Properties.TargetBigDataPool", Name: validation.Null, Rule: true ,
        Chain: []validation.Constraint{	{Target: "sparkJobDefinitionAzureResource.Properties.TargetBigDataPool.Type", Name: validation.Null, Rule: true, Chain: nil },
        	{Target: "sparkJobDefinitionAzureResource.Properties.TargetBigDataPool.ReferenceName", Name: validation.Null, Rule: true, Chain: nil },
        }},
        	{Target: "sparkJobDefinitionAzureResource.Properties.JobProperties", Name: validation.Null, Rule: true ,
        Chain: []validation.Constraint{	{Target: "sparkJobDefinitionAzureResource.Properties.JobProperties.File", Name: validation.Null, Rule: true, Chain: nil },
        	{Target: "sparkJobDefinitionAzureResource.Properties.JobProperties.DriverMemory", Name: validation.Null, Rule: true, Chain: nil },
        	{Target: "sparkJobDefinitionAzureResource.Properties.JobProperties.DriverCores", Name: validation.Null, Rule: true, Chain: nil },
        	{Target: "sparkJobDefinitionAzureResource.Properties.JobProperties.ExecutorMemory", Name: validation.Null, Rule: true, Chain: nil },
        	{Target: "sparkJobDefinitionAzureResource.Properties.JobProperties.ExecutorCores", Name: validation.Null, Rule: true, Chain: nil },
        	{Target: "sparkJobDefinitionAzureResource.Properties.JobProperties.NumExecutors", Name: validation.Null, Rule: true, Chain: nil },
        }},
        }}}}}); err != nil {
        return result, validation.NewError("synapse.SparkJobDefinitionClient", "DebugSparkJobDefinition", err.Error())
        }

        req, err := client.DebugSparkJobDefinitionPreparer(ctx, sparkJobDefinitionAzureResource)
    if err != nil {
    err = autorest.NewErrorWithError(err, "synapse.SparkJobDefinitionClient", "DebugSparkJobDefinition", nil , "Failure preparing request")
    return
    }

        result, err = client.DebugSparkJobDefinitionSender(req)
        if err != nil {
        err = autorest.NewErrorWithError(err, "synapse.SparkJobDefinitionClient", "DebugSparkJobDefinition", result.Response(), "Failure sending request")
        return
        }

    return
}

    // DebugSparkJobDefinitionPreparer prepares the DebugSparkJobDefinition request.
    func (client SparkJobDefinitionClient) DebugSparkJobDefinitionPreparer(ctx context.Context, sparkJobDefinitionAzureResource SparkJobDefinitionResource) (*http.Request, error) {
        urlParameters := map[string]interface{} {
        "endpoint": client.Endpoint,
        }

            const APIVersion = "2019-06-01-preview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsContentType("application/json; charset=utf-8"),
autorest.AsPost(),
autorest.WithCustomBaseURL("{endpoint}", urlParameters),
autorest.WithPath("/debugSparkJobDefinition"),
autorest.WithJSON(sparkJobDefinitionAzureResource),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // DebugSparkJobDefinitionSender sends the DebugSparkJobDefinition request. The method will close the
    // http.Response Body if it receives an error.
    func (client SparkJobDefinitionClient) DebugSparkJobDefinitionSender(req *http.Request) (future SparkJobDefinitionDebugSparkJobDefinitionFuture, err error) {
            var resp *http.Response
            resp, err = client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            if err != nil {
            return
            }
            future.Future, err = azure.NewFutureFromResponse(resp)
            return
            }

    // DebugSparkJobDefinitionResponder handles the response to the DebugSparkJobDefinition request. The method always
    // closes the http.Response Body.
    func (client SparkJobDefinitionClient) DebugSparkJobDefinitionResponder(resp *http.Response) (result SparkBatchJob, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusAccepted),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// DeleteSparkJobDefinition deletes a Spark Job Definition.
    // Parameters:
        // sparkJobDefinitionName - the spark job definition name.
func (client SparkJobDefinitionClient) DeleteSparkJobDefinition(ctx context.Context, sparkJobDefinitionName string) (result autorest.Response, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/SparkJobDefinitionClient.DeleteSparkJobDefinition")
        defer func() {
            sc := -1
        if result.Response != nil {
        sc = result.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.DeleteSparkJobDefinitionPreparer(ctx, sparkJobDefinitionName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "synapse.SparkJobDefinitionClient", "DeleteSparkJobDefinition", nil , "Failure preparing request")
    return
    }

        resp, err := client.DeleteSparkJobDefinitionSender(req)
        if err != nil {
        result.Response = resp
        err = autorest.NewErrorWithError(err, "synapse.SparkJobDefinitionClient", "DeleteSparkJobDefinition", resp, "Failure sending request")
        return
        }

        result, err = client.DeleteSparkJobDefinitionResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "synapse.SparkJobDefinitionClient", "DeleteSparkJobDefinition", resp, "Failure responding to request")
        }

    return
}

    // DeleteSparkJobDefinitionPreparer prepares the DeleteSparkJobDefinition request.
    func (client SparkJobDefinitionClient) DeleteSparkJobDefinitionPreparer(ctx context.Context, sparkJobDefinitionName string) (*http.Request, error) {
        urlParameters := map[string]interface{} {
        "endpoint": client.Endpoint,
        }

        pathParameters := map[string]interface{} {
        "sparkJobDefinitionName": autorest.Encode("path",sparkJobDefinitionName),
        }

            const APIVersion = "2019-06-01-preview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsDelete(),
autorest.WithCustomBaseURL("{endpoint}", urlParameters),
autorest.WithPathParameters("/sparkJobDefinitions/{sparkJobDefinitionName}",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // DeleteSparkJobDefinitionSender sends the DeleteSparkJobDefinition request. The method will close the
    // http.Response Body if it receives an error.
    func (client SparkJobDefinitionClient) DeleteSparkJobDefinitionSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            }

    // DeleteSparkJobDefinitionResponder handles the response to the DeleteSparkJobDefinition request. The method always
    // closes the http.Response Body.
    func (client SparkJobDefinitionClient) DeleteSparkJobDefinitionResponder(resp *http.Response) (result autorest.Response, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusNoContent),
            autorest.ByClosing())
            result.Response = resp
            return
    }

// ExecuteSparkJobDefinition executes the spark job definition.
    // Parameters:
        // sparkJobDefinitionName - the spark job definition name.
func (client SparkJobDefinitionClient) ExecuteSparkJobDefinition(ctx context.Context, sparkJobDefinitionName string) (result SparkJobDefinitionExecuteSparkJobDefinitionFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/SparkJobDefinitionClient.ExecuteSparkJobDefinition")
        defer func() {
            sc := -1
        if result.Response() != nil {
        sc = result.Response().StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.ExecuteSparkJobDefinitionPreparer(ctx, sparkJobDefinitionName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "synapse.SparkJobDefinitionClient", "ExecuteSparkJobDefinition", nil , "Failure preparing request")
    return
    }

        result, err = client.ExecuteSparkJobDefinitionSender(req)
        if err != nil {
        err = autorest.NewErrorWithError(err, "synapse.SparkJobDefinitionClient", "ExecuteSparkJobDefinition", result.Response(), "Failure sending request")
        return
        }

    return
}

    // ExecuteSparkJobDefinitionPreparer prepares the ExecuteSparkJobDefinition request.
    func (client SparkJobDefinitionClient) ExecuteSparkJobDefinitionPreparer(ctx context.Context, sparkJobDefinitionName string) (*http.Request, error) {
        urlParameters := map[string]interface{} {
        "endpoint": client.Endpoint,
        }

        pathParameters := map[string]interface{} {
        "sparkJobDefinitionName": autorest.Encode("path",sparkJobDefinitionName),
        }

            const APIVersion = "2019-06-01-preview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsPost(),
autorest.WithCustomBaseURL("{endpoint}", urlParameters),
autorest.WithPathParameters("/sparkJobDefinitions/{sparkJobDefinitionName}/execute",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ExecuteSparkJobDefinitionSender sends the ExecuteSparkJobDefinition request. The method will close the
    // http.Response Body if it receives an error.
    func (client SparkJobDefinitionClient) ExecuteSparkJobDefinitionSender(req *http.Request) (future SparkJobDefinitionExecuteSparkJobDefinitionFuture, err error) {
            var resp *http.Response
            resp, err = client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            if err != nil {
            return
            }
            future.Future, err = azure.NewFutureFromResponse(resp)
            return
            }

    // ExecuteSparkJobDefinitionResponder handles the response to the ExecuteSparkJobDefinition request. The method always
    // closes the http.Response Body.
    func (client SparkJobDefinitionClient) ExecuteSparkJobDefinitionResponder(resp *http.Response) (result SparkBatchJob, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusAccepted),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// GetSparkJobDefinition gets a Spark Job Definition.
    // Parameters:
        // sparkJobDefinitionName - the spark job definition name.
        // ifNoneMatch - eTag of the Spark Job Definition entity. Should only be specified for get. If the ETag matches
        // the existing entity tag, or if * was provided, then no content will be returned.
func (client SparkJobDefinitionClient) GetSparkJobDefinition(ctx context.Context, sparkJobDefinitionName string, ifNoneMatch string) (result SparkJobDefinitionResource, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/SparkJobDefinitionClient.GetSparkJobDefinition")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.GetSparkJobDefinitionPreparer(ctx, sparkJobDefinitionName, ifNoneMatch)
    if err != nil {
    err = autorest.NewErrorWithError(err, "synapse.SparkJobDefinitionClient", "GetSparkJobDefinition", nil , "Failure preparing request")
    return
    }

        resp, err := client.GetSparkJobDefinitionSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "synapse.SparkJobDefinitionClient", "GetSparkJobDefinition", resp, "Failure sending request")
        return
        }

        result, err = client.GetSparkJobDefinitionResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "synapse.SparkJobDefinitionClient", "GetSparkJobDefinition", resp, "Failure responding to request")
        }

    return
}

    // GetSparkJobDefinitionPreparer prepares the GetSparkJobDefinition request.
    func (client SparkJobDefinitionClient) GetSparkJobDefinitionPreparer(ctx context.Context, sparkJobDefinitionName string, ifNoneMatch string) (*http.Request, error) {
        urlParameters := map[string]interface{} {
        "endpoint": client.Endpoint,
        }

        pathParameters := map[string]interface{} {
        "sparkJobDefinitionName": autorest.Encode("path",sparkJobDefinitionName),
        }

            const APIVersion = "2019-06-01-preview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithCustomBaseURL("{endpoint}", urlParameters),
autorest.WithPathParameters("/sparkJobDefinitions/{sparkJobDefinitionName}",pathParameters),
autorest.WithQueryParameters(queryParameters))
        if len(ifNoneMatch) > 0 {
        preparer = autorest.DecoratePreparer(preparer,
        autorest.WithHeader("If-None-Match",autorest.String(ifNoneMatch)))
        }
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSparkJobDefinitionSender sends the GetSparkJobDefinition request. The method will close the
    // http.Response Body if it receives an error.
    func (client SparkJobDefinitionClient) GetSparkJobDefinitionSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            }

    // GetSparkJobDefinitionResponder handles the response to the GetSparkJobDefinition request. The method always
    // closes the http.Response Body.
    func (client SparkJobDefinitionClient) GetSparkJobDefinitionResponder(resp *http.Response) (result SparkJobDefinitionResource, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusNotModified),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// GetSparkJobDefinitionsByWorkspace lists spark job definitions.
func (client SparkJobDefinitionClient) GetSparkJobDefinitionsByWorkspace(ctx context.Context) (result SparkJobDefinitionsListResponsePage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/SparkJobDefinitionClient.GetSparkJobDefinitionsByWorkspace")
        defer func() {
            sc := -1
        if result.sjdlr.Response.Response != nil {
        sc = result.sjdlr.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        result.fn = client.getSparkJobDefinitionsByWorkspaceNextResults
    req, err := client.GetSparkJobDefinitionsByWorkspacePreparer(ctx)
    if err != nil {
    err = autorest.NewErrorWithError(err, "synapse.SparkJobDefinitionClient", "GetSparkJobDefinitionsByWorkspace", nil , "Failure preparing request")
    return
    }

        resp, err := client.GetSparkJobDefinitionsByWorkspaceSender(req)
        if err != nil {
        result.sjdlr.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "synapse.SparkJobDefinitionClient", "GetSparkJobDefinitionsByWorkspace", resp, "Failure sending request")
        return
        }

        result.sjdlr, err = client.GetSparkJobDefinitionsByWorkspaceResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "synapse.SparkJobDefinitionClient", "GetSparkJobDefinitionsByWorkspace", resp, "Failure responding to request")
        }
            if result.sjdlr.hasNextLink() && result.sjdlr.IsEmpty() {
            err = result.NextWithContext(ctx)
            }

    return
}

    // GetSparkJobDefinitionsByWorkspacePreparer prepares the GetSparkJobDefinitionsByWorkspace request.
    func (client SparkJobDefinitionClient) GetSparkJobDefinitionsByWorkspacePreparer(ctx context.Context) (*http.Request, error) {
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
autorest.WithPath("/sparkJobDefinitions"),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSparkJobDefinitionsByWorkspaceSender sends the GetSparkJobDefinitionsByWorkspace request. The method will close the
    // http.Response Body if it receives an error.
    func (client SparkJobDefinitionClient) GetSparkJobDefinitionsByWorkspaceSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            }

    // GetSparkJobDefinitionsByWorkspaceResponder handles the response to the GetSparkJobDefinitionsByWorkspace request. The method always
    // closes the http.Response Body.
    func (client SparkJobDefinitionClient) GetSparkJobDefinitionsByWorkspaceResponder(resp *http.Response) (result SparkJobDefinitionsListResponse, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

            // getSparkJobDefinitionsByWorkspaceNextResults retrieves the next set of results, if any.
            func (client SparkJobDefinitionClient) getSparkJobDefinitionsByWorkspaceNextResults(ctx context.Context, lastResults SparkJobDefinitionsListResponse) (result SparkJobDefinitionsListResponse, err error) {
            req, err := lastResults.sparkJobDefinitionsListResponsePreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "synapse.SparkJobDefinitionClient", "getSparkJobDefinitionsByWorkspaceNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.GetSparkJobDefinitionsByWorkspaceSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "synapse.SparkJobDefinitionClient", "getSparkJobDefinitionsByWorkspaceNextResults", resp, "Failure sending next results request")
            }
            result, err = client.GetSparkJobDefinitionsByWorkspaceResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "synapse.SparkJobDefinitionClient", "getSparkJobDefinitionsByWorkspaceNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

            // GetSparkJobDefinitionsByWorkspaceComplete enumerates all values, automatically crossing page boundaries as required.
            func (client SparkJobDefinitionClient) GetSparkJobDefinitionsByWorkspaceComplete(ctx context.Context) (result SparkJobDefinitionsListResponseIterator, err error) {
            if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/SparkJobDefinitionClient.GetSparkJobDefinitionsByWorkspace")
            defer func() {
            sc := -1
            if result.Response().Response.Response != nil {
            sc = result.page.Response().Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
            }()
            }
                    result.page, err = client.GetSparkJobDefinitionsByWorkspace(ctx)
                            return
            }

