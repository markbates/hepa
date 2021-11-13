package filters

import (
	"bytes"
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

func replace(b []byte, find string, replace string) []byte {
	if len(find) == 0 || len(replace) == 0 {
		return b
	}

	return bytes.ReplaceAll(b, []byte(find), []byte(replace))
}
