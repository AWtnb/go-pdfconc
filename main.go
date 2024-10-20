package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pdfcpu/pdfcpu/pkg/api"
)

const EXT = ".pdf"

func fromPipe() (lines []string, err error) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	return
}

func checkPiped() bool {
	stat, _ := os.Stdin.Stat()
	return (stat.Mode() & os.ModeCharDevice) == 0
}

func checkExt(paths []string) bool {
	for _, p := range paths {
		if filepath.Ext(p) != EXT {
			return false
		}
	}
	return true
}

func run(outname string) int {
	if !checkPiped() {
		fmt.Println("no stdin pipe")
		return 1
	}
	paths, err := fromPipe()
	if err != nil {
		fmt.Println(err)
		return 1
	}
	if !checkExt(paths) {
		fmt.Println("non-pdf file passed")
		return 1
	}
	if !strings.HasPrefix(outname, EXT) {
		outname = outname + EXT
	}
	p := filepath.Join(filepath.Dir(paths[0]), outname)
	if fi, err := os.Stat(p); err == nil && !fi.IsDir() {
		fmt.Println("outfile already exists")
		return 1
	}
	if err := api.MergeCreateFile(paths, p, true, nil); err != nil {
		fmt.Println(err)
		return 1
	}
	return 0
}

func main() {
	var (
		outname string
	)
	flag.StringVar(&outname, "outname", "conc", "output file name")
	flag.Parse()
	os.Exit(run(outname))
}
