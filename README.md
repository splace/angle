# Angle

an attempt at a general nice-n-clean (hiding the solution-space) way of using angles.

a 'typed' angle. potentially a pattern for other simular.

(could be compared with std. lib. Time/Duration.)

Overview/docs: [![GoDoc](https://godoc.org/github.com/splace/angle?status.svg)](https://godoc.org/github.com/splace/angle)

different scalings (degrees, radians etc) are out-of-context. Only for human readability or access to external processes. (like hardware acceleration requiring radians)

example of various human readable formats...

``` golang

func ExampleAngles() {
	fmt.Printf("%.1v degrees == %+[1]r == %+[1]v == %+.0[1]m == %+.2[1]l\n", Radian)
	fmt.Printf("%.1v degrees == %+[1]㎭ == %+.1[1]l == %+.0[1]s == %+.1[1]g\n", RightAngle)
	// Output:
	// 57.3 degrees == 1㎭ == 57.295784° == 3438′ == 57°18′44.82″
	// 90.0 degrees == 1.5707964㎭ == 90°0′0.0″ == 324002″ == 100.0ᵍ

}
```

# encoded as integers.

Here an Angle is a uint32 with its whole range representing one revolution.

Since its max approaches one rotation, its modulus behaviour matches a rotation modulus, so you get restricted within one revolution automatically. 

Notice: 'real' Angles arn't multipled by other angles.

Angles, rather than angle differences, are symetrical, no particular value is special, so a float representation with its higher precision closer to the zero value, is a mismatched behaviour.

Formula, say involving sin/cos, with intermediate steps involving small angles, needs to be handled with floats throughout (unless rounding errors when using this Angle is determined to be OK). these intermediate steps might be considered as not being Angles but Angles differences, so this might be expected. 

360 degrees (or 2Pi radians etc.) is just 0, and so is encoded/returned as 0 degrees. ( or 0 radians etc).

Power of two fractions of a rotation, are represented exactly, eg. 64*BinaryDegree==RightAngle, but in general multiplying a scaled angle can result in an in-exact representation, eg. 90*Degree!=RightAngle, (but RightAngle/90==Degree) use the usual approachs to limit rounding errors.

Note: constants report an out of range error when used beyond one rotation, replace with variables.

# Sector

encodes a region between two angles, necessererily requiring a direction. (Counter)Clockwise.

example: range clockwise 20 gradiens from 390 gradians, show degree.
Note 10gradians == 9degrees, so 9 divisions splitting 20 gradiens should be 10 angles in 2 degree steps.

[Sweep Iteration Example](https://go.dev/play/p/H5fRNCTrfnS)



