package angle

// angular region from an Angle up to a To.
type Sweep struct{
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

// an Angle with the direction required to get to it.
type To struct{
	Angle
	Direction
}

func (s Sweep) Contains(a Angle) bool {
	if s.Angle+s.To.Angle>s.Angle{
		return (a>=s.Angle && a<s.To.Angle)==s.To.Direction
	}
	return (a>=s.Angle || a<s.To.Angle)==s.To.Direction
}

func interpolate(a Angle,divs,i uint) Angle{ 
	return Angle(float64(a)*float64(i)/float64(divs))
}

func (s Sweep) Intermediate(divs,i uint) Angle {
	if s.To.Direction{
		return s.Angle+interpolate(s.To.Angle,divs,i)
	}
	return s.Angle-interpolate(s.To.Angle,divs,i)
}

func Over(s Sweep, steps uint) <-chan Angle{
	as:=make(chan Angle)
	go func(){
		for i:=uint(0);i<=steps;i++{
			as <- s.Intermediate(steps,i)
		}
		close(as)
	}()
	return (<-chan Angle)(as)
}

