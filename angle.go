package angle

import "fmt"
import "io"

type angle uint32

const (
	bits          = 32
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

func scalerAndUnit(scaler rune) (func(angle)float64, string){
	switch scaler {
	case 'r','㎭':
		return angle.Radians,"㎭"
	case 'm','′':
		return angle.Minutes,"′"
	case 's','″':
		return angle.Seconds,"″"
	case 'g','ᵍ':
		return angle.Gradians,"ᵍ"
	case 't':
		return angle.Rotations, "⟳"
	case 'f':
		return func(a angle)float64{return a.Rotations()*100},"%"
	case 'b':
		return angle.BinaryDegrees,"b"
	case 'd', 'v', '°':
		fallthrough
	default: 
		return angle.Degrees,"°"
	}
}

func (a angle) WriteCourse(w io.Writer){
	switch a >> (bits - 5) {
	case 0, 31:
		w.Write([]byte("N"))
	case 1, 2:
		w.Write([]byte("NNE"))
	case 3, 4:
		w.Write([]byte("NE"))
	case 5, 6:
		w.Write([]byte("ENE"))
	case 7, 8:
		w.Write([]byte("E"))
	case 9, 10:
		w.Write([]byte("NSE"))
	case 11, 12:
		w.Write([]byte("SE"))
	case 13, 14:
		w.Write([]byte("SSE"))
	case 15, 16:
		w.Write([]byte("S"))
	case 17, 18:
		w.Write([]byte("SSW"))
	case 19, 20:
		w.Write([]byte("SW"))
	case 21, 22:
		w.Write([]byte("WSW"))
	case 23, 24:
		w.Write([]byte("W"))
	case 25, 26:
		w.Write([]byte("WNW"))
	case 27, 28:
		w.Write([]byte("NW"))
	case 29, 30:
		w.Write([]byte("NNW"))
	}
}

func (a angle) Format(f fmt.State, r rune) {
	switch r {
	case 'l':
		fmt.Fprintf(f, `%+.0d%+.0m`, a,a%Degree)
		Angle{a%Minute}.Format(f,'s')		
	case 'c':
 		a.WriteCourse(f)
	default:
//		f.Write([]byte(string('|')))
		Angle{a}.Format(f,r)
//		f.Write([]byte(string('|')))
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

