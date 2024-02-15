package main

import (
	"flag"
	"fmt"
	"huffmanCoding/huffman"
	"os"
)

const EXTENSION = ".huff"

func main() {
	fileName := flag.String("f", "", "file to compress/decompress")
	compress := flag.Bool("c", false, "compress")
	decompress := flag.Bool("d", false, "decompress")
	referenceTableFile := flag.String("r", "", "reference table file")
	flag.Parse()

	if *fileName == "" {
		flag.PrintDefaults()
		return
	}

	if !*compress && !*decompress {
		fmt.Println("Please specify flags for what to do")
		flag.PrintDefaults()
		return
	}

	if *compress && *decompress {
		fmt.Println("Please specify only one flag for what to do")
		flag.PrintDefaults()
		return
	}

	if *decompress && *referenceTableFile == "" {
		fmt.Println("Please specify reference table file")
		flag.PrintDefaults()
		return
	}

	if _, err := os.Stat(*fileName); err != nil {
		fmt.Println("File does not exist")
		return
	}

	if *referenceTableFile != "" {
		if _, err := os.Stat(*referenceTableFile); err != nil {
			fmt.Println("Reference table file does not exist")
			return
		}
	}

	if *compress {
		huffman.Encode(*fileName)
	} else {
		huffman.Decode(*fileName, *referenceTableFile)
	}
}
