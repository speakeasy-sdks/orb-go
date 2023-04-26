// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"time"
)

// InvoiceCustomer - The customer receiving this invoice.
type InvoiceCustomer struct {
	ExternalCustomerID string `json:"external_customer_id"`
	ID                 string `json:"id"`
}

// InvoiceLineItemsGrouping - For configured prices that are split by a grouping key, this will be populated with the key and a value. The `amount` and `subtotal` will be the values for this particular grouping.
type InvoiceLineItemsGrouping struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// InvoiceLineItemsSubLineItemsMatrixConfig - Only available if `type` is `matrix`. Contains the values of the matrix that this `sub_line_item` represents.
type InvoiceLineItemsSubLineItemsMatrixConfig struct {
	// The ordered dimension values for this line item.
	DimensionValues []string `json:"dimension_values"`
}

// InvoiceLineItemsSubLineItemsTierConfig - Only available if `type` is `tier`. Contains the range of units in this tier and the unit amount.
type InvoiceLineItemsSubLineItemsTierConfig struct {
	FirstUnit  float64 `json:"first_unit"`
	LastUnit   float64 `json:"last_unit"`
	UnitAmount string  `json:"unit_amount"`
}

// InvoiceLineItemsSubLineItemsTypeEnum - An identifier for a sub line item that is specific to a pricing model.
type InvoiceLineItemsSubLineItemsTypeEnum string

const (
	InvoiceLineItemsSubLineItemsTypeEnumMatrix InvoiceLineItemsSubLineItemsTypeEnum = "matrix"
	InvoiceLineItemsSubLineItemsTypeEnumTier   InvoiceLineItemsSubLineItemsTypeEnum = "tier"
)

func (e InvoiceLineItemsSubLineItemsTypeEnum) ToPointer() *InvoiceLineItemsSubLineItemsTypeEnum {
	return &e
}

func (e *InvoiceLineItemsSubLineItemsTypeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "matrix":
		fallthrough
	case "tier":
		*e = InvoiceLineItemsSubLineItemsTypeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for InvoiceLineItemsSubLineItemsTypeEnum: %s", s)
	}
}

type InvoiceLineItemsSubLineItems struct {
	// The total amount for this sub line item.
	Amount string `json:"amount"`
	// Only available if `type` is `matrix`. Contains the values of the matrix that this `sub_line_item` represents.
	MatrixConfig *InvoiceLineItemsSubLineItemsMatrixConfig `json:"matrix_config,omitempty"`
	Name         string                                    `json:"name"`
	Quantity     float64                                   `json:"quantity"`
	// Only available if `type` is `tier`. Contains the range of units in this tier and the unit amount.
	TierConfig *InvoiceLineItemsSubLineItemsTierConfig `json:"tier_config,omitempty"`
	// An identifier for a sub line item that is specific to a pricing model.
	Type InvoiceLineItemsSubLineItemsTypeEnum `json:"type"`
}

type InvoiceLineItemsTaxAmounts struct {
	// The amount of additional tax incurred by this tax rate.
	Amount string `json:"amount"`
	// The human-readable description of the applied tax rate.
	TaxRateDescription string `json:"tax_rate_description"`
	// The tax rate percentage, out of 100.
	TaxRatePercentage string `json:"tax_rate_percentage"`
}

type InvoiceLineItems struct {
	// The final amount after any discounts or minimums.
	Amount   string                 `json:"amount"`
	Discount map[string]interface{} `json:"discount"`
	// The end date of the range of time applied for this line item's price.
	EndDate time.Time `json:"end_date"`
	// For configured prices that are split by a grouping key, this will be populated with the key and a value. The `amount` and `subtotal` will be the values for this particular grouping.
	Grouping InvoiceLineItemsGrouping `json:"grouping"`
	Minimum  map[string]interface{}   `json:"minimum"`
	// The name of the price associated with this line item.
	Name     string  `json:"name"`
	Quantity float64 `json:"quantity"`
	// The start date of the range of time applied for this line item's price.
	StartDate time.Time `json:"start_date"`
	// For complex pricing structures, the line item can be broken down further in `sub_line_items`.
	SubLineItems []InvoiceLineItemsSubLineItems `json:"sub_line_items"`
	// The line amount before any line item-specific discounts or minimums.
	Subtotal string `json:"subtotal"`
	// An array of tax rates and their incurred tax amounts. Empty if no tax integration is configured.
	TaxAmounts []InvoiceLineItemsTaxAmounts `json:"tax_amounts"`
}

