package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func main() {
	folderName := flag.String("f", "", "Name of the folder to be created")
	binaryName := flag.String("b", "", "Name of the binary to be created")
	flag.Parse()

	if *folderName == "" {
		flag.PrintDefaults()
		return
	}

	// folder should not have any spaces or special characters
	alphabetic := regexp.MustCompile("^[a-zA-Z]+$")
	if !alphabetic.MatchString(*folderName) {
		fmt.Println("Folder name should only contain alphabetic string")
		fmt.Println("Task aborted")
		return
	}

	// binary should not have any spaces or special characters
	if *binaryName != "" {
		if !alphabetic.MatchString(*binaryName) {
			fmt.Println("Binary name should only contain alphabetic string")
			fmt.Println("Task aborted")
			return
		}
	}

	// folder should not exist beforehand
	dirs, err := os.ReadDir("../")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, str := range dirs {
		if !strings.Contains(str.Name(), "/") && !strings.Contains(str.Name(), ".") {
			if *folderName == str.Name() {
				fmt.Println("Folder already exists")
				fmt.Println("Task Aborted")
				return
			}
		}
	}

	if *binaryName == "" {
		*binaryName = *folderName
	}

	// create folder
	err = os.Mkdir("../"+*folderName, 0755)
	if err != nil {
		fmt.Println(err)
		return
	}

	// run bash commands
	commands := []string{
		"cd ../" + *folderName,
		"go mod init " + *folderName,
		"go mod tidy",
		"touch main.go",
		"go work use .",
		"echo 'package main\n' > main.go",
		"echo 'import \"fmt\"\n' >> main.go",
		"echo 'func main() {' >> main.go",
		"echo '  fmt.Println(\"Hello, World!\")' >> main.go",
		"echo '}' >> main.go",
		"touch README.md",
		"echo '### " + *folderName + "' > README.md",
		"touch Makefile",
		"echo 'build:' >> Makefile",
		"echo '\tgo build -o " + *binaryName + " main.go' >> Makefile",
	}

	allCommands := strings.Join(commands, " && ")
	exec.Command("bash", "-c", allCommands).Run()
}
