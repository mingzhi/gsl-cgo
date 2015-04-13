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
#include <gsl/gsl_randist.h>
#include <gsl/gsl_rng.h>
*/
import "C"

//Exponential distribution
// p(x) dx = exp(-x/lambda) dx/lambda
// for x = 0 ... +infty
type Exponential struct {
	Lambda          float64
	RandomGenerator *RNG
}

func NewExponential(lambda float64, rng *RNG) (e *Exponential) {
	e = &Exponential{Lambda: lambda, RandomGenerator: rng}
	return
}

func (e *Exponential) Pdf(x float64) (p float64) {
	p = ExponentialPdf(x, e.Lambda)
	return
}

func (e *Exponential) Cdf(x float64) (p float64) {
	panic("Have not implement this function: Exponential.Cdf(float64)")
	return
}

func (e *Exponential) Float64() (x float64) {
	x = ExponentialRandomFloat64(e.RandomGenerator, e.Lambda)
	return
}

func ExponentialRandomFloat64(rd *RNG, lambda float64) (x float64) {
	x = float64(C.gsl_ran_exponential(rd.g, C.double(lambda)))
	return x
}

func ExponentialPdf(x, lambda float64) float64 {
	return float64(C.gsl_ran_exponential_pdf(C.double(x), C.double(lambda)))
}
