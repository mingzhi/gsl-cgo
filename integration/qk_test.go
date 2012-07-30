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

type float struct {
	value     float64
	precision float64
}

func testQK(f F, exp_result, exp_abserr, exp_resabs, exp_resasc float, qk QkFunction, a, b float64, prefix string, t *testing.T) {
	var result, abserr, resabs, resasc float64
	result, abserr, resabs, resasc = qk(f, a, b)
	if !assert.EqualFloat64(result, exp_result.value, exp_result.precision, 0) {
		t.Errorf("%s: result: %.15f, expected: %.15f\n", prefix, result, exp_result.value)
	}
	if !assert.EqualFloat64(abserr, exp_abserr.value, exp_abserr.precision, 0) {
		t.Errorf("%s: abserr: %.7f, expected: %.7f\n", prefix, abserr, exp_abserr.value)
	}
	if !assert.EqualFloat64(resabs, exp_resabs.value, exp_resabs.precision, 0) {
		t.Errorf("%s: resabs: %.15f, expected: %.15f\n", prefix, resabs, exp_resabs.value)
	}
	if !assert.EqualFloat64(resasc, exp_resasc.value, exp_resasc.precision, 0) {
		t.Errorf("%s: resasc: %.15f, expected: %.15f\n", prefix, resasc, exp_resasc.value)
	}

	result, abserr, resabs, resasc = qk(f, b, a)
	if !assert.EqualFloat64(result, -exp_result.value, exp_result.precision, 0) {
		t.Errorf("%s: result: %.15f, expected: %.15f\n", prefix, result, -exp_result.value)
	}
	if !assert.EqualFloat64(abserr, exp_abserr.value, exp_abserr.precision, 0) {
		t.Errorf("%s: abserr: %.7f, expected: %.7f\n", prefix, abserr, exp_abserr.value)
	}
	if !assert.EqualFloat64(resabs, exp_resabs.value, exp_resabs.precision, 0) {
		t.Errorf("%s: resabs: %.15f, expected: %.15f\n", prefix, resabs, exp_resabs.value)
	}
	if !assert.EqualFloat64(resasc, exp_resasc.value, exp_resasc.precision, 0) {
		t.Errorf("%s: resasc: %.15f, expected: %.15f\n", prefix, resasc, exp_resasc.value)
	}
}

func TestQk15(t *testing.T) {
	var (
		exp_result float
		exp_abserr float
		exp_resabs float
		exp_resasc float
		alpha      float64
		a, b       float64
		f          F
		qk         QkFunction
		prefix     string
	)
	qk = Qk15
	a = 0.0
	b = 1.0
	//Test the basic Gauss-Kronrod rules with a smooth positive function.
	exp_result = float{value: 7.716049357767090777E-02, precision: 1e-15}
	exp_abserr = float{value: 2.990224871000550874E-06, precision: 1e-7}
	exp_resabs = float{value: 7.716049357767090777E-02, precision: 1e-15}
	exp_resasc = float{value: 4.434273814139995384E-02, precision: 1e-15}

	alpha = 2.6
	f = f1{alpha: alpha}

	prefix = "Smooth"

	testQK(f, exp_result, exp_abserr, exp_resabs, exp_resasc, qk, a, b, prefix, t)

	// Now test the basic rules with a positive function that has a
	// singularity. This should give large values of abserr which would
	// find discrepancies in the abserr calculation.
	exp_result = float{value: 1.555688196612745777E+01, precision: 1e-7}
	exp_abserr = float{value: 2.350164577239293706E+01, precision: 1e-7}
	exp_resabs = float{value: 1.555688196612745777E+01, precision: 1e-7}
	exp_resasc = float{value: 2.350164577239293706E+01, precision: 1e-15}

	alpha = -0.9
	f = f1{alpha: alpha}

	prefix = "Singular"

	testQK(f, exp_result, exp_abserr, exp_resabs, exp_resasc, qk, a, b, prefix, t)

	// Test the basic Gauss-Kronrod rules with a smooth oscillating
	// function, over an unsymmetric range. This should find any
	// discrepancies in the abscissae.
	exp_result = float{value: -7.238969575483799046E-01, precision: 1e-15}
	exp_abserr = float{value: 8.760080200939757174E-06, precision: 1e-15}
	exp_resabs = float{value: 1.165564172429140788E+00, precision: 1e-15}
	exp_resasc = float{value: 9.334560307787327371E-01, precision: 1e-15}

	alpha = 1.3
	f = f3{alpha: alpha}

	a = 0.3
	b = 2.71

	prefix = "Oscill"

	testQK(f, exp_result, exp_abserr, exp_resabs, exp_resasc, qk, a, b, prefix, t)
}

