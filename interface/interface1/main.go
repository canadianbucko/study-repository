/*
	— Необходимо разработать модуль для работы с геометрическими фигурами

— Каждая фигура умеет вычислять свою площадь и периметр
— Разные фигуры реализуются по-разному, но взаимодействие с ними должно быть унифицировано через интерфейс
— Необходимо описать фигуры и функцию, которая сможет принимать различные виды фигур и высчитывать их площать и периметр
*/
package main

import "fmt"

type ShapeModule interface {
	Area() float64
	Perimeter() float64
}

type FigureCircle struct {
	diameter float64
}
type FigureSquare struct {
	height float64
	width  float64
}

func (f FigureCircle) Area() float64 {
	area := 3.14 * f.diameter
	return area
}
func (f FigureCircle) Perimeter() float64 {
	perimeter := 2 * 3.14 * f.diameter
	return perimeter
}

func (f FigureSquare) Area() float64 {
	area := f.height * f.width
	return area
}

func (f FigureSquare) Perimeter() float64 {
	perimeter := (f.height + f.width) * 2
	return perimeter
}

func area(shape ShapeModule) float64 {
	fmt.Println("высчитываю площадь")
	area := shape.Area()
	return area

}

func perimeter(shape ShapeModule) float64 {
	fmt.Println("считаю периметр")
	perimeter := shape.Perimeter()
	return perimeter
}

func main() {
	fCircle := FigureCircle{
		diameter: 20,
	}

	fSquare := FigureSquare{
		height: 10,
		width:  15,
	}

	circleArea := area(fCircle)
	squareArea := area(fSquare)
	circlePerimetr := perimeter(fCircle)
	squarePerimetr := perimeter(fSquare)

	fmt.Println(squareArea)
	fmt.Println(squarePerimetr)
	fmt.Println(circlePerimetr)

	fmt.Println(circleArea)
}
