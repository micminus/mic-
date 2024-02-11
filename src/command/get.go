package command

import (
	"os"
	"bufio"
	"strings"
)

func Get(line string) {
	reader := bufio.NewReader(os.Stdin) 		// prend l'input
	value, _ := reader.ReadString('\n')			// lit l'input

	line = strings.TrimPrefix(line, "get ")
	values := strings.Split(line, " ")
	name := string(values[1])

	name = strings.TrimPrefix(name, "'")
	name = strings.TrimSuffix(name, "'")

	types := []string{"hold", string(values[0])}

	write(name, value, types)
}