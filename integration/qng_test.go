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

import "testing"
import "bitbucket.org/mingzhi/goutils/assert"

func TestQng(t *testing.T) {
	{
		exp_result := 7.716049379303083211E-02
		exp_abserr := 9.424302199601294244E-08
		exp_neval := 21
		f := f1{alpha: 2.6}
		result, abserr, neval, status := Qng(f, 0.0, 1.0, 1e-1, 0.0)
		if status != 0 {
			t.Errorf("Don't expect Error code: %d\n", status)
		}
		if !assert.EqualFloat64(result, exp_result, 1e-15, 0) {
			t.Errorf("Result: %f, expect: %f\n", result, exp_result)
		}
		if !assert.EqualFloat64(abserr, exp_abserr, 1e-7, 0) {
			t.Errorf("Abserr: %f, expect: %f\n", abserr, exp_abserr)
		}
		if neval != exp_neval {
			t.Errorf("Neval: %d, expect: %d\n", neval, exp_neval)
		}
	}
}
