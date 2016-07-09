package dictionary

import "testing"

func TestWordMarshal(t *testing.T) {
	w := Word{Title:"Cat", Locale:"En"}
	e := `{"title":"Cat","locale":"En"}`

	r, err := w.Marshal()

	if err != nil {
		t.Error("Error while adding word to dictionary")
	}

	if string(r) != e {
		t.Errorf("Values doesnt match: %v : %v ", string(r), e)
	}
}
