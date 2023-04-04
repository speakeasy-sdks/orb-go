// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/orb-go/pkg/types"
	"time"
)

// UpcomingInvoiceCustomer - The customer receiving this invoice.
type UpcomingInvoiceCustomer struct {
	ExternalCustomerID string `json:"external_customer_id"`
	ID                 string `json:"id"`
}

// UpcomingInvoiceLineItemsGrouping - For configured prices that are split by a grouping key, this will be populated with the key and a value. The `amount` and `subtotal` will be the values for this particular grouping.
type UpcomingInvoiceLineItemsGrouping struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// UpcomingInvoiceLineItemsSubLineItemsMatrixConfig - Only available if `type` is `matrix`. Contains the values of the matrix that this `sub_line_item` represents.
type UpcomingInvoiceLineItemsSubLineItemsMatrixConfig struct {
	// The ordered dimension values for this line item.
	DimensionValues []string `json:"dimension_values"`
}

// UpcomingInvoiceLineItemsSubLineItemsTierConfig - Only available if `type` is `tier`. Contains the range of units in this tier and the unit amount.
type UpcomingInvoiceLineItemsSubLineItemsTierConfig struct {
	FirstUnit  float64 `json:"first_unit"`
	LastUnit   float64 `json:"last_unit"`
	UnitAmount string  `json:"unit_amount"`
}

// UpcomingInvoiceLineItemsSubLineItemsTypeEnum - An identifier for a sub line item that is specific to a pricing model.
type UpcomingInvoiceLineItemsSubLineItemsTypeEnum string

const (
	UpcomingInvoiceLineItemsSubLineItemsTypeEnumMatrix UpcomingInvoiceLineItemsSubLineItemsTypeEnum = "matrix"
	UpcomingInvoiceLineItemsSubLineItemsTypeEnumTier   UpcomingInvoiceLineItemsSubLineItemsTypeEnum = "tier"
)

func (e *UpcomingInvoiceLineItemsSubLineItemsTypeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "matrix":
		fallthrough
	case "tier":
		*e = UpcomingInvoiceLineItemsSubLineItemsTypeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for UpcomingInvoiceLineItemsSubLineItemsTypeEnum: %s", s)
	}
}

type UpcomingInvoiceLineItemsSubLineItems struct {
	// The total amount for this sub line item.
	Amount string `json:"amount"`
	// Only available if `type` is `matrix`. Contains the values of the matrix that this `sub_line_item` represents.
	MatrixConfig *UpcomingInvoiceLineItemsSubLineItemsMatrixConfig `json:"matrix_config,omitempty"`
	Name         string                                            `json:"name"`
	Quantity     float64                                           `json:"quantity"`
	// Only available if `type` is `tier`. Contains the range of units in this tier and the unit amount.
	TierConfig *UpcomingInvoiceLineItemsSubLineItemsTierConfig `json:"tier_config,omitempty"`
	// An identifier for a sub line item that is specific to a pricing model.
	Type UpcomingInvoiceLineItemsSubLineItemsTypeEnum `json:"type"`
}

type UpcomingInvoiceLineItems struct {
	// The final amount after any discounts or minimums.
	Amount string `json:"amount"`
	// The end date of the range of time applied for this line item's price.
	EndDate time.Time `json:"end_date"`
	// For configured prices that are split by a grouping key, this will be populated with the key and a value. The `amount` and `subtotal` will be the values for this particular grouping.
	Grouping UpcomingInvoiceLineItemsGrouping `json:"grouping"`
	// The name of the price associated with this line item.
	Name     string  `json:"name"`
	Quantity float64 `json:"quantity"`
	// The start date of the range of time applied for this line item's price.
	StartDate time.Time `json:"start_date"`
	// For complex pricing structures, the line item can be broken down further in `sub_line_items`.
	SubLineItems []UpcomingInvoiceLineItemsSubLineItems `json:"sub_line_items"`
	// The line amount before any line item-specific discounts or minimums.
	Subtotal string `json:"subtotal"`
}

// UpcomingInvoiceSubscription - The associated subscription for this invoice.
type UpcomingInvoiceSubscription struct {
	ID string `json:"id"`
}

// UpcomingInvoice - Upcoming invoices contain a line-by-line breakdown of an upcoming amount due for a subscription, including incurred usage in the billing period as well as any recurring fees.
//
// Although each line item will be invoiced on the `target_date`, each invoice line item may have distinct date ranges (e.g. for usage billed in arrears, the range may specify the current month whereas an in-advance recurring fees will be for the following month).
//
// Since an invoice resource has not been created for this upcoming invoice, this endpoint intentionally does not return an `id` field.
type UpcomingInvoice struct {
	// The final amount to be charged to the customer after all minimums and discounts have been applied. Only populated for non-pre-paid plans.
	AmountDue string `json:"amount_due"`
	// An ISO 4217 currency string or `credits`
	Currency string `json:"currency"`
	// The customer receiving this invoice.
	Customer UpcomingInvoiceCustomer `json:"customer"`
	// The breakdown of prices in this invoice.
	LineItems []UpcomingInvoiceLineItems `json:"line_items"`
	// The associated subscription for this invoice.
	Subscription UpcomingInvoiceSubscription `json:"subscription"`
	// The total before any discounts and minimums are applied.
	Subtotal string `json:"subtotal"`
	// The expected issue date of the invoice.
	TargetDate types.Date `json:"target_date"`
}
