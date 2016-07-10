package api

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"log"
	"io/ioutil"
	"os"
)

var routerTests = []struct {
        name string
	route string
	method string
	code int
}{
        {"Index", "/", "GET", 200},
        {"AddWord", "/add","POST", 200},
        {"TranslateWord", "/translate/en/word", "GET", 200},
        {"TranslateWord", "/translate/badlocale/word", "GET", 404},
        {"SomeExplicitRouter", "/this/is/bad/route", "GET", 404},
        {"SomeExplicitRouter", "/this/is/bad/route", "POST", 404},
}

// Entry point for testing
func TestMain(m *testing.M) {
    setUp()
    code := m.Run()
    os.Exit(code)
}

// Disable log output while testing
func setUp() {
	log.SetOutput(ioutil.Discard)

}

func TestCreateRouter(t *testing.T) {
	r := CreateRouter()

	for _, tt := range routerTests {
		a, e := http.NewRequest(tt.method, tt.route, nil)

		if  e != nil {
			t.Errorf("Failed fetching %s", tt.name)
		}

		w := httptest.NewRecorder()
		r.ServeHTTP(w, a)

		if w.Code != tt.code {
			t.Errorf("Incorrect %s response: expected %d got %d", tt.name, tt.code, w.Code)
		}

		t.Log()
        }
}
