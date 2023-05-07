package angle


type Sweep struct{
	Angle
	CW
}

func (s Sweep) Contains(a Angle) bool {
	if s.CW.bool{
		if s.Angle+s.CW.Angle>s.Angle{
			return a>=s.Angle && a<s.CW.Angle
		}
		return a>=s.Angle || a<s.CW.Angle
	}else{
		if s.Angle+s.CW.Angle>s.Angle{
			return !(a>=s.Angle && a<s.CW.Angle)
		}
		return !(a>=s.Angle || a<s.CW.Angle)
	}
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

type sweep [2]Angle

func (s sweep) interpolate(divs,i uint) float64{ 
	return float64(s[1]-s[0])*float64(i)/float64(divs)
}


type SweepCW sweep

func (s SweepCW) Contains(a Angle) bool {
	if s[1]>s[0]{
		return a>=s[0] && a<s[1]
	}
	return a>=s[0] || a<s[1]
}

func (s SweepCW) Intermediate(divs,i uint) Angle {
	return s[0]+Angle(sweep(s).interpolate(divs,i))
}

type SweepCCW sweep

func (s SweepCCW) Contains(a Angle) bool {
	return !SweepCW(sweep(s)).Contains(a)
}

func (s SweepCCW) Intermediate(divs,i uint) Angle {
	s[0],s[1]=s[1],s[0]
	return s[1]-Angle(sweep(s).interpolate(divs,i))
}

type Range interface{
	Contains(Angle)bool
	Intermediate(uint,uint) Angle
}

func NewRange(start Angle, offset CW) Range{
	if offset.bool{
		return SweepCW{start,start+offset.Angle}
	}
	return SweepCCW{start,start-offset.Angle}
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

