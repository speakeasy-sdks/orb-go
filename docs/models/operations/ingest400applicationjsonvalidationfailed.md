# Ingest400ApplicationJSONValidationFailed


## Fields

| Field                                                                              | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `IdempotencyKey`                                                                   | **string*                                                                          | :heavy_minus_sign:                                                                 | N/A                                                                                |
| `ValidationErrors`                                                                 | []*string*                                                                         | :heavy_minus_sign:                                                                 | An array of objects corresponding to validation failures for each idempotency_key. |