package main
import "fmt"

var totalMovimentos int

func main() {
	numeroDeDiscos := 3
	// Iniciar a torre de Hanoi com o número de discos especificado
	hanoi(numeroDeDiscos, "A", "C", "B")
	fmt.Printf("Resolvido em %d movimentos.\n", totalMovimentos)
}

func hanoi(n int, origem, destino, trabalho string) {
	if n > 0 {
		// Mover n-1 discos da origem para o trabalho
		hanoi(n-1, origem, trabalho, destino)

		// Mover o disco restante da origem para o destino
		fmt.Printf("Move peça %d: %s -> %s\n", n, origem, destino)
		totalMovimentos++

		// Mover os discos do trabalho para o destino usando a origem como espaço de trabalho
		hanoi(n-1, trabalho, destino, origem)
	}
}
