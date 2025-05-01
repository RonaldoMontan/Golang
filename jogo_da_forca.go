package main

// bibliotecas necessárias
import (
	"fmt"
	"math/rand"
	"reflect"
	"os/exec"
	"runtime"
	"os"
	"bufio"
	"strings"
)

func main() {

	num := rand.Intn(24)
	frutas := []string{"abacate", "pera", "caqui", "amora", "pitanga", "goiaba", "laranja", "banana", "acerola", "melao", "maça", "pitaya", "carambola", "jabuticaba", "melancia", "uva", "morango", "abacaxi", "caju", "coco", "framboesa", "kiwi", "manga", "maracuja"}
	adivinha := []string{}
	palpite := []string{}
	ja_foram := []string{}
	erradas := []string{}
	// letra := ""
	
	for _, valor := range frutas[num] {
		//fmt.Println(indice, string(valor))
		adivinha = append(adivinha, string(valor))
		palpite = append(palpite, "_")
	}
	
	fmt.Print("\n\tJogo da forca !\n\n")
	fmt.Print(">>> Valendo somente nomes de frutas <<<\n\n")
	fmt.Println("\t", palpite)
	
	for true {
		existeLetra := false

		letra := readLetter()

		if (check(ja_foram, letra) == true){
			fmt.Println("\n\n\t\t>>Essa letra já foi escolhida !<<")
			continue
		}
		
		for i := 0; i < len(adivinha); i++ {
			if adivinha[i] == letra {
				palpite[i] = letra
				existeLetra = true
			}
		}

		if !existeLetra  {
			erradas = append(erradas, letra)
		}

		ja_foram = append(ja_foram, letra)
		show(palpite, ja_foram, erradas)

		if reflect.DeepEqual(adivinha, palpite) == true {
			fmt.Println("\n\n\t\t>>Parabéns, Você acertou !!<<")
			break
		} else if len(erradas) == 5 {
			end(adivinha)
			break
		}
	}
}

func readLetter() string {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("\nInforme uma letra: ") 
    input = strings.TrimSpace(input)             // Remove espaços e \n
    return strings.ToLower(input)                // Padroniza para minúscula
}

func check(ja_foram []string, letra string) bool {	
	for _, l := range ja_foram {
		if l == letra {
			return true
		}
	}
	return false
}

func show(palpite, ja_foram , erradas []string) {
	clearTerminal()
	fmt.Print("\n\tJogo da forca !\n\n")
	fmt.Print(">>> Valendo somente nomes de frutas <<<\n\n")
	fmt.Println("\t ._______")
		fmt.Println("\t |/      *")
		fmt.Println("\t |")
		fmt.Println("\t |")
		fmt.Println("\t", palpite)
		fmt.Println("\n\n>Letras escolhidas<", ja_foram)
		fmt.Println("\n\n>Letras erradas<", erradas)
}

func end(adivinha []string) {
	fmt.Println("\t\tTentativas excedidas !")
	fmt.Println("\t\t", adivinha)
}

func clearTerminal() {

	var cmd *exec.Cmd
	
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}