// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Print the literals
//
//  1. Print a few integer literals
//
//  2. Print a few float literals
//
//  3. Print true and false bool literals
//
//  4. Print your name using a string literal
//
//  5. Print a non-english sentence using a string literal
//
// ---------------------------------------------------------

func main() {
	// Use fmt.Println()
	//  1. Print a few integer literals
	fmt.Println(
		-200, -10, 0, 50, 100,
	)
	//  2. Print a few float literals
	fmt.Println(
		-50.5, 20.5, -0., -1., 100.234,
	)
	//  3. Print true and false bool literals
	fmt.Println(
		true, false,
	)
	//  4. Print your name using a string literal
	fmt.Println(
		"Hello there,",
		"My name Tirza.",
	)
	//  5. Print a non-english sentence using a string literal
	fmt.Println("Merhaba, adım İnanç")
}
