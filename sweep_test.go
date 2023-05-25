package angle

import "fmt"

func ExampleSweepAngle() {
	fmt.Printf("%.1v degrees == %+[1]r == %+[1]v == %+.1[1]l == %+.2[1]f\n", Angle(Radian))
	fmt.Printf("%.1v degrees == %+[1]㎭ == %+.1[1]l == %+.1[1]g == %+.2[1]f\n", Angle(RightAngle))
	// Output:
	// 57.3 degrees == 1㎭ == 57.295784° == 57.3° == 15.92%
	// 90.0 degrees == 1.5707964㎭ == 90.0° == 100.0ᵍ == 25.00%

}

func ExampleSweepContains() {
	fmt.Println(Sector{RightAngle, Delta(3 * RightAngle),CW}.Contains(RightAngle))
	fmt.Println(Sector{3 * RightAngle, Delta(2 * RightAngle),CW}.Contains(0))
	fmt.Println(Sector{3 * RightAngle, Delta(2 * RightAngle),CW}.Contains(RightAngle))
	fmt.Println(Sector{0, Delta(2 * RightAngle),CW}.Contains(3 * RightAngle))
	fmt.Println(Sector{2 * RightAngle, Delta(0),CW}.Contains(RightAngle))
	fmt.Println(Sector{3 * RightAngle, Delta(RightAngle),CW}.Contains(0))

	fmt.Println(Sector{RightAngle, Delta(3 * RightAngle),CCW}.Contains(2 * RightAngle))
	fmt.Println(Sector{3 * RightAngle, Delta(2 * RightAngle),CCW}.Contains(0))
	fmt.Println(Sector{3 * RightAngle, Delta(2 * RightAngle),CCW}.Contains(2*RightAngle))
	fmt.Println(Sector{0, Delta(2 * RightAngle),CCW}.Contains(3 * RightAngle))
	fmt.Println(Sector{2 * RightAngle, Delta(0),CCW}.Contains(RightAngle))
	fmt.Println(Sector{3 * RightAngle, Delta(RightAngle),CCW}.Contains(0))
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
	s := Over(Sector{RightAngle, Delta(RightAngle),CW},2)
	fmt.Printf("%+.4t %+.4t %+.4t\n", <- s,<-s,<-s )
	s = Over(Sector{3*RightAngle, Delta(3*RightAngle),CW},2)
	fmt.Printf("%+.4t %+.4t %+.4t\n", <- s,<-s,<-s )
	s = Over(Sector{10 * Degree, Delta(340 * Degree),CW},2)
	fmt.Printf("%+.4d %+.4d %+.4d\n", <- s,<-s,<-s )
	s = Over(Sector{350 * Degree, Delta(20 * Degree),CW},2)
	fmt.Printf("%+.4d %+.4d %+.4d\n", <- s,<-s,<-s )


	s = Over(Sector{3*RightAngle, Delta(3*RightAngle),CCW},2)
	fmt.Printf("%+.4t %+.4t %+.4t\n", <- s,<-s,<-s )
	s = Over(Sector{2*RightAngle, Delta(RightAngle),CCW},2)
	fmt.Printf("%+.4t %+.4t %+.4t\n", <- s,<-s,<-s )
	s = Over(Sector{10 * Degree, Delta(340 * Degree),CCW},2)
	fmt.Printf("%+.4d %+.4d %+.4d\n", <- s,<-s,<-s )
	s = Over(Sector{350 * Degree, Delta(20 * Degree),CCW},2)
	fmt.Printf("%+.4d %+.4d %+.4d\n", <- s,<-s,<-s )
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

