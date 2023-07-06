package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

// This defines an io.Reader to read from Stdin.
type CustomReader struct{}

// Read reads data from stdin.
func (*CustomReader) Read(b []byte) (int, error) {
	fmt.Print("in > ")
	return os.Stdin.Read(b)
}

// CustomWriter defines an io.Writer to write to Stdout.
type CustomWriter struct{}

// Write writes data to stdout.
func (*CustomWriter) Write(b []byte) (int, error) {
	fmt.Print("out > ")
	return os.Stdout.Write(b)
}

func main() {
	// Instantiate reader and writer.
	var (
		reader CustomReader
		writer CustomWriter
	)

	if _, err := io.Copy(&writer, &reader); err != nil {
		log.Fatalln("Unable to read/write data")
	}
}
