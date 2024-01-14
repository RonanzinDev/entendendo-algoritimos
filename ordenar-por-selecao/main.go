package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

func buscarMenor(lista []int) int {
	menor := lista[0]
	menor_indice := 0
	for i := range lista {
		if lista[i] < menor {
			menor = lista[i]
			menor_indice = i
		}
	}
	return menor_indice
}

func ordenarPorSelecao(lista []int) []int {
	var novoArr []int
	for range lista {
		menor := buscarMenor(lista)
		novoArr = append(novoArr, lista[menor])
		lista = append(lista[:menor], lista[menor+1:]...)
	}
	return novoArr
}

func main() {
	t, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}

	err = trace.Start(t)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()
	lista := []int{5, 3, 6, 2, 10}
	fmt.Println(ordenarPorSelecao(lista))
}
