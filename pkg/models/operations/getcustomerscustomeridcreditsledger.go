// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Orb/pkg/models/shared"
	"encoding/json"
	"fmt"
	"net/http"
)

// GetCustomersCustomerIDCreditsLedgerEntryStatusEnum - Filters to a single status of ledger entry
type GetCustomersCustomerIDCreditsLedgerEntryStatusEnum string

const (
	GetCustomersCustomerIDCreditsLedgerEntryStatusEnumCommitted GetCustomersCustomerIDCreditsLedgerEntryStatusEnum = "committed"
	GetCustomersCustomerIDCreditsLedgerEntryStatusEnumPending   GetCustomersCustomerIDCreditsLedgerEntryStatusEnum = "pending"
)

func (e *GetCustomersCustomerIDCreditsLedgerEntryStatusEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "committed":
		fallthrough
	case "pending":
		*e = GetCustomersCustomerIDCreditsLedgerEntryStatusEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for GetCustomersCustomerIDCreditsLedgerEntryStatusEnum: %s", s)
	}
}

// GetCustomersCustomerIDCreditsLedgerEntryTypeEnum - Filter to a single type of ledger entry
type GetCustomersCustomerIDCreditsLedgerEntryTypeEnum string

const (
	GetCustomersCustomerIDCreditsLedgerEntryTypeEnumIncrement         GetCustomersCustomerIDCreditsLedgerEntryTypeEnum = "increment"
	GetCustomersCustomerIDCreditsLedgerEntryTypeEnumDecrement         GetCustomersCustomerIDCreditsLedgerEntryTypeEnum = "decrement"
	GetCustomersCustomerIDCreditsLedgerEntryTypeEnumExpirationChange  GetCustomersCustomerIDCreditsLedgerEntryTypeEnum = "expiration_change"
	GetCustomersCustomerIDCreditsLedgerEntryTypeEnumCreditBlockExpiry GetCustomersCustomerIDCreditsLedgerEntryTypeEnum = "credit_block_expiry"
)

func (e *GetCustomersCustomerIDCreditsLedgerEntryTypeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "increment":
		fallthrough
	case "decrement":
		fallthrough
	case "expiration_change":
		fallthrough
	case "credit_block_expiry":
		*e = GetCustomersCustomerIDCreditsLedgerEntryTypeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for GetCustomersCustomerIDCreditsLedgerEntryTypeEnum: %s", s)
	}
}

type GetCustomersCustomerIDCreditsLedgerRequest struct {
	CustomerID string `pathParam:"style=simple,explode=false,name=customer_id"`
	// Filters to a single status of ledger entry
	EntryStatus *GetCustomersCustomerIDCreditsLedgerEntryStatusEnum `queryParam:"style=form,explode=true,name=entry_status"`
	// Filter to a single type of ledger entry
	EntryType *GetCustomersCustomerIDCreditsLedgerEntryTypeEnum `queryParam:"style=form,explode=true,name=entry_type"`
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
