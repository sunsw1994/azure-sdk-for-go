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
    "context"
    "github.com/Azure/go-autorest/autorest"
    "github.com/Azure/go-autorest/autorest/to"
    "github.com/Azure/go-autorest/tracing"
    "net/http"
)

// The package's fully qualified name.
const fqdn = "$(go-sdk-folder)/services/preview/synapse/2020-02-01-preview/synapse/accesscontrol"

        // ErrorContract contains details when the response code indicates an error.
        type ErrorContract struct {
        // Error - The error details.
        Error *ErrorResponse `json:"error,omitempty"`
        }

        // ErrorDetail ...
        type ErrorDetail struct {
        Code *string `json:"code,omitempty"`
        Message *string `json:"message,omitempty"`
        Target *string `json:"target,omitempty"`
        }

        // ErrorResponse ...
        type ErrorResponse struct {
        Code *string `json:"code,omitempty"`
        Message *string `json:"message,omitempty"`
        Target *string `json:"target,omitempty"`
        Details *[]ErrorDetail `json:"details,omitempty"`
        }

        // ListRoleAssignmentDetails ...
        type ListRoleAssignmentDetails struct {
        autorest.Response `json:"-"`
        Value *[]RoleAssignmentDetails `json:"value,omitempty"`
        }

        // ListString ...
        type ListString struct {
        autorest.Response `json:"-"`
        Value *[]string `json:"value,omitempty"`
        }

        // Role synapse role details
        type Role struct {
        autorest.Response `json:"-"`
        // ID - Role ID
        ID *string `json:"id,omitempty"`
        // Name - Name of the Synapse role
        Name *string `json:"name,omitempty"`
        // IsBuiltIn - Is a built-in role or not
        IsBuiltIn *bool `json:"isBuiltIn,omitempty"`
        }

        // RoleAssignmentDetails role Assignment response details
        type RoleAssignmentDetails struct {
        autorest.Response `json:"-"`
        // ID - Role Assignment ID
        ID *string `json:"id,omitempty"`
        // RoleID - Role ID of the Synapse Built-In Role
        RoleID *string `json:"roleId,omitempty"`
        // PrincipalID - Object ID of the AAD principal or security-group
        PrincipalID *string `json:"principalId,omitempty"`
        }

        // RoleAssignmentOptions role Assignment request details
        type RoleAssignmentOptions struct {
        // RoleID - Role ID of the Synapse Built-In Role
        RoleID *string `json:"roleId,omitempty"`
        // PrincipalID - Object ID of the AAD principal or security-group
        PrincipalID *string `json:"principalId,omitempty"`
        }

        // RolesListResponse a list of Synapse roles available.
        type RolesListResponse struct {
        autorest.Response `json:"-"`
        // Value - List of Synapse roles.
        Value *[]Role `json:"value,omitempty"`
        // NextLink - The link to the next page of results, if any remaining results exist.
        NextLink *string `json:"nextLink,omitempty"`
        }

        // RolesListResponseIterator provides access to a complete listing of Role values.
        type RolesListResponseIterator struct {
            i int
            page RolesListResponsePage
        }
        // NextWithContext advances to the next value.  If there was an error making
        // the request the iterator does not advance and the error is returned.
        func (iter * RolesListResponseIterator) NextWithContext(ctx context.Context) (err error) {
        if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/RolesListResponseIterator.NextWithContext")
        defer func() {
        sc := -1
        if iter.Response().Response.Response != nil {
        sc = iter.Response().Response.Response.StatusCode
        }
        tracing.EndSpan(ctx, sc, err)
        }()
        }
        iter.i++
        if iter.i < len(iter. page.Values()) {
        return nil
        }
        err = iter.page.NextWithContext(ctx)
        if err != nil {
        iter. i--
        return err
        }
        iter.i = 0
        return nil
        }
        // Next advances to the next value.  If there was an error making
        // the request the iterator does not advance and the error is returned.
        // Deprecated: Use NextWithContext() instead.
        func (iter * RolesListResponseIterator) Next() error {
        return iter.NextWithContext(context.Background())
        }
        // NotDone returns true if the enumeration should be started or is not yet complete.
        func (iter RolesListResponseIterator) NotDone() bool {
        return iter.page.NotDone() && iter.i < len(iter. page.Values())
        }
        // Response returns the raw server response from the last page request.
        func (iter RolesListResponseIterator) Response() RolesListResponse {
        return iter.page.Response()
        }
        // Value returns the current value or a zero-initialized value if the
        // iterator has advanced beyond the end of the collection.
        func (iter RolesListResponseIterator) Value() Role {
        if !iter.page.NotDone() {
        return Role{}
        }
        return iter.page.Values()[iter.i]
        }
        // Creates a new instance of the RolesListResponseIterator type.
        func NewRolesListResponseIterator (page RolesListResponsePage) RolesListResponseIterator {
            return RolesListResponseIterator{page: page}
        }


            // IsEmpty returns true if the ListResult contains no values.
            func (rlr RolesListResponse) IsEmpty() bool {
            return rlr.Value == nil || len(*rlr.Value) == 0
            }

            // hasNextLink returns true if the NextLink is not empty.
            func (rlr RolesListResponse) hasNextLink() bool {
            return rlr.NextLink != nil && len(*rlr.NextLink) != 0
            }
                // rolesListResponsePreparer prepares a request to retrieve the next set of results.
                // It returns nil if no more results exist.
                func (rlr RolesListResponse) rolesListResponsePreparer(ctx context.Context) (*http.Request, error) {
                if !rlr.hasNextLink() {
                return nil, nil
                }
                return autorest.Prepare((&http.Request{}).WithContext(ctx),
                autorest.AsJSON(),
                autorest.AsGet(),
                autorest.WithBaseURL(to.String( rlr.NextLink)));
                }

        // RolesListResponsePage contains a page of Role values.
        type RolesListResponsePage struct {
            fn func(context.Context, RolesListResponse) (RolesListResponse, error)
            rlr RolesListResponse
        }

        // NextWithContext advances to the next page of values.  If there was an error making
        // the request the page does not advance and the error is returned.
        func (page * RolesListResponsePage) NextWithContext(ctx context.Context) (err error) {
        if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/RolesListResponsePage.NextWithContext")
        defer func() {
        sc := -1
        if page.Response().Response.Response != nil {
        sc = page.Response().Response.Response.StatusCode
        }
        tracing.EndSpan(ctx, sc, err)
        }()
        }
        for {
            next, err := page.fn(ctx, page.rlr)
            if err != nil {
            return err
            }
            page.rlr = next
            if !next.hasNextLink() || !next.IsEmpty() {
                break
            }
        }
        return nil
        }

        // Next advances to the next page of values.  If there was an error making
        // the request the page does not advance and the error is returned.
        // Deprecated: Use NextWithContext() instead.
        func (page * RolesListResponsePage) Next() error {
        return page.NextWithContext(context.Background())
        }
        // NotDone returns true if the page enumeration should be started or is not yet complete.
        func (page RolesListResponsePage) NotDone() bool {
        return !page.rlr.IsEmpty()
        }
        // Response returns the raw server response from the last page request.
        func (page RolesListResponsePage) Response() RolesListResponse {
        return page.rlr
        }
        // Values returns the slice of values for the current page or nil if there are no values.
        func (page RolesListResponsePage) Values() []Role {
        if page.rlr.IsEmpty() {
        return nil
        }
        return *page.rlr.Value
        }
        // Creates a new instance of the RolesListResponsePage type.
        func NewRolesListResponsePage (getNextPage func(context.Context, RolesListResponse) (RolesListResponse, error)) RolesListResponsePage {
            return RolesListResponsePage{fn: getNextPage}
        }

