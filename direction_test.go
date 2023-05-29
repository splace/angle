package angle

import "fmt"

func ExampleDirectionCourse() {
	fmt.Printf("%c %+[1]d\n", Direction(0))
	fmt.Printf("%c %+[1]d\n", Direction(Degree*180))
	fmt.Printf("%c %+[1]d\n", Direction(Degree*348))
	fmt.Printf("%c %+[1]d\n", Direction(Degree*349))
	// Output:
	// N 0°
	// S 180°
	// NNW 348°
	// N 349°
}
