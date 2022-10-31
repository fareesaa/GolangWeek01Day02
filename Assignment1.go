package main

import "fmt"

func main() {

	var roda int = 2

	switch roda {
	case 2:
		fmt.Printf("Roda %d: Motor", roda)
		break
	case 4:
		fmt.Printf("Roda %d: Mobil", roda)
		break
	default:
		fmt.Printf("Roda %d: Truk atau Bis", roda)
		break
	}
}
