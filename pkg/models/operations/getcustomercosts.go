// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Orb/pkg/models/shared"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// GetCustomerCostsViewMode - Controls whether Orb returns cumulative costs since the start of the billing period, or incremental day-by-day costs. If your customer has minimums or discounts, it's strongly recommended that you use the default cumulative behavior.
type GetCustomerCostsViewMode string

const (
	GetCustomerCostsViewModePeriodic   GetCustomerCostsViewMode = "periodic"
	GetCustomerCostsViewModeCumulative GetCustomerCostsViewMode = "cumulative"
)

func (e GetCustomerCostsViewMode) ToPointer() *GetCustomerCostsViewMode {
	return &e
}

func (e *GetCustomerCostsViewMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "periodic":
		fallthrough
	case "cumulative":
		*e = GetCustomerCostsViewMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetCustomerCostsViewMode: %v", v)
	}
}

type GetCustomerCostsRequest struct {
	// The Orb Customer ID
	CustomerID string `pathParam:"style=simple,explode=false,name=customer_id"`
	// Groups per-price costs by the key provided.
	GroupBy *string `queryParam:"style=form,explode=true,name=group_by"`
	// Costs returned are exclusive of `timeframe_end`.
	TimeframeEnd *string `queryParam:"style=form,explode=true,name=timeframe_end"`
	// Costs returned are inclusive of `timeframe_start`.
	TimeframeStart *time.Time `queryParam:"style=form,explode=true,name=timeframe_start"`
	// Controls whether Orb returns cumulative costs since the start of the billing period, or incremental day-by-day costs. If your customer has minimums or discounts, it's strongly recommended that you use the default cumulative behavior.
	ViewMode *GetCustomerCostsViewMode `queryParam:"style=form,explode=true,name=view_mode"`
}

type GetCustomerCosts200ApplicationJSONDataPerPriceCostsPriceGroups struct {
	// Grouping key to break down a single price's costs
	GroupingKey   string `json:"grouping_key"`
	GroupingValue string `json:"grouping_value"`
	// If the price is a matrix price, this is the second dimension key
	SecondaryGroupingKey   *string `json:"secondary_grouping_key,omitempty"`
	SecondaryGroupingValue *string `json:"secondary_grouping_value,omitempty"`
	// Total costs for this group for the timeframe. Note that this does not account for any minimums or discounts.
	Total string `json:"total"`
}