// InvoiceStatusEnum - The status of this invoice as known to Orb. Invoices that have been issued for past billing periods are marked `"issued"`. Invoices will be marked `"paid"` upon confirmation of successful automatic payment collection by Orb. Invoices synced to an external billing provider (such as Bill.com, QuickBooks, or Stripe Invoicing) will be marked as `"synced"`.
type InvoiceStatusEnum string

const (
	InvoiceStatusEnumIssued InvoiceStatusEnum = "issued"
	InvoiceStatusEnumPaid   InvoiceStatusEnum = "paid"
	InvoiceStatusEnumSynced InvoiceStatusEnum = "synced"
)

func (e InvoiceStatusEnum) ToPointer() *InvoiceStatusEnum {
	return &e
}

func (e *InvoiceStatusEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "issued":
		fallthrough
	case "paid":
		fallthrough
	case "synced":
		*e = InvoiceStatusEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for InvoiceStatusEnum: %s", s)
	}
}

// InvoiceSubscription - The associated subscription for this invoice.
type InvoiceSubscription struct {
	ID string `json:"id"`
}

// Invoice - An [`Invoice`](../reference/Orb-API.json/components/schemas/Invoice) is a fundamental billing entity, representing the request for payment for a single subscription. This includes a set of line items, which correspond to prices in the subscription's plan and can represent fixed recurring fees or usage-based fees. They are generated at the end of a billing period, or as the result of an action, such as a cancellation.
type Invoice struct {
	// This is the final amount required to be charged to the customer and reflects the application of the customer balance to the `total` of the invoice.
	AmountDue string `json:"amount_due"`
	// The creation time of the resource in Orb.
	CreatedAt time.Time `json:"created_at"`
	// An ISO 4217 currency string or `credits`
	Currency string `json:"currency"`
	// The customer receiving this invoice.
	Customer InvoiceCustomer `json:"customer"`
	// A list of Customer balance transactions that may be associated with this invoice.
	CustomerBalanceTransactions []CustomerBalanceTransaction `json:"customer_balance_transactions"`
	Discount                    map[string]interface{}       `json:"discount"`
	// When the invoice payment is due.
	DueDate time.Time `json:"due_date"`
	ID      string    `json:"id"`
	// Issue date of the invoice
	InvoiceDate time.Time `json:"invoice_date"`
	// The link to download the PDF representation of the `Invoice`.
	InvoicePdf string `json:"invoice_pdf"`
	// The breakdown of prices in this invoice.
	LineItems []InvoiceLineItems     `json:"line_items"`
	Minimum   map[string]interface{} `json:"minimum"`
	// The status of this invoice as known to Orb. Invoices that have been issued for past billing periods are marked `"issued"`. Invoices will be marked `"paid"` upon confirmation of successful automatic payment collection by Orb. Invoices synced to an external billing provider (such as Bill.com, QuickBooks, or Stripe Invoicing) will be marked as `"synced"`.
	Status *InvoiceStatusEnum `json:"status,omitempty"`
	// The associated subscription for this invoice.
	Subscription InvoiceSubscription `json:"subscription"`
	// The total before any discounts and minimums are applied.
	Subtotal string `json:"subtotal"`
	// The total after any minimums, discounts, and taxes have been applied.
	Total string `json:"total"`
}
