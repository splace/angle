package angle

import "strconv"
import "fmt"

type angle uint32

const (
	bits          = 32 // allow simple generation of different precision packages
	Degree  angle = 1 << (bits - 2) / 90
	Minute  angle = 1 << (bits - 2) / (90 * 60)
	Second  angle = 1 << (bits - 2) / (90 * 60 * 60)
	Radian  angle = (2935890503282001408) >> (64 - bits) // math.MaxUint64 / (2 * math.Pi )
	Gradian angle = 1 << (bits - 2) / 100

	// exact representation
	RightAngle   angle = 1 << (bits - 2)
	BinaryDegree angle = 1 << (bits - 8) // 256 per rotation.  equal to about about 1.42 degrees

	// internal optimisation
	degreeRecip       = 1.0 / float64(Degree)
	minuteRecip       = 1.0 / float64(Minute)
	secondRecip       = 1.0 / float64(Second)
	radianRecip       = 1.0 / float64(Radian)
	gradianRecip      = 1.0 / float64(Gradian)
	binaryDegreeRecip = 1.0 / float64(BinaryDegree)
)

func (a angle) Format(f fmt.State, r rune) {
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
		a -= angle(a.Degrees()) * Degree
		fmt.Fprintf(f, "%+.0m", a)
		a -= angle(a.Minutes()) * Minute
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
	case 't':
		vfn = a.Rotations
		r = '⟳'
	case 'f':
		vfn = func()float64{return a.Rotations()*100}
		r = '%'
	case 'b':
		vfn = a.BinaryDegrees
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

func (a angle) Degrees() float64 {
	return float64(a) * degreeRecip
}

func (a angle) Radians() float64 {
	return float64(a) * radianRecip
}

func (a angle) Minutes() float64 {
	return float64(a) * minuteRecip
}

func (a angle) Seconds() float64 {
	return float64(a) * secondRecip
}

func (a angle) Gradians() float64 {
	return float64(a) * gradianRecip
}


func (a angle) BinaryDegrees() float64 {
	return float64(a) * binaryDegreeRecip
}

func (a angle) Rotations() float64 {
	return float64(a) / (1<<bits)
}

// angle of fractional rotations
func Rotations(f float64) angle{
	return angle(f*(1<<bits))
}

