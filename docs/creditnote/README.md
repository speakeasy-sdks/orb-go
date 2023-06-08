# CreditNote

## Overview

The Credit Note resource represents a credit note that has been generated for a customer. Credit Notes are generated when a customer's billing interval has elapsed, and are updated when a customer's invoice is paid.

### Available Operations

* [List](#list) - List credit notes

## List

This endpoint returns a list of all [`Credit Note`](../reference/Orb-API.json/components/schemas/Credit-note)s for an account in a list format. 

The list of credit notes is ordered starting from the most recently created date. The response also includes [`pagination_metadata`](../api/pagination), which lets the caller retrieve the next page of results if they exist.

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
    res, err := s.CreditNote.List(ctx, operations.ListCreditNoteRequest{
        CustomerID: sdk.String("iure"),
        ExternalCustomerID: sdk.String("saepe"),
        SubscriptionID: sdk.String("quidem"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListCreditNote200ApplicationJSONObject != nil {
        // handle response
    }
}
```
