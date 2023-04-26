# Availability

## Overview

Actions related to API availability.

### Available Operations

* [Ping](#ping) - Check availability

## Ping

This endpoint allows you to test your connection to the Orb API and check the validity of your API key, passed in the `Authorization` header. This is particularly useful for checking that your environment is set up properly, and is a great choice for connectors and integrations.

This API does not have any side-effects or return any Orb resources.

### Example Usage

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
