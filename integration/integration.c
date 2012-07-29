#include <gsl/gsl_math.h>
#include <gsl/gsl_integration.h>
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