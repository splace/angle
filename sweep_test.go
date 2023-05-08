package angle

import "fmt"

func ExampleSweepContains() {
	fmt.Println(Sector{RightAngle, To{3 * RightAngle, CW}}.Contains(RightAngle))
	fmt.Println(Sector{3 * RightAngle, To{2 * RightAngle, CW}}.Contains(0))
	fmt.Println(Sector{3 * RightAngle, To{2 * RightAngle, CW}}.Contains(RightAngle))
	fmt.Println(Sector{0, To{2 * RightAngle, CW}}.Contains(3 * RightAngle))
	fmt.Println(Sector{2 * RightAngle, To{0, CW}}.Contains(RightAngle))
	fmt.Println(Sector{3 * RightAngle, To{RightAngle, CW}}.Contains(0))

	fmt.Println(Sector{RightAngle, To{3 * RightAngle, CCW}}.Contains(2 * RightAngle))
	fmt.Println(Sector{3 * RightAngle, To{2 * RightAngle, CCW}}.Contains(0))
	fmt.Println(Sector{3 * RightAngle, To{2 * RightAngle, CCW}}.Contains(RightAngle))
	fmt.Println(Sector{0, To{2 * RightAngle, CCW}}.Contains(3 * RightAngle))
	fmt.Println(Sector{2 * RightAngle, To{0, CCW}}.Contains(RightAngle))
	fmt.Println(Sector{3 * RightAngle, To{RightAngle, CCW}}.Contains(0))
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

// range clockwise 20 gradians from 390 gradians, show degree.
// Note 10gradiens == 9degrees, so 9 divisions splitting 20 gradiens should be 10 angles in 2 degree steps.
func ExampleSweepStepped() {
	s := Gradian * 10 // slightly more precise than using upward multiplier ie Gradian*390
	for a := range Over(Sector{-s, To{Gradian * 20, CW}}, 9) {
		fmt.Printf("%+.4v ", a)
	}
	// Output:
	// 351.0000° 353.0000° 355.0000° 357.0000° 359.0000° 1.0000° 3.0000° 5.0000° 7.0000° 9.0000°
}

func ExampleSweepDivided() {
	s := Sector{RightAngle, To{RightAngle, CW}}
	fmt.Printf("%+.4t %+.4t %+.4t\n", s.Intermediate(90, 0), s.Intermediate(90, 45), s.Intermediate(90, 90))
	s = Sector{3 * RightAngle, To{3 * RightAngle, CW}}
	fmt.Printf("%+.4t %+.4t %+.4t\n", s.Intermediate(90, 0), s.Intermediate(90, 45), s.Intermediate(90, 90))
	s = Sector{10 * Degree, To{340 * Degree, CW}}
	fmt.Printf("%+.4d %+.4d %+.4d\n", s.Intermediate(90, 0), s.Intermediate(90, 45), s.Intermediate(90, 90))
	s = Sector{350 * Degree, To{20 * Degree, CW}}
	fmt.Printf("%+.4d %+.4d %+.4d\n", s.Intermediate(90, 0), s.Intermediate(90, 45), s.Intermediate(90, 90))
	// Output:
	// 0.2500⟳ 0.3750⟳ 0.5000⟳
	// 0.7500⟳ 0.1250⟳ 0.5000⟳
	// 10.0000° 180.0000° 350.0000°
	// 350.0000° 360.0000° 10.0000°
}

func ExampleSweepCCWDivided() {
	s := Sector{3 * RightAngle, To{RightAngle, CCW}}
	fmt.Printf("%+.4t %+.4t %+.4t\n", s.Intermediate(90, 0), s.Intermediate(90, 45), s.Intermediate(90, 90))
	s = Sector{2 * RightAngle, To{3 * RightAngle, CCW}}
	fmt.Printf("%+.4t %+.4t %+.4t\n", s.Intermediate(90, 0), s.Intermediate(90, 45), s.Intermediate(90, 90))
	s = Sector{10 * Degree, To{20 * Degree, CCW}}
	fmt.Printf("%+.4d %+.4d %+.4d\n", s.Intermediate(90, 0), s.Intermediate(90, 45), s.Intermediate(90, 90))
	s = Sector{350 * Degree, To{340 * Degree, CCW}}
	fmt.Printf("%+.4d %+.4d %+.4d\n", s.Intermediate(90, 0), s.Intermediate(90, 45), s.Intermediate(90, 90))
	// Output:
	// 0.7500⟳ 0.6250⟳ 0.5000⟳
	// 0.5000⟳ 0.1250⟳ 0.7500⟳
	// 10.0000° 0.0000° 350.0000°
	// 350.0000° 180.0000° 10.0000°
}
