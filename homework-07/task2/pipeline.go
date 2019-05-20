/*CopyPast from treaching material
Перепишите программу-конвейер, ограничив количество передаваемых для обработки значений и обеспечив корректное завершение всех горутин.
*/
package main

import (
	"fmt"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// генерация
	go func() {
		for x := 0; x <= 10; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	// возведение в квадрат
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	// печать
	for q := range squares {
		fmt.Println(q)
	}
}
