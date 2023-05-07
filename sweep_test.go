package angle

import "fmt"




func ExampleSweepContains() {
	fmt.Println(Sweep{RightAngle,NewClockwise(3*RightAngle)}.Contains(RightAngle))
	fmt.Println(Sweep{3*RightAngle,NewClockwise(2*RightAngle)}.Contains(0))
	fmt.Println(Sweep{3*RightAngle,NewClockwise(2*RightAngle)}.Contains(RightAngle))
	fmt.Println(Sweep{0,NewClockwise(2*RightAngle)}.Contains(3*RightAngle))
	fmt.Println(Sweep{2*RightAngle,NewClockwise(0)}.Contains(RightAngle))
	fmt.Println(Sweep{3*RightAngle,NewClockwise(RightAngle)}.Contains(0))

	fmt.Println(Sweep{RightAngle,NewCounterClockwise(3*RightAngle)}.Contains(2*RightAngle))
	fmt.Println(Sweep{3*RightAngle,NewCounterClockwise(2*RightAngle)}.Contains(0))
	fmt.Println(Sweep{3*RightAngle,NewCounterClockwise(2*RightAngle)}.Contains(RightAngle))
	fmt.Println(Sweep{0,NewCounterClockwise(2*RightAngle)}.Contains(3*RightAngle))
	fmt.Println(Sweep{2*RightAngle,NewCounterClockwise(0)}.Contains(RightAngle))
	fmt.Println(Sweep{3*RightAngle,NewCounterClockwise(RightAngle)}.Contains(0))
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
	s:=Gradian*10 // slightly more precise than using bit multiplier ie Gradian*390
	for a:= range Over(Sweep{-s,NewClockwise(Gradian*20)},9){
		fmt.Printf("%+.4v ",a)
	}
	// Output:
	// 351.0000° 353.0000° 355.0000° 357.0000° 359.0000° 1.0000° 3.0000° 5.0000° 7.0000° 9.0000°
}

func ExampleSweepDivided() {
	s:=Sweep{RightAngle,NewClockwise(RightAngle)}
	fmt.Printf("%+.4t %+.4t %+.4t\n",s.Intermediate(90,0),s.Intermediate(90,45),s.Intermediate(90,90))
	s=Sweep{3*RightAngle,NewClockwise(3*RightAngle)}
	fmt.Printf("%+.4t %+.4t %+.4t\n",s.Intermediate(90,0),s.Intermediate(90,45),s.Intermediate(90,90))
	s=Sweep{10*Degree,NewClockwise(340*Degree)}
	fmt.Printf("%+.4d %+.4d %+.4d\n",s.Intermediate(90,0),s.Intermediate(90,45),s.Intermediate(90,90))
	s=Sweep{350*Degree,NewClockwise(20*Degree)}
	fmt.Printf("%+.4d %+.4d %+.4d\n",s.Intermediate(90,0),s.Intermediate(90,45),s.Intermediate(90,90))
	// Output:
	// 0.2500⟳ 0.3750⟳ 0.5000⟳
	// 0.7500⟳ 0.1250⟳ 0.5000⟳
	// 10.0000° 180.0000° 350.0000°
	// 350.0000° 360.0000° 10.0000°
}

func ExampleSweepCCWDivided() {
	s:=Sweep{3*RightAngle,NewCounterClockwise(RightAngle)}
	fmt.Printf("%+.4t %+.4t %+.4t\n",s.Intermediate(90,0),s.Intermediate(90,45),s.Intermediate(90,90))
	s=Sweep{2*RightAngle,NewCounterClockwise(3*RightAngle)}
	fmt.Printf("%+.4t %+.4t %+.4t\n",s.Intermediate(90,0),s.Intermediate(90,45),s.Intermediate(90,90))
	s=Sweep{10*Degree,NewCounterClockwise(20*Degree)}
	fmt.Printf("%+.4d %+.4d %+.4d\n",s.Intermediate(90,0),s.Intermediate(90,45),s.Intermediate(90,90))
	s=Sweep{350*Degree,NewCounterClockwise(340*Degree)}
	fmt.Printf("%+.4d %+.4d %+.4d\n",s.Intermediate(90,0),s.Intermediate(90,45),s.Intermediate(90,90))
	// Output:
	// 0.7500⟳ 0.6250⟳ 0.5000⟳
	// 0.5000⟳ 0.1250⟳ 0.7500⟳
	// 10.0000° 0.0000° 350.0000°
	// 350.0000° 180.0000° 10.0000°
}

