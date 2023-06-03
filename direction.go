package angle

import "fmt"
import "io"

type Direction Phase

func (a Direction) WriteCourse(w io.Writer) {
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

func (d Direction) Format(f fmt.State, r rune) {
	switch r {
	case 'c':
		d.WriteCourse(f)
	default:
		//		f.Write([]byte(string('|')))
		Phase(d).Format(f, r)
		//		f.Write([]byte(string('|')))
	}
}
