package bar

import (
	"fmt"
	. "github.com/daniel-fanjul-alcuten/libmjolnir"
	. "libfoobar/foo"
	"path/filepath"
)

type Bar struct {
	Files []*CFile
}

func NewBar(mjölnir *Mjölnir, basepath string, foo *Foo) *Bar {
	path, files := filepath.Join(basepath, "bar"), make([]*CFile, 50)
	for i := range files {
		name := fmt.Sprint("bar", i+1, ".c")
		files[i] = mjölnir.CFile(filepath.Join(path, name)).SetStd("c99").Includes(basepath).DependsOn(foo.Files...)
	}
	return &Bar{files}
}
