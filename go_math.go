package gomath

import "log"

// Add function
func Add(a, b int) int {
	log.Println(a, b)
	return a + b
}

// Sub function
func Sub(a, b int) int {
	return a - b
}

// Mul function
func Mul(a, b int) int {
	return a * b
}
