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
	"fmt"
	"html/template"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	// [TASK 2]
	err := drawRectangleWithLines("yellow", "blue", 400, 300, 4, "yellowrect.png")
	if err != nil {
		log.Fatalf("Draw rect error, %v", err)
	}

	// [TASK 3+5]
	fs := http.FileServer(http.Dir("img"))
	http.Handle("/", fs)
	http.HandleFunc("/hello/", helloPicture)
	http.HandleFunc("/chess/", helloChessPicture)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

// helloPicture - makes rectangle and show it on the page + parameters
func helloPicture(w http.ResponseWriter, r *http.Request) {
	// try get needed parameters for make specified rectangle
	bgColor := r.URL.Query().Get("bgcolor")
	lnColor := r.URL.Query().Get("lncolor")
	parts := r.URL.Query().Get("count")
	// get all keys
	keys := r.URL.Query()
	// DataT - structure for force to template parser
	type DataT struct {
		Title string
		Query []string
		Image string
	}
	// params slice
	params := make([]string, len(keys))
	i := 0
	// fill params of values
	for key, value := range keys {
		params[i] = fmt.Sprintf("%s=%s", key, value)
		i++
	}
	// template name
	tName := "task35.gohtml"
	fileimage := "httpRectangle.png"
	fileimagedir := "./img/" + fileimage
	data := DataT{
		Title: "Task-03",
		Query: params,
		Image: fileimage,
	}
	// parseInt from string
	intParts, _ := strconv.Atoi(parts)

	// make img for request parameters
	err := drawRectangleWithLines(bgColor, lnColor, 400, 400, intParts, fileimagedir)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "<html><body><h3 color=\"red\">Error: %v</h3></body></html>", err)
		return
	}
	// prepare template
	myT := template.Must(template.ParseGlob("./templates/*"))
	w.WriteHeader(http.StatusOK)                // 200 OK
	w.Header().Set("Content-Type", "text/html") //
	err = myT.ExecuteTemplate(w, tName, data)   // write to response body template with data
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "<html><body><h3 color=\"red\">Error: %v</h3></body></html>", err)
		return
	}

}

// helloChessPicture - makes chesslike rectangle and show it on the page + parameters
func helloChessPicture(w http.ResponseWriter, r *http.Request) {
	// try get needed parameters for make specified rectangle
	whiteColor := r.URL.Query().Get("whitecolor")
	blackColor := r.URL.Query().Get("blackcolor")
	parts := r.URL.Query().Get("count")
	// get all keys
	keys := r.URL.Query()
	// DataT - structure for force to template parser
	type DataT struct {
		Title string
		Query []string
		Image string
	}
	// params slice
	params := make([]string, len(keys))
	i := 0
	// fill params of values
	for key, value := range keys {
		params[i] = fmt.Sprintf("%s=%s", key, value)
		i++
	}
	// template name
	tName := "task35.gohtml"
	fileimage := "httpChessRectangle.png"
	fileimagedir := "./img/" + fileimage
	data := DataT{
		Title: "Task-05",
		Query: params,
		Image: fileimage,
	}
	// parseInt from string
	intParts, _ := strconv.Atoi(parts)

	// make img for request parameters
	err := drawChessLikeField(whiteColor, blackColor, 400, 400, intParts, fileimagedir)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "<html><body><h3 color=\"red\">Error: %v</h3></body></html>", err)
		return
	}
	// prepare template
	myT := template.Must(template.ParseGlob("./templates/*"))
	w.WriteHeader(http.StatusOK)                // 200 OK
	w.Header().Set("Content-Type", "text/html") //
	err = myT.ExecuteTemplate(w, tName, data)   // write to response body template with data
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "<html><body><h3 color=\"red\">Error: %v</h3></body></html>", err)
		return
	}

}

