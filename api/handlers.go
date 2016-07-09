package api

import (
	"net/http"
	"fmt"
	"github.com/styutnev/gotranslate/dictionary"
)

var l dictionary.Library

func GetIndex(w http.ResponseWriter, r *http.Request) {
	l.Add(dictionary.Word{Title:"Cat", Locale:"En"}, dictionary.Word{Title:"Kot", Locale:"Ru"})
	fmt.Fprint(w, "index\n")
}

func AddWord(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "add\n")
}

func TranslateWord(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "translate\n")
}