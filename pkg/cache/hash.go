package cache

import (
	"crypto/md5"
	"encoding/binary"
	"math"
)

// HashStr return hash of string key
//
// retturn 0 if string is empty
func HashStr(key string) uint64 {
	if key == "" {
		return 0
	}
	h := md5.New()

	h.Write([]byte(key))
	hashValue := h.Sum(nil)

	return binary.NativeEndian.Uint64(hashValue)
}

// HashInt64 return hash of int64 key
func HashInt64(key int64) uint64 {
	return uint64(math.Abs(float64(key)))
}

// HashFlaot64 return hash of float64 key
func HashFloat64(key float64) uint64 {
	return uint64((math.Abs(key) * 10000000.))
}
