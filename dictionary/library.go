package dictionary

// Library is mapping locales with dictionaries
type Library struct {
	Dictionaries map[string]*Dictionary
}

// Check if dictionary for locale exists
func (d *Library) Exists(locale string) bool {
	_, ok := d.Dictionaries[locale]

	return ok
}

// Appends new dictionary
func (d *Library) Append(locale string, dict *Dictionary) {
	d.Dictionaries[locale] = dict
}

// Add original and translation to dictionary
// Values will be added to two dictionaries to keep backward compatibility
func (d *Library) Add(o Word, t Word) {
	if false == d.Exists(o.Locale) {
		dict := &Dictionary{Translations: map[Word][]Word{}}
		d.Append(o.Locale, dict)
	}

	tr := d.Dictionaries[o.Locale]
	tr.Add(o, t)
}

// Add words to dictionary
func (d *Library) AddRecursive(o Word, t Word) {
	d.Add(o, t)
	r := d.Find(o)

	for _, w1 := range r {
		d.Add(w1, o)
		for _, w2 := range r {
			if w1 != w2 {
				d.Add(w1, w2)
			}
		}
	}
}

// Find all translations for word
func (d *Library) Find(w Word) []Word {
	if v, ok := d.Dictionaries[w.Locale]; ok {
		return v.Find(w)
	}

	return nil
}

// Translate word according to provided locale
func (d *Library) Translate(w Word, l string) []string {
	if v, ok := d.Dictionaries[w.Locale]; ok {
		tr := v.Translate(w, l)
		return tr
	}

	return nil
}
