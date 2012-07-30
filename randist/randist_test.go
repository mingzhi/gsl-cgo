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
package randist

import (
	"fmt"
	"github.com/mingzhi/gsl/integra"
	"math"
	"testing"
)

func TestExponential(t *testing.T) {
	var lambda float64
	var e *Exponential

	// [TODO] how to test FreeRandomGenerator()?

	lambda = 0.5
	e = NewExponential(lambda, RAND48)
	defer e.FreeRandomGenerator()
	ition := integration{cd: e, funcT: "pdf"}
	name := "ExponentialPdf"
	err := testPdf(ition, name)
	if err != nil {
		t.Errorf("%v\n", err)
	}
}

func TestPoisson(t *testing.T) {
	var lambda float64
	var e *Exponential

	// [TODO] how to test FreeRandomGenerator()?

	lambda = 0.5
	e = NewExponential(lambda, RAND48)
	defer e.FreeRandomGenerator()
	ition := integration{cd: e, funcT: "pdf"}
	name := "ExponentialPdf"
	err := testPdf(ition, name)
	if err != nil {
		t.Errorf("%v\n", err)
	}
}

/*
func TestGamma(t *testing.T) {
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	cd := NewGamma(1, 1, r)
	ition := integration{cd: cd, funcT: "pdf"}
	name := "GammaPdf"
	err := testPdf(ition, name)
	if err != nil {
		t.Errorf("%v\n", err)
	}
}
*/

const (
	N    = 100000
	BINS = 100
)

type integration struct {
	cd    ContinuousDistribution
	funcT string
}

func (i integration) Evaluate(x float64) (p float64) {
	switch i.funcT {
	case "cdf":
		p = i.cd.Cdf(x)

		break
	default:
		p = i.cd.Pdf(x)
	}
	return
}

func (i integration) RandFloat() (x float64) {
	return i.cd.RandomFloat64()
}

type TestError struct {
	message string
}

func (e TestError) Error() string {
	return e.message
}

func integrate(pdf integra.Function, a, b float64) float64 {
	n := 1000
	w, _ := integra.NewWorkspace(n)
	result, _, _ := integra.Qag(pdf, a, b, 1e-16, 1e-4, n, w, integra.Qk61)
	return result
}

func testPdf(rd integration, name string) (err error) {
	count := make([]float64, BINS)
	edge := make([]float64, BINS)
	p := make([]float64, BINS)

	a := -5.0
	b := +5.0
	dx := (b - a) / BINS
	var total, mean float64
	var attempts int
	n0 := 0
	n := N

	for i := 0; i < BINS; i++ {
		/* Compute the integral of p(x) from x to x+dx */

		x := a + float64(i)*dx

		if math.Abs(x) < 1e-10 {
			/* hit the origin exactly */
			x = 0.0
		}

		p[i] = integrate(rd, x, x+dx)
	}

trial:
	attempts++
	for i := n0; i < n; i++ {
		r := rd.RandFloat()
		total += r
		if a < r && r < b {
			u := (r - a) / dx
			bin, f := math.Modf(u)
			j := int(bin)
			if f == 0 {
				edge[j]++
			} else {
				count[j]++
			}
		}
	}

	/* Sort out where the hits on the edges should go */

	for i := 0; i < BINS; i++ {
		/* If the bin above is empty, its lower edge hits belong in the
		   lower bin */

		if i+1 < BINS && count[i+1] == 0 {
			count[i] += edge[i+1]
			edge[i+1] = 0
		}

		count[i] += edge[i]
		edge[i] = 0
	}

	mean = (total / float64(n))

	exception := math.IsInf(mean, 0) || math.IsNaN(mean)
	if exception {
		err = TestError{message: fmt.Sprintf("%s, finite mean, observed %g", name, mean)}
		return
	}
	exception_i := false
	for i := 0; i < BINS; i++ {
		x := a + float64(i)*dx
		d := math.Abs(count[i] - float64(n)*p[i])
		//d1 := math.Abs(count[i] - float64(n)*(rd.cd.Cdf(x+dx)-rd.cd.Cdf(x)))
		if math.IsNaN(p[i]) || math.IsInf(p[i], 0) {
			exception_i = true
		} else if p[i] != 0 {
			s := d / math.Sqrt(float64(n)*p[i])
			//s1 := d1 / math.Sqrt(float64(n)*p[i])
			exception_i = (s > 5 && d > 2) //|| (s1 > 5 && d1 > 2)
		} else {
			exception_i = (count[i] != 0)
		}

		/* Extend the sample if there is an outlier on the first attempt
		   to avoid spurious failures when running large numbers of tests. */
		if exception_i && attempts < 1 {
			n0 = n
			n = 2.0 * n
			goto trial
		}

		exception = exception && exception_i
		if exception_i {
			err = TestError{message: fmt.Sprintf("%s [%g,%g) (%g/%d=%g observed vs %g expected)",
				name, x, x+dx, count[i], n, count[i]/float64(n), p[i])}
			return
		}
	}

	if exception {
		err = TestError{message: fmt.Sprintf("%s, sampling against pdf over range [%g,%g) ",
			name, a, b)}
	}
	return
}
