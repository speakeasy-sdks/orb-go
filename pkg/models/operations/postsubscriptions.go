// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Orb/pkg/models/shared"
	"Orb/pkg/types"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// PostSubscriptionsRequestBodyExternalMarketplaceEnum - Optionally provide the name of the external marketplace that the subscription is attached to.
type PostSubscriptionsRequestBodyExternalMarketplaceEnum string

const (
	PostSubscriptionsRequestBodyExternalMarketplaceEnumGoogle PostSubscriptionsRequestBodyExternalMarketplaceEnum = "google"
	PostSubscriptionsRequestBodyExternalMarketplaceEnumAws    PostSubscriptionsRequestBodyExternalMarketplaceEnum = "aws"
	PostSubscriptionsRequestBodyExternalMarketplaceEnumAzure  PostSubscriptionsRequestBodyExternalMarketplaceEnum = "azure"
)

func (e PostSubscriptionsRequestBodyExternalMarketplaceEnum) ToPointer() *PostSubscriptionsRequestBodyExternalMarketplaceEnum {
	return &e
}

func (e *PostSubscriptionsRequestBodyExternalMarketplaceEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "google":
		fallthrough
	case "aws":
		fallthrough
	case "azure":
		*e = PostSubscriptionsRequestBodyExternalMarketplaceEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PostSubscriptionsRequestBodyExternalMarketplaceEnum: %v", v)
	}
}

type PostSubscriptionsRequestBodyPhaseOverrides struct {
	Discount map[string]interface{} `json:"discount,omitempty"`
	// The new minimum amount for the phase. Providing `null` will clear the existing minimum, if it exists.
	MinimumAmount *string `json:"minimum_amount,omitempty"`
	// The phase order that is being modified.
	Order *float64 `json:"order,omitempty"`
}

type PostSubscriptionsRequestBodyPriceOverrides7ModelTypeEnum string

const (
	PostSubscriptionsRequestBodyPriceOverrides7ModelTypeEnumTieredBps PostSubscriptionsRequestBodyPriceOverrides7ModelTypeEnum = "tiered_bps"
)

func (e PostSubscriptionsRequestBodyPriceOverrides7ModelTypeEnum) ToPointer() *PostSubscriptionsRequestBodyPriceOverrides7ModelTypeEnum {
	return &e
}

func (e *PostSubscriptionsRequestBodyPriceOverrides7ModelTypeEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "tiered_bps":
		*e = PostSubscriptionsRequestBodyPriceOverrides7ModelTypeEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PostSubscriptionsRequestBodyPriceOverrides7ModelTypeEnum: %v", v)
	}
}

type PostSubscriptionsRequestBodyPriceOverrides7TieredBpsConfigTiers struct {
	Bps            float64 `json:"bps"`
	MaximumAmount  string  `json:"maximum_amount"`
	MinimumAmount  string  `json:"minimum_amount"`
	PerUnitMaximum string  `json:"per_unit_maximum"`
}

type PostSubscriptionsRequestBodyPriceOverrides7TieredBpsConfig struct {
	Tiers []PostSubscriptionsRequestBodyPriceOverrides7TieredBpsConfigTiers `json:"tiers"`
}

// PostSubscriptionsRequestBodyPriceOverrides7 - Tiered BPS price override
type PostSubscriptionsRequestBodyPriceOverrides7 struct {
	Discount map[string]interface{} `json:"discount,omitempty"`
	ID       string                 `json:"id"`
	// The subscription's override minimum amount for this price.
	MinimumAmount   *string                                                    `json:"minimum_amount,omitempty"`
	ModelType       PostSubscriptionsRequestBodyPriceOverrides7ModelTypeEnum   `json:"model_type"`
	TieredBpsConfig PostSubscriptionsRequestBodyPriceOverrides7TieredBpsConfig `json:"tiered_bps_config"`
}

type PostSubscriptionsRequestBodyPriceOverrides6BulkBpsConfigTiers struct {
	Bps            float64 `json:"bps"`
	MaximumAmount  string  `json:"maximum_amount"`
	PerUnitMaximum string  `json:"per_unit_maximum"`
}

