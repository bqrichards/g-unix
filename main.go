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
	// Command-line arg parsing - https://gobyexample.com/command-line-flags
	filePtr := flag.String("file", "", "relative or absolute path to file")
	flag.Parse()

	if *(filePtr) == "" {
		fmt.Println("Welcome to GHexDump!\nThis is a hexdump command written in Golang.\nUsage: ghexdump <file>")
		return
	}

	// Open file
	f, err := os.Open(*filePtr)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	chunkSize := 16
	chunk := make([]byte, chunkSize)
	currentChunk := 0

	for {
		n, err := reader.Read(chunk) // load chunk into buffer
		if err != nil {
			if err != io.EOF {
				log.Fatal(f)
			}
			break
		}

		chunk = chunk[:n]

		// Print byte offset
		fmt.Printf("%08x  ", currentChunk)

		// Print bytes
		for index := range chunkSize {
			if index >= len(chunk) {
				fmt.Printf("   ")
			} else {
				value := chunk[index]
				fmt.Printf("%02x ", value)
			}

			// Print extra space every 8 bytes
			if (index+1)%8 == 0 {
				fmt.Printf(" ")
			}
		}

		// Print readable output
		fmt.Printf("|")
		for _, value := range chunk[:n] {
			if value < 32 || value > 126 {
				value = '.'
			}

			fmt.Printf("%s", string(value))
		}

		fmt.Printf("|\n")
		currentChunk += n
	}

	// Print empty last line
	fmt.Printf("%08x\n", currentChunk)
}
