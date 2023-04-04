# Orb

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
        Name: "excepturi",
        PaymentProvider: "bill.com",
        PaymentProviderID: "recusandae",
        ShippingAddress: &operations.PostCustomersRequestBodyShippingAddress{
            City: "Belleville",
            Country: "US",
            Line1: "quis",
            Line2: "veritatis",
            PostalCode: "03897-1889",
            State: "molestiae",
        },
        Timezone: "Etc/UTC",
    }

    ctx := context.Background()
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
