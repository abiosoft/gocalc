package main

import (
	"fmt"
)

//Result
var Result struct {
	last     float  //stores last result
	operator string //the next operator to use
	active   bool   //whether calculation has started
}
//perform calculation
func calculation(val float, operator string) {
	if !Result.active {
		Result.last = val
		Result.operator = operator
		Result.active = true
		return
	}
	switch Result.operator {
	case "+":
		Result.last += val
	case "-":
		Result.last -= val
	case "/":
		Result.last /= val
	case "x":
		Result.last *= val
	}
	Result.operator = operator
}
//resets the calculator
func reset() {
	Result.last = 0
	Result.operator = ""
	Result.active = false
}
//to generate formatted output
func getResult() string {
	return func() string {
		val := fmt.Sprintf("%f", Result.last)
		exp := fmt.Sprintf("%E", Result.last)
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
	}()
}
