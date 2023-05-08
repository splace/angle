package angle

// distinguishing type for Angles with a, potentially, problem-space defined zero, so accessible.
type Angle = angle

// an angular region from an Angle up to a To.
type Sector struct {
	Angle
	To
}

type Direction bool

const (
	Clockwise Direction = true
	CW
	ClounterClockwise = false
	CCW
)


type To struct {
	Angle
	Direction
}


func (s Sector) Contains(a angle) bool {
	if s.Angle+s.To.Angle > s.Angle {
		return (a >= s.Angle && a < s.To.Angle) == s.To.Direction
	}
	return (a >= s.Angle || a < s.To.Angle) == s.To.Direction
}

func interpolate(a angle, divs, i uint) angle {
	return angle(float64(a) * float64(i) / float64(divs))
}

func (s Sector) Intermediate(divs, i uint) angle {
	if s.To.Direction {
		return s.Angle + interpolate(s.To.Angle, divs, i)
	}
	return s.Angle - interpolate(s.To.Angle, divs, i)
}

func Over(s Sector, steps uint) <-chan angle {
	as := make(chan angle)
	go func() {
		for i := uint(0); i <= steps; i++ {
			as <- s.Intermediate(steps, i)
		}
		close(as)
	}()
	return (<-chan angle)(as)
}
