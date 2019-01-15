package main

import (
	"os"
	"strings"
	"fmt"
)

var cow = `
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||
`

func main() {
	input := strings.Join(os.Args[1:], " ")
	line := "--"
	for i := len(input); i > 0; i-- {
		line += "-"
	}
	text := fmt.Sprintf(" %s\n| %s |\n %s", line, input, line)
	fmt.Printf("%s%s", text, cow)
}