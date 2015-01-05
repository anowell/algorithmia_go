Algorithmia Go Client Library
-----------------------------

A minimal library for querying Algorithmia from Go.

## Usage

```go
import "algorithmia"
results, err := algorithmia.Query("kenny/Factor", api_key, input_data_reader)
```

## Example

See [examples/factor/factor.go] for a working example

    $ export ALGORITHMIA_API_KEY=11111111112222222222333333333300
    $ go run examples/factor.go 2379
    {"duration":0.00313746,"result":[3,41]}%

