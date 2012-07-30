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
		case 15:
		gsl_integration_qk15(&f, a, b, &result, &abserr, &resabs, &resasc);
		break;
		case 21:
		gsl_integration_qk21(&f, a, b, &result, &abserr, &resabs, &resasc);
		break;
		case 31:
		gsl_integration_qk31(&f, a, b, &result, &abserr, &resabs, &resasc);
		break;
		case 41:
		gsl_integration_qk41(&f, a, b, &result, &abserr, &resabs, &resasc);
		break;
		case 51:
		gsl_integration_qk51(&f, a, b, &result, &abserr, &resabs, &resasc);
		break;
		case 61:
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