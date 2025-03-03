package main

import "fmt"

func goConstants() {

	const PI = 3.1415926

	const (
		A = iota
		B
		C
		D
	)

	const (
		X = iota + 10
		Y
		Z
	)

	fmt.Printf("PI : %f\n\n", PI)

	fmt.Printf("constant value %d\n\n", D)

	fmt.Printf("XYZ values %04d %04d %04d\n\n", X, Y, Z)

	const (
		READ = 1 << iota
		WRITE
		EXECUTE
	)

	permissions := READ | WRITE | EXECUTE

	fmt.Printf("PERMISSION: %b\n\n", permissions)

	fmt.Printf("Flags values \nREAD: %d\nWRITE: %d\nEXECUTE: %d\n\n", READ, WRITE, EXECUTE)

	fmt.Printf("Flags binary \nREAD: %b\nWRITE: %b\nEXECUTE: %b\n\n", READ, WRITE, EXECUTE)

	const I int = 100
	const J = 100

	k := float32(20.12)

	fmt.Println("I*J", I*J)

	fmt.Println("J*k", J*k)

}
