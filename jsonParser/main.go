package main

import (
	"bufio"
	"fmt"
	"jsonParser/parser"
	"os"
	"strings"
)

func handleParsingFile(fileName string) bool {
	if fileName == "" {
		panic("please provide a file name")
	}

	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	fileContents := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fileContents += scanner.Text()
	}

	return parser.ParseJson(strings.Trim(fileContents, " "))
}

type Test struct {
	fileName string
	valid    bool
}

func main() {
	fileNames := []Test{
		{fileName: "tests/step1/valid.json", valid: true},
		{fileName: "tests/step1/invalid.json", valid: false},

		{fileName: "tests/step2/valid.json", valid: true},
		{fileName: "tests/step2/invalid.json", valid: false},
		{fileName: "tests/step2/valid2.json", valid: true},
		{fileName: "tests/step2/invalid2.json", valid: false},

		{fileName: "tests/step3/valid.json", valid: true},
		{fileName: "tests/step3/invalid.json", valid: false},

		{fileName: "tests/step4/valid.json", valid: true},
		{fileName: "tests/step4/invalid.json", valid: false},
		{fileName: "tests/step4/valid2.json", valid: true},

		// my custom test cases
		{fileName: "tests/step5/valid.json", valid: true},
	}

	for _, test := range fileNames {
		if handleParsingFile(test.fileName) != test.valid {
			panic("FAILED " + test.fileName)
		} else {
			fmt.Println("PASSED " + test.fileName)
		}
	}
}