type PostSubscriptionsRequestBodyPriceOverrides6BulkBpsConfig struct {
	Tiers []PostSubscriptionsRequestBodyPriceOverrides6BulkBpsConfigTiers `json:"tiers,omitempty"`
}

type PostSubscriptionsRequestBodyPriceOverrides6ModelTypeEnum string

const (
	PostSubscriptionsRequestBodyPriceOverrides6ModelTypeEnumBulkBps PostSubscriptionsRequestBodyPriceOverrides6ModelTypeEnum = "bulk_bps"
)

func (e PostSubscriptionsRequestBodyPriceOverrides6ModelTypeEnum) ToPointer() *PostSubscriptionsRequestBodyPriceOverrides6ModelTypeEnum {
	return &e
}

func (e *PostSubscriptionsRequestBodyPriceOverrides6ModelTypeEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "bulk_bps":
		*e = PostSubscriptionsRequestBodyPriceOverrides6ModelTypeEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PostSubscriptionsRequestBodyPriceOverrides6ModelTypeEnum: %v", v)
	}
}

// PostSubscriptionsRequestBodyPriceOverrides6 - Bulk BPS price override
type PostSubscriptionsRequestBodyPriceOverrides6 struct {
	BulkBpsConfig *PostSubscriptionsRequestBodyPriceOverrides6BulkBpsConfig `json:"bulk_bps_config,omitempty"`
	Discount      map[string]interface{}                                    `json:"discount,omitempty"`
	ID            string                                                    `json:"id"`
	// The subscription's override minimum amount for this price.
	MinimumAmount *string                                                  `json:"minimum_amount,omitempty"`
	ModelType     PostSubscriptionsRequestBodyPriceOverrides6ModelTypeEnum `json:"model_type"`
}

type PostSubscriptionsRequestBodyPriceOverrides5BpsConfig struct {
	Bps            float64 `json:"bps"`
	PerUnitMaximum string  `json:"per_unit_maximum"`
}

type PostSubscriptionsRequestBodyPriceOverrides5ModelTypeEnum string

const (
	PostSubscriptionsRequestBodyPriceOverrides5ModelTypeEnumBps PostSubscriptionsRequestBodyPriceOverrides5ModelTypeEnum = "bps"
)

func (e PostSubscriptionsRequestBodyPriceOverrides5ModelTypeEnum) ToPointer() *PostSubscriptionsRequestBodyPriceOverrides5ModelTypeEnum {
	return &e
}

func (e *PostSubscriptionsRequestBodyPriceOverrides5ModelTypeEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "bps":
		*e = PostSubscriptionsRequestBodyPriceOverrides5ModelTypeEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PostSubscriptionsRequestBodyPriceOverrides5ModelTypeEnum: %v", v)
	}
}

// PostSubscriptionsRequestBodyPriceOverrides5 - BPS price override
type PostSubscriptionsRequestBodyPriceOverrides5 struct {
	BpsConfig PostSubscriptionsRequestBodyPriceOverrides5BpsConfig `json:"bps_config"`
	Discount  map[string]interface{}                               `json:"discount,omitempty"`
	ID        string                                               `json:"id"`
	// The subscription's override minimum amount for this price.
	MinimumAmount *string                                                  `json:"minimum_amount,omitempty"`
	ModelType     PostSubscriptionsRequestBodyPriceOverrides5ModelTypeEnum `json:"model_type"`
}

type PostSubscriptionsRequestBodyPriceOverrides4ModelTypeEnum string

const (
	PostSubscriptionsRequestBodyPriceOverrides4ModelTypeEnumPackage PostSubscriptionsRequestBodyPriceOverrides4ModelTypeEnum = "package"
)

func (e PostSubscriptionsRequestBodyPriceOverrides4ModelTypeEnum) ToPointer() *PostSubscriptionsRequestBodyPriceOverrides4ModelTypeEnum {
	return &e
}

func (e *PostSubscriptionsRequestBodyPriceOverrides4ModelTypeEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "package":
		*e = PostSubscriptionsRequestBodyPriceOverrides4ModelTypeEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PostSubscriptionsRequestBodyPriceOverrides4ModelTypeEnum: %v", v)
	}
}

