package main

import (
	"bytes"
	"io"
	"os"
)

func main() {
	header := bytes.NewBufferString("----- HEADER -----\n")
	context := bytes.NewBufferString("Example of  io.MultiReader\n")
	footer := bytes.NewBufferString("----- FOOTER -----\n")

	reader := io.MultiReader(header, context, footer)
	io.Copy(os.Stdout, reader)
}
