package main

import "fmt"
import "math/rand"
import "strings"

/* Genera un número aleatorio entre 1 y 3 */
func turnoMaquina() int {
	return rand.Intn(3-1+1) + 1
}

/* Pide un número al usuario, debe estar entre 1 y 3 */
func turnoUsuario() int {
	var turno int

	fmt.Printf("TU TURNO\n")
	fmt.Printf("1) PIEDRA\n")
	fmt.Printf("2) PAPEL\n")
	fmt.Printf("3) TIJERA\n")
	fmt.Printf("> ")
	fmt.Scanln(&turno)
	// TODO: Validar rango del numero
	return turno
}

/* Dibuja las manos acordes con las jugadas */
func dibujar(maquina int, usuario int) {
	if usuario == 1 && maquina == 1 {
		fmt.Printf("" +
			"TU JUGADA         " + "           MAQUINA\n" +
			"    _______       " + "        _______   \n" +
			"---'   ____)      " + "      (____   '---\n" +
			"      (_____)     " + "     (_____)      \n" +
			"      (_____)     " + "     (_____)      \n" +
			"      (____)      " + "      (____)      \n" +
			"---.__(___)       " + "       (___)__.---\n")
	}
	// TODO: terminar el resto de manos
}

/* Imprime el veredicto de la jugada, llana a la función dibujar */
func jugar(maquina int, usuario int) {

	dibujar(maquina, usuario)

	if maquina == usuario {
		fmt.Printf("EMPATE\n")
	}
	if maquina == 1 && usuario == 2 {
		fmt.Printf("GANAS\n")
	}
	if maquina == 1 && usuario == 3 {
		fmt.Printf("PIERDES\n")
	}
	if maquina == 2 && usuario == 1 {
		fmt.Printf("PIERDES\n")
	}
	if maquina == 2 && usuario == 3 {
		fmt.Printf("GANAS\n")
	}
	if maquina == 3 && usuario == 1 {
		fmt.Printf("GANAS\n")
	}
	if maquina == 3 && usuario == 2 {
		fmt.Printf("PIERDES\n")
	}
}

func main() {
	var seguir string = "s"

	fmt.Printf("===================================\n")
	fmt.Printf("====  PIEDRA - PAPEL - TIJERA  ====\n")
	fmt.Printf("===================================\n")

	for strings.EqualFold(seguir, "s") {
		jugar(turnoMaquina(), turnoUsuario())

		fmt.Printf("SEGUIR JUGANDO? (S / N) > ")
		fmt.Scanln(&seguir)
	}
}
