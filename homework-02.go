/*
homework-02 is console util which doing five things:
- check parity
- check of divisibility by three
- print the sequence of Fibonacci numbers
- fill the slice with Fibonacci numbers
- fill the slice with simple numbers

Author: Artem K mailto:art.frela@gmail.com
Date: 2019-04-24
*/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	const taskCount = 5

	var input string
	var programMode int
	var greeting string

	greetingSet(&greeting, 0)
	//loop for continuous program execution
	for {
		fmt.Print(greeting)
		fmt.Scanln(&input)
		if strings.ToUpper(input) == "Q" {
			fmt.Println("Good buy!")
			return
		}
		if strings.ToUpper(input) == "H" {
			greetingSet(&greeting, 0)
			programMode = 0
			continue
		}

		if programMode == 0 { //main mode
			inputInt, _ := strconv.Atoi(input)
			wrongInputIs := true
			for t := 1; t <= taskCount; t++ {
				if inputInt == t {
					programMode = t
					greetingSet(&greeting, t)
					wrongInputIs = false
				}
			}
			if wrongInputIs {
				fmt.Println("Wrong choice <" + input + ">, try again!")
			}
			continue
		}

		if programMode == 1 { //currency
			//process input to Int
			inputInt, err := strconv.Atoi(input)
			if err != nil {
				fmt.Printf("You miss, try again. Error (%v)\n", err)
				continue
			}
			evenOdd := checkEvenOdd(inputInt)
			fmt.Printf("Your number <%d> is %s!\n", inputInt, evenOdd)
			continue
		}

		if programMode == 2 { //divide by tree
			inputInt, err := strconv.Atoi(input)
			if err != nil {
				fmt.Printf("You miss, try again. Error (%v)\n", err)
				continue
			}
			div3 := divByTree(inputInt)
			fmt.Printf("Your number <%d> is divisibility by 3? %s!\n", inputInt, div3)
			continue
		}

		if programMode == 3 {
			inputInt, err := strconv.Atoi(input)
			if err != nil || inputInt > 25 {
				fmt.Printf("You miss, try again. Error (%v)\n", err)
				continue
			}
			fibPrint(inputInt)
			continue
		}

		if programMode == 4 {
			inputInt, err := strconv.Atoi(input)
			if err != nil || inputInt > 93 {
				fmt.Printf("You miss, try again. Error (%v)\n", err)
				continue
			}
			fmt.Printf("%v\n", fibFillSlice(inputInt))
			continue
		}

		if programMode == 5 {
			inputInt, err := strconv.Atoi(input)
			if err != nil || (inputInt > 120 || inputInt < 2) {
				fmt.Printf("You miss, try again. Error (%v)\n", err)
				continue
			}
			fmt.Printf("%v\n", simpleNumbers(inputInt))
			continue
		}

	}

}

func greetingSet(greeting *string, opt int) {
	shell := "\033[1;38;5;207mMain menu (h), for exit (q):\033[0m"
	switch opt {
	case 0:
		*greeting = "\033[1;34mYou can choose which feature do you need:\n1-check parity\n2-check of divisibility by three\n3-print the sequence of Fibonacci numbers\n"
		*greeting += "4-fill the slice with Fibonacci numbers\n5-fill the slice with simple numbers (Sieve of Eratosfena)\n"
		*greeting += "Type some number 1, 2, 3, 4, 5 and press Enter.\033[0m " + shell + " "
	case 1:
		*greeting = "\033[1;34m<1>\033[0m " + shell + "\033[1;38;5;220mCheck parity. Type your number and press Enter:\033[0m "
	case 2:
		*greeting = "\033[1;34m<2>\033[0m " + shell + "\033[1;38;5;220mCheck of divisibility by three. Type your number and press Enter:\033[0m "
	case 3:
		*greeting = "\033[1;34m<3>\033[0m " + shell + "\033[1;38;5;220mPrint the sequence of Fibonacci numbers N<25. Type your number and press Enter:\033[0m "
	case 4:
		*greeting = "\033[1;34m<4>\033[0m " + shell + "\033[1;38;5;220mPrint the slice of Fibonacci numbers N<93. Type your number and press Enter:\033[0m "
	case 5:
		*greeting = "\033[1;34m<5>\033[0m " + shell + "\033[1;38;5;220mSieve of Eratosfena. Type your number 2<N<121 and press Enter:\033[0m "
	}
}

//checkEvenOdd - check parity, returns "odd" or "even" for Integer argument
func checkEvenOdd(number int) (res string) {
	if number%2 > 0 {
		res = "\033[1;31modd\033[0m"
	} else {
		res = "\033[1;5;92meven\033[0m"
	}
	return
}

//divByTree - check of divisibility by three, return YES or NO
func divByTree(number int) (res string) {
	if number%3 > 0 {
		res = "\033[1;31mNO\033[0m"
	} else {
		res = "\033[1;5;92mYES\033[0m"
	}
	return
}

//fibPrint - print Fibonacci N number, like table
// | n | 0 | 1 | 2 |...
// | Fn| 0 | 1 | 1 |...
func fibPrint(n int) {
	nRow := "| \033[1;5;92mn\033[0m\t| "
	fnRow := "| \033[1;31mFn\033[0m\t| "
	if n == 0 {
		nRow += fmt.Sprintf(" \033[1;5;92m%d\033[0m\t|", 0)
		fnRow += fmt.Sprintf(" \033[1;31m%d\033[0m\t|", 0)
		fmt.Println(nRow)
		fmt.Println(fnRow)
		return
	}
	nRow += fmt.Sprintf(" \033[1;5;92m%d\033[0m\t|", 0)
	fnRow += fmt.Sprintf(" \033[1;31m%d\033[0m\t|", 0)
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
		nRow += fmt.Sprintf(" \033[1;5;92m%d\033[0m\t|", i+1)
		fnRow += fmt.Sprintf(" \033[1;31m%d\033[0m\t|", a)
	}
	fmt.Println(nRow)
	fmt.Println(fnRow)
	return
}

//fibFillSlice - returns slice of Fibonacci numbers
func fibFillSlice(n int) []int {
	fib := make([]int, n+1)
	if n == 0 {
		fib[0] = 0
		return fib
	}
	fib[0], fib[1] = 0, 1
	for i := 2; i <= n; i++ {
		fib[i] = fib[i-2] + fib[i-1]
	}
	return fib
}

//simpleNumbers - returns slice of simple numbers (Sieve of Eratosfena)
//use https://ru.wikipedia.org/wiki/%D0%A0%D0%B5%D1%88%D0%B5%D1%82%D0%BE_%D0%AD%D1%80%D0%B0%D1%82%D0%BE%D1%81%D1%84%D0%B5%D0%BD%D0%B0
func simpleNumbers(n int) []int {
	if n < 2 { //protect n<2 argument
		return []int{0}
	}
	res := make([]int, 0)
	baseSlice := make([]bool, n+1)
	//to fill boolean array with the true values
	for i := 2; i <= n; i++ {
		baseSlice[i] = true
	}
	//mark not simple values as false
	for i := 2; i*i <= n; i++ {
		for j := i * i; j <= n; j += i {
			baseSlice[j] = false
		}
	}
	//fill res slice with simple numbers
	for i, v := range baseSlice {
		if v {
			res = append(res, i)
		}
	}
	return res
}
