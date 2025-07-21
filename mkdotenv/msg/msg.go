package msg

import(
	"os"
	"fmt"
	"slices"
	"strings"
	"github.com/pc-magas/mkdotenv/params"
)

// This is changed upon runtime.
var version = "dev"

func ExitError(msg string){
	fmt.Fprintln(os.Stderr,msg)
	PrintHelp()
	os.Exit(1)
}

func buildArgumentUsage() string {
	var paramUsage strings.Builder

	for _, meta := range params.GetFlagsMeta() {
		line := fmt.Sprintf("  --%s, -%s", meta.Name,meta.Name)
		for _,alias := range meta.Aliases {
			line+= fmt.Sprintf(", --%s, -%s",alias,alias)
		}
		
		line+="\t"
		if meta.Required {
			line+="REQUIRED"
		} else {
			line+="OPTIONAL"
		}

		line+=" "+meta.Usage
		line+="\n"
		paramUsage.WriteString(line)
	}

	return paramUsage.String()
}

func buildCommandUsage() string {
	groups := make(map[int][]params.FlagMeta)
	orders := []int{}

	for _, meta := range params.GetFlagsMeta() {
		order := meta.Order
		groups[order] = append(groups[order], meta)

		if !slices.Contains(orders, order) {
			orders = append(orders, order)
		}
	}

	slices.Sort(orders)

	var builder strings.Builder

	for _, order := range orders {
		flags := groups[order]
		builder.WriteString("\n\t  ")
		for _, meta := range flags {
			// Start with canonical flag
			part := fmt.Sprintf("--%s|-%s", meta.Name,meta.Name)
				
			for _,alias := range meta.Aliases {
					part+= fmt.Sprintf("|--%s|-%s",alias,alias)
			}

			if meta.Type == params.StringType {
				part+= fmt.Sprintf(" <%s>", strings.ReplaceAll(meta.Name, "-", "_"))
			}

			// Wrap optional flags in []
			if !meta.Required {
				part = "[ " + part + " ]"
			}

			// Append with a space
			builder.WriteString(" ")
			builder.WriteString(part)
		}

		builder.WriteString(" \\")
	}

	return builder.String()
}

func PrintHelp() {
	PrintVersion()
	executable:=os.Args[0]
	fmt.Fprintln(os.Stderr,"\nUsage:\n\t"+executable+" \\"+buildCommandUsage()+"\n")
	fmt.Fprint(os.Stderr,"Options:\n\n")
	fmt.Fprintln(os.Stderr,buildArgumentUsage())
}

func PrintVersion(){
	fmt.Fprintln(os.Stderr,"\nMkDotenv VERSION: ",version)
	fmt.Fprintln(os.Stderr,"Replace or add a variable into a .env file.")
}