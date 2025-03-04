package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// Validate file argument
	outFlag := flag.String("out", "", "relative or absolute path to output file")
	flag.Parse()

	if *(outFlag) == "" {
		fmt.Fprintf(os.Stderr, "Usage: tee <outfile>")
		return
	}

	// Create writer for file
	f, err := os.Create(*outFlag)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	fileWriter := bufio.NewWriter(f)

	// Create reader for stdin
	reader := bufio.NewReader(os.Stdin)

	// Create a multi-writer
	multiWriter := io.MultiWriter(os.Stdout, fileWriter)

	// Copy from stdin reader to stdout and file
	_, err = io.Copy(multiWriter, reader)
	if err != nil {
		log.Fatal(err)
	}

	// Flush file writer
	err = fileWriter.Flush()
	if err != nil {
		log.Fatal(err)
	}
}
