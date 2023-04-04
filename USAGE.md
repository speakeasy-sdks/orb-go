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