type PostSubscriptionsRequestBodyPriceOverrides4PackageConfig struct {
	PackageAmount string  `json:"package_amount"`
	PackageSize   float64 `json:"package_size"`
}

// PostSubscriptionsRequestBodyPriceOverrides4 - Package price override
type PostSubscriptionsRequestBodyPriceOverrides4 struct {
	Discount map[string]interface{} `json:"discount,omitempty"`
	ID       string                 `json:"id"`
	// The subscription's override minimum amount for this price.
	MinimumAmount *string                                                  `json:"minimum_amount,omitempty"`
	ModelType     PostSubscriptionsRequestBodyPriceOverrides4ModelTypeEnum `json:"model_type"`
	PackageConfig PostSubscriptionsRequestBodyPriceOverrides4PackageConfig `json:"package_config"`
}

type PostSubscriptionsRequestBodyPriceOverrides3BulkConfigTiers struct {
	MaximumUnits string `json:"maximum_units"`
	UnitAmount   string `json:"unit_amount"`
}

type PostSubscriptionsRequestBodyPriceOverrides3BulkConfig struct {
	Tiers []PostSubscriptionsRequestBodyPriceOverrides3BulkConfigTiers `json:"tiers"`
}

type PostSubscriptionsRequestBodyPriceOverrides3ModelTypeEnum string

const (
	PostSubscriptionsRequestBodyPriceOverrides3ModelTypeEnumBulk PostSubscriptionsRequestBodyPriceOverrides3ModelTypeEnum = "bulk"
)

func (e PostSubscriptionsRequestBodyPriceOverrides3ModelTypeEnum) ToPointer() *PostSubscriptionsRequestBodyPriceOverrides3ModelTypeEnum {
	return &e
}

func (e *PostSubscriptionsRequestBodyPriceOverrides3ModelTypeEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "bulk":
		*e = PostSubscriptionsRequestBodyPriceOverrides3ModelTypeEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PostSubscriptionsRequestBodyPriceOverrides3ModelTypeEnum: %v", v)
	}
}

// PostSubscriptionsRequestBodyPriceOverrides3 - Bulk price override
type PostSubscriptionsRequestBodyPriceOverrides3 struct {
	BulkConfig PostSubscriptionsRequestBodyPriceOverrides3BulkConfig `json:"bulk_config"`
	Discount   map[string]interface{}                                `json:"discount,omitempty"`
	ID         string                                                `json:"id"`
	// The subscription's override minimum amount for this price.
	MinimumAmount *string                                                  `json:"minimum_amount,omitempty"`
	ModelType     PostSubscriptionsRequestBodyPriceOverrides3ModelTypeEnum `json:"model_type"`
}

type PostSubscriptionsRequestBodyPriceOverrides2ModelTypeEnum string

const (
	PostSubscriptionsRequestBodyPriceOverrides2ModelTypeEnumUnit PostSubscriptionsRequestBodyPriceOverrides2ModelTypeEnum = "unit"
)

func (e PostSubscriptionsRequestBodyPriceOverrides2ModelTypeEnum) ToPointer() *PostSubscriptionsRequestBodyPriceOverrides2ModelTypeEnum {
	return &e
}

func (e *PostSubscriptionsRequestBodyPriceOverrides2ModelTypeEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "unit":
		*e = PostSubscriptionsRequestBodyPriceOverrides2ModelTypeEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PostSubscriptionsRequestBodyPriceOverrides2ModelTypeEnum: %v", v)
	}
}

type PostSubscriptionsRequestBodyPriceOverrides2UnitConfig struct {
	UnitAmount string `json:"unit_amount"`
}

// PostSubscriptionsRequestBodyPriceOverrides2 - Unit price override
type PostSubscriptionsRequestBodyPriceOverrides2 struct {
	Discount           map[string]interface{} `json:"discount,omitempty"`
	FixedPriceQuantity *int64                 `json:"fixed_price_quantity,omitempty"`
	ID                 string                 `json:"id"`
	// The subscription's override minimum amount for this price.
	MinimumAmount *string                                                  `json:"minimum_amount,omitempty"`
	ModelType     PostSubscriptionsRequestBodyPriceOverrides2ModelTypeEnum `json:"model_type"`
	UnitConfig    PostSubscriptionsRequestBodyPriceOverrides2UnitConfig    `json:"unit_config"`
}

