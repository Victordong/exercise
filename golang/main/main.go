package main

import (
	"fmt"
	"reflect"
)

type Wallet struct {
	Owner string
	Num   float64
}

func main() {
	w := &Wallet{
		Owner: "nicai",
		Num:   11,
	}
	wTypeOf := reflect.TypeOf(w)
	wValueOf := reflect.ValueOf(w)
	if wTypeOf.Kind() == reflect.Ptr {
		wTypeOf = wTypeOf.Elem()
	}
	if wValueOf.Kind() == reflect.Ptr {
		wValueOf = wValueOf.Elem()
	}
	fmt.Println(wValueOf.FieldByName("Owner").Type())
	if wValueOf.CanAddr() {
		fmt.Println(wValueOf.Addr())
		fmt.Println("can")
	} else {
		fmt.Println("can not")
		if wValueOf.FieldByName("Owner").CanSet() {
			//newValue := reflect.New(wValueOf.Elem().FieldByName("Owner").Type())
			//wValueOf.Elem().FieldByName("Owner").Set()
			fmt.Println("can set")
		} else {
			fmt.Println("can not set")
		}
	}
}
