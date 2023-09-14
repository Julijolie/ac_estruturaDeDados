package main

import (
    "fmt"
)

func main() {
    arr := []int{1, 2, 3, 4, 5, 6, 7}
    alvo := 9

    esquerda := 0
    direita := len(arr) - 1

    for esquerda < direita {
        soma := arr[esquerda] + arr[direita]

        if soma == alvo {
            fmt.Printf("Par encontrado: %d, %d\n", arr[esquerda], arr[direita])
            return
        } else if soma < alvo {
            esquerda++
        } else {
            direita--
        }
    }

    fmt.Println("(-1,-1)")
}