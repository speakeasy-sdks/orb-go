// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdk

import (
	"Orb/pkg/models/operations"
	"Orb/pkg/utils"
	"context"
	"fmt"
	"net/http"
	"strings"
)

// event - Actions related to event management.
type event struct {
	defaultClient  HTTPClient
	securityClient HTTPClient
	serverURL      string
	language       string
	sdkVersion     string
	genVersion     string
}

func newEvent(defaultClient, securityClient HTTPClient, serverURL, language, sdkVersion, genVersion string) *event {
	return &event{
		defaultClient:  defaultClient,
		securityClient: securityClient,
		serverURL:      serverURL,
		language:       language,
		sdkVersion:     sdkVersion,
		genVersion:     genVersion,
	}
}

// Deprecate - Deprecate single event
// This endpoint is used to deprecate a single usage event with a given `event_id`. `event_id` refers to the `idempotency_key` passed in during ingestion.
//
// This endpoint will mark the existing event as ignored. Note that if you attempt to re-ingest an event with the same `event_id` as a deprecated event, Orb will return an error.
//
// This is a powerful and audit-safe mechanism to retroactively deprecate a single event in cases where you need to:
// * no longer bill for an event that was improperly reported
// * no longer bill for an event based on the result of an external API call (ex. call to a payment gateway failed and the user should not be billed)
//
// If you want to only change specific properties of an event, but keep the event as part of the billing calculation, use the [Amend single event](../reference/Orb-API.json/paths/~1events~1{event_id}/put) endpoint instead.
//
// This API is always audit-safe. The process will still retain the deprecated event, though it will be ignored for billing calculations. For auditing and data fidelity purposes, Orb never overwrites or permanently deletes ingested usage data.
//
// ## Request validation
// * Orb does not accept an `idempotency_key` with the event in this endpoint, since this request is by design idempotent. On retryable errors, you should retry the request and assume the deprecation operation has not succeeded until receipt of a 2xx.
// * The event's `timestamp` must fall within the customer's current subscription's billing period, or within the grace period of the customer's current subscription's previous billing period. Orb does not allow deprecating events for billing periods that have already invoiced customers.
// * The `customer_id` or the `external_customer_id` of the original event ingestion request must identify a Customer resource within Orb, even if this event was ingested during the initial integration period. We do not allow deprecating events for customers not in the Orb system.
func (s *event) Deprecate(ctx context.Context, request operations.PutDeprecateEventsEventIDRequest) (*operations.PutDeprecateEventsEventIDResponse, error) {
	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/events/{event_id}/deprecate", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "PUT", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json;q=1, application/json;q=0")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s", s.language, s.sdkVersion, s.genVersion))

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.PutDeprecateEventsEventIDResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.PutDeprecateEventsEventID200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.PutDeprecateEventsEventID200ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 400:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.PutDeprecateEventsEventID400ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.PutDeprecateEventsEventID400ApplicationJSONObject = out
		}
	}

	return res, nil
}

