// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Orb/pkg/models/shared"
	"net/http"
)

type GetCustomersCustomerIDRequest struct {
	// Orb customer ID
	CustomerID string `pathParam:"style=simple,explode=false,name=customer_id"`
}

type GetCustomersCustomerIDResponse struct {
	ContentType string
	// Created
	Customer    *shared.Customer
	StatusCode  int
	RawResponse *http.Response
}