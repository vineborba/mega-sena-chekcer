package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Insira a sequência ganhadora, separando os numeros por espaço:")
	scanner.Scan()
	input := scanner.Text()
	win := strings.Fields(input)

	for {
		fmt.Println("Insira a sequência a verificar, separando os numeros por espaço")
		fmt.Println("Digite 0 para sair")
		scanner.Scan()
		text := scanner.Text()

		if len(strings.Fields(text)) < 6 {
			break
		}

		inputed := strings.Fields(text)
		hits := 0
		for _, i := range inputed {
			for _, n := range win {
				if i == n {
					hits++
				}
			}
		}
		fmt.Printf("Jogo %s, acertos %d\n", text, hits)
	}
}
