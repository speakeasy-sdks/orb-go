// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Orb/pkg/models/shared"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

// SchedulePlanChangeRequestBodyChangeOption - Determines the timing of the plan change
type SchedulePlanChangeRequestBodyChangeOption string

const (
	SchedulePlanChangeRequestBodyChangeOptionRequestedDate         SchedulePlanChangeRequestBodyChangeOption = "requested_date"
	SchedulePlanChangeRequestBodyChangeOptionEndOfSubscriptionTerm SchedulePlanChangeRequestBodyChangeOption = "end_of_subscription_term"
	SchedulePlanChangeRequestBodyChangeOptionImmediate             SchedulePlanChangeRequestBodyChangeOption = "immediate"
)

func (e SchedulePlanChangeRequestBodyChangeOption) ToPointer() *SchedulePlanChangeRequestBodyChangeOption {
	return &e
}

func (e *SchedulePlanChangeRequestBodyChangeOption) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "requested_date":
		fallthrough
	case "end_of_subscription_term":
		fallthrough
	case "immediate":
		*e = SchedulePlanChangeRequestBodyChangeOption(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SchedulePlanChangeRequestBodyChangeOption: %v", v)
	}
}

type SchedulePlanChangeRequestBodyPriceOverrides7ModelType string

const (
	SchedulePlanChangeRequestBodyPriceOverrides7ModelTypeTieredBps SchedulePlanChangeRequestBodyPriceOverrides7ModelType = "tiered_bps"
)

func (e SchedulePlanChangeRequestBodyPriceOverrides7ModelType) ToPointer() *SchedulePlanChangeRequestBodyPriceOverrides7ModelType {
	return &e
}

func (e *SchedulePlanChangeRequestBodyPriceOverrides7ModelType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "tiered_bps":
		*e = SchedulePlanChangeRequestBodyPriceOverrides7ModelType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SchedulePlanChangeRequestBodyPriceOverrides7ModelType: %v", v)
	}
}

type SchedulePlanChangeRequestBodyPriceOverrides7TieredBpsConfigTiers struct {
	Bps            float64 `json:"bps"`
	MaximumAmount  string  `json:"maximum_amount"`
	MinimumAmount  string  `json:"minimum_amount"`
	PerUnitMaximum string  `json:"per_unit_maximum"`
}

type SchedulePlanChangeRequestBodyPriceOverrides7TieredBpsConfig struct {
	Tiers []SchedulePlanChangeRequestBodyPriceOverrides7TieredBpsConfigTiers `json:"tiers"`
}

// SchedulePlanChangeRequestBodyPriceOverrides7 - Tiered BPS price override
type SchedulePlanChangeRequestBodyPriceOverrides7 struct {
	ID string `json:"id"`
	// The subscription's override minimum amount for this price.
	MinimumAmount   *string                                                     `json:"minimum_amount,omitempty"`
	ModelType       SchedulePlanChangeRequestBodyPriceOverrides7ModelType       `json:"model_type"`
	TieredBpsConfig SchedulePlanChangeRequestBodyPriceOverrides7TieredBpsConfig `json:"tiered_bps_config"`
}

type SchedulePlanChangeRequestBodyPriceOverrides6BulkBpsConfigTiers struct {
	Bps            float64 `json:"bps"`
	MaximumAmount  string  `json:"maximum_amount"`
	PerUnitMaximum string  `json:"per_unit_maximum"`
}

type SchedulePlanChangeRequestBodyPriceOverrides6BulkBpsConfig struct {
	Tiers []SchedulePlanChangeRequestBodyPriceOverrides6BulkBpsConfigTiers `json:"tiers,omitempty"`
}

type SchedulePlanChangeRequestBodyPriceOverrides6ModelType string

const (
	SchedulePlanChangeRequestBodyPriceOverrides6ModelTypeBulkBps SchedulePlanChangeRequestBodyPriceOverrides6ModelType = "bulk_bps"
)

