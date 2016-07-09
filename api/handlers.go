package api

import (
	"net/http"
	"fmt"
	"github.com/styutnev/gotranslate/dictionary"
)

var l = dictionary.NewLibrary()

func GetIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "index\n")
}

func AddWordPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "POST add\n")
}

func AddWordGet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "GET add\n")

}

func TranslateWord(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "translate\n")
}