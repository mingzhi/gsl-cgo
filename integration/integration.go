package integration

//#include <gsl/gsl_math.h>
//#include <gsl/gsl_integration.h>
import "C"
import "unsafe"

type F interface {
	Evaluate(float64) float64
}

const (
	GSL_INTEG_GAUSS15 = 1 /* 15 point Gauss-Kronrod rule */
	GSL_INTEG_GAUSS21 = 2 /* 21 point Gauss-Kronrod rule */
	GSL_INTEG_GAUSS31 = 3 /* 31 point Gauss-Kronrod rule */
	GSL_INTEG_GAUSS41 = 4 /* 41 point Gauss-Kronrod rule */
	GSL_INTEG_GAUSS51 = 5 /* 51 point Gauss-Kronrod rule */
	GSL_INTEG_GAUSS61 = 6 /* 61 point Gauss-Kronrod rule */
)

type ftype struct {
	function F
}

func (f *ftype) evaluate(x float64) float64 {
	return f.function.Evaluate(x)
}

//exported function for C module to deal with Go F Object
//export fevaluate
func fevaluate(x C.double, p unsafe.Pointer) C.double {
	f := (*ftype)(p)
	r := C.double(f.evaluate(float64(x)))
	return r
}
