package main
import (
	"fmt"
	"os"
	"jarren.dirk/monkeylang/repl"
)


func main() {
	fmt.Printf("Monkey REPL\n")
	repl.Start(os.Stdin, os.Stdout)
}
