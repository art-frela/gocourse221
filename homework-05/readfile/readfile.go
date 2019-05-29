/*
fixed case of readfile from lesson #5
Requirements:
	* demo for study purpose
	* console util
	* read file from the disk (not STDIN)
	* print content of file to console

Solution Description:
1. Don't use os.Args (it will be later)
2. Check exists of that file
3. Checks size of the file and say "Ctrl+c for interrupt of printing content that big file"
4. Print text of error. Two big diffren silent exit from program or print "file don't exist or permission denied"
5. Not differentiate between plain text and binary file

Rationale for the decision'
general:
	- file ops is unsafe, error handler is needed
	- user must to known what happen
use two diffrent way for reading content of file:
1. for small files - all in memory
2. for big files - use buffering

Author: Karpov A. mailto:art.frela@gmail.com
Date: 2019-05-10
*/
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

const buffersize int64 = 1024 * 1024 // 1048576Byte -> 1Mb

func main() {
	// specify of filename (lte it wiil be get from os.Args, or stdin)
	filename := "some-file" //"some-file", "10MB", cmd for generate random file 10Mb in the linux: dd if=/dev/urandom of=10MB count=10000 bs=1024
	fileIsBig := false      // flag for print a small or a big file

	// check exist of file
	fileInfo, err := os.Stat(filename)
	if os.IsNotExist(err) {
		fmt.Printf("file %s does not exist\n", filename)
		os.Exit(1)
	}

	// check size of file
	fileSize := fileInfo.Size()
	if fileSize > buffersize { // >1Mb
		fileIsBig = true
	}

	// print content
	if fileIsBig {
		// print big files
		fmt.Println(filename)
		f, err := os.Open(filename)
		if err != nil {
			fmt.Printf("Open file error, %v", err)
			os.Exit(1)
		}
		defer f.Close()
		buf := make([]byte, buffersize)
		fmt.Printf("file [%s] %d[Byte] content is:\n", filename, fileSize)
		for {
			n, err := f.Read(buf)
			if err != nil && err != io.EOF {
				fmt.Printf("Read file error, %v", err)
				os.Exit(1)
			}
			if n == 0 {
				break
			}
			fmt.Printf("%s", buf[:n])
		}

	} else {
		// print small files
		// read all contnet of file to memory
		fileContent, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("file [%s] %d[Byte] content is:\n%s", filename, fileSize, fileContent)
	}
}
