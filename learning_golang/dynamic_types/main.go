package main

import "fmt"

func main() {
	result := add(1, 2)
	fmt.Println(result)
}

// fuunction to demonstrate type checking and type switches
func printSomething(value interface{}) {
	stringVal, ok := value.(string)
	if ok {
		fmt.Println("String: ", stringVal)
		return
	}

	intVal, ok := value.(int)
	if ok {
		fmt.Println("Int: ", intVal)
		return
	}

	floatVal, ok := value.(float64)
	if ok {
		fmt.Println("Float: ", floatVal)
		return
	}

	// switch value.(type) {
	// case string:
	// 	fmt.Println("String: ", value)
	// case int:
	// 	fmt.Println("Int: ", value)
	// }
	// fmt.Println(value)
}

// function to demonstrate generics
func add[T int | float64 | string](a, b T) T {
	return a + b
}
