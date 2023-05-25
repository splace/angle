package angle

import "fmt"
import "strconv"

// Angle exposes an angle type for problem-space angles.
// angle/Angle is akin to Time/Duration (Angle commonly will be angle differences.)
// multiplying a Duration is fine, but multiplying Time is not the best type safety. hence here, unlike Time/Duration, the base type, angle, is package-local.
// angle's (like Time) have a common 'reference' zero but not a defined scaling center. making it possible to change the 'solution space' value that represents zero.
// Angle's (like Duration) have a problem-space and a scaling center zero that are the same as the solution-space zero and so can be multipled.
// Example Sector: doubling the From (angle) makes no sense in the problem-space, but doubling the Delta (=Angle) clearly represents twice the sector size. 
type Angle angle

func (a Angle) Format(f fmt.State, r rune) {
	sfn,u:=scalerAndUnit(r)
	if p, set := f.Precision(); set {
		f.Write([]byte(strconv.FormatFloat(sfn(angle(a)), 'f', p, bits)))
	} else {
		f.Write([]byte(strconv.FormatFloat(sfn(angle(a)), 'f', -1, bits)))
	}
	if f.Flag('+') {
		fmt.Fprint(f, u)
	}
}

// just for documentation
type Delta = Angle

// Sector is an angular region From an angle and of a Delta (Angle), in either direction.
// notice: Delta is Clockwise. that means for CCW this is set to 1 rotation minus the required sweep angle (simply -angle). see NewSector()
// this allows sweeps of upto 1 rotation in either direction, using a signed var to indicate direction would only allow upto half a rotation in either direction.
type Sector struct {
	From angle
	Delta
	Direction
}

func NewSector(f,a angle,d Direction) Sector{
	if d==CCW {
		a=-a
	}
	return Sector{f,Delta(a),d}
}

type Direction bool

const (
	Clockwise Direction = true
	CW
	CounterClockwise = false
	CCW
)


func (s Sector) Contains(a angle) bool {
	if s.From+angle(s.Delta) > s.From {
		return (a >= s.From && a <= angle(s.Delta)) == s.Direction
	}
	// sector crosses zero.
	return (a >= s.From || a <= angle(s.Delta)) == s.Direction
}


// return a sequence of Angle's (one more than steps) evenly dividing a sector
// Note: usually can simply range using a fixed Angle step, this function reduces rounding errors when the divisions are very small.
func Over(s Sector, steps uint) <-chan angle {
	as := make(chan angle)
	go func() {
		div:=1.0 / float64(steps)
		if s.Direction == CounterClockwise {
			for i := uint(0); i <= steps; i++ {
				as <- s.From - angle(float64(-s.Delta) * float64(i) *div)
			}
		}else{
			for i := uint(0); i <= steps; i++ {
				as <- s.From + angle(float64(s.Delta) * float64(i) * div)
			}
		}
		close(as)
	}()
	return (<-chan angle)(as)
}

var CWOver = Over

func CCWOver(s Sector, steps uint) <-chan angle {
	as := make(chan angle)
	go func() {
		div:=1.0 / float64(steps)
		if s.Direction == CounterClockwise {
			for  ;; steps-- {
				as <- s.From - angle(float64(-s.Delta) * float64(steps) * div)
				if steps==0 {break}
			}
			
		}else{
			for ;;steps-- {
				as <- s.From + angle(float64(s.Delta) * float64(steps) * div)
				if steps==0 {break}
			}
		}
		close(as)
	}()
	return (<-chan angle)(as)
}
