package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Path2TreeRes struct {
	TreeLevel string
	Err       error
}

func Paths2Tree(r io.Reader) <-chan Path2TreeRes {
	scanner := bufio.NewScanner(r)

	depth := 0
	var lastDir string
	pathHist := make(map[string]struct{}, 256)
	treeLevel := ""
	ch := make(chan Path2TreeRes)

	go func() {
		defer close(ch)
		for scanner.Scan() {
			treeLevel = ""
			path := scanner.Text()

			splitPath := strings.Split(path, "/")
			depth = len(splitPath) - 1

			// subdirectories
			lastSep := strings.LastIndexByte(path, '/')
			if lastSep >= 0 && depth > 0 && splitPath[len(splitPath)-2] != lastDir {
				pathToDir := path[:lastSep]
				// check if the path to item was visited before and print the
				// newly entered directory
				if _, ok := pathHist[pathToDir]; !ok {
					treeLevel += fmt.Sprintf("%s%s\n", strings.Repeat("|   ", depth-1), splitPath[len(splitPath)-2])
					lastDir = splitPath[len(splitPath)-2]
					pathHist[pathToDir] = struct{}{}
				}
			}

			// print the current item
			treePrefix := strings.Repeat("|   ", depth)
			treeLevel += fmt.Sprintf("%s%s\n", treePrefix, splitPath[len(splitPath)-1])

			if depth > 0 {
				lastDir = splitPath[len(splitPath)-2]
			}
			ch <- Path2TreeRes{treeLevel, nil}
		}
		if err := scanner.Err(); err != nil {
			ch <- Path2TreeRes{"", err}
		}
	}()

	return ch
}

func main() {
	ch := Paths2Tree(os.Stdin)
	for paths2TreeRes := range ch {
		if paths2TreeRes.Err != nil {
			log.Fatal(paths2TreeRes.Err)
		}
		fmt.Print(paths2TreeRes.TreeLevel)
	}
}
