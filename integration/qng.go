package integration

/*
#cgo pkg-config: gsl
#include "integration.h"
*/
import "C"
import "unsafe"

func Qng(f F, a, b, epsabs, epsrel float64) (result, abserr float64, neval, status int) {
	wf := C.wrapper_function(unsafe.Pointer(&f))
	r := C.qng(wf, C.double(a), C.double(b), C.double(epsabs), C.double(epsrel))
	result = float64(r.result)
	abserr = float64(r.abserr)
	status = int(r.status)
	neval = int(r.neval)
	return
}
