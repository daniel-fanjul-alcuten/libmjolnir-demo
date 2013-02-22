package main

import (
	"flag"
	. "github.com/daniel-fanjul-alcuten/libmjolnir"
	. "libfoobar/bar"
	. "libfoobar/config"
	. "libfoobar/foo"
	"log"
	"os"
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

	foo := NewFoo(mjölnir, basepath)
	bar := NewBar(mjölnir, basepath, foo)

	mjölnir.CLibrary("libmfoobar.a").Link(bar.Files...)
	if err := mjölnir.Build(); err != nil {
		log.Fatalln("error:", err)
	}
}
