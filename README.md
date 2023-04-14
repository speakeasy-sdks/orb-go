<div align="center">
    <picture>
        <source srcset="https://user-images.githubusercontent.com/6267663/229776363-b219eaec-e1aa-4192-9123-d8a8e0ab997d.svg" media="(prefers-color-scheme: dark)">
        <img src="https://user-images.githubusercontent.com/6267663/229776275-b670d564-fc2e-4843-b061-adf230737e3f.svg">
    </picture>
    <h1>Go SDK</h1>
   <p>The modern pricing platform to bill for seats, consumption, and everything in between.</p>
   <a href="https://docs.withorb.com/docs/orb-docs/overview"><img src="https://img.shields.io/static/v1?label=Docs&message=API Ref&color=5444e4&style=for-the-badge" /></a>
   <a href="https://github.com/speakeasy-sdks/orb-go/actions"><img src="https://img.shields.io/github/actions/workflow/status/speakeasy-sdks/orb-go/speakeasy_sdk_generation.yml?style=for-the-badge" /></a>
  <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-blue.svg?style=for-the-badge" /></a>
  <a href="https://github.com/speakeasy-sdks/orb-go/releases"><img src="https://img.shields.io/github/v/release/speakeasy-sdks/orb-go?sort=semver&style=for-the-badge" /></a>
</div>

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
    "Orb"
    "Orb/pkg/models/shared"
    "Orb/pkg/models/operations"
)

func main() {
    s := sdk.New(
        sdk.WithSecurity(shared.Security{
            BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.PostCustomersRequestBody{
        BillingAddress: &operations.PostCustomersRequestBodyBillingAddress{
            City: "Laruecester",
            Country: "US",
            Line1: "quibusdam",
            Line2: "unde",
            PostalCode: "58466-3428",
            State: "ipsa",
        },
        Currency: "delectus",
        Email: "Geraldine_Kreiger52@gmail.com",
        ExternalCustomerID: "iusto",
        Name: "Charlie Walsh II",
        PaymentProvider: "quickbooks",
        PaymentProviderID: "deserunt",
        ShippingAddress: &operations.PostCustomersRequestBodyShippingAddress{
            City: "West Ritaworth",
            Country: "US",
            Line1: "quo",
            Line2: "odit",
            PostalCode: "89478-4576",
            State: "dicta",
        },
        Timezone: "Etc/UTC",
    }

    res, err := s.Customer.Create(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Customer != nil {
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
* `GetBalance` - Get customer balance transactions
* `GetByExternalID` - Retrieve a customer by external ID
* `GetCosts` - View customer costs
* `GetCostsByExternalID` - View customer costs by external customer ID
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
* `GetByExternalID` - Retrieve a plan by external plan ID
* `List` - List plans

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
