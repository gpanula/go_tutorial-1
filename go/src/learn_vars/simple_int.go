package main

import "fmt"

func main() {
	var width, height = 100, 50 // declare two vars in a singal go
	fmt.Println("width is", width, "height is", height)
	// now assign values to width  and height 
	width = 75
	height = 200
	fmt.Println ("width is", width, "height", height)
}