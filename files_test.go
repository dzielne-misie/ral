package ral

import "testing"
import "github.com/dzielne-misie/ral/parsers"
import "fmt"

func TestExistingFile(t *testing.T) {
	files := make(map[string]*parsers.File)
	files["kaczka/dziwaczka"] = &parsers.File{Name: "kaczka/dziwaczka"}
	fmt.Println(files)
	repo := &Files{files: files}
	file := repo.Get("kaczka/dziwaczka")
	if file != files["kaczka/dziwaczka"] {
		t.Errorf("file and files[\"kaczka/dziwaczka\"] point to different files.")
	}
}

func TestNewFile(t *testing.T) {

}