func (e SchedulePlanChangeRequestBodyPriceOverrides6ModelType) ToPointer() *SchedulePlanChangeRequestBodyPriceOverrides6ModelType {
	return &e
}

func (e *SchedulePlanChangeRequestBodyPriceOverrides6ModelType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "bulk_bps":
		*e = SchedulePlanChangeRequestBodyPriceOverrides6ModelType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SchedulePlanChangeRequestBodyPriceOverrides6ModelType: %v", v)
	}
}

// SchedulePlanChangeRequestBodyPriceOverrides6 - Bulk BPS price override
type SchedulePlanChangeRequestBodyPriceOverrides6 struct {
	BulkBpsConfig *SchedulePlanChangeRequestBodyPriceOverrides6BulkBpsConfig `json:"bulk_bps_config,omitempty"`
	ID            string                                                     `json:"id"`
	// The subscription's override minimum amount for this price.
	MinimumAmount *string                                               `json:"minimum_amount,omitempty"`
	ModelType     SchedulePlanChangeRequestBodyPriceOverrides6ModelType `json:"model_type"`
}

type SchedulePlanChangeRequestBodyPriceOverrides5BpsConfig struct {
	Bps            float64 `json:"bps"`
	PerUnitMaximum string  `json:"per_unit_maximum"`
}

type SchedulePlanChangeRequestBodyPriceOverrides5ModelType string

const (
	SchedulePlanChangeRequestBodyPriceOverrides5ModelTypeBps SchedulePlanChangeRequestBodyPriceOverrides5ModelType = "bps"
)

func (e SchedulePlanChangeRequestBodyPriceOverrides5ModelType) ToPointer() *SchedulePlanChangeRequestBodyPriceOverrides5ModelType {
	return &e
}

func (e *SchedulePlanChangeRequestBodyPriceOverrides5ModelType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "bps":
		*e = SchedulePlanChangeRequestBodyPriceOverrides5ModelType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SchedulePlanChangeRequestBodyPriceOverrides5ModelType: %v", v)
	}
}

// SchedulePlanChangeRequestBodyPriceOverrides5 - BPS price override
type SchedulePlanChangeRequestBodyPriceOverrides5 struct {
	BpsConfig SchedulePlanChangeRequestBodyPriceOverrides5BpsConfig `json:"bps_config"`
	ID        string                                                `json:"id"`
	// The subscription's override minimum amount for this price.
	MinimumAmount *string                                               `json:"minimum_amount,omitempty"`
	ModelType     SchedulePlanChangeRequestBodyPriceOverrides5ModelType `json:"model_type"`
}

type SchedulePlanChangeRequestBodyPriceOverrides4ModelType string

const (
	SchedulePlanChangeRequestBodyPriceOverrides4ModelTypePackage SchedulePlanChangeRequestBodyPriceOverrides4ModelType = "package"
)

func (e SchedulePlanChangeRequestBodyPriceOverrides4ModelType) ToPointer() *SchedulePlanChangeRequestBodyPriceOverrides4ModelType {
	return &e
}

func (e *SchedulePlanChangeRequestBodyPriceOverrides4ModelType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "package":
		*e = SchedulePlanChangeRequestBodyPriceOverrides4ModelType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SchedulePlanChangeRequestBodyPriceOverrides4ModelType: %v", v)
	}
}

type SchedulePlanChangeRequestBodyPriceOverrides4PackageConfig struct {
	NumUnits   float64 `json:"num_units"`
	UnitAmount string  `json:"unit_amount"`
}

// SchedulePlanChangeRequestBodyPriceOverrides4 - Package price override
type SchedulePlanChangeRequestBodyPriceOverrides4 struct {
	ID string `json:"id"`
	// The subscription's override minimum amount for this price.
	MinimumAmount *string                                                   `json:"minimum_amount,omitempty"`
	ModelType     SchedulePlanChangeRequestBodyPriceOverrides4ModelType     `json:"model_type"`
	PackageConfig SchedulePlanChangeRequestBodyPriceOverrides4PackageConfig `json:"package_config"`
}

