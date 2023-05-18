package angle

import "fmt"

func ExampleSweepContains() {
	fmt.Println(NewCWSector(RightAngle, 3 * RightAngle).Contains(RightAngle))
	fmt.Println(NewCWSector(3 * RightAngle, 2 * RightAngle).Contains(0))
	fmt.Println(NewCWSector(3 * RightAngle, 2 * RightAngle).Contains(RightAngle))
	fmt.Println(NewCWSector(0, 2 * RightAngle).Contains(3 * RightAngle))
	fmt.Println(NewCWSector(2 * RightAngle, 0).Contains(RightAngle))
	fmt.Println(NewCWSector(3 * RightAngle, RightAngle).Contains(0))

	fmt.Println(NewCCWSector(RightAngle, 3 * RightAngle).Contains(2 * RightAngle))
	fmt.Println(NewCCWSector(3 * RightAngle, 2 * RightAngle).Contains(0))
	fmt.Println(NewCCWSector(3 * RightAngle, 2 * RightAngle).Contains(RightAngle))
	fmt.Println(NewCCWSector(0, 2 * RightAngle).Contains(3 * RightAngle))
	fmt.Println(NewCCWSector(2 * RightAngle, 0).Contains(RightAngle))
	fmt.Println(NewCCWSector(3 * RightAngle, RightAngle).Contains(0))
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

func ExampleSweepDivided() {
	s := NewCWSector(RightAngle, RightAngle)
	fmt.Printf("%+.4t %+.4t %+.4t\n", s.Intermediate(90, 0), s.Intermediate(90, 45), s.Intermediate(90, 90))
	s = NewCWSector(3 * RightAngle, 3 * RightAngle)
	fmt.Printf("%+.4t %+.4t %+.4t\n", s.Intermediate(90, 0), s.Intermediate(90, 45), s.Intermediate(90, 90))
	s = NewCWSector(10 * Degree, 340 * Degree)
	fmt.Printf("%+.4d %+.4d %+.4d\n", s.Intermediate(90, 0), s.Intermediate(90, 45), s.Intermediate(90, 90))
	s = NewCWSector(350 * Degree, 20 * Degree)
	fmt.Printf("%+.4d %+.4d %+.4d\n", s.Intermediate(90, 0), s.Intermediate(90, 45), s.Intermediate(90, 90))
	// Output:
	// 0.2500⟳ 0.3750⟳ 0.5000⟳
	// 0.7500⟳ 0.1250⟳ 0.5000⟳
	// 10.0000° 180.0000° 350.0000°
	// 350.0000° 360.0000° 10.0000°
}

func ExampleSweepCCWDivided() {
	s := NewCCWSector(3 * RightAngle, 3*RightAngle)
	fmt.Printf("%+.4t %+.4t %+.4t\n", s.Intermediate(90, 0), s.Intermediate(90, 45), s.Intermediate(90, 90))
	s = NewCCWSector(2 * RightAngle, RightAngle)
	fmt.Printf("%+.4t %+.4t %+.4t\n", s.Intermediate(90, 0), s.Intermediate(90, 45), s.Intermediate(90, 90))
	s = NewCCWSector(10 * Degree, 340 * Degree)
	fmt.Printf("%+.4d %+.4d %+.4d\n", s.Intermediate(90, 0), s.Intermediate(90, 45), s.Intermediate(90, 90))
	s = NewCCWSector(350 * Degree, 20 * Degree)
	fmt.Printf("%+.4d %+.4d %+.4d\n", s.Intermediate(90, 0), s.Intermediate(90, 45), s.Intermediate(90, 90))
	// Output:
	// 0.7500⟳ 0.6250⟳ 0.5000⟳
	// 0.5000⟳ 0.1250⟳ 0.7500⟳
	// 10.0000° 360.0000° 350.0000°
	// 350.0000° 180.0000° 10.0000°
}
