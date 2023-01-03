package helper

import (
	"fmt"
	"reflect"
)

func MapIt(fromStruct interface{}, toStruct interface{}) (outStruct interface{}) {
	fmt.Print("Mapping Starts")

	e := reflect.ValueOf(&fromStruct).Elem()
	f := reflect.ValueOf(&toStruct).Elem()

	for i := 0; i < e.NumField(); i++ {
		for j := 0; j < f.NumField(); j++ {

		}
	}

	return outStruct
}