type SchedulePlanChangeRequestBodyPriceOverrides3BulkConfigTiers struct {
	MaximumUnits string `json:"maximum_units"`
	UnitAmount   string `json:"unit_amount"`
}

type SchedulePlanChangeRequestBodyPriceOverrides3BulkConfig struct {
	Tiers []SchedulePlanChangeRequestBodyPriceOverrides3BulkConfigTiers `json:"tiers"`
}

type SchedulePlanChangeRequestBodyPriceOverrides3ModelType string

const (
	SchedulePlanChangeRequestBodyPriceOverrides3ModelTypeBulk SchedulePlanChangeRequestBodyPriceOverrides3ModelType = "bulk"
)

func (e SchedulePlanChangeRequestBodyPriceOverrides3ModelType) ToPointer() *SchedulePlanChangeRequestBodyPriceOverrides3ModelType {
	return &e
}

func (e *SchedulePlanChangeRequestBodyPriceOverrides3ModelType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "bulk":
		*e = SchedulePlanChangeRequestBodyPriceOverrides3ModelType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SchedulePlanChangeRequestBodyPriceOverrides3ModelType: %v", v)
	}
}

// SchedulePlanChangeRequestBodyPriceOverrides3 - Bulk price override
type SchedulePlanChangeRequestBodyPriceOverrides3 struct {
	BulkConfig SchedulePlanChangeRequestBodyPriceOverrides3BulkConfig `json:"bulk_config"`
	ID         string                                                 `json:"id"`
	// The subscription's override minimum amount for this price.
	MinimumAmount *string                                               `json:"minimum_amount,omitempty"`
	ModelType     SchedulePlanChangeRequestBodyPriceOverrides3ModelType `json:"model_type"`
}

type SchedulePlanChangeRequestBodyPriceOverrides2ModelType string

const (
	SchedulePlanChangeRequestBodyPriceOverrides2ModelTypeUnit SchedulePlanChangeRequestBodyPriceOverrides2ModelType = "unit"
)

func (e SchedulePlanChangeRequestBodyPriceOverrides2ModelType) ToPointer() *SchedulePlanChangeRequestBodyPriceOverrides2ModelType {
	return &e
}

func (e *SchedulePlanChangeRequestBodyPriceOverrides2ModelType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "unit":
		*e = SchedulePlanChangeRequestBodyPriceOverrides2ModelType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SchedulePlanChangeRequestBodyPriceOverrides2ModelType: %v", v)
	}
}

type SchedulePlanChangeRequestBodyPriceOverrides2UnitConfig struct {
	UnitAmount string `json:"unit_amount"`
}

// SchedulePlanChangeRequestBodyPriceOverrides2 - Unit price override
type SchedulePlanChangeRequestBodyPriceOverrides2 struct {
	ID string `json:"id"`
	// The subscription's override minimum amount for this price.
	MinimumAmount *string                                                `json:"minimum_amount,omitempty"`
	ModelType     SchedulePlanChangeRequestBodyPriceOverrides2ModelType  `json:"model_type"`
	UnitConfig    SchedulePlanChangeRequestBodyPriceOverrides2UnitConfig `json:"unit_config"`
}

type SchedulePlanChangeRequestBodyPriceOverrides1ModelType string

const (
	SchedulePlanChangeRequestBodyPriceOverrides1ModelTypeTiered SchedulePlanChangeRequestBodyPriceOverrides1ModelType = "tiered"
)

func (e SchedulePlanChangeRequestBodyPriceOverrides1ModelType) ToPointer() *SchedulePlanChangeRequestBodyPriceOverrides1ModelType {
	return &e
}

func (e *SchedulePlanChangeRequestBodyPriceOverrides1ModelType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "tiered":
		*e = SchedulePlanChangeRequestBodyPriceOverrides1ModelType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SchedulePlanChangeRequestBodyPriceOverrides1ModelType: %v", v)
	}
}

