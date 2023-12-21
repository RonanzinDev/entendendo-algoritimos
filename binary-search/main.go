package main

import "fmt"

// A busca por um elemento começa pelo meio do slice. Faz mais sentido do que a busca convencional(com for) quando o tamanho do slide for maior. Invez de ficar passando por cada elemento do slice como na forma tradicional, começamos a busca no meio da lista(ordenada), caso o elemento desejado não seja esse que esta na metade, verificamos se é maiorou menor. Se o elemento da metade for maior que o elemento desejado, logo ficamos apenas com os numeros menor daquele que estava na metade.
// Por exemplo:
// Lista = [1,2,3,4,5,6,7,8,9,10]
// Numero desejado = 3
// Começariasmo pela metade -> 5
// 5 é maior que 3. Logo a nova lista sera [1,2,3,4]
// Ai pela metada de novo -> 2
// 3 é menor que 2. Logo a nova lista sera = [3,4]
// Pela metada de novo -> 3
// Pronto achamos. Pelo metodo tradicional fariamos 5 buscas passando por cada numero (1,2,3,4,5), já pela busca binaria fizemos apenas 3. Log 2 n

func binarySearch(lista []int, numero int) int {
	inicio := 0
	final := len(lista) - 1
	for inicio <= final {
		meio := (inicio + final) / 2
		numeroAtual := lista[meio]
		if numeroAtual == numero {
			return numeroAtual
		}
		if numeroAtual >= numero {
			final = meio - 1
		}
		if numeroAtual < numero {
			inicio = meio + 1
		}
	}
	return -1
}

func main() {
	items := []int{1, 2, 3, 4, 5, 6, 7, 9, 10}
	fmt.Println(binarySearch(items, 3))
}
