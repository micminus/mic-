package command

import (
	"strings"
	"regexp"
	"fmt"
	"os"
)

func Set(line string) {
	line2 := ""
	if strings.Contains(line, "§") {
		line2 = strings.TrimSuffix(line, "§")
	}

	re1 := regexp.MustCompile(`set\s*`)
	if line2 == "" {
		line2 = re1.ReplaceAllString(line, "")
	} else {
		line2 = re1.ReplaceAllString(line2, "")
	}

	name := string(strings.Split(line2, " ")[0])
	nvalue := string(strings.Split(line2, " ")[2])

	re2 := regexp.MustCompile(`["']`)
	nvalue = re2.ReplaceAllString(nvalue, "")

	values := get(name)

	if values[0] == "store" {
		fmt.Println("Error: can't set a store variable")
		fmt.Println("Line : " + line)
		os.Exit(0)
	}

	write(name, nvalue, values)
}