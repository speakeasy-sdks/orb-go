// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Orb/pkg/models/shared"
	"Orb/pkg/types"
	"net/http"
)

type CreateInvoiceLineItemRequestBody struct {
	// The total amount in the invoice's currency to add to the line item.
	Amount string `json:"amount"`
	// An date string to specify the line item's end date in the customer's timezone.
	EndDate types.Date `json:"end_date"`
	// The id of the `Invoice` to add this line item.
	InvoiceID string `json:"invoice_id"`
	// The item name associated with this line item. If an item with the same name exists in Orb, that item will be associated with the line item.
	Name string `json:"name"`
	// The number of units on the line item
	Quantity float64 `json:"quantity"`
	// An date string to specify the line item's start date in the customer's timezone.
	StartDate types.Date `json:"start_date"`
}

type CreateInvoiceLineItemResponse struct {
	ContentType string
	// Created
	InvoiceLineItem *shared.InvoiceLineItem
	StatusCode      int
	RawResponse     *http.Response
}
