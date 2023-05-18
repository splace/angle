package angle_test


import "fmt"
import . "../angle"  // remove 'go.mod' for local testing.

// range clockwise 20 gradians from 390 gradians, show values in degrees.
// Notice: 10gradians == 9degrees, so 9 divisions splitting 20 gradiens should be 10 angles in 2 degree steps.
func ExampleAngle_testRangeOverSector() {
	for a := range Over(Sector{Gradian*390, Gradian * 20, CW}, 9) {
		fmt.Printf("%+.4v ", a)
	}
	// Output:
	// 351.0000° 353.0000° 355.0000° 357.0000° 359.0000° 1.0000° 3.0000° 5.0000° 7.0000° 9.0000°
}