type SchedulePlanChangeRequestBodyPriceOverrides1TieredConfigTiers struct {
	FirstUnit  *string `json:"first_unit,omitempty"`
	LastUnit   *string `json:"last_unit,omitempty"`
	UnitAmount *string `json:"unit_amount,omitempty"`
}

type SchedulePlanChangeRequestBodyPriceOverrides1TieredConfig struct {
	Tiers []SchedulePlanChangeRequestBodyPriceOverrides1TieredConfigTiers `json:"tiers,omitempty"`
}

// SchedulePlanChangeRequestBodyPriceOverrides1 - Tiered price override
type SchedulePlanChangeRequestBodyPriceOverrides1 struct {
	// price_id
	ID string `json:"id"`
	// The subscription's override minimum amount for this price.
	MinimumAmount *string                                                  `json:"minimum_amount,omitempty"`
	ModelType     SchedulePlanChangeRequestBodyPriceOverrides1ModelType    `json:"model_type"`
	TieredConfig  SchedulePlanChangeRequestBodyPriceOverrides1TieredConfig `json:"tiered_config"`
}

type SchedulePlanChangeRequestBodyPriceOverridesType string

const (
	SchedulePlanChangeRequestBodyPriceOverridesTypeSchedulePlanChangeRequestBodyPriceOverrides1 SchedulePlanChangeRequestBodyPriceOverridesType = "schedule-plan-change_requestBody_price_overrides_1"
	SchedulePlanChangeRequestBodyPriceOverridesTypeSchedulePlanChangeRequestBodyPriceOverrides2 SchedulePlanChangeRequestBodyPriceOverridesType = "schedule-plan-change_requestBody_price_overrides_2"
	SchedulePlanChangeRequestBodyPriceOverridesTypeSchedulePlanChangeRequestBodyPriceOverrides3 SchedulePlanChangeRequestBodyPriceOverridesType = "schedule-plan-change_requestBody_price_overrides_3"
	SchedulePlanChangeRequestBodyPriceOverridesTypeSchedulePlanChangeRequestBodyPriceOverrides4 SchedulePlanChangeRequestBodyPriceOverridesType = "schedule-plan-change_requestBody_price_overrides_4"
	SchedulePlanChangeRequestBodyPriceOverridesTypeSchedulePlanChangeRequestBodyPriceOverrides5 SchedulePlanChangeRequestBodyPriceOverridesType = "schedule-plan-change_requestBody_price_overrides_5"
	SchedulePlanChangeRequestBodyPriceOverridesTypeSchedulePlanChangeRequestBodyPriceOverrides6 SchedulePlanChangeRequestBodyPriceOverridesType = "schedule-plan-change_requestBody_price_overrides_6"
	SchedulePlanChangeRequestBodyPriceOverridesTypeSchedulePlanChangeRequestBodyPriceOverrides7 SchedulePlanChangeRequestBodyPriceOverridesType = "schedule-plan-change_requestBody_price_overrides_7"
)

type SchedulePlanChangeRequestBodyPriceOverrides struct {
	SchedulePlanChangeRequestBodyPriceOverrides1 *SchedulePlanChangeRequestBodyPriceOverrides1
	SchedulePlanChangeRequestBodyPriceOverrides2 *SchedulePlanChangeRequestBodyPriceOverrides2
	SchedulePlanChangeRequestBodyPriceOverrides3 *SchedulePlanChangeRequestBodyPriceOverrides3
	SchedulePlanChangeRequestBodyPriceOverrides4 *SchedulePlanChangeRequestBodyPriceOverrides4
	SchedulePlanChangeRequestBodyPriceOverrides5 *SchedulePlanChangeRequestBodyPriceOverrides5
	SchedulePlanChangeRequestBodyPriceOverrides6 *SchedulePlanChangeRequestBodyPriceOverrides6
	SchedulePlanChangeRequestBodyPriceOverrides7 *SchedulePlanChangeRequestBodyPriceOverrides7

	Type SchedulePlanChangeRequestBodyPriceOverridesType
}

