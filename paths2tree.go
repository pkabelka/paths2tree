package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/pkabelka/litu"
)

type Path2TreeRes struct {
	TreeLevel string
	Err error
}

func Paths2Tree(r io.Reader) <-chan Path2TreeRes {
	type PathHist struct {
		depth int
		path []string
	}

	scanner := bufio.NewScanner(r)

	depth := 0
	var lastDir string
	pathHist := make([]PathHist, 0, 256)
	treeLevel := ""
	ch := make(chan Path2TreeRes)

	go func() {
		defer close(ch)
		for scanner.Scan() {
			treeLevel = ""
			path := scanner.Text()
			splitPath := strings.Split(path, "/")
			depth = len(splitPath)-1

			if depth > 0 && splitPath[len(splitPath)-2] != lastDir {
				found := false
				for _, e := range pathHist {
					if e.depth == depth && litu.Equal(e.path, splitPath[:len(splitPath)-1]) {
						found = true
					}
				}
				if !found {
					treeLevel += fmt.Sprintf("%s%s\n", strings.Repeat("|   ", depth-1), splitPath[len(splitPath)-2])
					lastDir = splitPath[len(splitPath)-2]
					pathHist = append(pathHist, PathHist{depth, splitPath[:len(splitPath)-1]})
				}
			}

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
