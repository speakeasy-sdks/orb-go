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