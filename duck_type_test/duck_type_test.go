package duck_type_test

import (
	"fmt"
	"math/rand"
	"testing"
	"unsafe"
)

type Animal interface {
	Say()
}

type Pet struct {
	Name string
}

func (p *Pet) String() string {
	return fmt.Sprintf("hello, %s", p.Name)
}

func (p *Pet) Say() {
	p.Name = "Jack"
	fmt.Printf("hello, %s, address is %x\n", p.Name, unsafe.Pointer(&p.Name))
}

func show(animal Animal) {
	animal.Say()
}

func TestDuckType(t *testing.T) {
	pet := new(Pet)
	pet.Name = "Lily"

	// pet = &Pet{Name: "Lily"}
	t.Logf("address is :%x\n", unsafe.Pointer(&pet.Name))
	pet.Say()
	t.Log(pet.String())
	show(pet)
}

func sum(numbers []int, result chan int) {
	ret := 0
	for _, val := range numbers {
		ret += val
	}
	result <- ret
}

func TestChan(t *testing.T) {
	numbers := []int{}
	for i := 0; i < 100; i++ {
		numbers = append(numbers, (int)(rand.Int31n(1000)))
	}

	length := len(numbers)

	result := make(chan int, 2)

	for i := 0; i < 2; i++ {
		go sum(numbers[:length/2], result)
		go sum(numbers[length/2:], result)
	}
	a, b := <-result, <-result
	t.Log(a + b)

	t.Log(5 / 2)
}
