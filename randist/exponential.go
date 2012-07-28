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
	randomGenerator *C.gsl_rng
}

func NewExponential(lambda float64, rng int) (e *Exponential) {
	rg := newRandomGenerator(rng)
	e = &Exponential{Lambda: lambda, randomGenerator: rg}
	return
}

func (e *Exponential) SetRandomGenerator(rng int) {
	rg := newRandomGenerator(rng)
	e.randomGenerator = rg
}

func (e *Exponential) FreeRandomGenerator() {
	freeRandomGenerator(e.randomGenerator)
}

func (e *Exponential) Pdf(x float64) (p float64) {
	p = ExponentialPdf(x, e.Lambda)
	return
}

func (e *Exponential) Cdf(x float64) (p float64) {
	panic("Have not implement this function: Exponential.Cdf(float64)")
	return
}

func (e *Exponential) RandomFloat64() (x float64) {
	x = ExponentialRandomFloat64(e.randomGenerator, e.Lambda)
	return
}

func ExponentialRandomFloat64(rd *C.gsl_rng, lambda float64) (x float64) {
	x = float64(C.gsl_ran_exponential(rd, C.double(lambda)))
	return x
}

func ExponentialPdf(x, lambda float64) float64 {
	return float64(C.gsl_ran_exponential_pdf(C.double(x), C.double(lambda)))
}
