// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Orb/pkg/models/shared"
	"encoding/json"
	"fmt"
	"net/http"
)

// GetCustomersCustomerIDCreditsLedgerEntryStatus - Filters to a single status of ledger entry
type GetCustomersCustomerIDCreditsLedgerEntryStatus string

const (
	GetCustomersCustomerIDCreditsLedgerEntryStatusCommitted GetCustomersCustomerIDCreditsLedgerEntryStatus = "committed"
	GetCustomersCustomerIDCreditsLedgerEntryStatusPending   GetCustomersCustomerIDCreditsLedgerEntryStatus = "pending"
)

func (e GetCustomersCustomerIDCreditsLedgerEntryStatus) ToPointer() *GetCustomersCustomerIDCreditsLedgerEntryStatus {
	return &e
}

func (e *GetCustomersCustomerIDCreditsLedgerEntryStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "committed":
		fallthrough
	case "pending":
		*e = GetCustomersCustomerIDCreditsLedgerEntryStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetCustomersCustomerIDCreditsLedgerEntryStatus: %v", v)
	}
}

// GetCustomersCustomerIDCreditsLedgerEntryType - Filter to a single type of ledger entry
type GetCustomersCustomerIDCreditsLedgerEntryType string

const (
	GetCustomersCustomerIDCreditsLedgerEntryTypeIncrement         GetCustomersCustomerIDCreditsLedgerEntryType = "increment"
	GetCustomersCustomerIDCreditsLedgerEntryTypeDecrement         GetCustomersCustomerIDCreditsLedgerEntryType = "decrement"
	GetCustomersCustomerIDCreditsLedgerEntryTypeExpirationChange  GetCustomersCustomerIDCreditsLedgerEntryType = "expiration_change"
	GetCustomersCustomerIDCreditsLedgerEntryTypeCreditBlockExpiry GetCustomersCustomerIDCreditsLedgerEntryType = "credit_block_expiry"
)

func (e GetCustomersCustomerIDCreditsLedgerEntryType) ToPointer() *GetCustomersCustomerIDCreditsLedgerEntryType {
	return &e
}

func (e *GetCustomersCustomerIDCreditsLedgerEntryType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "increment":
		fallthrough
	case "decrement":
		fallthrough
	case "expiration_change":
		fallthrough
	case "credit_block_expiry":
		*e = GetCustomersCustomerIDCreditsLedgerEntryType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetCustomersCustomerIDCreditsLedgerEntryType: %v", v)
	}
}

type GetCustomersCustomerIDCreditsLedgerRequest struct {
	CustomerID string `pathParam:"style=simple,explode=false,name=customer_id"`
	// Filters to a single status of ledger entry
	EntryStatus *GetCustomersCustomerIDCreditsLedgerEntryStatus `queryParam:"style=form,explode=true,name=entry_status"`
	// Filter to a single type of ledger entry
	EntryType *GetCustomersCustomerIDCreditsLedgerEntryType `queryParam:"style=form,explode=true,name=entry_type"`
	// Filter to ledger entries that affect at least this amount
	MinimumAmount *float64 `queryParam:"style=form,explode=true,name=minimum_amount"`
}

type GetCustomersCustomerIDCreditsLedger200ApplicationJSONPaginationMetadata struct {
	HasMore    bool   `json:"has_more"`
	NextCursor string `json:"next_cursor"`
}

// GetCustomersCustomerIDCreditsLedger200ApplicationJSON - OK
type GetCustomersCustomerIDCreditsLedger200ApplicationJSON struct {
	Data               []shared.CreditLedgerEntry                                              `json:"data"`
	PaginationMetadata GetCustomersCustomerIDCreditsLedger200ApplicationJSONPaginationMetadata `json:"pagination_metadata"`
}

type GetCustomersCustomerIDCreditsLedgerResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	GetCustomersCustomerIDCreditsLedger200ApplicationJSONObject *GetCustomersCustomerIDCreditsLedger200ApplicationJSON
}
