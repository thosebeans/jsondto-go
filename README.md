# jsondto-go

jsondto-go provides __data transfer objects__ for working with JSON-data.  

## Installation

```sh
go get github.com/thosebeans/jsondto-go/jsondto
```
## Usage

```go
package main

import(
    "fmt"
    
    "github.com/thosebeans/jsondto-go/jsondto"
)

func main() {
    o := new(jsondto.Object)
    o.Put(jsondto.String("ip"),  jsondto.String("192.168.2.2"))
    o.Put(jsondto.String("port"), jsondto.Int(8080))
    
    data,_ := o.MarshalJSON()
    fmt.Println(string(data))
}
```

## License

[Unlicense](https://unlicense.org)