func TestQk21(t *testing.T) {
	var (
		exp_result float
		exp_abserr float
		exp_resabs float
		exp_resasc float
		alpha      float64
		a, b       float64
		f          F
		prefix     string
		qk         QkFunction
	)
	qk = Qk21
	a = 0.0
	b = 1.0

	//Test the basic Gauss-Kronrod rules with a smooth positive function.
	exp_result = float{value: 7.716049379303084599E-02, precision: 1e-15}
	exp_abserr = float{value: 9.424302194248481445E-08, precision: 1e-7}
	exp_resabs = float{value: 7.716049379303084599E-02, precision: 1e-15}
	exp_resasc = float{value: 4.434311425038358484E-02, precision: 1e-15}

	alpha = 2.6
	f = f1{alpha: alpha}

	prefix = "Smooth"

	testQK(f, exp_result, exp_abserr, exp_resabs, exp_resasc, qk, a, b, prefix, t)

	// Now test the basic rules with a positive function that has a
	// singularity. This should give large values of abserr which would
	// find discrepancies in the abserr calculation.
	exp_result = float{value: 1.799045317938126232E+01, precision: 1e-7}
	exp_abserr = float{value: 2.782360287710622515E+01, precision: 1e-5}
	exp_resabs = float{value: 1.799045317938126232E+01, precision: 1e-7}
	exp_resasc = float{value: 2.782360287710622515E+01, precision: 1e-15}

	alpha = -0.9
	f = f1{alpha: alpha}

	prefix = "Singular"
	testQK(f, exp_result, exp_abserr, exp_resabs, exp_resasc, qk, a, b, prefix, t)

	//Test the basic Gauss-Kronrod rules with a smooth oscillating
	// function, over an unsymmetric range. This should find any
	// discrepancies in the abscissae.
	exp_result = float{value: -7.238969575482959717E-01, precision: 1e-15}
	exp_abserr = float{value: 7.999213141433641888E-11, precision: 1e-15}
	exp_resabs = float{value: 1.150829032708484023E+00, precision: 1e-15}
	exp_resasc = float{value: 9.297591249133687619E-01, precision: 1e-15}

	alpha = 1.3
	f = f3{alpha: alpha}

	a = 0.3
	b = 2.71

	prefix = "Oscill"
	testQK(f, exp_result, exp_abserr, exp_resabs, exp_resasc, qk, a, b, prefix, t)
}

func TestQk31(t *testing.T) {
	var (
		exp_result float
		exp_abserr float
		exp_resabs float
		exp_resasc float
		alpha      float64
		a, b       float64
		f          F
		prefix     string
		qk         QkFunction
	)
	qk = Qk31
	a = 0.0
	b = 1.0

	//Test the basic Gauss-Kronrod rules with a smooth positive function.
	exp_result = float{value: 7.716049382494900855E-02, precision: 1e-15}
	exp_abserr = float{value: 1.713503193600029893E-09, precision: 1e-7}
	exp_resabs = float{value: 7.716049382494900855E-02, precision: 1e-15}
	exp_resasc = float{value: 4.427995051868838933E-02, precision: 1e-15}

	alpha = 2.6
	f = f1{alpha: alpha}

	prefix = "Smooth"
	testQK(f, exp_result, exp_abserr, exp_resabs, exp_resasc, qk, a, b, prefix, t)

	// Now test the basic rules with a positive function that has a
	// singularity. This should give large values of abserr which would
	// find discrepancies in the abserr calculation.
	exp_result = float{value: 2.081873305159121657E+01, precision: 1e-7}
	exp_abserr = float{value: 3.296500137482590276E+01, precision: 1e-5}
	exp_resabs = float{value: 2.081873305159121301E+01, precision: 1e-7}
	exp_resasc = float{value: 3.296500137482590276E+01, precision: 1e-7}

	alpha = -0.9
	f = f1{alpha: alpha}

	prefix = "Singular"
	testQK(f, exp_result, exp_abserr, exp_resabs, exp_resasc, qk, a, b, prefix, t)

	//Test the basic Gauss-Kronrod rules with a smooth oscillating
	// function, over an unsymmetric range. This should find any
	// discrepancies in the abscissae.
	exp_result = float{value: -7.238969575482959717E-01, precision: 1e-15}
	exp_abserr = float{value: 1.285805464427459261E-14, precision: 1e-15}
	exp_resabs = float{value: 1.158150602093290571E+00, precision: 1e-15}
	exp_resasc = float{value: 9.277828092501518853E-01, precision: 1e-15}

	alpha = 1.3
	f = f3{alpha: alpha}

	a = 0.3
	b = 2.71

	prefix = "Oscill"
	testQK(f, exp_result, exp_abserr, exp_resabs, exp_resasc, qk, a, b, prefix, t)
}

