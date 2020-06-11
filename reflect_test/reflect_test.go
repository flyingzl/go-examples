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

func (s Student) String() string {
	return fmt.Sprintf("invoke String(): Name=%s", s.Name)
}

// 自动填充结构体
func fillBySetting(v interface{}, params map[string]interface{}) error {
	retType := reflect.TypeOf(v)
	if retType.Kind() != reflect.Ptr {
		return fmt.Errorf("%v should be pointer", v)
	} else {
		if retType.Elem().Kind() != reflect.Struct {
			return fmt.Errorf("%v should be struct", v)
		}
	}

	retVal := reflect.ValueOf(v).Elem()

	for key, val := range params {
		if realKey := retVal.FieldByName(key); !realKey.IsValid() || !realKey.CanSet() {
			continue
		} else {
			if realKey.Kind() == reflect.TypeOf(val).Kind() {
				realKey.Set(reflect.ValueOf(val))
			}
		}
	}
	return nil

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
	//  不传递任何参数
	fmt.Println(retValue.MethodByName("String").Call([]reflect.Value{}))
}

func TestReflect(t *testing.T) {
	u := Student{Name: "larry"}
	detect(u)
	fmt.Printf("-------------------\n")
	detect(&u)
}

type Dog struct {
	AnimalID string
	Name     string
	Age      int
}

type Person struct {
	UserID string
	Name   string
	Age    int
}

func TestReflectWithAmbious(t *testing.T) {
	dog := Dog{AnimalID: "xxxx"}
	person := new(Person)
	params := map[string]interface{}{
		"AnimalID": "ANIMAL_ID",
		"Name":     "James",
		"Age":      20,
	}
	if err := fillBySetting(&dog, params); err != nil {
		t.Error(err)
	}
	if err := fillBySetting(person, params); err != nil {
		t.Error(err)
	}
	t.Log(dog)
	t.Log(person)
}
