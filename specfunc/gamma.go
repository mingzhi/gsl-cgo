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
package specfunc

//#cgo pkg-config: gsl
//#include <gsl/gsl_sf_gamma.h>
import "C"

/* Log[Gamma(x)], x not a negative integer
 * Uses real Lanczos method.
 * Returns the real part of Log[Gamma[x]] when x < 0,
 * i.e. Log[|Gamma[x]|].
 */
func LnGamma(x float64) (r float64) {
	r = float64(C.gsl_sf_lngamma(C.double(x)))
	return
}

/* Gamma(x), x not a negative integer
 * Uses real Lanczos method.
 */
func Gamma(x float64) (r float64) {
	r = float64(C.gsl_sf_gamma(C.double(x)))
	return
}

/* Regulated Gamma Function, x > 0
 * Gamma^*(x) = Gamma(x)/(Sqrt[2Pi] x^(x-1/2) exp(-x))
 *            = (1 + 1/(12x) + ...),  x->Inf
 */
func GammaStar(x float64) (r float64) {
	r = float64(C.gsl_sf_gammastar(C.double(x)))
	return
}

/* 1/Gamma(x)
 * Uses real Lanczos method.
 */
func GammaInv(x float64) (r float64) {
	//r = float64(C.gsl_sf_gammainv(C.double(x)))
	r = 1.0 / Gamma(x)
	return
}

/* x^n / n!
 *
 * x >= 0.0, n >= 0
 */
func TaylorCoeff(n int, x float64) (r float64) {
	r = float64(C.gsl_sf_taylorcoeff(C.int(n), C.double(x)))
	return r
}

/* n!
 *
 */
func Fact(n uint32) (r float64) {
	r = float64(C.gsl_sf_fact(C.uint(n)))
	return r
}

/* n!! = n(n-2)(n-4) ... 
 *
 */
func DoubleFact(n uint32) (r float64) {
	r = float64(C.gsl_sf_doublefact(C.uint(n)))
	return r
}

/* log(n!) 
 * Faster than ln(Gamma(n+1)) for n < 170; defers for larger n.
 */
func LnFact(n uint32) (r float64) {
	r = float64(C.gsl_sf_lnfact(C.uint(n)))
	return r
}

/* log(n!!) 
 */
func LnDoubleFact(n uint32) (r float64) {
	r = float64(C.gsl_sf_lndoublefact(C.uint(n)))
	return r
}

/* log(n choose m)
 */
func LnChoose(n, m uint32) (r float64) {
	r = float64(C.gsl_sf_lnchoose(C.uint(n), C.uint(m)))
	return
}

/* n choose m
 */
func Choose(n, m uint32) (r float64) {
	r = float64(C.gsl_sf_choose(C.uint(n), C.uint(m)))
	return
}

/* Logarithm of Pochhammer (Apell) symbol
 *   log( (a)_x )
 *   where (a)_x := Gamma[a + x]/Gamma[a]
 *
 * a > 0, a+x > 0
 *
 */
func LnPoch(a, x float64) (r float64) {
	r = float64(C.gsl_sf_lnpoch(C.double(a), C.double(x)))
	return
}

/* Pochhammer (Apell) symbol
 *   (a)_x := Gamma[a + x]/Gamma[x]
 *
 * a != neg integer, a+x != neg integer
 *
 */
func Poch(a, x float64) (r float64) {
	r = float64(C.gsl_sf_poch(C.double(a), C.double(x)))
	return
}

/* Relative Pochhammer (Apell) symbol
 *   ((a,x) - 1)/x
 *   where (a,x) = (a)_x := Gamma[a + x]/Gamma[a]
 *
 */
func PochRel(a, x float64) (r float64) {
	r = float64(C.gsl_sf_pochrel(C.double(a), C.double(x)))
	return
}

/* Normalized Incomplete Gamma Function
 *
 * Q(a,x) = 1/Gamma(a) Integral[ t^(a-1) e^(-t), {t,x,Infinity} ]
 *
 * a >= 0, x >= 0
 *   Q(a,0) := 1
 *   Q(0,x) := 0, x != 0
 *
 */
func GammaIncQ(a, x float64) (r float64) {
	r = float64(C.gsl_sf_gamma_inc_Q(C.double(a), C.double(x)))
	return
}

/* Complementary Normalized Incomplete Gamma Function
 *
 * P(a,x) = 1/Gamma(a) Integral[ t^(a-1) e^(-t), {t,0,x} ]
 *
 * a > 0, x >= 0
 *
 * exceptions: GSL_EDOM
 */
func GammaIncP(a, x float64) (r float64) {
	r = float64(C.gsl_sf_gamma_inc_P(C.double(a), C.double(x)))
	return
}

/* Non-normalized Incomplete Gamma Function
 *
 * Gamma(a,x) := Integral[ t^(a-1) e^(-t), {t,x,Infinity} ]
 *
 * x >= 0.0
 *   Gamma(a, 0) := Gamma(a)
 *
 */
func GammaInc(a, x float64) (r float64) {
	r = float64(C.gsl_sf_gamma_inc(C.double(a), C.double(x)))
	return
}

/* Logarithm of Beta Function
 * Log[B(a,b)]
 *
 * a > 0, b > 0
 */
func LnBeta(a, b float64) (r float64) {
	r = float64(C.gsl_sf_lnbeta(C.double(a), C.double(b)))
	return
}

/* Beta Function
 * B(a,b)
 *
 * a > 0, b > 0
 * exceptions: GSL_EDOM, GSL_EOVRFLW, GSL_EUNDRFLW
 */
func Beta(a, b float64) (r float64) {
	r = float64(C.gsl_sf_beta(C.double(a), C.double(b)))
	return
}

/* Normalized Incomplete Beta Function
 * B_x(a,b)/B(a,b)
 *
 * a > 0, b > 0, 0 <= x <= 1
 * exceptions: GSL_EDOM, GSL_EUNDRFLW
 */
func BetaInc(a, b, x float64) (r float64) {
	r = float64(C.gsl_sf_beta_inc(C.double(a), C.double(b), C.double(x)))
	return
}
