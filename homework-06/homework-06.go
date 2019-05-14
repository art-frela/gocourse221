/*
homework-06

1. Дополнить код из раздела тестирование функцией подсчёта суммы переданных элементов и тестом для этой функции.
2. Дополнить пример из раздела пакет img изображением горизонтальных и вертикальных линий. Воспользуйтесь статьей https://4gophers.ru/articles/rabota-s-izobrazheniyami/.
3. Дополнить функцию hello() нашего http сервера так, чтобы принять и отобразить один GET параметр.
4. * Написать функцию для вычисления корней квадратного уравнения (алгоритм можно найти в википедии) и тесты к ней.
5. ** Написать программу, генерирующую png файл с рисунком шахматной доски.

Detail explain see README.md

Author: Karpov A. mailto:art.frela@gmail.com
Date: 2019-05-15
*/

package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

func main() {
	// [TASK 2]
	type rectBorder struct {
		x0, y0, x1, y1 int
		color color.RGBA
	}
	gray := color.RGBA{0, 0, 0, 0.1}
	ectangle := rectBorder{
		0,0,800,800, gray
	}

	green := color.RGBA{0, 255, 0, 255}
	red := color.RGBA{200, 30, 30, 255}
	rectImg := image.NewRGBA(image.Rect(0, 0, 200, 200))

	draw.Draw(rectImg, rectImg.Bounds(), &image.Uniform{green}, image.ZP, draw.Src)

	file, err := os.Create("rectangle.png")
	if err != nil {
		log.Fatalf("Failed create file: %s", err)
	}
    defer file.Close()
	png.Encode(file, rectImg)
}