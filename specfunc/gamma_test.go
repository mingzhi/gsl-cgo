package specfunc

import "testing"
import "bitbucket.org/mingzhi/goutils/assert"

func TestLnGamma(t *testing.T) {
	var (
		x float64
		r float64
		e float64
		d float64
	)

	// case 1
	x = -0.1
	r = LnGamma(x)
	e = 2.368961332728788655
	d = 1e-6
	if !assert.EqualFloat64(r, e, d, 1) {
		t.Errorf("LnGamma: result %.5f, expected %.5f\n", r, e)
	}

	// case 2
	x = -1.0 / 256.0
	r = LnGamma(x)
	e = 5.547444766967471595
	d = 1e-6
	if !assert.EqualFloat64(r, e, d, 1) {
		t.Errorf("LnGamma: result %.5f, expected %.5f\n", r, e)
	}

	// case 3
	x = 1.0e-08
	r = LnGamma(x)
	e = 18.420680738180208905
	d = 1e-6
	if !assert.EqualFloat64(r, e, d, 1) {
		t.Errorf("LnGamma: result %.5f, expected %.5f\n", r, e)
	}

	// case 4
	x = 0.1
	r = LnGamma(x)
	e = 2.252712651734205
	d = 1e-6
	if !assert.EqualFloat64(r, e, d, 1) {
		t.Errorf("LnGamma: result %.5f, expected %.5f\n", r, e)
	}

	// case 5
	x = 100.0
	r = LnGamma(x)
	e = 359.1342053695753
	d = 1e-6
	if !assert.EqualFloat64(r, e, d, 1) {
		t.Errorf("LnGamma: result %.5f, expected %.5f\n", r, e)
	}

	// case 6
	x = -100 - 1.0/65536.0
	r = LnGamma(x)
	e = -352.6490910117097874
	d = 1e-6
	if !assert.EqualFloat64(r, e, d, 1) {
		t.Errorf("LnGamma: result %.5f, expected %.5f\n", r, e)
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

func TestGammaStar(t *testing.T) {
	var (
		x float64
		e float64
		r float64
		d float64
	)
	x, e = 1.0e-08, 3989.423555759890865
	r = GammaStar(x)
	d = 1e-6
	if !assert.EqualFloat64(r, e, d, 1) {
		t.Errorf("GammaStar: result %.5f, expected %.5f\n", r, e)
	}

	x, e = 1.0e-05, 126.17168469882690233
	r = GammaStar(x)
	d = 1e-6
	if !assert.EqualFloat64(r, e, d, 1) {
		t.Errorf("GammaStar: result %.5f, expected %.5f\n", r, e)
	}

	x, e = 0.001, 12.708492464364073506
	r = GammaStar(x)
	d = 1e-6
	if !assert.EqualFloat64(r, e, d, 1) {
		t.Errorf("GammaStar: result %.5f, expected %.5f\n", r, e)
	}
	x, e = 1.5, 1.0563442442685598666
	r = GammaStar(x)
	d = 1e-6
	if !assert.EqualFloat64(r, e, d, 1) {
		t.Errorf("GammaStar: result %.5f, expected %.5f\n", r, e)
	}

	x, e = 3.0, 1.0280645179187893045
	r = GammaStar(x)
	d = 1e-6
	if !assert.EqualFloat64(r, e, d, 1) {
		t.Errorf("GammaStar: result %.5f, expected %.5f\n", r, e)
	}

	x, e = 9.0, 1.0092984264218189715
	r = GammaStar(x)
	d = 1e-6
	if !assert.EqualFloat64(r, e, d, 1) {
		t.Errorf("GammaStar: result %.5f, expected %.5f\n", r, e)
	}

	x, e = 11.0, 1.0076024283104962850
	r = GammaStar(x)
	d = 1e-6
	if !assert.EqualFloat64(r, e, d, 1) {
		t.Errorf("GammaStar: result %.5f, expected %.5f\n", r, e)
	}

	x, e = 100.0, 1.0008336778720121418
	r = GammaStar(x)
	d = 1e-6
	if !assert.EqualFloat64(r, e, d, 1) {
		t.Errorf("GammaStar: result %.5f, expected %.5f\n", r, e)
	}

	x, e = 1.0e+05, 1.0000008333336805529
	r = GammaStar(x)
	d = 1e-6
	if !assert.EqualFloat64(r, e, d, 1) {
		t.Errorf("GammaStar: result %.5f, expected %.5f\n", r, e)
	}

	x, e = 1.0e+20, 1.0
	r = GammaStar(x)
	d = 1e-6
	if !assert.EqualFloat64(r, e, d, 1) {
		t.Errorf("GammaStar: result %.5f, expected %.5f\n", r, e)
	}
}

func TestGammaInv(t *testing.T) {
	var (
		x float64
		e float64
		r float64
		d float64
	)
	x, e = 1.0, 1.0
	r, d = GammaInv(x), 1e-6
	if !assert.EqualFloat64(r, e, d, 1) {
		t.Errorf("GammaInv: result %.5f, expected %.5f\n", r, e)
	}

	x, e = 2.0, 1.0
	r, d = GammaInv(x), 1e-6
	if !assert.EqualFloat64(r, e, d, 1) {
		t.Errorf("GammaInv: result %.5f, expected %.5f\n", r, e)
	}

	x, e = 3.0, 0.5
	r, d = GammaInv(x), 1e-6
	if !assert.EqualFloat64(r, e, d, 1) {
		t.Errorf("GammaInv: result %.5f, expected %.5f\n", r, e)
	}

	x, e = 4.0, 1.0/6.0
	r, d = GammaInv(x), 1e-6
	if !assert.EqualFloat64(r, e, d, 1) {
		t.Errorf("GammaInv: result %.5f, expected %.5f\n", r, e)
	}

	x, e = 10.0, 1.0/362880.0
	r, d = GammaInv(x), 1e-6
	if !assert.EqualFloat64(r, e, d, 1) {
		t.Errorf("GammaInv: result %.5f, expected %.5f\n", r, e)
	}

	x, e = 100.0, 1.0715102881254669232e-156
	r, d = GammaInv(x), 1e-6
	if !assert.EqualFloat64(r, e, d, 1) {
		t.Errorf("GammaInv: result %.5f, expected %.5f\n", r, e)
	}
	/*
		x, e = 0.0, 0.0
		r, d = GammaInv(x), 1e-6
		if !assert.EqualFloat64(r, e, d, 1) {
			t.Errorf("GammaInv: result %.5f, expected %.5f\n", r, e)
		}


		x, e = -1.0, 0.0
		r, d = GammaInv(x), 1e-6
		if !assert.EqualFloat64(r, e, d, 1) {
			t.Errorf("GammaInv: result %.5f, expected %.5f\n", r, e)
		}


		x, e = -2.0, 0.0
		r, d = GammaInv(x), 1e-6
		if !assert.EqualFloat64(r, e, d, 1) {
			t.Errorf("GammaInv: result %.5f, expected %.5f\n", r, e)
		}

		x, e = -3.0, 0.0
		r, d = GammaInv(x), 1e-6
		if !assert.EqualFloat64(r, e, d, 1) {
			t.Errorf("GammaInv: result %.5f, expected %.5f\n", r, e)
		}

		x, e = -4.0, 0.0
		r, d = GammaInv(x), 1e-6
		if !assert.EqualFloat64(r, e, d, 1) {
			t.Errorf("GammaInv: result %.5f, expected %.5f\n", r, e)
		}

		x, e = -10.5, -1.0/2.640121820547716316e-07
		r, d = GammaInv(x), 1e-6
		if !assert.EqualFloat64(r, e, d, 1) {
			t.Errorf("GammaInv: result %.5f, expected %.5f\n", r, e)
		}

		x, e = -11.25, 1.0/6.027393816261931672e-08
		r, d = GammaInv(x), 1e-6
		if !assert.EqualFloat64(r, e, d, 1) {
			t.Errorf("GammaInv: result %.5f, expected %.5f\n", r, e)
		}

		x, e = -1.0+1.0/65536.0, -1.0/65536.42280587818970
		r, d = GammaInv(x), 1e-6
		if !assert.EqualFloat64(r, e, d, 1) {
			t.Errorf("GammaInv: result %.5f, expected %.5f\n", r, e)
		}
	*/
}

func TestTaylorCoeff(t *testing.T) {
	var (
		x float64
		n int
		e float64
		d float64
		r float64
	)
	n, x, e = 10, 1.0/1048576.0, 1.7148961854776073928e-67
	r = TaylorCoeff(n, x)
	d = 1e-6
	if !assert.EqualFloat64(r, e, d, 1) {
		t.Errorf("TaylorCoeff: result %.5f, expected %.5f\n", r, e)
	}
	n, x, e = 10, 1.0/1024.0, 2.1738891788497900281e-37
	r = TaylorCoeff(n, x)
	d = 1e-6
	if !assert.EqualFloat64(r, e, d, 1) {
		t.Errorf("TaylorCoeff: result %.5f, expected %.5f\n", r, e)
	}
	n, x, e = 10, 1.0, 2.7557319223985890653e-07
	r = TaylorCoeff(n, x)
	d = 1e-6
	if !assert.EqualFloat64(r, e, d, 1) {
		t.Errorf("TaylorCoeff: result %.5f, expected %.5f\n", r, e)
	}
	n, x, e = 10, 5.0, 2.6911444554673721340
	r = TaylorCoeff(n, x)
	d = 1e-6
	if !assert.EqualFloat64(r, e, d, 1) {
		t.Errorf("TaylorCoeff: result %.5f, expected %.5f\n", r, e)
	}
	n, x, e = 10, 500.0, 2.6911444554673721340e+20
	r = TaylorCoeff(n, x)
	d = 1e-6
	if !assert.EqualFloat64(r, e, d, 1) {
		t.Errorf("TaylorCoeff: result %.5f, expected %.5f\n", r, e)
	}
	n, x, e = 100, 100.0, 1.0715102881254669232e+42
	r = TaylorCoeff(n, x)
	d = 1e-6
	if !assert.EqualFloat64(r, e, d, 1) {
		t.Errorf("TaylorCoeff: result %.5f, expected %.5f\n", r, e)
	}
	n, x, e = 1000, 200.0, 2.6628790558154746898e-267
	r = TaylorCoeff(n, x)
	d = 1e-6
	if !assert.EqualFloat64(r, e, d, 1) {
		t.Errorf("TaylorCoeff: result %.5f, expected %.5f\n", r, e)
	}
	n, x, e = 1000, 500.0, 2.3193170139740855074e+131
	r = TaylorCoeff(n, x)
	d = 1e-6
	if !assert.EqualFloat64(r, e, d, 1) {
		t.Errorf("TaylorCoeff: result %.5f, expected %.5f\n", r, e)
	}
}

func TestFact(t *testing.T) {
	var (
		n uint32
		e float64
		r float64
		d float64
	)
	n, e = 0, 1.0
	r, d = Fact(n), 1e-6
	if !assert.EqualFloat64(r, e, d, 1) {
		t.Errorf("Fact: result %.5f, expected %.5f\n", r, e)
	}

	n, e = 1, 1.0
	r, d = Fact(n), 1e-6
	if !assert.EqualFloat64(r, e, d, 1) {
		t.Errorf("Fact: result %.5f, expected %.5f\n", r, e)
	}

	n, e = 7, 5040.0
	r, d = Fact(n), 1e-6
	if !assert.EqualFloat64(r, e, d, 1) {
		t.Errorf("Fact: result %.5f, expected %.5f\n", r, e)
	}

	n, e = 33, 8.683317618811886496e+36
	r, d = Fact(n), 1e-6
	if !assert.EqualFloat64(r, e, d, 1) {
		t.Errorf("Fact: result %.5f, expected %.5f\n", r, e)
	}

}

func TestLnFact(t *testing.T) {
	var (
		n uint32
		e float64
		r float64
		d float64
	)
	n, e = 0, 0.0
	r, d = LnFact(n), 1e-6
	if !assert.EqualFloat64(r, e, d, 1) {
		t.Errorf("LnFact: result %.5f, expected %.5f\n", r, e)
	}

	n, e = 1, 0.0
	r, d = LnFact(n), 1e-6
	if !assert.EqualFloat64(r, e, d, 1) {
		t.Errorf("LnFact: result %.5f, expected %.5f\n", r, e)
	}

	n, e = 7, 8.525161361065414300
	r, d = LnFact(n), 1e-6
	if !assert.EqualFloat64(r, e, d, 1) {
		t.Errorf("LnFact: result %.5f, expected %.5f\n", r, e)
	}

	n, e = 33, 85.05446701758151741
	r, d = LnFact(n), 1e-6
	if !assert.EqualFloat64(r, e, d, 1) {
		t.Errorf("LnFact: result %.5f, expected %.5f\n", r, e)
	}

}

func TestLnDoubleFact(t *testing.T) {
	var (
		n uint32
		e float64
		r float64
		d float64
	)
	n, e = 0, 0.0
	r, d = LnDoubleFact(n), 1e-6
	if !assert.EqualFloat64(r, e, d, 1) {
		t.Errorf("LnDoubleFact: result %.5f, expected %.5f\n", r, e)
	}

	n, e = 7, 4.653960350157523371
	r, d = LnDoubleFact(n), 1e-6
	if !assert.EqualFloat64(r, e, d, 1) {
		t.Errorf("LnDoubleFact: result %.5f, expected %.5f\n", r, e)
	}

	n, e = 33, 43.292252022541719660
	r, d = LnDoubleFact(n), 1e-6
	if !assert.EqualFloat64(r, e, d, 1) {
		t.Errorf("LnDoubleFact: result %.5f, expected %.5f\n", r, e)
	}

	n, e = 34, 45.288575519655959140
	r, d = LnDoubleFact(n), 1e-6
	if !assert.EqualFloat64(r, e, d, 1) {
		t.Errorf("LnDoubleFact: result %.5f, expected %.5f\n", r, e)
	}

	n, e = 1034, 3075.6383796271197707
	r, d = LnDoubleFact(n), 1e-6
	if !assert.EqualFloat64(r, e, d, 1) {
		t.Errorf("LnDoubleFact: result %.5f, expected %.5f\n", r, e)
	}

	n, e = 1035, 3078.8839081731809169
	r, d = LnDoubleFact(n), 1e-6
	if !assert.EqualFloat64(r, e, d, 1) {
		t.Errorf("LnDoubleFact: result %.5f, expected %.5f\n", r, e)
	}

}

func TestLnChoose(t *testing.T) {
	var (
		n, m    uint32
		r, e, d float64
	)
	n, m, e = 7, 3, 3.555348061489413680
	r, d = LnChoose(n, m), 1e-6
	if !assert.EqualFloat64(r, e, d, 1) {
		t.Errorf("LnChoose: result %.5f, expected %.5f\n", r, e)
	}

	n, m, e = 5, 2, 2.302585092994045684
	r, d = LnChoose(n, m), 1e-6
	if !assert.EqualFloat64(r, e, d, 1) {
		t.Errorf("LnChoose: result %.5f, expected %.5f\n", r, e)
	}
}

func TestChoose(t *testing.T) {
	var (
		n, m    uint32
		r, e, d float64
	)

	n, m, e = 7, 3, 35
	r, d = Choose(n, m), 1e-6
	if !assert.EqualFloat64(r, e, d, 1) {
		t.Errorf("Choose: result %.5f, expected %.5f\n", r, e)
	}

	n, m, e = 7, 4, 35
	r, d = Choose(n, m), 1e-6
	if !assert.EqualFloat64(r, e, d, 1) {
		t.Errorf("Choose: result %.5f, expected %.5f\n", r, e)
	}

	n, m, e = 5, 2, 10
	r, d = Choose(n, m), 1e-6
	if !assert.EqualFloat64(r, e, d, 1) {
		t.Errorf("Choose: result %.5f, expected %.5f\n", r, e)
	}

	n, m, e = 5, 3, 10
	r, d = Choose(n, m), 1e-6
	if !assert.EqualFloat64(r, e, d, 1) {
		t.Errorf("Choose: result %.5f, expected %.5f\n", r, e)
	}

	n, m, e = 500, 495, 255244687600.0
	r, d = Choose(n, m), 1e-6
	if !assert.EqualFloat64(r, e, d, 1) {
		t.Errorf("Choose: result %.5f, expected %.5f\n", r, e)
	}

	n, m, e = 500, 5, 255244687600.0
	r, d = Choose(n, m), 1e-6
	if !assert.EqualFloat64(r, e, d, 1) {
		t.Errorf("Choose: result %.5f, expected %.5f\n", r, e)
	}

	n, m, e = 500, 200, 5.054949849935532221e+144
	r, d = Choose(n, m), 1e-6
	if !assert.EqualFloat64(r, e, d, 1) {
		t.Errorf("Choose: result %.5f, expected %.5f\n", r, e)
	}

	n, m, e = 500, 300, 5.054949849935532221e+144
	r, d = Choose(n, m), 1e-6
	if !assert.EqualFloat64(r, e, d, 1) {
		t.Errorf("Choose: result %.5f, expected %.5f\n", r, e)
	}
}
