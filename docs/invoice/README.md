# Invoice

## Overview

Actions related to invoice management.

### Available Operations

* [Get](#get) - Retrieve an Invoice
* [GetUpcoming](#getupcoming) - Retrieve upcoming invoice
* [List](#list) - List invoices

## Get

This endpoint is used to fetch an [`Invoice`](../reference/Orb-API.json/components/schemas/Invoice) given an identifier.

### Example Usage

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
    res, err := s.Invoice.Get(ctx, operations.GetInvoiceInvoiceIDRequest{
        InvoiceID: "quasi",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Invoice != nil {
        // handle response
    }
}
```

## GetUpcoming

This endpoint can be used to fetch the [`UpcomingInvoice`](../reference/Orb-API.json/components/schemas/Upcoming%20Invoice) for the current billing period given a subscription.

### Example Usage

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
    res, err := s.Invoice.GetUpcoming(ctx, operations.GetInvoicesUpcomingRequest{
        SubscriptionID: "iure",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpcomingInvoice != nil {
        // handle response
    }
}
```

## List

This endpoint returns a list of all [`Invoice`](../reference/Orb-API.json/components/schemas/Invoice)s for an account in a list format. 

The list of invoices is ordered starting from the most recently issued invoice date. The response also includes `pagination_metadata`, which lets the caller retrieve the next page of results if they exist.

### Example Usage

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
    res, err := s.Invoice.List(ctx, operations.ListInvoicesRequest{
        CustomerID: sdk.String("doloribus"),
        ExternalCustomerID: sdk.String("debitis"),
        SubscriptionID: sdk.String("eius"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListInvoices200ApplicationJSONObject != nil {
        // handle response
    }
}
```
