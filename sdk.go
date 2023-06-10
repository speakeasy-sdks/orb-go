// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdk

import (
	"Orb/pkg/models/shared"
	"Orb/pkg/utils"
	"fmt"
	"net/http"
	"time"
)

// ServerList contains the list of servers available to the SDK
var ServerList = []string{
	// Production server
	"https://api.withorb.com/v1",
}

// HTTPClient provides an interface for suplying the SDK with a custom HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// Bool provides a helper function to return a pointer to a bool
func Bool(b bool) *bool { return &b }

// Int provides a helper function to return a pointer to an int
func Int(i int) *int { return &i }

// Int64 provides a helper function to return a pointer to an int64
func Int64(i int64) *int64 { return &i }

// Float32 provides a helper function to return a pointer to a float32
func Float32(f float32) *float32 { return &f }

// Float64 provides a helper function to return a pointer to a float64
func Float64(f float64) *float64 { return &f }

type sdkConfiguration struct {
	DefaultClient     HTTPClient
	SecurityClient    HTTPClient
	Security          *shared.Security
	ServerURL         string
	ServerIndex       int
	Language          string
	OpenAPIDocVersion string
	SDKVersion        string
	GenVersion        string
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	return ServerList[c.ServerIndex], nil
}

// SDK - API Reference: Orb powers usage-based billing for the fastest-growing software companies.
// Orb's API is built with the following principles in mind:
//
// 1. **Predictable developer experience**: Where applicable, the Orb API uses industry-standard patterns such as cursor-based pagination and standardized error output. To help with debugging in critical API actions, the API always strives to provide detailed and actionable error messages. Aliases such as external customer IDs aid in fast integration times.
// 2. **Reliably real time**: Orb's event-based APIs, such as event ingestion are designed to handle extremely high throughput and scale with concurrent load. Orb also provides a real-time event-level credits ledger and a highly performant webhooks architecture.
// 3. **Flexibility at the forefront**: Features like timezone localization and the ability to amend historical usage show the flexible nature of the platform.
type SDK struct {
	// Availability - The Availability resource represents a customer's availability. Availability is created when a customer's invoice is paid, and is updated when a customer's transaction is refunded.
	Availability *availability
	// Coupon - The Coupon resource represents a discount that can be applied to a customer's invoice. Coupons can be applied to a customer's invoice either by the customer or by the Orb API.
	Coupon *coupon
	// Credit - The Credits resource represents a customer's credits. Credits are created when a customer's invoice is paid, and are updated when a customer's transaction is refunded.
	Credit *credit
	// CreditNote - The Credit Note resource represents a credit note that has been generated for a customer. Credit Notes are generated when a customer's billing interval has elapsed, and are updated when a customer's invoice is paid.
	CreditNote *creditNote
	// Customer - The Customer resource represents a customer of your service. Customers are created when a customer is created in your service, and are updated when a customer's information is updated in your service.
	Customer *customer
	// Event - The Event resource represents an event that has been created for a customer. Events are created when a customer's invoice is paid, and are updated when a customer's transaction is refunded.
	Event *event
	// Invoice - The Invoice resource represents an invoice that has been generated for a customer. Invoices are generated when a customer's billing interval has elapsed, and are updated when a customer's invoice is paid.
	Invoice *invoice
	// Plan - The Plan resource represents a plan that can be subscribed to by a customer. Plans define the amount of credits that a customer will receive, the price of the plan, and the billing interval.
	Plan *plan
	// Subscription - The Subscription resource represents a customer's subscription to a plan. Subscriptions are created when a customer subscribes to a plan, and are updated when a customer's plan is changed.
	Subscription *subscription

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*SDK)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *SDK) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *SDK) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServerIndex allows the overriding of the default server by index
func WithServerIndex(serverIndex int) SDKOption {
	return func(sdk *SDK) {
		if serverIndex < 0 || serverIndex >= len(ServerList) {
			panic(fmt.Errorf("server index %d out of range", serverIndex))
		}

		sdk.sdkConfiguration.ServerIndex = serverIndex
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *SDK) {
		sdk.sdkConfiguration.DefaultClient = client
	}
}

// WithSecurity configures the SDK to use the provided security details
func WithSecurity(security shared.Security) SDKOption {
	return func(sdk *SDK) {
		sdk.sdkConfiguration.Security = &security
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *SDK {
	sdk := &SDK{
		sdkConfiguration: sdkConfiguration{
			Language:          "go",
			OpenAPIDocVersion: "1.0",
			SDKVersion:        "0.13.0",
			GenVersion:        "2.39.0",
		},
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk.sdkConfiguration.DefaultClient == nil {
		sdk.sdkConfiguration.DefaultClient = &http.Client{Timeout: 60 * time.Second}
	}
	if sdk.sdkConfiguration.SecurityClient == nil {
		if sdk.sdkConfiguration.Security != nil {
			sdk.sdkConfiguration.SecurityClient = utils.ConfigureSecurityClient(sdk.sdkConfiguration.DefaultClient, sdk.sdkConfiguration.Security)
		} else {
			sdk.sdkConfiguration.SecurityClient = sdk.sdkConfiguration.DefaultClient
		}
	}

	sdk.Availability = newAvailability(sdk.sdkConfiguration)

	sdk.Coupon = newCoupon(sdk.sdkConfiguration)

	sdk.Credit = newCredit(sdk.sdkConfiguration)

	sdk.CreditNote = newCreditNote(sdk.sdkConfiguration)

	sdk.Customer = newCustomer(sdk.sdkConfiguration)

	sdk.Event = newEvent(sdk.sdkConfiguration)

	sdk.Invoice = newInvoice(sdk.sdkConfiguration)

	sdk.Plan = newPlan(sdk.sdkConfiguration)

	sdk.Subscription = newSubscription(sdk.sdkConfiguration)

	return sdk
}
