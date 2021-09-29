package main

import (
	"container/list"
	"fmt"
)


func main(){

	// list is a sequence of elements with each elements connected to another
	// via link forward and backward.These elements can be removed and added as well

	var intList list.List

	// adding element to list.
	intList.PushBack(1)
	intList.PushBack(2)
	intList.PushBack(3)
	intList.PushBack(4)
	intList.PushBack(5)

	// iterate thru the list and check if there's an element, elements will be accessed thru the Value() method
	for element := intList.Front(); element != nil; element = element.Next(){
		fmt.Println(element.Value.(int))
	}


}