type PostSubscriptionsRequestBodyPriceOverrides1ModelTypeEnum string

const (
	PostSubscriptionsRequestBodyPriceOverrides1ModelTypeEnumTiered PostSubscriptionsRequestBodyPriceOverrides1ModelTypeEnum = "tiered"
)

func (e PostSubscriptionsRequestBodyPriceOverrides1ModelTypeEnum) ToPointer() *PostSubscriptionsRequestBodyPriceOverrides1ModelTypeEnum {
	return &e
}

func (e *PostSubscriptionsRequestBodyPriceOverrides1ModelTypeEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "tiered":
		*e = PostSubscriptionsRequestBodyPriceOverrides1ModelTypeEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PostSubscriptionsRequestBodyPriceOverrides1ModelTypeEnum: %v", v)
	}
}

type PostSubscriptionsRequestBodyPriceOverrides1TieredConfigTiers struct {
	FirstUnit  *string `json:"first_unit,omitempty"`
	LastUnit   *string `json:"last_unit,omitempty"`
	UnitAmount *string `json:"unit_amount,omitempty"`
}

type PostSubscriptionsRequestBodyPriceOverrides1TieredConfig struct {
	Tiers []PostSubscriptionsRequestBodyPriceOverrides1TieredConfigTiers `json:"tiers,omitempty"`
}

// PostSubscriptionsRequestBodyPriceOverrides1 - Tiered price override
type PostSubscriptionsRequestBodyPriceOverrides1 struct {
	Discount map[string]interface{} `json:"discount,omitempty"`
	// price_id
	ID string `json:"id"`
	// The subscription's override minimum amount for this price.
	MinimumAmount *string                                                  `json:"minimum_amount,omitempty"`
	ModelType     PostSubscriptionsRequestBodyPriceOverrides1ModelTypeEnum `json:"model_type"`
	TieredConfig  PostSubscriptionsRequestBodyPriceOverrides1TieredConfig  `json:"tiered_config"`
}

type PostSubscriptionsRequestBodyPriceOverridesType string

const (
	PostSubscriptionsRequestBodyPriceOverridesTypePostSubscriptionsRequestBodyPriceOverrides1 PostSubscriptionsRequestBodyPriceOverridesType = "post-subscriptions_requestBody_price_overrides_1"
	PostSubscriptionsRequestBodyPriceOverridesTypePostSubscriptionsRequestBodyPriceOverrides2 PostSubscriptionsRequestBodyPriceOverridesType = "post-subscriptions_requestBody_price_overrides_2"
	PostSubscriptionsRequestBodyPriceOverridesTypePostSubscriptionsRequestBodyPriceOverrides3 PostSubscriptionsRequestBodyPriceOverridesType = "post-subscriptions_requestBody_price_overrides_3"
	PostSubscriptionsRequestBodyPriceOverridesTypePostSubscriptionsRequestBodyPriceOverrides4 PostSubscriptionsRequestBodyPriceOverridesType = "post-subscriptions_requestBody_price_overrides_4"
	PostSubscriptionsRequestBodyPriceOverridesTypePostSubscriptionsRequestBodyPriceOverrides5 PostSubscriptionsRequestBodyPriceOverridesType = "post-subscriptions_requestBody_price_overrides_5"
	PostSubscriptionsRequestBodyPriceOverridesTypePostSubscriptionsRequestBodyPriceOverrides6 PostSubscriptionsRequestBodyPriceOverridesType = "post-subscriptions_requestBody_price_overrides_6"
	PostSubscriptionsRequestBodyPriceOverridesTypePostSubscriptionsRequestBodyPriceOverrides7 PostSubscriptionsRequestBodyPriceOverridesType = "post-subscriptions_requestBody_price_overrides_7"
)

