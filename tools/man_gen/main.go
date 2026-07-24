package main

import (
	"path/filepath"
	"runtime"
	"path"
	"man_gen/man"
	"man_gen/html"
	"man_gen/resolver"
	"fmt"
)

func main() {

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("cannot get caller info")
	}

	man_path := path.Join(filepath.Dir(filename),"..","..","man")
	html_path := path.Join(filepath.Dir(filename),"..","..","webpage","content")

	version_file := path.Join(filepath.Dir(filename),"..","..","VERSION")
	
	man_file := path.Join(man_path,"mkdotenv.1")
	man.MakeManpage(man_file,version_file)
	
	html_file := path.Join(html_path,"manpage_main.njk")
	html.MakeHtml(html_file)

	parser_folder := path.Join(filepath.Dir(filename),"..","..","mkdotenv","secret")
	resolver.GenerateDocsForSecretResolvers(parser_folder)
}
