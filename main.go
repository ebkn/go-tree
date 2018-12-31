package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var (
	dirCount  = 0
	fileCount = 0
)

func printTree(dir string, pre string) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	for i, f := range files {
		isLast := i == len(files)-1
		flinfo, err := os.Stat(dir + "/" + f.Name())
		if err != nil {
			log.Fatal(err.Error())
		}

		var branch string
		if isLast {
			branch = "└"
		} else {
			branch = "├"
		}
		fmt.Println(fmt.Sprintf("%s%s── %s", pre, branch, f.Name()))

		if flinfo.IsDir() {
			dirCount += 1
			var blanks string
			if isLast {
				blanks = pre + "     "
			} else {
				blanks = pre + "│   "
			}
			printTree(dir+"/"+f.Name(), blanks)
		} else {
			fileCount += 1
		}
	}
	return nil
}

func main() {
	var dir string
	if len(os.Args) > 1 {
		dir = os.Args[1]
	} else {
		dir = "."
	}
	if _, err := os.Stat(dir); err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(dir)
	if err := printTree(dir, ""); err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(fmt.Sprintf("\n%d directories, %d files.", dirCount, fileCount))
}
