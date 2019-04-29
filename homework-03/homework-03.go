/*
homework-03 is console util which must doing three things:
- print a structure
- emulate a queue FIFO
- save a phonebook to the disk
Author: Artem K mailto:art.frela@gmail.com
Date: 2019-04-29
*/

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

//TrackCar - structure for Track car
type TrackCar struct {
	VIN             string  `json:"vin"`
	Mark            string  `json:"mark"`
	Model           string  `json:"model"`
	YearManufacture int     `json:"yearmanufacture"`
	EngineIsRun     bool    `json:"engineisrun"`
	WindowIsOpen    bool    `json:"windowisopen"`
	BodySizeTon     int     `json:"bodysizeton"`
	BodyOccupancy   float64 `json:"bodyoccupancy"`
}

//PassageCar - structure for passage car
type PassageCar struct {
	VIN               string  `json:"vin"`
	Mark              string  `json:"mark"`
	Model             string  `json:"model"`
	YearManufacture   int     `json:"yearmanufacture"`
	EngineIsRun       bool    `json:"engineisrun"`
	WindowIsOpen      bool    `json:"windowisopen"`
	TrunkVolumeLitter int     `json:"trunkvolumelitter"`
	TrunkOccupancy    float64 `json:"trunkoccupancy"`
}

//Message - message structure
type Message struct {
	ID   int    `json:"id"`
	Body string `json:"body"`
}

func main() {

	// [TASK-1]
	// Описать несколько структур — любой легковой автомобиль и грузовик.
	// Структуры должны содержать марку авто, год выпуска, объем багажника/кузова,
	// запущен ли двигатель, открыты ли окна, насколько заполнен объем багажника.
	car1 := PassageCar{
		VIN:               "536DFS@$3465657",
		Mark:              "Skoda",
		Model:             "Ocavia",
		YearManufacture:   2009,
		EngineIsRun:       false,
		WindowIsOpen:      false,
		TrunkVolumeLitter: 588,
		TrunkOccupancy:    5.5,
	}

	truck1 := TrackCar{
		VIN:             "536DFS@$34656456",
		Mark:            "Mercedes",
		Model:           "Actros",
		YearManufacture: 2018,
		EngineIsRun:     false,
		WindowIsOpen:    false,
		BodySizeTon:     30,
		BodyOccupancy:   0,
	}

	printSome("After Initialazed: ", car1, truck1)

	// [TASK-2]
	// Инициализировать несколько экземпляров структур. Применить к ним различные действия.
	// Вывести значения свойств экземпляров в консоль.

	// to Run cars
	car1.EngineIsRun = true
	truck1.EngineIsRun = true
	printSome("Engine run: ", car1, truck1)

	// to open windows
	car1.WindowIsOpen = true
	truck1.WindowIsOpen = true
	printSome("Windows open: ", car1, truck1)
	// to fill 45% truck1
	truck1.BodyOccupancy = 45.4
	printSome("to fill 45% truck1: ", car1, truck1)

	// [TASK-3]
	// * Реализовать очередь. Это структура данных, работающая по принципу FIFO
	// (First Input — first output, или «первым зашел — первым вышел»).
	printSome("\n\n\nTASK-3:")
	var queueOfMessages = make([]Message, 0, 20)

	// simple push message to queue
	examplePushOrderedMessage(5, &queueOfMessages)

	printSome("Basic queue: ", queueOfMessages)

	reverseQueue := reverseQueueMessage(queueOfMessages)
	for ix, msg := range reverseQueue {
		fmt.Printf("Index=%d, Message=%+v\n", ix, msg)
	}

	printSome("\n\n\nTASK-4:")
	// [TASK-4]
	// * Внести в телефонный справочник дополнительные данные.
	// Реализовать сохранение json-файла на диске с помощью пакета ioutil при изменении данных.

	addressBook := make(map[string][]int)

	addressBook["Alex"] = []int{89996543210}
	addressBook["zoro"] = []int{89030011234}
	addressBook["Bob"] = []int{89167243812}
	saveAddressbook("addressbook.json", true, addressBook)
	addressBookBefore := addressBook
	addressBook["Bob"] = append(addressBook["Bob"], 89155243627)
	if mapEq(addressBookBefore, addressBook) {
		saveAddressbook("addressbook.json", true, addressBook)
	}

	fmt.Println("Done.")
}

// printSome - simple analog of fmt.Println but in a column
func printSome(args ...interface{}) {
	for _, val := range args {
		fmt.Printf("%+v\n", val)
	}
}

// examplePushOrderedMessage - simple filler for slice of Messages
func examplePushOrderedMessage(count int, queue *[]Message) {
	for i := 1; i <= count; i++ {
		tmpMsg := Message{
			i,
			fmt.Sprintf("Message # %d", i),
		}
		*queue = append(*queue, tmpMsg)
	}
}

// reverseQueueMessage - make reverse slice, input = Message slice, output = reverse Message slice
func reverseQueueMessage(inputQ []Message) (output []Message) {
	maxIxInput := len(inputQ) - 1
	for ix := maxIxInput; ix >= 0; ix-- {
		output = append(output, inputQ[ix])
	}
	return
}

// equal - compare maps, if equal return true
func mapEq(left, right map[string][]int) bool {
	if len(left) != len(right) {
		return false
	}

	for lix, lval := range left {
		if rval, ok := right[lix]; !ok || !sliceEq(rval, lval) {
			return false
		}
	}
	return true
}

// sliceEq - compares two slices, if eq return true
func sliceEq(left, right []int) bool {
	if len(left) != len(right) {
		return false
	}

	for lix, lval := range left {
		if rval := right[lix]; lval != rval {
			return false
		}
	}
	return true
}

// saveAddressbook - saves the addressbook to dist, json format
func saveAddressbook(filename string, pretty bool, data map[string][]int) (err error) {
	var dataSave []byte
	if pretty {
		dataSave, err = json.MarshalIndent(data, "", "\t")
	} else {
		dataSave, err = json.Marshal(data)
	}
	if err != nil {
		return
	}
	//save to disk
	err = ioutil.WriteFile(filename, dataSave, 0644)
	return
}