// Ingest - Ingest events
// Orb's event ingestion model and API is designed around two core principles:
//
// 1. **Data fidelity**: The accuracy of your billing model depends on a robust foundation of events. Orb's API protocol encourages usage patterns that ensure that your data is consistently complete and correct.
// 2. **Fast integration**: Sending events into Orb requires no tedious setup steps or explicit field schema for your event shape, making it instant to start streaming in usage in real-time.
//
// ## Event shape
//
// Events are the starting point for all usage calculations in the system, and are simple at their core:
//
// ```json
//
//	{
//	  // customer_id and external_customer_id are used to
//	  // attribute usage to a given Customer. Exactly one of these
//	  // should be specified in a given ingestion event.
//
//	  // `customer_id` is the Orb generated identifier for the Customer,
//	  // which is returned from the Create Customer API call.
//	  customer_id: string,
//
//	  // external_customer_id is an alternate identifier which is associated
//	  // with a Customer at creation time. This is treated as an alias for
//	  // customer_id, and is usually set to an identifier native to your system.
//	  external_customer_id: string,
//
//	  // A string name identifying the event, usually a usage
//	  // action. By convention, this should not contain any whitespace.
//	  event_name: string,
//
//	  // An ISO 8601 format date with no timezone offset.
//	  // This should represent the time that usage occurred
//	  // and is important to attribute usage to a given
//	  // billing period. See the notes below on determining the timestamp.
//	  // e.g. 2020-12-09T16:09:53Z
//	  timestamp: string,
//
//	  // A unique value, generated by the client, that is
//	  // used to de-duplicate events.
//	  // Exactly one event with a given
//	  // idempotency key will be ingested, which allows for
//	  // safe request retries.
//	  idempotency_key: string
//
//	  // Optional custom metadata to attach to the event.
//	  // This might include a numeric value used for aggregation,
//	  // or a string/boolean value used for filtering.
//	  // The schema of this dictionary need not be pre-declared, and
//	  // properties can be added at any time.
//	  properties: {
//	    [key: string]?: string | number | boolean,
//	  },
//	}
//
// ```
//
// ## Required fields
// Because events streamed to Orb are meant to be as flexible as possible, there are only a few required fields in every event.
//
// - We recommend that `idempotency_key` are unique strings that you generated with V4 UUIDs, but only require that they uniquely identify an event (i.e. don’t collide).
// - The `timestamp` field in the event body will be used to determine which billable period a given event falls into. For example, with a monthly billing cycle starting from the first of December, Orb will calculate metrics based on events that fall into the range `12-01 00:00:00 <= timestamp < 01-01 00:00:00`.
//
// ## Logging metadata
//
// Orb allows tagging events with metadata using a flexible properties dictionary. Since Orb does not enforce a rigid schema for this field-set, key-value pairs can be added dynamically as your events evolve.
//
// This dictionary can be helpful for a wide variety of use cases:
//
// - Numeric properties on events like `compute_time_ms` can later be inputs to our flexible query engine to determine usage.
// - Logging a region or cluster with each event can help you provide customers more granular visibility into their usage.
//
// We encourage logging this metadata with an eye towards future use cases to ensure full coverage for historical data. The datatype of the value in the properties dictionary is important for metric creation from an event source. Values that you wish to numerically aggregate should be of numeric type in the event.
//
// ## Determining event timestamp
// For cases where usage is being reported in real time as it is occurring, timestamp should correspond to the time that usage occurred.
//
// In cases where usage is reported in aggregate for a historical timeframe at a regular interval, we recommend setting the event `timestamp` to the midpoint of the interval. As an example, if you have an hourly reporter that sends data once an hour for the previous hour of usage, setting the `timestamp` to the half-hour mark will ensure that the usage is counted within the correct period.
//
// Note that other time-related fields (e.g. time elapsed) can be added to the properties dictionary as necessary.
//
// In cases where usage is reported in aggregate for a historical timeframe, the timestamp must be within the grace period set for your account. Events with `timestamp < current_time - grace_period` will not be accepted as a valid event, and will throw validation errors. Enforcing the grace period enables Orb to accurately map usage to the correct billing cycle and ensure that all usage is billed for in the corresponding billing period.
//
// ## Event validation
//
// Orb’s validation ensures that you recognize errors in your events as quickly as possible, and the API provides informative error messages to help you fix problems quickly.
//
// We validate the following:
//
// - Exactly one of `customer_id` and `external_customer_id` should be specified.
// - If specified, `customer_id` must identify a Customer resource within Orb. We do not support sending events for customers that have not been provisioned. Similarly, if specified, `external_customer_id` must be an identifier that is associated with an Orb Customer resource. Note: During our initial integration period, this enforcement will be temporarily turned into a warning to ensure smooth customer migration.
// - `timestamp` must conform to ISO 8601 and represent a timestamp at most 1 hour in the future. This timestamp should be sent in UTC timezone (no timezone offset).
//
// ## Idempotency and retry semantics
//
// Orb's idempotency guarantees allow you to implement safe retry logic in the event of network or machine failures, ensuring data fidelity. Each event in the request payload is associated with an idempotency key, and Orb guarantees that a single idempotency key will be successfully ingested at most once.
//
// - Successful responses return a 200 HTTP status code. The response contains information about previously processed events.
// - Requests that return a `4xx` HTTP status code indicate a payload error and contain at least one event with a validation failure. An event with a validation failure can be re-sent to the ingestion endpoint (after the payload is fixed) with the original idempotency key since that key is not marked as processed.
// - Requests that return a `5xx` HTTP status code indicate a server-side failure. These requests should be retried in their entirety.
//
// ## API usage and limits
// The ingestion API is designed made for real-time streaming ingestion and architected for high throughput. Even if events are later deemed unnecessary or filtered out, we encourage you to log them to Orb if they may be relevant to billing calculations in the future.
//
// To take advantage of the real-time features of the Orb platform and avoid any chance of dropped events by producers, we recommend reporting events to Orb frequently. Optionally, events can also be briefly aggregated at the source, as this API accepts an array of event bodies.
//
// Orb does not currently enforce a hard rate-limit for API usage or a maximum request payload size, but please give us a heads up if you’re changing either of these factors by an order of magnitude from initial setup.
//
// ## Testing in debug mode
// The ingestion API supports a debug mode, which returns additional verbose output to indicate which event idempotency keys were newly ingested or duplicates from previous requests. To enable this mode, mark `debug=true` as a query parameter.
//
// If `debug=true` is not specified, the response will only contain `validation_failed`. Orb will still honor the idempotency guarantees set [here](https://docs.withorb.com/docs/orb-docs/event-ingestion#event-volume-and-concurrency) in all cases.
//
// We strongly recommend that you only use debug mode as part of testing your initial Orb integration. Once you're ready to switch to production, disable debug mode to take advantage of improved performance and maximal throughput.
//
// #### Example: ingestion response with `debug=true`
//
// ```json
//
//	{
//	  "debug": {
//	    "duplicate": [],
//	    "ingested": [
//	      "B7E83HDMfJPAunXW",
//	      "SJs5DQJ3TnwSqEZE",
//	      "8SivfDsNKwCeAXim"
//	    ]
//	  },
//	  "validation_failed": []
//	}
//
// ```
//
// #### Example: ingestion response with `debug=false`
//
// ```json
//
//	{
//	  "validation_failed": []
//	}
//
// ```
func (s *event) Ingest(ctx context.Context, request operations.PostIngestRequest) (*operations.PostIngestResponse, error) {
	baseURL := s.serverURL
	url := strings.TrimSuffix(baseURL, "/") + "/ingest"

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "RequestBody", "json")
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json;q=1, application/json;q=0")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s", s.language, s.sdkVersion, s.genVersion))

	req.Header.Set("Content-Type", reqContentType)

	if err := utils.PopulateQueryParams(ctx, req, request, nil); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.PostIngestResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.PostIngest200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.PostIngest200ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 400:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.PostIngest400ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.PostIngest400ApplicationJSONObject = out
		}
	}

	return res, nil
}

