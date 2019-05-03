/*
Author: Artem K., mailto:art.frela@gmail.com
Date: 2019-05-03
Task:
1. Написать свой интерфейс и создать несколько структур, удовлетворяющих ему.
2. Создать псевдоним типа телефонной книги и реализовать для него интерфейс Sort{}.
3. Дописать функцию, которая будет выводить справку к калькулятору. Она должна вызываться при вводе слова help вместо выражения.
4. * Написать функцию, которая будет получать позицию коня на шахматной доске, а возвращать массив из точек, в которые конь сможет сделать ход.
	Точку следует обозначить как структуру, содержащую x и y типа int
	Получение значений и их запись в точку должны происходить только с помощью отдельных методов. В них надо проводить проверку на то, что такая точка может существовать на шахматной доске.

*/
package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"gocourse221/calculator"
	"gocourse221/chees"
	"io/ioutil"
	"math/rand"
	"os"
	"sort"
	"time"
)

const (
	userxml  string = "phonebook.xml"
	userjson string = "phonebook.json"
)

func main() {
	// [TASK 1 demo]
	fmt.Printf("Task 1 - implement Storager interface, as JSON and XML storage\n")
	jStorage := storageJSONfile{
		filename: userjson,
	}
	xStorage := storageXMLfile{
		filename: userxml,
	}
	jStorage.readStorage()
	xStorage.readStorage()
	rand.Seed(time.Now().UnixNano())
	someAge := rand.Intn(100)
	someName := fmt.Sprintf("SomeName-%d", someAge)
	usr := User{
		NickName: someName,
		Age:      someAge,
		Phones:   []Phone{"+79017776655", "+7108033334455"},
	}
	juid, _ := jStorage.Insert(usr)
	xuid, _ := xStorage.Insert(usr)
	fmt.Printf("For user %v, juid=%d, xuid=%d\n", usr, juid, xuid)

	// [TASK 2 demo]
	fmt.Printf("Task 2 - implement Sort{} interface, for PhoneBook\n")

	var phonebook PhoneBook
	phonebook = PhoneBook(jStorage.data)
	fmt.Println("Storage before Sort")
	phonebook.String()
	sort.Sort(phonebook)
	fmt.Println("Storage after Sort")
	phonebook.String()

	// [TASK 3 demo]
	fmt.Printf("\n\nTask 3 - envelop Calculator, add HELP text\n")
	input := ""
	for {
		fmt.Print("type \"exit\" or \"help\"> ")
		if _, err := fmt.Scanln(&input); err != nil {
			fmt.Println(err)
			continue
		}

		if input == "exit" {
			break
		}
		if input == "help" {
			calculator.HelpPrint()
			continue
		}

		if res, err := calculator.Calculate(input); err == nil {
			fmt.Printf("Результат: %v\n", res)
		} else {
			fmt.Println("Не удалось произвести вычисление")
		}
	}

	// // [TASK 4 demo]

	fmt.Printf("\n\nTask 4 - chees possible moves for figures\n")
	// CheesFields - field for chees play
	var CheesField chees.Field
	CheesField.Init()

	for {
		CheesField.PrintField()
		fmt.Print("type \"exit\", type position X-Y (example 2-8 for kNight-black) for define possible moves > ")
		if _, err := fmt.Scanln(&input); err != nil {
			fmt.Println(err)
			continue
		}

		if input == "exit" {
			break
		}
		pos, err := chees.ProcessingInputToPosition(input)
		if err != nil {
			fmt.Printf("Some error happen, %v\n", err)
			continue
		}
		err = CheesField.SetPossibleMoves(pos)
		if err != nil {
			fmt.Printf("Some error happen, %v\n", err)
			continue
		}
		CheesField.PrintPossibleMoves()
		fmt.Print("Press Enter for continue > ")
		if _, err := fmt.Scanln(&input); err != nil {
			fmt.Println(err)
			continue
		}
	}

}

// [TASK 1]

// Storager - interface for storage
type Storager interface {
	readStorage() error
	saveStorage() error
	GetByID(int) (User, error)
	Insert(User) (int, error)
}

// [Implementation for JSON Storage]

// storageJSONfile - implementation fo Storager interface
type storageJSONfile struct {
	filename string
	data     []User
}

