package angle_test

import "fmt"

//import . "github.com/splace/angle"
import . "../angle" // remove 'go.mod' for local testing.

func ExampleDirection_testSweepContains() {
	fmt.Println(
		Sector{Direction(RightAngle),3 * RightAngle, CW}.Contains(Direction(RightAngle)),
		Sector{Direction(3 * RightAngle),2 * RightAngle, CW}.Contains(Direction(0)),
		Sector{Direction(3 * RightAngle),2 * RightAngle, CW}.Contains(Direction(RightAngle)),
		Sector{Direction(0),2 * RightAngle, CW}.Contains(Direction(3*RightAngle)),
		Sector{Direction(2 * RightAngle),0, CW}.Contains(Direction(RightAngle)),
		Sector{Direction(3 * RightAngle),RightAngle, CW}.Contains(Direction(0)),
	)
	fmt.Println(
		Sector{Direction(RightAngle),RightAngle, CCW}.Contains(Direction(2*RightAngle)),
		Sector{Direction(3 * RightAngle),2 * RightAngle, CCW}.Contains(Direction(0)),
		Sector{Direction(3 * RightAngle),2 * RightAngle, CCW}.Contains(Direction(2*RightAngle)),
		Sector{Direction(0),2 * RightAngle, CCW}.Contains(Direction(3*RightAngle)),
		Sector{Direction(2 * RightAngle),0, CCW}.Contains(Direction(RightAngle)),
		Sector{Direction(3 * RightAngle),3 * RightAngle, CCW}.Contains(Direction(0)),
	)
	// Output:
	// true true true false false true
	// true false false true true false
}

// range clockwise 20 gradians from 390 gradians in 9 steps, show values in degrees.
// Notice: 10gradians == 9degrees, so 9 divisions splitting 20 gradiens should be 10 angles in 2 degree steps.
func ExampleDirection_testRangeOverSector() {
	for a := range Over(NewSector(Direction(Gradian*390), Gradian*20, CW), 9) {
		fmt.Printf("%+.3v ", a)
	}
	fmt.Println()
	// Output:
	// 351.000° 353.000° 355.000° 357.000° 359.000° 1.000° 3.000° 5.000° 7.000° 9.000°
}

func ExampleDirection_testRangeOverSectorVarious2() {
	s := NewSector(Direction(Gradian*10), Gradian*20, CCW)
	for a := range Over(s, 9) {
		fmt.Printf("%+.3v,", a)
	}
	fmt.Println()
	// reverse
	for a := range CCWOver(s, 9) {
		fmt.Printf("%+.3v,", a)
	}
	fmt.Println()
	// the other half of a rotation
	s.Turn = !s.Turn
	for a := range Over(s, 9) {
		fmt.Printf("%+.3v,", a)
	}
	fmt.Println()
	// the other half of a rotation in reverse
	for a := range CCWOver(s, 9) {
		fmt.Printf("%+.3v,", a)
	}
	// Output:
	// 9.000°,7.000°,5.000°,3.000°,1.000°,359.000°,357.000°,355.000°,353.000°,351.000°,
	// 351.000°,353.000°,355.000°,357.000°,359.000°,1.000°,3.000°,5.000°,7.000°,9.000°,
	// 9.000°,47.000°,85.000°,123.000°,161.000°,199.000°,237.000°,275.000°,313.000°,351.000°,
	// 351.000°,313.000°,275.000°,237.000°,199.000°,161.000°,123.000°,85.000°,47.000°,9.000°,
}
