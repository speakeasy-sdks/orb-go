// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/orb-go/pkg/models/shared"
	"net/http"
)

type PostSubscriptionsSubscriptionIDUnschedulePendingPlanChangesRequest struct {
	SubscriptionID string `pathParam:"style=simple,explode=false,name=subscription_id"`
}

type PostSubscriptionsSubscriptionIDUnschedulePendingPlanChangesResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	Subscription *shared.Subscription
}
