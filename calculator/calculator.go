package calculator

import (
	// Пакет для работы с ошибками
	"errors" // Описания типов для построения бинарных деревьев
	"fmt"
	"go/ast"    // Пакет для обработки исходных данных
	"go/parser" // Содержит константы для математической лексики +, -, /, * и т.д.
	"go/token"  // Математические константы и выражения
	"math"      // Конвертер строк
	"strconv"   // Приведение имен констант к единому виду
	"strings"
)

// Calculate рассчитает данные выражения, переданного в строке
func Calculate(expr string) (float64, error) {
	// Поиск управляющих операторов в строке и построение дерева
	// Вернет корневой элемент в переменную root все математические действия, построенные
	// в виде дерева
	// Node будет содержать в себе действие для вычисления
	// exprNode() даст доступ к ветвям дерева, исходящим из корня/ноды
	//
	// type Expr interface {
	//  Node
	//  exprNode()
	// }
	root, err := parser.ParseExpr(expr)

	if err != nil {
		return -1, err
	} else {
		return eval(root)
	}
}

// HelpPrint - print decription of functionality of package Caclulator
func HelpPrint() {
	desc := "Калькулятор позволяет вычислять ряд выражений с заданными аргументами.\n"
	desc += "Для этого вам неободимо написать выражение которое вы хотите вычислить и нажать клавишу ввода\n"
	desc += "Допускается использование вложенных выражений max(sin(30), cos(30))\n"
	desc += "Вам доступны следующие примитивы:\n\n"
	for _, fun := range funcMap {
		desc += fmt.Sprintf("Функция: %s (%s)\nПример использования %s\n\n", fun.Name, fun.Help, fun.Example)
	}
	fmt.Println(desc)
}

type Func struct {
	Name    string
	Args    int
	Func    func(args ...float64) float64
	Help    string
	Example string
}

var funcMap map[string]Func

// Создание массива функций для обработки данных
func init() {
	funcMap = make(map[string]Func)
	funcMap["sqrt"] = Func{
		Name: "sqrt",
		Args: 1,
		Func: func(args ...float64) float64 {
			return math.Sqrt(args[0])
		},
		Help:    "Sqrt(x) returns the square root of x.",
		Example: "sqrt(9) //3",
	}
	funcMap["abs"] = Func{
		Name: "abs",
		Args: 1,
		Func: func(args ...float64) float64 {
			return math.Abs(args[0])
		},
		Help:    "Abs(x) returns the absolute value of x.",
		Example: "abs(-5) //5",
	}
	funcMap["log"] = Func{
		Name: "log",
		Args: 2,
		Func: func(args ...float64) float64 {
			return math.Log(args[0]) / math.Log(args[1])
		},
		Help:    "log(x, y) the ratio of natural logarithms of two numbers ln(x)/ln(y)",
		Example: "log(5,7) //0.827087...",
	}
	funcMap["ln"] = Func{
		Name: "ln",
		Args: 1,
		Func: func(args ...float64) float64 {
			return math.Log(args[0])
		},
		Help:    "Ln(x) returns the natural logarithm of x.",
		Example: "ln(5) //1.609437...",
	}
	funcMap["sin"] = Func{
		Name: "sin",
		Args: 1,
		Func: func(args ...float64) float64 {
			return math.Sin(args[0] * math.Pi / 180)
		},
		Help:    "sin(x) returns the sine of the degree argument x.",
		Example: "sin(30) //~0.5...",
	}
	funcMap["cos"] = Func{
		Name: "cos",
		Args: 1,
		Func: func(args ...float64) float64 {
			return math.Cos(args[0] * math.Pi / 180)
		},
		Help:    "cos(x) returns the cosine of the degree argument x.",
		Example: "cos(60) //~0.5...",
	}
	funcMap["max"] = Func{
		Name: "max",
		Args: 2,
		Func: func(args ...float64) float64 {
			return math.Max(args[0], args[1])
		},
		Help:    "max(x,y) returns the larger of x or y.",
		Example: "max(60, 30) //60...",
	}
	funcMap["min"] = Func{
		Name: "min",
		Args: 2,
		Func: func(args ...float64) float64 {
			return math.Min(args[0], args[1])
		},
		Help:    "min(x,y) returns the smaller of x or y.",
		Example: "min(60, 30) //30...",
	}
}

// Разбор полученных данных
func eval(expr ast.Expr) (float64, error) {
	switch expr.(type) {
	case *ast.BasicLit:
		return basic(expr.(*ast.BasicLit))
	// Вложенные вычисления
	case *ast.ParenExpr:
		return eval(expr.(*ast.ParenExpr).X)
	// Обработка посредством математических функций
	case *ast.CallExpr:
		return call(expr.(*ast.CallExpr))
	// Случай для обработки констант
	case *ast.Ident:
		return ident(expr.(*ast.Ident))
	default:
		return -1, errors.New("Не удалось распознать оператор")
	}
}

// Разбор чисел
func basic(lit *ast.BasicLit) (float64, error) {
	switch lit.Kind {
	case token.INT:
		i, err := strconv.ParseInt(lit.Value, 10, 64)

		if err != nil {
			return -1, err
		} else {
			return float64(i), nil
		}
	case token.FLOAT:
		i, err := strconv.ParseFloat(lit.Value, 64)

		if err != nil {
			return -1, err
		} else {
			return i, nil
		}
	default:
		return -1, errors.New("Неизвестный аргумент")
	}
}

// Обработка констант
func ident(id *ast.Ident) (float64, error) {
	switch n := strings.ToLower(id.Name); n {
	case "pi":
		return math.Pi, nil
	case "e":
		return math.E, nil
	case "phi":
		return math.Phi, nil
	default:
		return -1, errors.New("Неизвестная константа " + n)
	}
}

// Обработка функциональных операторов с помощью созданного массива функций для обработки
func call(c *ast.CallExpr) (float64, error) {
	switch t := c.Fun.(type) {
	case *ast.Ident:
	default:
		_ = t
		return -1, errors.New("Неизвестный тип функции")
	}

	ident := c.Fun.(*ast.Ident)

	args := make([]float64, len(c.Args))
	for i, expr := range c.Args {
		var err error
		args[i], err = eval(expr)
		if err != nil {
			return -1, err
		}
	}

	name := strings.ToLower(ident.Name)

	if val, ok := funcMap[name]; ok {
		if len(args) == val.Args {
			return val.Func(args...), nil
		} else {
			return -1, errors.New("Слишком много аргументов для " + name)
		}
	} else {
		return -1, errors.New("Неизвестная функция " + name)
	}
}
