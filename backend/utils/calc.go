package utils

import (
	"errors"
	"math"
)

// Cosine ref: https://github.com/gaspiman/cosine_similarity
func Cosine(a []float64, b []float64) (cosine float64, err error) {
	if len(a) != len(b) {
		return 0.0, errors.New("vector Length Different")
	}
	sum := 0.0
	s1 := 0.0
	s2 := 0.0
	for k := 0; k < len(a); k++ {
		sum += a[k] * b[k]
		s1 += math.Pow(a[k], 2)
		s2 += math.Pow(b[k], 2)
	}
	if s1 == 0 || s2 == 0 {
		return 0.0, errors.New("vectors should not be null (all zeros)")
	}
	return sum / (math.Sqrt(s1) * math.Sqrt(s2)), nil
}
