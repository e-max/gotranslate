package api

import (
	"net/http"
	"fmt"
	"github.com/styutnev/gotranslate/dictionary"
	"encoding/json"
	"github.com/gorilla/mux"
	"strings"
)

type ErrorResponse struct {
	Code int `json:"code"`
	Message string `json:"message"`
}

type TranslateResponse struct {
	Result string `json:"result"`
}

type Request struct {
	Original dictionary.Word
	Translate dictionary.Word
}

var l = dictionary.NewLibrary()

func GetIndex(w http.ResponseWriter, r *http.Request) {
	w1 := dictionary.Word{Title:"cat", Locale:"en"}
	t1 := dictionary.Word{Title:"katzen", Locale:"de"}
	t2 := dictionary.Word{Title:"neko", Locale:"jp"}
	l.Add(w1, t1)
	l.Add(w1, t2)
	fmt.Fprint(w, "Translate index\n")

	return
}

func AddWordPost(w http.ResponseWriter, r *http.Request) {
	var i Request
	err := json.NewDecoder(r.Body).Decode(&i)

	if err != nil {
		panic(err)
	}

	l.AddRecursive(i.Original, i.Translate)
	w.WriteHeader(http.StatusOK)

	return
}

func TranslateWord(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	a := dictionary.Word{Title: v["title"], Locale: v["olocale"]}
	t := l.Translate(a, v["tlocale"])

	if t != nil {
		resp := TranslateResponse{Result:strings.Join(t,",")}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(resp); err != nil {
			panic(err)
		}
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		e := ErrorResponse{Code:404, Message:"Word not found"}

		if err := json.NewEncoder(w).Encode(e); err != nil {
			panic(err)
		}
	}

	return
}