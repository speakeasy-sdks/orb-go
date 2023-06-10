// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"time"
)

// AmendUsageRequestBodyEventsProperties - A dictionary of custom properties. Values in this dictionary must be numeric, boolean, or strings. Nested dictionaries are disallowed.
type AmendUsageRequestBodyEventsProperties struct {
}

type AmendUsageRequestBodyEvents struct {
	// A name to meaningfully identify the action or event type.
	EventName string `json:"event_name"`
	// A dictionary of custom properties. Values in this dictionary must be numeric, boolean, or strings. Nested dictionaries are disallowed.
	Properties AmendUsageRequestBodyEventsProperties `json:"properties"`
	// An ISO 8601 format date with no timezone offset (i.e. UTC). This should represent the time that usage was recorded, and is particularly important to attribute usage to a given billing period.
	Timestamp string `json:"timestamp"`
}

type AmendUsageRequestBody struct {
	Events []AmendUsageRequestBodyEvents `json:"events"`
}

type AmendUsageRequest struct {
	RequestBody *AmendUsageRequestBody `request:"mediaType=application/json"`
	CustomerID  string                 `pathParam:"style=simple,explode=false,name=customer_id"`
	// This bound is exclusive (i.e. events before this timestamp will be updated)
	TimeframeEnd time.Time `queryParam:"style=form,explode=true,name=timeframe_end"`
	// This bound is inclusive (i.e. events with this timestamp onward, inclusive will be updated)
	TimeframeStart time.Time `queryParam:"style=form,explode=true,name=timeframe_start"`
}

type AmendUsage400ApplicationJSONValidationErrors struct {
	IdempotencyKey *string `json:"idempotency_key,omitempty"`
	// An array of strings corresponding to each validation failure
	ValidationErrors []string `json:"validation_errors,omitempty"`
}

// AmendUsage400ApplicationJSON - Bad Request
type AmendUsage400ApplicationJSON struct {
	// HTTP Code
	Status int64 `json:"status"`
	// Error message
	Title string `json:"title"`
	Type  string `json:"type"`
	// Contains all failing validation events.
	ValidationErrors []AmendUsage400ApplicationJSONValidationErrors `json:"validation_errors"`
}

type AmendUsage200ApplicationJSONDuplicate struct {
}

// AmendUsage200ApplicationJSON - OK
type AmendUsage200ApplicationJSON struct {
	// An array of strings, corresponding to idempotency_key's marked as duplicates (previously ingested)
	Duplicate []AmendUsage200ApplicationJSONDuplicate `json:"duplicate,omitempty"`
	// An array of strings, corresponding to idempotency_key's which were successfully ingested.
	Ingested []string `json:"ingested,omitempty"`
}

type AmendUsageResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	AmendUsage200ApplicationJSONObject *AmendUsage200ApplicationJSON
	// Bad Request
	AmendUsage400ApplicationJSONObject *AmendUsage400ApplicationJSON
}
