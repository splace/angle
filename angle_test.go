package angle

import "fmt"
import "testing"

func ExampleAnglePrint() {
	fmt.Printf("%.1v degrees == %+[1]r == %+[1]v == %+.0[1]l == %+.2[1]f\n", Radian)
	fmt.Printf("%.1v degrees == %+[1]㎭ == %+.0[1]l == %+.1[1]g == %+.2[1]f\n", RightAngle)
	// Output:
	// 57.3 degrees == 1㎭ == 57.295784° == 57°18′45″ == 15.92%
	// 90.0 degrees == 1.5707964㎭ == 90°0′0″ == 100.0ᵍ == 25.00%
}

func ExampleAngle() {
	var ra = RightAngle
	fmt.Printf("%+.2d\n", ra)
	fmt.Printf("%+.2d\n", 0-ra)
	fmt.Printf("%+.2d\n", 53*ra)
	// Output:
	// 90.00°
	// 270.00°
	// 90.00°
}

func ExampleAngleMin() {
	fmt.Printf("%+.5s == %.4rμ㎭\n", Angle(1), Angle(1)*1000000)
	// Output:
	// 0.00030″ == 0.0015μ㎭
}

func ExampleAngleBinary() {
	fmt.Printf("%+.2[1]b == %+.0[1]l\n", Angle(1<<(bits-8)))
	// Output:
	// 1.00b == 1°24′23″
}

func ExampleAngleMultiply() {
	fmt.Printf("%+.4d\n", Second*60*60)
	// Output:
	// 1.0000°
}

func ExampleAngleMultiplyOverflow() {
	a := Degree // have to use var in order to prevent out of range message
	fmt.Printf("%+.3d\n", a*1000)
	fmt.Printf("%+.3d\n", -a*1000)
	// Output:
	// 280.000°
	// 80.000°
}

func ExampleAngleTurns() {
	fmt.Printf("%+.2t\n", RightAngle)
	// Output:
	// 0.25⟳
}

func ExampleAngleRotations() {
	fmt.Printf("%+.2t\n", Rotations(.5))
	fmt.Printf("%+.2t\n", Rotations(-2.5))
	// Output:
	// 0.50⟳
	// 0.50⟳
}

