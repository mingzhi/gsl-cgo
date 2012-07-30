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

#include <gsl/gsl_errno.h>
#include "integration.h"
#include "_cgo_export.h"

gsl_function
wrapper_function (void * p)
{
  gsl_function f_new;
  f_new.function = &fevaluate ;
  f_new.params = p ;
  return f_new;
}

my_result
qng(const gsl_function f, double a, double b, double epsabs, double epsrel) 
{
	double result = 0, abserr = 0;
	size_t neval;
	int status;
	my_result r;

	status = gsl_integration_qng(&f, a, b, epsabs, epsrel, &result, &abserr, &neval);
	r.result = result;
	r.abserr = abserr;
	r.status = status;
	r.neval = neval;

	return r;
}

my_result
qk(const gsl_function f, int rule, double a, double b)
{
	my_result r;

	double result = 0, abserr = 0, resabs = 0, resasc = 0;
	switch (rule) {
		case GSL_INTEG_GAUSS15:
		gsl_integration_qk15(&f, a, b, &result, &abserr, &resabs, &resasc);
		break;
		case GSL_INTEG_GAUSS21:
		gsl_integration_qk21(&f, a, b, &result, &abserr, &resabs, &resasc);
		break;
		case GSL_INTEG_GAUSS31:
		gsl_integration_qk31(&f, a, b, &result, &abserr, &resabs, &resasc);
		break;
		case GSL_INTEG_GAUSS41:
		gsl_integration_qk41(&f, a, b, &result, &abserr, &resabs, &resasc);
		break;
		case GSL_INTEG_GAUSS51:
		gsl_integration_qk51(&f, a, b, &result, &abserr, &resabs, &resasc);
		break;
		case GSL_INTEG_GAUSS61:
		gsl_integration_qk61(&f, a, b, &result, &abserr, &resabs, &resasc);
		break;
		default:
		gsl_integration_qk15(&f, a, b, &result, &abserr, &resabs, &resasc);
		break;
	}
	
	r.result = result;
	r.abserr = abserr;
	r.resabs = resabs;
	r.resasc = resasc;
	r.status = GSL_SUCCESS;

	return r;
}

my_result qag(const gsl_function f, 
	double a, double b, 
	double epsabs, double epsrel, 
	size_t limit, int key)
{
	my_result r;

	int status;
	double result = 0, abserr = 0;
	gsl_integration_workspace * w = gsl_integration_workspace_alloc(limit) ;
	status = gsl_integration_qag (&f, a, b, 
								  epsabs, epsrel, w->limit,
                                  key, w,
                                  &result, &abserr) ;
	r.result = result;
	r.abserr = abserr;
	r.status = status;

	gsl_integration_workspace_free(w);
	return r;
}

my_result qags(const gsl_function f, 
	double a, double b, 
	double epsabs, double epsrel, 
	size_t limit)
{
	my_result r;

	int status;
	double result = 0, abserr = 0;
	gsl_integration_workspace * w = gsl_integration_workspace_alloc(limit) ;
	status = gsl_integration_qags (&f, a, b, 
								  epsabs, epsrel, w->limit,
                                  w,
                                  &result, &abserr) ;
	r.result = result;
	r.abserr = abserr;
	r.status = status;

	gsl_integration_workspace_free(w);
	return r;
}
