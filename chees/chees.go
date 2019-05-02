package chees

type CheesField struct {
	X      XAxis
	Y      YAxis
	Armies [2]CheesArmy
}

type XAxis [8]int
type YAxis [8]int

type Position [2]int

type CheesFigure struct {
	Name            string
	Count           int
	Directions      []MoveDirection
	Shifts          []MoveShift
	CurrentPosition Position
}

type MoveDirection string

type MoveShift [2]int

type CheesArmy struct {
	Color   string
	Figures []CheesFigure
}

var ChField CheesField

// placeIsBorderOccupied - check position in the border of chees field and occupied,
// return true if in border and donn't occupied
func (chees CheesField) placeIsBorderAndFree(p Position) bool {
	//check is border
	isBorder := p[0] > 0 && p[0] <= len(chees.X) && p[1] > 0 && p[1] <= len(chees.Y)
	isFree := false
	for _, army := range chees.Armies {
		for _, fig := range army.Figures {
			if fig.CurrentPosition != p {
				isFree = true
			}
		}
	}
	return isBorder && isFree
}

//
func (fig CheesFigure) PossibleMoves(pos Position, chees CheesField) (possPos []Position) {
	//iterate by movies
	for _, dir := range fig.Directions {
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
	for _, shift := range fig.Shifts {
		tmpPos := Position{pos[0] + shift[0], pos[1] + shift[1]}
		if chees.placeIsBorderAndFree(tmpPos) {
			possPos = append(possPos, tmpPos)
		}
	}
	return possPos
}

// init - initialize chees field
func (chees *CheesField) init() {
	//make field axies
	chees.X = XAxis{1, 2, 3, 4, 5, 6, 7, 8}
	chees.Y = YAxis{1, 2, 3, 4, 5, 6, 7, 8}
	//make armies
	//white
	pawns := make([]CheesFigure, 8, 8)
	for i := 1; i <= len(chees.X); i++ {
		pawn := CheesFigure{
			Name:            "pawn",
			Count:           8,
			Shifts:          []MoveShift{{0, 1}},
			CurrentPosition: Position{i, 2},
		}
		pawns = append(pawns, pawn)
	}

	kNights := make([]CheesFigure, 2, 2)
	kNights = append(kNights, CheesFigure{
		Name:            "kNight",
		Count:           2,
		Shifts:          []MoveShift{{1, 2}, {2, 1}, {2, -1}, {1, -2}, {-1, -2}, {-2, -1}, {-2, 1}, {-1, 2}},
		CurrentPosition: Position{2, 1},
	})
	kNights = append(kNights, CheesFigure{
		Name:            "kNight",
		Count:           2,
		Shifts:          []MoveShift{{1, 2}, {2, 1}, {2, -1}, {1, -2}, {-1, -2}, {-2, -1}, {-2, 1}, {-1, 2}},
		CurrentPosition: Position{7, 1},
	})

	armyfigs := make([]CheesFigure, 16, 16)
	army = append(army, pawns)
	army = append(army, kNights)

	whiteArmy := CheesArmy{
		Color:   "white",
		Figures: army,
	}

	//black army
	bpawns := make([]CheesFigure, 8, 8)
	for i := 1; i <= len(chees.X); i++ {
		pawn := CheesFigure{
			Name:            "pawn",
			Count:           8,
			Shifts:          []MoveShift{{0, 1}},
			CurrentPosition: Position{i, 7},
		}
		bpawns = append(bpawns, pawn)
	}

	bkNights := make([]CheesFigure, 2, 2)
	bkNights = append(bkNights, CheesFigure{
		Name:            "kNight",
		Count:           2,
		Shifts:          []MoveShift{{1, 2}, {2, 1}, {2, -1}, {1, -2}, {-1, -2}, {-2, -1}, {-2, 1}, {-1, 2}},
		CurrentPosition: Position{2, 1},
	})
	bkNights = append(kNights, CheesFigure{
		Name:            "kNight",
		Count:           2,
		Shifts:          []MoveShift{{1, 2}, {2, 1}, {2, -1}, {1, -2}, {-1, -2}, {-2, -1}, {-2, 1}, {-1, 2}},
		CurrentPosition: Position{7, 1},
	})

	barmy := make([]CheesFigure, 16, 16)
	barmy = append(army, bpawns)
	barmy = append(army, bkNights)

	blackArmy := CheesArmy{
		Color:   "black",
		Figures: barmy,
	}
	armies := make([]CheesArmy)
	armies = append(armies, whiteArmy)
	armies = append(armies, blackArmy)

	chees.Armies = armies
}

func init() {
	//init chees field, armies, figures...
	/*
		Фигура	Русское сокращение	Английское сокращение
		Король	Кр					K (king)
		Ферзь	Ф					Q (queen)
		Ладья	Л					R (rook)
		Конь	К					N (kNight)
		Слон	С					B (bishop)
		Пешка	п или ничего		p (pawn) или ничего
	*/

}
