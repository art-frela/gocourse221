/*
gocp - simple console util, which process copying files
	   analog of linux cp program
	   one thread executing

implement 3 flags:
 -i ask user for to confirm if existing file should be over written in the copying process
 -r for recursive, to copy all the subdirectories and files in a given directory and preserve the tree structure
 -v for verbose, shows files being copied one by one


 Author: Karpov A.
 Date: 2019-05-12
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

const buffersize int64 = 1073741824 // 1073741824/(1024*1024*1024)=1 eq 1Gb for processing big files

func main() {

	// flags set and Parse
	confirm := flag.Bool("i", false, "confirm if existing file should be over written in the copying process")
	recursive := flag.Bool("r", false, "for recursive, to copy all the subdirectories and files in a given directory and preserve the tree structure")
	verbose := flag.Bool("v", false, "for verbose, shows files being copied one by one")
	flag.Parse()

	osArgs := os.Args        // all arguments of cmd
	arguments := flag.Args() // all arguments excluding flags

	// check command format
	if len(arguments) < 2 {
		fmt.Fprintf(os.Stderr, "using %s [-i] [-r] [-v] source_file [file2 | file3 | ... fileN] target_file\n", filepath.Base(osArgs[0]))
		os.Exit(1)
	}

	// recursive execute copying
	if *recursive {
		if isDir(arguments[0]) && isDir(arguments[1]) {
			basedir := arguments[0]
			dst := arguments[1]
			//
			copyRecurse(basedir, dst, *confirm, *verbose)
			return
		}
	}

	// not recursive
	// fill source_files
	files := arguments[:len(arguments)-1]
	dst := arguments[len(arguments)-1:][0]

	// process every source file
	for _, src := range files {
		if isDir(dst) {
			dst = filepath.Join(dst, filepath.Base(src))
		}
		err := copyfile(src, dst, buffersize, *confirm, *verbose)
		if err != nil {
			//fmt.Fprintf(os.Stderr, "du: %v\n", err)
			fmt.Fprintf(os.Stderr, "GOCP error: %v\n", err)
			os.Exit(1)
		}
	}

}

// copyRecurse - recursive search based on and basedir / subdirs, creates the same structure in destdir and copies files there
func copyRecurse(basedir, destdir string, confirm, verbose bool) {
	//
	err := filepath.Walk(basedir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
		}
		if path == basedir {
			// skip root(basedir) dir
			return nil
		}
		// relevant path to source file
		suffixpath, _ := filepath.Rel(basedir, path)
		// abs path to source file
		source := path
		// make abs path to destination file/directory
		destination := filepath.Join(destdir, suffixpath)

		if info.IsDir() {
			// make subfolder for destination
			return os.MkdirAll(destination, info.Mode().Perm())
		}
		return copyfile(source, destination, buffersize, confirm, verbose)
	})
	if err != nil {
		// print error to stderr and continue
		fmt.Fprintf(os.Stderr, "error walking the path %q: %v\n", basedir, err)
	}

}

// copyfile - simple operation for copying file
func copyfile(src, dst string, buffersize int64, confirm, verbose bool) (err error) {
	rewrite := true
	srcFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !srcFileStat.Mode().IsRegular() {
		return fmt.Errorf("The %s is not a regular file.", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	_, err = os.Stat(dst)
	if err == nil {
		if confirm {
			rewrite = ynCheck("File <" + dst + "> exist, rewrite them [default is Y]?")
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
	// verbose print result when all is OK
	if verbose {
		fmt.Printf("%s  ->   %s\n", src, dst)
	}

	return
}

// isDir simple check Dir or not, return true if path is directory
func isDir(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fileInfo.IsDir()
}

// ynCheck - simple processing user input Y/N, returns true/false respectively.
// Default is TRUE!
func ynCheck(phrase string) bool {
	input := ""
	fmt.Printf("%s>", phrase)
	_, err := fmt.Scanln(&input)
	if err != nil {
		return true //default is true
	}
	yn := strings.ToLower(input)
	if yn == "n" {
		return false
	}
	return true
}
