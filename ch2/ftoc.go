// ftoc prints two fahrenheit to celsius conversions
package main

import "fmt"

const freezingF, boilingF = 32.0, 212.0

func main() {
	fmt.Printf("%g\u00b0F = %g\u00b0C\n", freezingF, fToC(freezingF))
	fmt.Printf("%g\u00b0F = %g\u00b0C\n", boilingF, fToC(boilingF))
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}