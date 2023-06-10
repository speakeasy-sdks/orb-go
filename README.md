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
)

func main() {
    s := sdk.New(
        sdk.WithSecurity(shared.Security{
            APIKeyAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Availability.Ping(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.Ping200ApplicationJSONObject != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [Availability](docs/sdks/availability/README.md)

* [Ping](docs/sdks/availability/README.md#ping) - Check availability

### [Coupon](docs/sdks/coupon/README.md)

* [Archive](docs/sdks/coupon/README.md#archive) - Archive a coupon
* [Create](docs/sdks/coupon/README.md#create) - Create a coupon
* [Fetch](docs/sdks/coupon/README.md#fetch) - Retrieve a coupon
* [List](docs/sdks/coupon/README.md#list) - List coupons
* [ListSubscriptions](docs/sdks/coupon/README.md#listsubscriptions) - List subscriptions for a coupon

### [Credit](docs/sdks/credit/README.md)

* [AddByExternalID](docs/sdks/credit/README.md#addbyexternalid) - Add credit ledger entry by external customer ID
* [Create](docs/sdks/credit/README.md#create) - Add credit ledger entry
* [Fetch](docs/sdks/credit/README.md#fetch) - Retrieve credit balance
* [FetchByExternalID](docs/sdks/credit/README.md#fetchbyexternalid) - Retrieve credit balance by external customer ID
* [FetchLedger](docs/sdks/credit/README.md#fetchledger) - View credits ledger
* [FetchLedgerByExternalID](docs/sdks/credit/README.md#fetchledgerbyexternalid) - View credits ledger by external customer ID

### [CreditNote](docs/sdks/creditnote/README.md)

* [List](docs/sdks/creditnote/README.md#list) - List credit notes

### [Customer](docs/sdks/customer/README.md)

* [Amend](docs/sdks/customer/README.md#amend) - Amend customer usage
* [AmendByExternalID](docs/sdks/customer/README.md#amendbyexternalid) - Amend customer usage by external ID
* [Create](docs/sdks/customer/README.md#create) - Create customer
* [CreateTransaction](docs/sdks/customer/README.md#createtransaction) - Create a customer balance transaction
* [Delete](docs/sdks/customer/README.md#delete) - Delete a customer
* [Fetch](docs/sdks/customer/README.md#fetch) - Retrieve a customer
* [FetchByExternalID](docs/sdks/customer/README.md#fetchbyexternalid) - Retrieve a customer by external ID
* [FetchCosts](docs/sdks/customer/README.md#fetchcosts) - View customer costs
* [FetchCostsByExternalID](docs/sdks/customer/README.md#fetchcostsbyexternalid) - View customer costs by external customer ID
* [FetchTransactions](docs/sdks/customer/README.md#fetchtransactions) - Get customer balance transactions
* [List](docs/sdks/customer/README.md#list) - List customers
* [UpdateByExternalID](docs/sdks/customer/README.md#updatebyexternalid) - Update a customer by external ID
* [UpdateCustomer](docs/sdks/customer/README.md#updatecustomer) - Update customer

### [Event](docs/sdks/event/README.md)

* [Amend](docs/sdks/event/README.md#amend) - Amend single event
* [CloseBackfill](docs/sdks/event/README.md#closebackfill) - Close a backfill
* [Create](docs/sdks/event/README.md#create) - Create a backfill
* [DeprecateEvent](docs/sdks/event/README.md#deprecateevent) - Deprecate single event
* [Ingest](docs/sdks/event/README.md#ingest) - Ingest events
* [ListBackfills](docs/sdks/event/README.md#listbackfills) - List backfills
* [RevertBackfill](docs/sdks/event/README.md#revertbackfill) - Revert a backfill
* [Search](docs/sdks/event/README.md#search) - Search events

### [Invoice](docs/sdks/invoice/README.md)

* [Create](docs/sdks/invoice/README.md#create) - Create invoice line item
* [Fetch](docs/sdks/invoice/README.md#fetch) - Retrieve an Invoice
* [FetchUpcoming](docs/sdks/invoice/README.md#fetchupcoming) - Retrieve upcoming invoice
* [List](docs/sdks/invoice/README.md#list) - List invoices
* [Void](docs/sdks/invoice/README.md#void) - Void an invoice

### [Plan](docs/sdks/plan/README.md)

* [Fetch](docs/sdks/plan/README.md#fetch) - Retrieve a plan
* [GetByExternalID](docs/sdks/plan/README.md#getbyexternalid) - Retrieve a plan by external plan ID
* [List](docs/sdks/plan/README.md#list) - List plans

### [Subscription](docs/sdks/subscription/README.md)

* [Cancel](docs/sdks/subscription/README.md#cancel) - Cancel subscription
* [Create](docs/sdks/subscription/README.md#create) - Create subscription
* [Create](docs/sdks/subscription/README.md#create) - Create subscription
* [Fetch](docs/sdks/subscription/README.md#fetch) - Retrieve a subscription
* [FetchCosts](docs/sdks/subscription/README.md#fetchcosts) - View subscription costs
* [FetchSchedule](docs/sdks/subscription/README.md#fetchschedule) - View subscription schedule
* [FetchUsage](docs/sdks/subscription/README.md#fetchusage) - View subscription usage
* [List](docs/sdks/subscription/README.md#list) - List subscriptions
* [SchedulePlanChange](docs/sdks/subscription/README.md#scheduleplanchange) - Schedule plan change
* [UnscheduleCancellation](docs/sdks/subscription/README.md#unschedulecancellation) - Unschedule pending cancellation
* [UnschedulePlanChange](docs/sdks/subscription/README.md#unscheduleplanchange) - Unschedule pending plan changes
* [UpdateFixedFeeQuantity](docs/sdks/subscription/README.md#updatefixedfeequantity) - Update fixed fee quantity
<!-- End SDK Available Operations -->

### Maturity

This SDK is in beta and therefore, we recommend pinning usage to a specific package version.
This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated and maintained programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release !

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
