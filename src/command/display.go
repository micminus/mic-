package command

import (
	"fmt"
	"strings"
	"regexp"
	"os"
)

func Display(line string) {
	todisplay := strings.Split(line, " ")
	todisplay = todisplay[1:]
	
	todisplayfinal := strings.Join(todisplay, " ")
	todisplayfinal = strings.ReplaceAll(todisplayfinal, "\n", "")

	// on check si il faut add des valeurs
	if strings.Contains(todisplayfinal, "§") {
		todisplayfinal = strings.TrimSuffix(todisplayfinal, "§")
	}

	if !strings.HasPrefix(todisplayfinal, "\"") && !strings.HasSuffix(todisplayfinal, "\"") && !strings.HasPrefix(todisplayfinal, "'") && !strings.HasSuffix(todisplayfinal, "'") {
		defer func() {
			if r := recover(); r != nil {
				// on ferme le programme
				fmt.Println("Error: missing quotes or variable unknown")
				fmt.Println("Line : " + line)
				fmt.Println("exemples : \"test\" or 'test'")
				os.Exit(0)
			}
		}()

		kk := strings.TrimSpace(todisplayfinal)
		list := get(kk)

		list[2] = strings.TrimSpace(list[2])

		if list[1] == "str" {
			todisplayfinal = strings.Replace(todisplayfinal, kk, list[2], -1)
			fmt.Println(todisplayfinal)
		} else {
			fmt.Println("error: not a string")
			os.Exit(0)
		}
	}

	if strings.Contains(todisplayfinal, "+") {
		ok := strings.Split(todisplayfinal, "+")
		re := regexp.MustCompile(`["']\s*\+\s*["']|["']\s*\+\s*|\s*\+\s*["']|\s*\+\s*`)
		todisplayfinal = re.ReplaceAllString(todisplayfinal, "")
		todisplayfinal = strings.ReplaceAll(todisplayfinal, "\n", "")

		for _, kk := range ok {
			func() {
				defer func() {
					if r := recover(); r != nil {
					}
				}()

				kk = strings.TrimSpace(kk)
				list := get(kk)

				list[2] = strings.TrimSpace(list[2])
				
				if list[1] == "str" {
					todisplayfinal = strings.Replace(todisplayfinal, kk, list[2], -1)
				} else {
					fmt.Println("Error: not a string")
					fmt.Println("Line : " + line)
					os.Exit(0)
				}
			}()
		}
	}

	if strings.Contains(todisplayfinal, "\"") {
		todisplayfinal = strings.TrimPrefix(todisplayfinal, "\"")
		todisplayfinal = strings.TrimSuffix(todisplayfinal, "\"")
		fmt.Println(todisplayfinal)
	} else if strings.Contains(todisplayfinal, "'") {
		todisplayfinal = strings.TrimPrefix(todisplayfinal, "'")
		todisplayfinal = strings.TrimSuffix(todisplayfinal, "'")
		fmt.Println(todisplayfinal)
	}
}
