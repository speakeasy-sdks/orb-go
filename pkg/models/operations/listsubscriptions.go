// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Orb/pkg/models/shared"
	"net/http"
)

type ListSubscriptionsRequest struct {
	CustomerID         *string `queryParam:"style=form,explode=true,name=customer_id"`
	ExternalCustomerID *string `queryParam:"style=form,explode=true,name=external_customer_id"`
}

// ListSubscriptions200ApplicationJSON - OK
type ListSubscriptions200ApplicationJSON struct {
	Data               []shared.Subscription  `json:"data,omitempty"`
	PaginationMetadata map[string]interface{} `json:"pagination_metadata,omitempty"`
}

type ListSubscriptionsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	ListSubscriptions200ApplicationJSONObject *ListSubscriptions200ApplicationJSON
}
