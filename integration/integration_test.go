package integration

import "math"

type f1 struct {
	alpha float64
}

func (f f1) Evaluate(x float64) float64 {
	return math.Pow(x, f.alpha) * math.Log(1.0/x)
}

type f3 struct {
	alpha float64
}

func (f f3) Evaluate(x float64) float64 {
	return math.Cos(math.Pow(2.0, f.alpha) * math.Sin(x))
}
