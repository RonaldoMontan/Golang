package main

import (
	"fmt"
	"math/rand"
	"reflect"
)

func main() {

	num := rand.Intn(7)
	frutas := []string{"abacate", "pera", "laranja", "melao", "maça", "uva", "morango"}
	adivinha := []string{}
	palpite := []string{}
	letra := ""

	for _, valor := range frutas[num] {
		//fmt.Println(indice, string(valor))
		adivinha = append(adivinha, string(valor))
		palpite = append(palpite, "_")
	}

	fmt.Print("\n\tJogo da forca !\n\n")
	fmt.Print(">>> Valendo somente nomes de frutas <<<\n\n")
	fmt.Println(palpite)

	for true {
		fmt.Println("\n\nInforme uma letra: ")
		fmt.Scanln(&letra)

		for i := 0; i < len(adivinha); i++ {
			if adivinha[i] == letra {
				palpite[i] = letra
			}
		}
		fmt.Println(palpite)

		if reflect.DeepEqual(adivinha, palpite) == true {
			break
		}
	}
	fmt.Println("\n\nParabéns, Você acertou !!")

}
