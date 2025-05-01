package main

// bibliotecas necessárias
import (
	"fmt"
	"math/rand"
	"reflect"
	"os/exec"
	"runtime"
	"os"
	"strings"
	"time"
)

func main() {

	num := rand.Intn(29)
	fruit := []string{"abacate", "pera", "caqui", "amora", "pitanga", "goiaba", "laranja", "banana", "acerola", "melao", "maça", "pitaya", 
						"carambola", "jabuticaba", "melancia", "uva", "morango", "abacaxi", "caju", "coco", "framboesa", "kiwi", "manga", 
						"maracuja", "ameixa", "mamao", "tangerina", "pessego", "nectarina"}
	word := []string{}
	guess := []string{}
	chosen := []string{}
	wrong := []string{}
	
	for _, value := range fruit[num] {
		word = append(word, string(value))
		guess = append(guess, "_")
	}
	
	fmt.Print("\n\tJogo da forca !\n\n")
	fmt.Print(">>> Valendo somente nomes de frutas <<<\n\n")
	fmt.Println("\t", guess)
	
	for true {
		existLetter := false

		letter := readLetter()

		if (check(chosen, letter) == true){
			fmt.Println("\n\n\t\t>>Essa letra já foi escolhida !<<")
			continue
		}
		
		for i := 0; i < len(word); i++ {
			if word[i] == letter {
				guess[i] = letter
				existLetter = true
			}
		}

		if !existLetter  {
			wrong = append(wrong, letter)
		}

		chosen = append(chosen, letter)
		show(guess, chosen, wrong)

		if reflect.DeepEqual(word, guess) == true {
			fmt.Println("\n\n\t\t>>Parabéns, Você acertou !!<<")
			break
		} else if len(wrong) == 5 {
			end(word)
			break
		}
	}
	menu()
}

func menu() {
	var option int

	fmt.Println("\n\n\t\t>>Deseja jogar novamente ?<<")
	fmt.Println("\t\t1 - Sim")
	fmt.Println("\t\t2 - Não")
	fmt.Scanln(&option)

	if option == 1 {
		clearTerminal()
		main()
	} else if option == 2 {
		clearTerminal()
		fmt.Println("\t\t>>Obrigado por jogar<<")
		time.Sleep(10 * time.Second)
	} else {
		clearTerminal()
		fmt.Println("\t\t>>Opção inválida<<")
		menu()
	}
}

func readLetter() string {
	var input string
    fmt.Print("\nInforme uma letra: ")
	fmt.Scanln(&input)
    input = strings.TrimSpace(input)             // Remove espaços e \n
    return strings.ToLower(input)                // Padroniza para minúscula
}

func check(chosen []string, letter string) bool {	
	for _, l := range chosen {
		if l == letter {
			return true
		}
	}
	return false
}

func show(guess, chosen , wrong []string) {
	clearTerminal()
	fmt.Print("\n\tJogo da forca !\n\n")
	fmt.Print(">>> Valendo somente nomes de frutas <<<\n\n")
	fmt.Println("\t ._______")
		fmt.Println("\t |/      *")
		fmt.Println("\t |")
		fmt.Println("\t |")
		fmt.Println("\t", guess)
		fmt.Println("\n\n>Letras escolhidas<", chosen)
		fmt.Println("\n\n>Letras erradas<", wrong)
}

func end(word []string) {
	fmt.Println("\n\t\tTentativas excedidas !")
	fmt.Println("\t\t", word)
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