package singleton_test

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

// Person a person
type Person struct {
	Name string
}

func (p *Person) greet(messages string) {
	fmt.Printf("hello %s\n", p.Name)
}

var once sync.Once
var person *Person

// CreatePerson used  to create a singleton object
func NewPerson(name string) *Person {
	once.Do(func() {
		person = &Person{}
		person.Name = name
	})
	return person
}

func TestSingleton(t *testing.T) {
	wg := new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			person := NewPerson("Larry")
			fmt.Printf("pointer address is %x\n", unsafe.Pointer(person))
			wg.Done()
		}(i)
	}
	wg.Wait()
}
