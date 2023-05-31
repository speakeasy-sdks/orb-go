// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type DiscountDiscountType string

const (
	DiscountDiscountTypePercentage DiscountDiscountType = "percentage"
	DiscountDiscountTypeTrial      DiscountDiscountType = "trial"
	DiscountDiscountTypeUsage      DiscountDiscountType = "usage"
	DiscountDiscountTypeAmount     DiscountDiscountType = "amount"
)

func (e DiscountDiscountType) ToPointer() *DiscountDiscountType {
	return &e
}

func (e *DiscountDiscountType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "percentage":
		fallthrough
	case "trial":
		fallthrough
	case "usage":
		fallthrough
	case "amount":
		*e = DiscountDiscountType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DiscountDiscountType: %v", v)
	}
}

type Discount struct {
	// Only available if discount_type is `amount`.
	AmountDiscount *string `json:"amount_discount,omitempty"`
	// List of price_ids that this discount applies to. For plan/plan phase discounts, this can be a subset of prices.
	AppliesToPriceIds []string             `json:"applies_to_price_ids,omitempty"`
	DiscountType      DiscountDiscountType `json:"discount_type"`
	// Only available if discount_type is `percentage`.This is a number between 0 and 1.
	PercentageDiscount *float64 `json:"percentage_discount,omitempty"`
	// Only available if discount_type is `trial`
	TrialAmountDiscount *string `json:"trial_amount_discount,omitempty"`
	// Only available if discount_type is `usage`. Number of usage units that this discount is for
	UsageDiscount *float64 `json:"usage_discount,omitempty"`
}
