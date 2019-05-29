/*
The gogrep - simple console utils which prints lines matching a pattern

Implement:
 - [x] simple search in file
 - [x] -c Count lines where appearance of word
 - [x] -r Recursive search
 - [x] -v Invert the sense of matching
 - [x] -w Select only those lines containing matches that form whole words
TODO: fix replace problem in the simpleCheck function

Author: Karpov A. mailto:art.frela@gmail.com
Date: 2019-05-12
*/
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	// settings program
	appName := os.Args[0]
	help := "using " + appName + ":\ncat <file name> | " + appName + " [OPTIONS] <string>\n<command> | " + appName + " [OPTIONS] <string>\n" + appName + " [OPTIONS] <string> [file name]"

	minusC := flag.Bool("c", false, "Count the appearance of words.")
	minusR := flag.Bool("r", false, "Recursive search")
	minusV := flag.Bool("v", false, "Invert the sense of matching")
	minusW := flag.Bool("w", false, "Select only those lines containing matches that form whole words")
	minusH := flag.Bool("h", false, "Print help using")
	flag.Parse()

	// print help if needed
	if *minusH {
		fmt.Println(help)
		os.Exit(0)
	}
	// validate input data - must be
	arguments := flag.Args()
	if len(arguments) < 1 { // too little
		fmt.Println(help)
		os.Exit(1)
	}
	// set pattern
	pattern := arguments[0]
	// for recursive processing
	if *minusR {
		basedir := arguments[1]
		results := grepRecurse(basedir, pattern, *minusC, *minusV, *minusW)
		for _, v := range results {
			fmt.Printf("%s\n", v)
		}
		os.Exit(0)
	}
	// not recursive
	var f *os.File
	filename := ""
	// define input method file/stdin
	if len(arguments) == 1 {
		f = os.Stdin
	} else {
		filename = arguments[1]
		fileHandler, err := os.Open(filename)
		if err != nil {
			fmt.Printf("error opening %s: %s", filename, err)
			os.Exit(1)
		}
		f = fileHandler
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	if *minusC {
		fmt.Printf("%d\n", countMatches(scanner, pattern, *minusV, *minusW))
		os.Exit(0)
	}
	outputs := simpleCheck(scanner, pattern, *minusV, *minusW)
	for _, v := range outputs {
		fmt.Println(v)
	}
}

// simpleCheck - returns slice of strings with matches
func simpleCheck(scanner *bufio.Scanner, pattern string, minusV, minusW bool) []string {
	reverse := make([]string, 0)
	matches := make([]string, 0)

	pattern = makePattern(pattern, minusW)

	for scanner.Scan() {
		res, ok := markMatch(scanner.Text(), pattern)
		if ok {
			matches = append(matches, res)
		} else {
			reverse = append(reverse, res)
		}
	}
	if minusV {
		return reverse
	}
	return matches
}

// countMatches - returns count of matches in the input
func countMatches(scanner *bufio.Scanner, pattern string, minusV, minusW bool) int {
	count := 0
	pattern = makePattern(pattern, minusW)
	for scanner.Scan() {
		matchMe := regexp.MustCompile(pattern)
		if minusV {
			if !matchMe.MatchString(scanner.Text()) {
				count++
			}
		} else {
			if matchMe.MatchString(scanner.Text()) {
				count++
			}
		}

	}
	return count
}

// makePattern - returns the specified string
func makePattern(pattern string, minusW bool) string {
	if minusW {
		return fmt.Sprintf(`\b(%s)\b`, pattern)
	}
	return fmt.Sprintf(`(?i)%s`, pattern)
}

// markMatch - returns string with colored matches and flag match exist or not
func markMatch(input, pattern string) (string, bool) {
	matchMe := regexp.MustCompile(pattern)
	repl := []byte(textWrap(matchMe.FindString(pattern), "red"))
	result := string(matchMe.ReplaceAll([]byte(input), repl))
	resOK := matchMe.MatchString(input)
	return result, resOK
}

// textWrap - wrapper text for coloring console output
func textWrap(text, color string) (res string) {
	switch color {
	case "green":
		res = "\033[1;5;92m" + text + "\033[0m"
	case "red":
		res = "\033[1;31m" + text + "\033[0m"
	default:
		res = "\033[1;38;5;220m" + text + "\033[0m"
	}
	return
}

// grepRecurse - recursive search based on and basedir / subdirs, check content for matching with template
// return slice of filenames and matched strings with colored matches
func grepRecurse(basedir, pattern string, minusC, minusV, minusW bool) (result []string) {
	//
	err := filepath.Walk(basedir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("prevent panic by handling failure accessing a path %q: %v", path, err)
		}

		if info.IsDir() || !info.Mode().IsRegular() {
			//skip processing of search in directory or simbol link files
			return nil
		}
		f, err := os.Open(path)
		if err != nil {
			return fmt.Errorf("error opening: %s", err)
		}
		defer f.Close()

		mode := info.Mode()
		if mode&0111 != 0 {
			// executing file, exclude from reading
			result = append(result, fmt.Sprintf("executing file %s, matches", textWrap(path, "green")))
			return nil
		}

		scanner := bufio.NewScanner(f)
		if minusC {
			c := countMatches(scanner, pattern, minusV, minusW)
			result = append(result, fmt.Sprintf("%s: %d", textWrap(path, "green"), c))
			return nil
		}
		for _, v := range simpleCheck(scanner, pattern, minusV, minusW) {
			result = append(result, fmt.Sprintf("%s: %s", textWrap(path, "green"), v))
		}
		return nil
	})

	if err != nil {
		// print error to stderr and continue
		fmt.Fprintf(os.Stderr, "error walking the path %q: %v", basedir, err)
	}
	return
}
