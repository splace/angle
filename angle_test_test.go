package angle_test


import "fmt"
import . "../angle"  // remove 'go.mod' for local testing.

// range clockwise 20 gradians from 390 gradians, show values in degrees.
// Notice: 10gradians == 9degrees, so 9 divisions splitting 20 gradiens should be 10 angles in 2 degree steps.
func ExampleAngle_testRangeOverSector() {
	for a := range Over(NewCWSector(Gradian*390, Gradian * 20), 9) {
		fmt.Printf("%+.3v ", a)
	}
	fmt.Println()
	// Output:
	// 351.000° 353.000° 355.000° 357.000° 359.000° 1.000° 3.000° 5.000° 7.000° 9.000° 

}

func ExampleAngle_testRangeOverSectorVarious() {
	s:=NewCCWSector(Gradian*10, Gradian * 20)
	for a := range Over(s, 9) {
		fmt.Printf("%+.3v,", a)
	}
	fmt.Println()
	// reverse
	for a := range ReverseOver(s, 9) {
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
	for a := range ReverseOver(s, 9) {
		fmt.Printf("%+.3v,", a)
	}
	// Output:
	// 9.000°,7.000°,5.000°,3.000°,1.000°,359.000°,357.000°,355.000°,353.000°,351.000°,
	// 351.000°,353.000°,355.000°,357.000°,359.000°,1.000°,3.000°,5.000°,7.000°,9.000°,
	// 9.000°,47.000°,85.000°,123.000°,161.000°,199.000°,237.000°,275.000°,313.000°,351.000°,
	// 351.000°,313.000°,275.000°,237.000°,199.000°,161.000°,123.000°,85.000°,47.000°,9.000°,
}


