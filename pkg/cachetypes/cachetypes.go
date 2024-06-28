package cachetypes

import (
	"cache/pkg/cache"
	"reflect"
)

type Int int

func (i Int) Hash() uint64 {
	return cache.HashInt64(int64(i))
}

type Int8 int

func (i Int8) Hash() uint64 {
	return cache.HashInt64(int64(i))
}

type Int16 int

func (i Int16) Hash() uint64 {
	return cache.HashInt64(int64(i))
}

type Int32 int

func (i Int32) Hash() uint64 {
	return cache.HashInt64(int64(i))
}

type Int64 int

func (i Int64) Hash() uint64 {
	return cache.HashInt64(int64(i))
}

type String string

func (s String) Hash() uint64 {
	return cache.HashStr(string(s))
}

type User struct {
	ID    Int
	Name  String
	Roles [4]string
}

func (u User) Hash() uint64 {
	var hash uint64
	fields := reflect.VisibleFields(reflect.TypeOf(u))

	for _, f := range fields {
		_ = f
	}

	value := reflect.ValueOf(u)
	for i := 0; i < value.NumField(); i++ {
		hash += structHash(value.FieldByIndex([]int{i}))
	}

	return hash
}

func structHash(v reflect.Value) uint64 {
	var hash uint64

	t := v.Kind()

	switch t {
	case reflect.Int:
		hash = cache.HashInt64(v.Int())
	case reflect.String:
		hash = cache.HashStr(v.String())
	case reflect.Array:
		for i := 0; i < v.Len(); i++ {
			hash += structHash(v.Index(i))
		}
	}

	return hash
}
