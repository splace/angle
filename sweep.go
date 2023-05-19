package angle

// Angle exposes an angle type for problem-space angles.
// this is akin to time/duration (Angles commonly will be angle differences.)
// multiplying a duration is fine, but multiplying time is not the best type safety. hence here, unlike time/duration, angle is package-local.
// angle's (like time) have a common 'reference' zero but not a defined scaling center. making it possible to change the 'solution space' value that represents zero.
// Angle's (like duration) have a problem-space and a scaling center zero that are the same as the solution-space zero and so can be multipled.
// Example Sector: doubling the From (angle) makes no sense in the problem-space, but doubling the Width (Angle) clearly represents twice the sector size. 
type Angle angle

// Sector is an angular region From an angle of a Width Angle, in either direction.
// notice: Width Angle is always positive. small CCW sweeps are stored as a large Angle. (because of the modular behaviour this large Angle is simply -Angle))
// this allows sweeps of upto 1 rotation in either direction, using a signed var to indicate direction would only allow upto half a rotation in either direction.
type Sector struct {
	From angle
	Width Angle
	Direction
}

func NewCWSector(s,d angle)Sector{
	return Sector{s,Angle(d),CW}
}

func NewCCWSector(s,d angle)Sector{
	return Sector{s,Angle(-d),CCW}
}

type Direction bool

const (
	Clockwise Direction = true
	CW
	ClounterClockwise = false
	CCW
)


func (s Sector) Contains(a angle) bool {
	if s.From+angle(s.Width) > s.From {
		return (a >= s.From && a <= angle(s.Width)) == s.Direction
	}
	// sector crosses zero.
	return (a >= s.From || a <= angle(s.Width)) == s.Direction
}

func interpolate(a angle, divs, i uint) angle {
	return angle(float64(a) * float64(i) / float64(divs))
}

// return the angle for the indexed division
func (s Sector) Intermediate(divs, i uint) angle {
	if s.Direction {
		return s.From + interpolate(angle(s.Width), divs, i)
	}
	return s.From - interpolate(-angle(s.Width), divs, i)
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
