// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"time"
)

// CreditLedgerEntryCreditBlock - Credit block that the entry affected
type CreditLedgerEntryCreditBlock struct {
	// Complete timestamp with date and time for when this block expires
	ExpiryDate string `json:"expiry_date"`
	ID         string `json:"id"`
	// How much, in USD, a customer paid for a single credit in this block
	PerUnitCostBasis string `json:"per_unit_cost_basis"`
}

type CreditLedgerEntryCustomer struct {
	ExternalCustomerID string `json:"external_customer_id"`
	ID                 string `json:"id"`
}

// CreditLedgerEntryEntryStatus - Committed entries are older than the ingestion grace period, and cannot change. Pending entries are newer than the grace period and are subject to updates
type CreditLedgerEntryEntryStatus string

const (
	CreditLedgerEntryEntryStatusCommitted CreditLedgerEntryEntryStatus = "committed"
	CreditLedgerEntryEntryStatusPending   CreditLedgerEntryEntryStatus = "pending"
)

func (e CreditLedgerEntryEntryStatus) ToPointer() *CreditLedgerEntryEntryStatus {
	return &e
}

func (e *CreditLedgerEntryEntryStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "committed":
		fallthrough
	case "pending":
		*e = CreditLedgerEntryEntryStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreditLedgerEntryEntryStatus: %v", v)
	}
}

type CreditLedgerEntryEntryType string

const (
	CreditLedgerEntryEntryTypeIncrement         CreditLedgerEntryEntryType = "increment"
	CreditLedgerEntryEntryTypeDecrement         CreditLedgerEntryEntryType = "decrement"
	CreditLedgerEntryEntryTypeExpirationChange  CreditLedgerEntryEntryType = "expiration_change"
	CreditLedgerEntryEntryTypeCreditBlockExpiry CreditLedgerEntryEntryType = "credit_block_expiry"
)

func (e CreditLedgerEntryEntryType) ToPointer() *CreditLedgerEntryEntryType {
	return &e
}

func (e *CreditLedgerEntryEntryType) UnmarshalJSON(data []byte) error {
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
		*e = CreditLedgerEntryEntryType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreditLedgerEntryEntryType: %v", v)
	}
}

// CreditLedgerEntryMetadata - User-specified metadata dictionary that's specified when adding a ledger entry. This contains key/value pairs if metadata is specified, but otherwise is an empty dictionary.
type CreditLedgerEntryMetadata struct {
}

// CreditLedgerEntry - A credit ledger entry is a single entry in the customer balance ledger. More details about working with real-time balances are [here](../guides/product-catalog/prepurchase).
//
// To support late and out-of-order event reporting, ledger entries are marked as either __committed_ or _pending_. Committed entries are finalized and will not change. Pending entries can be updated up until the event reporting grace period.
type CreditLedgerEntry struct {
	// Number of credits that were impacted. Required on creation for increment and decrement entries.
	Amount    *float64  `json:"amount,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	// Credit block that the entry affected
	CreditBlock CreditLedgerEntryCreditBlock `json:"credit_block"`
	Customer    CreditLedgerEntryCustomer    `json:"customer"`
	// Optional metadata that can be specified when adding ledger results via the API
	Description   string  `json:"description"`
	EndingBalance float64 `json:"ending_balance"`
	// Committed entries are older than the ingestion grace period, and cannot change. Pending entries are newer than the grace period and are subject to updates
	EntryStatus CreditLedgerEntryEntryStatus `json:"entry_status"`
	EntryType   CreditLedgerEntryEntryType   `json:"entry_type"`
	EventID     *string                      `json:"event_id,omitempty"`
	ID          string                       `json:"id"`
	// The position in which this entry appears in the ledger, starting at 0
	LedgerSequenceNumber float64 `json:"ledger_sequence_number"`
	// User-specified metadata dictionary that's specified when adding a ledger entry. This contains key/value pairs if metadata is specified, but otherwise is an empty dictionary.
	Metadata CreditLedgerEntryMetadata `json:"metadata"`
	// In the case of an expiration change ledger entry, this represents the expiration time of the new block.
	NewBlockExpiryDate *string `json:"new_block_expiry_date,omitempty"`
	PriceID            *string `json:"price_id,omitempty"`
	StartingBalance    float64 `json:"starting_balance"`
}
