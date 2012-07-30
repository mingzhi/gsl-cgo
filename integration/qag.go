package integration

/*
#cgo pkg-config: gsl
#include "integration.h"
*/
import "C"
import "unsafe"

func Qag(f F, a, b float64, epsabs, epsrel float64, limit int, key int) (result, abserr float64, status int) {
	wf := C.wrapper_function(unsafe.Pointer(&f))
	r := C.qag(wf, C.double(a), C.double(b), C.double(epsabs), C.double(epsrel), C.size_t(limit), C.int(key))
	result = float64(r.result)
	abserr = float64(r.abserr)
	status = int(r.status)
	return
}

func Qags(f F, a, b float64, epsabs, epsrel float64, limit int) (result, abserr float64, status int) {
	wf := C.wrapper_function(unsafe.Pointer(&f))
	r := C.qags(wf, C.double(a), C.double(b), C.double(epsabs), C.double(epsrel), C.size_t(limit))
	result = float64(r.result)
	abserr = float64(r.abserr)
	status = int(r.status)
	return
}