func TestQk41(t *testing.T) {
	var (
		exp_result float
		exp_abserr float
		exp_resabs float
		exp_resasc float
		alpha      float64
		a, b       float64
		f          F
		prefix     string
		qk         QkFunction
	)
	qk = Qk41
	a = 0.0
	b = 1.0

	//Test the basic Gauss-Kronrod rules with a smooth positive function.
	exp_result = float{value: 7.716049382681375302E-02, precision: 1e-15}
	exp_abserr = float{value: 9.576386660975511224E-11, precision: 1e-7}
	exp_resabs = float{value: 7.716049382681375302E-02, precision: 1e-15}
	exp_resasc = float{value: 4.421521169637691873E-02, precision: 1e-15}

	alpha = 2.6
	f = f1{alpha: alpha}

	prefix = "Smooth"
	testQK(f, exp_result, exp_abserr, exp_resabs, exp_resasc, qk, a, b, prefix, t)

	// Now test the basic rules with a positive function that has a
	// singularity. This should give large values of abserr which would
	// find discrepancies in the abserr calculation.
	exp_result = float{value: 2.288677623903126701E+01, precision: 1e-15}
	exp_abserr = float{value: 3.671538820274916048E+01, precision: 1e-5}
	exp_resabs = float{value: 2.288677623903126701E+01, precision: 1e-15}
	exp_resasc = float{value: 3.671538820274916048E+01, precision: 1e-7}

	alpha = -0.9
	f = f1{alpha: alpha}

	prefix = "Singular"
	testQK(f, exp_result, exp_abserr, exp_resabs, exp_resasc, qk, a, b, prefix, t)

	//Test the basic Gauss-Kronrod rules with a smooth oscillating
	// function, over an unsymmetric range. This should find any
	// discrepancies in the abscissae.
	exp_result = float{value: -7.238969575482959717E-01, precision: 1e-15}
	exp_abserr = float{value: 1.286535726271015626E-14, precision: 1e-15}
	exp_resabs = float{value: 1.158808363486595328E+00, precision: 1e-15}
	exp_resasc = float{value: 9.264382258645686985E-01, precision: 1e-7}

	alpha = 1.3
	f = f3{alpha: alpha}

	a = 0.3
	b = 2.71

	prefix = "Oscill"
	testQK(f, exp_result, exp_abserr, exp_resabs, exp_resasc, qk, a, b, prefix, t)
}

