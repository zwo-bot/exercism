// Package triangle checks the Kind of a triangle
package triangle

import "math"

type Kind string

const (
	NaT = "NaT" // not a triangle
	Equ = "Equ" // equilateral
	Iso = "Iso" // isosceles
	Sca = "Sca" // scalene
)

// KindKindFromSides returs the Kind of a triangle
func KindFromSides(a, b, c float64) Kind {
	var k Kind
	for _, v := range []float64{a, b, c} {
		if math.IsInf(v, 0) || math.IsNaN(v) || v < 0 {
			k = NaT
			return k
		}
	}
	if a+b < c || a+c < b || b+c < a || a+b+c == 0 {
		k = NaT
		return k
	}
	set := make(map[float64]bool)
	set[a] = true
	set[b] = true
	set[c] = true

	switch len(set) {
	case 1:
		k = Equ
	case 2:
		k = Iso
	case 3:
		k = Sca
	}
	return k
}
