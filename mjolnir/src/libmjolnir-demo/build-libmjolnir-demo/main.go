package main

import (
	"flag"
	. "github.com/daniel-fanjul-alcuten/libmjolnir"
	. "libfoobar/bar"
	. "libfoobar/config"
	. "libfoobar/foo"
	. "libmjolnir-demo/baz"
	"log"
	"os"
	"path/filepath"
)

func main() {
	log.SetFlags(0)
	flag.Parse()

	basepath, err := os.Getwd()
	if err != nil {
		log.Fatalln("error:", err)
	}

	mjölnir := NewMjölnir()
	Configure(mjölnir, basepath)

	foobarpath := filepath.Join(basepath, "libfoobar")

	foo := NewFoo(mjölnir, foobarpath)
	bar := NewBar(mjölnir, foobarpath, foo)
	baz := NewBaz(mjölnir, basepath, bar)
	main := mjölnir.CFile(filepath.Join(basepath, "main.c")).DependsOn(baz.Files...)

	mjölnir.CExecutable("mmain").Link(main)
	if err := mjölnir.Build(); err != nil {
		log.Fatalln("error:", err)
	}
}
