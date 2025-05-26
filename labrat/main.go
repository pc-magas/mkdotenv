package main
import(
	"os"
	"flag"
	"fmt"
)


func main() {
	flagSet := flag.NewFlagSet("params", flag.ContinueOnError)
	flagSet.String("name", "default", "your name")
	flagSet.Int("age", 0, "your age")
	
	err := flagSet.Parse(os.Args[1:])
	if err != nil {
		fmt.Println("Error parsing flags:", err)
		return
	}

	flagSet.Visit(func(f *flag.Flag){
		fmt.Printf("Flag %s Value: %s\n",f.Name,f.Value.String())
	})
}
