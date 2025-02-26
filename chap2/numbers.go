package main

import "fmt"

func numbers() {

	c1 := 1 + 7i
	c2 := complex(2, 5)

	fmt.Printf("complexe c1: %T\n", c1)
	fmt.Printf("complexe c2: %T\n", c2)

	c3 := complex64(c1 + c2)

	fmt.Printf("complexe c3: %T\n", c3)

	var a int = 9

	var b int = 7

	c := int64(a + b)

	fmt.Printf("c is a %T\n", c)

	d := a / b

	e := float64(a) / float64(b)

	fmt.Printf("d (%d) is a %T\n", d, d)

	fmt.Printf("e (%.3f) is a %T\n", e, e)

}
