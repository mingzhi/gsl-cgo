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
