/*CopyPast from teaching material and modifying
Уберите из первого примера (Фибоначчи и спиннер) функцию, вычисляющую числа Фибоначчи.
Как теперь заставить спиннер вращаться в течение некоего времени? В течение 10 секунд?
*/
package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	n := flag.Int("n", 10, "[n] seconds duration of spinner")
	flag.Parse()
	go spinner(50 * time.Millisecond)
	time.Sleep(time.Duration(*n) * time.Second)
	fmt.Println("done")
}

func spinner(delay time.Duration) {
	for {
		for _, r := range "-\\|/" {
			fmt.Printf("%c\r", r)
			time.Sleep(delay)
		}
	}
}
