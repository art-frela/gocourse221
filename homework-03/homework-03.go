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
	"fmt"
)

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

func main() {

	//task - 1
	//Описать несколько структур — любой легковой автомобиль и грузовик.
	//Структуры должны содержать марку авто, год выпуска, объем багажника/кузова,
	//запущен ли двигатель, открыты ли окна, насколько заполнен объем багажника.
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

	//task 2
	//Инициализировать несколько экземпляров структур. Применить к ним различные действия.
	//Вывести значения свойств экземпляров в консоль.

	//to Run cars
	car1.EngineIsRun = true
	truck1.EngineIsRun = true
	printSome("Engine run: ", car1, truck1)

	//to open windows
	car1.WindowIsOpen = true
	truck1.WindowIsOpen = true
	printSome("Windows open: ", car1, truck1)
	//to fill 45% truck1
	truck1.BodyOccupancy = 45.4
	printSome("to fill 45% truck1: ", car1, truck1)

}

func printSome(structurs ...interface{}) {
	for _, val := range structurs {
		fmt.Printf("%+v\n", val)
	}
}
