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

func ExampleSectorOver() {
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

func ExampleSectorReverse() {
	s:=NewSector(Direction(RightAngle*3),Direction(RightAngle),CW)
	rs:=s.Opposite()  
	for d:=range Over(Sector{Angle:340*Degree},17){
		fmt.Printf("%+v %v %v\n", d, s.Contains(d),rs.Contains(d))
	}
	// Output:
	//0° true false
	//20° true false
	//40° true false
	//60° true false
	//80° true false
	//100° false true
	//120° false true
	//140° false true
	//160° false true
	//180° false true
	//200° false true
	//220° false true
	//240° false true
	//260° false true
	//280° true false
	//300° true false
	//320° true false
	//340° true false
}