type PostSubscriptionsRequestBodyPriceOverrides struct {
	PostSubscriptionsRequestBodyPriceOverrides1 *PostSubscriptionsRequestBodyPriceOverrides1
	PostSubscriptionsRequestBodyPriceOverrides2 *PostSubscriptionsRequestBodyPriceOverrides2
	PostSubscriptionsRequestBodyPriceOverrides3 *PostSubscriptionsRequestBodyPriceOverrides3
	PostSubscriptionsRequestBodyPriceOverrides4 *PostSubscriptionsRequestBodyPriceOverrides4
	PostSubscriptionsRequestBodyPriceOverrides5 *PostSubscriptionsRequestBodyPriceOverrides5
	PostSubscriptionsRequestBodyPriceOverrides6 *PostSubscriptionsRequestBodyPriceOverrides6
	PostSubscriptionsRequestBodyPriceOverrides7 *PostSubscriptionsRequestBodyPriceOverrides7

	Type PostSubscriptionsRequestBodyPriceOverridesType
}

func CreatePostSubscriptionsRequestBodyPriceOverridesPostSubscriptionsRequestBodyPriceOverrides1(postSubscriptionsRequestBodyPriceOverrides1 PostSubscriptionsRequestBodyPriceOverrides1) PostSubscriptionsRequestBodyPriceOverrides {
	typ := PostSubscriptionsRequestBodyPriceOverridesTypePostSubscriptionsRequestBodyPriceOverrides1

	return PostSubscriptionsRequestBodyPriceOverrides{
		PostSubscriptionsRequestBodyPriceOverrides1: &postSubscriptionsRequestBodyPriceOverrides1,
		Type: typ,
	}
}

func CreatePostSubscriptionsRequestBodyPriceOverridesPostSubscriptionsRequestBodyPriceOverrides2(postSubscriptionsRequestBodyPriceOverrides2 PostSubscriptionsRequestBodyPriceOverrides2) PostSubscriptionsRequestBodyPriceOverrides {
	typ := PostSubscriptionsRequestBodyPriceOverridesTypePostSubscriptionsRequestBodyPriceOverrides2

	return PostSubscriptionsRequestBodyPriceOverrides{
		PostSubscriptionsRequestBodyPriceOverrides2: &postSubscriptionsRequestBodyPriceOverrides2,
		Type: typ,
	}
}

func CreatePostSubscriptionsRequestBodyPriceOverridesPostSubscriptionsRequestBodyPriceOverrides3(postSubscriptionsRequestBodyPriceOverrides3 PostSubscriptionsRequestBodyPriceOverrides3) PostSubscriptionsRequestBodyPriceOverrides {
	typ := PostSubscriptionsRequestBodyPriceOverridesTypePostSubscriptionsRequestBodyPriceOverrides3

	return PostSubscriptionsRequestBodyPriceOverrides{
		PostSubscriptionsRequestBodyPriceOverrides3: &postSubscriptionsRequestBodyPriceOverrides3,
		Type: typ,
	}
}

func CreatePostSubscriptionsRequestBodyPriceOverridesPostSubscriptionsRequestBodyPriceOverrides4(postSubscriptionsRequestBodyPriceOverrides4 PostSubscriptionsRequestBodyPriceOverrides4) PostSubscriptionsRequestBodyPriceOverrides {
	typ := PostSubscriptionsRequestBodyPriceOverridesTypePostSubscriptionsRequestBodyPriceOverrides4

	return PostSubscriptionsRequestBodyPriceOverrides{
		PostSubscriptionsRequestBodyPriceOverrides4: &postSubscriptionsRequestBodyPriceOverrides4,
		Type: typ,
	}
}

func CreatePostSubscriptionsRequestBodyPriceOverridesPostSubscriptionsRequestBodyPriceOverrides5(postSubscriptionsRequestBodyPriceOverrides5 PostSubscriptionsRequestBodyPriceOverrides5) PostSubscriptionsRequestBodyPriceOverrides {
	typ := PostSubscriptionsRequestBodyPriceOverridesTypePostSubscriptionsRequestBodyPriceOverrides5

	return PostSubscriptionsRequestBodyPriceOverrides{
		PostSubscriptionsRequestBodyPriceOverrides5: &postSubscriptionsRequestBodyPriceOverrides5,
		Type: typ,
	}
}

