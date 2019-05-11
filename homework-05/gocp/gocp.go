/*
gocp - simple console util, which process copying files
	   analog of linux cp program
implement 3 flags:
 -i ask user for to confirm if existing file should be over written in the copuing process
 -r for recursive, to copy all the subdirectories and files in a given directory and preserve the tree structure
 -v for verbose, shows files being copied one by one


 Author: Karpov A.
 Date: 2019-05-11
*/
package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const buffersize int64 = 1073741824 // 1073741824/(1024*1024*1024)=1 eq 1Gb

func main() {
	// flags set and Parse
	confirm := flag.Bool("i", false, "confirm if existing file should be over written in the copying process")
	recursive := flag.Bool("r", false, "for recursive, to copy all the subdirectories and files in a given directory and preserve the tree structure")
	verbose := flag.Bool("v", false, "for verbose, shows files being copied one by one")
	flag.Parse()

	arguments := flag.Args()
	fmt.Println(*confirm, *verbose, *recursive, arguments)
	// check command format
	if len(arguments) < 2 {
		fmt.Printf("using %s [-i] [-r] [-v] source_file [file2 | file3 | ... fileN] target_file\n", filepath.Base(arguments[0]))
		os.Exit(1)
	}
	// process source_file
	files := arguments[:len(arguments)-1]

	for _, src := range files {
		dst := arguments[len(arguments)-1:][0]
		if isDir(arguments[len(arguments)-1:][0]) {
			dst = filepath.Join(dst, filepath.Base(src))
		}
		fmt.Printf("%s -> %s\n", src, dst)
		err := copy(src, dst, buffersize, *confirm, *recursive, *verbose)
		if err != nil {
			fmt.Printf("Error %v\n", err)
			os.Exit(1)
		}
		if *verbose {
			fmt.Printf("file %s copied to %s successfully\n", src, dst)
		}
	}

}

// isDir simple check Dir or not, return true if path is directory
func isDir(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fileInfo.IsDir()
}

// ynCheck -
func ynCheck(phrase string) bool {
	input := ""
	fmt.Printf("%s>", phrase)
	_, err := fmt.Scanln(&input)
	if err != nil {
		return false
	}
	yn := strings.ToLower(input)
	if yn == "y" {
		return true
	}
	return false
}

// copy -
func copy(src, dst string, buffersize int64, confirm, recursive, verbose bool) (err error) {
	rewrite := true
	srcFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !srcFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file.", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	_, err = os.Stat(dst)
	if err == nil {
		if confirm {
			rewrite = ynCheck("File " + dst + " exist, rewrite them? type Y/N and press Enter?")
		}
		if !rewrite {
			return nil //fmt.Errorf("File %s exists and will not be rewritten", dst)
		}
	}

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()

	// buf - size of buffer for copying files
	buf := make([]byte, buffersize)

	// copying cicle
	for {
		// read piece of data
		n, err := source.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}
		// write peice to dest file
		if _, err := destination.Write(buf[:n]); err != nil {
			return err
		}
	}
	return
}
