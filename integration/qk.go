package integration

/*
#cgo pkg-config: gsl
#include "integration.h"
*/
import "C"
import "unsafe"

type QkFunction func(F, float64, float64) (float64, float64, float64, float64)

func Qk15(f F, a, b float64) (float64, float64, float64, float64) {
	key := GSL_INTEG_GAUSS15
	return Qk(f, a, b, key)
}

func Qk21(f F, a, b float64) (float64, float64, float64, float64) {
	key := GSL_INTEG_GAUSS21
	return Qk(f, a, b, key)
}

func Qk31(f F, a, b float64) (float64, float64, float64, float64) {
	key := GSL_INTEG_GAUSS31
	return Qk(f, a, b, key)
}

func Qk41(f F, a, b float64) (float64, float64, float64, float64) {
	key := GSL_INTEG_GAUSS41
	return Qk(f, a, b, key)
}

func Qk51(f F, a, b float64) (float64, float64, float64, float64) {
	key := GSL_INTEG_GAUSS51
	return Qk(f, a, b, key)
}

func Qk61(f F, a, b float64) (float64, float64, float64, float64) {
	key := GSL_INTEG_GAUSS61
	return Qk(f, a, b, key)
}

func Qk(f F, a, b float64, key int) (result, abserr, resabs, resasc float64) {
	wf := C.wrapper_function(unsafe.Pointer(&f))

	r := C.qk(wf, C.int(key), C.double(a), C.double(b))
	result = float64(r.result)
	abserr = float64(r.abserr)
	resabs = float64(r.resabs)
	resasc = float64(r.resasc)
	return
}
