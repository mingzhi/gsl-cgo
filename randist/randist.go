/*
 *   Copyright (C) 2012 Mingzhi Lin
 *
 * Permission is hereby granted, free of charge, to any person obtaining 
 * a copy of this software and associated documentation files (the "Software"), 
 * to deal in the Software without restriction, including without limitation 
 * the rights to use, copy, modify, merge, publish, distribute, sublicense, 
 * and/or sell copies of the Software, and to permit persons to whom the 
 * Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included 
 * in all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS 
 * OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, 
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE 
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER 
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, 
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE 
 * SOFTWARE.
 */
package randist

//#cgo pkg-config: gsl
//#include <gsl/gsl_rng.h>
import "C"

type ContinuousDistribution interface {
	Pdf(float64) float64
	Cdf(float64) float64
	RandomFloat64() float64
}

type DiscreteDistribution interface {
	Pdf(int) float64
	Cdf(int) float64
	RandomInt() int
}

const (
	BOROSH13 = iota
	CMRG
	COVEYOU
	FISHMAN18
	FISHMAN20
	FISHMAN2X
	GFSR4
	KNUTHRAN
	KNUTHRAN2
	KNUTHRAN2002
	LECUYER21
	MINSTD
	MRG
	MT19937
	MT19937_1999
	MT19937_1998
	R250
	RAN0
	RAN1
	RAN2
	RAN3
	RAND
	RAND48
	RANDOM128_BSD
	RANDOM128_GLIBC2
	RANDOM128_LIBC5
	RANDOM256_BSD
	RANDOM256_GLIBC2
	RANDOM256_LIBC5
	RANDOM32_BSD
	RANDOM32_GLIBC2
	RANDOM32_LIBC5
	RANDOM64_BSD
	RANDOM64_GLIBC2
	RANDOM64_LIBC5
	RANDOM8_BSD
	RANDOM8_GLIBC2
	RANDOM8_LIBC5
	RANDOM_BSD
	RANDOM_GLIBC2
	RANDOM_LIBC5
	RANDU
	RANF
	RANLUX
	RANLUX389
	RANLXD1
	RANLXD2
	RANLXS0
	RANLXS1
	RANLXS2
	RANMAR
	SLATEC
	TAUS
	TAUS2
	TAUS113
	TRANSPUTER
	TT800
	UNI
	UNI32
	VAX
	WATERMAN14
	ZUF
)