func TestQk51(t *testing.T) {
	var (
		exp_result float
		exp_abserr float
		exp_resabs float
		exp_resasc float
		alpha      float64
		a, b       float64
		f          F
		prefix     string
		qk         QkFunction
	)
	qk = Qk51
	a = 0.0
	b = 1.0

	//Test the basic Gauss-Kronrod rules with a smooth positive function.
	exp_result = float{value: 7.716049382708510540E-02, precision: 1e-15}
	exp_abserr = float{value: 1.002079980317363772E-11, precision: 1e-5}
	exp_resabs = float{value: 7.716049382708510540E-02, precision: 1e-15}
	exp_resasc = float{value: 4.416474291216854892E-02, precision: 1e-15}

	alpha = 2.6
	f = f1{alpha: alpha}

	prefix = "Smooth"
	testQK(f, exp_result, exp_abserr, exp_resabs, exp_resasc, qk, a, b, prefix, t)

	// Now test the basic rules with a positive function that has a
	// singularity. This should give large values of abserr which would
	// find discrepancies in the abserr calculation.
	exp_result = float{value: 2.449953612016972215E+01, precision: 1e-7}
	exp_abserr = float{value: 3.967771249391228849E+01, precision: 1e-5}
	exp_resabs = float{value: 2.449953612016972215E+01, precision: 1e-7}
	exp_resasc = float{value: 3.967771249391228849E+01, precision: 1e-7}

	alpha = -0.9
	f = f1{alpha: alpha}

	prefix = "Singular"
	testQK(f, exp_result, exp_abserr, exp_resabs, exp_resasc, qk, a, b, prefix, t)

	//Test the basic Gauss-Kronrod rules with a smooth oscillating
	// function, over an unsymmetric range. This should find any
	// discrepancies in the abscissae.
	exp_result = float{value: -7.238969575482961938E-01, precision: 1e-7}
	exp_abserr = float{value: 1.285290995039385778E-14, precision: 1e-15}
	exp_resabs = float{value: 1.157687209264406381E+00, precision: 1e-7}
	exp_resasc = float{value: 9.264666884071264263E-01, precision: 1e-7}

	alpha = 1.3
	f = f3{alpha: alpha}

	a = 0.3
	b = 2.71

	prefix = "Oscill"
	testQK(f, exp_result, exp_abserr, exp_resabs, exp_resasc, qk, a, b, prefix, t)
}

func TestQk61(t *testing.T) {
	var (
		exp_result float
		exp_abserr float
		exp_resabs float
		exp_resasc float
		alpha      float64
		a, b       float64
		f          F
		prefix     string
		qk         QkFunction
	)
	qk = Qk61
	a = 0.0
	b = 1.0

	//Test the basic Gauss-Kronrod rules with a smooth positive function.
	exp_result = float{value: 7.716049382713800753E-02, precision: 1e-15}
	exp_abserr = float{value: 1.566060362296155616E-12, precision: 1e-15}
	exp_resabs = float{value: 7.716049382713800753E-02, precision: 1e-15}
	exp_resasc = float{value: 4.419287685934316506E-02, precision: 1e-15}

	alpha = 2.6
	f = f1{alpha: alpha}

	prefix = "Smooth"
	testQK(f, exp_result, exp_abserr, exp_resabs, exp_resasc, qk, a, b, prefix, t)

	// Now test the basic rules with a positive function that has a
	// singularity. This should give large values of abserr which would
	// find discrepancies in the abserr calculation.
	exp_result = float{value: 2.583030240976628988E+01, precision: 1e-15}
	exp_abserr = float{value: 4.213750493076978643E+01, precision: 1e-7}
	exp_resabs = float{value: 2.583030240976628988E+01, precision: 1e-15}
	exp_resasc = float{value: 4.213750493076978643E+01, precision: 1e-7}

	alpha = -0.9
	f = f1{alpha: alpha}

	prefix = "Singular"
	testQK(f, exp_result, exp_abserr, exp_resabs, exp_resasc, qk, a, b, prefix, t)

	//Test the basic Gauss-Kronrod rules with a smooth oscillating
	// function, over an unsymmetric range. This should find any
	// discrepancies in the abscissae.
	exp_result = float{value: -7.238969575482959717E-01, precision: 1e-15}
	exp_abserr = float{value: 1.286438572027470736E-14, precision: 1e-15}
	exp_resabs = float{value: 1.158720854723590099E+00, precision: 1e-15}
	exp_resasc = float{value: 9.270469641771273972E-01, precision: 1e-15}

	alpha = 1.3
	f = f3{alpha: alpha}

	a = 0.3
	b = 2.71

	prefix = "Oscill"
	testQK(f, exp_result, exp_abserr, exp_resabs, exp_resasc, qk, a, b, prefix, t)
}
