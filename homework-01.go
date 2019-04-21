/*
homework-01 is console util which doing three things:
- currency converter (rubles to usd dollars)
- calculates the area, perimeter, hypotenuse for a right triangle
- calculates the amount of the deposit in the bank and the annual interest

Author: Artem K mailto:art.frela@gmail.com
*/
package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {

	const (
		usdRate float64 = 64.01
	)

	var input string
	var programMode int

	var greeting string

	greetingSet(&greeting, 0, usdRate)
	//loop for continuous program execution
	for {
		fmt.Print(greeting)
		log.Println("Programmode=", programMode)
		fmt.Scanln(&input)
		if strings.ToUpper(input) == "Q" {
			fmt.Println("Good buy!")
			return
		}
		if strings.ToUpper(input) == "H" {
			greetingSet(&greeting, 0, usdRate)
			programMode = 0
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

		// switch programMode {
		// case 0:
		// 	greetingReset(&greeting)
		// 	continue
		// case 1:
		// 	greetingSet(1)
		// 	fmt.Println("")
		// }
		greeting := "<" + input + "> (main menu (h), exit (q))>"
		switch input {
		case "1":
			programMode = 1
			for {
				fmt.Printf("%s Type how much rubles do you want to convert to USD(USA), rate is %d rub/1 usd", greeting)
				fmt.Scanln(&input)

			}

		}
		// if err != nil {
		// 	fmt.Printf("Wrong input <%s>, need 1, 2, 3\n", inputOption)
		// 	continue
		// }
	}

}

func greetingSet(greeting *string, opt int, usdrate float64) {
	switch opt {
	case 0:
		*greeting = "You can choose which feature do you need:\n1-currency converter\n2-calculator for right triangle\n3-bank deposite calculator\n"
		*greeting += "Type some number 1, 2, 3 and press Enter. Main menu (h), for exit (q): "
	case 1:
		*greeting = "<1> (main menu (h), exit (q))>Currency calculartor. USD Rate =" + strconv.FormatFloat(usdrate, 'f', 2, 32) + "rub/1usd. Type how much rubles do you want convert to USD?:"
	case 2:
		*greeting = "You can choose which feature do you need:\n1-currency converter\n2-calculator for right triangle\n3-bank deposite calculator\n"
		*greeting += "Type some number 1, 2, 3 and press Enter. Main menu (h), for exit (q): "
	case 3:
		*greeting = "You can choose which feature do you need:\n1-currency converter\n2-calculator for right triangle\n3-bank deposite calculator\n"
		*greeting += "Type some number 1, 2, 3 and press Enter. Main menu (h), for exit (q): "
	}
}

//currencyConv - currency converter (rubles to usd dollars)
func currencyConv(usdRate, rubles float64) string {
	//only in one direction converter
	return strconv.FormatFloat(rubles/usdRate, 'f', 2, 32)
}
