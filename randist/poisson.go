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

type Poisson struct {
	Lambda          float64
	randomGenerator *C.gsl_rng
}

func NewPoisson(lambda float64, rng int) (psn *Poisson) {
	rg := newRandomGenerator(rng)
	psn = &Poisson{Lambda: lambda, randomGenerator: rg}
	return
}

func (psn *Poisson) FreeRandomGenerator() {
	freeRandomGenerator(psn.randomGenerator)
}

func (psn *Poisson) SetRandomGenerator(i int) {
	// free the previous generator first
	psn.FreeRandomGenerator()
	rg := newRandomGenerator(i)
	psn.randomGenerator = rg
}

func (psn *Poisson) Pdf(k int) float64 {
	return PoissonPdf(k, psn.Lambda)
}

func (psn *Poisson) Cdf(k int) float64 {
	panic("Have not implemented this function: Poisson.Cdf")
	return 0
}

func (psn *Poisson) RandomInt() int {
	return PoissonRandomInt(psn.randomGenerator, psn.Lambda)
}

func PoissonPdf(k int, lambda float64) float64 {
	if k < 0 {
		return 0
	}
	return float64(C.gsl_ran_poisson_pdf(C.uint(k), C.double(lambda)))
}

func PoissonRandomInt(rng *C.gsl_rng, lambda float64) int {
	return int(C.gsl_ran_poisson(rng, C.double(lambda)))
}
