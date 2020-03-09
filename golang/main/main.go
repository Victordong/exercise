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
		if wValueOf.FieldByName("Owner").CanSet() {
			//newValue := reflect.New(wValueOf.Elem().FieldByName("Owner").Type())
			//wValueOf.Elem().FieldByName("Owner").Set()
			v := reflect.ValueOf("2313123")
			wValueOf.FieldByName("Owner").Set(v)
			vNum := reflect.ValueOf(1.1)
			wValueOf.FieldByName("Num").Set(vNum)
			fmt.Println(w)
			fmt.Println("can set")
		} else {
			fmt.Println("can not set")
		}
	} else {
		fmt.Println("can not")
	}
}
