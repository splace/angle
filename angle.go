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

)

func (a Angle) Degrees() float64 {
	return float64(a) * 1.0 / float64(Degree)
}

func (a Angle) Radians() float64 {
	return float64(a) * 1.0 / float64(Radian)
}

func (a Angle) Minutes() float64 {
	return float64(a) * 1.0 / float64(Minute)
}

func (a Angle) Seconds() float64 {
	return float64(a) * 1.0 / float64(Second)
}

func (a Angle) Gradians() float64 {
	return float64(a) * 1.0 / float64(Gradian)
}

func (a Angle) BinaryDegrees() float64 {
	return float64(a) * 1.0 / float64(BinaryDegree)
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


