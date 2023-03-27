package utils

import "golang.org/x/exp/constraints"

func SliceMin[T constraints.Ordered](slice []T) (index int, m T) {
	for i, e := range slice {
		if i == 0 || e < m {
			m = e
			index = i
		}
	}
	return
}
