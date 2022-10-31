// You can edit this code!
// Click here and start typing.
package main

// Importing fmt
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Calling main
func main4() {
	scanner := bufio.NewScanner(os.Stdin)

	var to string
	var message string
	var lenghtStr int
	var pil int

	fmt.Print("To: ")
	fmt.Scanf("%s", &to)
	fmt.Print("Message: ")
	scanner.Scan()
	message = scanner.Text()
	fmt.Println("Menu")
	fmt.Println("1.Kirim Pesan!")
	fmt.Println("2.Batalkan Pesan!")
	fmt.Println("")
	fmt.Print("Silahkan pilih menu: ")
	fmt.Scanf("%d", &pil)
	fmt.Println("====================")

	switch pil {
	case 1:

		fmt.Println("Message: ", message)
		lenghtStr = len(message)
		fmt.Printf("Jumlah karakter pada pesan: %d", lenghtStr)
		nofA := strings.Count(message, "a")
		nofB := strings.Count(message, "b")
		nofC := strings.Count(message, "c")
		nofD := strings.Count(message, "d")
		nofE := strings.Count(message, "e")
		nofF := strings.Count(message, "f")
		nofG := strings.Count(message, "g")
		nofH := strings.Count(message, "h")
		nofI := strings.Count(message, "i")
		nofJ := strings.Count(message, "j")
		nofK := strings.Count(message, "k")
		nofL := strings.Count(message, "l")
		nofM := strings.Count(message, "m")
		nofN := strings.Count(message, "n")
		nofO := strings.Count(message, "o")
		nofP := strings.Count(message, "p")
		nofQ := strings.Count(message, "q")
		nofR := strings.Count(message, "r")
		nofS := strings.Count(message, "s")
		nofT := strings.Count(message, "t")
		nofU := strings.Count(message, "u")
		nofV := strings.Count(message, "v")
		nofW := strings.Count(message, "w")
		nofX := strings.Count(message, "x")
		nofY := strings.Count(message, "y")
		nofZ := strings.Count(message, "z")
		nofSpace := strings.Count(message, " ")

		fmt.Println("")
		fmt.Println("A = ", nofA)
		fmt.Println("B = ", nofB)
		fmt.Println("C = ", nofC)
		fmt.Println("D = ", nofD)
		fmt.Println("E = ", nofE)
		fmt.Println("F = ", nofF)
		fmt.Println("G = ", nofG)
		fmt.Println("H = ", nofH)
		fmt.Println("I = ", nofI)
		fmt.Println("J = ", nofJ)
		fmt.Println("K = ", nofK)
		fmt.Println("L = ", nofL)
		fmt.Println("M = ", nofM)
		fmt.Println("N = ", nofN)
		fmt.Println("O = ", nofO)
		fmt.Println("P = ", nofP)
		fmt.Println("Q = ", nofQ)
		fmt.Println("R = ", nofR)
		fmt.Println("S = ", nofS)
		fmt.Println("T = ", nofT)
		fmt.Println("U = ", nofU)
		fmt.Println("V = ", nofV)
		fmt.Println("W = ", nofW)
		fmt.Println("X = ", nofX)
		fmt.Println("Y = ", nofY)
		fmt.Println("Z = ", nofZ)
		fmt.Println("Space = ", nofSpace)

		fmt.Println("Pesan Terkirim!")
		break

	default:
		fmt.Println("Pesan diurungkan!")
		break
	}
}
