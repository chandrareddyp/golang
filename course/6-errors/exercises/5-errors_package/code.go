package main

import (
	"errors"
	"fmt"
)

func divide(x, y float64) (float64, error) {
	var f float64
	if y == 0 {
		return f,errors.New("Cannot proceed, divisor is 0")
	}
	return x / y, nil
}

// don't edit below this line

func test(x, y float64) {
	defer fmt.Println("====================================")
	fmt.Printf("Dividing %.2f by %.2f ...\n", x, y)
	quotient, err := divide(x, y)
	if x>y{
		err=errors.Join(errors.New("Dividend is greater than divisor"), err)	
	}
	
	if errors.Is(err, errors.New("Cannot proceed, divisor is 0")){
		fmt.Println("BOTH Errors same")
	}else{
	
		fmt.Println("BOTH Errors different")
	}
	err = errors.Join(errors.New("LAST ERROR"), err)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Unwrapping the error")
	err = errors.Unwrap(err)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Quotient: %.2f\n", quotient)
	return
}

func main() {
	test(10, 0)
	test(10, 2)
	test(15, 30)
	test(6, 3)
}
