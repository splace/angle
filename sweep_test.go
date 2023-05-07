package angle

import "fmt"


func ExampleSweepContains() {
	fmt.Println(Sweep{RightAngle,NewCWTo(3*RightAngle)}.Contains(RightAngle))
	fmt.Println(Sweep{3*RightAngle,NewCWTo(2*RightAngle)}.Contains(0))
	fmt.Println(Sweep{3*RightAngle,NewCWTo(2*RightAngle)}.Contains(RightAngle))
	fmt.Println(Sweep{0,NewCWTo(2*RightAngle)}.Contains(3*RightAngle))
	fmt.Println(Sweep{2*RightAngle,NewCWTo(0)}.Contains(RightAngle))
	fmt.Println(Sweep{3*RightAngle,NewCWTo(RightAngle)}.Contains(0))

	fmt.Println(Sweep{RightAngle,NewCCWTo(3*RightAngle)}.Contains(2*RightAngle))
	fmt.Println(Sweep{3*RightAngle,NewCCWTo(2*RightAngle)}.Contains(0))
	fmt.Println(Sweep{3*RightAngle,NewCCWTo(2*RightAngle)}.Contains(RightAngle))
	fmt.Println(Sweep{0,NewCCWTo(2*RightAngle)}.Contains(3*RightAngle))
	fmt.Println(Sweep{2*RightAngle,NewCCWTo(0)}.Contains(RightAngle))
	fmt.Println(Sweep{3*RightAngle,NewCCWTo(RightAngle)}.Contains(0))
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
	s:=Gradian*10 // slightly more precise than using upward multiplier ie Gradian*390
	for a:= range Over(Sweep{-s,NewCWTo(Gradian*20)},9){
		fmt.Printf("%+.4v ",a)
	}
	// Output:
	// 351.0000° 353.0000° 355.0000° 357.0000° 359.0000° 1.0000° 3.0000° 5.0000° 7.0000° 9.0000°
}

func ExampleSweepDivided() {
	s:=Sweep{RightAngle,NewCWTo(RightAngle)}
	fmt.Printf("%+.4t %+.4t %+.4t\n",s.Intermediate(90,0),s.Intermediate(90,45),s.Intermediate(90,90))
	s=Sweep{3*RightAngle,NewCWTo(3*RightAngle)}
	fmt.Printf("%+.4t %+.4t %+.4t\n",s.Intermediate(90,0),s.Intermediate(90,45),s.Intermediate(90,90))
	s=Sweep{10*Degree,NewCWTo(340*Degree)}
	fmt.Printf("%+.4d %+.4d %+.4d\n",s.Intermediate(90,0),s.Intermediate(90,45),s.Intermediate(90,90))
	s=Sweep{350*Degree,NewCWTo(20*Degree)}
	fmt.Printf("%+.4d %+.4d %+.4d\n",s.Intermediate(90,0),s.Intermediate(90,45),s.Intermediate(90,90))
	// Output:
	// 0.2500⟳ 0.3750⟳ 0.5000⟳
	// 0.7500⟳ 0.1250⟳ 0.5000⟳
	// 10.0000° 180.0000° 350.0000°
	// 350.0000° 360.0000° 10.0000°
}

func ExampleSweepCCWDivided() {
	s:=Sweep{3*RightAngle,NewCCWTo(RightAngle)}
	fmt.Printf("%+.4t %+.4t %+.4t\n",s.Intermediate(90,0),s.Intermediate(90,45),s.Intermediate(90,90))
	s=Sweep{2*RightAngle,NewCCWTo(3*RightAngle)}
	fmt.Printf("%+.4t %+.4t %+.4t\n",s.Intermediate(90,0),s.Intermediate(90,45),s.Intermediate(90,90))
	s=Sweep{10*Degree,NewCCWTo(20*Degree)}
	fmt.Printf("%+.4d %+.4d %+.4d\n",s.Intermediate(90,0),s.Intermediate(90,45),s.Intermediate(90,90))
	s=Sweep{350*Degree,NewCCWTo(340*Degree)}
	fmt.Printf("%+.4d %+.4d %+.4d\n",s.Intermediate(90,0),s.Intermediate(90,45),s.Intermediate(90,90))
	// Output:
	// 0.7500⟳ 0.6250⟳ 0.5000⟳
	// 0.5000⟳ 0.1250⟳ 0.7500⟳
	// 10.0000° 0.0000° 350.0000°
	// 350.0000° 180.0000° 10.0000°
}

