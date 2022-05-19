package main

import (
	"fmt"
)

func main() {
	fmt.Printf("===================================\n")
	fmt.Printf("====  PIEDRA - PAPEL - TIJERA  ====\n")
	fmt.Printf("===================================\n")

	switch menu() {
	case 1:
		juegoUnJugador()
		break
	case 2:
		juegoMultijugador()
		break
	}
}
