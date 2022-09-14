package main

//as bibliotecas necessárias
import (
	"fmt"
	"math/rand"
	"reflect"
)

// Unica função e principal
func main() {

	num := rand.Intn(24)
	frutas := []string{"abacate", "pera", "caqui", "amora", "pitanga", "goiaba", "laranja", "banana", "acerola", "melao", "maça", "pitaya", "carambola", "jabuticaba", "melancia", "uva", "morango", "abacaxi", "caju", "coco", "framboesa", "kiwi", "manga", "maracuja"}
	adivinha := []string{}
	palpite := []string{}
	ja_foram := []string{}
	letra := ""

	for _, valor := range frutas[num] {
		//fmt.Println(indice, string(valor))
		adivinha = append(adivinha, string(valor))
		palpite = append(palpite, "_")
	}

	fmt.Print("\n\tJogo da forca !\n\n")
	fmt.Print(">>> Valendo somente nomes de frutas <<<\n\n")
	fmt.Println("\t", palpite)

	for true {
		fmt.Print("\n\nInforme uma letra: ")
		fmt.Scanln(&letra)

		for i := 0; i < len(adivinha); i++ {
			if adivinha[i] == letra {
				palpite[i] = letra
			}
		}
		ja_foram = append(ja_foram, letra)
		fmt.Println("\t ._______")
		fmt.Println("\t |/      *")
		fmt.Println("\t |")
		fmt.Println("\t |")
		fmt.Println("\t", palpite)
		fmt.Println("\n\n>Letras escolhidas<", ja_foram)

		if reflect.DeepEqual(adivinha, palpite) == true {
			break
		}
	}
	fmt.Println("\n\n\t\t>>Parabéns, Você acertou !!<<")

}
