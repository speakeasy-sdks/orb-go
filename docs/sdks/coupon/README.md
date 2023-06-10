# Coupon

## Overview

The Coupon resource represents a discount that can be applied to a customer's invoice. Coupons can be applied to a customer's invoice either by the customer or by the Orb API.

### Available Operations

* [Archive](#archive) - Archive a coupon
* [Create](#create) - Create a coupon
* [Fetch](#fetch) - Retrieve a coupon
* [List](#list) - List coupons
* [ListSubscriptions](#listsubscriptions) - List subscriptions for a coupon

## Archive

This endpoint allows a coupon to be archived. Archived coupons can no longer be redeemed, and will be hidden from lists of active coupons. Additionally, once a coupon is archived, its redemption code can be reused for a different coupon.

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
    res, err := s.Coupon.Archive(ctx, operations.ArchiveCouponRequest{
        CouponID: "corrupti",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Coupon != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.ArchiveCouponRequest](../../models/operations/archivecouponrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.ArchiveCouponResponse](../../models/operations/archivecouponresponse.md), error**


## Create

This endpoint allows the creation of coupons, which can then be redeemed at subscription creation or plan change.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"Orb"
	"Orb/pkg/models/shared"
)

func main() {
    s := sdk.New(
        sdk.WithSecurity(shared.Security{
            APIKeyAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Coupon.Create(ctx, shared.CouponInput{
        Discount: shared.Discount{
            AmountDiscount: sdk.String("provident"),
            AppliesToPriceIds: []string{
                "quibusdam",
                "unde",
                "nulla",
            },
            DiscountType: shared.DiscountDiscountTypePercentage,
            PercentageDiscount: sdk.Float64(0.15),
            TrialAmountDiscount: sdk.String("corrupti"),
            UsageDiscount: sdk.Float64(8472.52),
        },
        DurationInMonths: sdk.Int64(423655),
        ID: sdk.String("9a674e0f-467c-4c87-96ed-151a05dfc2dd"),
        MaxRedemptions: sdk.Int64(978619),
        RedemptionCode: "molestiae",
        TimesRedeemed: sdk.Int64(799159),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `request`                                                | [shared.CouponInput](../../models/shared/couponinput.md) | :heavy_check_mark:                                       | The request object to use for the request.               |


### Response

**[*operations.CreateCouponResponse](../../models/operations/createcouponresponse.md), error**


## Fetch

This endpoint retrieves a coupon by its ID. To fetch coupons by their redemption code, use the [List coupons](list-coupons) endpoint with the `redemption_code` parameter.

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
    res, err := s.Coupon.Fetch(ctx, operations.FetchCouponRequest{
        CouponID: "quod",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Coupon != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.FetchCouponRequest](../../models/operations/fetchcouponrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.FetchCouponResponse](../../models/operations/fetchcouponresponse.md), error**


## List

This endpoint returns a list of all [coupons](../reference/Orb-API.json/components/schemas/Coupon) for an account in a list format. 

The list of coupons is ordered starting from the most recently created coupon. The response also includes [`pagination_metadata`](../api/pagination), which lets the caller retrieve the next page of results if they exist. More information about pagination can be found in the [Pagination-metadata schema](../reference/Orb-API.json/components/schemas/Pagination-metadata).

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
    res, err := s.Coupon.List(ctx, operations.ListCouponsRequest{
        RedemptionCode: sdk.String("esse"),
        ShowArchived: sdk.Bool(false),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListCoupons200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.ListCouponsRequest](../../models/operations/listcouponsrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.ListCouponsResponse](../../models/operations/listcouponsresponse.md), error**


## ListSubscriptions

This endpoint returns a list of all subscriptions that have redeemed a given coupon as a [paginated](../api/pagination) list, ordered starting from the most recently created subscription. For a full discussion of the subscription resource, see [Subscription](../reference/Orb-API.json/components/schemas/Subscription).

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
    res, err := s.Coupon.ListSubscriptions(ctx, operations.ListCouponSubscriptionsRequest{
        CouponID: "totam",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListCouponSubscriptions200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.ListCouponSubscriptionsRequest](../../models/operations/listcouponsubscriptionsrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.ListCouponSubscriptionsResponse](../../models/operations/listcouponsubscriptionsresponse.md), error**

