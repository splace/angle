package angle

import "fmt"
import "strconv"

// Direction/Angle is akin to Time/Duration.
// multiplying a Angle is fine, but multiplying a Direction is not the best type safety.
// Direction's (like Time) have a common 'reference' zero but not a defined scaling center. making it possible to change the 'solution space' value that represents zero.
// Angle's (like Duration) have a problem-space and a scaling center zero that are the same as the solution-space zero and so can be multipled.
// Example Sector: doubling the Direction makes no sense in the problem-space, but doubling the Angle clearly represents twice the sector size.
type Angle Direction

const (
	bits          = 32
	Degree  Angle = 1 << (bits - 2) / 90
	Minute  Angle = 1 << (bits - 2) / (90 * 60)
	Second  Angle = 1 << (bits - 2) / (90 * 60 * 60)
	Radian  Angle = (2935890503282001408) >> (64 - bits) // math.MaxUint64 / (2 * math.Pi )
	Gradian Angle = 1 << (bits - 2) / 100

	// exact representation
	RightAngle   Angle = 1 << (bits - 2)
	BinaryDegree Angle = 1 << (bits - 8) // 256 per rotation.  equal to about about 1.42 degrees

	// internal optimisation
	degreeRecip       = 1.0 / float64(Degree)
	minuteRecip       = 1.0 / float64(Minute)
	secondRecip       = 1.0 / float64(Second)
	radianRecip       = 1.0 / float64(Radian)
	gradianRecip      = 1.0 / float64(Gradian)
	binaryDegreeRecip = 1.0 / float64(BinaryDegree)
)

func (a Angle) Degrees() float64 {
	return float64(a) * degreeRecip
}

func (a Angle) Radians() float64 {
	return float64(a) * radianRecip
}

func (a Angle) Minutes() float64 {
	return float64(a) * minuteRecip
}

func (a Angle) Seconds() float64 {
	return float64(a) * secondRecip
}

func (a Angle) Gradians() float64 {
	return float64(a) * gradianRecip
}

func (a Angle) BinaryDegrees() float64 {
	return float64(a) * binaryDegreeRecip
}

func (a Angle) Rotations() float64 {
	return float64(a) / (1 << bits)
}

// angle of fractional rotations
func Rotations(f float64) Angle {
	return Angle(f * (1 << bits))
}

func (a Angle) Format(f fmt.State, r rune) {
	sfn, u := scalerAndUnit(r)
	if r == 'l' {
		fmt.Fprintf(f, `%+.0d%+.0m`, a, a%Degree)
		Angle(a%Minute).Format(f, 's')
		return
	}
	if p, set := f.Precision(); set {
		f.Write([]byte(strconv.FormatFloat(sfn(a), 'f', p, bits)))
	} else {
		f.Write([]byte(strconv.FormatFloat(sfn(a), 'f', -1, bits)))
	}
	if f.Flag('+') {
		fmt.Fprint(f, u)
	}
}

// Sector is an angular region From a Direction of an Angle, turning either w ay.
// notice: the Angle is Clockwise. this means for CCW this is set to 1 rotation minus the required sweep Angle (simply -Angle). see NewSector()
// this allows sweeps of upto 1 rotation turning either way. using a signed var, to indicate turn way, would only allow upto half a rotation in either.
type Sector struct {
	From Direction
	Angle
	Turn
}

func NewSector(from Direction, by Angle, d Turn) Sector {
	if d == CCW {
		by = -by
	}
	return Sector{from, by, d}
}

type Turn bool

const (
	Clockwise Turn = true
	CW
	CounterClockwise = false
	CCW
)

func (s Sector) Contains(a Direction) bool {
	if s.From+Direction(s.Angle) > s.From {
		return (a >= s.From && a <= Direction(s.Angle)) == s.Turn
	}
	// sector crosses zero.
	return (a >= s.From || a <= Direction(s.Angle)) == s.Turn
}

// return a sequence of Angle's (one more than steps) evenly dividing a sector
// Note: usually can simply range using a fixed Angle step, this function reduces rounding errors when the divisions are very small.
func Over(s Sector, steps uint) <-chan Direction {
	as := make(chan Direction)
	go func() {
		div := 1.0 / float64(steps)
		if s.Turn == CounterClockwise {
			for i := uint(0); i <= steps; i++ {
				as <- s.From - Direction(float64(-s.Angle)*float64(i)*div)
			}
		} else {
			for i := uint(0); i <= steps; i++ {
				as <- s.From + Direction(float64(s.Angle)*float64(i)*div)
			}
		}
		close(as)
	}()
	return (<-chan Direction)(as)
}

var CWOver = Over

func CCWOver(s Sector, steps uint) <-chan Direction {
	as := make(chan Direction)
	go func() {
		div := 1.0 / float64(steps)
		if s.Turn == CounterClockwise {
			for ; ; steps-- {
				as <- s.From - Direction(float64(-s.Angle)*float64(steps)*div)
				if steps == 0 {
					break
				}
			}

		} else {
			for ; ; steps-- {
				as <- s.From + Direction(float64(s.Angle)*float64(steps)*div)
				if steps == 0 {
					break
				}
			}
		}
		close(as)
	}()
	return (<-chan Direction)(as)
}
