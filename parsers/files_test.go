package parsers

import "testing"
import "sync"

// File pointer availalbe in the map
func TestExistingFile(t *testing.T) {
	fls := make(map[string]*File)
	fls["kaczka/dziwaczka"] = &File{Name: "kaczka/dziwaczka"}
	repo := &files{mutex: new(sync.RWMutex), files: fls}
	file := repo.Get("kaczka/dziwaczka")
	if file != fls["kaczka/dziwaczka"] {
		t.Errorf("file and files[\"kaczka/dziwaczka\"] point to different files.")
	}
}

// File pointer not available in the map
func TestNewFile(t *testing.T) {
	fls := make(map[string]*File)
	repo := &files{files: fls, mutex: new(sync.RWMutex)}
	file := repo.Get("lis/witalis")
	if file.Name != "lis/witalis" {
		t.Errorf("New File has not been created")
	}
}

// Getting the map from the files
func TestGetMap(t *testing.T) {
	fls := make(map[string]*File)
	fls["pies/pankracy"] = &File{Name: "pies/pankracy"}
	repo := &files{files: fls, mutex: new(sync.RWMutex)}
	filesMap := repo.GetMap()
	assertMapLen(filesMap, 1, t)
	dog := repo.Get("pies/pankracy")
	if dog.Name != "pies/pankracy" {
		t.Errorf("pies/pankracy does not seem to be himself!")
	}
	assertMapLen(filesMap, 1, t)
	wolf := repo.Get("wilk/i/zając")
	if wolf.Name != "wilk/i/zając" {
		t.Errorf("Wolf and rabbit do not seem to be themselves!")
	}
	assertMapLen(filesMap, 2, t)
}

// Checks number of element in map
func assertMapLen(files map[string]*File, l int, t *testing.T) {
	if len(files) != l {
		t.Errorf("There should be %d element(s) in the map - %d found", l, len(files))
	}
}
