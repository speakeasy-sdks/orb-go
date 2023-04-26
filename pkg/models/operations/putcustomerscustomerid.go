// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Orb/pkg/models/shared"
	"encoding/json"
	"fmt"
	"net/http"
)

// PutCustomersCustomerIDRequestBodyBillingAddress - The customer's billing address; all fields in the address are optional. This address appears on customer invoices.
type PutCustomersCustomerIDRequestBodyBillingAddress struct {
	City *string `json:"city,omitempty"`
	// Two-letter country code (ISO 3166-1 alpha-2).
	Country *string `json:"country,omitempty"`
	Line1   *string `json:"line1,omitempty"`
	Line2   *string `json:"line2,omitempty"`
	// ZIP or postal code
	PostalCode *string `json:"postal_code,omitempty"`
	State      *string `json:"state,omitempty"`
}

// PutCustomersCustomerIDRequestBodyPaymentProviderEnum - This is used for creating charges or invoices in an external system via Orb. When not in test mode:
// - the connection must first be configured in the Orb webapp.
// - if the provider is an invoicing provider (`stripe_invoice`, `quickbooks`, `bill.com`), any product mappings must first be configured with the Orb team.
type PutCustomersCustomerIDRequestBodyPaymentProviderEnum string

const (
	PutCustomersCustomerIDRequestBodyPaymentProviderEnumStripeInvoice PutCustomersCustomerIDRequestBodyPaymentProviderEnum = "stripe_invoice"
	PutCustomersCustomerIDRequestBodyPaymentProviderEnumQuickbooks    PutCustomersCustomerIDRequestBodyPaymentProviderEnum = "quickbooks"
	PutCustomersCustomerIDRequestBodyPaymentProviderEnumBillCom       PutCustomersCustomerIDRequestBodyPaymentProviderEnum = "bill.com"
	PutCustomersCustomerIDRequestBodyPaymentProviderEnumStripeCharge  PutCustomersCustomerIDRequestBodyPaymentProviderEnum = "stripe_charge"
	PutCustomersCustomerIDRequestBodyPaymentProviderEnumNull          PutCustomersCustomerIDRequestBodyPaymentProviderEnum = "null"
)

func (e PutCustomersCustomerIDRequestBodyPaymentProviderEnum) ToPointer() *PutCustomersCustomerIDRequestBodyPaymentProviderEnum {
	return &e
}

func (e *PutCustomersCustomerIDRequestBodyPaymentProviderEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "stripe_invoice":
		fallthrough
	case "quickbooks":
		fallthrough
	case "bill.com":
		fallthrough
	case "stripe_charge":
		fallthrough
	case "null":
		*e = PutCustomersCustomerIDRequestBodyPaymentProviderEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for PutCustomersCustomerIDRequestBodyPaymentProviderEnum: %s", s)
	}
}

// PutCustomersCustomerIDRequestBodyShippingAddress - The customer's shipping address; all fields in the address are optional. Note that downstream tax calculations are based on the shipping address.
type PutCustomersCustomerIDRequestBodyShippingAddress struct {
	City *string `json:"city,omitempty"`
	// Two-letter country code (ISO 3166-1 alpha-2).
	Country *string `json:"country,omitempty"`
	Line1   *string `json:"line1,omitempty"`
	Line2   *string `json:"line2,omitempty"`
	// ZIP or postal code
	PostalCode *string `json:"postal_code,omitempty"`
	State      *string `json:"state,omitempty"`
}

// PutCustomersCustomerIDRequestBody - The external payments or invoicing solution connected to this customer.
type PutCustomersCustomerIDRequestBody struct {
	// The customer's billing address; all fields in the address are optional. This address appears on customer invoices.
	BillingAddress *PutCustomersCustomerIDRequestBodyBillingAddress `json:"billing_address,omitempty"`
	// A valid customer email, to be used for invoicing and notifications.
	Email *string `json:"email,omitempty"`
	Name  *string `json:"name,omitempty"`
	// This is used for creating charges or invoices in an external system via Orb. When not in test mode:
	// - the connection must first be configured in the Orb webapp.
	// - if the provider is an invoicing provider (`stripe_invoice`, `quickbooks`, `bill.com`), any product mappings must first be configured with the Orb team.
	PaymentProvider *PutCustomersCustomerIDRequestBodyPaymentProviderEnum `json:"payment_provider,omitempty"`
	// The ID of this customer in an external payments solution, such as Stripe. This is used for creating charges or invoices in the external system via Orb.
	PaymentProviderID *string `json:"payment_provider_id,omitempty"`
	// The customer's shipping address; all fields in the address are optional. Note that downstream tax calculations are based on the shipping address.
	ShippingAddress *PutCustomersCustomerIDRequestBodyShippingAddress `json:"shipping_address,omitempty"`
}

type PutCustomersCustomerIDRequest struct {
	RequestBody *PutCustomersCustomerIDRequestBody `request:"mediaType=application/json"`
	// Orb customer ID
	CustomerID string `pathParam:"style=simple,explode=false,name=customer_id"`
}

type PutCustomersCustomerIDResponse struct {
	ContentType string
	// OK
	Customer    *shared.Customer
	StatusCode  int
	RawResponse *http.Response
}
