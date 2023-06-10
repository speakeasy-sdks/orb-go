# Ingest400ApplicationJSONDebug

Optional debug information (only present when debug=true is passed to the endpoint). Contains ingested and duplicate event idempotency keys.


## Fields

| Field              | Type               | Required           | Description        |
| ------------------ | ------------------ | ------------------ | ------------------ |
| `Duplicate`        | []*string*         | :heavy_minus_sign: | N/A                |
| `Ingested`         | []*string*         | :heavy_minus_sign: | N/A                |