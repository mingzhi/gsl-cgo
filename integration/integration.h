#include <gsl/gsl_math.h>
#include <gsl/gsl_integration.h>

struct my_result_struct {
	double result;
	double abserr;
	double resabs;
	double resasc;
	int status;
	size_t neval;
};
typedef struct my_result_struct my_result;

gsl_function wrapper_function(void * p);

my_result qng(const gsl_function f, double a, double b, double epsabs, double epsrel);

my_result qk(const gsl_function f, int rule, double a, double b);

//my_result qcheb(const gsl_function  f, double a, double b);

my_result qag(const gsl_function f, 
	double a, double b, 
	double epsabs, double epsrel, 
	size_t limit, int key);

my_result qags(const gsl_function f, 
	double a, double b, 
	double epsabs, double epsrel, 
	size_t limit);