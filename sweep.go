package angle

type sweep struct{
	start,end Angle
}

func (s sweep) interpolate(divs,i uint) float64{ 
	return float64(s.end-s.start)*float64(i)/float64(divs)
}

type SweepCW sweep

func (s SweepCW) Contains(a Angle) bool {
	if s.end>s.start{
		return a>=s.start && a<s.end
	}
	return a>=s.start || a<s.end
}

func (s SweepCW) Intermediate(divs,i uint) Angle {
	return s.start+Angle(sweep(s).interpolate(divs,i))
}

type SweepCCW sweep

func (s SweepCCW) Contains(a Angle) bool {
	return !SweepCW(sweep(s)).Contains(a)
}

func (s SweepCCW) Intermediate(divs,i uint) Angle {
	return s.start-Angle(sweep{s.end,s.start}.interpolate(divs,i))
}


