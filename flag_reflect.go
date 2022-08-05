package main

import (
	"errors"
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

type Rules map[string][]string

var PersonCheck = Rules{"Name": {"NotEmpty"}, "Age": {"NotEmpty"}, "Sex": {"NotEmpty"}}

func Check(st interface{}, rule Rules) error {
	typ := reflect.TypeOf(st)
	val := reflect.ValueOf(st)
	// fmt.Println(typ)
	// fmt.Println(val)
	if val.Kind() != reflect.Struct {
		return errors.New("struct expect")
	}
	num := val.NumField()
	for i := 0; i < num; i++ {
		struct_field := typ.Field(i)
		struct_val := val.Field(i)
		// fmt.Println(struct_field.Name)
		// fmt.Println(struct_val)
		if len(rule[struct_field.Name]) > 0 {
			for _, v := range rule[struct_field.Name] {
				switch {
				case v == "NotEmpty":
					if isBlack(struct_val) {
						return errors.New(struct_field.Name + "值不能为空")
					}
				default:
					return nil
				}

			}
		}
	}
	return nil

}

func isBlack(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.String:
		return value.Len() == 0
	case reflect.Bool:
		return value.Bool() == false
	case reflect.Interface, reflect.Ptr:
		return value.IsNil()
	}
	return reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
}

func main() {
	p := Person{
		Name: "xiaoming",
		Age:  0,
	}
	if err := Check(p, PersonCheck); err != nil {
		fmt.Println(err.Error())
	}
}
