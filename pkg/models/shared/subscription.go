// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"time"
)

type SubscriptionStatusEnum string

const (
	SubscriptionStatusEnumActive   SubscriptionStatusEnum = "active"
	SubscriptionStatusEnumEnded    SubscriptionStatusEnum = "ended"
	SubscriptionStatusEnumUpcoming SubscriptionStatusEnum = "upcoming"
)

func (e *SubscriptionStatusEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "active":
		fallthrough
	case "ended":
		fallthrough
	case "upcoming":
		*e = SubscriptionStatusEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for SubscriptionStatusEnum: %s", s)
	}
}

// Subscription - A subscription represents the purchase of a plan by a customer.
//
// By default, subscriptions begin on the day that they're created and renew automatically for each billing cycle at the cadence that's configured in the plan definition.
//
// Subscriptions also default to **beginning of month alignment**, which means the first invoice issued for the subscription will have pro-rated charges between the `start_date` and the first of the following month. Subsequent billing periods will always start and end on a month boundary (e.g. subsequent month starts for monthly billing).
//
// Depending on the plan configuration, any _flat_ recurring fees will be billed either at the beginning (in-advance) or end (in-arrears) of each billing cycle. Plans default to **in-advance billing**. Usage-based fees are billed in arrears as usage is accumulated. In the normal course of events, you can expect an invoice to contain usage-based charges for the previous period, and a recurring fee for the following period.
type Subscription struct {
	// The current plan phase that is active, only if the subscription's plan has phases.
	ActivePlanPhaseOrder *float64  `json:"active_plan_phase_order,omitempty"`
	CreatedAt            time.Time `json:"created_at"`
	// The end of the current billing period. This is an exclusive timestamp, such that the instant returned is not part of the billing period. Set to null for subscriptions that are not currently active.
	CurrentBillingPeriodEndDate *time.Time `json:"current_billing_period_end_date,omitempty"`
	// The start of the current billing period. This is an inclusive timestamp; the instant returned is exactly the beginning of the billing period. Set to null if the subscription is not currently active.
	CurrentBillingPeriodStartDate *time.Time `json:"current_billing_period_start_date,omitempty"`
	// A customer is a buyer of your products, and the other party to the billing relationship.
	//
	// In Orb, customers are assigned system generated identifiers automatically, but it's often desirable to have these match existing identifiers in your system. To avoid having to denormalize Orb ID information, you can pass in an `external_customer_id` with your own identifier. See [Customer ID Aliases](../docs/Customer-ID-Aliases.md) for further information about how these aliases work in Orb.
	//
	// In addition to having an identifier in your system, a customer may exist in a payment provider solution like Stripe. Use the `payment_provider_id` and the `payment_provider` enum field to express this mapping.
	//
	// A customer also has a timezone (from the standard [IANA timezone database](https://www.iana.org/time-zones)), which defaults to your account's timezone. See [Timezone localization](../docs/Timezone-localization.md) for information on what this timezone parameter influences within Orb.
	Customer Customer `json:"customer"`
	// The date Orb stops billing for this subscription.
	EndDate time.Time `json:"end_date"`
	ID      string    `json:"id"`
	Plan    Plan      `json:"plan"`
	// The date Orb starts billing for this subscription.
	StartDate time.Time              `json:"start_date"`
	Status    SubscriptionStatusEnum `json:"status"`
}
