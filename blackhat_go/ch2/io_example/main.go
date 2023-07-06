package main

import (
	"fmt"
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

	// Create buffer to hold input/output.
	input := make([]byte, 4096)

	// Use reader to read input.
	s, err := reader.Read(input)
	if err != nil {
		log.Fatalln("Unable to read data")
	}
	fmt.Printf("Read %d bytes from stdin\n", s)

	// Use writer to write output.
	s, err = writer.Write(input)
	if err != nil {
		log.Fatalln("Unable to write data")
	}
	fmt.Printf("Wrote %d bytes to stdout\n", s)
}
