package command

import (
	"strings"
	"regexp"
	"fmt"
	"os"
)

func Store(line string) {
	line2 := ""
	if strings.Contains(line, "§") {
		line2 = strings.TrimSuffix(line, "§")
	}

	re1 := regexp.MustCompile(`store\s*`)
	if line2 == "" {
		line2 = re1.ReplaceAllString(line, "")
	} else {
		line2 = re1.ReplaceAllString(line2, "")
	}

	types := []string{"store", string(strings.Split(line2, " ")[0])}
	name := string(strings.Split(line2, " ")[1])

	re2 := regexp.MustCompile(`\s*=\s*`)
	value := re2.ReplaceAllString(line2, "")
	value = strings.Split(value, name)[1]
	
	re3 := regexp.MustCompile(`["']`)
	value = re3.ReplaceAllString(value, "")

	if checktype(types[1], value) {
		write(name, value, types)
	} else {
		fmt.Println("Error: invalid type")
		os.Exit(0)
	}
}