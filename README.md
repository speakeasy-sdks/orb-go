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

import(
	"context"
	"log"
	"Orb"
	"Orb/pkg/models/operations"
)

func main() {
    s := sdk.New(
        sdk.WithSecurity(shared.Security{
            BearerAuth: "YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Customer.Create(ctx, operations.PostCustomersRequestBody{
        BillingAddress: &operations.PostCustomersRequestBodyBillingAddress{
            City: sdk.String("Laruecester"),
            Country: sdk.String("US"),
            Line1: sdk.String("quibusdam"),
            Line2: sdk.String("unde"),
            PostalCode: sdk.String("58466-3428"),
            State: sdk.String("ipsa"),
        },
        Currency: sdk.String("delectus"),
        Email: "Geraldine_Kreiger52@gmail.com",
        ExternalCustomerID: sdk.String("iusto"),
        Name: "Charlie Walsh II",
        PaymentProvider: operations.PostCustomersRequestBodyPaymentProviderQuickbooks.ToPointer(),
        PaymentProviderID: sdk.String("deserunt"),
        ShippingAddress: &operations.PostCustomersRequestBodyShippingAddress{
            City: sdk.String("West Ritaworth"),
            Country: sdk.String("US"),
            Line1: sdk.String("quo"),
            Line2: sdk.String("odit"),
            PostalCode: sdk.String("89478-4576"),
            State: sdk.String("dicta"),
        },
        Timezone: sdk.String("Etc/UTC"),
    })
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


### [Availability](docs/availability/README.md)

* [Ping](docs/availability/README.md#ping) - Check availability

### [Credits](docs/credits/README.md)

* [Get](docs/credits/README.md#get) - Add credit ledger entry
* [GetCredits](docs/credits/README.md#getcredits) - Retrieve credit balance
* [GetCreditsLedger](docs/credits/README.md#getcreditsledger) - View credits ledger

### [Customer](docs/customer/README.md)

* [Create](docs/customer/README.md#create) - Create customer
* [Get](docs/customer/README.md#get) - Retrieve a customer
* [GetBalance](docs/customer/README.md#getbalance) - Get customer balance transactions
* [GetByExternalID](docs/customer/README.md#getbyexternalid) - Retrieve a customer by external ID
* [GetCosts](docs/customer/README.md#getcosts) - View customer costs
* [GetCostsByExternalID](docs/customer/README.md#getcostsbyexternalid) - View customer costs by external customer ID
* [List](docs/customer/README.md#list) - List customers
* [Update](docs/customer/README.md#update) - Update customer
* [UpdateByExternalID](docs/customer/README.md#updatebyexternalid) - Update a customer by external ID
* [UpdateUsage](docs/customer/README.md#updateusage) - Amend customer usage
* [UpdateUsageByExternalID](docs/customer/README.md#updateusagebyexternalid) - Amend customer usage by external ID

### [Event](docs/event/README.md)

* [Deprecate](docs/event/README.md#deprecate) - Deprecate single event
* [Ingest](docs/event/README.md#ingest) - Ingest events
* [Search](docs/event/README.md#search) - Search events
* [Update](docs/event/README.md#update) - Amend single event

### [Invoice](docs/invoice/README.md)

* [Get](docs/invoice/README.md#get) - Retrieve an Invoice
* [GetUpcoming](docs/invoice/README.md#getupcoming) - Retrieve upcoming invoice
* [List](docs/invoice/README.md#list) - List invoices

### [Plan](docs/plan/README.md)

* [Get](docs/plan/README.md#get) - Retrieve a plan
* [GetByExternalID](docs/plan/README.md#getbyexternalid) - Retrieve a plan by external plan ID
* [List](docs/plan/README.md#list) - List plans

### [Subscription](docs/subscription/README.md)

* [Cancel](docs/subscription/README.md#cancel) - Cancel subscription
* [ChangeSchedule](docs/subscription/README.md#changeschedule) - Schedule plan change
* [Create](docs/subscription/README.md#create) - Create subscription
* [Get](docs/subscription/README.md#get) - Retrieve a subscription
* [GetCost](docs/subscription/README.md#getcost) - View subscription costs
* [GetSchedule](docs/subscription/README.md#getschedule) - View subscription schedule
* [GetUsage](docs/subscription/README.md#getusage) - View subscription usage
* [List](docs/subscription/README.md#list) - List subscriptions
* [Unschedule](docs/subscription/README.md#unschedule) - Unschedule pending plan changes
<!-- End SDK Available Operations -->

### Maturity

This SDK is in beta and therefore, we recommend pinning usage to a specific package version.
This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated and maintained programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release !

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
