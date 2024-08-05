package go_dd

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

// Dump prints the value and type of the given variable(s) and exits the program.
// It also prints the caller function name for better debugging context.
func Dump(v ...interface{}) {
	// Get the caller information using runtime.Caller
	pc, _, _, _ := runtime.Caller(1)
	callerName := runtime.FuncForPC(pc).Name()

	// Split the caller name to extract the function name
	parts := strings.Split(callerName, ".")
	funcName := parts[len(parts)-1]

	// Print the caller function name
	fmt.Printf("Called from: %s\n", funcName)

	// Iterate over the provided variables and print their type and value
	for _, value := range v {
		printValue(value)
	}

	// Print a separator for better readability
	fmt.Println("---")
}

// printValue prints the type and value of the given variable.
func printValue(v interface{}) {
	// Use reflection to get the type and value of the variable
	val := reflect.ValueOf(v)

	// Print the type of the variable
	fmt.Printf("Type: %s\n", val.Type())

	// Print the value of the variable in a detailed format
	fmt.Printf("Value: %#v\n", v)
}
