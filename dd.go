package go_dd

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

// Dump prints the value and type of the given variable(s) and exits the program.
// It also prints the caller function name for better debugging context.
// Dump 打印传递的变量的类型和值，以及调用该函数的函数名称。
// 该函数主要用于调试，快速定位调用位置及变量的值。
// 参数:
//
//	v: 可变参数，允许传递多个值进行类型和值的打印。
func Dump(v ...interface{}) {
	// 使用 runtime.Caller 获取调用者信息
	pc, _, _, _ := runtime.Caller(1)
	// 根据程序计数器获取调用者的函数名称
	callerName := runtime.FuncForPC(pc).Name()

	// 拆分调用者名称以提取函数名称
	parts := strings.Split(callerName, ".")
	// 函数名称是拆分后的最后一部分
	funcName := parts[len(parts)-1]

	// 打印调用者函数名称
	fmt.Printf("从 %s 被调用\n", funcName)

	// 遍历提供的变量并打印其类型和值
	for _, value := range v {
		printValue(value)
	}

	// 打印分隔线以便于阅读
	fmt.Println("---")
}

// printValue prints the type and value of the given variable.
// printValue 打印变量的类型和值。
// 该函数使用反射来确定传递给函数的变量的类型和值，并打印它们。
// 参数:
//
//	v: 类型为interface{}的参数，可以是任意类型的值。
func printValue(v interface{}) {
	// 使用反射获取变量的类型和值
	val := reflect.ValueOf(v)

	// 打印变量的类型
	fmt.Printf("Type: %s\n", val.Type())

	// 以详细的格式打印变量的值
	fmt.Printf("Value: %#v\n", v)
}
