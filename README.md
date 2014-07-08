half
====

This package provides IEEE 754 compliant functions for converting to and
from half-precision, 16-bit, floating point numbers. It properly handles
NaNs, infinities, and subnormal numbers.

###Purpose

This package is primarily intended for use with OpenGL. By specifying vertexes
or textures as 16-bit floats, you can save video memory, improve buffer
upload times, and draw faster. Besides OpenGL, 16-bit floats are also used in
many image formats, such as OpenEXR, for high dynamic range and deep color.

###Limitations
The range and precision of a 16-bit float are much less than a 32-bit float.
They can represent positive and negative values as small as 0.000070 and as
large as 65472.0 with approximately 3 decimal digits of precision. For
example, pi is rounded to 3.140625.

###Credit
The conversion functions are based on a Stack Overflow post by Paul
"Phernost" Tessier:
http://stackoverflow.com/questions/1659440/32-bit-to-16-bit-floating-point-conversion
