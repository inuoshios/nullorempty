# nullorempty

The `nullorempty` package provides a utility function to check if a given value is `nil` or empty. This function can handle various types like `nil`, `string`, `slice`, `map`, `array`, and `pointer`, and determines whether they are "empty" based on their type.

## Installation

To use the `nullorempty` package, simply import it into your Go project:

```go
package main

import (
    "fmt"
    "github.com/inuoshios/nullorempty"
)

func main() {
    fmt.Println(nullorempty.IsNullOrEmpty(nil))         // true
    fmt.Println(nullorempty.IsNullOrEmpty(""))          // true
    fmt.Println(nullorempty.IsNullOrEmpty("Hello"))     // false
    fmt.Println(nullorempty.IsNullOrEmpty([]int{}))     // true
    fmt.Println(nullorempty.IsNullOrEmpty(map[string]int{})) // true
    fmt.Println(nullorempty.IsNullOrEmpty([]int{1, 2})) // false
}
````
