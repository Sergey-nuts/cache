package cache

import (
	"fmt"
	"sync"
)

var (
	ErrNotFound = fmt.Errorf("key not in cache")

	// default count backets for cache
	BucketsCount = 5
)

type Hashable interface {
	comparable
	Hash() uint64
}

type cache[K Hashable, V any] struct {
	// buckets slice
	bs []bucket[K, V]
}

type bucket[K Hashable, V any] struct {
	mu sync.RWMutex
	m  map[K]V
}

// New(BucketsCount) is constructor for cashe
//
// use BucketsCount for default value
func New[K Hashable, V any](n int) *cache[K, V] {
	sl := make([]bucket[K, V], n)

	for i := range n {
		sl[i] = newBucket[K, V]()
	}

	return &cache[K, V]{bs: sl}
}

func newBucket[K Hashable, V any]() bucket[K, V] {
	return bucket[K, V]{
		m: make(map[K]V),
	}
}

// Set adds pair key,value in cache
func (c *cache[K, V]) Set(key K, value V) error {
	hash := key.Hash()
	n := hash % uint64(len(c.bs))

	err := c.bs[n].set(key, value)
	if err != nil {
		return err
	}

	return nil
}

// Get return value by key from cache,
//
// return err if key not in cache
func (c *cache[K, V]) Get(key K) (V, error) {
	hash := key.Hash()
	n := hash % uint64(len(c.bs))

	value, err := c.bs[n].get(key)
	if err != nil {
		return *new(V), err
	}

	return value, nil
}

// Delete remove pair key,value from cache
func (c *cache[K, V]) Delete(key K) error {
	hash := key.Hash()
	n := hash % uint64(len(c.bs))

	err := c.bs[n].delete(key)
	if err != nil {
		return err
	}

	return nil
}

func (b *bucket[K, V]) set(key K, value V) error {
	b.mu.Lock()
	b.m[key] = value
	b.mu.Unlock()

	return nil
}

func (b *bucket[K, V]) get(key K) (V, error) {
	b.mu.RLock()
	v, ok := b.m[key]
	b.mu.RUnlock()
	if !ok {
		return *new(V), ErrNotFound
	}

	return v, nil
}

func (b *bucket[K, V]) delete(key K) error {
	b.mu.Lock()
	delete(b.m, key)
	b.mu.Unlock()

	return nil
}
