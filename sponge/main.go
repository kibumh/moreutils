package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func printError(msg string) {
	fmt.Fprintf(os.Stderr, "%s: %s\n", filepath.Base(os.Args[0]), msg)
}

func main() {
	if len(os.Args) != 2 {
		printError("missing file operand")
		os.Exit(1)
	}

	buf := bytes.Buffer{}
	if _, err := io.Copy(&buf, os.Stdin); err != nil {
		printError(err.Error())
		os.Exit(1)

	}
	f, err := os.OpenFile(os.Args[1], os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		printError(err.Error())
		os.Exit(1)
	}
	if _, err := io.Copy(f, &buf); err != nil {
		printError(err.Error())
		os.Exit(1)
	}
	if err = f.Close(); err != nil {
		os.Exit(1)
	}
}
