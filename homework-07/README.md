## HOMEWORK #7

### Tasks 

* [ ] Уберите из первого примера (Фибоначчи и спиннер) функцию, вычисляющую числа Фибоначчи. Как теперь заставить спиннер вращаться в течение некоего времени? В течение 10 секунд?
./task1/spinner.go -n 5
* [ ] Перепишите программу-конвейер, ограничив количество передаваемых для обработки значений и обеспечив корректное завершение всех горутин.
* [ ] Добавьте для time-сервера возможность его корректного завершения при вводе команды exit.
* [ ] * Доработайте фрагмент из раздела “Буферизованные каналы”, превратив его в полноценное приложение. Найдите три реальных зеркала (либо три разных сайта), замерьте их отклик, а затем протестируйте работу приложения, убедившись что функция mirroredQuery() действительно выполняет свою функцию.
* [ ] ** Попробуйте смоделировать с помощью sync.WaitGroup гонку автомобилей. Каждый автомобиль должен быть представлен отдельной горутиной со случайно устанавливаемой скоростью (math/rand) и случайным временем готовности к старту. Программа должна ожидать готовности всех машин, обеспечивать одновременный старт и фиксировать финиш каждой машины.