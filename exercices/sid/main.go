package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ProcessLine(line, old, new string) (found bool, res string, occ int) {
	found = strings.Contains(line, old)

	if found {
		occ = strings.Count(line, old)
		res = strings.Replace(line, old, new, -1)
	}

	return found, res, occ
}

func FindReplaceFile(src, old, new string) (occ int, lines []int, err error) {
	// Opening src file
	keywordsToMatch := [3]string{old, strings.ToLower(old), strings.ToUpper(old)}
	filePath, _err := os.Open(src)
	newFile, _err := os.Create(src + ".new")
	defer filePath.Close()
	defer newFile.Close()

	scanner := bufio.NewScanner(filePath)
	writer := bufio.NewWriter(newFile)
	defer writer.Flush()
	var lineNumber int = 0

	for scanner.Scan() {
		lineNumber += 1
		line := scanner.Text()
		newLine := line
		for _, keyword := range keywordsToMatch {
			found, res, _occ := ProcessLine(newLine, keyword, new)
			if found {
				newLine = res
				lines = append(lines, lineNumber)
				occ += _occ
			}
		}
		l := newLine + "\n"
		writer.Write([]byte(l))
	}
	return occ, lines, _err
}

func main() {
	var src = "README.md"
	var old = "Go"
	var new = "Rust"

	occ, lines, err := FindReplaceFile(src, old, new)

	if err != nil {
		fmt.Printf("%v", err)
	} else {
		fmt.Printf("== Summary ==\nNumber of occurences of Go: %v\nNumber of lines: %v\nLines: %v\n== End of Summary ==\n", occ, len(lines), lines)
	}
}
