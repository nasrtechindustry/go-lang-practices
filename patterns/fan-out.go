package patterns

import (
	"fmt"
	"sync"
)


// "Concurrent Execution using Goroutines and WaitGroup (Fan-Out)"
func RunFanOut() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	person := Person{Name: "Juma Khan", Age: 40}

	go func() {
		defer wg.Done()
		person.ShowName()
	}()

	go func() {
		defer wg.Done()
		person.ShowAge()
	}()

	wg.Wait()
	println("The program ends to execute here")

}

type Person struct {
	Name string
	Age  int
}

func (n Person) ShowName() {
	fmt.Println("My name is: ", n.Name)
}

func (n Person) ShowAge() {
	fmt.Println("And my Age is: ", n.Age)
}
