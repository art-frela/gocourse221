/*
homework-01 is console util which doing three things:
- currency converter (rubles to usd dollars)
- calculates the area, perimeter, hypotenuse for a right triangle
- calculates the amount of the deposit in the bank and the annual interest

Author: Artem K mailto:art.frela@gmail.com
Date: 2019-04-21
*/
package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {

	const (
		usdRate       float64 = 64.01
		termOfDeposit int     = 5
	)

	var input, inputB string
	var programMode int
	var greeting string

	greetingSet(&greeting, 0, usdRate)
	//loop for continuous program execution
	for {
		fmt.Print(greeting)
		fmt.Scanln(&input, &inputB)
		if strings.ToUpper(input) == "Q" {
			fmt.Println("Good buy!")
			return
		}
		if strings.ToUpper(input) == "H" {
			greetingSet(&greeting, 0, usdRate)
			programMode = 0
			continue
		}

		if programMode == 0 { //main mode
			//
			switch input {
			case "1":
				programMode = 1
				greetingSet(&greeting, 1, usdRate)
				continue
			case "2":
				programMode = 2
				greetingSet(&greeting, 2, usdRate)
				continue
			case "3":
				programMode = 3
				greetingSet(&greeting, 3, usdRate)
				continue
			default:
				fmt.Println("Wrong choice <" + input + ">, try again!")
				continue
			}
		}

		if programMode == 1 { //currency
			//process input to float64
			rubles, err := strconv.ParseFloat(input, 64)
			if err != nil {
				fmt.Printf("You miss, try again. Error (%v)\n", err)
				continue
			}
			usd := currencyConv(usdRate, rubles)
			fmt.Printf("On your %s rubles you can buy %s dollar(s)\n", strconv.FormatFloat(rubles, 'f', 2, 32), usd)
			continue

		}

		if programMode == 2 { //triangle
			catA, err := strconv.ParseFloat(input, 64)
			if err != nil {
				fmt.Printf("You miss, try again. Error (%v)\n", err)
				continue
			}
			catB, err := strconv.ParseFloat(inputB, 64)
			if err != nil {
				fmt.Printf("You miss, try again. Error (%v)\n", err)
				continue
			}
			hypo := triangleHypoten(catA, catB)
			area := triangleArea(catA, catB)
			perim := trianglePerim(catA, catB)

			fmt.Printf("Your triangle has: Cathetuses %f, %f\t Hypotenuse=%f\tArea=%f\tPerimeter=%f\n", catA, catB, hypo, area, perim)
			continue
		}

		if programMode == 3 {
			amount, err := strconv.ParseFloat(input, 64)
			if err != nil {
				fmt.Printf("You miss, try again. Error (%v)\n", err)
				continue
			}
			percent, err := strconv.ParseFloat(inputB, 64)
			if err != nil {
				fmt.Printf("You miss, try again. Error (%v)\n", err)
				continue
			}

			profitCalc(amount, percent, termOfDeposit)
			continue
		}

	}

}

func greetingSet(greeting *string, opt int, usdrate float64) {
	switch opt {
	case 0:
		*greeting = "You can choose which feature do you need:\n1-currency converter\n2-calculator for right triangle\n3-bank deposite calculator\n"
		*greeting += "Type some number 1, 2, 3 and press Enter. Main menu (h), for exit (q): "
	case 1:
		*greeting = "<1> (main menu (h), exit (q))>Currency calculator. USD Rate=" + strconv.FormatFloat(usdrate, 'f', 2, 32) + "rub/1usd. Type how much rubles do you want convert to USD?: "
	case 2:
		*greeting = "<2> (main menu (h), exit (q))>Right triangle. Type value of 2 cathetus: "
	case 3:
		*greeting = "<3> (main menu (h), exit (q))>Bank deposit calculator (5yrs),"
		*greeting += "type your amount (>0) and percent (0%-50%) and press Enter: "
	}
}

//currencyConv - currency converter (rubles to usd dollars)
func currencyConv(usdRate, rubles float64) string {
	//only in one direction converter
	return strconv.FormatFloat(rubles/usdRate, 'f', 2, 32)
}

//triangleHypoten - calculates hypotenuse for right triangle
func triangleHypoten(a, b float64) (c float64) {
	if a <= 0 || b <= 0 { //simple validation of input values
		return c //default = 0
	}
	c = math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
	return c
}

//triangleArea - calculates area for right triangle
func triangleArea(a, b float64) (c float64) {
	if a <= 0 || b <= 0 { //simple validation of input values
		return c //default = 0
	}
	c = a * b / 2
	return c
}

//trianglePerim - calculates perimeter for right triangle
func trianglePerim(a, b float64) (p float64) {
	if a <= 0 || b <= 0 { //simple validation of input values
		return p //default = 0
	}
	c := triangleHypoten(a, b)
	p = a + b + c
	return p
}

//profitCalc - calculates the profit by the amount and interest and a fixed deposit period of 5 years.
func profitCalc(amount, percent float64, term int) {
	if amount <= 0 || percent > 50 { //simple validation of input values
		fmt.Println("Wrong values, try again!")
		return
	}

	var resultSum = amount
	var perYearSum float64
	percent /= 100
	for y := 1; y <= term; y++ { //thousandths and smaller fractions were counted
		perYearSum = resultSum * percent
		resultSum += perYearSum
		fmt.Printf("For %d year you earn %fye and can withdraw from the account %fye\n", y, perYearSum, resultSum)
	}
	profit := resultSum - amount
	fmt.Printf("The total profit for %d years was %s ye\n", term, strconv.FormatFloat(profit, 'f', 2, 32))
}
