package angle

import "fmt"

// range clockwise 20 gradians from 390 gradians, show degree.
// Note 10gradiens == 9degrees, so 9 divisions splitting 20 gradiens should be 10 angles in 2 degree steps.
func ExampleSweepStepped() {
	s:=Gradian*10 // slightly more precise than using bit multiplier ie Gradian*390
	for a:= range Over(NewRange(-s,Gradian*20,true),9){
		fmt.Printf("%+.4v ",a)
	}
	// Output:
	// 351.0000° 353.0000° 355.0000° 357.0000° 359.0000° 1.0000° 3.0000° 5.0000° 7.0000° 9.0000°
}

func ExampleSweepConains() {
	fmt.Println(SweepCW{RightAngle,3*RightAngle}.Contains(RightAngle))
	fmt.Println(SweepCW{3*RightAngle,2*RightAngle}.Contains(0))
	fmt.Println(SweepCW{3*RightAngle,2*RightAngle}.Contains(RightAngle))
	fmt.Println(SweepCW{0,2*RightAngle}.Contains(3*RightAngle))
	fmt.Println(SweepCW{2*RightAngle,0}.Contains(RightAngle))
	fmt.Println(SweepCW{3*RightAngle,RightAngle}.Contains(0))

	fmt.Println(SweepCCW{RightAngle,3*RightAngle}.Contains(2*RightAngle))
	fmt.Println(SweepCCW{3*RightAngle,2*RightAngle}.Contains(0))
	fmt.Println(SweepCCW{3*RightAngle,2*RightAngle}.Contains(RightAngle))
	fmt.Println(SweepCCW{0,2*RightAngle}.Contains(3*RightAngle))
	fmt.Println(SweepCCW{2*RightAngle,0}.Contains(RightAngle))
	fmt.Println(SweepCCW{3*RightAngle,RightAngle}.Contains(0))
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
	s:=SweepCW{RightAngle,2*RightAngle}
	fmt.Printf("%+.4t %+.4t %+.4t\n",s.Intermediate(90,0),s.Intermediate(90,45),s.Intermediate(90,90))
	s=SweepCW{3*RightAngle,2*RightAngle}
	fmt.Printf("%+.4t %+.4t %+.4t\n",s.Intermediate(90,0),s.Intermediate(90,45),s.Intermediate(90,90))
	s=SweepCW{10*Degree,350*Degree}
	fmt.Printf("%+.4d %+.4d %+.4d\n",s.Intermediate(90,0),s.Intermediate(90,45),s.Intermediate(90,90))
	s=SweepCW{350*Degree,10*Degree}
	fmt.Printf("%+.4d %+.4d %+.4d\n",s.Intermediate(90,0),s.Intermediate(90,45),s.Intermediate(90,90))
	// Output:
	// 0.2500⟳ 0.3750⟳ 0.5000⟳
	// 0.7500⟳ 0.1250⟳ 0.5000⟳
	// 10.0000° 180.0000° 350.0000°
	// 350.0000° 360.0000° 10.0000°
}

func ExampleSweepCCWDivided() {
	s:=SweepCCW{3*RightAngle,2*RightAngle}
	fmt.Printf("%+.4t %+.4t %+.4t\n",s.Intermediate(90,0),s.Intermediate(90,45),s.Intermediate(90,90))
	s=SweepCCW{2*RightAngle,3*RightAngle}
	fmt.Printf("%+.4t %+.4t %+.4t\n",s.Intermediate(90,0),s.Intermediate(90,45),s.Intermediate(90,90))
	s=SweepCCW{10*Degree,350*Degree}
	fmt.Printf("%+.4d %+.4d %+.4d\n",s.Intermediate(90,0),s.Intermediate(90,45),s.Intermediate(90,90))
	s=SweepCCW{350*Degree,10*Degree}
	fmt.Printf("%+.4d %+.4d %+.4d\n",s.Intermediate(90,0),s.Intermediate(90,45),s.Intermediate(90,90))
	// Output:
	// 0.7500⟳ 0.6250⟳ 0.5000⟳
	// 0.5000⟳ 0.1250⟳ 0.7500⟳
	// 10.0000° 360.0000° 350.0000°
	// 350.0000° 180.0000° 10.0000°
}

