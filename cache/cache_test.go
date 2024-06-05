package cache_test

import (
	"cache/cache"
	"testing"
)

type Int int

func (i Int) Hash() uint64 {
	return uint64(i)
}

func TestCache(t *testing.T) {
	c1 := cache.New[Int, any](cache.BucketsCount)
	var key Int = 1
	want := "first"
	err := c1.Set(key, want)
	if err != nil {
		t.Errorf("Set(%v,%v) = %v, want %v", key, want, err, nil)
	}

	got, err := c1.Get(key)
	if err != nil {
		t.Errorf("Get(%v) = %v, want %v, error=%v", key, got, want, err)
	}

	err = c1.Delete(key)
	if err != nil {
		t.Errorf("Delete(%v) error=%v", key, err)
	}

	_, err = c1.Get(key)
	if err == nil {
		t.Errorf("Get(%v) = %v, want error=%v", key, err, cache.ErrNotFound)
	}
}
