package angle


type Sweep struct{
	Angle
	CW
}

func (s Sweep) Contains(a Angle) bool {
	if s.Angle+s.CW.Angle>s.Angle{
		return (a>=s.Angle && a<s.CW.Angle)==s.CW.bool
	}
	return (a>=s.Angle || a<s.CW.Angle)==s.CW.bool
}

func (s Sweep) interpolate(divs,i uint) float64{ 
	return float64(s.CW.Angle)*float64(i)/float64(divs)
}

func (s Sweep) Intermediate(divs,i uint) Angle {
	if s.CW.bool{
		return s.Angle+Angle(s.interpolate(divs,i))
	}
	return s.Angle-Angle(s.interpolate(divs,i))
}

type CW struct{
	Angle
	bool
}


func NewClockwise(a Angle)CW{
	return CW{a,true}
}

func NewCounterClockwise(a Angle)CW{
	return CW{a,false}
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

