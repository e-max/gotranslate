package dictionary

import "testing"

func TestLibraryAdd(t *testing.T) {
	d := NewLibrary()

	w := Word{Title:"Cat", Locale:"En"}
	t1 := Word{Title:"Katzen", Locale:"De"}
	t2 := Word{Title:"Neko", Locale:"Jp"}

	d.Add(w, t1)
	d.Add(w, t2)

	if len(d.Dictionaries[w.Locale].Translations) != 1 {
		t.Error("Add failed")
	}
}

func TestLibraryAddRecursive(t *testing.T) {
	d := NewLibrary()

	w1 := Word{Title:"Cat", Locale:"En"}
	t1 := Word{Title:"Katzen", Locale:"De"}
	t2 := Word{Title:"Neko", Locale:"Jp"}

	d.AddRecursive(w1, t1) // Cat[En] => Katzen[De]
	d.AddRecursive(t2, w1) // Neko[Jp] => Cat[En]
	d.AddRecursive(t1, t2) // Katzen[De] => Neko[Jp]

	if len(d.Dictionaries[w1.Locale].Translations[w1]) != 2 {
		t.Error("Add failed")
	}
}

func TestLibraryFind(t *testing.T) {
	d := NewLibrary()
	w := Word{Title:"Cat", Locale:"En"}
	t1 := Word{Title:"Katzen", Locale:"De"}
	t2 := Word{Title:"Neko", Locale:"Jp"}
	t3 := Word{Title:"Gato", Locale:"Es"}

	d.Add(w, t1)
	d.Add(t2, w)
	d.Add(t1, t2)

	ws := d.Find(w)

	if len(ws) != 1 {
		t.Errorf("Not all results were found for %v - %v", w.Title, len(ws))
	}

	ws = d.Find(t1)

	if len(ws) != 1 {
		t.Errorf("Not all results were found for %v - %v", w.Title, len(ws))
	}

	ws = d.Find(t3)

	if ws != nil {
		t.Errorf("Not all results were found for %v - %v", w.Title, len(ws))
	}
}

func TestLibraryTranslate(t *testing.T) {
	d := NewLibrary()
	w1 := Word{Title:"Cat", Locale:"En"}
	w2 := Word{Title:"Kot", Locale:"Ru"}
	t1 := Word{Title:"Katzen", Locale:"De"}
	t2 := Word{Title:"Neko", Locale:"Jp"}
	t3 := Word{Title:"Gato", Locale:"Es"}


	d.AddRecursive(w1, t1) // Cat[En] => Katzen[De]
	d.AddRecursive(t1, t2) // Katzen[De] => Neko[Jp]
	d.AddRecursive(w1, w2) // Cat[En] => Kot[ru]

	ws := d.Translate(w1, t1.Locale)

	if len(ws) != 1 {
		t.Error("Not all results were found")
	}

	ws = d.Translate(t1, w1.Locale)

	if len(ws) != 1 {
		t.Error("Not all results were found")
	}

	if ws[0] != w1.Title {
		t.Error("Not all results were found")
	}

	ws = d.Translate(t1, t2.Locale)

	if len(ws) != 1 {
		t.Error("Not all results were found")
	}

	if ws[0] != t2.Title {
		t.Error("Not all results were found")
	}

	ws = d.Translate(t3, t2.Locale)

	if ws != nil {
		t.Error("Not all results were found")
	}
}
