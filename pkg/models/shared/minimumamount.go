// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type MinimumAmount struct {
	// List of price_ids that this minimum amount applies to. For plan/plan phase minimums, this can be a subset of prices.
	AppliesToPriceIds []string `json:"applies_to_price_ids"`
	// Minimum amount applied
	MinimumAmount string `json:"minimum_amount"`
}