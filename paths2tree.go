package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type PathHist struct {
	depth int
	dir string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	depth := 0
	var lastDir string
	pathHist := make([]PathHist, 0, 64)

	for scanner.Scan() {
		path := scanner.Text()
		splitPath := strings.Split(path, "/")
		depth = len(splitPath)-1

		if depth > 0 && splitPath[len(splitPath)-2] != lastDir {
			found := false
			for _, e := range pathHist {
				if e.depth == depth && e.dir == splitPath[len(splitPath)-2] {
					found = true
				}
			}
			if !found {
				fmt.Printf("%s%s\n", strings.Repeat("|   ", depth-1), splitPath[len(splitPath)-2])
				lastDir = splitPath[len(splitPath)-2]
				pathHist = append(pathHist, PathHist{depth, lastDir})
			}
		}

		treePrefix := strings.Repeat("|   ", depth)
		fmt.Printf("%s%s\n", treePrefix, splitPath[len(splitPath)-1])
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
