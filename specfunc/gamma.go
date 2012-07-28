package specfunc

//#cgo pkg-config: gsl
//#include <gsl/gsl_sf_gamma.h>
import "C"

/* Log[Gamma(x)], x not a negative integer
 * Uses real Lanczos method.
 * Returns the real part of Log[Gamma[x]] when x < 0,
 * i.e. Log[|Gamma[x]|].
 */
func Lngamma(x float64) (r float64) {
	r = float64(C.gsl_sf_lngamma(C.double(x)))
	return
}

/* Gamma(x), x not a negative integer
 * Uses real Lanczos method.
 */
func Gamma(x float64) (r float64) {
	r = float64(C.gsl_sf_gamma(C.double(x)))
	return
}

/* Regulated Gamma Function, x > 0
 * Gamma^*(x) = Gamma(x)/(Sqrt[2Pi] x^(x-1/2) exp(-x))
 *            = (1 + 1/(12x) + ...),  x->Inf
 */
func Gammastar(x float64) (r float64) {
	r = float64(C.gsl_sf_gammastar(C.double(x)))
	return
}
