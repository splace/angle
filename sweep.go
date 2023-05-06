package angle

type sweep struct{
	start,end Angle
}


type SweepCW sweep

func (s SweepCW) Contains(a Angle) bool {
	if s.end>s.start{
		return a>=s.start && a<=s.end
	}
	return a>=s.start || a<=s.end
}

func (s SweepCW) Intermediate(divs,i uint) Angle {
	return s.start+Angle(float64(Rotation)*interpolate(s.start,s.end,divs,i))
}

type SweepCCW sweep

func (s SweepCCW) Contains(a Angle) bool {
	if s.end<s.start{
		return a<=s.start && a>=s.end
	}
	return a<=s.start || a>=s.end
}

func (s SweepCCW) Intermediate(divs,i uint) Angle {
	return s.start-Angle(float64(Rotation)*interpolate(s.end,s.start,divs,i))
}

func interpolate(s,e Angle, divs,i uint) float64{ 
	return (e-s).Rotations()*float64(i)/float64(divs)
}

