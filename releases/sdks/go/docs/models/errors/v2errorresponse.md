# V2ErrorResponse


## Fields

| Field                                                                                        | Type                                                                                         | Required                                                                                     | Description                                                                                  | Example                                                                                      |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `Details`                                                                                    | **string*                                                                                    | :heavy_minus_sign:                                                                           | N/A                                                                                          | https://play.numscript.org/?payload=eyJlcnJvciI6ImFjY291bnQgaGFkIGluc3VmZmljaWVudCBmdW5kcyJ9 |
| `ErrorCode`                                                                                  | [V2ErrorsEnum](../../models/errors/v2errorsenum.md)                                          | :heavy_check_mark:                                                                           | N/A                                                                                          | VALIDATION                                                                                   |
| `ErrorMessage`                                                                               | *string*                                                                                     | :heavy_check_mark:                                                                           | N/A                                                                                          | [VALIDATION] invalid 'cursor' query param                                                    |