package reflect_test

import (
	"fmt"
	"reflect"
	"testing"
)

type Student struct {
	Name string `json:"userName"`
}

func (s *Student) Greet(msg string) {
	fmt.Printf("%s say: %s\n", s.Name, msg)
}

func detect(v interface{}) {

	retType := reflect.TypeOf(v)
	retValue := reflect.ValueOf(v)
	fmt.Println(retType, retValue)

	// 检测是否 为 指针
	if retType.Kind() == reflect.Ptr {
		if field, ok := retType.Elem().FieldByName("Name"); ok {
			fmt.Println(field.Tag.Get("json"))
		}

		// 检测是否允许修改
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
	fnVal := retValue.MethodByName("Greet")
	// Greet方法只能在使用指针时才调用
	if fnVal.IsValid() {
		fnVal.Call(params)
	}
}

func TestReflect(t *testing.T) {
	u := Student{Name: "larry"}
	detect(u)
	fmt.Printf("-------------------\n")
	detect(&u)
}
