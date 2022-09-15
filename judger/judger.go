package judger

import (
	"bytes"
	"errors"
	"log"
	"os"
	"os/exec"
	"strings"
)

// exec the certain input for source code
func Judge(src string, test string) (bool, error) {
	filename := getFileName(src)
	compile := exec.Command("g++", "-o", filename, src)
	run := exec.Command(filename)
	clean := exec.Command("rm", filename)
	var out bytes.Buffer
	in, err := os.Open(test + "/input.txt")
	defer in.Close()
	if err != nil {
		return false, errors.New("Error opening the input file")
	}
	correct, err := os.ReadFile(test + "/correct.text")
	if err != nil {
		return false, errors.New("Error closing the correct file")
	}

	run.Stdin = in
	run.Stdout = &out
	err = compile.Run()
	defer func() {
		if err = clean.Run(); err != nil {
			log.Fatal("Error cleanig the binary")
		}
	}()
	if err != nil {
		return false, errors.New("Error compiling the Code")
	}
	err = run.Run()
	if err != nil {
		return false, errors.New("Error runnning the binary")
	}
	return bytes.Compare(correct, out.Bytes()) != 0, nil
}

// get filename without extension
func getFileName(filename string) string {
	var sb strings.Builder
	runes := []rune(filename)
	var back int
	for i := len(runes) - 1; i >= 0; i-- {
		if runes[i] == rune('.') {
			back = i
			break
		}
	}
	for i := 0; i < back; i++ {
		sb.WriteRune(runes[i])
	}
	return sb.String()
}