// readStorage - read data from file
func (storage *storageJSONfile) readStorage() (err error) {
	f, err := os.Open(storage.filename)
	if err != nil {
		return
	}
	defer f.Close()
	//Read file and Decode to structure
	filecontent, _ := ioutil.ReadAll(f)
	err = json.Unmarshal(filecontent, &storage.data)
	return
}

// readStorage - read data from file
func (storage storageJSONfile) saveStorage() (err error) {
	dataSave, err := json.Marshal(storage.data)
	//save to disk
	err = ioutil.WriteFile(storage.filename, dataSave, 0644)
	return
}

// GetByID - implement method for interface Storager
// returns User data by UID
func (storage storageJSONfile) GetByID(uid int) (u User, err error) {
	err = fmt.Errorf("User with ID=%d is't exist", uid)
	for _, usr := range storage.data {
		if usr.UID == uid {
			return usr, nil
		}
	}
	return u, err
}

// NewID - implement method for interface Storager
// returns User data by UID
func (storage storageJSONfile) NewID() (uid int) {
	for _, usr := range storage.data {
		if usr.UID > uid {
			uid = usr.UID
		}
	}
	uid++ //next ID
	return
}

// Insert - implement method for interface Storager Insert(User) (int, error)
// returns uid for new user in the storage
func (storage *storageJSONfile) Insert(usr User) (uid int, err error) {
	uid = storage.NewID()
	usr.UID = uid
	storage.data = append(storage.data, usr)
	err = storage.saveStorage()
	return
}

// [implementation for XML Storage]

// storageJSONfile - implementation fo Storager interface
type storageXMLfile struct {
	filename string
	data     []User
}

// readStorage - read data from file
func (storage *storageXMLfile) readStorage() (err error) {
	f, err := os.Open(storage.filename)
	if err != nil {
		return
	}
	defer f.Close()
	//Read file and Decode to structure
	filecontent, _ := ioutil.ReadAll(f)
	err = xml.Unmarshal(filecontent, &storage.data)
	return
}

// readStorage - read data from file
func (storage storageXMLfile) saveStorage() (err error) {
	dataSave, err := xml.Marshal(storage.data)
	//save to disk
	err = ioutil.WriteFile(storage.filename, dataSave, 0644)
	return
}

// GetByID - implement method for interface Storager
// returns User data by UID
func (storage storageXMLfile) GetByID(uid int) (u User, err error) {
	err = fmt.Errorf("User with ID=%d is't exist", uid)
	for _, usr := range storage.data {
		if usr.UID == uid {
			return usr, nil
		}
	}
	return u, err
}

// NewID - implement method for interface Storager
// returns User data by UID
func (storage storageXMLfile) NewID() (uid int) {
	for _, usr := range storage.data {
		if usr.UID > uid {
			uid = usr.UID
		}
	}
	uid++ //next ID
	return
}

// Insert - implement method for interface Storager Insert(User) (int, error)
// returns uid for new user in the storage
func (storage *storageXMLfile) Insert(usr User) (uid int, err error) {
	uid = storage.NewID()
	usr.UID = uid
	storage.data = append(storage.data, usr)
	err = storage.saveStorage()
	return
}

// [Task 2 Implement Sort{}]
/*
type Interface interface {
        // Len is the number of elements in the collection.
        Len() int
        // Less reports whether the element with
        // index i should sort before the element with index j.
        Less(i, j int) bool
        // Swap swaps the elements with indexes i and j.
        Swap(i, j int)
}
*/
// Len - count of contacts at the phonebook
func (pb PhoneBook) Len() int {
	return len(pb)
}

// Less
func (pb PhoneBook) Less(i, j int) bool {
	return pb[i].Age < pb[j].Age
}

// Swap
func (pb PhoneBook) Swap(i, j int) {
	pb[i], pb[j] = pb[j], pb[i]
}

// String - returns string representation of storage
func (pb PhoneBook) String() {
	for _, val := range pb {
		fmt.Printf("UID:%d, Name:%s, Age:%d\n", val.UID, val.NickName, val.Age)
	}
}

// [TASK 3 Evaluate calculator]

// Phone - type for phone numbers
type Phone string

// User - structure for describe of humans
type User struct {
	XMLName  xml.Name `xml:"Users" json:"-"`
	UID      int      `xml:"UId" json:"uid"`
	NickName string   `xml:"NickName" json:"nickname"`
	Age      int      `xml:"Age" json:"age"`
	Phones   []Phone  `xml:"Phones>Phone" json:"phones"`
}

// PhoneBook - type of slice Users
type PhoneBook []User
