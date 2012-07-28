package specfunc

//#cgo pkg-config: gsl
//#include <gsl/gsl_sf_gamma.h>
import "C"

func Lngamma(x float64) (r float64) {
	r = float64(C.gsl_sf_lngamma(C.double(x)))
	return
}