func CreateSchedulePlanChangeRequestBodyPriceOverridesSchedulePlanChangeRequestBodyPriceOverrides1(schedulePlanChangeRequestBodyPriceOverrides1 SchedulePlanChangeRequestBodyPriceOverrides1) SchedulePlanChangeRequestBodyPriceOverrides {
	typ := SchedulePlanChangeRequestBodyPriceOverridesTypeSchedulePlanChangeRequestBodyPriceOverrides1

	return SchedulePlanChangeRequestBodyPriceOverrides{
		SchedulePlanChangeRequestBodyPriceOverrides1: &schedulePlanChangeRequestBodyPriceOverrides1,
		Type: typ,
	}
}

func CreateSchedulePlanChangeRequestBodyPriceOverridesSchedulePlanChangeRequestBodyPriceOverrides2(schedulePlanChangeRequestBodyPriceOverrides2 SchedulePlanChangeRequestBodyPriceOverrides2) SchedulePlanChangeRequestBodyPriceOverrides {
	typ := SchedulePlanChangeRequestBodyPriceOverridesTypeSchedulePlanChangeRequestBodyPriceOverrides2

	return SchedulePlanChangeRequestBodyPriceOverrides{
		SchedulePlanChangeRequestBodyPriceOverrides2: &schedulePlanChangeRequestBodyPriceOverrides2,
		Type: typ,
	}
}

func CreateSchedulePlanChangeRequestBodyPriceOverridesSchedulePlanChangeRequestBodyPriceOverrides3(schedulePlanChangeRequestBodyPriceOverrides3 SchedulePlanChangeRequestBodyPriceOverrides3) SchedulePlanChangeRequestBodyPriceOverrides {
	typ := SchedulePlanChangeRequestBodyPriceOverridesTypeSchedulePlanChangeRequestBodyPriceOverrides3

	return SchedulePlanChangeRequestBodyPriceOverrides{
		SchedulePlanChangeRequestBodyPriceOverrides3: &schedulePlanChangeRequestBodyPriceOverrides3,
		Type: typ,
	}
}

func CreateSchedulePlanChangeRequestBodyPriceOverridesSchedulePlanChangeRequestBodyPriceOverrides4(schedulePlanChangeRequestBodyPriceOverrides4 SchedulePlanChangeRequestBodyPriceOverrides4) SchedulePlanChangeRequestBodyPriceOverrides {
	typ := SchedulePlanChangeRequestBodyPriceOverridesTypeSchedulePlanChangeRequestBodyPriceOverrides4

	return SchedulePlanChangeRequestBodyPriceOverrides{
		SchedulePlanChangeRequestBodyPriceOverrides4: &schedulePlanChangeRequestBodyPriceOverrides4,
		Type: typ,
	}
}

func CreateSchedulePlanChangeRequestBodyPriceOverridesSchedulePlanChangeRequestBodyPriceOverrides5(schedulePlanChangeRequestBodyPriceOverrides5 SchedulePlanChangeRequestBodyPriceOverrides5) SchedulePlanChangeRequestBodyPriceOverrides {
	typ := SchedulePlanChangeRequestBodyPriceOverridesTypeSchedulePlanChangeRequestBodyPriceOverrides5

	return SchedulePlanChangeRequestBodyPriceOverrides{
		SchedulePlanChangeRequestBodyPriceOverrides5: &schedulePlanChangeRequestBodyPriceOverrides5,
		Type: typ,
	}
}

