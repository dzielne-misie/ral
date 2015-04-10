package parsers

// Files struct helps to make sure that only unique File objects exist
// (File is considered unique as long as its Name - a path - is unique)
type files struct {
	ch    chan *File
	files map[string]*File
}

// This is a constructor, of some sorts ;) We need to be able
func NewFiles(ch chan *File) *files {
	return &files{ch: ch, files: make(map[string]*File)}
}

// Function gets the file from the repository or creates a new instance (if there is not file with such a name registered)
func (f *files) Get(name string) {
	_, e := f.files[name]
	//	fmt.Println(file, e)
	if e == false {

		f.files[name] = &File{Name: name}
		//		fmt.Println("!!!!!!!!", f.files[name].Name)
		//		f.files[name] = file
	}
	//	fmt.Println("~~~~~~", f.files[name].Name)
	f.ch <- f.files[name]
}

// Retrieves all the exiting files. Attention: this function is not goroutine safe! (maps are not - https://blog.golang.org/go-maps-in-action#TOC_6.)
func (f *files) GetMap() *map[string]*File {
	return &f.files
}
