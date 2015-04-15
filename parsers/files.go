package parsers

import "sync"

// Files struct helps to make sure that only unique File objects exist
// (File is considered unique as long as its Name - a path - is unique)
type files struct {
	mutex *sync.RWMutex
	files map[string]*File
}

// This is a constructor, of some sorts ;) We need to be able
func NewFiles() *files {
	return &files{mutex: new(sync.RWMutex), files: make(map[string]*File)}
}

// Function gets the file from the repository or creates a new instance (if there is not file with such a name registered)
func (f *files) Get(name string) *File {
	f.mutex.Lock()
	_, e := f.files[name]
	if e == false {
		f.files[name] = &File{Name: name}
	}
	file := f.files[name]
	f.mutex.Unlock()
	return file
}

// Retrieves all the exiting files.
func (f *files) GetMap() map[string]*File {
	f.mutex.RLock()
	files := f.files
	f.mutex.RUnlock()
	return files
}
