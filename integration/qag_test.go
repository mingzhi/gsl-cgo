package integration

import "testing"
import "bitbucket.org/mingzhi/goutils/assert"

func TestQag(t *testing.T) {
	{

		exp_result := 7.716049382715854665E-02
		exp_abserr := 6.679384885865053037E-12

		alpha := 2.6
		f := f1{alpha: alpha}
		result, abserr, status := Qag(f, 0.0, 1.0, 0.0, 1e-10, 1000, GSL_INTEG_GAUSS15)
		if !assert.EqualFloat64(result, exp_result, 1e-15, 0) {
			t.Errorf("Smooth: result %.10f, expected %.10f\n", result, exp_result)
		}
		if !assert.EqualFloat64(abserr, exp_abserr, 1e-6, 0) {
			t.Errorf("Smooth: abserr %f, expected %f\n", abserr, exp_abserr)
		}
		if status != 0 {
			t.Errorf("Smooth: do not expected error %v\n", status)
		}

		result, abserr, status = Qag(f, 1.0, 0.0, 0.0, 1e-10, 1000, GSL_INTEG_GAUSS15)
		if !assert.EqualFloat64(result, -exp_result, 1e-15, 0) {
			t.Errorf("Smooth: result %.10f, expected %.10f\n", result, -exp_result)
		}
		if !assert.EqualFloat64(abserr, exp_abserr, 1e-6, 0) {
			t.Errorf("Smooth: abserr %f, expected %f\n", abserr, exp_abserr)
		}
		if status != 0 {
			t.Errorf("Smooth: do not expected error %v\n", status)
		}
	}

	/* Test the same function using an absolute error bound and the
	   21-point rule */

	{
		exp_result := 7.716049382716050342E-02
		exp_abserr := 2.227969521869139532E-15

		alpha := 2.6
		f := f1{alpha: alpha}
		result, abserr, status := Qag(f, 0.0, 1.0, 1e-14, 0.0, 1000, GSL_INTEG_GAUSS21)
		if !assert.EqualFloat64(result, exp_result, 1e-15, 0) {
			t.Errorf("Smooth: result %.10f, expected %.10f\n", result, exp_result)
		}
		if !assert.EqualFloat64(abserr, exp_abserr, 1e-6, 0) {
			t.Errorf("Smooth: abserr %f, expected %f\n", abserr, exp_abserr)
		}
		if status != 0 {
			t.Errorf("Smooth: do not expected error %v\n", status)
		}

		result, abserr, status = Qag(f, 1.0, 0.0, 1e-14, 0.0, 1000, GSL_INTEG_GAUSS21)
		if !assert.EqualFloat64(result, -exp_result, 1e-15, 0) {
			t.Errorf("Smooth: result %.10f, expected %.10f\n", result, -exp_result)
		}
		if !assert.EqualFloat64(abserr, exp_abserr, 1e-6, 0) {
			t.Errorf("Smooth: abserr %f, expected %f\n", abserr, exp_abserr)
		}
		if status != 0 {
			t.Errorf("Smooth: do not expected error %v\n", status)
		}
	}

	/* Adaptive integration of an oscillatory function which terminates because
	   of roundoff error, uses the 31-pt rule 
	{
		exp_result := -7.238969575482959717E-01
		exp_abserr := 1.285805464427459261E-14

		alpha := 1.3
		f := f3{alpha: alpha}
		result, abserr, status := Qag(f, 0.3, 2.71, 1e-14, 0.0, 1000, GSL_INTEG_GAUSS31)
		if !assert.EqualFloat64(result, exp_result, 1e-15, 0) {
			t.Errorf("Oscill: result %.10f, expected %.10f\n", result, exp_result)
		}
		if !assert.EqualFloat64(abserr, exp_abserr, 1e-6, 0) {
			t.Errorf("Oscill: abserr %f, expected %f\n", abserr, exp_abserr)
		}
		if status != 0 {
			t.Errorf("Oscill: do not expected error %v\n", status)
		}

		result, abserr, status = Qag(f, 2.71, 0.3, 1e-14, 0.0, 1000, GSL_INTEG_GAUSS31)
		if !assert.EqualFloat64(result, -exp_result, 1e-15, 0) {
			t.Errorf("Oscill: result %.10f, expected %.10f\n", result, -exp_result)
		}
		if !assert.EqualFloat64(abserr, exp_abserr, 1e-6, 0) {
			t.Errorf("Oscill: abserr %f, expected %f\n", abserr, exp_abserr)
		}
		if status != 0 {
			t.Errorf("Oscill: do not expected error %v\n", status)
		}
	}
	*/

}

func TestQags(t *testing.T) {
	{

		exp_result := 7.716049382715789440E-02
		exp_abserr := 2.216394961010438404E-12

		alpha := 2.6
		f := f1{alpha: alpha}
		result, abserr, status := Qags(f, 0.0, 1.0, 0.0, 1e-10, 1000)
		if !assert.EqualFloat64(result, exp_result, 1e-15, 0) {
			t.Errorf("Smooth: result %.10f, expected %.10f\n", result, exp_result)
		}
		if !assert.EqualFloat64(abserr, exp_abserr, 1e-6, 0) {
			t.Errorf("Smooth: abserr %f, expected %f\n", abserr, exp_abserr)
		}
		if status != 0 {
			t.Errorf("Smooth: do not expected error %v\n", status)
		}

		result, abserr, status = Qags(f, 1.0, 0.0, 0.0, 1e-10, 1000)
		if !assert.EqualFloat64(result, -exp_result, 1e-15, 0) {
			t.Errorf("Smooth: result %.10f, expected %.10f\n", result, -exp_result)
		}
		if !assert.EqualFloat64(abserr, exp_abserr, 1e-6, 0) {
			t.Errorf("Smooth: abserr %f, expected %f\n", abserr, exp_abserr)
		}
		if status != 0 {
			t.Errorf("Smooth: do not expected error %v\n", status)
		}
	}

}
