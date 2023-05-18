package angle

import "fmt"

func ExampleSweepContains() {
	fmt.Println(Sector{RightAngle, 3 * RightAngle, CW}.Contains(RightAngle))
	fmt.Println(Sector{3 * RightAngle, 2 * RightAngle, CW}.Contains(0))
	fmt.Println(Sector{3 * RightAngle, 2 * RightAngle, CW}.Contains(RightAngle))
	fmt.Println(Sector{0, 2 * RightAngle, CW}.Contains(3 * RightAngle))
	fmt.Println(Sector{2 * RightAngle, 0, CW}.Contains(RightAngle))
	fmt.Println(Sector{3 * RightAngle, RightAngle, CW}.Contains(0))

	fmt.Println(Sector{RightAngle, 3 * RightAngle, CCW}.Contains(2 * RightAngle))
	fmt.Println(Sector{3 * RightAngle, 2 * RightAngle, CCW}.Contains(0))
	fmt.Println(Sector{3 * RightAngle, 2 * RightAngle, CCW}.Contains(RightAngle))
	fmt.Println(Sector{0, 2 * RightAngle, CCW}.Contains(3 * RightAngle))
	fmt.Println(Sector{2 * RightAngle, 0, CCW}.Contains(RightAngle))
	fmt.Println(Sector{3 * RightAngle, RightAngle, CCW}.Contains(0))
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
	s := Sector{RightAngle, RightAngle, CW}
	fmt.Printf("%+.4t %+.4t %+.4t\n", s.Intermediate(90, 0), s.Intermediate(90, 45), s.Intermediate(90, 90))
	s = Sector{3 * RightAngle, 3 * RightAngle, CW}
	fmt.Printf("%+.4t %+.4t %+.4t\n", s.Intermediate(90, 0), s.Intermediate(90, 45), s.Intermediate(90, 90))
	s = Sector{10 * Degree, 340 * Degree, CW}
	fmt.Printf("%+.4d %+.4d %+.4d\n", s.Intermediate(90, 0), s.Intermediate(90, 45), s.Intermediate(90, 90))
	s = Sector{350 * Degree, 20 * Degree, CW}
	fmt.Printf("%+.4d %+.4d %+.4d\n", s.Intermediate(90, 0), s.Intermediate(90, 45), s.Intermediate(90, 90))
	// Output:
	// 0.2500⟳ 0.3750⟳ 0.5000⟳
	// 0.7500⟳ 0.1250⟳ 0.5000⟳
	// 10.0000° 180.0000° 350.0000°
	// 350.0000° 360.0000° 10.0000°
}

func ExampleSweepCCWDivided() {
	s := Sector{3 * RightAngle, 3*RightAngle, CCW}
	fmt.Printf("%+.4t %+.4t %+.4t\n", s.Intermediate(90, 0), s.Intermediate(90, 45), s.Intermediate(90, 90))
	s = Sector{2 * RightAngle, RightAngle, CCW}
	fmt.Printf("%+.4t %+.4t %+.4t\n", s.Intermediate(90, 0), s.Intermediate(90, 45), s.Intermediate(90, 90))
	s = Sector{10 * Degree, 340 * Degree, CCW}
	fmt.Printf("%+.4d %+.4d %+.4d\n", s.Intermediate(90, 0), s.Intermediate(90, 45), s.Intermediate(90, 90))
	s = Sector{350 * Degree, 20 * Degree, CCW}
	fmt.Printf("%+.4d %+.4d %+.4d\n", s.Intermediate(90, 0), s.Intermediate(90, 45), s.Intermediate(90, 90))
	// Output:
	// 0.7500⟳ 0.6250⟳ 0.5000⟳
	// 0.5000⟳ 0.1250⟳ 0.7500⟳
	// 10.0000° 360.0000° 350.0000°
	// 350.0000° 180.0000° 10.0000°
}
