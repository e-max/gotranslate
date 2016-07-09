// Package provides structures and methods for different kinds of dictionaries
package dictionary

// Interface for collectable structures
type Collector interface {
	Add(Word, Word)
	Exists(Word) bool
	Find(Word) ([]Word, error)
	Translate(string, string) []string
}

// Dictionary is mapping for words
// One word can be mapped to different words (translations, meanings, etc.)
type Dictionary struct {
	Translations map[Word][]Word
}

// Create new Dictionary
func NewDictionary() *Dictionary {
	d := new(Dictionary)
	d.Translations = make(map[Word][]Word)

	return d
}

// Check if words is mapped
func (d *Dictionary) Exists(o Word) bool  {
	_, ok := d.Translations[o]

	return ok
}

// Add word translation
// If translation exists - return error
func (d *Dictionary) Add(o Word, t Word) {
	if v, ok := d.Translations[o]; ok {
		// Mapping was found - values must not contain duplicates
		f := false
		for _, ws := range v {
			if ws.Title == t.Title {
				f = true
			}
		}

		if !f {
			d.Translations[o] = append(d.Translations[o], t)
		}
	} else {
		// No mapping was found
		d.Translations[o] = append(d.Translations[o], t)
	}
}

// Find word translations
// If there is no translation for provided word - return error
func (d *Dictionary) Find(w Word) []Word {
	if val, ok := d.Translations[w]; ok {
		return val
	}

	return nil
}

// Translate word into locale
// If there is no translation - return error
func (d *Dictionary) Translate(w Word, l string) []string {
	r := []string{}

	if val, ok := d.Translations[w]; ok {
		for _, e := range val {
			if e.Locale == l {
				r = append(r, e.Title)
			}
		}

		return r
	}

	return nil
}
