# Invoice

## Overview

The Invoice resource represents an invoice that has been generated for a customer. Invoices are generated when a customer's billing interval has elapsed, and are updated when a customer's invoice is paid.

### Available Operations

* [Create](#create) - Create invoice line item
* [Fetch](#fetch) - Retrieve an Invoice
* [FetchUpcoming](#fetchupcoming) - Retrieve upcoming invoice
* [List](#list) - List invoices
* [Void](#void) - Void an invoice

## Create

This creates a one-off fixed fee [Invoice line item](../reference/Orb-API.json/components/schemas/Invoice-line-item) on an [Invoice](../reference/Orb-API.json/components/schemas/Invoice). This can only be done for invoices that are in a `draft` status.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"Orb"
	"Orb/pkg/models/operations"
	"Orb/pkg/types"
)

func main() {
    s := sdk.New(
        sdk.WithSecurity(shared.Security{
            APIKeyAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Invoice.Create(ctx, operations.CreateInvoiceLineItemRequestBody{
        Amount: "accusantium",
        EndDate: types.MustDateFromString("2022-01-07"),
        InvoiceID: "quidem",
        Name: "Colleen Johnston PhD",
        Quantity: 3654.96,
        StartDate: types.MustDateFromString("2022-12-13"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.InvoiceLineItem != nil {
        // handle response
    }
}
```

## Fetch

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
            APIKeyAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Invoice.Fetch(ctx, operations.FetchInvoiceRequest{
        InvoiceID: "fugiat",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Invoice != nil {
        // handle response
    }
}
```

## FetchUpcoming

This endpoint can be used to fetch the [`Upcoming Invoice`](../reference/Orb-API.json/components/schemas/UpcomingInvoice) for the current billing period given a subscription.

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
            APIKeyAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Invoice.FetchUpcoming(ctx, operations.FetchUpcomingInvoiceRequest{
        SubscriptionID: "amet",
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

The list of invoices is ordered starting from the most recently issued invoice date. The response also includes [`pagination_metadata`](../api/pagination), which lets the caller retrieve the next page of results if they exist.

By default, this only returns invoices that are `issued`, `paid`, or `synced`.

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
            APIKeyAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Invoice.List(ctx, operations.ListInvoicesRequest{
        CustomerID: sdk.String("aut"),
        ExternalCustomerID: sdk.String("cumque"),
        Status: sdk.String("corporis"),
        SubscriptionID: sdk.String("hic"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListInvoices200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## Void

This endpoint allows an invoice's status to be set the `void` status. This can only be done to invoices that are in the `issued` status.

If the associated invoice has used the customer balance to change the amount due, the customer balance operation will be reverted. For example, if the invoice used $10 of customer balance, that amount will be added back to the customer balance upon voiding.

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
            APIKeyAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Invoice.Void(ctx, operations.PostInvoicesInvoiceIDVoidRequest{
        InvoiceID: "libero",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Invoice != nil {
        // handle response
    }
}
```
