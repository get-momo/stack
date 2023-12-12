# CurrencyCloudConfig


## Fields

| Field                                                                         | Type                                                                          | Required                                                                      | Description                                                                   | Example                                                                       |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `APIKey`                                                                      | *string*                                                                      | :heavy_check_mark:                                                            | N/A                                                                           | XXX                                                                           |
| `Endpoint`                                                                    | **string*                                                                     | :heavy_minus_sign:                                                            | The endpoint to use for the API. Defaults to https://devapi.currencycloud.com | XXX                                                                           |
| `LoginID`                                                                     | *string*                                                                      | :heavy_check_mark:                                                            | Username of the API Key holder                                                | XXX                                                                           |
| `Name`                                                                        | *string*                                                                      | :heavy_check_mark:                                                            | N/A                                                                           | My CurrencyCloud Account                                                      |
| `PollingPeriod`                                                               | **string*                                                                     | :heavy_minus_sign:                                                            | The frequency at which the connector will fetch transactions                  | 60s                                                                           |