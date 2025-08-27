package main

import (
	"fmt"
	"reflect"
)

func main() {

}

func getType(i interface{}) {
	t := reflect.TypeOf(i)

	switch t.Kind() {
	case reflect.Int:
		fmt.Println("Integer")
	case reflect.String:
		fmt.Println("String")
	case reflect.Bool:
		fmt.Println("Boolean")
	case reflect.Chan:
		fmt.Println("Chan")
	}

	// 	switch v := x.(type) {
	//    case int:
	//         return "int"
	//    case string:
	//         return "string"
	//    case bool:
	//         return "bool"
}
