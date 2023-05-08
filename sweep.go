package angle

// distinguishes not necessarily absolute angles  
type Angle = Absolute

// an angular region from an Angle up to a To.
type Sector struct {
	Absolute
	To
}

type Direction bool

const (
	Clockwise Direction = true
	CW
	ClounterClockwise = false
	CCW
)

// a relative Angle.
type To struct {
	Angle
	Direction
}

func (s Sector) Contains(a Absolute) bool {
	if s.Absolute+s.To.Angle > s.Absolute {
		return (a >= s.Absolute && a < s.To.Angle) == s.To.Direction
	}
	return (a >= s.Absolute || a < s.To.Angle) == s.To.Direction
}

func interpolate(a Absolute, divs, i uint) Absolute {
	return Absolute(float64(a) * float64(i) / float64(divs))
}

func (s Sector) Intermediate(divs, i uint) Absolute {
	if s.To.Direction {
		return s.Absolute + interpolate(s.To.Angle, divs, i)
	}
	return s.Absolute - interpolate(s.To.Angle, divs, i)
}

func Over(s Sector, steps uint) <-chan Absolute {
	as := make(chan Absolute)
	go func() {
		for i := uint(0); i <= steps; i++ {
			as <- s.Intermediate(steps, i)
		}
		close(as)
	}()
	return (<-chan Absolute)(as)
}
