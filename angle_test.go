package angle

import "fmt"
import "testing"

func ExampleAngle() {
	fmt.Printf("%.1v degrees == %+[1]r == %+[1]v == %+.0[1]m == %+.2[1]l\n", Radian)
	fmt.Printf("%.1v degrees == %+[1]㎭ == %+.1[1]l == %+.0[1]s == %+.1[1]g == %+.1[1]f\n", RightAngle)
	// Output:
	// 57.3 degrees == 1㎭ == 57.295784° == 3438′ == 57°18′44.82″
	// 90.0 degrees == 1.5707964㎭ == 90°0′0.0″ == 324002″ == 100.0ᵍ

}

func ExampleAngleAdd() {
	var ra = RightAngle
	fmt.Printf("%+.2d\n", ra)
	fmt.Printf("%+.2d\n", 0-ra)
	fmt.Printf("%+.2d\n", 53*ra)
	// Output:
	// 90.00°
	// 270.00°
	// 90.00°
}

func ExampleAngleCourse() {
	fmt.Printf("%c %+[1]d\n", angle(0))
	fmt.Printf("%c %+[1]d\n", Degree*180)
	fmt.Printf("%c %+[1]d\n", Degree*348)
	fmt.Printf("%c %+[1]d\n", Degree*349)
	// Output:
	// N 0°
	// S 180°
	// NNW 348°
	// N 349°
}

func ExampleAngleMin() {
	fmt.Printf("%.4[1]rμ㎭\n", angle(1)*1000000)
	// Output:
	// 0.0015μ㎭
}

func ExampleAngleBinary() {
	fmt.Printf("%+.2[1]d %+.0[1]l\n", angle(1<<(bits-8)))
	// Output:
	// 1.41° 1°24′23″
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


func TestAngle(t *testing.T) {

	//	t.Error(args ...any)
	//// equivalent to Log followed by Fail.

	//	t.Errorf(format string, args ...any)
	//// equivalent to Logf followed by Fail.

	//	Fail()
	//// marks the function as having failed but continues execution.

	//	t.FailNow()
	//// marks the function as having failed and stops its execution by calling runtime.Goexit (which then runs all deferred calls in the current goroutine). Execution will continue at the next test or benchmark. FailNow must be called from the goroutine running the test or benchmark function,not from other goroutines created during the test. Calling FailNow does not stop those other goroutines.

	//	t.Failed() bool
	//// reports whether the function has failed.

	//	t.Fatal(args ...any)
	//// equivalent to Log followed by FailNow.

	//	t.Fatalf(format string, args ...any)
	//// equivalent to Logf followed by FailNow.

}

func BenchmarkAngle(b *testing.B) {

	//	b.StartTimer()
	////    StartTimer starts timing a test. This function is called automatically before a benchmark starts, but it can also be used to resume timing after a call to StopTimer.

	for i := 0; i < b.N; i++ {
	}

	//	b.StopTimer()
	////    StopTimer stops timing a test. This can be used to pause the timer while performing complex initialization that you don't want to measure.

}
