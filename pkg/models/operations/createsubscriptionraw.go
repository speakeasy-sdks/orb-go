// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Orb/pkg/models/shared"
	"net/http"
)

type CreateSubscriptionRawResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	Subscription *shared.Subscription
}