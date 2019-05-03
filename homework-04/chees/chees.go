/*
package chees
implement simple logic behavior of chees figures on the field

Implement only for kNights (Horses)
Side of the field is ignored
Verification of occupied positions
Field has  coordinates of position in the digit format A=1, B=2 ...
Autor: Artem K., mailto:art.frela@gmail.com
Date: 2019-05-03
*/
package chees

import (
	"fmt"
	"strconv"
	"strings"
)

// Field - structure for chees filed
type Field struct {
	X           xAxis
	Y           yAxis
	Armies      [2]cheesArmy
	SelectedFig cheesFigure
}

type xAxis [8]int
type yAxis [8]int

// Position - coordinates for figures X - Y int
type Position [2]int

type cheesFigure struct {
	name                string
	count               int
	ixArmy              int
	ixPosition          int
	color               string
	directions          []moveDirection
	shifts              []moveShift
	currentPosition     Position
	perspectivePosition []Position
}

type moveDirection string

type moveShift [2]int

type cheesArmy struct {
	color   string
	figures []cheesFigure
}

// placeIsBorderOccupied - check position in the border of chees field and occupied,
// return true if in border and donn't occupied
func (chees Field) placeIsBorderAndFree(p Position) bool {
	//check is border
	isBorder := p[0] > 0 && p[0] <= len(chees.X) && p[1] > 0 && p[1] <= len(chees.Y)
	isFree := true
	for _, army := range chees.Armies {
		for _, fig := range army.figures {
			if fig.currentPosition == p {
				isFree = false
			}
		}
	}
	return isBorder && isFree
}

// defineFigureAtPosition - fill Army index and position index and return CheesFigure at that position
func (chees *Field) defineFigureAtPosition(pos Position) (fig cheesFigure) {
	for ixa, a := range chees.Armies {
		for ixp, f := range a.figures {
			if pos == f.currentPosition {
				chees.Armies[ixa].figures[ixp].ixPosition = ixp
				return f
			}
		}
	}
	return
}

// SetPossibleMoves - calculate and fill PerspectivePosition for figure on the current position
func (chees *Field) SetPossibleMoves(pos Position) (err error) {
	//iterate by movies
	var possPos []Position
	//define a figure at the specified position
	fig := chees.defineFigureAtPosition(pos)
	if fig.name == "" {
		err = fmt.Errorf("selected position <%v> is empty", pos)
		return
	}
	//
	for _, dir := range fig.directions {
		switch dir {
		case "OX":
			for _, x := range chees.X {
				if x != pos[0] {
					tmpPos := Position{x, pos[1]}
					if chees.placeIsBorderAndFree(tmpPos) {
						possPos = append(possPos, tmpPos)
					}

				}
			}
		case "OY":
			for _, y := range chees.Y {
				if y != pos[1] {
					tmpPos := Position{pos[0], y}
					if chees.placeIsBorderAndFree(tmpPos) {
						possPos = append(possPos, tmpPos)
					}

				}
			}
		case "XY":
			//don't implemented...
		}
	}
	//
	for _, shift := range fig.shifts {
		tmpPos := Position{pos[0] + shift[0], pos[1] + shift[1]}
		if chees.placeIsBorderAndFree(tmpPos) {
			possPos = append(possPos, tmpPos)
		}
	}
	//debug
	//fmt.Printf("DEBUG INFO: possPos is %v\n", possPos)
	chees.Armies[fig.ixArmy].figures[fig.ixPosition].perspectivePosition = possPos
	chees.SelectedFig = chees.defineFigureAtPosition(pos)
	chees.SelectedFig.perspectivePosition = possPos
	return
}

// PrintPossibleMoves for the selected figure
func (chees Field) PrintPossibleMoves() {
	fmt.Printf("Possible movies for %s (%s):\n", chees.SelectedFig.name, chees.SelectedFig.color)
	for i, v := range chees.SelectedFig.perspectivePosition {
		fmt.Printf("Option-%d = %v\n", i, v)
	}
}

// PrintField prints all figures at the field
func (chees Field) PrintField() {

	for _, a := range chees.Armies {
		fmt.Printf("Army of %s\n", a.color)
		for _, f := range a.figures {
			if f.name != "" {
				fmt.Printf("Figure [%s-%s] position is %v\n", f.name, f.color, f.currentPosition)
			}

		}
	}
}

// ProcessingInputToPosition - process input text and make Position element
func ProcessingInputToPosition(input string) (pos Position, err error) {
	poselements := strings.Split(input, "-")
	//fmt.Printf("slice %v, len(slice)=%d\n", poselements, len(poselements))
	if len(poselements) != 2 {
		err = fmt.Errorf("Wrong input format (%s), need x-y", input)
		return
	}
	pos[0], err = strconv.Atoi(poselements[0])
	if err != nil {
		return
	}
	pos[1], err = strconv.Atoi(poselements[1])
	return
}

// Init - initialize chees field
func (chees *Field) Init() {
	//make field axies
	chees.X = xAxis{1, 2, 3, 4, 5, 6, 7, 8}
	chees.Y = yAxis{1, 2, 3, 4, 5, 6, 7, 8}
	//make armies
	//white
	army := make([]cheesFigure, 16, 16)
	for i := 1; i <= len(chees.X); i++ {
		pawn := cheesFigure{
			name:            "pawn",
			count:           8,
			ixArmy:          0,
			color:           "white",
			shifts:          []moveShift{{0, 1}},
			currentPosition: Position{i, 2},
		}
		army = append(army, pawn)
	}

	army = append(army, cheesFigure{
		name:            "kNight",
		count:           2,
		ixArmy:          0,
		color:           "white",
		shifts:          []moveShift{{1, 2}, {2, 1}, {2, -1}, {1, -2}, {-1, -2}, {-2, -1}, {-2, 1}, {-1, 2}},
		currentPosition: Position{3, 5},
	})
	army = append(army, cheesFigure{
		name:            "kNight",
		count:           2,
		ixArmy:          0,
		color:           "white",
		shifts:          []moveShift{{1, 2}, {2, 1}, {2, -1}, {1, -2}, {-1, -2}, {-2, -1}, {-2, 1}, {-1, 2}},
		currentPosition: Position{4, 4},
	})

	whiteArmy := cheesArmy{
		color:   "white",
		figures: army,
	}

	//black army
	barmy := make([]cheesFigure, 16, 16)

	for i := 1; i <= len(chees.X); i++ {
		pawn := cheesFigure{
			name:            "pawn",
			count:           8,
			ixArmy:          1,
			color:           "black",
			shifts:          []moveShift{{0, 1}},
			currentPosition: Position{i, 7},
		}
		barmy = append(barmy, pawn)
	}

	barmy = append(barmy, cheesFigure{
		name:            "kNight",
		count:           2,
		ixArmy:          1,
		color:           "black",
		shifts:          []moveShift{{1, 2}, {2, 1}, {2, -1}, {1, -2}, {-1, -2}, {-2, -1}, {-2, 1}, {-1, 2}},
		currentPosition: Position{2, 8},
	})
	barmy = append(barmy, cheesFigure{
		name:            "kNight",
		count:           2,
		ixArmy:          1,
		color:           "black",
		shifts:          []moveShift{{1, 2}, {2, 1}, {2, -1}, {1, -2}, {-1, -2}, {-2, -1}, {-2, 1}, {-1, 2}},
		currentPosition: Position{5, 6},
	})

	blackArmy := cheesArmy{
		color:   "black",
		figures: barmy,
	}
	armies := [2]cheesArmy{whiteArmy, blackArmy}

	chees.Armies = armies
}
