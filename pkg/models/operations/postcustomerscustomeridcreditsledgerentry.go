// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Orb/pkg/models/shared"
	"Orb/pkg/types"
	"encoding/json"
	"fmt"
	"net/http"
)

type PostCustomersCustomerIDCreditsLedgerEntryRequestBodyEntryTypeEnum string

const (
	PostCustomersCustomerIDCreditsLedgerEntryRequestBodyEntryTypeEnumIncrement        PostCustomersCustomerIDCreditsLedgerEntryRequestBodyEntryTypeEnum = "increment"
	PostCustomersCustomerIDCreditsLedgerEntryRequestBodyEntryTypeEnumDecrement        PostCustomersCustomerIDCreditsLedgerEntryRequestBodyEntryTypeEnum = "decrement"
	PostCustomersCustomerIDCreditsLedgerEntryRequestBodyEntryTypeEnumExpirationChange PostCustomersCustomerIDCreditsLedgerEntryRequestBodyEntryTypeEnum = "expiration_change"
)

func (e *PostCustomersCustomerIDCreditsLedgerEntryRequestBodyEntryTypeEnum) UnmarshalJSON(data []byte) error {
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
		*e = PostCustomersCustomerIDCreditsLedgerEntryRequestBodyEntryTypeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for PostCustomersCustomerIDCreditsLedgerEntryRequestBodyEntryTypeEnum: %s", s)
	}
}

type PostCustomersCustomerIDCreditsLedgerEntryRequestBody struct {
	Amount float64 `json:"amount"`
	// The ID of the block affected by an `expiration_change`
	BlockID *string `json:"block_id,omitempty"`
	// Optional metadata that can be specified when adding ledger results via the API. For example, this can be used to note an increment refers to trial credits, or for noting corrections as a result of an incident, etc.
	Description *string                                                           `json:"description,omitempty"`
	EntryType   PostCustomersCustomerIDCreditsLedgerEntryRequestBodyEntryTypeEnum `json:"entry_type"`
	// A future date (specified in YYYY-MM-DD format) that denotes when this credit balance should expire.
	//
	//
	ExpiryDate *types.Date `json:"expiry_date,omitempty"`
	// Can only be specified when `entry_type=increment`. How much, in USD, a customer paid for a single credit in this block
	PerUnitCostBasis *string `json:"per_unit_cost_basis,omitempty"`
	// A future date (specified in YYYY-MM-DD) used for `expiration_change`
	TargetExpiryDate *types.Date `json:"target_expiry_date,omitempty"`
}

type PostCustomersCustomerIDCreditsLedgerEntryRequest struct {
	RequestBody *PostCustomersCustomerIDCreditsLedgerEntryRequestBody `request:"mediaType=application/json"`
	CustomerID  string                                                `pathParam:"style=simple,explode=false,name=customer_id"`
}

type PostCustomersCustomerIDCreditsLedgerEntryResponse struct {
	ContentType string
	// OK
	CreditLedgerEntry *shared.CreditLedgerEntry
	StatusCode        int
	RawResponse       *http.Response
}
