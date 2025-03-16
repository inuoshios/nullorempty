package nullorempty

import "reflect"

// IsNullOrEmpty checks if the provided value is nil or empty.
// It returns true if the value is nil, an empty string, or an empty slice (including arrays and maps).
// It returns false otherwise.
//
// Supported types:
// - nil: Always returns true.
// - string: Returns true if the string is empty ("").
// - slice: Returns true if the slice has no elements.
// - map: Returns true if the map has no elements.
// - array: Returns true if the array has no elements.
// - pointer: Returns true if the pointer is nil.
//
// Example usage:
//
//	package main
//
//	import (
//	    "fmt"
//	    "github.com/inuoshios/nullorempty"
//	)
//
//	func main() {
//	    fmt.Println(nullorempty.IsNullOrEmpty(nil))         // true
//	    fmt.Println(nullorempty.IsNullOrEmpty(""))          // true
//	    fmt.Println(nullorempty.IsNullOrEmpty("Hello"))     // false
//	    fmt.Println(nullorempty.IsNullOrEmpty([]int{}))     // true
//	    fmt.Println(nullorempty.IsNullOrEmpty(map[string]int{})) // true
//	    fmt.Println(nullorempty.IsNullOrEmpty([]int{1, 2})) // false
//	}
//
// Returns:
// - bool: true if the value is nil or empty, false otherwise.
func IsNullOrEmpty(val any) bool {
	if val == nil {
		return true
	}

	v := reflect.ValueOf(val)

	switch v.Kind() {
	case reflect.String:
		return v.String() == ""
	case reflect.Array, reflect.Slice:
		return v.Len() == 0
	case reflect.Map:
		return v.Len() == 0
	case reflect.Ptr:
		return v.IsNil()
	}

	return true
}
