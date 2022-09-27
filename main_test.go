package main

import (
	"fmt"
	"testing"
)

// func calculate() int {
// 	a := 1
// 	b := 2
// 	c := a + b
// 	fmt.Println(c)
// 	return c
// }

func TestCalculate(t *testing.T) {
	// expectedoutput := 3
	// a := 1
	// b := 2
	// c := a + b
	// // calculate()
	// output := c
	// if expectedoutput != output {
	// 	t.Errorf("Failed ! got %v want %c", output, expectedoutput)
	// } else {
	// 	t.Logf("Success !")
	// }
	output := Calculate()
	expectedoutput := 3
	fmt.Println(output)
	if expectedoutput != output {
		t.Fail()
	}

}
