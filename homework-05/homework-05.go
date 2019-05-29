/*
Tasks:
1. Изучите статью “Time in Go: A primer”. В письменном виде кратко изложите своё мнение: что из этой статьи стоило бы добавить в методичку?
   Если считаете что не стоит - так и напишите, приведя свои аргументы.
2. Что бы вы изменили в программе чтения из файла? Приведите исправленный вариант, обоснуйте свое решение в комментарии.
3. Самостоятельно изучите пакет encoding/csv, напишите пример, иллюстрирующий его применение.
4. * Напишите утилиту для копирования файлов, используя пакет flags.
5. ** Напишите упрощенный аналог утилиты grep.

  [TASK - 1 ]
  1. Для данного пакета и всех пакетов вообще полезно указывать для чего они в основном используются/могут использоваться
  2. пакет time для работы со временем:
	- получение текущего времени (для замеров скорости, вывода в журналы и пр.)
	- анализ текста и преобразование его в формат времени (парсинг) - важно, а тут ни слова. нужно упомянуть и пример.
	- таймшифтинг, так же ни слова про него, а это частый кейс - сдвинуть текущую дату на день или час(на интервла в общем)
	- сравнение времени: что [после], что [до], разницу в абсолютных значениях
т.е. описать базовые возможности и за подробностями отправлять в исп литературу.

	[TASK -2]
	See ./readfile/readFile.go
	Rationale for the decision
	general:
		- file ops is unsafe, error handler is needed
		- user must to known what happen
	use two diffrent way for reading content of file:
	1. for small files - all in memory
	2. for big files - use buffering

	[TASK - 3 - use csvuse package]
	for delete maked csv files use:
	cd /path/to/gocourse221/homework-05
	rm -rf *.csv
*/
package main

import (
	"fmt"
	"log"

	"./csvuse"
)

const (
	fileCSV string = "MyRetailCountingData.csv"
)

func main() {
	fmt.Println("Hello Gophers!")
	// [TASK 3] demo CSV processing
	// makes one file with counting data of 5 shops
	// reads that file and makes separete file for every shop and day
	// STEP BY STEP
	// 1. Generates "big" CSV file with counting data of 5 shops
	// 2. Print first 20 strings of "big" CSV file (you can see random order of data and ";" separator)
	// 3. Separate data from that file to different files for shops and days
	// 4. Print first 5 strings of those files (you can see filtered data with "," separator)

	// Generates "big" CSV file with counting data of 5 shops
	genCountFile, err := csvuse.GenerateCountData(fileCSV)
	if err != nil {
		log.Fatalf("Error of generate random CSV file, %v", err)
	}
	// Print first 20 strings of "big" CSV file (you can see random order of data and ";" separator)
	fmt.Println("content source file", genCountFile.FileName, "is:\n", genCountFile.String(20))

	// Separate data from that file to different files for shops and days
	sepFiles, err := genCountFile.Separate()
	if err != nil {
		log.Fatalf("Error of separate file, %v", err)
	}
	// Print first 5 strings of those files (you can see filtered data with "," separator)
	for _, fileShop := range sepFiles {
		fmt.Println(fileShop.FileName, "\n", fileShop.String(5))
	}

}
