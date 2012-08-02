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
package randist

/*
#cgo pkg-config: gsl
#include <gsl/gsl_randist.h>
#include <gsl/gsl_rng.h>
*/
import "C"

// Random Number Generator
type RNG struct {
	g *C.gsl_rng
}

func NewRNG(i int) *RNG {
	return &RNG{g: newRandomGenerator(i)}
}

func (r *RNG) SetSeed(seed int) {
	setSeed(r.g, C.ulong(seed))
}

// Free memory allocated to C.gsl_rng
func (r *RNG) Free() {
	freeRandomGenerator(r.g)
}

func newRandomGenerator(i int) *C.gsl_rng {
	var r *C.gsl_rng
	types := []*C.gsl_rng_type{
		C.gsl_rng_borosh13,
		C.gsl_rng_cmrg,
		C.gsl_rng_coveyou,
		C.gsl_rng_fishman18,
		C.gsl_rng_fishman20,
		C.gsl_rng_fishman2x,
		C.gsl_rng_gfsr4,
		C.gsl_rng_knuthran,
		C.gsl_rng_knuthran2,
		C.gsl_rng_knuthran2002,
		C.gsl_rng_lecuyer21,
		C.gsl_rng_minstd,
		C.gsl_rng_mrg,
		C.gsl_rng_mt19937,
		C.gsl_rng_mt19937_1999,
		C.gsl_rng_mt19937_1998,
		C.gsl_rng_r250,
		C.gsl_rng_ran0,
		C.gsl_rng_ran1,
		C.gsl_rng_ran2,
		C.gsl_rng_ran3,
		C.gsl_rng_rand,
		C.gsl_rng_rand48,
		C.gsl_rng_random128_bsd,
		C.gsl_rng_random128_glibc2,
		C.gsl_rng_random128_libc5,
		C.gsl_rng_random256_bsd,
		C.gsl_rng_random256_glibc2,
		C.gsl_rng_random256_libc5,
		C.gsl_rng_random32_bsd,
		C.gsl_rng_random32_glibc2,
		C.gsl_rng_random32_libc5,
		C.gsl_rng_random64_bsd,
		C.gsl_rng_random64_glibc2,
		C.gsl_rng_random64_libc5,
		C.gsl_rng_random8_bsd,
		C.gsl_rng_random8_glibc2,
		C.gsl_rng_random8_libc5,
		C.gsl_rng_random_bsd,
		C.gsl_rng_random_glibc2,
		C.gsl_rng_random_libc5,
		C.gsl_rng_randu,
		C.gsl_rng_ranf,
		C.gsl_rng_ranlux,
		C.gsl_rng_ranlux389,
		C.gsl_rng_ranlxd1,
		C.gsl_rng_ranlxd2,
		C.gsl_rng_ranlxs0,
		C.gsl_rng_ranlxs1,
		C.gsl_rng_ranlxs2,
		C.gsl_rng_ranmar,
		C.gsl_rng_slatec,
		C.gsl_rng_taus,
		C.gsl_rng_taus2,
		C.gsl_rng_taus113,
		C.gsl_rng_transputer,
		C.gsl_rng_tt800,
		C.gsl_rng_uni,
		C.gsl_rng_uni32,
		C.gsl_rng_vax,
		C.gsl_rng_waterman14,
		C.gsl_rng_zuf,
	}
	if i >= len(types) {
		panic("Can not find gsl rand generator!")
	}
	r = C.gsl_rng_alloc(types[i])
	return r
}

func freeRandomGenerator(r *C.gsl_rng) {
	C.gsl_rng_free(r)
}

func setSeed(r *C.gsl_rng, seed C.ulong) {
	C.gsl_rng_set(r, seed)
}
