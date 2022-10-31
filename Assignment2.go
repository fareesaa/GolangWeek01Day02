package main

import "fmt"

func main2() {
	var a int = 1
	var i int

	for i < 9 {
		i++
		var hasil int = a + i
		fmt.Printf("%d + %d  = %d", a, i, hasil)
		fmt.Println("")
	}

	fmt.Println("=================")

	var length int = 5
	var aster string

	for i := length; i > 0; i-- {
		for j := 1; j <= i; j++ {
			aster = " "
			fmt.Printf("%s", aster)
		}
		for i1 := length; i1 >= i; i1-- {
			aster = "*"
			fmt.Printf("%s", aster)
		}
		fmt.Println()
	}
}
