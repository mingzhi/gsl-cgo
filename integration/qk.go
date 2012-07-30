package integration

/*
#cgo pkg-config: gsl
#include "integration.h"
*/
import "C"
import "unsafe"

type QkFunction func(F, float64, float64) (float64, float64, float64, float64)

func Qk15(f F, a, b float64) (float64, float64, float64, float64) {
	rule := 15
	return qk(f, rule, a, b)
}

func Qk21(f F, a, b float64) (float64, float64, float64, float64) {
	rule := 21
	return qk(f, rule, a, b)
}

func Qk31(f F, a, b float64) (float64, float64, float64, float64) {
	rule := 31
	return qk(f, rule, a, b)
}

func Qk41(f F, a, b float64) (float64, float64, float64, float64) {
	rule := 41
	return qk(f, rule, a, b)
}

func Qk51(f F, a, b float64) (float64, float64, float64, float64) {
	rule := 51
	return qk(f, rule, a, b)
}

func Qk61(f F, a, b float64) (float64, float64, float64, float64) {
	rule := 61
	return qk(f, rule, a, b)
}

func qk(f F, rule int, a, b float64) (result, abserr, resabs, resasc float64) {
	wf := C.wrapper_function(unsafe.Pointer(&f))

	r := C.qk(wf, C.int(rule), C.double(a), C.double(b))
	result = float64(r.result)
	abserr = float64(r.abserr)
	resabs = float64(r.resabs)
	resasc = float64(r.resasc)
	return
}
