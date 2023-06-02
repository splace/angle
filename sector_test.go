package angle

import "fmt"

func ExampleSectorContains() {
	fmt.Println(Sector{Direction(RightAngle), 3 * RightAngle}.Contains(Direction(RightAngle)))
	fmt.Println(Sector{Direction(3 * RightAngle), 2 * RightAngle}.Contains(0))
	fmt.Println(Sector{Direction(3 * RightAngle), 2 * RightAngle}.Contains(Direction(RightAngle)))
	fmt.Println(Sector{Direction(0), 2 * RightAngle}.Contains(Direction(3 * RightAngle)))
	fmt.Println(Sector{Direction(2 * RightAngle), 0}.Contains(Direction(RightAngle)))
	fmt.Println(Sector{Direction(3 * RightAngle), RightAngle}.Contains(0))

	fmt.Println(Sector{0, 3 * RightAngle}.Contains(Direction(2 * RightAngle)))
	fmt.Println(Sector{Direction(RightAngle), 2 * RightAngle}.Contains(0))
	fmt.Println(Sector{Direction(RightAngle), 2 * RightAngle}.Contains(Direction(2 * RightAngle)))
	fmt.Println(Sector{Direction(2 * RightAngle), 0}.Contains(Direction(3 * RightAngle)))
	fmt.Println(Sector{Direction(2 * RightAngle), 2 * RightAngle}.Contains(Direction(3 * RightAngle)))
	fmt.Println(Sector{Direction(2 * RightAngle), RightAngle}.Contains(0))
	// Output:
	// true
	// true
	// true
	// false
	// false
	// true
	// true
	// false
	// true
	// false
	// true
	// false
}

// Example (see Sector): doubling the Direction makes no sense in the problem-space, but doubling the Angle/Phase clearly represents twice the sector size.
func ExampleSector() {
	s := Over(Sector{Direction(RightAngle), RightAngle}, 2)
	fmt.Printf("%+.4t %+.4t %+.4t\n", <-s, <-s, <-s)
	s = Over(Sector{Direction(3 * RightAngle), 3 * RightAngle}, 2)
	fmt.Printf("%+.4t %+.4t %+.4t\n", <-s, <-s, <-s)
	s = Over(Sector{Direction(10 * Degree), 340 * Degree}, 2)
	fmt.Printf("%+.4d %+.4d %+.4d\n", <-s, <-s, <-s)
	s = Over(Sector{Direction(350 * Degree), 20 * Degree}, 2)
	fmt.Printf("%+.4d %+.4d %+.4d\n", <-s, <-s, <-s)

	s = Over(Sector{0, 3 * RightAngle}, 2)
	fmt.Printf("%+.4t %+.4t %+.4t\n", <-s, <-s, <-s)
	s = Over(Sector{Direction(2 * RightAngle), 3 * RightAngle}, 2)
	fmt.Printf("%+.4t %+.4t %+.4t\n", <-s, <-s, <-s)
	s = Over(Sector{Direction(10 * Degree), 340 * Degree}, 2)
	fmt.Printf("%+.4d %+.4d %+.4d\n", <-s, <-s, <-s)
	s = Over(Sector{Direction(350 * Degree), 20 * Degree}, 2)
	fmt.Printf("%+.4d %+.4d %+.4d\n", <-s, <-s, <-s)
	// Output:
	// 0.2500⟳ 0.3750⟳ 0.5000⟳
	// 0.7500⟳ 0.1250⟳ 0.5000⟳
	// 10.0000° 180.0000° 350.0000°
	// 350.0000° 360.0000° 10.0000°
	// 0.0000⟳ 0.3750⟳ 0.7500⟳
	// 0.5000⟳ 0.8750⟳ 0.2500⟳
	// 10.0000° 180.0000° 350.0000°
	// 350.0000° 360.0000° 10.0000°
}
