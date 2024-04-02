// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Orb/pkg/models/shared"
	"encoding/json"
	"fmt"
	"net/http"
)

// PostSubscriptionsSubscriptionIDCancelCancelOptionEnum - Determines the timing of subscription cancellation
type PostSubscriptionsSubscriptionIDCancelCancelOptionEnum string

const (
	PostSubscriptionsSubscriptionIDCancelCancelOptionEnumEndOfSubscriptionTerm PostSubscriptionsSubscriptionIDCancelCancelOptionEnum = "end_of_subscription_term"
	PostSubscriptionsSubscriptionIDCancelCancelOptionEnumImmediate             PostSubscriptionsSubscriptionIDCancelCancelOptionEnum = "immediate"
)

func (e *PostSubscriptionsSubscriptionIDCancelCancelOptionEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "end_of_subscription_term":
		fallthrough
	case "immediate":
		*e = PostSubscriptionsSubscriptionIDCancelCancelOptionEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for PostSubscriptionsSubscriptionIDCancelCancelOptionEnum: %s", s)
	}
}

type PostSubscriptionsSubscriptionIDCancelRequest struct {
	// Determines the timing of subscription cancellation
	CancelOption   PostSubscriptionsSubscriptionIDCancelCancelOptionEnum `queryParam:"style=form,explode=true,name=cancel_option"`
	SubscriptionID string                                                `pathParam:"style=simple,explode=false,name=subscription_id"`
}

type PostSubscriptionsSubscriptionIDCancelResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	Subscription *shared.Subscription
}