package angle_test


import "fmt"
import . "../angle"  // remove 'go.mod' for local testing.

// range clockwise 20 gradians from 390 gradians, show values in degrees.
// Notice: 10gradians == 9degrees, so 9 divisions splitting 20 gradiens should be 10 angles in 2 degree steps.
func ExampleAngle_testRangeOverSector() {
	for a := range Over(Sector{Gradian*390, Gradian * 20, CW}, 9) {
		fmt.Printf("%+.3v ", a)
	}
	fmt.Println()
	for a := range Over(Sector{Gradian*10, Gradian * 380, CCW}, 9) {
		fmt.Printf("%+.3v ", a)
	}
	fmt.Println()
	// Output:
	// 351.000° 353.000° 355.000° 357.000° 359.000° 1.000° 3.000° 5.000° 7.000° 9.000° 
	// 9.000° 7.000° 5.000° 3.000° 1.000° 359.000° 357.000° 355.000° 353.000° 351.000° 
}

