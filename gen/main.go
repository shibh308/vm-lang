package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	arg_len := len(os.Args) - 1

	if arg_len < 1 {
		_, _ = fmt.Fprintln(os.Stderr, "missing argument")
		os.Exit(1)
	}
	readPath := os.Args[1]
	f, err := os.Open(readPath)
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	s := string(b)
	writePath := "./out.scbc"
	if arg_len >= 2 {
		writePath = os.Args[2]
	}

	tokens := lexString(strings.TrimSpace(s))
	root := parseTokenSlice(tokens)
	root.generateByteCode()
	root.printByteCode()
	root.writeByteCode(writePath)
}
