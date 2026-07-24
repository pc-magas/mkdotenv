package resolver

import(
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"man_gen/resolver/scan"
	"man_gen/resolver/parser"
	"man_gen/resolver/format"
)

func GenerateDocsForSecretResolvers(secret_folder string){

	files := scan.ScanDirectory(secret_folder)
	
	for _, file := range files {

		f, err := os.Open(file)
		
		if err != nil {
			panic(err)
		}
		defer file.Close()
		fmt.Println("Generating Doc for "+file)
    	spec := parser.ParseComment(f)
		GenerateHtmlPage(spec)
	}
}

func GenerateManPage(spec parser.CommandSpec) {

}

func GenerateHtmlPage(spec parser.CommandSpec) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("cannot get caller info")
	}
	path := path.Join(path.Dir(filename),"..","..","..","webpage","content",spec.Command+".html")
	fmt.Println("Generating Html file "+path)

	f, err := os.Open(file)
		
	if err != nil {
		panic(err)
	}
	defer file.Close()

	f.WriteString(format.GenerateName(spec,true))
	f.WriteString(format.GenerateUsage(spec,true))
	f.WriteString(format.GenerateDescription(spec,true))
	f.WriteString(format.GenerateParamsDescription(spec.true))
	f.WriteString(format.GenerateFieldDescription(spec,true))
}