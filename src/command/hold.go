package command

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func checktype(types string, value string) bool {
	if types == "int" {
		re := regexp.MustCompile(`^[-+]?[0-9]+$`)
		return re.MatchString(value)
	}
	if types == "float" {
		re := regexp.MustCompile(`^[-+]?[0-9]+\.[0-9]+$`)
		return re.MatchString(value)
	}
	if types == "str" {
		return true
	}
	if types == "char" {
		char := strings.Split(value, "")
		if len(char) == 1 {
			return true
		} else {
			return false
		}
	}
	if types == "bool" {
		re := regexp.MustCompile(`^(true|false)$`)
		return re.MatchString(value)
	}
	if types == "list" {
		re := regexp.MustCompile(`^\[.*\]$`)
		return re.MatchString(value)
	}
	if types == "dict" {
		re := regexp.MustCompile(`^\{.*\}$`)
		return re.MatchString(value)
	}

	return false
}

func Hold(line string) {
	line2 := ""
	if strings.Contains(line, "§") {
		line2 = strings.TrimSuffix(line, "§")
	}

	re1 := regexp.MustCompile(`hold\s*`)
	if line2 == "" {
		line2 = re1.ReplaceAllString(line, "")
	} else {
		line2 = re1.ReplaceAllString(line2, "")
	}
	

	types := []string{"hold", string(strings.Split(line2, " ")[0])}
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
		fmt.Println("Line : " + line)
		os.Exit(0)
	}
}
