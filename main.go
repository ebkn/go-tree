package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var (
	dirCount  = 0
	fileCount = 0
	showIcon  *bool
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
		var icon string
		if *showIcon {
			if flinfo.IsDir() {
				icon = string(0x1f4c1) + " "
			} else {
				icon = string(0x1f4c4) + " "
			}
		}
		fmt.Println(fmt.Sprintf("%s%s── %s%s", pre, branch, icon, f.Name()))

		if flinfo.IsDir() {
			dirCount++
			var blanks string
			if isLast {
				blanks = pre + "     "
			} else {
				blanks = pre + "│   "
			}
			printTree(dir+"/"+f.Name(), blanks)
		} else {
			fileCount++
		}
	}
	return nil
}

func main() {
	showIcon = flag.Bool("icon", false, "show icon")
	flag.Parse()
	args := flag.Args()
	var dir string
	if len(args) > 0 {
		dir = args[0]
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
