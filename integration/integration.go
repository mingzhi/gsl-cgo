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
