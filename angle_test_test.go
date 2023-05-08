package angle_test

import "fmt"
import . "../angle"  // for this to work need to remove 'go.mod' ? this mod/sum/v2/install has gone way too complex.

// range clockwise 20 gradians from 390 gradians, show degree.
// Notice: 10gradians == 9degrees, so 9 divisions splitting 20 gradiens should be 10 angles in 2 degree steps.
func ExampleAngleTestRangeOverSector() {
	for a := range Over(Sector{Gradian*390, To{Gradian * 20, CW}}, 9) {
		fmt.Printf("%+.4v ", a)
	}
	// Output:
	// 351.0000° 353.0000° 355.0000° 357.0000° 359.0000° 1.0000° 3.0000° 5.0000° 7.0000° 9.0000°
}

