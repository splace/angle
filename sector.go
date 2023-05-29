package angle


type Turn bool

const (
	Clockwise Turn = false
	CW
	CounterClockwise = true
	CCW
)


// Sector is an angular region starting at a Direction and with an Angle width.
// notice: the Angle is Clockwise.
type Sector struct {
	Direction
	Angle
}

func NewSector(from,too Direction, t Turn) Sector {
	if t==CW {
		return Sector{from, Angle(too-from)}
	}
	return Sector{too, Angle(from-too)}
}

func (s Sector) Contains(a Direction) bool {
	if u:=s.Direction+Direction(s.Angle); u < s.Direction {
		// sector crosses zero
		return a >= s.Direction || a <= u
	}else{
		return a >= s.Direction && a <= u
	}
}

// return a sequence of Angle's (one more than steps) evenly dividing a sector
// Note: usually can simply range using a fixed Angle step, this function reduces rounding errors when the divisions are very small.
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

var CWOver = Over

func CCWOver(s Sector, steps uint) <-chan Direction {
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
