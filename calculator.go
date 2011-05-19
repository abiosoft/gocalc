package main

import (
	"fmt"
)

// Stores the result of calculations
var result struct {
	last     float32 // stores last result
	operator string  // the next operator to use
	active   bool    // whether calculation has started
}

// Perform calculation
func Calculation(val float32, operator string) {
	if !result.active {
		result.last = val
		result.operator = operator
		result.active = true
		return
	}
	switch result.operator {
	case "+":
		result.last += val
	case "-":
		result.last -= val
	case "/":
		result.last /= val
	case "x":
		result.last *= val
	}
	result.operator = operator
}

// Resets the calculator
func Reset() {
	result.last = 0
	result.operator = ""
	result.active = false
}

// Generate formatted output, returns a function
func GetResult() string {
	val := fmt.Sprintf("%f", result.last)
	exp := fmt.Sprintf("%E", result.last)
	trim := val[0 : len(val)-7]
	if len(trim) <= 10 {
		if val[len(val)-7:] == ".000000" {
			return trim
		}
		if len(val) <= 10 {
			return func() string {
				for i := len(val) - 1; i >= len(val)-7; i-- {
					if val[i] != '0' {
						return val[0 : i+1]
					}
				}
				return val
			}()
		}
	}
	return exp
}
