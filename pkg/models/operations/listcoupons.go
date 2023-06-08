// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Orb/pkg/models/shared"
	"net/http"
)

type ListCouponsRequest struct {
	// Filter to coupons matching this redemption code.
	RedemptionCode *string `queryParam:"style=form,explode=true,name=redemption_code"`
	// Show archived coupons as well (by default, this endpoint only returns active coupons).
	ShowArchived *bool `queryParam:"style=form,explode=true,name=show_archived"`
}

// ListCoupons200ApplicationJSON - OK
type ListCoupons200ApplicationJSON struct {
	Data               []shared.Coupon            `json:"data,omitempty"`
	PaginationMetadata *shared.PaginationMetadata `json:"pagination_metadata,omitempty"`
}

type ListCouponsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	ListCoupons200ApplicationJSONObject *ListCoupons200ApplicationJSON
}