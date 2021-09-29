package main

import (
	"fmt"
)

//gets the power series of integer a and returns tuple of square of a
// and cube of a
func powerSeries(a int) (int,int) {

	return a*a, a*a*a
  
  }

func main(){

	// tuples are sorted list of elements. A data structure that groups data and they are immutable.

	var square int 
	var cube int 
	square, cube = powerSeries(3)

	fmt.Println("Square ---->", square, "Cube ---->", cube)
}


