# AddMetadataOnTransactionRequest


## Fields

| Field                                                                                                              | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        | Example                                                                                                            |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `idempotency_key`                                                                                                  | *Optional[str]*                                                                                                    | :heavy_minus_sign:                                                                                                 | Use an idempotency key                                                                                             |                                                                                                                    |
| `request_body`                                                                                                     | dict[str, *str*]                                                                                                   | :heavy_minus_sign:                                                                                                 | metadata                                                                                                           |                                                                                                                    |
| `async_`                                                                                                           | *Optional[bool]*                                                                                                   | :heavy_minus_sign:                                                                                                 | Set async mode.                                                                                                    | true                                                                                                               |
| `dry_run`                                                                                                          | *Optional[bool]*                                                                                                   | :heavy_minus_sign:                                                                                                 | Set the dryRun mode. Dry run mode doesn't add the logs to the database or publish a message to the message broker. | true                                                                                                               |
| `ledger`                                                                                                           | *str*                                                                                                              | :heavy_check_mark:                                                                                                 | Name of the ledger.                                                                                                | ledger001                                                                                                          |
| `txid`                                                                                                             | *int*                                                                                                              | :heavy_check_mark:                                                                                                 | Transaction ID.                                                                                                    | 1234                                                                                                               |