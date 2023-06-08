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