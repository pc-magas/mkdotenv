package main

import (
	"path/filepath"
	"runtime"
	"path"
	"man_gen/man"
	"man_gen/html"
)

func main() {

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("cannot get caller info")
	}

	version_file := path.Join(filepath.Dir(filename),"..","..","VERSION")
	
	man_file := path.Join(filepath.Dir(filename),"..","..","man","mkdotenv.1")
	man.MakeManpage(man_file,version_file)
	
	html_file := path.Join(filepath.Dir(filename),"..","..","webpage","template","content","manpage_main.njk")
	html.MakeHtml(html_file)
}
