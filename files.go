package ral

import "github.com/dzielne-misie/ral/parsers"

// Files struct helps to make sure that only unique parsers.File objects exist
// (parsers.File is considered unique as long as its Name - a path - is unique)
type Files struct {
	files map[string]*parsers.File
}

// Function gets the file from the repository or creates a new instance (if there is not file with such a name registered)
func (f *Files) Get(name string) *parsers.File {
	file := f.files[name]
	if file.Name == "" {
		file = &parsers.File{Name: name}
		f.files[name] = file
	}
	return file
}
