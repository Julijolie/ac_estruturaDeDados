package main
import "fmt"

var totalMovimentos int

func main() {
	numeroDeDiscos := 3
	hanoi(numeroDeDiscos, "A", "C", "B")
	fmt.Printf("Resolvido em %d movimentos.\n", totalMovimentos)
}

func hanoi(n int, origem, destino, trabalho string) {
	if n > 0 {

		hanoi(n-1, origem, trabalho, destino)

		fmt.Printf("Move peça %d: %s -> %s\n", n, origem, destino)
		totalMovimentos++

		hanoi(n-1, trabalho, destino, origem)
	}
}
