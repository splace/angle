package angle_test

import "fmt"

//import . "github.com/splace/angle"
import . "../angle" // remove 'go.mod' for local testing.

func ExampleDirection_testSectorContains() {
	fmt.Println(
		Sector{Direction(RightAngle), 3 * RightAngle}.Contains(Direction(RightAngle)),
		Sector{Direction(3 * RightAngle), 2 * RightAngle}.Contains(Direction(0)),
		Sector{Direction(3 * RightAngle), 2 * RightAngle}.Contains(Direction(RightAngle)),
		Sector{Direction(0), 2 * RightAngle}.Contains(Direction(3*RightAngle)),
		Sector{Direction(2 * RightAngle), 0}.Contains(Direction(RightAngle)),
		Sector{Direction(3 * RightAngle), RightAngle}.Contains(Direction(0)),
	)
	fmt.Println(
		Sector{0, 3 * RightAngle}.Contains(Direction(2*RightAngle)),
		Sector{Direction(RightAngle), 2 * RightAngle}.Contains(0),
		Sector{Direction(3 * RightAngle), 2 * RightAngle}.Contains(Direction(2*RightAngle)),
		Sector{Direction(2 * RightAngle), 2 * RightAngle}.Contains(Direction(3*RightAngle)),
		Sector{0, 2 * RightAngle}.Contains(Direction(RightAngle)),
		Sector{Direction(RightAngle), RightAngle}.Contains(0),
	)
	// Output:
	// true true false false false false
	// true false false true true false
}

// range clockwise 20 gradians from 390 gradians in 9 steps, show values in degrees.
// Notice: 10gradians == 9degrees, so 9 divisions splitting 20 gradiens should be 10 angles in 2 degree steps.
func ExampleDirection_testRangeOverSector() {
	for a := range Over(NewSectorCW(Direction(Gradian*390), Direction(Gradian*10)), 9) {
		fmt.Printf("%+.3v ", a)
	}
	fmt.Println()
	// Output:
	// 351.000° 353.000° 355.000° 357.000° 359.000° 1.000° 3.000° 5.000° 7.000° 9.000°
}

func ExampleDirection_testRangeOverSectorBothWays() {
	s := NewSectorCCW(Direction(Gradian*10), Direction(Gradian*390))
	for a := range OverCCW(s, 9) {
		fmt.Printf("%+.3v,", a)
	}
	fmt.Println()
	// reverse
	for a := range OverCW(s, 9) {
		fmt.Printf("%+.3v,", a)
	}
	fmt.Println()
	// Output:
	// 9.000°,7.000°,5.000°,3.000°,1.000°,359.000°,357.000°,355.000°,353.000°,351.000°,
	// 351.000°,353.000°,355.000°,357.000°,359.000°,1.000°,3.000°,5.000°,7.000°,9.000°,
}
