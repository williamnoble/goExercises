package main

import (
	"fmt"
	"reflect"
)

type Cabinet struct {
	s reflect.Value
}

func NewCabinet(t reflect.Type) *Cabinet {
	return &Cabinet{
		s: reflect.MakeSlice(reflect.SliceOf(t), 0, 10),
	}
}

func (c *Cabinet) Put(val interface{}) {
	if reflect.ValueOf(val).Type() != c.s.Type().Elem() {
		panic(fmt.Sprintf("Put: Cannot put a %T into a slice of %s", val, c.s.Type().Elem()))
	}
	c.s = reflect.Append(c.s, reflect.ValueOf(val))
}

func (c *Cabinet) Get(retref interface{}) {
	retref = c.s.Index(0)
	c.s = c.s.Slice(1, c.s.Len())
}
func main() {
	f := 3.14152 //pi
	g := 0.1
	fmt.Println(reflect.TypeOf(f))
	fmt.Println(reflect.ValueOf(f))
	c := NewCabinet(reflect.TypeOf(f))
	c.Put(f)
	fmt.Println(c.s.Index(0))
	c.Get(&g)
	fmt.Printf("reflectExample: %f (%T)\n", g, g)
}
