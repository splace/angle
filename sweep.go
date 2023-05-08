package angle

// an angular region from an Angle up to a To.
type Sector struct {
	angle
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
	angle
	Direction
}

// distinguishing type for Angles with a, potentially, problem-space defined zero.
type Angle angle

func NewOffset(a Angle,d Direction) To{
	return To{angle(a),d}
}

func (s Sector) Contains(a angle) bool {
	if s.angle+s.To.angle > s.angle {
		return (a >= s.angle && a < s.To.angle) == s.To.Direction
	}
	return (a >= s.angle || a < s.To.angle) == s.To.Direction
}

func interpolate(a angle, divs, i uint) angle {
	return angle(float64(a) * float64(i) / float64(divs))
}

func (s Sector) Intermediate(divs, i uint) angle {
	if s.To.Direction {
		return s.angle + interpolate(s.To.angle, divs, i)
	}
	return s.angle - interpolate(s.To.angle, divs, i)
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
