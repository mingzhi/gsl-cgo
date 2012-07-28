package specfunc

import "testing"
import "bitbucket.org/mingzhi/goutils/assert"

func TestLngamma(t *testing.T) {
	var (
		x float64
		r float64
		e float64
		d float64
	)

	// case 1
	x = -0.1
	r = Lngamma(x)
	e = 2.368961332728788655
	d = 1e-6
	if !assert.EqualFloat64(r, e, d, 1) {
		t.Errorf("Lngamma: result %.5f, expected %.5f\n", r, e)
	}

	// case 2
	x = -1.0 / 256.0
	r = Lngamma(x)
	e = 5.547444766967471595
	d = 1e-6
	if !assert.EqualFloat64(r, e, d, 1) {
		t.Errorf("Lngamma: result %.5f, expected %.5f\n", r, e)
	}

	// case 3
	x = 1.0e-08
	r = Lngamma(x)
	e = 18.420680738180208905
	d = 1e-6
	if !assert.EqualFloat64(r, e, d, 1) {
		t.Errorf("Lngamma: result %.5f, expected %.5f\n", r, e)
	}

	// case 4
	x = 0.1
	r = Lngamma(x)
	e = 2.252712651734205
	d = 1e-6
	if !assert.EqualFloat64(r, e, d, 1) {
		t.Errorf("Lngamma: result %.5f, expected %.5f\n", r, e)
	}

	// case 5
	x = 100.0
	r = Lngamma(x)
	e = 359.1342053695753
	d = 1e-6
	if !assert.EqualFloat64(r, e, d, 1) {
		t.Errorf("Lngamma: result %.5f, expected %.5f\n", r, e)
	}

	// case 6
	x = -100 - 1.0/65536.0
	r = Lngamma(x)
	e = -352.6490910117097874
	d = 1e-6
	if !assert.EqualFloat64(r, e, d, 1) {
		t.Errorf("Lngamma: result %.5f, expected %.5f\n", r, e)
	}
}

func TestGamma(t *testing.T) {
	var (
		x float64
		r float64
		e float64
		d float64
	)

	x = 1.0 + 1.0/4096.0
	r = Gamma(x)
	e = 0.9998591371459403421
	d = 1e-6
	if !assert.EqualFloat64(r, e, d, 1) {
		t.Errorf("Gamma: result %.5f, expected %.5f\n", r, e)
	}

	x = 1.0 + 1.0/32.0
	r = Gamma(x)
	e = 0.9829010992836269148
	d = 1e-6
	if !assert.EqualFloat64(r, e, d, 1) {
		t.Errorf("Gamma: result %.5f, expected %.5f\n", r, e)
	}

	x = 2.0 + 1.0/256.0
	r = Gamma(x)
	e = 1.0016577903733583299
	d = 1e-6
	if !assert.EqualFloat64(r, e, d, 1) {
		t.Errorf("Gamma: result %.5f, expected %.5f\n", r, e)
	}

	x = 9.0
	r = Gamma(x)
	e = 40320.0
	d = 1e-6
	if !assert.EqualFloat64(r, e, d, 1) {
		t.Errorf("Gamma: result %.5f, expected %.5f\n", r, e)
	}

	x = 10.0
	r = Gamma(x)
	e = 362880.0
	d = 1e-6
	if !assert.EqualFloat64(r, e, d, 1) {
		t.Errorf("Gamma: result %.5f, expected %.5f\n", r, e)
	}

	x = 100.0
	r = Gamma(x)
	e = 9.332621544394415268e+155
	d = 1e-6
	if !assert.EqualFloat64(r, e, d, 1) {
		t.Errorf("Gamma: result %.5f, expected %.5f\n", r, e)
	}

	x = 170.0
	r = Gamma(x)
	e = 4.269068009004705275e+304
	d = 1e-6
	if !assert.EqualFloat64(r, e, d, 1) {
		t.Errorf("Gamma: result %.5f, expected %.5f\n", r, e)
	}

	x = 171.0
	r = Gamma(x)
	e = 7.257415615307998967e+306
	d = 1e-6
	if !assert.EqualFloat64(r, e, d, 1) {
		t.Errorf("Gamma: result %.5f, expected %.5f\n", r, e)
	}

	x = -10.5
	r = Gamma(x)
	e = -2.640121820547716316e-07
	d = 1e-6
	if !assert.EqualFloat64(r, e, d, 1) {
		t.Errorf("Gamma: result %.5f, expected %.5f\n", r, e)
	}

	x = -11.25
	r = Gamma(x)
	e = 6.027393816261931672e-08
	d = 1e-6
	if !assert.EqualFloat64(r, e, d, 1) {
		t.Errorf("Gamma: result %.5f, expected %.5f\n", r, e)
	}

	x = -1.0 + 1.0/65536.0
	r = Gamma(x)
	e = -65536.42280587818970
	d = 1e-6
	if !assert.EqualFloat64(r, e, d, 1) {
		t.Errorf("Gamma: result %.5f, expected %.5f\n", r, e)
	}

}
