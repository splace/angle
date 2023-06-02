package angle

import "fmt"
import "strconv"

// Phase/Direction are restricted to a single revolution, no multi-turn. these types achieve this simply by using an unsigned int representation with its whole range representing one revolution.
// Phase's (like time.Duration) have a problem-space and a scaling center zero that are the same as the solution-space zero and so can be multipled.
// Direction's (like time.Time) have a common 'reference' zero but not a defined scaling center. making it possible to change the 'solution space' value that represents zero.
type Phase uint32

type Angle = Phase

const (
	bits          = 32
	Degree  Phase = 1 << (bits - 1) / 180
	Minute  Phase = 1 << (bits - 1) / (180 * 60)
	Second  Phase = 1 << (bits - 1) / (180 * 60 * 60)
	Radian  Phase = (2935890503282001408) >> (64 - bits) // math.MaxUint64 / (2 * math.Pi )
	Gradian Phase = 1 << (bits - 1) / 200

	// exact representation
	RightAngle   Phase = 1 << (bits - 2)
	BinaryDegree Phase = 1 << (bits - 8) // 256 per rotation.  equal to about about 1.42 degrees

)

func (a Phase) Degrees() float64 {
	return float64(a) / float64(Degree)
}

func (a Phase) Radians() float64 {
	return float64(a) / float64(Radian)
}

func (a Phase) Minutes() float64 {
	return float64(a) / float64(Minute)
}

func (a Phase) Seconds() float64 {
	return float64(a) / float64(Second)
}

func (a Phase) Gradians() float64 {
	return float64(a) / float64(Gradian)
}

func (a Phase) BinaryDegrees() float64 {
	return float64(a) / float64(BinaryDegree)
}

func (a Phase) Rotations() float64 {
	return float64(a) / (1 << bits)
}

// phase of fractionals of a rotation
func Rotations(f float64) Phase {
	return Phase(f * (1 << bits))
}

func (a Phase) Format(f fmt.State, r rune) {
	sfn, u := scalerAndUnit(r)
	if r == 'l' {
		fmt.Fprintf(f, `%+.0d%+.0m`, a, a%Degree)
		Phase(a%Minute).Format(f, 's')
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
