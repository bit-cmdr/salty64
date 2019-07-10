# Salty64

This is a [Go](https://golang.org) package that is used to generate a hash of a string with a minimum level of security. The ideal scenario is if you need to generate an ID for a record that is based on one or more fields combined together and will be the same every time it is generated. This is not meant to be a secure string or used for passwords in any way, simply as a hash to identify a record in a way that does not leak anything like an email address or name in plain text.

## Installation

```sh
go get -u github.com/bit-cmdr/salty64
```

## Testing

```sh
go test .
```

## Benchmarking

```sh
go test -bench=. -benchmem
```

## Example

```go
package main

import (
	"github.com/bit-cmdr/salty64"
	"fmt"
)

func main() {
    email := "testemail@test.email"

    s, err := salty64.NewShaker("secret", 2)
    if err != nil {
        fmt.Printf("error: %s\n", err)
        return
    }

    hash, err := s.Encode(email)
    if err != nil {
        fmt.Printf("error: %s\n", err)
        return
    }

    fmt.Printf("hash: %s\n", hash)
}
```
