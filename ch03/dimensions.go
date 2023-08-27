package main

import "fmt"

func paintNeed(width float64, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("a width of %0.2f is invalid", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("a height of %0.2f is invalid", height)
	}
	area := width * height
	return area / 10.0, nil
}

func main() {
	var total float64
	amount, err := paintNeed(4.2, 3)
	fmt.Println(err)
	fmt.Printf("%0.2f liters needed\n", amount)
	total += amount
	amount, err = paintNeed(5.2, 3.5)
	fmt.Printf("%0.2f liters needed\n", amount)
	total += amount
	fmt.Printf("Total: %0.2f liters\n", total)
}
