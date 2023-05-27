package angle

import "fmt"

func ExampleSweepAngle() {
	fmt.Printf("%+.0[1]l %[1]c\n",Direction(RightAngle))	
	fmt.Printf("%+.0[1]l %[1]c\n",Direction(RightAngle*3/2))
	// Output:
	// 90°0′0″ E
	// 135°0′0″ SE
}

func ExampleSweepContains() {
	fmt.Println(Sector{Direction(RightAngle),3 * RightAngle, CW}.Contains(Direction(RightAngle)))
	fmt.Println(Sector{Direction(3 * RightAngle),2 * RightAngle, CW}.Contains(Direction(0)))
	fmt.Println(Sector{Direction(3 * RightAngle),2 * RightAngle, CW}.Contains(Direction(RightAngle)))
	fmt.Println(Sector{Direction(0),2 * RightAngle, CW}.Contains(Direction(3 * RightAngle)))
	fmt.Println(Sector{Direction(2 * RightAngle),0, CW}.Contains(Direction(RightAngle)))
	fmt.Println(Sector{Direction(3 * RightAngle),RightAngle, CW}.Contains(Direction(0)))

	fmt.Println(Sector{Direction(RightAngle),3 * RightAngle, CCW}.Contains(Direction(2 * RightAngle)))
	fmt.Println(Sector{Direction(3 * RightAngle),2 * RightAngle, CCW}.Contains(Direction(0)))
	fmt.Println(Sector{Direction(3 * RightAngle),2 * RightAngle, CCW}.Contains(Direction(2 * RightAngle)))
	fmt.Println(Sector{Direction(0),2 * RightAngle, CCW}.Contains(Direction(3 * RightAngle)))
	fmt.Println(Sector{Direction(2 * RightAngle),0, CCW}.Contains(Direction(RightAngle)))
	fmt.Println(Sector{Direction(3 * RightAngle),RightAngle, CCW}.Contains(Direction(0)))
	// Output:
	// true
	// true
	// true
	// false
	// false
	// true
	// false
	// false
	// false
	// true
	// true
	// false
}

func ExampleSweepSector() {
	s := Over(Sector{Direction(RightAngle),RightAngle, CW}, 2)
	fmt.Printf("%+.4t %+.4t %+.4t\n", <-s, <-s, <-s)
	s = Over(Sector{Direction(3 * RightAngle),3 * RightAngle, CW}, 2)
	fmt.Printf("%+.4t %+.4t %+.4t\n", <-s, <-s, <-s)
	s = Over(Sector{Direction(10 * Degree),340 * Degree, CW}, 2)
	fmt.Printf("%+.4d %+.4d %+.4d\n", <-s, <-s, <-s)
	s = Over(Sector{Direction(350 * Degree),20 * Degree, CW}, 2)
	fmt.Printf("%+.4d %+.4d %+.4d\n", <-s, <-s, <-s)

	s = Over(Sector{Direction(3 * RightAngle),3 * RightAngle, CCW}, 2)
	fmt.Printf("%+.4t %+.4t %+.4t\n", <-s, <-s, <-s)
	s = Over(Sector{Direction(2 * RightAngle),RightAngle, CCW}, 2)
	fmt.Printf("%+.4t %+.4t %+.4t\n", <-s, <-s, <-s)
	s = Over(Sector{Direction(10 * Degree),340 * Degree, CCW}, 2)
	fmt.Printf("%+.4d %+.4d %+.4d\n", <-s, <-s, <-s)
	s = Over(Sector{Direction(350 * Degree),20 * Degree, CCW}, 2)
	fmt.Printf("%+.4d %+.4d %+.4d\n", <-s, <-s, <-s)
	// Output:
	// 0.2500⟳ 0.3750⟳ 0.5000⟳
	// 0.7500⟳ 0.1250⟳ 0.5000⟳
	// 10.0000° 180.0000° 350.0000°
	// 350.0000° 360.0000° 10.0000°
	// 0.7500⟳ 0.6250⟳ 0.5000⟳
	// 0.5000⟳ 0.1250⟳ 0.7500⟳
	// 10.0000° 360.0000° 350.0000°
	// 350.0000° 180.0000° 10.0000°
}


func ExampleSweepCourse() {
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

