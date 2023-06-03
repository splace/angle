package angle

type Turn bool

const (
	Clockwise Turn = false
	CW
	CounterClockwise = true
	CCW
)

// Sector is an angular region starting at a Direction having an Angle width.
type Sector struct {
	Direction
	Angle
}

// Sector between two Directions.
// using the Turn (WRT the first parameter direction) to indicate which of the two possible regions.
func NewSector(from, too Direction, t Turn) Sector {
	if t == CW {
		return Sector{from, Angle(too - from)}
	}
	return Sector{too, Angle(from - too)}
}

func (s Sector) Contains(a Direction) bool {
	if u := s.Direction + Direction(s.Angle); u < s.Direction {
		// sector crosses zero
		return a >= s.Direction || a <= u
	} else {
		return a >= s.Direction && a <= u
	}
}

func (s Sector) Reverse()Sector{
	return Sector{Direction(s.Angle+Angle(s.Direction)),-s.Angle}  
}

// return Direction's (one more than steps) evenly dividing a sector
// Note: can simply range using a fixed Angle step, but this function can be used to reduce rounding errors particularly when the divisions are very small.
func Over(s Sector, steps uint) <-chan Direction {
	as := make(chan Direction)
	go func() {
		div := 1.0 / float64(steps)
		for i := uint(0); i <= steps; i++ {
			as <- s.Direction + Direction(float64(s.Angle)*float64(i)*div)
		}
		close(as)
	}()
	return (<-chan Direction)(as)
}

var OverCW = Over

func OverCCW(s Sector, steps uint) <-chan Direction {
	as := make(chan Direction)
	go func() {
		div := 1.0 / float64(steps)
		for ; ; steps-- {
			as <- s.Direction + Direction(float64(s.Angle)*float64(steps)*div)
			if steps == 0 {
				break
			}
		}
		close(as)
	}()
	return (<-chan Direction)(as)
}
