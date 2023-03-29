// Created by: Janet
// Created on: Sep 2020
//
// This program finds the volume of a cylinder

package main

import (
	"fmt"
	"math"
)

func main() {
	// input
	var radius float64
	var height float64
	fmt.Print("Enter the radius of the cylinder (in cm): ")
	fmt.Scanln(&radius)
	fmt.Print("Enter the height of the cylinder (in cm): ")
	fmt.Scanln(&height)
	fmt.Println()

	// process
	volume := 3.14 * math.Pow(radius, 2) * height

	// output
	fmt.Printf("The volume of the cylinder is: %.2f cm³\n", volume)

	 fmt.Println("\nDone.")
}