func CreatePostSubscriptionsRequestBodyPriceOverridesPostSubscriptionsRequestBodyPriceOverrides6(postSubscriptionsRequestBodyPriceOverrides6 PostSubscriptionsRequestBodyPriceOverrides6) PostSubscriptionsRequestBodyPriceOverrides {
	typ := PostSubscriptionsRequestBodyPriceOverridesTypePostSubscriptionsRequestBodyPriceOverrides6

	return PostSubscriptionsRequestBodyPriceOverrides{
		PostSubscriptionsRequestBodyPriceOverrides6: &postSubscriptionsRequestBodyPriceOverrides6,
		Type: typ,
	}
}

func CreatePostSubscriptionsRequestBodyPriceOverridesPostSubscriptionsRequestBodyPriceOverrides7(postSubscriptionsRequestBodyPriceOverrides7 PostSubscriptionsRequestBodyPriceOverrides7) PostSubscriptionsRequestBodyPriceOverrides {
	typ := PostSubscriptionsRequestBodyPriceOverridesTypePostSubscriptionsRequestBodyPriceOverrides7

	return PostSubscriptionsRequestBodyPriceOverrides{
		PostSubscriptionsRequestBodyPriceOverrides7: &postSubscriptionsRequestBodyPriceOverrides7,
		Type: typ,
	}
}

func (u *PostSubscriptionsRequestBodyPriceOverrides) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	postSubscriptionsRequestBodyPriceOverrides1 := new(PostSubscriptionsRequestBodyPriceOverrides1)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&postSubscriptionsRequestBodyPriceOverrides1); err == nil {
		u.PostSubscriptionsRequestBodyPriceOverrides1 = postSubscriptionsRequestBodyPriceOverrides1
		u.Type = PostSubscriptionsRequestBodyPriceOverridesTypePostSubscriptionsRequestBodyPriceOverrides1
		return nil
	}

	postSubscriptionsRequestBodyPriceOverrides2 := new(PostSubscriptionsRequestBodyPriceOverrides2)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&postSubscriptionsRequestBodyPriceOverrides2); err == nil {
		u.PostSubscriptionsRequestBodyPriceOverrides2 = postSubscriptionsRequestBodyPriceOverrides2
		u.Type = PostSubscriptionsRequestBodyPriceOverridesTypePostSubscriptionsRequestBodyPriceOverrides2
		return nil
	}

	postSubscriptionsRequestBodyPriceOverrides3 := new(PostSubscriptionsRequestBodyPriceOverrides3)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&postSubscriptionsRequestBodyPriceOverrides3); err == nil {
		u.PostSubscriptionsRequestBodyPriceOverrides3 = postSubscriptionsRequestBodyPriceOverrides3
		u.Type = PostSubscriptionsRequestBodyPriceOverridesTypePostSubscriptionsRequestBodyPriceOverrides3
		return nil
	}

	postSubscriptionsRequestBodyPriceOverrides4 := new(PostSubscriptionsRequestBodyPriceOverrides4)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&postSubscriptionsRequestBodyPriceOverrides4); err == nil {
		u.PostSubscriptionsRequestBodyPriceOverrides4 = postSubscriptionsRequestBodyPriceOverrides4
		u.Type = PostSubscriptionsRequestBodyPriceOverridesTypePostSubscriptionsRequestBodyPriceOverrides4
		return nil
	}

	postSubscriptionsRequestBodyPriceOverrides5 := new(PostSubscriptionsRequestBodyPriceOverrides5)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&postSubscriptionsRequestBodyPriceOverrides5); err == nil {
		u.PostSubscriptionsRequestBodyPriceOverrides5 = postSubscriptionsRequestBodyPriceOverrides5
		u.Type = PostSubscriptionsRequestBodyPriceOverridesTypePostSubscriptionsRequestBodyPriceOverrides5
		return nil
	}

	postSubscriptionsRequestBodyPriceOverrides6 := new(PostSubscriptionsRequestBodyPriceOverrides6)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&postSubscriptionsRequestBodyPriceOverrides6); err == nil {
		u.PostSubscriptionsRequestBodyPriceOverrides6 = postSubscriptionsRequestBodyPriceOverrides6
		u.Type = PostSubscriptionsRequestBodyPriceOverridesTypePostSubscriptionsRequestBodyPriceOverrides6
		return nil
	}

	postSubscriptionsRequestBodyPriceOverrides7 := new(PostSubscriptionsRequestBodyPriceOverrides7)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&postSubscriptionsRequestBodyPriceOverrides7); err == nil {
		u.PostSubscriptionsRequestBodyPriceOverrides7 = postSubscriptionsRequestBodyPriceOverrides7
		u.Type = PostSubscriptionsRequestBodyPriceOverridesTypePostSubscriptionsRequestBodyPriceOverrides7
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u PostSubscriptionsRequestBodyPriceOverrides) MarshalJSON() ([]byte, error) {
	if u.PostSubscriptionsRequestBodyPriceOverrides1 != nil {
		return json.Marshal(u.PostSubscriptionsRequestBodyPriceOverrides1)
	}

	if u.PostSubscriptionsRequestBodyPriceOverrides2 != nil {
		return json.Marshal(u.PostSubscriptionsRequestBodyPriceOverrides2)
	}

	if u.PostSubscriptionsRequestBodyPriceOverrides3 != nil {
		return json.Marshal(u.PostSubscriptionsRequestBodyPriceOverrides3)
	}

	if u.PostSubscriptionsRequestBodyPriceOverrides4 != nil {
		return json.Marshal(u.PostSubscriptionsRequestBodyPriceOverrides4)
	}

	if u.PostSubscriptionsRequestBodyPriceOverrides5 != nil {
		return json.Marshal(u.PostSubscriptionsRequestBodyPriceOverrides5)
	}

	if u.PostSubscriptionsRequestBodyPriceOverrides6 != nil {
		return json.Marshal(u.PostSubscriptionsRequestBodyPriceOverrides6)
	}

	if u.PostSubscriptionsRequestBodyPriceOverrides7 != nil {
		return json.Marshal(u.PostSubscriptionsRequestBodyPriceOverrides7)
	}

	return nil, nil
}

