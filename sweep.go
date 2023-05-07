package angle

type isClockwise bool

type To struct{
	Angle
	isClockwise
}

func NewCWTo(a Angle)To{
	return To{a,true}
}

func NewCCWTo(a Angle)To{
	return To{a,false}
}

type Sweep struct{
	Angle
	To
}

func (s Sweep) Contains(a Angle) bool {
	if s.Angle+s.To.Angle>s.Angle{
		return (a>=s.Angle && a<s.To.Angle)==s.To.isClockwise
	}
	return (a>=s.Angle || a<s.To.Angle)==s.To.isClockwise
}

func interpolate(a Angle,divs,i uint) Angle{ 
	return Angle(float64(a)*float64(i)/float64(divs))
}

func (s Sweep) Intermediate(divs,i uint) Angle {
	if s.To.isClockwise{
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

