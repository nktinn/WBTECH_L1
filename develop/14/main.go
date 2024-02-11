package main

import (
	"fmt"
	"reflect"
)

/*
Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.
*/

// Type Switch

func GetType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("%d is an integer\n", v)
	case float32:
		fmt.Printf("%f is float32\n", v)
	case float64:
		fmt.Printf("%f is float64\n", v)
	case string:
		fmt.Printf("%q is a string\n", v)
	case bool:
		fmt.Printf("%t is a bool\n", v)
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

// Рефлексия

func GetTypeReflect(i interface{}) {
	t := reflect.TypeOf(i)
	switch t.Kind() {
	case reflect.Int:
		fmt.Printf("%d is an integer\n", i.(int))
	case reflect.Float32:
		fmt.Printf("%f is float32\n", i.(float32))
	case reflect.Float64:
		fmt.Printf("%f is float64\n", i.(float64))
	case reflect.String:
		fmt.Printf("%q is a string\n", i.(string))
	case reflect.Bool:
		fmt.Printf("%t is a bool\n", i.(bool))
	default:
		fmt.Printf("I don't know about type %T!\n", t)
	}
}

func main() {
	var i interface{}

	i = 14
	GetType(i)
	GetTypeReflect(i)

	i = "fourteen"
	GetType(i)
	GetTypeReflect(i)

	i = true
	GetType(i)
	GetTypeReflect(i)

	i = 14.0
	GetType(i)
	GetTypeReflect(i)
}
