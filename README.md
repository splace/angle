# angle
dealing with angles in a sensible way, or, angles typed

formats as human readable as used in various fields

``` golang

func ExampleAngles() {
	fmt.Printf("%.1v degrees == %+[1]r == %+[1]v == %+.0[1]m == %+.2[1]l\n", Radian)
	fmt.Printf("%.1v degrees == %+[1]㎭ == %+.1[1]l == %+.0[1]s == %+.1[1]g\n", RightAngle)
	// Output:
	// 57.3 degrees == 1㎭ == 57.295784° == 3438′ == 57°18′44.82″
	// 90.0 degrees == 1.5707964㎭ == 90°0′0.0″ == 324002″ == 100.0ᵍ

}


```

converters to all major unit scales.

# angles encoded as ints.

An Angle is a uint32 with its whole range as one revolution.

Since its max approaches one rotation, its modulus behaviour matches a rotation modulus, so you get restricted within one revolution automatically.

Note: constants report an out of range error when used beyond one rotation, replace with variables.
Angles dont make sense when multipled by other angles.

Where a float representation would have higher precision the closer to zero value, Angle has fixed precision and also away from zero is more precise than a float32.

360 degrees (or 2Pi radians etc.) is just 0, and so is encoded/returned as 0 degrees. ( or 0 radians etc).

Power of two fractions of a rotation are represented exactly, eg. 64*BinaryDegree==RightAngle, but in general multiplying a unit can result in an in-exact representation, eg. 90*Degree!=RightAngle, (but RightAngle/90==Degree) use the usual approachs to limit rounding errors.
