# Plan

## Overview

Actions related to plan management.

### Available Operations

* [Get](#get) - Retrieve a plan
* [GetByExternalID](#getbyexternalid) - Retrieve a plan by external plan ID
* [List](#list) - List plans

## Get

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
            BearerAuth: "YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Plan.Get(ctx, operations.GetPlansPlanIDRequest{
        PlanID: "maxime",
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
            BearerAuth: "YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Plan.GetByExternalID(ctx, operations.GetPlansExternalPlanIDRequest{
        Plan: &shared.Plan{
            BasePlanID: sdk.String("deleniti"),
            CreatedAt: types.MustTimeFromString("2022-02-08T00:19:59.821Z"),
            Currency: "architecto",
            Description: "architecto",
            Discount: map[string]interface{}{
                "ullam": "expedita",
                "nihil": "repellat",
                "quibusdam": "sed",
                "saepe": "pariatur",
            },
            ExternalPlanID: sdk.String("accusantium"),
            ID: "28921cdd-c692-4601-bb57-6b0d5f0d30c5",
            InvoicingCurrency: "hic",
            Minimum: map[string]interface{}{
                "nobis": "dolores",
                "quis": "totam",
                "dignissimos": "eaque",
            },
            Name: "Mr. Robin Daugherty",
            PlanPhases: []shared.PlanPhase{
                shared.PlanPhase{
                    Description: sdk.String("dolor"),
                    Discount: map[string]interface{}{
                        "nostrum": "hic",
                        "recusandae": "omnis",
                        "facilis": "perspiciatis",
                        "voluptatem": "porro",
                    },
                    Duration: sdk.Int64(164694),
                    DurationUnit: shared.PlanPhaseDurationUnitQuarterly,
                    Minimum: map[string]interface{}{
                        "eaque": "occaecati",
                        "rerum": "adipisci",
                        "asperiores": "earum",
                    },
                    Name: sdk.String("Sheryl Parisian"),
                    Order: sdk.Int64(589910),
                },
                shared.PlanPhase{
                    Description: sdk.String("nobis"),
                    Discount: map[string]interface{}{
                        "delectus": "quaerat",
                        "quos": "aliquid",
                        "dolorem": "dolorem",
                    },
                    Duration: sdk.Int64(222443),
                    DurationUnit: shared.PlanPhaseDurationUnitMonthly,
                    Minimum: map[string]interface{}{
                        "hic": "excepturi",
                    },
                    Name: sdk.String("Adrian Kuhn"),
                    Order: sdk.Int64(680545),
                },
            },
            Prices: []shared.Price{
                shared.Price{
                    BillableMetric: &shared.PriceBillableMetric{
                        ID: sdk.String("100674eb-f692-480d-9ba7-7a89ebf737ae"),
                    },
                    BpsConfig: &shared.PriceBpsConfig{
                        Bps: sdk.Float64(2633.22),
                        PerUnitMaximum: sdk.String("aspernatur"),
                    },
                    BulkBpsConfig: &shared.PriceBulkBpsConfig{
                        Tiers: []shared.PriceBulkBpsConfigTiers{
                            shared.PriceBulkBpsConfigTiers{
                                Bps: sdk.Float64(2292.19),
                                MaximumAmount: sdk.String("optio"),
                                PerUnitMaximum: sdk.String("accusamus"),
                            },
                        },
                    },
                    BulkConfig: &shared.PriceBulkConfig{
                        Tiers: []shared.PriceBulkConfigTiers{
                            shared.PriceBulkConfigTiers{
                                MaximumUnits: sdk.String("saepe"),
                                UnitAmount: sdk.String("suscipit"),
                            },
                            shared.PriceBulkConfigTiers{
                                MaximumUnits: sdk.String("deserunt"),
                                UnitAmount: sdk.String("provident"),
                            },
                        },
                    },
                    Cadence: shared.PriceCadenceAnnual.ToPointer(),
                    CreatedAt: types.MustTimeFromString("2021-06-10T21:59:23.045Z"),
                    Currency: "USD",
                    Discount: map[string]interface{}{
                        "alias": "at",
                        "quaerat": "tempora",
                        "vel": "quod",
                    },
                    FixedPriceQuantity: sdk.Float64(8853.38),
                    ID: sdk.String("2af7a73c-f3be-4453-b870-b326b5a73429"),
                    MatrixConfig: &shared.PriceMatrixConfig{
                        DefaultUnitAmount: sdk.String("maxime"),
                        Dimensions: []string{
                            "soluta",
                            "dicta",
                            "laborum",
                            "totam",
                        },
                        MatrixValues: []shared.PriceMatrixConfigMatrixValues{
                            shared.PriceMatrixConfigMatrixValues{
                                DimensionValues: []string{
                                    "dolores",
                                },
                                UnitAmount: sdk.String("distinctio"),
                            },
                            shared.PriceMatrixConfigMatrixValues{
                                DimensionValues: []string{
                                    "aliquid",
                                    "quam",
                                    "molestias",
                                },
                                UnitAmount: sdk.String("temporibus"),
                            },
                        },
                    },
                    Minimum: map[string]interface{}{
                        "neque": "fugit",
                    },
                    ModelType: shared.PriceModelTypeTiered.ToPointer(),
                    Name: sdk.String("Ashley Hermiston"),
                    PackageConfig: &shared.PricePackageConfig{
                        PackageAmount: sdk.String("voluptatem"),
                        PackageSize: sdk.Float64(7653.26),
                    },
                    PlanPhaseOrder: 7469.94,
                    TieredBpsConfig: &shared.PriceTieredBpsConfig{
                        Tiers: []shared.PriceTieredBpsConfigTiers{
                            shared.PriceTieredBpsConfigTiers{
                                Bps: sdk.Float64(925.96),
                                MaximumAmount: sdk.String("saepe"),
                                MinimumAmount: sdk.String("ipsum"),
                                PerUnitMaximum: sdk.String("veritatis"),
                            },
                            shared.PriceTieredBpsConfigTiers{
                                Bps: sdk.Float64(7492.55),
                                MaximumAmount: sdk.String("quos"),
                                MinimumAmount: sdk.String("tempore"),
                                PerUnitMaximum: sdk.String("cupiditate"),
                            },
                            shared.PriceTieredBpsConfigTiers{
                                Bps: sdk.Float64(456.14),
                                MaximumAmount: sdk.String("delectus"),
                                MinimumAmount: sdk.String("dolorem"),
                                PerUnitMaximum: sdk.String("dolore"),
                            },
                        },
                    },
                    TieredConfig: &shared.PriceTieredConfig{
                        Tiers: []shared.PriceTieredConfigTiers{
                            shared.PriceTieredConfigTiers{
                                FirstUnit: sdk.String("adipisci"),
                                LastUnit: sdk.String("dolorum"),
                                UnitAmount: sdk.String("architecto"),
                            },
                            shared.PriceTieredConfigTiers{
                                FirstUnit: sdk.String("quae"),
                                LastUnit: sdk.String("aut"),
                                UnitAmount: sdk.String("quas"),
                            },
                        },
                    },
                    UnitConfig: &shared.PriceUnitConfig{
                        UnitAmount: sdk.String("itaque"),
                    },
                },
                shared.Price{
                    BillableMetric: &shared.PriceBillableMetric{
                        ID: sdk.String("0adcf4b9-2187-49fc-a953-f73ef7fbc7ab"),
                    },
                    BpsConfig: &shared.PriceBpsConfig{
                        Bps: sdk.Float64(8742.88),
                        PerUnitMaximum: sdk.String("ducimus"),
                    },
                    BulkBpsConfig: &shared.PriceBulkBpsConfig{
                        Tiers: []shared.PriceBulkBpsConfigTiers{
                            shared.PriceBulkBpsConfigTiers{
                                Bps: sdk.Float64(8445.5),
                                MaximumAmount: sdk.String("illum"),
                                PerUnitMaximum: sdk.String("sequi"),
                            },
                            shared.PriceBulkBpsConfigTiers{
                                Bps: sdk.Float64(6178.77),
                                MaximumAmount: sdk.String("impedit"),
                                PerUnitMaximum: sdk.String("aut"),
                            },
                        },
                    },
                    BulkConfig: &shared.PriceBulkConfig{
                        Tiers: []shared.PriceBulkConfigTiers{
                            shared.PriceBulkConfigTiers{
                                MaximumUnits: sdk.String("exercitationem"),
                                UnitAmount: sdk.String("nulla"),
                            },
                            shared.PriceBulkConfigTiers{
                                MaximumUnits: sdk.String("fugit"),
                                UnitAmount: sdk.String("porro"),
                            },
                            shared.PriceBulkConfigTiers{
                                MaximumUnits: sdk.String("maiores"),
                                UnitAmount: sdk.String("doloribus"),
                            },
                            shared.PriceBulkConfigTiers{
                                MaximumUnits: sdk.String("iusto"),
                                UnitAmount: sdk.String("eligendi"),
                            },
                        },
                    },
                    Cadence: shared.PriceCadenceMonthly.ToPointer(),
                    CreatedAt: types.MustTimeFromString("2022-05-12T14:13:14.359Z"),
                    Currency: "USD",
                    Discount: map[string]interface{}{
                        "ipsam": "ea",
                        "aspernatur": "vel",
                    },
                    FixedPriceQuantity: sdk.Float64(8221.18),
                    ID: sdk.String("436813f1-6d9f-45fc-a6c5-56146c3e250f"),
                    MatrixConfig: &shared.PriceMatrixConfig{
                        DefaultUnitAmount: sdk.String("libero"),
                        Dimensions: []string{
                            "aut",
                        },
                        MatrixValues: []shared.PriceMatrixConfigMatrixValues{
                            shared.PriceMatrixConfigMatrixValues{
                                DimensionValues: []string{
                                    "aliquam",
                                    "fugit",
                                    "accusamus",
                                    "inventore",
                                },
                                UnitAmount: sdk.String("non"),
                            },
                            shared.PriceMatrixConfigMatrixValues{
                                DimensionValues: []string{
                                    "dolorum",
                                },
                                UnitAmount: sdk.String("laborum"),
                            },
                            shared.PriceMatrixConfigMatrixValues{
                                DimensionValues: []string{
                                    "velit",
                                    "eum",
                                    "autem",
                                    "nobis",
                                },
                                UnitAmount: sdk.String("quas"),
                            },
                        },
                    },
                    Minimum: map[string]interface{}{
                        "nulla": "voluptas",
                        "libero": "quasi",
                        "tempora": "numquam",
                        "explicabo": "provident",
                    },
                    ModelType: shared.PriceModelTypeUnit.ToPointer(),
                    Name: sdk.String("Megan Kuhlman"),
                    PackageConfig: &shared.PricePackageConfig{
                        PackageAmount: sdk.String("esse"),
                        PackageSize: sdk.Float64(5245.93),
                    },
                    PlanPhaseOrder: 6832.82,
                    TieredBpsConfig: &shared.PriceTieredBpsConfig{
                        Tiers: []shared.PriceTieredBpsConfigTiers{
                            shared.PriceTieredBpsConfigTiers{
                                Bps: sdk.Float64(6956.26),
                                MaximumAmount: sdk.String("fugiat"),
                                MinimumAmount: sdk.String("ut"),
                                PerUnitMaximum: sdk.String("eum"),
                            },
                            shared.PriceTieredBpsConfigTiers{
                                Bps: sdk.Float64(3799.27),
                                MaximumAmount: sdk.String("assumenda"),
                                MinimumAmount: sdk.String("eos"),
                                PerUnitMaximum: sdk.String("praesentium"),
                            },
                        },
                    },
                    TieredConfig: &shared.PriceTieredConfig{
                        Tiers: []shared.PriceTieredConfigTiers{
                            shared.PriceTieredConfigTiers{
                                FirstUnit: sdk.String("veritatis"),
                                LastUnit: sdk.String("ipsa"),
                                UnitAmount: sdk.String("id"),
                            },
                            shared.PriceTieredConfigTiers{
                                FirstUnit: sdk.String("quidem"),
                                LastUnit: sdk.String("neque"),
                                UnitAmount: sdk.String("quo"),
                            },
                            shared.PriceTieredConfigTiers{
                                FirstUnit: sdk.String("illum"),
                                LastUnit: sdk.String("quo"),
                                UnitAmount: sdk.String("fuga"),
                            },
                            shared.PriceTieredConfigTiers{
                                FirstUnit: sdk.String("eius"),
                                LastUnit: sdk.String("eos"),
                                UnitAmount: sdk.String("voluptas"),
                            },
                        },
                    },
                    UnitConfig: &shared.PriceUnitConfig{
                        UnitAmount: sdk.String("ab"),
                    },
                },
            },
            Product: shared.PlanProduct{
                CreatedAt: types.MustTimeFromString("2022-12-24T22:15:46.522Z"),
                ID: "4e523c7e-0bc7-4178-a479-6f2a70c68828",
                Name: "Lee O'Conner",
            },
            TrialConfig: &shared.PlanTrialConfig{
                TrialPeriod: sdk.Float64(1288.6),
                TrialPeriodUnit: shared.PlanTrialConfigTrialPeriodUnitDays,
            },
        },
        ExternalPlanID: "minima",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## List

This endpoint returns a list of all [plans](../reference/Orb-API.json/components/schemas/Plan) for an account in a list format. 

The list of plans is ordered starting from the most recently created plan. The response also includes [`pagination_metadata`](../reference/Orb-API.json/components/schemas/Pagination-metadata), which lets the caller retrieve the next page of results if they exist. More information about pagination can be found in the [Pagination-metadata schema](../reference/Orb-API.json/components/schemas/Pagination-metadata).


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
            BearerAuth: "YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Plan.List(ctx, operations.ListPlansRequestBody{
        Data: []shared.Plan{
            shared.Plan{
                BasePlanID: sdk.String("fugit"),
                CreatedAt: types.MustTimeFromString("2022-07-09T22:36:51.892Z"),
                Currency: "ratione",
                Description: "explicabo",
                Discount: map[string]interface{}{
                    "occaecati": "atque",
                    "et": "esse",
                    "eveniet": "accusamus",
                    "veritatis": "esse",
                },
                ExternalPlanID: sdk.String("quod"),
                ID: "be61e6b7-b95b-4c0a-b3c2-0c4f3789fd87",
                InvoicingCurrency: "quasi",
                Minimum: map[string]interface{}{
                    "error": "sint",
                    "pariatur": "possimus",
                    "quia": "eveniet",
                    "asperiores": "facere",
                },
                Name: "Marilyn Botsford",
                PlanPhases: []shared.PlanPhase{
                    shared.PlanPhase{
                        Description: sdk.String("tenetur"),
                        Discount: map[string]interface{}{
                            "earum": "vel",
                        },
                        Duration: sdk.Int64(447378),
                        DurationUnit: shared.PlanPhaseDurationUnitMonthly,
                        Minimum: map[string]interface{}{
                            "illum": "soluta",
                            "accusantium": "aliquam",
                            "sapiente": "dicta",
                        },
                        Name: sdk.String("Delores Hermiston IV"),
                        Order: sdk.Int64(185232),
                    },
                    shared.PlanPhase{
                        Description: sdk.String("quibusdam"),
                        Discount: map[string]interface{}{
                            "deleniti": "itaque",
                            "dolorum": "architecto",
                        },
                        Duration: sdk.Int64(609178),
                        DurationUnit: shared.PlanPhaseDurationUnitAnnual,
                        Minimum: map[string]interface{}{
                            "at": "et",
                        },
                        Name: sdk.String("Mrs. Cynthia Hansen"),
                        Order: sdk.Int64(614465),
                    },
                },
                Prices: []shared.Price{
                    shared.Price{
                        BillableMetric: &shared.PriceBillableMetric{
                            ID: sdk.String("08086a18-4039-44c2-a071-f93f5f0642da"),
                        },
                        BpsConfig: &shared.PriceBpsConfig{
                            Bps: sdk.Float64(8070.23),
                            PerUnitMaximum: sdk.String("dignissimos"),
                        },
                        BulkBpsConfig: &shared.PriceBulkBpsConfig{
                            Tiers: []shared.PriceBulkBpsConfigTiers{
                                shared.PriceBulkBpsConfigTiers{
                                    Bps: sdk.Float64(9894.1),
                                    MaximumAmount: sdk.String("nemo"),
                                    PerUnitMaximum: sdk.String("quae"),
                                },
                                shared.PriceBulkBpsConfigTiers{
                                    Bps: sdk.Float64(3127.53),
                                    MaximumAmount: sdk.String("porro"),
                                    PerUnitMaximum: sdk.String("quod"),
                                },
                                shared.PriceBulkBpsConfigTiers{
                                    Bps: sdk.Float64(2883.98),
                                    MaximumAmount: sdk.String("ab"),
                                    PerUnitMaximum: sdk.String("adipisci"),
                                },
                            },
                        },
                        BulkConfig: &shared.PriceBulkConfig{
                            Tiers: []shared.PriceBulkConfigTiers{
                                shared.PriceBulkConfigTiers{
                                    MaximumUnits: sdk.String("id"),
                                    UnitAmount: sdk.String("suscipit"),
                                },
                                shared.PriceBulkConfigTiers{
                                    MaximumUnits: sdk.String("velit"),
                                    UnitAmount: sdk.String("culpa"),
                                },
                                shared.PriceBulkConfigTiers{
                                    MaximumUnits: sdk.String("est"),
                                    UnitAmount: sdk.String("recusandae"),
                                },
                            },
                        },
                        Cadence: shared.PriceCadenceMonthly.ToPointer(),
                        CreatedAt: types.MustTimeFromString("2021-09-23T14:56:28.537Z"),
                        Currency: "USD",
                        Discount: map[string]interface{}{
                            "quos": "vel",
                            "labore": "possimus",
                        },
                        FixedPriceQuantity: sdk.Float64(7065.75),
                        ID: sdk.String("b675fd5e-60b3-475e-94f6-fbee41f33317"),
                        MatrixConfig: &shared.PriceMatrixConfig{
                            DefaultUnitAmount: sdk.String("a"),
                            Dimensions: []string{
                                "consectetur",
                                "corporis",
                                "harum",
                                "laboriosam",
                            },
                            MatrixValues: []shared.PriceMatrixConfigMatrixValues{
                                shared.PriceMatrixConfigMatrixValues{
                                    DimensionValues: []string{
                                        "libero",
                                        "vitae",
                                        "accusamus",
                                        "similique",
                                    },
                                    UnitAmount: sdk.String("tempora"),
                                },
                            },
                        },
                        Minimum: map[string]interface{}{
                            "voluptas": "voluptas",
                        },
                        ModelType: shared.PriceModelTypeBulk.ToPointer(),
                        Name: sdk.String("Gayle Parisian"),
                        PackageConfig: &shared.PricePackageConfig{
                            PackageAmount: sdk.String("dolores"),
                            PackageSize: sdk.Float64(5039.34),
                        },
                        PlanPhaseOrder: 4492.92,
                        TieredBpsConfig: &shared.PriceTieredBpsConfig{
                            Tiers: []shared.PriceTieredBpsConfigTiers{
                                shared.PriceTieredBpsConfigTiers{
                                    Bps: sdk.Float64(3044.68),
                                    MaximumAmount: sdk.String("officiis"),
                                    MinimumAmount: sdk.String("temporibus"),
                                    PerUnitMaximum: sdk.String("ullam"),
                                },
                                shared.PriceTieredBpsConfigTiers{
                                    Bps: sdk.Float64(2377.42),
                                    MaximumAmount: sdk.String("cum"),
                                    MinimumAmount: sdk.String("blanditiis"),
                                    PerUnitMaximum: sdk.String("quas"),
                                },
                            },
                        },
                        TieredConfig: &shared.PriceTieredConfig{
                            Tiers: []shared.PriceTieredConfigTiers{
                                shared.PriceTieredConfigTiers{
                                    FirstUnit: sdk.String("nesciunt"),
                                    LastUnit: sdk.String("culpa"),
                                    UnitAmount: sdk.String("corrupti"),
                                },
                                shared.PriceTieredConfigTiers{
                                    FirstUnit: sdk.String("pariatur"),
                                    LastUnit: sdk.String("totam"),
                                    UnitAmount: sdk.String("hic"),
                                },
                                shared.PriceTieredConfigTiers{
                                    FirstUnit: sdk.String("exercitationem"),
                                    LastUnit: sdk.String("nobis"),
                                    UnitAmount: sdk.String("sit"),
                                },
                                shared.PriceTieredConfigTiers{
                                    FirstUnit: sdk.String("rerum"),
                                    LastUnit: sdk.String("sed"),
                                    UnitAmount: sdk.String("reiciendis"),
                                },
                            },
                        },
                        UnitConfig: &shared.PriceUnitConfig{
                            UnitAmount: sdk.String("explicabo"),
                        },
                    },
                    shared.Price{
                        BillableMetric: &shared.PriceBillableMetric{
                            ID: sdk.String("fb7b194a-276b-4269-96fe-1f08f4294e36"),
                        },
                        BpsConfig: &shared.PriceBpsConfig{
                            Bps: sdk.Float64(5799.12),
                            PerUnitMaximum: sdk.String("quos"),
                        },
                        BulkBpsConfig: &shared.PriceBulkBpsConfig{
                            Tiers: []shared.PriceBulkBpsConfigTiers{
                                shared.PriceBulkBpsConfigTiers{
                                    Bps: sdk.Float64(2716.53),
                                    MaximumAmount: sdk.String("tempora"),
                                    PerUnitMaximum: sdk.String("voluptate"),
                                },
                                shared.PriceBulkBpsConfigTiers{
                                    Bps: sdk.Float64(9700.76),
                                    MaximumAmount: sdk.String("ex"),
                                    PerUnitMaximum: sdk.String("sit"),
                                },
                                shared.PriceBulkBpsConfigTiers{
                                    Bps: sdk.Float64(2484.13),
                                    MaximumAmount: sdk.String("officiis"),
                                    PerUnitMaximum: sdk.String("praesentium"),
                                },
                                shared.PriceBulkBpsConfigTiers{
                                    Bps: sdk.Float64(7086.09),
                                    MaximumAmount: sdk.String("quaerat"),
                                    PerUnitMaximum: sdk.String("incidunt"),
                                },
                            },
                        },
                        BulkConfig: &shared.PriceBulkConfig{
                            Tiers: []shared.PriceBulkConfigTiers{
                                shared.PriceBulkConfigTiers{
                                    MaximumUnits: sdk.String("debitis"),
                                    UnitAmount: sdk.String("rem"),
                                },
                                shared.PriceBulkConfigTiers{
                                    MaximumUnits: sdk.String("sit"),
                                    UnitAmount: sdk.String("nobis"),
                                },
                            },
                        },
                        Cadence: shared.PriceCadenceMonthly.ToPointer(),
                        CreatedAt: types.MustTimeFromString("2022-09-02T17:12:20.523Z"),
                        Currency: "USD",
                        Discount: map[string]interface{}{
                            "reiciendis": "nulla",
                            "magni": "aperiam",
                            "saepe": "numquam",
                            "veniam": "in",
                        },
                        FixedPriceQuantity: sdk.Float64(8892.34),
                        ID: sdk.String("1858b6a8-9fbe-43a5-aa8e-4824d0ab4075"),
                        MatrixConfig: &shared.PriceMatrixConfig{
                            DefaultUnitAmount: sdk.String("sit"),
                            Dimensions: []string{
                                "quas",
                                "repudiandae",
                                "corporis",
                            },
                            MatrixValues: []shared.PriceMatrixConfigMatrixValues{
                                shared.PriceMatrixConfigMatrixValues{
                                    DimensionValues: []string{
                                        "ex",
                                        "sed",
                                        "sit",
                                    },
                                    UnitAmount: sdk.String("vel"),
                                },
                            },
                        },
                        Minimum: map[string]interface{}{
                            "saepe": "error",
                            "consequatur": "incidunt",
                        },
                        ModelType: shared.PriceModelTypeMatrix.ToPointer(),
                        Name: sdk.String("Ms. Opal Buckridge"),
                        PackageConfig: &shared.PricePackageConfig{
                            PackageAmount: sdk.String("quidem"),
                            PackageSize: sdk.Float64(5390.74),
                        },
                        PlanPhaseOrder: 6719.57,
                        TieredBpsConfig: &shared.PriceTieredBpsConfig{
                            Tiers: []shared.PriceTieredBpsConfigTiers{
                                shared.PriceTieredBpsConfigTiers{
                                    Bps: sdk.Float64(9488.61),
                                    MaximumAmount: sdk.String("laboriosam"),
                                    MinimumAmount: sdk.String("alias"),
                                    PerUnitMaximum: sdk.String("amet"),
                                },
                                shared.PriceTieredBpsConfigTiers{
                                    Bps: sdk.Float64(6471.97),
                                    MaximumAmount: sdk.String("voluptate"),
                                    MinimumAmount: sdk.String("unde"),
                                    PerUnitMaximum: sdk.String("reiciendis"),
                                },
                                shared.PriceTieredBpsConfigTiers{
                                    Bps: sdk.Float64(5887.4),
                                    MaximumAmount: sdk.String("repellendus"),
                                    MinimumAmount: sdk.String("delectus"),
                                    PerUnitMaximum: sdk.String("voluptates"),
                                },
                            },
                        },
                        TieredConfig: &shared.PriceTieredConfig{
                            Tiers: []shared.PriceTieredConfigTiers{
                                shared.PriceTieredConfigTiers{
                                    FirstUnit: sdk.String("est"),
                                    LastUnit: sdk.String("quidem"),
                                    UnitAmount: sdk.String("reprehenderit"),
                                },
                            },
                        },
                        UnitConfig: &shared.PriceUnitConfig{
                            UnitAmount: sdk.String("facere"),
                        },
                    },
                    shared.Price{
                        BillableMetric: &shared.PriceBillableMetric{
                            ID: sdk.String("a8a50ce1-87f8-46bc-973d-689eee9526f8"),
                        },
                        BpsConfig: &shared.PriceBpsConfig{
                            Bps: sdk.Float64(8717.86),
                            PerUnitMaximum: sdk.String("error"),
                        },
                        BulkBpsConfig: &shared.PriceBulkBpsConfig{
                            Tiers: []shared.PriceBulkBpsConfigTiers{
                                shared.PriceBulkBpsConfigTiers{
                                    Bps: sdk.Float64(3793.56),
                                    MaximumAmount: sdk.String("repudiandae"),
                                    PerUnitMaximum: sdk.String("atque"),
                                },
                                shared.PriceBulkBpsConfigTiers{
                                    Bps: sdk.Float64(5413.81),
                                    MaximumAmount: sdk.String("sunt"),
                                    PerUnitMaximum: sdk.String("recusandae"),
                                },
                                shared.PriceBulkBpsConfigTiers{
                                    Bps: sdk.Float64(6806.97),
                                    MaximumAmount: sdk.String("repellendus"),
                                    PerUnitMaximum: sdk.String("labore"),
                                },
                            },
                        },
                        BulkConfig: &shared.PriceBulkConfig{
                            Tiers: []shared.PriceBulkConfigTiers{
                                shared.PriceBulkConfigTiers{
                                    MaximumUnits: sdk.String("doloremque"),
                                    UnitAmount: sdk.String("repudiandae"),
                                },
                                shared.PriceBulkConfigTiers{
                                    MaximumUnits: sdk.String("dicta"),
                                    UnitAmount: sdk.String("accusantium"),
                                },
                                shared.PriceBulkConfigTiers{
                                    MaximumUnits: sdk.String("beatae"),
                                    UnitAmount: sdk.String("dolores"),
                                },
                                shared.PriceBulkConfigTiers{
                                    MaximumUnits: sdk.String("enim"),
                                    UnitAmount: sdk.String("laboriosam"),
                                },
                            },
                        },
                        Cadence: shared.PriceCadenceAnnual.ToPointer(),
                        CreatedAt: types.MustTimeFromString("2021-04-24T18:03:33.752Z"),
                        Currency: "USD",
                        Discount: map[string]interface{}{
                            "saepe": "consequuntur",
                            "occaecati": "officiis",
                        },
                        FixedPriceQuantity: sdk.Float64(5979.37),
                        ID: sdk.String("73e922a5-7a15-4be3-a060-807e2b6e3ab8"),
                        MatrixConfig: &shared.PriceMatrixConfig{
                            DefaultUnitAmount: sdk.String("rem"),
                            Dimensions: []string{
                                "ad",
                                "repellat",
                            },
                            MatrixValues: []shared.PriceMatrixConfigMatrixValues{
                                shared.PriceMatrixConfigMatrixValues{
                                    DimensionValues: []string{
                                        "perspiciatis",
                                        "nihil",
                                    },
                                    UnitAmount: sdk.String("mollitia"),
                                },
                            },
                        },
                        Minimum: map[string]interface{}{
                            "alias": "maiores",
                            "reiciendis": "dolores",
                        },
                        ModelType: shared.PriceModelTypeTieredBps.ToPointer(),
                        Name: sdk.String("Joanne Parisian DVM"),
                        PackageConfig: &shared.PricePackageConfig{
                            PackageAmount: sdk.String("omnis"),
                            PackageSize: sdk.Float64(3092.51),
                        },
                        PlanPhaseOrder: 4776.46,
                        TieredBpsConfig: &shared.PriceTieredBpsConfig{
                            Tiers: []shared.PriceTieredBpsConfigTiers{
                                shared.PriceTieredBpsConfigTiers{
                                    Bps: sdk.Float64(2840),
                                    MaximumAmount: sdk.String("culpa"),
                                    MinimumAmount: sdk.String("adipisci"),
                                    PerUnitMaximum: sdk.String("debitis"),
                                },
                                shared.PriceTieredBpsConfigTiers{
                                    Bps: sdk.Float64(5145.13),
                                    MaximumAmount: sdk.String("eum"),
                                    MinimumAmount: sdk.String("nemo"),
                                    PerUnitMaximum: sdk.String("recusandae"),
                                },
                            },
                        },
                        TieredConfig: &shared.PriceTieredConfig{
                            Tiers: []shared.PriceTieredConfigTiers{
                                shared.PriceTieredConfigTiers{
                                    FirstUnit: sdk.String("provident"),
                                    LastUnit: sdk.String("quis"),
                                    UnitAmount: sdk.String("eum"),
                                },
                                shared.PriceTieredConfigTiers{
                                    FirstUnit: sdk.String("reiciendis"),
                                    LastUnit: sdk.String("provident"),
                                    UnitAmount: sdk.String("aspernatur"),
                                },
                            },
                        },
                        UnitConfig: &shared.PriceUnitConfig{
                            UnitAmount: sdk.String("ullam"),
                        },
                    },
                    shared.Price{
                        BillableMetric: &shared.PriceBillableMetric{
                            ID: sdk.String("1a5a9da6-60ff-457b-baad-4f9efc1b4512"),
                        },
                        BpsConfig: &shared.PriceBpsConfig{
                            Bps: sdk.Float64(7652.71),
                            PerUnitMaximum: sdk.String("quae"),
                        },
                        BulkBpsConfig: &shared.PriceBulkBpsConfig{
                            Tiers: []shared.PriceBulkBpsConfigTiers{
                                shared.PriceBulkBpsConfigTiers{
                                    Bps: sdk.Float64(2419.01),
                                    MaximumAmount: sdk.String("aspernatur"),
                                    PerUnitMaximum: sdk.String("eum"),
                                },
                            },
                        },
                        BulkConfig: &shared.PriceBulkConfig{
                            Tiers: []shared.PriceBulkConfigTiers{
                                shared.PriceBulkConfigTiers{
                                    MaximumUnits: sdk.String("rem"),
                                    UnitAmount: sdk.String("at"),
                                },
                                shared.PriceBulkConfigTiers{
                                    MaximumUnits: sdk.String("impedit"),
                                    UnitAmount: sdk.String("eos"),
                                },
                            },
                        },
                        Cadence: shared.PriceCadenceQuarterly.ToPointer(),
                        CreatedAt: types.MustTimeFromString("2022-11-19T04:16:20.363Z"),
                        Currency: "USD",
                        Discount: map[string]interface{}{
                            "beatae": "cupiditate",
                            "provident": "earum",
                        },
                        FixedPriceQuantity: sdk.Float64(7453.98),
                        ID: sdk.String("fd0e9fe6-c632-4ca3-aed0-117996312fde"),
                        MatrixConfig: &shared.PriceMatrixConfig{
                            DefaultUnitAmount: sdk.String("ipsa"),
                            Dimensions: []string{
                                "nihil",
                                "molestiae",
                            },
                            MatrixValues: []shared.PriceMatrixConfigMatrixValues{
                                shared.PriceMatrixConfigMatrixValues{
                                    DimensionValues: []string{
                                        "esse",
                                        "praesentium",
                                    },
                                    UnitAmount: sdk.String("maiores"),
                                },
                            },
                        },
                        Minimum: map[string]interface{}{
                            "vel": "architecto",
                            "fugiat": "doloremque",
                            "dicta": "odio",
                            "tempora": "esse",
                        },
                        ModelType: shared.PriceModelTypePackage.ToPointer(),
                        Name: sdk.String("Jeanne Beer II"),
                        PackageConfig: &shared.PricePackageConfig{
                            PackageAmount: sdk.String("fugiat"),
                            PackageSize: sdk.Float64(7137.67),
                        },
                        PlanPhaseOrder: 3996.67,
                        TieredBpsConfig: &shared.PriceTieredBpsConfig{
                            Tiers: []shared.PriceTieredBpsConfigTiers{
                                shared.PriceTieredBpsConfigTiers{
                                    Bps: sdk.Float64(3813.97),
                                    MaximumAmount: sdk.String("aliquid"),
                                    MinimumAmount: sdk.String("perferendis"),
                                    PerUnitMaximum: sdk.String("eum"),
                                },
                                shared.PriceTieredBpsConfigTiers{
                                    Bps: sdk.Float64(3747.53),
                                    MaximumAmount: sdk.String("iste"),
                                    MinimumAmount: sdk.String("id"),
                                    PerUnitMaximum: sdk.String("ab"),
                                },
                                shared.PriceTieredBpsConfigTiers{
                                    Bps: sdk.Float64(6253.58),
                                    MaximumAmount: sdk.String("possimus"),
                                    MinimumAmount: sdk.String("voluptates"),
                                    PerUnitMaximum: sdk.String("mollitia"),
                                },
                            },
                        },
                        TieredConfig: &shared.PriceTieredConfig{
                            Tiers: []shared.PriceTieredConfigTiers{
                                shared.PriceTieredConfigTiers{
                                    FirstUnit: sdk.String("libero"),
                                    LastUnit: sdk.String("ad"),
                                    UnitAmount: sdk.String("deleniti"),
                                },
                                shared.PriceTieredConfigTiers{
                                    FirstUnit: sdk.String("enim"),
                                    LastUnit: sdk.String("vitae"),
                                    UnitAmount: sdk.String("repellendus"),
                                },
                                shared.PriceTieredConfigTiers{
                                    FirstUnit: sdk.String("ex"),
                                    LastUnit: sdk.String("quo"),
                                    UnitAmount: sdk.String("ex"),
                                },
                            },
                        },
                        UnitConfig: &shared.PriceUnitConfig{
                            UnitAmount: sdk.String("ut"),
                        },
                    },
                },
                Product: shared.PlanProduct{
                    CreatedAt: types.MustTimeFromString("2022-04-15T09:59:51.518Z"),
                    ID: "08b61891-baa0-4fe1-ade0-08e6f8c5f350",
                    Name: "Jimmie Russel",
                },
                TrialConfig: &shared.PlanTrialConfig{
                    TrialPeriod: sdk.Float64(3732.16),
                    TrialPeriodUnit: shared.PlanTrialConfigTrialPeriodUnitDays,
                },
            },
            shared.Plan{
                BasePlanID: sdk.String("culpa"),
                CreatedAt: types.MustTimeFromString("2022-09-10T19:23:06.016Z"),
                Currency: "inventore",
                Description: "deleniti",
                Discount: map[string]interface{}{
                    "tempora": "dolor",
                },
                ExternalPlanID: sdk.String("consequatur"),
                ID: "10421813-d520-48ec-a7e2-53b668451c6c",
                InvoicingCurrency: "laboriosam",
                Minimum: map[string]interface{}{
                    "consequuntur": "voluptatem",
                    "exercitationem": "necessitatibus",
                    "quasi": "nisi",
                    "at": "vero",
                },
                Name: "Kerry Dickinson",
                PlanPhases: []shared.PlanPhase{
                    shared.PlanPhase{
                        Description: sdk.String("occaecati"),
                        Discount: map[string]interface{}{
                            "voluptate": "blanditiis",
                            "officia": "voluptas",
                        },
                        Duration: sdk.Int64(254025),
                        DurationUnit: shared.PlanPhaseDurationUnitQuarterly,
                        Minimum: map[string]interface{}{
                            "eius": "aspernatur",
                            "ducimus": "nesciunt",
                            "fuga": "laudantium",
                        },
                        Name: sdk.String("Joyce Leffler III"),
                        Order: sdk.Int64(159845),
                    },
                    shared.PlanPhase{
                        Description: sdk.String("consectetur"),
                        Discount: map[string]interface{}{
                            "cupiditate": "reiciendis",
                        },
                        Duration: sdk.Int64(746837),
                        DurationUnit: shared.PlanPhaseDurationUnitMonthly,
                        Minimum: map[string]interface{}{
                            "eos": "occaecati",
                            "iste": "magni",
                            "inventore": "fuga",
                        },
                        Name: sdk.String("Jan Reichel"),
                        Order: sdk.Int64(328379),
                    },
                    shared.PlanPhase{
                        Description: sdk.String("praesentium"),
                        Discount: map[string]interface{}{
                            "magnam": "temporibus",
                            "quos": "commodi",
                            "itaque": "commodi",
                            "totam": "earum",
                        },
                        Duration: sdk.Int64(267207),
                        DurationUnit: shared.PlanPhaseDurationUnitAnnual,
                        Minimum: map[string]interface{}{
                            "voluptatem": "ipsam",
                            "vel": "alias",
                            "quasi": "non",
                            "maiores": "enim",
                        },
                        Name: sdk.String("Orville Nolan"),
                        Order: sdk.Int64(444587),
                    },
                    shared.PlanPhase{
                        Description: sdk.String("est"),
                        Discount: map[string]interface{}{
                            "sint": "accusamus",
                            "impedit": "hic",
                        },
                        Duration: sdk.Int64(900103),
                        DurationUnit: shared.PlanPhaseDurationUnitAnnual,
                        Minimum: map[string]interface{}{
                            "voluptas": "debitis",
                            "delectus": "quae",
                        },
                        Name: sdk.String("Grant Padberg"),
                        Order: sdk.Int64(538869),
                    },
                },
                Prices: []shared.Price{
                    shared.Price{
                        BillableMetric: &shared.PriceBillableMetric{
                            ID: sdk.String("c2beb477-373c-48d7-af64-d1db1f2c4310"),
                        },
                        BpsConfig: &shared.PriceBpsConfig{
                            Bps: sdk.Float64(4246.63),
                            PerUnitMaximum: sdk.String("ea"),
                        },
                        BulkBpsConfig: &shared.PriceBulkBpsConfig{
                            Tiers: []shared.PriceBulkBpsConfigTiers{
                                shared.PriceBulkBpsConfigTiers{
                                    Bps: sdk.Float64(8777.51),
                                    MaximumAmount: sdk.String("excepturi"),
                                    PerUnitMaximum: sdk.String("eum"),
                                },
                            },
                        },
                        BulkConfig: &shared.PriceBulkConfig{
                            Tiers: []shared.PriceBulkConfigTiers{
                                shared.PriceBulkConfigTiers{
                                    MaximumUnits: sdk.String("ut"),
                                    UnitAmount: sdk.String("perspiciatis"),
                                },
                            },
                        },
                        Cadence: shared.PriceCadenceQuarterly.ToPointer(),
                        CreatedAt: types.MustTimeFromString("2022-03-25T02:57:12.529Z"),
                        Currency: "USD",
                        Discount: map[string]interface{}{
                            "iste": "itaque",
                            "alias": "nisi",
                            "itaque": "velit",
                            "laborum": "non",
                        },
                        FixedPriceQuantity: sdk.Float64(2244.67),
                        ID: sdk.String("7000ae6b-6bc9-4b8f-b59e-ac55a9741d31"),
                        MatrixConfig: &shared.PriceMatrixConfig{
                            DefaultUnitAmount: sdk.String("inventore"),
                            Dimensions: []string{
                                "ad",
                            },
                            MatrixValues: []shared.PriceMatrixConfigMatrixValues{
                                shared.PriceMatrixConfigMatrixValues{
                                    DimensionValues: []string{
                                        "ex",
                                        "nemo",
                                        "soluta",
                                    },
                                    UnitAmount: sdk.String("libero"),
                                },
                            },
                        },
                        Minimum: map[string]interface{}{
                            "dolorum": "odio",
                            "fugit": "alias",
                            "magni": "vel",
                        },
                        ModelType: shared.PriceModelTypeUnit.ToPointer(),
                        Name: sdk.String("Pauline Durgan"),
                        PackageConfig: &shared.PricePackageConfig{
                            PackageAmount: sdk.String("et"),
                            PackageSize: sdk.Float64(2153.98),
                        },
                        PlanPhaseOrder: 6022.29,
                        TieredBpsConfig: &shared.PriceTieredBpsConfig{
                            Tiers: []shared.PriceTieredBpsConfigTiers{
                                shared.PriceTieredBpsConfigTiers{
                                    Bps: sdk.Float64(7143.76),
                                    MaximumAmount: sdk.String("maxime"),
                                    MinimumAmount: sdk.String("quia"),
                                    PerUnitMaximum: sdk.String("quia"),
                                },
                                shared.PriceTieredBpsConfigTiers{
                                    Bps: sdk.Float64(3421.37),
                                    MaximumAmount: sdk.String("omnis"),
                                    MinimumAmount: sdk.String("libero"),
                                    PerUnitMaximum: sdk.String("dicta"),
                                },
                                shared.PriceTieredBpsConfigTiers{
                                    Bps: sdk.Float64(6633.18),
                                    MaximumAmount: sdk.String("libero"),
                                    MinimumAmount: sdk.String("fugiat"),
                                    PerUnitMaximum: sdk.String("officia"),
                                },
                                shared.PriceTieredBpsConfigTiers{
                                    Bps: sdk.Float64(5546.03),
                                    MaximumAmount: sdk.String("placeat"),
                                    MinimumAmount: sdk.String("sit"),
                                    PerUnitMaximum: sdk.String("iusto"),
                                },
                            },
                        },
                        TieredConfig: &shared.PriceTieredConfig{
                            Tiers: []shared.PriceTieredConfigTiers{
                                shared.PriceTieredConfigTiers{
                                    FirstUnit: sdk.String("voluptates"),
                                    LastUnit: sdk.String("inventore"),
                                    UnitAmount: sdk.String("aperiam"),
                                },
                            },
                        },
                        UnitConfig: &shared.PriceUnitConfig{
                            UnitAmount: sdk.String("totam"),
                        },
                    },
                },
                Product: shared.PlanProduct{
                    CreatedAt: types.MustTimeFromString("2022-03-31T09:16:11.300Z"),
                    ID: "b0672d1a-d879-4eeb-9665-b85efbd02bae",
                    Name: "Mamie Torp",
                },
                TrialConfig: &shared.PlanTrialConfig{
                    TrialPeriod: sdk.Float64(4835.18),
                    TrialPeriodUnit: shared.PlanTrialConfigTrialPeriodUnitDays,
                },
            },
        },
        PaginationMetadata: map[string]interface{}{
            "odit": "explicabo",
            "corporis": "error",
            "earum": "adipisci",
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```
