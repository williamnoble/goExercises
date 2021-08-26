package main

import (
	"fmt"
)

type unknown []interface{}

func (c *unknown) Put(elem interface{}) {
	//*c = append(*c, elem)
}

func (c *unknown) Get() interface{} {
	elem := (*c)[0]
	//	*c = (*c)[1:]
	return elem
}

func main() {
	intContainer := &unknown{}
	intContainer.Put(7)
	intContainer.Put(42)
	intContainer.Put(84)
	intContainer.Put(43)
	fmt.Println(intContainer)
	elem, ok := intContainer.Get().(int) // assert that the actual type is int
	if !ok {
		fmt.Println("Unable to read an int from intContainer")
	}
	fmt.Printf("assertExample: %d (%T)\n", elem, elem)

}