// Search - Search events
// This endpoint returns a filtered set of events for an account in a paginated list format.
//
// Note that this is a `POST` endpoint rather than a `GET` endpoint because it employs a JSON body for search criteria rather than query parameters, allowing for a more flexible search syntax.
//
// Note that a search criteria _must_ be specified. Currently, Orb supports the following criteria:
// - `event_ids`: This is an explicit array of IDs to filter by. Note that an event's ID is the `idempotency_key` that was originally used for ingestion.
// - `invoice_id`: This is an issued Orb invoice ID (see also [List Invoices](../reference/Orb-API.json/paths/~1invoices/get)). Orb will fetch all events that were used to calculate the invoice. In the common case, this will be a list of events whose `timestamp` property falls within the billing period specified by the invoice.
//
// By default, Orb does not return _deprecated_ events in this endpoint.
//
// By default, Orb will not throw a `404` if no events matched, Orb will return an empty array for `data` instead.
func (s *event) Search(ctx context.Context, request operations.PostEventsSearchRequestBody) (*operations.PostEventsSearchResponse, error) {
	baseURL := s.serverURL
	url := strings.TrimSuffix(baseURL, "/") + "/events/search"

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "Request", "json")
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s", s.language, s.sdkVersion, s.genVersion))

	req.Header.Set("Content-Type", reqContentType)

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.PostEventsSearchResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.PostEventsSearch200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.PostEventsSearch200ApplicationJSONObject = out
		}
	}

	return res, nil
}

// Update - Amend single event
// This endpoint is used to amend a single usage event with a given `event_id`. `event_id` refers to the `idempotency_key` passed in during ingestion. The event will maintain its existing `event_id` after the amendment.
//
// This endpoint will mark the existing event as ignored, and Orb will only use the new event passed in the body of this request as the source of truth for that `event_id`. Note that a single event can be amended any number of times, so the same event can be overwritten in subsequent calls to this endpoint, or overwritten using the [Amend customer usage](../reference/Orb-API.json/paths/~1customers~1{customer_id}~1usage/patch) endpoint. Only a single event with a given `event_id` will be considered the source of truth at any given time.
//
// This is a powerful and audit-safe mechanism to retroactively update a single event in cases where you need to:
// * update an event with new metadata as you iterate on your pricing model
// * update an event based on the result of an external API call (ex. call to a payment gateway succeeded or failed)
//
// This amendment API is always audit-safe. The process will still retain the original event, though it will be ignored for billing calculations. For auditing and data fidelity purposes, Orb never overwrites or permanently deletes ingested usage data.
//
// ## Request validation
// * The `timestamp` of the new event must match the `timestamp` of the existing event already ingested. As with ingestion, all timestamps must be sent in ISO8601 format with UTC timezone offset.
// * The `customer_id` or `external_customer_id` of the new event must match the `customer_id` or `external_customer_id` of the existing event already ingested. Exactly one of `customer_id` and `external_customer_id` should be specified, and similar to ingestion, the ID must identify a Customer resource within Orb. Unlike ingestion, for event amendment, we strictly enforce that the Customer must be in the Orb system, even during the initial integration period. We do not allow updating the `Customer` an event is associated with.
// * Orb does not accept an `idempotency_key` with the event in this endpoint, since this request is by design idempotent. On retryable errors, you should retry the request and assume the amendment operation has not succeeded until receipt of a 2xx.
// * The event's `timestamp` must fall within the customer's current subscription's billing period, or within the grace period of the customer's current subscription's previous billing period.
func (s *event) Update(ctx context.Context, request operations.PutEventsEventIDRequest) (*operations.PutEventsEventIDResponse, error) {
	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/events/{event_id}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "RequestBody", "json")
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "PUT", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json;q=1, application/json;q=0")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s", s.language, s.sdkVersion, s.genVersion))

	req.Header.Set("Content-Type", reqContentType)

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.PutEventsEventIDResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.PutEventsEventID200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.PutEventsEventID200ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 400:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.PutEventsEventID400ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.PutEventsEventID400ApplicationJSONObject = out
		}
	}

	return res, nil
}
