package parsers

import "testing"

// File pointer availalbe in the map
func TestExistingFile(t *testing.T) {
	ch := make(chan *File)
	fls := make(map[string]*File)
	fls["kaczka/dziwaczka"] = &File{Name: "kaczka/dziwaczka"}
	repo := &files{files: fls, ch: ch}
	go repo.Get("kaczka/dziwaczka")
	file := <-ch
	if file != fls["kaczka/dziwaczka"] {
		t.Errorf("file and files[\"kaczka/dziwaczka\"] point to different files.")
	}
}

// File pointer not available in the map
func TestNewFile(t *testing.T) {
	ch := make(chan *File)
	fls := make(map[string]*File)
	repo := &files{files: fls, ch: ch}
	go repo.Get("lis/witalis")
	file := <-ch
	if file.Name != "lis/witalis" {
		t.Errorf("New File has not been created")
	}
}

// Getting the map from the files
func TestGetMap(t *testing.T) {
	ch := make(chan *File)
	fls := make(map[string]*File)
	fls["pies/pankracy"] = &File{Name: "pies/pankracy"}
	repo := &files{files: fls, ch: ch}
	filesMap := repo.GetMap()
	assertMapLen(*filesMap, 1, t)
	go repo.Get("pies/pankracy")
	dog := <-ch
	if dog.Name != "pies/pankracy" {
		t.Errorf("pies/pankracy does not seem to be himself!")
	}
	assertMapLen(*filesMap, 1, t)
	go repo.Get("wilk/i/zając")
	wolf := <-ch
	if wolf.Name != "wilk/i/zając" {
		t.Errorf("Wolf and rabbit do not seem to be themselves!")
	}
	assertMapLen(*filesMap, 2, t)
}

// Checks number of element in map
func assertMapLen(files map[string]*File, l int, t *testing.T) {
	if len(files) != l {
		t.Errorf("There should be %d element(s) in the map - %d found", l, len(files))
	}
}
