package main

import (
	"fmt"
)

type FooAdapter interface {
	Read()
	Write()
	Close()
}

type Foo struct {
	adapter FooAdapter
}

func NewFoo(a FooAdapter) *Foo {
	return &Foo{adapter: a}
}

type StubAdapter struct {
	FooAdapter
}

func(s *StubAdapter) Read() {
	fmt.Println("implemented")
}


func doSomething() {
	f := NewFoo(&StubAdapter{})

	f.adapter.Read()
	f.adapter.Write()
}

func main() {
	doSomething()
}