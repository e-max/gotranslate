package dictionary

import "testing"

func TestDictionary(t *testing.T) {
	d := NewDictionary()

	if len(d.Translations) > 0 {
		t.Error("Empty dict is not empty")
	}
}

func TestDictionaryExists(t *testing.T) {
	d := NewDictionary()
	o1 := Word{Title:"Cat", Locale:"En"}
	tr := Word{Title:"Katzen", Locale:"De"}

	e := d.Exists(o1)

	if e != false {
		t.Error("Check if word exists failed")
	}

	d.Add(o1, tr)

	e = d.Exists(o1)

	if e != true {
		t.Error("Check if word exists failed")
	}
}

func TestDictionaryAdd(t *testing.T) {
	d := NewDictionary()
	o1 := Word{Title:"Cat", Locale:"En"}
	o2 := Word{Title:"Cat", Locale:"En"}
	tr := Word{Title:"Katzen", Locale:"De"}

	d.Add(o1, tr)

	if len(d.Translations) != 1 {
		t.Error("Word was not added")
	}

	// Adding same word to dictionary must not change dictionary size
	d.Add(o2, tr)

	if len(d.Translations) != 1 {
		t.Error("Word was not added")
	}
}

func TestDictionaryFind(t *testing.T) {
	d := NewDictionary()
	w1 := Word{Title:"Cat", Locale:"En"}
	w2 := Word{Title:"Кот", Locale:"Ru"}
	t1 := Word{Title:"Katzen", Locale:"De"}
	t2 := Word{Title:"Neko", Locale:"Jp"}

	d.Add(w1, t1)
	d.Add(w1, t2)

	ws := d.Find(w1)

	if len(ws) != 2 {
		t.Error("Not all results were found")
	}

	ws = d.Find(w2)

	if len(ws) != 0 {
		t.Error("Not all results were found")
	}
}

func TestDictionaryTranslate(t *testing.T) {
	d := NewDictionary()
	w1 := Word{Title:"Cat", Locale:"En"}
	w2 := Word{Title:"Кот", Locale:"Ru"}
	t1 := Word{Title:"Katzen", Locale:"De"}
	t2 := Word{Title:"Neko", Locale:"Jp"}
	d.Add(w1, t1)
	d.Add(w1, t2)

	ws := d.Translate(w1, "De")

	if len(ws) != 1 {
		t.Error("Not all results were found")
	}

	ws = d.Translate(w2, "En")

	if ws != nil {
		t.Error("Translation should be found")
	}
}
