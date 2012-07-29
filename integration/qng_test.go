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
