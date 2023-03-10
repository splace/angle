# angle
dealing with angles in a sensible way, or, angles typed

# angles encoded as ints.

An Angle is a uint32 with its whole range as one revolution.

Since its max approaches one rotation, its modulus behaviour matches a rotation modulus, so you get restricted within one revolution automatically.

Note: constants report an out of range error when used beyond one rotation, replace with variables.
Angles dont make sense when multipled by other angles.

Where a float representation would have higher precision the closer to zero value, Angle has fixed precision and also away from zero is more precise than a float32.

360 degrees (or 2Pi radians etc.) is just 0, and so is encoded/returned as 0 degrees. ( or 0 radians etc).

Power of two fractions of a rotation are represented exactly, eg. 64*BinaryDegree==RightAngle, but in general multiplying a unit can result in an in-exact representation, eg. 90*Degree!=RightAngle, (but RightAngle/90==Degree) use the usual approachs to limit rounding errors.
