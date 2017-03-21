package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func readLine(path string) []string {
	lines := []string{}
	inFile, err := os.Open(path)
	if err != nil {
		fmt.Printf("Issue with file open")
	}
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func main() {
	lines := readLine(os.Args[1])
	kernelOptionRegexp := regexp.MustCompile(`\s*append`)
	for _, line := range lines {
		if kernelOptionRegexp.MatchString(line) {
			options := strings.Join(strings.Fields(line)[1:], " ")
			fmt.Printf("%#v", options)
		}
	}
}
