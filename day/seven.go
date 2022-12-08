// Copy and paste this each day to get started
package day

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Directory struct {
	path   string
	parent string
	size   int
}

func Seven(filepath string) {
	rootDir := Directory{"/", "", 0}

	fileSystem := map[string]Directory{"/": rootDir}

	f, err := os.Open(filepath)
	check(err)
	defer f.Close()
	sc := bufio.NewScanner(f)

	activeDir := rootDir
	cmd := ""

	for sc.Scan() {
		line := sc.Text()
		if line != "" {
			if line[0:1] == "$" {
				// if previous command was ls, add activeDir and dirSize to a slice for later reference
				cmd = strings.TrimSpace(strings.SplitAfterN(line, " ", 3)[1])
				if cmd == "cd" {
					cdPath := strings.SplitAfterN(line, "$ cd ", 2)[1]
					if cdPath == ".." {
						activeDir = fileSystem[activeDir.parent]
					} else if cdPath == "/" {
						activeDir = fileSystem["/"]
					} else {
						if activeDir.path == "/" {
							activeDir = fileSystem[activeDir.path+cdPath]
						} else {
							activeDir = fileSystem[activeDir.path+"/"+cdPath]
						}

					}
				}
			} else if line[0:3] == "dir" {
				dirName := activeDir.path + "/" + strings.TrimSpace(strings.SplitAfterN(line, " ", 2)[1])
				if activeDir.path == "/" {
					dirName = "/" + strings.TrimSpace(strings.SplitAfterN(line, " ", 2)[1])
				}
				if _, ok := fileSystem[dirName]; ok {
					continue
				}
				newDir := Directory{dirName, activeDir.path, 0}
				fileSystem[dirName] = newDir
			} else {
				fileSize, err := strconv.Atoi(strings.TrimSpace(strings.SplitAfterN(line, " ", 2)[0]))
				check(err)
				updateDir := Directory{activeDir.path, activeDir.parent, activeDir.size + fileSize}
				fileSystem[activeDir.path] = updateDir
				activeDir = fileSystem[activeDir.path]
			}
		}
	}
	sumTotal := 0
	grandTotal := 0
	for _, v := range fileSystem {
		fmt.Println("path:", v.path, "parent", v.parent, "size:", v.size)
		if v.path != "" {
			grandTotal += v.size
		}
		if v.size < 100000 {
			sumTotal += v.size
		}
	}

	delDir := []int{}
	fmt.Println("totalFileSize", grandTotal)
	for _, v := range fileSystem {
		if v.size >= (70000000 - grandTotal) {
			delDir = append(delDir, v.size)
		}
	}
	sort.Sort(sort.IntSlice(delDir))
	fmt.Println(delDir)
}
