# go-json-datatypes
--
    import "github.com/segmentio/go-json-datatypes"

Get JSON datatype for given value.

## Example


```go

package main

import "github.com/segmentio/go-json-datatypes"
import "fmt"

func main() {
    fmt.Println(Datatype("hello world"))
    // String
    fmt.Println(Datatype(15))
    // Number
    fmt.Println(Datatype([]string{"hello"}))
    // Array
    // ..
}
```

## License

MIT