type PostSubscriptionsRequestBody struct {
	// Align billing periods with the subscription's start_date. If this is not provided, this defaults to aligning billing periods with the start of the month.
	AlignBillingWithSubscriptionStartDate *bool `json:"align_billing_with_subscription_start_date,omitempty"`
	// The ID of the customer to subscribe.
	CustomerID *string `json:"customer_id,omitempty"`
	// The external ID of the customer to subscribe, as an alternate to passing the `customer_id`.
	ExternalCustomerID *string `json:"external_customer_id,omitempty"`
	// Optionally provide the name of the external marketplace that the subscription is attached to.
	ExternalMarketplace *PostSubscriptionsRequestBodyExternalMarketplaceEnum `json:"external_marketplace,omitempty"`
	// The reporting ID to associate this subscription with the external marketplace. Required if external_marketplace is specified.
	ExternalMarketplaceReportingID *string `json:"external_marketplace_reporting_id,omitempty"`
	// The external ID of the plan, which can be used in place of the `plan_id`.
	ExternalPlanID *string `json:"external_plan_id,omitempty"`
	// The subscription's override minimum amount for the plan.
	MinimumAmount *string `json:"minimum_amount,omitempty"`
	// Optionally provide a list of minimum amount or discount overrides for phases on the plan.
	PhaseOverrides []PostSubscriptionsRequestBodyPhaseOverrides `json:"phase_overrides,omitempty"`
	// The plan that the given customer should be subscribed to. The plan determines the pricing and cadence of the subscription.
	PlanID *string `json:"plan_id,omitempty"`
	// Optionally provide a list of overrides for prices on the plan
	PriceOverrides []PostSubscriptionsRequestBodyPriceOverrides `json:"price_overrides,omitempty"`
	// The date that Orb should start billing for the subscription, localized to the customer's timezone. If this is not provided, this defaults to the current date in the customer's timezone.
	StartDate *types.Date `json:"start_date,omitempty"`
}

type PostSubscriptionsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	Subscription *shared.Subscription
}
