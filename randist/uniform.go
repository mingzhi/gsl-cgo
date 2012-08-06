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

/*
#cgo pkg-config: gsl
#include <gsl/gsl_rng.h>
*/
import "C"

type Uniform struct {
	RandomGenerator *RNG
}

func (uni *Uniform) RandomFloat64() float64 {
	return UniformRandomFloat64(uni.RandomGenerator)
}

func (uni *Uniform) RandomPos() float64 {
	return UniformRandomPos(uni.RandomGenerator)
}

func (uni *Uniform) RandomInt(n int) int {
	return UniformRandomInt(uni.RandomGenerator, n)
}

func UniformRandomFloat64(rng *RNG) float64 {
	return float64(C.gsl_rng_uniform(rng.g))
}

func UniformRandomPos(rng *RNG) float64 {
	return float64(C.gsl_rng_uniform_pos(rng.g))
}

func UniformRandomInt(rng *RNG, n int) int {
	return int(C.gsl_rng_uniform_int(rng.g, C.ulong(n)))
}

func UniformGet(rng *RNG) int {
	return int(C.gsl_rng_get(rng.g))
}
