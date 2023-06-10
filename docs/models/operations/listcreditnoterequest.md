# ListCreditNoteRequest


## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `CustomerID`                                                               | **string*                                                                  | :heavy_minus_sign:                                                         | Filter by a specific customer (cannot be used with `external_customer_id`) |
| `ExternalCustomerID`                                                       | **string*                                                                  | :heavy_minus_sign:                                                         | Filter by a specific customer (cannot be used with `customer_id`)          |
| `SubscriptionID`                                                           | **string*                                                                  | :heavy_minus_sign:                                                         | Filter by a specific subscription                                          |