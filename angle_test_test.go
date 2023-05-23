package angle_test


import "fmt"
//import . "github.com/splace/angle"
import . "../angle"  // remove 'go.mod' for local testing.

func ExampleAngle_testDelta() {
	fmt.Printf("%.1v degrees == %+[1]r == %+[1]v == %+.1[1]l == %+.2[1]f\n", Delta{Radian})
	fmt.Printf("%.1v degrees == %+[1]㎭ == %+.1[1]l == %+.1[1]g == %+.2[1]f\n", Delta{RightAngle})
	// Output:
	// 57.3 degrees == 1㎭ == 57.295784° == 57.3° == 15.92%
	// 90.0 degrees == 1.5707964㎭ == 90.0° == 100.0ᵍ == 25.00%
}

func ExampleAngle_testSweepContains() {
	fmt.Println(
		Sector{RightAngle, Delta{3 * RightAngle},CW}.Contains(RightAngle),
		Sector{3 * RightAngle, Delta{2 * RightAngle},CW}.Contains(0),
		Sector{3 * RightAngle, Delta{2 * RightAngle},CW}.Contains(RightAngle),
		Sector{0, Delta{2 * RightAngle},CW}.Contains(3 * RightAngle),
		Sector{2 * RightAngle, Delta{0},CW}.Contains(RightAngle),
		Sector{3 * RightAngle, Delta{RightAngle},CW}.Contains(0),
	)
	fmt.Println(
		Sector{RightAngle, Delta{RightAngle},CCW}.Contains(2 * RightAngle),
		Sector{3 * RightAngle, Delta{2 * RightAngle},CCW}.Contains(0),
		Sector{3 * RightAngle, Delta{2 * RightAngle},CCW}.Contains(2*RightAngle),
		Sector{0, Delta{2 * RightAngle},CCW}.Contains(3 * RightAngle),
		Sector{2 * RightAngle, Delta{0},CCW}.Contains(RightAngle),
		Sector{3 * RightAngle, Delta{3*RightAngle},CCW}.Contains(0),
	)
	// Output:
	// true true true false false true
	// true false false true true false
}

// range clockwise 20 gradians from 390 gradians in 9 steps, show values in degrees.
// Notice: 10gradians == 9degrees, so 9 divisions splitting 20 gradiens should be 10 angles in 2 degree steps.
func ExampleAngle_testRangeOverSector() {
	for a := range Over(NewSector(Gradian*390, Gradian * 20,CW), 9) {
		fmt.Printf("%+.3v ", a)
	}
	fmt.Println()
	// Output:
	// 351.000° 353.000° 355.000° 357.000° 359.000° 1.000° 3.000° 5.000° 7.000° 9.000° 
}

func ExampleAngle_testRangeOverSectorVarious2() {
	s:=NewSector(Gradian*10, Gradian * 20,CCW)
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
	s.Direction=!s.Direction
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



