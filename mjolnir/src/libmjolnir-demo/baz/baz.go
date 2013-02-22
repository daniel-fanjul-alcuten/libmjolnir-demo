package baz

import (
	"fmt"
	. "github.com/daniel-fanjul-alcuten/libmjolnir"
	. "libfoobar/bar"
	"path/filepath"
)

type Baz struct {
	Files []*CFile
}

func NewBaz(mjölnir *Mjölnir, basepath string, bar *Bar) *Baz {
	path, files := filepath.Join(basepath, "baz"), make([]*CFile, 50)
	for i := range files {
		name := fmt.Sprint("baz", i+1, ".c")
		files[i] = mjölnir.CFile(filepath.Join(path, name)).SetStd("c99").DependsOn(bar.Files...)
	}
	return &Baz{files}
}
