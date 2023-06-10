# Plan

## Overview

The Plan resource represents a plan that can be subscribed to by a customer. Plans define the amount of credits that a customer will receive, the price of the plan, and the billing interval.

### Available Operations

* [Fetch](#fetch) - Retrieve a plan
* [GetByExternalID](#getbyexternalid) - Retrieve a plan by external plan ID
* [List](#list) - List plans

## Fetch

This endpoint is used to fetch [plan](../reference/Orb-API.json/components/schemas/Plan) details given a plan identifier. It returns information about the prices included in the plan and their configuration, as well as the product that the plan is attached to.

## Serialized prices
Orb supports a few different pricing models out of the box. Each of these models is serialized differently in a given [Price](../reference/Orb-API.json/components/schemas/Price) object. The `model_type` field determines the key for the configuration object that is present. A detailed explanation of price types can be found in the [Price schema](../reference/Orb-API.json/components/schemas/Price). 

## Phases
Orb supports plan phases, also known as contract ramps. For plans with phases, the serialized prices refer to all prices across all phases.

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
    res, err := s.Plan.Fetch(ctx, operations.GetPlansPlanIDRequest{
        PlanID: "debitis",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Plan != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetPlansPlanIDRequest](../../models/operations/getplansplanidrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.GetPlansPlanIDResponse](../../models/operations/getplansplanidresponse.md), error**


## GetByExternalID

This endpoint is used to fetch [plan](../reference/Orb-API.json/components/schemas/Plan) details given an external_plan_id identifier. It returns information about the prices included in the plan and their configuration, as well as the product that the plan is attached to.

## Serialized prices
Orb supports a few different pricing models out of the box. Each of these models is serialized differently in a given [Price](../reference/Orb-API.json/components/schemas/Price) object. The `model_type` field determines the key for the configuration object that is present. A detailed explanation of price types can be found in the [Price schema](../reference/Orb-API.json/components/schemas/Price). 

### Example Usage

```go
package main

import(
	"context"
	"log"
	"Orb"
	"Orb/pkg/models/operations"
	"Orb/pkg/models/shared"
	"Orb/pkg/types"
)

func main() {
    s := sdk.New(
        sdk.WithSecurity(shared.Security{
            APIKeyAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Plan.GetByExternalID(ctx, operations.GetPlansExternalPlanIDRequest{
        Plan: &shared.Plan{
            BasePlan: &shared.PlanBasePlan{
                ExternalPlanID: sdk.String("eius"),
                ID: sdk.String("c8b711e5-b7fd-42ed-8289-21cddc692601"),
                Name: sdk.String("Rickey Hintz"),
            },
            BasePlanID: sdk.String("nam"),
            CreatedAt: types.MustTimeFromString("2022-02-18T18:29:26.833Z"),
            Currency: "nemo",
            DefaultInvoiceMemo: sdk.String("voluptatibus"),
            Description: "perferendis",
            Discount: shared.Discount{
                AmountDiscount: sdk.String("fugiat"),
                AppliesToPriceIds: []string{
                    "aut",
                },
                DiscountType: shared.DiscountDiscountTypePercentage,
                PercentageDiscount: sdk.Float64(0.15),
                TrialAmountDiscount: sdk.String("cumque"),
                UsageDiscount: sdk.Float64(3599.78),
            },
            ExternalPlanID: sdk.String("hic"),
            ID: "bb258705-3202-4c73-95fe-9b90c28909b3",
            InvoicingCurrency: "asperiores",
            Minimum: shared.MinimumAmount{
                AppliesToPriceIds: []string{
                    "modi",
                    "iste",
                    "dolorum",
                    "deleniti",
                },
                MinimumAmount: "pariatur",
            },
            Name: "Loren Renner",
            NetTerms: sdk.Int64(554242),
            PlanPhases: []shared.PlanPhase{
                shared.PlanPhase{
                    Description: sdk.String("dolorem"),
                    Discount: shared.Discount{
                        AmountDiscount: sdk.String("dolorem"),
                        AppliesToPriceIds: []string{
                            "qui",
                        },
                        DiscountType: shared.DiscountDiscountTypePercentage,
                        PercentageDiscount: sdk.Float64(0.15),
                        TrialAmountDiscount: sdk.String("ipsum"),
                        UsageDiscount: sdk.Float64(9443.73),
                    },
                    Duration: sdk.Int64(569574),
                    DurationUnit: shared.PlanPhaseDurationUnitAnnual,
                    Minimum: shared.MinimumAmount{
                        AppliesToPriceIds: []string{
                            "dignissimos",
                            "reiciendis",
                        },
                        MinimumAmount: "amet",
                    },
                    Name: sdk.String("Mr. Bradley Bogan"),
                    Order: sdk.Int64(487838),
                },
                shared.PlanPhase{
                    Description: sdk.String("quaerat"),
                    Discount: shared.Discount{
                        AmountDiscount: sdk.String("accusamus"),
                        AppliesToPriceIds: []string{
                            "voluptatibus",
                            "voluptas",
                            "natus",
                        },
                        DiscountType: shared.DiscountDiscountTypePercentage,
                        PercentageDiscount: sdk.Float64(0.15),
                        TrialAmountDiscount: sdk.String("eos"),
                        UsageDiscount: sdk.Float64(5424.99),
                    },
                    Duration: sdk.Int64(24678),
                    DurationUnit: shared.PlanPhaseDurationUnitAnnual,
                    Minimum: shared.MinimumAmount{
                        AppliesToPriceIds: []string{
                            "soluta",
                        },
                        MinimumAmount: "dolorum",
                    },
                    Name: sdk.String("Colleen Pagac"),
                    Order: sdk.Int64(896672),
                },
            },
            Prices: []shared.Price{
                shared.Price{
                    BillableMetric: &shared.PriceBillableMetric{
                        ID: sdk.String("f737ae42-03ce-45e6-a95d-8a0d446ce2af"),
                    },
                    BpsConfig: &shared.PriceBpsConfig{
                        Bps: sdk.Float64(4561.3),
                        PerUnitMaximum: sdk.String("harum"),
                    },
                    BulkBpsConfig: &shared.PriceBulkBpsConfig{
                        Tiers: []shared.PriceBulkBpsConfigTiers{
                            shared.PriceBulkBpsConfigTiers{
                                Bps: sdk.Float64(2155.07),
                                MaximumAmount: sdk.String("quisquam"),
                                PerUnitMaximum: sdk.String("tenetur"),
                            },
                            shared.PriceBulkBpsConfigTiers{
                                Bps: sdk.Float64(2294.42),
                                MaximumAmount: sdk.String("tempore"),
                                PerUnitMaximum: sdk.String("accusamus"),
                            },
                        },
                    },
                    BulkConfig: &shared.PriceBulkConfig{
                        Tiers: []shared.PriceBulkConfigTiers{
                            shared.PriceBulkConfigTiers{
                                MaximumUnits: sdk.String("enim"),
                                UnitAmount: sdk.String("dolorem"),
                            },
                            shared.PriceBulkConfigTiers{
                                MaximumUnits: sdk.String("sapiente"),
                                UnitAmount: sdk.String("totam"),
                            },
                        },
                    },
                    Cadence: shared.PriceCadenceMonthly.ToPointer(),
                    CreatedAt: types.MustTimeFromString("2022-04-16T06:31:32.314Z"),
                    Currency: "USD",
                    Discount: shared.Discount{
                        AmountDiscount: sdk.String("neque"),
                        AppliesToPriceIds: []string{
                            "vel",
                        },
                        DiscountType: shared.DiscountDiscountTypePercentage,
                        PercentageDiscount: sdk.Float64(0.15),
                        TrialAmountDiscount: sdk.String("libero"),
                        UsageDiscount: sdk.Float64(3741.7),
                    },
                    FixedPriceQuantity: sdk.Float64(6462.65),
                    ID: sdk.String("73429cdb-1a84-422b-b679-d2322715bf0c"),
                    MatrixConfig: &shared.PriceMatrixConfig{
                        DefaultUnitAmount: sdk.String("soluta"),
                        Dimensions: []string{
                            "et",
                            "saepe",
                            "ipsum",
                        },
                        MatrixValues: []shared.PriceMatrixConfigMatrixValues{
                            shared.PriceMatrixConfigMatrixValues{
                                DimensionValues: []string{
                                    "quos",
                                    "tempore",
                                    "cupiditate",
                                },
                                UnitAmount: sdk.String("aperiam"),
                            },
                        },
                    },
                    Minimum: shared.MinimumAmount{
                        AppliesToPriceIds: []string{
                            "dolorem",
                            "dolore",
                            "labore",
                            "adipisci",
                        },
                        MinimumAmount: "dolorum",
                    },
                    ModelType: shared.PriceModelTypeUnit.ToPointer(),
                    Name: sdk.String("Margaret Luettgen MD"),
                    PackageConfig: &shared.PricePackageConfig{
                        PackageAmount: sdk.String("repellendus"),
                        PackageSize: sdk.Float64(7851.53),
                    },
                    PlanPhaseOrder: 9843.3,
                    TieredBpsConfig: &shared.PriceTieredBpsConfig{
                        Tiers: []shared.PriceTieredBpsConfigTiers{
                            shared.PriceTieredBpsConfigTiers{
                                Bps: sdk.Float64(7034.95),
                                MaximumAmount: sdk.String("cupiditate"),
                                MinimumAmount: sdk.String("qui"),
                                PerUnitMaximum: sdk.String("quae"),
                            },
                            shared.PriceTieredBpsConfigTiers{
                                Bps: sdk.Float64(5123.93),
                                MaximumAmount: sdk.String("odio"),
                                MinimumAmount: sdk.String("occaecati"),
                                PerUnitMaximum: sdk.String("voluptatibus"),
                            },
                        },
                    },
                    TieredConfig: &shared.PriceTieredConfig{
                        Tiers: []shared.PriceTieredConfigTiers{
                            shared.PriceTieredConfigTiers{
                                FirstUnit: sdk.String("vero"),
                                LastUnit: sdk.String("omnis"),
                                UnitAmount: sdk.String("quis"),
                            },
                            shared.PriceTieredConfigTiers{
                                FirstUnit: sdk.String("ipsum"),
                                LastUnit: sdk.String("delectus"),
                                UnitAmount: sdk.String("voluptate"),
                            },
                            shared.PriceTieredConfigTiers{
                                FirstUnit: sdk.String("consectetur"),
                                LastUnit: sdk.String("vero"),
                                UnitAmount: sdk.String("tenetur"),
                            },
                            shared.PriceTieredConfigTiers{
                                FirstUnit: sdk.String("dignissimos"),
                                LastUnit: sdk.String("hic"),
                                UnitAmount: sdk.String("distinctio"),
                            },
                        },
                    },
                    UnitConfig: &shared.PriceUnitConfig{
                        UnitAmount: sdk.String("quod"),
                    },
                },
                shared.Price{
                    BillableMetric: &shared.PriceBillableMetric{
                        ID: sdk.String("7abd74dd-39c0-4f5d-acff-7c70a45626d4"),
                    },
                    BpsConfig: &shared.PriceBpsConfig{
                        Bps: sdk.Float64(1898.48),
                        PerUnitMaximum: sdk.String("ex"),
                    },
                    BulkBpsConfig: &shared.PriceBulkBpsConfig{
                        Tiers: []shared.PriceBulkBpsConfigTiers{
                            shared.PriceBulkBpsConfigTiers{
                                Bps: sdk.Float64(1206.57),
                                MaximumAmount: sdk.String("dolor"),
                                PerUnitMaximum: sdk.String("maiores"),
                            },
                            shared.PriceBulkBpsConfigTiers{
                                Bps: sdk.Float64(978.44),
                                MaximumAmount: sdk.String("ex"),
                                PerUnitMaximum: sdk.String("nulla"),
                            },
                            shared.PriceBulkBpsConfigTiers{
                                Bps: sdk.Float64(5692.11),
                                MaximumAmount: sdk.String("voluptatibus"),
                                PerUnitMaximum: sdk.String("nostrum"),
                            },
                        },
                    },
                    BulkConfig: &shared.PriceBulkConfig{
                        Tiers: []shared.PriceBulkConfigTiers{
                            shared.PriceBulkConfigTiers{
                                MaximumUnits: sdk.String("quisquam"),
                                UnitAmount: sdk.String("saepe"),
                            },
                            shared.PriceBulkConfigTiers{
                                MaximumUnits: sdk.String("ea"),
                                UnitAmount: sdk.String("impedit"),
                            },
                            shared.PriceBulkConfigTiers{
                                MaximumUnits: sdk.String("corporis"),
                                UnitAmount: sdk.String("veniam"),
                            },
                            shared.PriceBulkConfigTiers{
                                MaximumUnits: sdk.String("aliquid"),
                                UnitAmount: sdk.String("inventore"),
                            },
                        },
                    },
                    Cadence: shared.PriceCadenceAnnual.ToPointer(),
                    CreatedAt: types.MustTimeFromString("2022-03-24T01:04:28.722Z"),
                    Currency: "USD",
                    Discount: shared.Discount{
                        AmountDiscount: sdk.String("consectetur"),
                        AppliesToPriceIds: []string{
                            "aspernatur",
                            "minima",
                            "eaque",
                            "a",
                        },
                        DiscountType: shared.DiscountDiscountTypePercentage,
                        PercentageDiscount: sdk.Float64(0.15),
                        TrialAmountDiscount: sdk.String("libero"),
                        UsageDiscount: sdk.Float64(139.48),
                    },
                    FixedPriceQuantity: sdk.Float64(114.27),
                    ID: sdk.String("8c42e141-aac3-466c-8dd6-b14429074747"),
                    MatrixConfig: &shared.PriceMatrixConfig{
                        DefaultUnitAmount: sdk.String("esse"),
                        Dimensions: []string{
                            "fuga",
                            "reprehenderit",
                            "quidem",
                        },
                        MatrixValues: []shared.PriceMatrixConfigMatrixValues{
                            shared.PriceMatrixConfigMatrixValues{
                                DimensionValues: []string{
                                    "eum",
                                    "suscipit",
                                },
                                UnitAmount: sdk.String("assumenda"),
                            },
                            shared.PriceMatrixConfigMatrixValues{
                                DimensionValues: []string{
                                    "praesentium",
                                },
                                UnitAmount: sdk.String("quisquam"),
                            },
                            shared.PriceMatrixConfigMatrixValues{
                                DimensionValues: []string{
                                    "ipsa",
                                },
                                UnitAmount: sdk.String("id"),
                            },
                            shared.PriceMatrixConfigMatrixValues{
                                DimensionValues: []string{
                                    "neque",
                                    "quo",
                                    "illum",
                                },
                                UnitAmount: sdk.String("quo"),
                            },
                        },
                    },
                    Minimum: shared.MinimumAmount{
                        AppliesToPriceIds: []string{
                            "eius",
                            "eos",
                            "voluptas",
                        },
                        MinimumAmount: "ab",
                    },
                    ModelType: shared.PriceModelTypeBps.ToPointer(),
                    Name: sdk.String("Rhonda Toy"),
                    PackageConfig: &shared.PricePackageConfig{
                        PackageAmount: sdk.String("sequi"),
                        PackageSize: sdk.Float64(7791.92),
                    },
                    PlanPhaseOrder: 4598.56,
                    TieredBpsConfig: &shared.PriceTieredBpsConfig{
                        Tiers: []shared.PriceTieredBpsConfigTiers{
                            shared.PriceTieredBpsConfigTiers{
                                Bps: sdk.Float64(446.12),
                                MaximumAmount: sdk.String("distinctio"),
                                MinimumAmount: sdk.String("quod"),
                                PerUnitMaximum: sdk.String("dignissimos"),
                            },
                            shared.PriceTieredBpsConfigTiers{
                                Bps: sdk.Float64(769.56),
                                MaximumAmount: sdk.String("nihil"),
                                MinimumAmount: sdk.String("totam"),
                                PerUnitMaximum: sdk.String("accusamus"),
                            },
                            shared.PriceTieredBpsConfigTiers{
                                Bps: sdk.Float64(3068.1),
                                MaximumAmount: sdk.String("odio"),
                                MinimumAmount: sdk.String("occaecati"),
                                PerUnitMaximum: sdk.String("commodi"),
                            },
                            shared.PriceTieredBpsConfigTiers{
                                Bps: sdk.Float64(9594.34),
                                MaximumAmount: sdk.String("dolores"),
                                MinimumAmount: sdk.String("deserunt"),
                                PerUnitMaximum: sdk.String("molestiae"),
                            },
                        },
                    },
                    TieredConfig: &shared.PriceTieredConfig{
                        Tiers: []shared.PriceTieredConfigTiers{
                            shared.PriceTieredConfigTiers{
                                FirstUnit: sdk.String("porro"),
                                LastUnit: sdk.String("eum"),
                                UnitAmount: sdk.String("quas"),
                            },
                        },
                    },
                    UnitConfig: &shared.PriceUnitConfig{
                        UnitAmount: sdk.String("praesentium"),
                    },
                },
                shared.Price{
                    BillableMetric: &shared.PriceBillableMetric{
                        ID: sdk.String("282aa482-562f-4222-a981-7ee17cbe61e6"),
                    },
                    BpsConfig: &shared.PriceBpsConfig{
                        Bps: sdk.Float64(6900.25),
                        PerUnitMaximum: sdk.String("molestiae"),
                    },
                    BulkBpsConfig: &shared.PriceBulkBpsConfig{
                        Tiers: []shared.PriceBulkBpsConfigTiers{
                            shared.PriceBulkBpsConfigTiers{
                                Bps: sdk.Float64(5801.97),
                                MaximumAmount: sdk.String("minima"),
                                PerUnitMaximum: sdk.String("distinctio"),
                            },
                            shared.PriceBulkBpsConfigTiers{
                                Bps: sdk.Float64(7567.79),
                                MaximumAmount: sdk.String("sit"),
                                PerUnitMaximum: sdk.String("culpa"),
                            },
                            shared.PriceBulkBpsConfigTiers{
                                Bps: sdk.Float64(7313.98),
                                MaximumAmount: sdk.String("adipisci"),
                                PerUnitMaximum: sdk.String("cumque"),
                            },
                        },
                    },
                    BulkConfig: &shared.PriceBulkConfig{
                        Tiers: []shared.PriceBulkConfigTiers{
                            shared.PriceBulkConfigTiers{
                                MaximumUnits: sdk.String("consequatur"),
                                UnitAmount: sdk.String("minus"),
                            },
                        },
                    },
                    Cadence: shared.PriceCadenceAnnual.ToPointer(),
                    CreatedAt: types.MustTimeFromString("2022-04-21T00:17:42.407Z"),
                    Currency: "USD",
                    Discount: shared.Discount{
                        AmountDiscount: sdk.String("esse"),
                        AppliesToPriceIds: []string{
                            "provident",
                            "a",
                            "nulla",
                        },
                        DiscountType: shared.DiscountDiscountTypePercentage,
                        PercentageDiscount: sdk.Float64(0.15),
                        TrialAmountDiscount: sdk.String("quas"),
                        UsageDiscount: sdk.Float64(4572.23),
                    },
                    FixedPriceQuantity: sdk.Float64(974.68),
                    ID: sdk.String("f99dd2ef-d121-4aa6-b1e6-74bdb04f1575"),
                    MatrixConfig: &shared.PriceMatrixConfig{
                        DefaultUnitAmount: sdk.String("nisi"),
                        Dimensions: []string{
                            "voluptatum",
                        },
                        MatrixValues: []shared.PriceMatrixConfigMatrixValues{
                            shared.PriceMatrixConfigMatrixValues{
                                DimensionValues: []string{
                                    "ex",
                                    "deleniti",
                                    "itaque",
                                    "dolorum",
                                },
                                UnitAmount: sdk.String("architecto"),
                            },
                        },
                    },
                    Minimum: shared.MinimumAmount{
                        AppliesToPriceIds: []string{
                            "tenetur",
                            "quasi",
                            "at",
                        },
                        MinimumAmount: "et",
                    },
                    ModelType: shared.PriceModelTypePackage.ToPointer(),
                    Name: sdk.String("Joann Bogan"),
                    PackageConfig: &shared.PricePackageConfig{
                        PackageAmount: sdk.String("iste"),
                        PackageSize: sdk.Float64(8395.13),
                    },
                    PlanPhaseOrder: 330.74,
                    TieredBpsConfig: &shared.PriceTieredBpsConfig{
                        Tiers: []shared.PriceTieredBpsConfigTiers{
                            shared.PriceTieredBpsConfigTiers{
                                Bps: sdk.Float64(156.06),
                                MaximumAmount: sdk.String("laudantium"),
                                MinimumAmount: sdk.String("eum"),
                                PerUnitMaximum: sdk.String("mollitia"),
                            },
                            shared.PriceTieredBpsConfigTiers{
                                Bps: sdk.Float64(680.74),
                                MaximumAmount: sdk.String("corrupti"),
                                MinimumAmount: sdk.String("non"),
                                PerUnitMaximum: sdk.String("voluptatem"),
                            },
                            shared.PriceTieredBpsConfigTiers{
                                Bps: sdk.Float64(2211.61),
                                MaximumAmount: sdk.String("occaecati"),
                                MinimumAmount: sdk.String("numquam"),
                                PerUnitMaximum: sdk.String("impedit"),
                            },
                        },
                    },
                    TieredConfig: &shared.PriceTieredConfig{
                        Tiers: []shared.PriceTieredConfigTiers{
                            shared.PriceTieredConfigTiers{
                                FirstUnit: sdk.String("voluptas"),
                                LastUnit: sdk.String("aut"),
                                UnitAmount: sdk.String("dignissimos"),
                            },
                        },
                    },
                    UnitConfig: &shared.PriceUnitConfig{
                        UnitAmount: sdk.String("dicta"),
                    },
                },
            },
            Product: shared.PlanProduct{
                CreatedAt: types.MustTimeFromString("2021-02-22T18:20:18.826Z"),
                ID: "3f5f0642-dac7-4af5-95cc-413aa63aae8d",
                Name: "Dora Luettgen",
            },
            TrialConfig: &shared.PlanTrialConfig{
                TrialPeriod: sdk.Float64(8225.6),
                TrialPeriodUnit: shared.PlanTrialConfigTrialPeriodUnitDays,
            },
        },
        ExternalPlanID: "facilis",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Plan != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GetPlansExternalPlanIDRequest](../../models/operations/getplansexternalplanidrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.GetPlansExternalPlanIDResponse](../../models/operations/getplansexternalplanidresponse.md), error**


## List

This endpoint returns a list of all [plans](../reference/Orb-API.json/components/schemas/Plan) for an account in a list format. 

The list of plans is ordered starting from the most recently created plan. The response also includes [`pagination_metadata`](../api/pagination), which lets the caller retrieve the next page of results if they exist.


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
            APIKeyAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Plan.List(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.ListPlans200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ListPlansResponse](../../models/operations/listplansresponse.md), error**

