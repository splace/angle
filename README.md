# angle

an Angle type. 

encoded as an unsigned integer with its whole range representing one revolution, creating a 'modulus arithmetic' type, the zero so being transparent in the problem-space. only up to 1 revolution can be represented, higher/lower vales are automatically converted to the same angle within one revolution. so multi-turn angles need to held as an Angle plus an integer number of rotations.

this modulus behaviour occurs fundamentally, so this type doesn't require/have its own maths operators.

potentially a pattern for other problem-space types with similar behaviour.

Overview/docs: [![GoDoc](https://godoc.org/github.com/splace/angle?status.svg)](https://godoc.org/github.com/splace/angle)

different scalings (degrees, radians etc) are out-of-context. Only for human readability or access to external processes. (like hardware acceleration requiring radians)

example of various human readable formats...

``` golang

func ExampleAngles() {
	fmt.Printf("%.1v degrees == %+[1]r == %+[1]v == %+.0[1]m == %+.0[1]l\n", Radian)
	fmt.Printf("%.1v degrees == %+[1]㎭ == %+.0[1]l == %+.0[1]s == %+.1[1]g\n", RightAngle)
	// Output:
	// 57.3 degrees == 1㎭ == 57.295784° == 3438′ == 57°18′44″
	// 90.0 degrees == 1.5707964㎭ == 90°0′0″ == 324002″ == 100.0ᵍ
}
```

# encoded as integers.

angles are symmetrical, no particular value is special, so a float representation with its higher precision closer to the zero value, is a mismatched behaviour.

angles (Angle/Phase and Direction types) shouldn't be multiplied with themselves or each other.

and, specifically, Direction's shouldn't be added/subtracted from other Directions.

Formula, say involving sin/cos, with intermediate steps involving small angles, need to be handled with floats throughout (unless rounding errors when using this angle is determined to be OK). these intermediate steps might be considered as not being angles but Angle/Phase differences, so this might be expected. 

360 degrees (or 2Pi radians etc.) is just 0, and so is encoded/returned as 0 degrees. ( or 0 radians etc).

Power of two fractions of a rotation, are represented exactly, eg. 64*BinaryDegree==RightAngle, but in general multiplying a scaled angle can result in an in-exact representation, eg. 90*Degree!=RightAngle, (but RightAngle/90==Degree) use the usual approaches to limit rounding errors.

Note: constants report an out of range error when used beyond one rotation, replace with variables.

# Sector

a Sector encodes an angular region using a Direction and an Angle/Phase, the Angle being the size of the region, in a positive sense, from the Direction. so Sections don't have a winding.

example: range over a Sector clockwise from 390 gradians to 10 gradians, show degree.

[Iteration Example](https://go.dev/play/p/j30uc46iTBb)

Note: 10 gradians == 9 degrees, so 9 divisions splitting 20 gradians should be 10 angles in 2 degree steps.



