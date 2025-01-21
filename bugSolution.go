```go
package main

import "fmt"

func main() {
    var m map[string]int
    // Check if the map is nil and initialize if it is.
    if m == nil {
        m = make(map[string]int)
    }
    m["foo"] = 1
    fmt.Println(m["foo"])
}
```