// drawRectangle With Lines - makes a rectangle with a COLOR background, with LNCOLOR lines that divide the field into countParts elements along each axis
// save the result to a file named filename
func drawRectangleWithLines(bgcolor, lncolor string, dx, dy, countParts int, filename string) error {
	type rectBorder struct {
		x0, y0, x1, y1 int
		color          color.RGBA
	}
	// colors map for simplified rgba decode
	colors := make(map[string]color.RGBA)
	colors["green"] = color.RGBA{0, 255, 0, 255}
	colors["red"] = color.RGBA{200, 30, 30, 255}
	colors["gray"] = color.RGBA{192, 192, 192, 255}
	colors["white"] = color.RGBA{255, 255, 255, 255}
	colors["black"] = color.RGBA{0, 0, 0, 255}
	colors["blue"] = color.RGBA{0, 0, 255, 255}
	colors["yellow"] = color.RGBA{255, 255, 0, 255}
	colors["pink"] = color.RGBA{255, 0, 255, 255}

	bgColorRGBA, ok := colors[bgcolor]
	if !ok {
		bgColorRGBA = colors["gray"]
	}

	lnColorRGBA, ok := colors[lncolor]
	if !ok {
		lnColorRGBA = colors["red"]
	}
	if countParts == 0 {
		countParts = 8
	}

	rectangle := rectBorder{
		0, 0, dx, dy, bgColorRGBA,
	}

	rectImg := image.NewRGBA(image.Rect(rectangle.x0, rectangle.y0, rectangle.x1, rectangle.y1))

	draw.Draw(rectImg, rectImg.Bounds(), &image.Uniform{rectangle.color}, image.ZP, draw.Src)

	dX := rectangle.x1 - rectangle.x0
	dY := rectangle.y1 - rectangle.y0
	stepX := dX / countParts
	stepY := dY / countParts

	// draw horizont lines
	for y := stepY; y <= dY-stepY; y += stepY {
		for ddx := 0; ddx < dX; ddx++ {
			rectImg.Set(ddx, y, lnColorRGBA)
		}

	}
	// draw vertical lines
	for x := stepX; x <= dX-stepX; x += stepX {
		for ddy := 0; ddy < dY; ddy++ {
			rectImg.Set(x, ddy, lnColorRGBA)
		}

	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	png.Encode(file, rectImg)
	return nil
}

// drawChessLikeField - makes a rectangle with a WHITECOLOR background, with BLACKCOLOR rectangles that fill the field
// save the result to a file named filename
func drawChessLikeField(whitecolor, blackcolor string, dx, dy, countParts int, filename string) error {
	type rectBorder struct {
		x0, y0, x1, y1 int
		color          color.RGBA
	}
	// colors map for simplified rgba decode
	colors := make(map[string]color.RGBA)
	colors["green"] = color.RGBA{0, 255, 0, 255}
	colors["red"] = color.RGBA{200, 30, 30, 255}
	colors["gray"] = color.RGBA{192, 192, 192, 255}
	colors["white"] = color.RGBA{255, 255, 255, 255}
	colors["black"] = color.RGBA{0, 0, 0, 255}
	colors["blue"] = color.RGBA{0, 0, 255, 255}
	colors["yellow"] = color.RGBA{255, 255, 0, 255}
	colors["pink"] = color.RGBA{255, 0, 255, 255}

	whiteColorRGBA, ok := colors[whitecolor]
	if !ok {
		whiteColorRGBA = colors["white"]
	}

	blackColorRGBA, ok := colors[blackcolor]
	if !ok {
		blackColorRGBA = colors["black"]
	}

	lnColorRGBA, ok := colors[blackcolor]
	if !ok {
		lnColorRGBA = colors["black"]
	}
	if countParts <= 0 {
		countParts = 8
	}

	rectangle := rectBorder{
		0, 0, dx, dy, whiteColorRGBA,
	}
	// field prepare
	rectImg := image.NewRGBA(image.Rect(rectangle.x0, rectangle.y0, rectangle.x1, rectangle.y1))
	// draw background white color
	// func Draw(dst Image, r image.Rectangle, src image.Image, sp image.Point, op Op)
	draw.Draw(rectImg, rectImg.Bounds(), &image.Uniform{rectangle.color}, image.ZP, draw.Src)

	dX := rectangle.x1 - rectangle.x0
	dY := rectangle.y1 - rectangle.y0
	stepX := dX / countParts
	stepY := dY / countParts

	// prepare coordinates of for black rectangles
	countBlackFigs := countParts * countParts / 2
	coords := make([]rectBorder, countBlackFigs)

	for i := 0; i < countParts; i++ {
		for j := 0; j < countParts; j++ {
			ij := (i + j) % 2
			if !(ij > 0) {
				coords = append(coords, rectBorder{
					stepX * j, stepY * i, stepX*j + stepX, stepY*i + stepY, blackColorRGBA,
				})
			}
		}
	}

	// field prepare

	mask := image.NewRGBA(image.Rect(rectangle.x0, rectangle.y0, rectangle.x1, rectangle.y1))
	draw.Draw(mask, mask.Bounds(), &image.Uniform{blackColorRGBA}, image.ZP, draw.Src)
	// func Draw(dst Image, r image.Rectangle, src image.Image, sp image.Point, op Op)
	// func DrawMask(dst Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, op Op)
	// draw black rectangls on the white field
	for _, coord := range coords {
		blackImg := image.NewRGBA(image.Rect(coord.x0, coord.y0, coord.x1, coord.y1))
		draw.Draw(blackImg, blackImg.Bounds(), &image.Uniform{blackColorRGBA}, image.ZP, draw.Src)

		draw.DrawMask(rectImg, rectImg.Bounds(), mask, image.ZP, blackImg, image.ZP, draw.Over)
	}

	// draw horizont lines
	for y := 0; y <= dY; y += stepY {
		for ddx := 0; ddx < dX; ddx++ {
			rectImg.Set(ddx, y, lnColorRGBA)
		}

	}
	// draw vertical lines
	for x := 0; x <= dX; x += stepX {
		for ddy := 0; ddy < dY; ddy++ {
			rectImg.Set(x, ddy, lnColorRGBA)
		}

	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	png.Encode(file, rectImg)
	return nil
}
