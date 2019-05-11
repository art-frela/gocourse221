/*Package csvuse is demo implementation working with CSV.
General use:
generate CSV file with counting data of 5 shops with ";" separator.
after that - read generated file and separate data by shop and day
and  save them to specified CSV file with "," separator

Author: Karpov A. mailto:art.frela@gmail.com
Date: 2019-05-11
*/
package csvuse

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"
)

const (
	linesRead int = 200 // count strings in the buffer size for processing input CSV (may be CSV size of Gb)
)

var (
	wg sync.WaitGroup
)

// A CSVFile is structure for describe file with example data
type CSVFile struct {
	FileName string
}

// String - textual representation of the structure
func (src *CSVFile) String(countRow int) (content string) {
	// print filecontent
	f, err := os.Open(src.FileName)
	if err != nil {
		return fmt.Sprintf("%v", err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	strings.Trim(content, " ")
	for i := 0; i <= countRow; i++ {
		line, err := r.ReadString('\n')
		if err != nil {
			return fmt.Sprintf("%v", err)
		}
		content += line
	}
	return

}

// Separate - read CSV file, process content and make specified file for every shop and day.
// For reading and separating CSV file using SEMICOLON separator (";").
// For new CSV files using COMMA (",") separator.
func (src *CSVFile) Separate() (sepFiles []CSVFile, err error) {
	fullSepFiles := make([]CSVFile, 0, 5)
	f, err := os.Open(src.FileName)
	if err != nil {
		return
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.Comma = ';'
	bufRecords := make([][]string, 0, 10)
	i := 0
	isHeadRec := true
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		if isHeadRec {
			isHeadRec = false
			// you can insert there a validation input data by header fields
			continue
		}

		if i <= linesRead {
			bufRecords = append(bufRecords, record)
			//fmt.Println("###DEBUG bufRecords###", i, bufRecords)
			i++
		} else {
			bufRecords = append(bufRecords, record)
			i = 0
			tmpCSVFiles, err := addDataToSeparateFiles(bufRecords)
			if err != nil {
				return nil, err
			}
			for _, fn := range tmpCSVFiles {
				fullSepFiles = append(fullSepFiles, fn)
			}
			// clean buffer of data
			bufRecords = make([][]string, 0, 10)
		}
	}
	// fill uniq slice of CSVFile
	for _, v0 := range fullSepFiles {
		add := true
		for _, v1 := range sepFiles {
			if v1.FileName == v0.FileName {
				add = false
			}
		}
		if add {
			sepFiles = append(sepFiles, v0)
		}
	}

	return
}

// GenerateCountData - make file with example data
func GenerateCountData(filename string) (srcFile CSVFile, err error) {
	srcFile.FileName = filename
	// generate counting data file
	// ShopID;DateStart;Interval;Enters;Exits
	// 220;2019-05-05 12:00:00;900;3;4
	// 221;2019-05-05 12:00:00;900;22;15
	// shops - example network with 5th shops
	shops := [5]string{"Shop111", "Shop222", "Shop333", "Shop444", "Shop555"}

	// dates - example dates of counting data which will be place to the file
	dates := [3]string{"2019-05-06", "2019-05-07", "2019-05-08"}

	// newrecords - chan for recieve example record and insert them to the file
	// channels and goroutines using for random order data at the generate file
	newrecords := make(chan string)

	for _, shop := range shops {
		for _, date := range dates {
			wg.Add(1)
			go genExampleDataIntByInt(shop, date, newrecords)

		}
	}

	// start reciever string and past them to the file
	go func() {
		f, err := os.Create(filename)
		if err != nil {
			return
		}
		defer f.Close()
		//newrecords <-
		_, err = io.WriteString(f, "ShopID;DateStart;Interval;Enters;Exits\n")
		if err != nil {
			return
		}
		for oneRec := range newrecords {
			_, err := io.WriteString(f, oneRec)
			if err != nil {
				return
			}
		}
	}()

	wg.Wait()
	close(newrecords)

	return
}

// genExampleDataIntByInt - send to string channel new record for CSV EXAMPLE FILE.
// Helper for GenerateCountData.
func genExampleDataIntByInt(shop, date string, out chan<- string) {
	interval := "900"
	startDate, err := time.Parse("2006-01-02 15:04:05", date+" 00:00:00")
	if err != nil {
		log.Fatalf("For %s and %s dateParse error, %v\n", shop, date, err)
	}
	for i := 1; i <= 96; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		in := r.Int31n(100)
		o := r.Int31n(100)
		dts := startDate.Format("2006-01-02 15:04:05")
		out <- fmt.Sprintf("%s;%s;%s;%d;%d\n", shop, dts, interval, in, o)
		startDate = startDate.Add(900 * time.Second) //shift 900sec = 15minutes
	}
	wg.Done()
}

// addDataToSeparateFiles - write [][]string to CSV files, separated by shop and day, returns slice of CSV Files.
// Helper for Separate() method of CSVFile.
func addDataToSeparateFiles(data [][]string) (sepFiles []CSVFile, err error) {
	bufData := make(map[string]map[string][][]string)
	bufCount := make(map[string]int)
	//fmt.Printf("-=#DEBUG bufData - %+v#=-\n\n\n", bufData)
	for _, rec := range data {
		day := rec[1][:10]
		shop := rec[0]
		//fmt.Println("DEBUG filename for shop", day, shop)
		if bufData[shop] == nil {
			tmpShop := make(map[string][][]string)
			tmpShop[day] = append(tmpShop[day], rec)
			//fmt.Printf("DEBUG bufData - %+v; tmpShop - %+v", bufData, tmpShop)
			bufData[shop] = tmpShop
			bufCount[shop] += len(rec)
		} else {
			bufData[shop][day] = append(bufData[shop][day], rec)
			bufCount[shop] += len(rec)
		}
	}

	for shop, day := range bufData {
		for shopDay, shopData := range day {
			fn := fmt.Sprintf("%s_%s.csv", shop, shopDay)
			// debug
			//fmt.Println("DEBUG filename for shop = ", fn)
			//
			f, err := os.OpenFile(fn, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
			if err != nil {
				return nil, err
			}
			defer f.Close()
			w := csv.NewWriter(f)
			w.WriteAll(shopData)

			if err := w.Error(); err != nil {
				return nil, err
			}
			tmpSepFile := CSVFile{fn}
			sepFiles = append(sepFiles, tmpSepFile)

		}
	}

	return
}
