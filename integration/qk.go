/*
 *   Copyright (C) 2012 Mingzhi Lin
 *
 * Permission is hereby granted, free of charge, to any person obtaining 
 * a copy of this software and associated documentation files (the "Software"), 
 * to deal in the Software without restriction, including without limitation 
 * the rights to use, copy, modify, merge, publish, distribute, sublicense, 
 * and/or sell copies of the Software, and to permit persons to whom the 
 * Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included 
 * in all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS 
 * OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, 
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE 
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER 
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, 
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE 
 * SOFTWARE.
 */
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
