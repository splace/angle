package angle

type sweep struct{
	from,until Angle
}

func (s sweep) interpolate(divs,i uint) float64{ 
	return float64(s.until-s.from)*float64(i)/float64(divs)
}

type SweepCW sweep

func (s SweepCW) Contains(a Angle) bool {
	if s.until>s.from{
		return a>=s.from && a<s.until
	}
	return a>=s.from || a<s.until
}

func (s SweepCW) Intermediate(divs,i uint) Angle {
	return s.from+Angle(sweep(s).interpolate(divs,i))
}

type SweepCCW sweep

func (s SweepCCW) Contains(a Angle) bool {
	return !SweepCW(sweep(s)).Contains(a)
}

func (s SweepCCW) Intermediate(divs,i uint) Angle {
	return s.from-Angle(sweep{s.until,s.from}.interpolate(divs,i))
}


