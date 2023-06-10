# CreateDiscount

The subscription's override discount for this price.


## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `AmountDiscount`                                                                 | **string*                                                                        | :heavy_minus_sign:                                                               | Only allowed if discount_type is amount                                          |
| `DiscountType`                                                                   | [*CreateDiscountDiscountType](../../models/shared/creatediscountdiscounttype.md) | :heavy_minus_sign:                                                               | N/A                                                                              |
| `PercentageDiscount`                                                             | **string*                                                                        | :heavy_minus_sign:                                                               | Only allowed if discount_type is percentage                                      |
| `UsageDiscount`                                                                  | **string*                                                                        | :heavy_minus_sign:                                                               | Only allowed if discount_type is usage                                           |