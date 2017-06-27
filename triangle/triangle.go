package triangle

import "math"

const testVersion = 3

func KindFromSides(a, b, c float64) Kind {
	triangle := Triangle{a, b, c}
	switch {
	case !triangle.isValid():
		return NaT
	case triangle.isEquilateral():
		return Equ
	case triangle.isIsosceles():
		return Iso
	default:
		return Sca
	}
}

type Kind uint8

const (
	NaT = 0
	Equ = 1
	Iso = 2
	Sca = 3
)

type Triangle struct {
	side1, side2, side3 float64
}

func (t Triangle) isValid() bool {
	return t.side1 > 0 && t.side2 > 0 && t.side3 > 0 &&
		!math.IsInf(t.side1, 1) && !math.IsInf(t.side2, 1) && !math.IsInf(t.side3, 1) &&
		t.side1+t.side2 >= t.side3 &&
		t.side1+t.side3 >= t.side2 &&
		t.side2+t.side3 >= t.side1
}

func (t Triangle) isEquilateral() bool {
	return t.side1 == t.side2 && t.side2 == t.side3
}

func (t Triangle) isIsosceles() bool {
	return t.side1 == t.side2 || t.side2 == t.side3 || t.side3 == t.side1
}
