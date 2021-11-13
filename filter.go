package hepa

import (
	"github.com/markbates/hepa/filters"
)

// Filter can be implemented to filter out unwanted data.
// from a slice of bytes.
type Filter = filters.Filter

// FilterFn is a convenience type for Filter.
type FilterFn = filters.FilterFn

// func Rinse(p Purifier, s, r []byte) Purifier {
// 	return WithFunc(p, func(b []byte) ([]byte, error) {
// 		b = bytes.ReplaceAll(b, s, r)
// 		return b, nil
// 	})
// }

// func Clean(p Purifier, s []byte) Purifier {
// 	return WithFunc(p, func(b []byte) ([]byte, error) {
// 		if bytes.Contains(b, s) {
// 			return []byte{}, nil
// 		}
// 		return b, nil
// 	})
// }

var Noop = filters.Noop
