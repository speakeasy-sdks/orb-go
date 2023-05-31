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
            APIKeyAuth: "YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Plan.Fetch(ctx, operations.GetPlansPlanIDRequest{
        PlanID: "nobis",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Plan != nil {
        // handle response
    }
}
```

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
            APIKeyAuth: "YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Plan.GetByExternalID(ctx, operations.GetPlansExternalPlanIDRequest{
        Plan: &shared.Plan{
            BasePlan: &shared.PlanBasePlan{
                ExternalPlanID: sdk.String("dolores"),
                ID: sdk.String("58705320-2c73-4d5f-a9b9-0c28909b3fe4"),
                Name: sdk.String("Omar Leuschke"),
            },
            BasePlanID: sdk.String("nobis"),
            CreatedAt: types.MustTimeFromString("2021-01-26T22:08:21.462Z"),
            Currency: "quaerat",
            DefaultInvoiceMemo: sdk.String("quos"),
            Description: "aliquid",
            Discount: shared.Discount{
                AmountDiscount: sdk.String("dolorem"),
                AppliesToPriceIds: []string{
                    "dolor",
                },
                DiscountType: shared.DiscountDiscountTypePercentage,
                PercentageDiscount: sdk.Float64(0.15),
                TrialAmountDiscount: sdk.String("qui"),
                UsageDiscount: sdk.Float64(2187.49),
            },
            ExternalPlanID: sdk.String("hic"),
            ID: "9b77f3a4-1006-474e-bf69-280d1ba77a89",
            InvoicingCurrency: "necessitatibus",
            Minimum: shared.MinimumAmount{
                AppliesToPriceIds: []string{
                    "asperiores",
                    "nihil",
                    "ipsum",
                },
                MinimumAmount: "voluptate",
            },
            Name: "Elbert Gislason I",
            NetTerms: sdk.Int64(758379),
            PlanPhases: []shared.PlanPhase{
                shared.PlanPhase{
                    Description: sdk.String("ad"),
                    Discount: shared.Discount{
                        AmountDiscount: sdk.String("saepe"),
                        AppliesToPriceIds: []string{
                            "deserunt",
                            "provident",
                        },
                        DiscountType: shared.DiscountDiscountTypePercentage,
                        PercentageDiscount: sdk.Float64(0.15),
                        TrialAmountDiscount: sdk.String("minima"),
                        UsageDiscount: sdk.Float64(8310.49),
                    },
                    Duration: sdk.Int64(519711),
                    DurationUnit: shared.PlanPhaseDurationUnitQuarterly,
                    Minimum: shared.MinimumAmount{
                        AppliesToPriceIds: []string{
                            "at",
                        },
                        MinimumAmount: "quaerat",
                    },
                    Name: sdk.String("Gina Schmeler"),
                    Order: sdk.Int64(679880),
                },
                shared.PlanPhase{
                    Description: sdk.String("a"),
                    Discount: shared.Discount{
                        AmountDiscount: sdk.String("esse"),
                        AppliesToPriceIds: []string{
                            "iusto",
                            "ipsum",
                            "quisquam",
                        },
                        DiscountType: shared.DiscountDiscountTypePercentage,
                        PercentageDiscount: sdk.Float64(0.15),
                        TrialAmountDiscount: sdk.String("tenetur"),
                        UsageDiscount: sdk.Float64(2294.42),
                    },
                    Duration: sdk.Int64(730856),
                    DurationUnit: shared.PlanPhaseDurationUnitAnnual,
                    Minimum: shared.MinimumAmount{
                        AppliesToPriceIds: []string{
                            "enim",
                            "dolorem",
                        },
                        MinimumAmount: "sapiente",
                    },
                    Name: sdk.String("Christian Balistreri"),
                    Order: sdk.Int64(153694),
                },
                shared.PlanPhase{
                    Description: sdk.String("vel"),
                    Discount: shared.Discount{
                        AmountDiscount: sdk.String("libero"),
                        AppliesToPriceIds: []string{
                            "deserunt",
                            "quam",
                        },
                        DiscountType: shared.DiscountDiscountTypePercentage,
                        PercentageDiscount: sdk.Float64(0.15),
                        TrialAmountDiscount: sdk.String("ipsum"),
                        UsageDiscount: sdk.Float64(2776.28),
                    },
                    Duration: sdk.Int64(186458),
                    DurationUnit: shared.PlanPhaseDurationUnitQuarterly,
                    Minimum: shared.MinimumAmount{
                        AppliesToPriceIds: []string{
                            "pariatur",
                            "soluta",
                            "dicta",
                            "laborum",
                        },
                        MinimumAmount: "totam",
                    },
                    Name: sdk.String("Kelly Daniel"),
                    Order: sdk.Int64(396060),
                },
                shared.PlanPhase{
                    Description: sdk.String("quam"),
                    Discount: shared.Discount{
                        AmountDiscount: sdk.String("molestias"),
                        AppliesToPriceIds: []string{
                            "qui",
                            "neque",
                            "fugit",
                            "magni",
                        },
                        DiscountType: shared.DiscountDiscountTypePercentage,
                        PercentageDiscount: sdk.Float64(0.15),
                        TrialAmountDiscount: sdk.String("odio"),
                        UsageDiscount: sdk.Float64(1248.33),
                    },
                    Duration: sdk.Int64(355613),
                    DurationUnit: shared.PlanPhaseDurationUnitAnnual,
                    Minimum: shared.MinimumAmount{
                        AppliesToPriceIds: []string{
                            "voluptatem",
                            "cumque",
                            "soluta",
                            "nobis",
                        },
                        MinimumAmount: "et",
                    },
                    Name: sdk.String("Dale Boehm"),
                    Order: sdk.Int64(731694),
                },
            },
            Prices: []shared.Price{
                shared.Price{
                    BillableMetric: &shared.PriceBillableMetric{
                        ID: sdk.String("0f3443a1-108e-40ad-8f4b-921879fce953"),
                    },
                    BpsConfig: &shared.PriceBpsConfig{
                        Bps: sdk.Float64(9615.71),
                        PerUnitMaximum: sdk.String("voluptate"),
                    },
                    BulkBpsConfig: &shared.PriceBulkBpsConfig{
                        Tiers: []shared.PriceBulkBpsConfigTiers{
                            shared.PriceBulkBpsConfigTiers{
                                Bps: sdk.Float64(8788.7),
                                MaximumAmount: sdk.String("tenetur"),
                                PerUnitMaximum: sdk.String("dignissimos"),
                            },
                        },
                    },
                    BulkConfig: &shared.PriceBulkConfig{
                        Tiers: []shared.PriceBulkConfigTiers{
                            shared.PriceBulkConfigTiers{
                                MaximumUnits: sdk.String("distinctio"),
                                UnitAmount: sdk.String("quod"),
                            },
                            shared.PriceBulkConfigTiers{
                                MaximumUnits: sdk.String("odio"),
                                UnitAmount: sdk.String("similique"),
                            },
                            shared.PriceBulkConfigTiers{
                                MaximumUnits: sdk.String("facilis"),
                                UnitAmount: sdk.String("vero"),
                            },
                            shared.PriceBulkConfigTiers{
                                MaximumUnits: sdk.String("ducimus"),
                                UnitAmount: sdk.String("dolore"),
                            },
                        },
                    },
                    Cadence: shared.PriceCadenceQuarterly.ToPointer(),
                    CreatedAt: types.MustTimeFromString("2022-06-02T04:41:10.492Z"),
                    Currency: "USD",
                    Discount: shared.Discount{
                        AmountDiscount: sdk.String("natus"),
                        AppliesToPriceIds: []string{
                            "aut",
                            "voluptatibus",
                            "exercitationem",
                            "nulla",
                        },
                        DiscountType: shared.DiscountDiscountTypePercentage,
                        PercentageDiscount: sdk.Float64(0.15),
                        TrialAmountDiscount: sdk.String("fugit"),
                        UsageDiscount: sdk.Float64(7804.27),
                    },
                    FixedPriceQuantity: sdk.Float64(9818.3),
                    ID: sdk.String("f7c70a45-626d-4436-813f-16d9f5fce6c5"),
                    MatrixConfig: &shared.PriceMatrixConfig{
                        DefaultUnitAmount: sdk.String("veniam"),
                        Dimensions: []string{
                            "inventore",
                            "magnam",
                        },
                        MatrixValues: []shared.PriceMatrixConfigMatrixValues{
                            shared.PriceMatrixConfigMatrixValues{
                                DimensionValues: []string{
                                    "consectetur",
                                    "recusandae",
                                    "aspernatur",
                                    "minima",
                                },
                                UnitAmount: sdk.String("eaque"),
                            },
                            shared.PriceMatrixConfigMatrixValues{
                                DimensionValues: []string{
                                    "libero",
                                    "aut",
                                    "aut",
                                    "deleniti",
                                },
                                UnitAmount: sdk.String("impedit"),
                            },
                        },
                    },
                    Minimum: shared.MinimumAmount{
                        AppliesToPriceIds: []string{
                            "fugit",
                            "accusamus",
                        },
                        MinimumAmount: "inventore",
                    },
                    ModelType: shared.PriceModelTypeBulk.ToPointer(),
                    Name: sdk.String("Sonya Osinski"),
                    PackageConfig: &shared.PricePackageConfig{
                        PackageAmount: sdk.String("eum"),
                        PackageSize: sdk.Float64(4205.39),
                    },
                    PlanPhaseOrder: 7521.35,
                    TieredBpsConfig: &shared.PriceTieredBpsConfig{
                        Tiers: []shared.PriceTieredBpsConfigTiers{
                            shared.PriceTieredBpsConfigTiers{
                                Bps: sdk.Float64(8296.03),
                                MaximumAmount: sdk.String("nulla"),
                                MinimumAmount: sdk.String("voluptas"),
                                PerUnitMaximum: sdk.String("libero"),
                            },
                            shared.PriceTieredBpsConfigTiers{
                                Bps: sdk.Float64(965.49),
                                MaximumAmount: sdk.String("tempora"),
                                MinimumAmount: sdk.String("numquam"),
                                PerUnitMaximum: sdk.String("explicabo"),
                            },
                            shared.PriceTieredBpsConfigTiers{
                                Bps: sdk.Float64(5919.35),
                                MaximumAmount: sdk.String("ipsa"),
                                MinimumAmount: sdk.String("molestiae"),
                                PerUnitMaximum: sdk.String("magnam"),
                            },
                        },
                    },
                    TieredConfig: &shared.PriceTieredConfig{
                        Tiers: []shared.PriceTieredConfigTiers{
                            shared.PriceTieredConfigTiers{
                                FirstUnit: sdk.String("eius"),
                                LastUnit: sdk.String("esse"),
                                UnitAmount: sdk.String("esse"),
                            },
                            shared.PriceTieredConfigTiers{
                                FirstUnit: sdk.String("rem"),
                                LastUnit: sdk.String("fuga"),
                                UnitAmount: sdk.String("reprehenderit"),
                            },
                        },
                    },
                    UnitConfig: &shared.PriceUnitConfig{
                        UnitAmount: sdk.String("quidem"),
                    },
                },
                shared.Price{
                    BillableMetric: &shared.PriceBillableMetric{
                        ID: sdk.String("d466d28c-10ab-43cd-8a42-51904e523c7e"),
                    },
                    BpsConfig: &shared.PriceBpsConfig{
                        Bps: sdk.Float64(446.12),
                        PerUnitMaximum: sdk.String("distinctio"),
                    },
                    BulkBpsConfig: &shared.PriceBulkBpsConfig{
                        Tiers: []shared.PriceBulkBpsConfigTiers{
                            shared.PriceBulkBpsConfigTiers{
                                Bps: sdk.Float64(4908.19),
                                MaximumAmount: sdk.String("inventore"),
                                PerUnitMaximum: sdk.String("nihil"),
                            },
                            shared.PriceBulkBpsConfigTiers{
                                Bps: sdk.Float64(5188.35),
                                MaximumAmount: sdk.String("accusamus"),
                                PerUnitMaximum: sdk.String("aliquam"),
                            },
                            shared.PriceBulkBpsConfigTiers{
                                Bps: sdk.Float64(4884.1),
                                MaximumAmount: sdk.String("occaecati"),
                                PerUnitMaximum: sdk.String("commodi"),
                            },
                            shared.PriceBulkBpsConfigTiers{
                                Bps: sdk.Float64(9594.34),
                                MaximumAmount: sdk.String("dolores"),
                                PerUnitMaximum: sdk.String("deserunt"),
                            },
                        },
                    },
                    BulkConfig: &shared.PriceBulkConfig{
                        Tiers: []shared.PriceBulkConfigTiers{
                            shared.PriceBulkConfigTiers{
                                MaximumUnits: sdk.String("accusantium"),
                                UnitAmount: sdk.String("porro"),
                            },
                            shared.PriceBulkConfigTiers{
                                MaximumUnits: sdk.String("eum"),
                                UnitAmount: sdk.String("quas"),
                            },
                        },
                    },
                    Cadence: shared.PriceCadenceMonthly.ToPointer(),
                    CreatedAt: types.MustTimeFromString("2022-06-19T07:05:05.913Z"),
                    Currency: "USD",
                    Discount: shared.Discount{
                        AmountDiscount: sdk.String("fugit"),
                        AppliesToPriceIds: []string{
                            "mollitia",
                            "incidunt",
                            "atque",
                        },
                        DiscountType: shared.DiscountDiscountTypePercentage,
                        PercentageDiscount: sdk.Float64(0.15),
                        TrialAmountDiscount: sdk.String("explicabo"),
                        UsageDiscount: sdk.Float64(3256.85),
                    },
                    FixedPriceQuantity: sdk.Float64(3926.76),
                    ID: sdk.String("2f222e98-17ee-417c-be61-e6b7b95bc0ab"),
                    MatrixConfig: &shared.PriceMatrixConfig{
                        DefaultUnitAmount: sdk.String("adipisci"),
                        Dimensions: []string{
                            "consequuntur",
                            "consequatur",
                            "minus",
                            "quaerat",
                        },
                        MatrixValues: []shared.PriceMatrixConfigMatrixValues{
                            shared.PriceMatrixConfigMatrixValues{
                                DimensionValues: []string{
                                    "esse",
                                },
                                UnitAmount: sdk.String("blanditiis"),
                            },
                            shared.PriceMatrixConfigMatrixValues{
                                DimensionValues: []string{
                                    "a",
                                    "nulla",
                                    "quas",
                                },
                                UnitAmount: sdk.String("esse"),
                            },
                            shared.PriceMatrixConfigMatrixValues{
                                DimensionValues: []string{
                                    "a",
                                },
                                UnitAmount: sdk.String("error"),
                            },
                            shared.PriceMatrixConfigMatrixValues{
                                DimensionValues: []string{
                                    "pariatur",
                                    "possimus",
                                    "quia",
                                },
                                UnitAmount: sdk.String("eveniet"),
                            },
                        },
                    },
                    Minimum: shared.MinimumAmount{
                        AppliesToPriceIds: []string{
                            "facere",
                            "veritatis",
                            "consequuntur",
                            "quasi",
                        },
                        MinimumAmount: "similique",
                    },
                    ModelType: shared.PriceModelTypeTieredBps.ToPointer(),
                    Name: sdk.String("Mandy Berge"),
                    PackageConfig: &shared.PricePackageConfig{
                        PackageAmount: sdk.String("in"),
                        PackageSize: sdk.Float64(2586.84),
                    },
                    PlanPhaseOrder: 7276.97,
                    TieredBpsConfig: &shared.PriceTieredBpsConfig{
                        Tiers: []shared.PriceTieredBpsConfigTiers{
                            shared.PriceTieredBpsConfigTiers{
                                Bps: sdk.Float64(7422.38),
                                MaximumAmount: sdk.String("accusantium"),
                                MinimumAmount: sdk.String("aliquam"),
                                PerUnitMaximum: sdk.String("sapiente"),
                            },
                            shared.PriceTieredBpsConfigTiers{
                                Bps: sdk.Float64(1197.71),
                                MaximumAmount: sdk.String("ullam"),
                                MinimumAmount: sdk.String("reprehenderit"),
                                PerUnitMaximum: sdk.String("ullam"),
                            },
                            shared.PriceTieredBpsConfigTiers{
                                Bps: sdk.Float64(3917.74),
                                MaximumAmount: sdk.String("aut"),
                                MinimumAmount: sdk.String("voluptatum"),
                                PerUnitMaximum: sdk.String("qui"),
                            },
                            shared.PriceTieredBpsConfigTiers{
                                Bps: sdk.Float64(8453.58),
                                MaximumAmount: sdk.String("ex"),
                                MinimumAmount: sdk.String("deleniti"),
                                PerUnitMaximum: sdk.String("itaque"),
                            },
                        },
                    },
                    TieredConfig: &shared.PriceTieredConfig{
                        Tiers: []shared.PriceTieredConfigTiers{
                            shared.PriceTieredConfigTiers{
                                FirstUnit: sdk.String("architecto"),
                                LastUnit: sdk.String("omnis"),
                                UnitAmount: sdk.String("tenetur"),
                            },
                            shared.PriceTieredConfigTiers{
                                FirstUnit: sdk.String("quasi"),
                                LastUnit: sdk.String("at"),
                                UnitAmount: sdk.String("et"),
                            },
                            shared.PriceTieredConfigTiers{
                                FirstUnit: sdk.String("voluptate"),
                                LastUnit: sdk.String("ipsa"),
                                UnitAmount: sdk.String("minima"),
                            },
                        },
                    },
                    UnitConfig: &shared.PriceUnitConfig{
                        UnitAmount: sdk.String("veritatis"),
                    },
                },
                shared.Price{
                    BillableMetric: &shared.PriceBillableMetric{
                        ID: sdk.String("339d0808-6a18-4403-94c2-6071f93f5f06"),
                    },
                    BpsConfig: &shared.PriceBpsConfig{
                        Bps: sdk.Float64(3100.67),
                        PerUnitMaximum: sdk.String("consequuntur"),
                    },
                    BulkBpsConfig: &shared.PriceBulkBpsConfig{
                        Tiers: []shared.PriceBulkBpsConfigTiers{
                            shared.PriceBulkBpsConfigTiers{
                                Bps: sdk.Float64(6387.62),
                                MaximumAmount: sdk.String("maxime"),
                                PerUnitMaximum: sdk.String("dignissimos"),
                            },
                            shared.PriceBulkBpsConfigTiers{
                                Bps: sdk.Float64(6400.24),
                                MaximumAmount: sdk.String("asperiores"),
                                PerUnitMaximum: sdk.String("nemo"),
                            },
                            shared.PriceBulkBpsConfigTiers{
                                Bps: sdk.Float64(653.04),
                                MaximumAmount: sdk.String("quaerat"),
                                PerUnitMaximum: sdk.String("porro"),
                            },
                            shared.PriceBulkBpsConfigTiers{
                                Bps: sdk.Float64(8018.36),
                                MaximumAmount: sdk.String("labore"),
                                PerUnitMaximum: sdk.String("ab"),
                            },
                        },
                    },
                    BulkConfig: &shared.PriceBulkConfig{
                        Tiers: []shared.PriceBulkConfigTiers{
                            shared.PriceBulkConfigTiers{
                                MaximumUnits: sdk.String("fuga"),
                                UnitAmount: sdk.String("id"),
                            },
                        },
                    },
                    Cadence: shared.PriceCadenceMonthly.ToPointer(),
                    CreatedAt: types.MustTimeFromString("2022-05-14T14:45:53.115Z"),
                    Currency: "USD",
                    Discount: shared.Discount{
                        AmountDiscount: sdk.String("est"),
                        AppliesToPriceIds: []string{
                            "totam",
                            "fugiat",
                            "vel",
                            "ducimus",
                        },
                        DiscountType: shared.DiscountDiscountTypePercentage,
                        PercentageDiscount: sdk.Float64(0.15),
                        TrialAmountDiscount: sdk.String("quos"),
                        UsageDiscount: sdk.Float64(4278.34),
                    },
                    FixedPriceQuantity: sdk.Float64(2870.51),
                    ID: sdk.String("dbb675fd-5e60-4b37-9ed4-f6fbee41f333"),
                    MatrixConfig: &shared.PriceMatrixConfig{
                        DefaultUnitAmount: sdk.String("beatae"),
                        Dimensions: []string{
                            "a",
                            "debitis",
                        },
                        MatrixValues: []shared.PriceMatrixConfigMatrixValues{
                            shared.PriceMatrixConfigMatrixValues{
                                DimensionValues: []string{
                                    "harum",
                                    "laboriosam",
                                },
                                UnitAmount: sdk.String("ipsa"),
                            },
                        },
                    },
                    Minimum: shared.MinimumAmount{
                        AppliesToPriceIds: []string{
                            "libero",
                            "vitae",
                            "accusamus",
                            "similique",
                        },
                        MinimumAmount: "tempora",
                    },
                    ModelType: shared.PriceModelTypeTiered.ToPointer(),
                    Name: sdk.String("Melanie Hirthe"),
                    PackageConfig: &shared.PricePackageConfig{
                        PackageAmount: sdk.String("dolorum"),
                        PackageSize: sdk.Float64(2378.07),
                    },
                    PlanPhaseOrder: 7955.35,
                    TieredBpsConfig: &shared.PriceTieredBpsConfig{
                        Tiers: []shared.PriceTieredBpsConfigTiers{
                            shared.PriceTieredBpsConfigTiers{
                                Bps: sdk.Float64(5039.34),
                                MaximumAmount: sdk.String("in"),
                                MinimumAmount: sdk.String("dolore"),
                                PerUnitMaximum: sdk.String("aliquam"),
                            },
                        },
                    },
                    TieredConfig: &shared.PriceTieredConfig{
                        Tiers: []shared.PriceTieredConfigTiers{
                            shared.PriceTieredConfigTiers{
                                FirstUnit: sdk.String("temporibus"),
                                LastUnit: sdk.String("ullam"),
                                UnitAmount: sdk.String("adipisci"),
                            },
                            shared.PriceTieredConfigTiers{
                                FirstUnit: sdk.String("cum"),
                                LastUnit: sdk.String("blanditiis"),
                                UnitAmount: sdk.String("quas"),
                            },
                            shared.PriceTieredConfigTiers{
                                FirstUnit: sdk.String("hic"),
                                LastUnit: sdk.String("nesciunt"),
                                UnitAmount: sdk.String("culpa"),
                            },
                            shared.PriceTieredConfigTiers{
                                FirstUnit: sdk.String("corrupti"),
                                LastUnit: sdk.String("pariatur"),
                                UnitAmount: sdk.String("totam"),
                            },
                        },
                    },
                    UnitConfig: &shared.PriceUnitConfig{
                        UnitAmount: sdk.String("hic"),
                    },
                },
            },
            Product: shared.PlanProduct{
                CreatedAt: types.MustTimeFromString("2022-04-01T23:17:58.998Z"),
                ID: "0b2f2fb7-b194-4a27-ab26-916fe1f08f42",
                Name: "Herbert Treutel",
            },
            TrialConfig: &shared.PlanTrialConfig{
                TrialPeriod: sdk.Float64(5799.12),
                TrialPeriodUnit: shared.PlanTrialConfigTrialPeriodUnitDays,
            },
        },
        ExternalPlanID: "quos",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Plan != nil {
        // handle response
    }
}
```

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
            APIKeyAuth: "YOUR_BEARER_TOKEN_HERE",
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
