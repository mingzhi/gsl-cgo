package randist

/*
#cgo pkg-config: gsl
#include <gsl/gsl_randist.h>
#include <gsl/gsl_rng.h>
*/
import "C"

type Gaussian struct {
	Sigma           float64
	RandomGenerator *RNG
}

func NewGaussian(sigma float64, rng *RNG) *Gaussian {
	var g Gaussian
	g.RandomGenerator = rng
	g.Sigma = sigma
	return &g
}

func (g *Gaussian) Float64() float64 {
	var x float64
	x = GaussianRandomFloat64(g.RandomGenerator, g.Sigma)
	return x
}

func GaussianRandomFloat64(rd *RNG, sigma float64) float64 {
	var x float64
	x = float64(C.gsl_ran_gaussian(rd.g, C.double(sigma)))
	return x
}