// GetCustomerCosts200ApplicationJSONDataPerPriceCosts - Price's contributions for the timeframe, excluding any transforms (minimums and discounts).
type GetCustomerCosts200ApplicationJSONDataPerPriceCosts struct {
	// Orb supports a few different pricing models out of the box. Each of these models is serialized differently in a given Price object. The model_type field determines the key for the configuration object that is present.
	//
	// ## Unit pricing
	// With unit pricing, each unit costs a fixed amount.
	// ```json
	// {
	//     ...
	//     "model_type": "unit",
	//     "unit_config": {
	//         "unit_amount": "0.50"
	//     }
	//     ...
	// }
	// ```
	//
	// ## Tiered pricing
	// In tiered pricing, the cost of a given unit depends on the tier range that it falls into, where each tier range is defined by an upper and lower bound. For example, the first ten units may cost $0.50 each and all units thereafter may cost $0.10 each.
	// ```json
	// {
	//     ...
	//     "model_type": "tiered",
	//     "tiered_config": {
	//         "tiers": [
	//            {
	//                 "first_unit": 1,
	//                 "last_unit": 10,
	//                 "unit_amount": "0.50"
	//             },
	//             {
	//                 "first_unit": 11,
	//                 "last_unit": null,
	//                 "unit_amount": "0.10"
	//             }
	//         ]
	//     }
	//     ...
	// }
	// ```
	//
	// ## Bulk pricing
	// Bulk pricing applies when the number of units determine the cost of all units. For example, if you've bought less than 10 units, they may each be $0.50 for a total of $5.00. Once you've bought more than 10 units, all units may now be priced at $0.40 (i.e. 101 units total would be $40.40).
	// ```json
	// {
	//     ...
	//     "model_type": "bulk",
	//     "bulk_config": {
	//         "tiers": [
	//             {
	//                 "maximum_units": 10,
	//                 "unit_amount": "0.50"
	//             },
	//             {
	//                 "maximum_units": 1000,
	//                 "unit_amount": "0.40"
	//             }
	//           ]
	//     }
	//     ...
	// }
	// ```
	//
	// ## Package pricing
	// Package pricing defines the size or granularity of a unit for billing purposes. For example, if the package size is set to 5, then 4 units will be billed as 5 and 6 units will be billed at 10.
	// ```json
	// {
	//     ...
	//     "model_type": "package",
	//     "package_config": {
	//        "package_amount": "0.80",
	//        "package_size": 10
	//     }
	//     ...
	//  }
	// ```
	//
	// ## BPS pricing
	// BPS pricing specifies a per-event (e.g. per-payment) rate in one hundredth of a percent (the number of basis points to charge), as well as a cap per event to assess. For example, this would allow you to assess a fee of 0.25% on every payment you process, with a maximum charge of $25 per payment.
	// ```json
	// {
	//     ...
	//     "model_type": "bps",
	//     "bps_config": {
	//        "bps": 125,
	//        "per_unit_maximum": "11.00"
	//     }
	//     ...
	//  }
	// ```
	//
	// ## Bulk BPS pricing
	// Bulk BPS pricing specifies BPS parameters in a tiered manner, dependent on the total quantity across all events. Similar to bulk pricing, the BPS parameters of a given event depends on the tier range that the billing period falls into. Each tier range is defined by an upper bound. For example, after $1.5M of payment volume is reached, each individual payment may have a lower cap or a smaller take-rate.
	// ```json
	//     ...
	//     "model_type": "bulk_bps",
	//     "bulk_bps_config": {
	//         "tiers": [
	//            {
	//                 "maximum_amount": "1000000.00",
	//                 "bps": 125,
	//                 "per_unit_maximum": "19.00"
	//            },
	//           {
	//                 "maximum_amount": null,
	//                 "bps": 115,
	//                 "per_unit_maximum": "4.00"
	//             }
	//         ]
	//     }
	//     ...
	// }
	// ```
	//
	// ## Tiered BPS pricing
	// Tiered BPS pricing specifies BPS parameters in a graduated manner, where an event's applicable parameter is a function of its marginal addition to the period total. Similar to tiered pricing, the BPS parameters of a given event depends on the tier range that it falls into, where each tier range is defined by an upper and lower bound. For example, the first few payments may have a 0.8 BPS take-rate and all payments after a specific volume may incur a take-rate of 0.5 BPS each.
	// ```json
	//     ...
	//     "model_type": "tiered_bps",
	//     "tiered_bps_config": {
	//         "tiers": [
	//            {
	//                 "minimum_amount": "0",
	//                 "maximum_amount": "1000000.00",
	//                 "bps": 125,
	//                 "per_unit_maximum": "19.00"
	//            },
	//           {
	//                 "minimum_amount": "1000000.00",
	//                 "maximum_amount": null,
	//                 "bps": 115,
	//                 "per_unit_maximum": "4.00"
	//             }
	//         ]
	//     }
	//     ...
	// }
	// ```
	//
	// ## Matrix pricing
	// Matrix pricing defines a set of unit prices in a one or two-dimensional matrix. `dimensions` defines the two event property values evaluated in this pricing model. In a one-dimensional matrix, the second value is `null`. Every configuration has a list of `matrix_values` which give the unit prices for specified property values. In a one-dimensional matrix, the matrix values will have `dimension_values` where the second value of the pair is null. If an event does not match any of the dimension values in the matrix, it will resort to the `default_unit_amount`.
	// ```json
	// ...
	// "model_type": "matrix"
	// "matrix_config": {
	//     "default_unit_amount": "3.00",
	//     "dimensions": [
	//         "cluster_name",
	//         "region"
	//     ],
	//     "matrix_values": [
	//         {
	//             "dimension_values": [
	//                 "alpha",
	//                 "west"
	//             ],
	//             "unit_amount": "2.00"
	//         },
	//         ...
	//     ]
	// }
	// ...
	// ```
	//
	// ### Fixed fees
	// Fixed fees are prices that are applied independent of usage quantites, and follow unit pricing. They also have an additional parameter `fixed_price_quantity`. If the Price represents a fixed cost, this represents the quantity of units applied.
	// ```json
	// {
	//     ...
	//     "id": "price_id",
	//     "model_type": "unit",
	//     "unit_config": {
	//        "unit_amount": "2.00"
	//     },
	//     "fixed_price_quantity": 3.0
	//     ...
	// }
	// ```
	//
	Price *shared.Price `json:"price,omitempty"`
	// If a `group_by` attribute is passed in, array of costs per `grouping_key`, `grouping_value` or `secondary_grouping_key`, `secondary_grouping_value`.
	PriceGroups []GetCustomerCosts200ApplicationJSONDataPerPriceCostsPriceGroups `json:"price_groups"`
	// Price's contributions for the timeframe, excluding any minimums and discounts.
	Subtotal *string `json:"subtotal,omitempty"`
	// Price's contributions for the timeframe, including any minimums and discounts.
	Total *string `json:"total,omitempty"`
}

type GetCustomerCosts200ApplicationJSONData struct {
	PerPriceCosts []GetCustomerCosts200ApplicationJSONDataPerPriceCosts `json:"per_price_costs"`
	// Total costs for the timeframe, excluding minimums and discounts.
	Subtotal       string    `json:"subtotal"`
	TimeframeEnd   time.Time `json:"timeframe_end"`
	TimeframeStart time.Time `json:"timeframe_start"`
	// Total costs for the timeframe, including minimums and discounts.
	Total string `json:"total"`
}

// GetCustomerCosts200ApplicationJSON - OK
type GetCustomerCosts200ApplicationJSON struct {
	Data               []GetCustomerCosts200ApplicationJSONData `json:"data"`
	PaginationMetadata map[string]interface{}                   `json:"pagination_metadata"`
}

type GetCustomerCostsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	GetCustomerCosts200ApplicationJSONObject *GetCustomerCosts200ApplicationJSON
}