func CreateSchedulePlanChangeRequestBodyPriceOverridesSchedulePlanChangeRequestBodyPriceOverrides6(schedulePlanChangeRequestBodyPriceOverrides6 SchedulePlanChangeRequestBodyPriceOverrides6) SchedulePlanChangeRequestBodyPriceOverrides {
	typ := SchedulePlanChangeRequestBodyPriceOverridesTypeSchedulePlanChangeRequestBodyPriceOverrides6

	return SchedulePlanChangeRequestBodyPriceOverrides{
		SchedulePlanChangeRequestBodyPriceOverrides6: &schedulePlanChangeRequestBodyPriceOverrides6,
		Type: typ,
	}
}

func CreateSchedulePlanChangeRequestBodyPriceOverridesSchedulePlanChangeRequestBodyPriceOverrides7(schedulePlanChangeRequestBodyPriceOverrides7 SchedulePlanChangeRequestBodyPriceOverrides7) SchedulePlanChangeRequestBodyPriceOverrides {
	typ := SchedulePlanChangeRequestBodyPriceOverridesTypeSchedulePlanChangeRequestBodyPriceOverrides7

	return SchedulePlanChangeRequestBodyPriceOverrides{
		SchedulePlanChangeRequestBodyPriceOverrides7: &schedulePlanChangeRequestBodyPriceOverrides7,
		Type: typ,
	}
}

func (u *SchedulePlanChangeRequestBodyPriceOverrides) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	schedulePlanChangeRequestBodyPriceOverrides1 := new(SchedulePlanChangeRequestBodyPriceOverrides1)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&schedulePlanChangeRequestBodyPriceOverrides1); err == nil {
		u.SchedulePlanChangeRequestBodyPriceOverrides1 = schedulePlanChangeRequestBodyPriceOverrides1
		u.Type = SchedulePlanChangeRequestBodyPriceOverridesTypeSchedulePlanChangeRequestBodyPriceOverrides1
		return nil
	}

	schedulePlanChangeRequestBodyPriceOverrides2 := new(SchedulePlanChangeRequestBodyPriceOverrides2)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&schedulePlanChangeRequestBodyPriceOverrides2); err == nil {
		u.SchedulePlanChangeRequestBodyPriceOverrides2 = schedulePlanChangeRequestBodyPriceOverrides2
		u.Type = SchedulePlanChangeRequestBodyPriceOverridesTypeSchedulePlanChangeRequestBodyPriceOverrides2
		return nil
	}

	schedulePlanChangeRequestBodyPriceOverrides3 := new(SchedulePlanChangeRequestBodyPriceOverrides3)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&schedulePlanChangeRequestBodyPriceOverrides3); err == nil {
		u.SchedulePlanChangeRequestBodyPriceOverrides3 = schedulePlanChangeRequestBodyPriceOverrides3
		u.Type = SchedulePlanChangeRequestBodyPriceOverridesTypeSchedulePlanChangeRequestBodyPriceOverrides3
		return nil
	}

	schedulePlanChangeRequestBodyPriceOverrides4 := new(SchedulePlanChangeRequestBodyPriceOverrides4)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&schedulePlanChangeRequestBodyPriceOverrides4); err == nil {
		u.SchedulePlanChangeRequestBodyPriceOverrides4 = schedulePlanChangeRequestBodyPriceOverrides4
		u.Type = SchedulePlanChangeRequestBodyPriceOverridesTypeSchedulePlanChangeRequestBodyPriceOverrides4
		return nil
	}

	schedulePlanChangeRequestBodyPriceOverrides5 := new(SchedulePlanChangeRequestBodyPriceOverrides5)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&schedulePlanChangeRequestBodyPriceOverrides5); err == nil {
		u.SchedulePlanChangeRequestBodyPriceOverrides5 = schedulePlanChangeRequestBodyPriceOverrides5
		u.Type = SchedulePlanChangeRequestBodyPriceOverridesTypeSchedulePlanChangeRequestBodyPriceOverrides5
		return nil
	}

	schedulePlanChangeRequestBodyPriceOverrides6 := new(SchedulePlanChangeRequestBodyPriceOverrides6)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&schedulePlanChangeRequestBodyPriceOverrides6); err == nil {
		u.SchedulePlanChangeRequestBodyPriceOverrides6 = schedulePlanChangeRequestBodyPriceOverrides6
		u.Type = SchedulePlanChangeRequestBodyPriceOverridesTypeSchedulePlanChangeRequestBodyPriceOverrides6
		return nil
	}

	schedulePlanChangeRequestBodyPriceOverrides7 := new(SchedulePlanChangeRequestBodyPriceOverrides7)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&schedulePlanChangeRequestBodyPriceOverrides7); err == nil {
		u.SchedulePlanChangeRequestBodyPriceOverrides7 = schedulePlanChangeRequestBodyPriceOverrides7
		u.Type = SchedulePlanChangeRequestBodyPriceOverridesTypeSchedulePlanChangeRequestBodyPriceOverrides7
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SchedulePlanChangeRequestBodyPriceOverrides) MarshalJSON() ([]byte, error) {
	if u.SchedulePlanChangeRequestBodyPriceOverrides1 != nil {
		return json.Marshal(u.SchedulePlanChangeRequestBodyPriceOverrides1)
	}

	if u.SchedulePlanChangeRequestBodyPriceOverrides2 != nil {
		return json.Marshal(u.SchedulePlanChangeRequestBodyPriceOverrides2)
	}

	if u.SchedulePlanChangeRequestBodyPriceOverrides3 != nil {
		return json.Marshal(u.SchedulePlanChangeRequestBodyPriceOverrides3)
	}

	if u.SchedulePlanChangeRequestBodyPriceOverrides4 != nil {
		return json.Marshal(u.SchedulePlanChangeRequestBodyPriceOverrides4)
	}

	if u.SchedulePlanChangeRequestBodyPriceOverrides5 != nil {
		return json.Marshal(u.SchedulePlanChangeRequestBodyPriceOverrides5)
	}

	if u.SchedulePlanChangeRequestBodyPriceOverrides6 != nil {
		return json.Marshal(u.SchedulePlanChangeRequestBodyPriceOverrides6)
	}

	if u.SchedulePlanChangeRequestBodyPriceOverrides7 != nil {
		return json.Marshal(u.SchedulePlanChangeRequestBodyPriceOverrides7)
	}

	return nil, nil
}

