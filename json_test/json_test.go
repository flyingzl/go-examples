package json_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"time"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestJson(t *testing.T) {
	d := map[string]int{"x": 2, "y": 22}
	z, _ := json.Marshal(d)
	fmt.Println(z)
	result := map[string]int{}
	_ = json.Unmarshal(z, &result)
	fmt.Println(result)

	user := User{"larry", 4}
	v, _ := json.Marshal(user)
	json.NewEncoder(os.Stdout).Encode(user)

	newUser := &User{}
	if err := json.Unmarshal(v, newUser); err != nil {
		panic(err)
	} else {
		fmt.Println(newUser)
	}

	bytes := []byte(`{"num":6.13,"strs":["a","b"]}`)
	var data map[string]interface{}
	if err := json.Unmarshal(bytes, &data); err != nil {
		panic(err)
	}

	strs := data["strs"].([]interface{})
	fmt.Println(strs[0])

	weekDay := time.Now().Weekday()
	fmt.Println(weekDay)
}
