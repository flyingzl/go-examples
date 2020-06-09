package reflect_test

import (
	"fmt"
	"reflect"
	"testing"
)

type Student struct {
	Name string `json:"userName"`
}

func (s Student) Greet(msg string) {
	fmt.Printf("%s say: %s\n", s.Name, msg)
}

func detect(v interface{}) {

	retType := reflect.TypeOf(v)
	retValue := reflect.ValueOf(v)
	fmt.Println(retType, retValue)

	if retType.Kind() == reflect.Ptr {
		if field, ok := retType.Elem().FieldByName("Name"); ok {
			fmt.Println(field.Tag.Get("json"))
		}

		nameVal := retValue.Elem().FieldByName("Name")
		if nameVal.CanSet() {
			nameVal.SetString("Jim")
		}

		fmt.Println(retValue.Elem().FieldByName("Name").String())

	} else {
		if field, ok := retType.FieldByName("Name"); ok {
			fmt.Println(field.Tag.Get("json"))
		}
		fmt.Println(retValue.FieldByName("Name").String())

	}

	params := []reflect.Value{reflect.ValueOf("hello")}
	reflect.ValueOf(v).MethodByName("Greet").Call(params)
}

func TestReflect(t *testing.T) {
	u := Student{Name: "larry"}
	detect(u)
	fmt.Printf("-------------------\n")
	detect(&u)
}
