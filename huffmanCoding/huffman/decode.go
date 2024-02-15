package huffman

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Decode(fileName string, referenceTableFile string) {
	referenceTable := make(map[rune]int)
	referenceFile, err := os.Open(referenceTableFile)
	if err != nil {
		fmt.Println("Error reading reference table file")
		return
	}
	defer referenceFile.Close()
	scanner := bufio.NewScanner(referenceFile)
	for scanner.Scan() {
		// on each line, first is the character, second is the bitRepresentation
		text := scanner.Text()
		count, err := strconv.Atoi(text[1:])
		if err != nil {
			fmt.Println("Invalid reference table file")
			return
		}

		referenceTable[rune(text[0])] = count
	}

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error reading file")
		return
	}
	defer file.Close()

	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		for _, ch := range text {
			fmt.Print(ch)
		}
	}
}
