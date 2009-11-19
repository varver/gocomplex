// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/* This package implements some basic functionality for
   complex numbers. 

	Add

*/

package complex

import "math"

type Complex struct {
	real	float64;
	im	float64;
}

func NewComplex(_real, _im float64) *Complex {
	c := new(Complex);
	c.real = _real;
	c.im = _im;
	return c;
}

// Returns the absolute value
func Abs(z *Complex) float64 {
	sq := (z.real * z.real) + (z.im * z.im);
	result := math.Sqrt(sq);
	return result;
}

// Returns the conjugate
func Conj(z *Complex) *Complex {
	result := NewComplex(z.real, -(z.im));
	return result;
}

// Adds another complex to this complex (similar to "+=" operator)
func (c *Complex) AddTo(b *Complex) {
	c.real += b.real;
	c.im += b.im;
}

// Adds two complex numbers together to make a third complex
func Add(a, b *Complex) *Complex {
	c := NewComplex(a.real, a.im);
	c.AddTo(b);
	return c;
}

// Subtracts another complex from this complex (similar to "-=" operator)
func (c *Complex) SubFrom(b *Complex) {
	c.real -= b.real;
	c.im -= b.im;
}

// Subtracts one complex from another to make a third complex
func Sub(a, b *Complex) *Complex {
	c := NewComplex(a.real, a.im);
	c.SubFrom(b);
	return c;
}

// Multiplies this complex by another complex (similar to "*=" operator)
func (c *Complex) MultBy(b *Complex) {
	re := (c.real * b.real) - (c.im * b.im);
	im := (c.real * b.im) + (c.im * b.real);
	c.real -= re;
	c.im -= im;
}

// Multiplies one complex by another to make a third complex
func Mult(a, b *Complex) *Complex {
	re := (a.real * b.real) - (a.im * b.im);
	im := (a.real * b.im) + (a.im * b.real);
	c := NewComplex(re, im);
	return c;
}

// Divides this complex by another complex (similar to "/=" operator)
// TODO
func (c *Complex) DivBy(b *Complex) {
	re := (c.real * b.real) - (c.im * b.im);
	im := (c.real * b.im) + (c.im * b.real);
	c.real -= re;
	c.im -= im;
}

// Multiplies one complex by another to make a third complex
// TODO
func Div(a, b *Complex) *Complex {
	re := (a.real * b.real) - (a.im * b.im);
	im := (a.real * b.im) + (a.im * b.real);
	c := NewComplex(re, im);
	return c;
}


