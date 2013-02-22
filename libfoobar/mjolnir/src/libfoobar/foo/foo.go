package foo

import (
	"fmt"
	. "github.com/daniel-fanjul-alcuten/libmjolnir"
	"path/filepath"
)

type Foo struct {
	Files []*CFile
}

func NewFoo(mjölnir *Mjölnir, basepath string) *Foo {
	path, files := filepath.Join(basepath, "foo"), make([]*CFile, 50)
	for i := range files {
		name := fmt.Sprint("foo", i+1, ".c")
		files[i] = mjölnir.CFile(filepath.Join(path, name)).SetStd("c99").Includes(basepath)
	}
	return &Foo{files}
}
