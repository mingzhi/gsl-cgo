package integration

//#include <gsl/gsl_math.h>
//#include <gsl/gsl_integration.h>
import "C"
import "unsafe"

type F interface {
	Evaluate(float64) float64
}

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
