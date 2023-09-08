# CountTransactionsRequest


## Fields

| Field                                                                                                                                                                           | Type                                                                                                                                                                            | Required                                                                                                                                                                        | Description                                                                                                                                                                     | Example                                                                                                                                                                         |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `account`                                                                                                                                                                       | *Optional[str]*                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                              | Filter transactions with postings involving given account, either as source or destination (regular expression placed between ^ and $).                                         | users:001                                                                                                                                                                       |
| `destination`                                                                                                                                                                   | *Optional[str]*                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                              | Filter transactions with postings involving given account at destination (regular expression placed between ^ and $).                                                           | users:001                                                                                                                                                                       |
| `end_time`                                                                                                                                                                      | [date](https://docs.python.org/3/library/datetime.html#date-objects)                                                                                                            | :heavy_minus_sign:                                                                                                                                                              | Filter transactions that occurred before this timestamp.<br/>The format is RFC3339 and is exclusive (for example, "2023-01-02T15:04:01Z" excludes the first second of 4th minute).<br/> |                                                                                                                                                                                 |
| `ledger`                                                                                                                                                                        | *str*                                                                                                                                                                           | :heavy_check_mark:                                                                                                                                                              | Name of the ledger.                                                                                                                                                             | ledger001                                                                                                                                                                       |
| `metadata`                                                                                                                                                                      | dict[str, *str*]                                                                                                                                                                | :heavy_minus_sign:                                                                                                                                                              | Filter transactions by metadata key value pairs. Nested objects can be used like this -> metadata[key]=value1&metadata[a.nested.key]=value2                                     |                                                                                                                                                                                 |
| `reference`                                                                                                                                                                     | *Optional[str]*                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                              | Filter transactions by reference field.                                                                                                                                         | ref:001                                                                                                                                                                         |
| `source`                                                                                                                                                                        | *Optional[str]*                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                              | Filter transactions with postings involving given account at source (regular expression placed between ^ and $).                                                                | users:001                                                                                                                                                                       |
| `start_time`                                                                                                                                                                    | [date](https://docs.python.org/3/library/datetime.html#date-objects)                                                                                                            | :heavy_minus_sign:                                                                                                                                                              | Filter transactions that occurred after this timestamp.<br/>The format is RFC3339 and is inclusive (for example, "2023-01-02T15:04:01Z" includes the first second of 4th minute).<br/> |                                                                                                                                                                                 |