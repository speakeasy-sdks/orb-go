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
        PaymentProvider: operations.PostCustomersRequestBodyPaymentProviderEnumQuickbooks.ToPointer(),
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