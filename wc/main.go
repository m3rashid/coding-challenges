package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"
	"wc/count"
)

func main() {
	isCountWord := flag.Bool("w", false, "count words in the given file")
	isCountLine := flag.Bool("l", false, "count lines in the given file")
	isCountByte := flag.Bool("c", false, "count bytes in the given file")
	isCountChar := flag.Bool("m", false, "count characters in the given file")
	flag.Parse()
	otherArgs := flag.Args()

	if !*isCountWord && !*isCountLine && !*isCountByte && !*isCountChar {
		*isCountWord = true
		*isCountLine = true
		*isCountByte = true
	}

	cmdArgs := count.Args{
		IsCountWord: *isCountWord,
		IsCountLine: *isCountLine,
		IsCountByte: *isCountByte,
		IsCountChar: *isCountChar,
	}

	fileName := ""
	var contents bytes.Buffer

	if len(otherArgs) == 0 {
		// check if there is content from unix pipe
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Fprintln(&contents, line)
		}

		count.HandleArgs(cmdArgs, contents, count.PIPE)
	} else {
		fileName = otherArgs[0]
		file, err := os.Open(fileName)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Fprintln(&contents, line)
		}
		count.HandleArgs(cmdArgs, contents, count.FILE)
	}

	if fileName != "" {
		fmt.Println(fileName)
	} else {
		fmt.Println()
	}
}
