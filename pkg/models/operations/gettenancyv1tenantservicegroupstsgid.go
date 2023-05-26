// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"sase-tenancy/pkg/models/shared"
)

type GetTenancyV1TenantServiceGroupsTsgIDRequest struct {
	// A unique identifier for the tenant service group.
	//
	TsgID string `pathParam:"style=simple,explode=false,name=tsg_id"`
}

type GetTenancyV1TenantServiceGroupsTsgIDResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Successful response.
	TenantServiceGroup *shared.TenantServiceGroup
}
