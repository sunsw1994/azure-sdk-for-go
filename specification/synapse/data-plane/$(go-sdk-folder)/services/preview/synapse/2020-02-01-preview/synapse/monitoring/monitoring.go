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
)

// MonitoringClient is the client for the Monitoring methods of the Synapse service.
type MonitoringClient struct {
    BaseClient
}
// NewMonitoringClient creates an instance of the MonitoringClient client.
func NewMonitoringClient(endpoint string) MonitoringClient {
    return MonitoringClient{ New(endpoint)}
}

// GetSparkJobList get list of spark applications for the workspace.
    // Parameters:
        // xMsClientRequestID - can provide a guid, which is helpful for debugging and to provide better customer
        // support
func (client MonitoringClient) GetSparkJobList(ctx context.Context, APIVersion string, xMsClientRequestID string) (result SparkJobListViewResponse, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/MonitoringClient.GetSparkJobList")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.GetSparkJobListPreparer(ctx, APIVersion, xMsClientRequestID)
    if err != nil {
    err = autorest.NewErrorWithError(err, "synapse.MonitoringClient", "GetSparkJobList", nil , "Failure preparing request")
    return
    }

        resp, err := client.GetSparkJobListSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "synapse.MonitoringClient", "GetSparkJobList", resp, "Failure sending request")
        return
        }

        result, err = client.GetSparkJobListResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "synapse.MonitoringClient", "GetSparkJobList", resp, "Failure responding to request")
        }

    return
}

    // GetSparkJobListPreparer prepares the GetSparkJobList request.
    func (client MonitoringClient) GetSparkJobListPreparer(ctx context.Context, APIVersion string, xMsClientRequestID string) (*http.Request, error) {
        urlParameters := map[string]interface{} {
        "endpoint": client.Endpoint,
        }

        queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithCustomBaseURL("{endpoint}", urlParameters),
autorest.WithPath("/monitoring/workloadTypes/spark/Applications"),
autorest.WithQueryParameters(queryParameters))
        if len(xMsClientRequestID) > 0 {
        preparer = autorest.DecoratePreparer(preparer,
        autorest.WithHeader("x-ms-client-request-id",autorest.String(xMsClientRequestID)))
        }
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSparkJobListSender sends the GetSparkJobList request. The method will close the
    // http.Response Body if it receives an error.
    func (client MonitoringClient) GetSparkJobListSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            }

    // GetSparkJobListResponder handles the response to the GetSparkJobList request. The method always
    // closes the http.Response Body.
    func (client MonitoringClient) GetSparkJobListResponder(resp *http.Response) (result SparkJobListViewResponse, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// GetSQLJobQueryString get SQL OD/DW Query for the workspace.
    // Parameters:
        // xMsClientRequestID - can provide a guid, which is helpful for debugging and to provide better customer
        // support
func (client MonitoringClient) GetSQLJobQueryString(ctx context.Context, APIVersion string, xMsClientRequestID string, filter string, orderby string, skip string) (result SQLQueryStringDataModel, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/MonitoringClient.GetSQLJobQueryString")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.GetSQLJobQueryStringPreparer(ctx, APIVersion, xMsClientRequestID, filter, orderby, skip)
    if err != nil {
    err = autorest.NewErrorWithError(err, "synapse.MonitoringClient", "GetSQLJobQueryString", nil , "Failure preparing request")
    return
    }

        resp, err := client.GetSQLJobQueryStringSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "synapse.MonitoringClient", "GetSQLJobQueryString", resp, "Failure sending request")
        return
        }

        result, err = client.GetSQLJobQueryStringResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "synapse.MonitoringClient", "GetSQLJobQueryString", resp, "Failure responding to request")
        }

    return
}

    // GetSQLJobQueryStringPreparer prepares the GetSQLJobQueryString request.
    func (client MonitoringClient) GetSQLJobQueryStringPreparer(ctx context.Context, APIVersion string, xMsClientRequestID string, filter string, orderby string, skip string) (*http.Request, error) {
        urlParameters := map[string]interface{} {
        "endpoint": client.Endpoint,
        }

        queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }
        if len(filter) > 0 {
        queryParameters["filter"] = autorest.Encode("query",filter)
        }
        if len(orderby) > 0 {
        queryParameters["$orderby"] = autorest.Encode("query",orderby)
        }
        if len(skip) > 0 {
        queryParameters["skip"] = autorest.Encode("query",skip)
        }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithCustomBaseURL("{endpoint}", urlParameters),
autorest.WithPath("/monitoring/workloadTypes/sql/querystring"),
autorest.WithQueryParameters(queryParameters))
        if len(xMsClientRequestID) > 0 {
        preparer = autorest.DecoratePreparer(preparer,
        autorest.WithHeader("x-ms-client-request-id",autorest.String(xMsClientRequestID)))
        }
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSQLJobQueryStringSender sends the GetSQLJobQueryString request. The method will close the
    // http.Response Body if it receives an error.
    func (client MonitoringClient) GetSQLJobQueryStringSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            }

    // GetSQLJobQueryStringResponder handles the response to the GetSQLJobQueryString request. The method always
    // closes the http.Response Body.
    func (client MonitoringClient) GetSQLJobQueryStringResponder(resp *http.Response) (result SQLQueryStringDataModel, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

