package main

import "fmt"

func main() {
	var originalNumber int
	var convertedNumber float32

	fmt.Scanf("%d", &originalNumber) // Use %d for integer input and & to pass the address

	fmt.Printf("Original number: %d\n", originalNumber)

	convertedNumber = float32(originalNumber) // type conversion or type casting

	fmt.Printf("Converted number: %f\n", convertedNumber)

	originalNumber = int(convertedNumber) // again type casting

	fmt.Printf("float to int converted number: %v\n", originalNumber)
}
