#include <gsl/gsl_math.h>
#include <gsl/gsl_integration.h>
gsl_function wrapper_function(void * p);
struct my_result_struct {
	double result;
	double abserr;
	int status;
	size_t neval;
};
typedef struct my_result_struct my_result;