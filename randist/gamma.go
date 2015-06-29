package randist

/*
#cgo pkg-config: gsl
#include <gsl/gsl_randist.h>
#include <gsl/gsl_rng.h>
*/
import "C"

type Gamma struct {
	Scale           float64
	Shape           float64
	RandomGenerator *RNG
}

func NewGamma(shape, scale float64, rng *RNG) *Gamma {
	var g Gamma
	g.RandomGenerator = rng
	g.Shape = shape
	g.Scale = scale
	return &g
}

func (g *Gamma) Float64() float64 {
	var x float64
	x = GammaRandomFloat64(g.RandomGenerator, g.Shape, g.Scale)
	return x
}

func GammaRandomFloat64(rd *RNG, shape, scale float64) float64 {
	var x float64
	x = float64(C.gsl_ran_gamma(rd.g, C.double(shape), C.double(scale)))
	return x
}
