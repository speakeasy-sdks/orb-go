// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Orb/pkg/models/shared"
	"net/http"
)

type FetchUpcomingInvoiceRequest struct {
	SubscriptionID string `queryParam:"style=form,explode=true,name=subscription_id"`
}

type FetchUpcomingInvoiceResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	UpcomingInvoice *shared.UpcomingInvoice
}