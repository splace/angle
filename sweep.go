package angle

type Direction bool

type Turn struct{
	Angle
	Direction
}

func NewClockwise(a Angle)Turn{
	return Turn{a,true}
}

func NewCounterClockwise(a Angle)Turn{
	return Turn{a,false}
}


type Sweep struct{
	Angle
	Turn
}

func (s Sweep) Contains(a Angle) bool {
	if s.Angle+s.Turn.Angle>s.Angle{
		return (a>=s.Angle && a<s.Turn.Angle)==s.Turn.Direction
	}
	return (a>=s.Angle || a<s.Turn.Angle)==s.Turn.Direction
}

func (s Sweep) interpolate(divs,i uint) float64{ 
	return float64(s.Turn.Angle)*float64(i)/float64(divs)
}

func (s Sweep) Intermediate(divs,i uint) Angle {
	if s.Turn.Direction{
		return s.Angle+Angle(s.interpolate(divs,i))
	}
	return s.Angle-Angle(s.interpolate(divs,i))
}



func Over(r interface{Intermediate(uint,uint) Angle}, steps uint) <-chan Angle{
	as:=make(chan Angle)
	go func(){
		for i:=uint(0);i<=steps;i++{
			as <- r.Intermediate(steps,i)
		}
		close(as)
	}()
	return (<-chan Angle)(as)
}

