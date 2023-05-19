package angle

// Angle exposes an angle type for problem-space angles.
// this is akin to time/duration (Angles commonly will be angle differences.)
// multiplying a duration is fine, but multiplying time is not the best type safety. hence here, unlike time/duration, angle is package-local.
// angle's (like time) have a common 'reference' zero but not a defined scaling center. making it possible to change the 'solution space' value that represents zero.
// Angle's (like duration) have a problem-space and a scaling center zero that are the same as the solution-space zero and so can be multipled.
// Example Sector: doubling the From (angle) makes no sense in the problem-space, but doubling the Width (Angle) clearly represents twice the sector size. 
type Angle struct{
	angle
}

type Delta = Angle

// Sector is an angular region From an angle and of a Delta (Angle), in either direction.
// notice: Delta is Clockwise. that means to get a small CCW delta, this is set to 1 rotation minus the required sweep angle. (due to modulus; simply -angle) 
// this is baked into the NewCCWSector() where the sweep angle parameter is then in the direction indicated by the constructor. 
// this allows sweeps of upto 1 rotation in either direction, using a signed var to indicate direction would only allow upto half a rotation in either direction.
type Sector struct {
	angle
	Delta
	Direction
}

func NewCWSector(s,d angle)Sector{
	return Sector{s,Angle{d},CW}
}

func NewCCWSector(s,d angle)Sector{
	return Sector{s,Angle{-d},CCW}
}

type Direction bool

const (
	Clockwise Direction = true
	CW
	CounterClockwise = false
	CCW
)


func (s Sector) Contains(a angle) bool {
	if s.angle+s.Delta.angle > s.angle {
		return (a >= s.angle && a <=s.Delta.angle) == s.Direction
	}
	// sector crosses zero.
	return (a >= s.angle || a <= s.Delta.angle) == s.Direction
}


// return a sequence of Angle's (one more than steps) evenly dividing a sector
// Note: usually can simply range using a fixed Angle step, this function reduces rounding errors when the divisions are very small.
// Direction species which way to sequence over the sector.
func Over(s Sector, steps uint) <-chan angle {
	as := make(chan angle)
	go func() {
		if s.Direction == CounterClockwise {
			for i := uint(0); i <= steps; i++ {
				as <- s.angle - angle(float64(-s.Delta.angle) * float64(i) / float64(steps))
			}
		}else{
			for i := uint(0); i <= steps; i++ {
				as <- s.angle + angle(float64(s.Delta.angle) * float64(i) / float64(steps))
			}
		}
		close(as)
	}()
	return (<-chan angle)(as)
}

func ReverseOver(s Sector, steps uint) <-chan angle {
	as := make(chan angle)
	go func() {
		if s.Direction == CounterClockwise {
			for i := steps; ; i-- {
				as <- s.angle - angle(float64(-s.Delta.angle) * float64(i) / float64(steps))
				if i==0 {break}
			}
			
		}else{
			for i := steps;; i-- {
				as <- s.angle + angle(float64(s.Delta.angle) * float64(i) / float64(steps))
				if i==0 {break}
			}
		}
		close(as)
	}()
	return (<-chan angle)(as)
}
