// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"sase-tenancy/pkg/models/shared"
)

type GetTenancyV1TenantServiceGroupsSecurity struct {
	Bearer string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
}

type GetTenancyV1TenantServiceGroupsRequest struct {
	// Indicates whether the response structure lists groups in
	// their hierarchy, or as an array of TSGs without regard to
	// hierarchy. Default is false (don't show hierarchy).
	//
	// If false, the order of the TSGs in the result array is not
	// guaranteed.
	//
	Hierarchy *bool `queryParam:"style=form,explode=true,name=hierarchy"`
}

// GetTenancyV1TenantServiceGroups200ApplicationJSON - Successful response.
type GetTenancyV1TenantServiceGroups200ApplicationJSON struct {
	// Total count of the items
	Count int64                       `json:"count"`
	Items []shared.TenantServiceGroup `json:"items"`
}

type GetTenancyV1TenantServiceGroupsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Successful response.
	GetTenancyV1TenantServiceGroups200ApplicationJSONObject *GetTenancyV1TenantServiceGroups200ApplicationJSON
}
