package integration

/*
#cgo pkg-config: gsl
#include "integration.h"
my_result
qng(const gsl_function f, double a, double b, double epsabs, double epsrel) {
	double result = 0, abserr = 0;
	size_t neval;
	int status;
	my_result r;

	status = gsl_integration_qng(&f, a, b, epsabs, epsrel, &result, &abserr, &neval);
	r.result = result;
	r.abserr = abserr;
	r.status = status;
	r.neval = neval;

	return r;
}
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
