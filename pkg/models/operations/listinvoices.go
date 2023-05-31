// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Orb/pkg/models/shared"
	"net/http"
)

type ListInvoicesRequest struct {
	// Filter by a specific customer (cannot be used with `external_customer_id`)
	CustomerID *string `queryParam:"style=form,explode=true,name=customer_id"`
	// Filter by a specific customer (cannot be used with `customer_id`)
	ExternalCustomerID *string `queryParam:"style=form,explode=true,name=external_customer_id"`
	// Filter to invoices of a particular status (`draft`, `issued`, `paid`, `void`, `synced`); this param can be used to filter to a single status (`?status="draft"`) or a set of statuses (`?status[]=paid&status[]=void`)
	Status interface{} `queryParam:"style=form,explode=true,name=status"`
	// Filter by a specific subscription
	SubscriptionID *string `queryParam:"style=form,explode=true,name=subscription_id"`
}

// ListInvoices200ApplicationJSON - OK
type ListInvoices200ApplicationJSON struct {
	Data               []shared.Invoice           `json:"data,omitempty"`
	PaginationMetadata *shared.PaginationMetadata `json:"pagination_metadata,omitempty"`
}

type ListInvoicesResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	ListInvoices200ApplicationJSONObject *ListInvoices200ApplicationJSON
}
