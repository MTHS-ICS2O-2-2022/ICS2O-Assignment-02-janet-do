// Created by: Janet
// Created on: Sep 2020
//
// This program finds the area of a triangle

package main

import "fmt"

func main() {
	// input
	var radius float64
	var height float64
	fmt.Print("Enter the radius of the cylinder (in cm): ")
	fmt.Scanln(&radius)
	fmt.Print("Enter the height of the cylinder (in cm): ")
	fmt.Scanln(&height)

	// process
	volume := 3.14 * math.Pow(radius, 2) * height

	// output
	fmt.Printf("The volume of the cylinder is: %.2f cm³\n", volume)
}
