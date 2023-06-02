package angle

import "fmt"

func ExampleDirection() {
	f := "%c (%+.2[1]d)\n"
	fmt.Printf(f, Direction(0))
	fmt.Printf(f, Direction(RightAngle))
	fmt.Printf(f, Direction(RightAngle/2))
	fmt.Printf(f, Direction(Degree*348))
	fmt.Printf(f, Direction(Degree*349))
	// Output:
	// N (0.00°)
	// E (90.00°)
	// NE (45.00°)
	// NNW (348.00°)
	// N (349.00°)
}
