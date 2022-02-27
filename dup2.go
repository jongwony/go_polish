package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	files := os.Args[1:]
	dup_info := make(map[string]map[string]int)
	if len(files) == 0 {
		countLines(os.Stdin, dup_info, "")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, dup_info, arg)
		}
	}
	for filename, d := range dup_info {
		for line, n := range d {
			if n > 1 {
				fmt.Printf("[%s]\t%d\t%s\n", filename, n, line)
			}
		}
	}
}

func countLines(f *os.File, dup_info map[string]map[string]int, filename string) {
	input := bufio.NewScanner(f)
	counts := make(map[string]int)
	for input.Scan() {
		counts[input.Text()]++
	}
	dup_info[filename] = counts
}
