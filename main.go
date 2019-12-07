package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var appVer string
var buildVer string

func main() {
	help := flag.Bool("help", false, "Prints this screen")
	infile := flag.String("f", "", "in-file containing filepaths to be copied")
	outfolder := flag.String("o", "", "Base out folder where to copy files")
	//prefix := flag.String("p", "", "Optional prefix to prepend to filepaths read from in-file")
	flag.Parse()

	if *help {
		printUsage(false)
		return
	}

	if *infile == "" || *outfolder == "" {
		printUsage(true)
		return
	}

	files, err := readLines(*infile)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if f == "" {
			continue
		}
		// get folder name
		dir := filepath.Dir(f)
		// make sure path exist
		os.MkdirAll(filepath.Join(*outfolder, dir), os.ModePerm)

		//copy file
		outpath := filepath.Join(*outfolder, f)
		err = copy(f, outpath)
		if err != nil {
			fmt.Printf("error copying %s to %s [%s]\n", f, outpath, err.Error())
		} else {
			fmt.Printf("copied %s\n", f)
		}
	}
}

// Copy the src file to dst. Any existing file will be overwritten and will not
// copy file attributes.
func copy(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		txt = strings.ReplaceAll(txt, "/", "\\")
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func printUsage(iserr bool) {
	fmt.Println("filecp - utility to copy files from different locations to specific folder with hierarchy")
	fmt.Printf("Version %s build %s\n", appVer, buildVer)
	if iserr {
		fmt.Println(" ")
		fmt.Println("Error:  Insufficient input")
		fmt.Println(" ")
	}
	fmt.Println("Syntax:  filecp -f [infile] -o [destfolder]")
	fmt.Println(" ")
	flag.Usage()
	fmt.Println(" ")
	if iserr {
		os.Exit(1)
	}
	os.Exit(0)
}

//GOOS=windows GOARCH=386 go build -o filecp.exe
