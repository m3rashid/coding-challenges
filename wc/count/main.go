package count

import (
	"bufio"
	"bytes"
	"fmt"
)

const (
	PIPE = "PIPE"
	FILE = "FILE"
)

type Args struct {
	IsCountWord bool
	IsCountLine bool
	IsCountByte bool
	IsCountChar bool
}

func resetBufferScanner(contents *bytes.Buffer) *bufio.Scanner {
	var newContents bytes.Buffer
	newContents.Write(contents.Bytes())
	return bufio.NewScanner(&newContents)
}

func GetCharCount(scanner *bufio.Scanner) int {
	count := 0
	scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {
		count++
	}
	return count
}

func GetWordCount(scanner *bufio.Scanner) int {
	count := 0
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		count++
	}
	return count
}

func GetLineCount(scanner *bufio.Scanner) int {
	count := 0
	for scanner.Scan() {
		count++
	}
	return count
}

func GetByteCount(scanner *bufio.Scanner) int {
	count := 0
	scanner.Split(bufio.ScanBytes)
	for scanner.Scan() {
		count++
	}
	return count
}

func HandleArgs(args Args, contents bytes.Buffer, extraLen int, cmdType string) {
	if cmdType != FILE && cmdType != PIPE {
		panic("Invalid command type")
	}

	if args.IsCountLine {
		scanner := resetBufferScanner(&contents)
		res := GetLineCount(scanner)
		fmt.Print(res, "  ")
	}

	if args.IsCountWord {
		scanner := resetBufferScanner(&contents)
		res := GetWordCount(scanner)
		fmt.Print(res, "  ")
	}

	if args.IsCountByte {
		scanner := resetBufferScanner(&contents)
		res := GetByteCount(scanner) + extraLen
		fmt.Print(res, "  ")
	}

	if args.IsCountChar {
		scanner := resetBufferScanner(&contents)
		res := GetCharCount(scanner) + extraLen
		fmt.Print(res, "  ")
	}
}
