package main

import (
	"path/filepath"
	"runtime"
	"path"
	"man_gen/man"
)

func main() {

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("cannot get caller info")
	}

	man_file := path.Join(filepath.Dir(filename),"..","..","man","mkdotenv.1")
	version_file := path.Join(filepath.Dir(filename),"..","..","VERSION")

	man.MakeManpage(man_file,version_file)
}
