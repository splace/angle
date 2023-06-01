package angle

import "fmt"

func ExampleDirection() {
	f:="%c (%+.1[1]d)\n"
	fmt.Printf(f, Direction(0))
	fmt.Printf(f, Direction(RightAngle))
	fmt.Printf(f, Direction(RightAngle/2))
	fmt.Printf(f, Direction(Degree*348))
	fmt.Printf(f, Direction(Degree*349))
	// Output:
	// N (0.0°)
	// E (90.0°)
	// NE (45.0°)
	// NNW (348.0°)
	// N (349.0°)
}
