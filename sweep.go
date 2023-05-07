package angle

type Direction bool

type To struct{
	Angle
	Direction
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
		return (a>=s.Angle && a<s.To.Angle)==s.To.Direction
	}
	return (a>=s.Angle || a<s.To.Angle)==s.To.Direction
}

func (s Sweep) interpolate(divs,i uint) float64{ 
	return float64(s.To.Angle)*float64(i)/float64(divs)
}

func (s Sweep) Intermediate(divs,i uint) Angle {
	if s.To.Direction{
		return s.Angle+Angle(s.interpolate(divs,i))
	}
	return s.Angle-Angle(s.interpolate(divs,i))
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

