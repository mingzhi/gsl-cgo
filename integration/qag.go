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
