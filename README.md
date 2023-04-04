# github.com/speakeasy-sdks/orb-go

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/orb-go
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->
```go
package main

import (
    "context"
    "log"
    "github.com/speakeasy-sdks/orb-go"
    "github.com/speakeasy-sdks/orb-go/pkg/models/shared"
    "github.com/speakeasy-sdks/orb-go/pkg/models/operations"
)

func main() {
    s := sdk.New(
        sdk.WithSecurity(shared.Security{
            BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Availability.Ping(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetPing200ApplicationJSONObject != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### Availability

* `Ping` - Check availability

### Credits

* `Create` - Add credit ledger entry
* `GetCredits` - Retrieve credit balance
* `GetCreditsLedger` - View credits ledger

### Customer

* `Create` - Create customer
* `Get` - Retrieve a customer
* `Get` - View customer costs by external customer ID
* `GetBalance` - Get customer balance transactions
* `GetByExternalID` - Retrieve a customer by external ID
* `GetCosts` - View customer costs
* `List` - List customers
* `Update` - Update customer
* `UpdateByExternalID` - Update a customer by external ID
* `UpdateUsage` - Amend customer usage
* `UpdateUsageByExternalID` - Amend customer usage by external ID

### Event

* `Deprecate` - Deprecate single event
* `Ingest` - Ingest events
* `Search` - Search events
* `Update` - Amend single event

### Invoice

* `Get` - Retrieve an Invoice
* `GetUpcoming` - Retrieve upcoming invoice
* `List` - List invoices

### Plan

* `Get` - Retrieve a plan
* `Get` - Retrieve a plan by external plan ID
* `Get` - List plans

### Subscription

* `Cancel` - Cancel subscription
* `ChangeSchedule` - Schedule plan change
* `Create` - Create subscription
* `Get` - Retrieve a subscription
* `GetCost` - View subscription costs
* `GetSchedule` - View subscription schedule
* `GetUsage` - View subscription usage
* `List` - List subscriptions
* `Unschedule` - Unschedule pending plan changes
<!-- End SDK Available Operations -->

### Maturity

This SDK is in beta and therefore, we recommend pinning usage to a specific package version.
This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated and maintained programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release !

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
