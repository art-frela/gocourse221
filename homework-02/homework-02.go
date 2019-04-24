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

	const (
		maxFiboTbl   = 25
		maxFiboSlice = 93
		maxSimple    = 500
		maxSimpleIx  = 3571
		taskCount    = 5
		menuCmd      = "Main menu (h), for exit (q): "
		greeting0    = `You can choose which feature do you need:
1-check parity
2-check of divisibility by three
3-print the sequence of Fibonacci numbers
4-fill the slice with Fibonacci numbers
5-fill the slice with simple numbers (Sieve of Eratosfena)
Type some number 1, 2, 3, 4, 5 and press Enter.`
		greeting1 = "Check parity. Type your number and press Enter: "
		greeting2 = "Check of divisibility by three. Type your number and press Enter: "
		greeting3 = "Print the sequence of Fibonacci numbers N<25. Type your number and press Enter: "
		greeting4 = "Print the slice of Fibonacci numbers N<93. Type your number and press Enter: "
		greeting5 = "Sieve of Eratosfena. Type your number 2<N<500 and press Enter: "
	)
	var input string
	var programMode int
	var taskMenu = []string{
		greeting0 + " " + menuCmd,
		"<1> " + menuCmd + " " + greeting1,
		"<2> " + menuCmd + " " + greeting2,
		"<3> " + menuCmd + " " + greeting3,
		"<4> " + menuCmd + " " + greeting4,
		"<5> " + menuCmd + " " + greeting5,
	}

	// loop for continuous program execution
	for {
		fmt.Print(taskMenu[programMode])
		fmt.Scanln(&input)
		if strings.ToUpper(input) == "Q" {
			fmt.Println("Good buy!")
			return
		}
		if strings.ToUpper(input) == "H" {
			programMode = 0
			continue
		}
		//
		inputInt, err := strconv.Atoi(input)
		if err != nil {
			fmt.Printf("You miss, try again. Error (%v)\n", err)
			continue
		}

		if programMode == 0 { //main mode
			wrongInputIs := true
			for t := 1; t <= taskCount; t++ {
				if inputInt == t {
					programMode = t
					wrongInputIs = false
				}
			}
			if wrongInputIs {
				fmt.Println("Wrong choice <" + input + ">, try again!")
			}
			continue
		}

		if programMode == 1 { // divide by 2 (even/odd)
			evenOdd := checkRemain(inputInt, 2)
			verdict := fmt.Sprintf("Your number <%d> is even!\n", inputInt)
			if evenOdd {
				verdict = fmt.Sprintf("Your number <%d> is odd!\n", inputInt)
			}
			fmt.Println(verdict)
			continue
		}

		if programMode == 2 { // divide by 3
			div3 := checkRemain(inputInt, 3)
			verdict := fmt.Sprintf("Your number <%d> is divisibility by 3? Yes it is!\n", inputInt)
			if div3 {
				verdict = fmt.Sprintf("Your number <%d> is divisibility by 3? No it isn't!\n", inputInt)
			}
			fmt.Println(verdict)
			continue
		}

		if programMode == 3 { // print Fibonacci numbers
			if inputInt > maxFiboTbl {
				fmt.Printf("You miss, try again. Error (%v)\n", err)
				continue
			}
			fiboPrint(inputInt)
			continue
		}

		if programMode == 4 { // print slice with Fibonacci numbers
			if inputInt > maxFiboSlice {
				fmt.Printf("You miss, try again. Error (%v)\n", err)
				continue
			}
			fmt.Printf("%v\n", fiboFillSlice(inputInt))
			continue
		}

		if programMode == 5 { //
			if inputInt < 2 || inputInt > maxSimple {
				fmt.Printf("You miss, try again. Error (%v)\n", err)
				continue
			}
			fmt.Printf("%v\n", simpleNumbers(inputInt, maxSimpleIx))
			continue
		}

	}
}

// checkRemain - check parity, returns "odd" or "even" for Integer argument
func checkRemain(dividend, divisor int) (isremain bool) {
	if dividend%divisor > 0 {
		isremain = true
	}
	return
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

// fiboPrint - print Fibonacci N number, like table
// | n | 0 | 1 | 2 |...
// | Fn| 0 | 1 | 1 |...
func fiboPrint(n int) {
	nRow := "| " + textWrap("n", "green") + "\t| "
	fnRow := "| " + textWrap("Fn", "red") + "\t| "
	nRow += fmt.Sprintf(" %s\t|", textWrap("0", "green"))
	fnRow += fmt.Sprintf(" %s\t|", textWrap("0", "red"))
	if n == 0 {
		fmt.Println(nRow)
		fmt.Println(fnRow)
		return
	}
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
		si := strconv.Itoa(i + 1)
		sa := strconv.Itoa(a)
		nRow += fmt.Sprintf(" %s\t|", textWrap(si, "green"))
		fnRow += fmt.Sprintf(" %s\t|", textWrap(sa, "red"))
	}
	fmt.Println(nRow)
	fmt.Println(fnRow)
	return
}

// fiboFillSlice - returns slice of Fibonacci numbers
func fiboFillSlice(n int) []int {
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

// simpleNumbers - returns slice of N<100 simple numbers (Sieve of Eratosfena)
// use https://ru.wikipedia.org/wiki/%D0%A0%D0%B5%D1%88%D0%B5%D1%82%D0%BE_%D0%AD%D1%80%D0%B0%D1%82%D0%BE%D1%81%D1%84%D0%B5%D0%BD%D0%B0
func simpleNumbers(n, ix int) []int {
	if n < 2 { //protect n<2 argument
		return []int{0}
	}
	allSimple := make([]int, 0)
	baseSlice := make([]bool, ix+1)

	// mark not simple values as true
	for i := 2; i*i <= ix; i++ {
		for j := i * i; j <= ix; j += i {
			baseSlice[j] = true
		}
	}
	// fill allSimple slice with simple numbers
	for i, v := range baseSlice {
		if !v {
			allSimple = append(allSimple, i)
		}
	}
	res := allSimple[:n]
	return res
}
