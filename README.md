# IBAN validation

Code and infrastructure to validate IBAN written in Golang.

## Building

Ensure go is installed on your machine

To install external dependencies navigate to project's ./web dir and run:

```cmd

go get

```

## Web server

1. Run server

```go

go run ./web/iban-web-server.go

```

Default port is 8080

2. Now one can perform calls:

```cmd

curl -i http://127.0.0.1:8080/iban/validate/MT31MALT0110000000000000000012

```

Response:

```json

{
    "success": false,
    "message": "IBAN string has incorrect length"
}

```

Batch processing:

```cmd

curl --location --request POST "http://127.0.0.1:8080/iban/validate" --header "Content-Type: application/json"   --data "{\"ibans\":[\"MT31MALT01100000000000000000123\",\"GB82WEST1234569876543\"]}"

```

Response:

```json

[
    {
        "success": true,
        "message": ""
    },
    {
        "success": false,
        "message": "IBAN string has incorrect length"
    }
]

```

## Testing

```go

go ./iban/test

```