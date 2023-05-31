package angle

import "fmt"

func ExampleDirection() {
	fmt.Printf("%c %+.1[1]d\n", Direction(0))
	fmt.Printf("%c %+.1[1]d\n", Direction(RightAngle))
	fmt.Printf("%c %+.1[1]d\n", Direction(RightAngle/2))
	fmt.Printf("%c %+.1[1]d\n", Direction(Degree*348))
	fmt.Printf("%c %+.1[1]d\n", Direction(Degree*349))
	// Output:
	// N 0.0°
	// E 90.0°
	// NE 45.0°
	// NNW 348.0°
	// N 349.0°
}
