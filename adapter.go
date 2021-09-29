package main

import "fmt"


type IProcess interface{
	process ()
}

type Adapter struct{
	adaptee Adaptee
}

// adapter class/struct has a process method that invokes the convert method on the adaptee.

func (adapter Adapter) process() {
	fmt.Println("Adapter process -----")
	adapter.adaptee.convert()
}

type Adaptee struct{
	adapterType int
}

func (adaptee Adaptee) convert(){
	fmt.Println("adaptee convert method")
}

func main() {
	var processor IProcess = Adapter{}
	processor.process()
}