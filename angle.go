package angle

import "strconv"
import "fmt"

// # angles encoded as ints.
// An Angle is a uint32 with its whole range as one revolution.
// Since its max approaches one rotation, its modulus behaviour matches a rotation modulus, so you get restricted within one revolution automatically.
// Note: constants report an out of range error when used beyond one rotation, replace with variables.
// Angles dont make sense when multipled by other angles.
// Where a float representation would have higher precision the closer to zero value, Angle has fixed precision and also away from zero is more precise than a float32.
// 360 degrees (or 2Pi radians etc.) is just 0, and so is encoded/returned as 0 degrees. ( or 0 radians etc).
// Power of two fractions of a rotation are represented exactly, eg. 64*BinaryDegree==RightAngle, but in general multiplying a unit can result in an in-exact representation, eg. 90*Degree!=RightAngle, (but RightAngle/90==Degree) use the usual approachs to limit rounding errors.
type Angle uint32

const (
	bits         = 32  // allow simple generation of differnt precision packages
	Degree Angle = 1 << (bits - 2) / 90
	Minute Angle = 1 << (bits - 2) / (90 * 60)
	Second Angle = 1 << (bits - 2) / (90 * 60 * 60)
	Radian Angle = (2935890503282001408) >> (64 - bits) // math.MaxUint64 / (2 * math.Pi )
	Gradian Angle = 1 << (bits - 2) / 100

	// exact representions
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

func (a Angle) Format(f fmt.State, r rune) {
	var vfn func() float64
	switch r {
	case 'r':
		r = '㎭'
		fallthrough
	case '㎭':
		vfn = a.Radians
	case 'm':
		r = '′'
		fallthrough
	case '′':
		vfn = a.Minutes
	case 'g':
		r = 'ᵍ'
		fallthrough
	case 'ᵍ':
		vfn = a.Gradians
	case 'l':
		fmt.Fprintf(f, "%+.0d", a)
		a -= Angle(a.Degrees()) * Degree
		fmt.Fprintf(f, "%+.0m", a)
		a -= Angle(a.Minutes()) * Minute
		fallthrough
	case 's':
		r = '″'
		fallthrough
	case '″':
		vfn = a.Seconds
		r = '″'
	case 'c':
		switch a >> (bits - 5) {
		case 0, 31:
			f.Write([]byte("N"))
		case 1, 2:
			f.Write([]byte("NNE"))
		case 3, 4:
			f.Write([]byte("NE"))
		case 5, 6:
			f.Write([]byte("ENE"))
		case 7, 8:
			f.Write([]byte("E"))
		case 9, 10:
			f.Write([]byte("NSE"))
		case 11, 12:
			f.Write([]byte("SE"))
		case 13, 14:
			f.Write([]byte("SSE"))
		case 15, 16:
			f.Write([]byte("S"))
		case 17, 18:
			f.Write([]byte("SSW"))
		case 19, 20:
			f.Write([]byte("SW"))
		case 21, 22:
			f.Write([]byte("WSW"))
		case 23, 24:
			f.Write([]byte("W"))
		case 25, 26:
			f.Write([]byte("WNW"))
		case 27, 28:
			f.Write([]byte("NW"))
		case 29, 30:
			f.Write([]byte("NNW"))
		}
		return
	case 'b':
		vfn = a.BinaryDegree
	case 'd', 'v':
		r = '°'
		fallthrough
	case '°':
		fallthrough
	default:
		vfn = a.Degrees
	}
	if p, set := f.Precision(); set {
		f.Write([]byte(strconv.FormatFloat(vfn(), 'f', p, bits)))
	} else {
		f.Write([]byte(strconv.FormatFloat(vfn(), 'f', -1, bits)))
	}
	if f.Flag('+') {
		fmt.Fprint(f, string(r))
	}
}

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

func (a Angle) BinaryDegree() float64 {
	return float64(a) * binaryDegreeRecip
}

// helper for a range of angles that might cross zero
func (a Angle) Between(start,end Angle) bool {
	if end>start{
		return a>=start && a<=end
	}
	return a>=start || a<=end
}

