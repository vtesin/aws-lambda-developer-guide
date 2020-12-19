// Dup3 Loads the whole file into memory and then counts the occurences
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)

		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}

		for _, line := range strings.Split(string(data), "\n") {
			counts[strings.Join([]string{filename, line}, "\t")]++
		}
		for line, n := range counts {
			if n > 1 {
				fileAndLine := strings.Split(line, "\t")
				fmt.Printf("%d\t%s\tfile: %s\n", n, fileAndLine[1], fileAndLine[0])
			}
		}
	}
}
