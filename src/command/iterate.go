package command

import (
	"fmt"
	"strings"
	"regexp"
	"strconv"
	"os"
)

var info = make(map[string]string)

func End() {
	if info["type"] == "nb" {
		k, err := strconv.Atoi(info["nb_of_iter"])

		if err != nil {
			fmt.Println("Error during converting to int the number of iterations")
			os.Exit(0)
		}

		monophase := strings.Split(info["var"], "as")

		jsp := get(strings.TrimSpace(monophase[0]))
		if jsp != nil {
			fmt.Println("Error: variable " + strings.TrimSpace(monophase[0])+ " already exists")
			os.Exit(0)
		}

		for i := 0; i < k; i++ {
			write(strings.TrimSpace(monophase[0]), strconv.Itoa(i), []string{"store", strings.TrimSpace(monophase[1])})
			for _, j := range please() {
				common(j)
			}
		}

		delete(strings.TrimSpace(monophase[0]))
		liste = []string{}
	} else if info["type"] == "cond" {
		fmt.Println("Error: not implemented")
		os.Exit(0)
	}
}

func Iterate(line string) {
	line2 := ""
	etime := []int{}
	etimee := ""
	if strings.Contains(line, "§") {
		line2 = strings.TrimSuffix(line, "§")
	}

	re1 := regexp.MustCompile(`iterate\s*`)
	if line2 == "" {
		line2 = re1.ReplaceAllString(line, "")
	} else {
		line2 = re1.ReplaceAllString(line2, "")
	}

	execution_time := strings.Split(line2, "with")[0]
	execution_time = strings.TrimSpace(execution_time)

	variable := ""
	defer func() {
			if r := recover(); r != nil {
				info["type"] = "cond"
				info["nb_of_iter"] = strconv.Itoa(len(etime))
			}
		}()

	variable = strings.Split(line2, "with")[1]
	variable = strings.TrimSpace(variable)
	variable = strings.TrimSuffix(variable, "{")
	variable = strings.TrimSpace(variable)

	vare := variable

	if strings.HasPrefix(execution_time, "range") {
		execution_time = strings.TrimPrefix(execution_time, "range(")
		l := strings.Split(execution_time, ",")[1]
		t, err := strconv.Atoi(strings.TrimSpace(strings.ReplaceAll(l, ")", "")))
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		s, err2 := strconv.Atoi(strings.TrimSpace(strings.Split(execution_time, ",")[0]))
		if err2 != nil {
			fmt.Println(err2)
		}

		for i := s; i < t+1; i++ {
			etime = append(etime, i)
		}

		info["type"] = "nb"
		etimee = strconv.Itoa(len(etime))
	} else if strings.HasPrefix(execution_time, "len") {
		execution_time = strings.TrimPrefix(execution_time, "len(")
		execution_time = strings.TrimSuffix(execution_time, ")")

		t := get(execution_time)
		if t[1] == "int" {
			etimee = t[2]
		} else if t[1] == "str" {
			etimee = strconv.Itoa(len(t[2]))
		} else if t[1] == "list" {
			etimee = strconv.Itoa(len(t[2]))
		} else if t[1] == "dict" {
			etimee = strconv.Itoa(len(t[2]))
		} else if t[1] == "bool" {
			fmt.Println("Error: bool type doesn't have len")
			os.Exit(0)
		} else if t[1] == "char" {
			etimee = strconv.Itoa(len(t[2]))
		} else if t[1] == "float" {
			fmt.Println("Error: float type doesn't have len")
			os.Exit(0)
		}

		info["type"] = "nb"
	}

	info["var"] = vare
	info["nb_of_iter"] = etimee
}