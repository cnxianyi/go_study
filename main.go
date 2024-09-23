// main.go
package main

import (
	"fmt"
	"os"

	"github.com/cnxianyi/go_study/src/grammar"
)

func main() {
	grammar.PrintGrammar()
	grammar.AllType()

	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
