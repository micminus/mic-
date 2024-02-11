package command

import (
	"strings"
	//"fmt"
)

var variables = make(map[string][]string)
var liste = make([]string, 0)

func write(name string, value string, types []string) {
	/*
	Donner le nom, le type et les valeurs pour quelles soit enregistrer
	-----------------------------
	Args :
		name : le nom de la variable (ex : x) => string
		value : la valeur de la variable (ex : 5) => string
		types : le type de la variable (ex : ["store", "int"]) => []string 
	-----------------------------
	Exemple :
		write("x", "5", ["store", "int"]) => x = ["store", "int", "5"]
	-----------------------------
	Rappel :
		Store => variable constante
		Hold => variable non constante

		Type de variable :
			int => entier
			str => chaine de caractère
			char => caractère
			bool => booleen
			list => liste
			dict => dictionnaire
	*/

	variables[name] = []string{types[0], types[1], value} 		// ex : x = ["store", "int", "5"]
}

func delete(name string) {
	variables[name] = nil
}

func get(name string) []string {
	/*
	Donner le nom de la variable pour obtenir les valeurs associé
	---------------------------------
	Args :
		name : le nom de la variable (ex: x) => string
	---------------------------------
	Retour :
		[]string => ["store", "int", "5"]
	---------------------------------
	Exemple :
		get("x") => ["store", "int", "5"]
	---------------------------------
	Rappel :
		Premier élément : constante ou non
		Deuxième élément : type de la variable
		Troisième élément : valeur
	*/

	return variables[name]
}

func List(line string, tab string) {
	line = strings.TrimPrefix(line, tab)

	liste = append(liste, line)
}

func common(line string) {
	tokens := map[string]func(string){
		"display": func(input string) { Display(input) },
		"get":     func(input string) { Get(input) },
		"hold":    func(input string) { Hold(input) },
		"store":   func(input string) { Store(input) },
		"set":     func(input string) { Set(input) },
		"§":       func(input string) {},
	}

	for token, function := range tokens {
		if strings.HasPrefix(line, token) {
			function(line) // Appel de la fonction associée au token
		}
	}
}

func please() []string {
	return liste
}