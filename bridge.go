package main

import "fmt"

//IdrawShape interface.

type IdrawShape interface{
	drawShape(x[5] float32, y[5]float32)
}

//DrawShape class/struct.
type DrawShape struct{}


// drawShape will draw the shape given the cordinatess.
func (drawShape DrawShape) drawShape(x[5]float32, y[5]float32){
	fmt.Println("--- drawing shape ----")
}

//Icontour interface.
type Icontour interface{
	drawContour(x[5]float32,y[5]float32)
	resizeByFactor(factor int)
}


//drawContour constructor.
type DrawContour struct{
	x[5] float32
	y[5] float32
	shape DrawShape
	factor int
}
//DrawContour method drawContour given the coordinates

func (contour DrawContour) drawContour(x[5]float32,y[5]float32){
	fmt.Println("drawing contour")
	contour.shape.drawShape(contour.x,contour.y)
}

//DrawContour method resizeByFactor given factor
func (contour DrawContour) resizeByFactor(factor int) {
	contour.factor = factor
   }
   // main method
   func main() {
   var x = [5]float32{1,2,3,4,5}
   var y = [5]float32{1,2,3,4,5}
   var contour Icontour = DrawContour{x,y,DrawShape{},2}
   contour.drawContour(x,y)
	contour.resizeByFactor(2)
   }