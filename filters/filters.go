package filters

import (
	"bytes"
	"fmt"
	"path/filepath"
)

// Filter can be implemented to filter out unwanted data.
// from a slice of bytes.
type Filter interface {
	Filter([]byte) ([]byte, error)
}

// FilterFn is a convenience type for Filter.
type FilterFn func([]byte) ([]byte, error)

func (f FilterFn) Filter(b []byte) ([]byte, error) {
	if f == nil {
		return b, nil
	}
	return f(b)
}

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

func Noop() FilterFn {
	return func(b []byte) ([]byte, error) {
		return b, nil
	}
}

func replace(b []byte, s string, r string) []byte {

	// s = strings.TrimPrefix(s, string(filepath.Separator))

	// fmt.Printf("TODO >> filters.go:48 s %[1]T %[1]v\n", s)

	r = fmt.Sprintf("$%s%s", r, string(filepath.Separator))
	// fmt.Printf("TODO >> filters.go:50 r %[1]T %[1]v\n", r)

	// fmt.Printf("TODO >> filters.go:53 string(b) %[1]T %[1]v\n", string(b))
	return bytes.ReplaceAll(b, []byte(s), []byte(r))
}
