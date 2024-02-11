package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"os/exec"
	"runtime"

	"./command"
)

var space = ""


func ChangeTab(line string) string {
	line2 := line
	for {
		if !strings.HasPrefix(line2, " ") {
			break
		}
		space += " "
		line2 = strings.TrimPrefix(line2, " ")
	}

	return line
}

func Launch(filepath string) {
	tokens := map[string]func(string){
		"display": func(input string) { command.Display(input) },
		"get":     func(input string) { command.Get(input) },
		"hold":    func(input string) { command.Hold(input) },
		"store":   func(input string) { command.Store(input) },
		"set":     func(input string) { command.Set(input) },
		"iterate": func(input string) { command.Iterate(input) },
		"§":       func(input string) {},
	}
	ct_tokens := []string{
		"iterate",
	}

	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var shouldchangetab bool
	for scanner.Scan() {
		line := scanner.Text()

		if shouldchangetab {
			ChangeTab(line)
			shouldchangetab = false
		}

		if strings.HasPrefix(line, "}") {
			space = ""
			command.End()
		}

		for token, function := range tokens {
			if strings.HasPrefix(line, space + token) {
				if space == "" {
					function(line) // Appel de la fonction associée au token

					for _, ctoken := range ct_tokens {
						if ctoken == token {
							shouldchangetab = true
						}
					}

				} else {
					command.List(line, space)
				}
			}

		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Erreur lors de la lecture du fichier:", err)
	}
}

func main() {
	command := os.Args[1]
	args := []string{}
	for i := 2; i < len(os.Args); i++ {
		args = append(args, os.Args[i])
	}
	switch command {
		case "run":
			Launch(args[0])
		case "help":
			fmt.Println("help - Affiche l'aide")
			fmt.Println("run [file] - Interprete le fichier")
			fmt.Println("doc - Ouvre la documentation")
		case "doc":
			// on ouvre le fichier doc.html
			filename := "C:/Users/valen/Documents/test/eee/langage_en_cours/doc/doc.html"
			var err error
			switch runtime.GOOS {
			case "linux":
				err = exec.Command("xdg-open", filename).Start()
			case "windows":
				err = exec.Command("cmd", "/C", "start", "", filename).Start()
			case "darwin":
				err = exec.Command("open", filename).Start()
			default:
				err = exec.Command("xdg-open", filename).Start()
			}
			if err != nil {
				panic(err)
			}
		default:
			fmt.Println("Commande inconnue")
			fmt.Println("help pour afficher l'aide")
			os.Exit(0)
	}
}