package angle

// distinguishing type for Angle's that have, potentially, problem-space defined zero, so can/should be accessible.
type Angle = angle

// an angular region from an Angle to a To.
type Sector struct {
	From angle
	Angle
	Direction
}

type Direction bool

const (
	Clockwise Direction = true
	CW
	ClounterClockwise = false
	CCW
)


func (s Sector) Contains(a angle) bool {
	if s.From+s.Angle > s.From {
		return (a >= s.From && a < s.Angle) == s.Direction
	}
	// sector crosses zero.
	return (a >= s.From || a < s.Angle) == s.Direction
}

func interpolate(a angle, divs, i uint) angle {
	return angle(float64(a) * float64(i) / float64(divs))
}

// return the angle for the indexed division
func (s Sector) Intermediate(divs, i uint) angle {
	if s.Direction {
		return s.From + interpolate(s.Angle, divs, i)
	}
	return s.From - interpolate(-s.Angle, divs, i)
}

// return a sequence of Angle's (one more than steps) evenly dividing a sector
// Note: usually can simply range using a fixed Angle step, this function reduces rounding errors when the divisions are very small. 
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
