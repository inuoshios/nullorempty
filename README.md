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

# License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

# Contributing
Please feel free to contribute to this project by submitting pull requests or opening issues. Your contributions are highly appreciated!

# Authors
- [Inu John](https://github.com/inuoshios)