type SchedulePlanChangeRequestBody struct {
	// Reset billing periods to be aligned with the plan change’s effective date.
	AlignBillingWithPlanChangeDate *bool `json:"align_billing_with_plan_change_date,omitempty"`
	// The date that the plan change should take effect. This parameter can only be passed if the `change_option` is `requested_date`.
	ChangeDate *time.Time `json:"change_date,omitempty"`
	// Determines the timing of the plan change
	ChangeOption SchedulePlanChangeRequestBodyChangeOption `json:"change_option"`
	// Redemption code to be used for this subscription. If the coupon cannot be found by its redemption code, or cannot be redeemed, an error response will be returned and the plan change will not be scheduled.
	CouponRedemptionCode *string `json:"coupon_redemption_code,omitempty"`
	// The external_plan_id of the plan that the given subscription should be switched to. Note that either this property or `plan_id` must be specified.
	ExternalPlanID *string `json:"external_plan_id,omitempty"`
	// The subscription's override minimum amount for the plan.
	MinimumAmount *string `json:"minimum_amount,omitempty"`
	// The plan that the given subscription should be switched to. Note that either this property or `external_plan_id` must be specified.
	PlanID *string `json:"plan_id,omitempty"`
	// Optionally provide a list of overrides for prices on the plan
	PriceOverrides []SchedulePlanChangeRequestBodyPriceOverrides `json:"price_overrides,omitempty"`
}

type SchedulePlanChangeRequest struct {
	RequestBody    *SchedulePlanChangeRequestBody `request:"mediaType=application/json"`
	SubscriptionID string                         `pathParam:"style=simple,explode=false,name=subscription_id"`
}

type SchedulePlanChangeResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	Subscription *shared.Subscription
}
