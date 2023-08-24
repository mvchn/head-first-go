package main

import "fmt"

func paintNeed(width float64, height float64) {
	area := width * height
	fmt.Printf("%.2f liters needed\n", area/10.0)
}

func main() {
	paintNeed(4.2, 3.0)
	paintNeed(5.2, 3.5)
	paintNeed(5.0, 3.3)
}
