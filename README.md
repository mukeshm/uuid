# uuid
simple UUID v4 library

## Getting Started

### Installing

Use Go and run `go get`:

```sh
$ go get github.com/mukeshm/uuid
```
This will retrieve the library and place it in your `$GOPATH`

### Using UUID

Import uuid

```go
import "github.com/mukeshm/uuid"
```
This will import `uuid` library into your package

### Generating a new UUID v4

```go
u, err := uuid.GenerateV4()
uuid := u.String()
```