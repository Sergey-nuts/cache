package cachetypes

import (
	"cache/pkg/cache"
	"testing"
)

type Hashable interface {
	comparable
	Hash() uint64
}

type caseHash[T Hashable] []*struct {
	key  T
	want uint64
}

func testHash[T cache.Hashable](cases caseHash[T]) func(*testing.T) {
	return func(t *testing.T) {
		t.Helper()

		for _, tt := range cases {
			// if got :=  cache.Hash(tt.key); got != tt.want {}
			if got := tt.key.Hash(); got != tt.want {
				t.Errorf("hash() = %v, want %v", got, tt.want)
			}
		}
	}
}

func TestHash(t *testing.T) {
	t.Run("test hash int", testHash(caseHash[Int]{
		{key: 23, want: 23},
		{key: 0, want: 0},
	}))
	t.Run("test hash int8", testHash(caseHash[Int8]{
		{key: 23, want: 23},
		{key: 0, want: 0},
	}))
	t.Run("test hash int16", testHash(caseHash[Int16]{
		{key: -128, want: 128},
		{key: 0, want: 0},
	}))
	t.Run("test hash int32", testHash(caseHash[Int32]{
		{key: 10000, want: 10000},
		{key: 0, want: 0},
	}))
	t.Run("test hash int64", testHash(caseHash[Int64]{
		{key: -23, want: 23},
		{key: 0, want: 0},
	}))
	t.Run("test hash string", testHash(caseHash[String]{
		{key: "2", want: 7147015373226909384},
		{key: "string", want: 2323256545573297332},
	}))
	t.Run("test hash User struct", testHash(caseHash[User]{
		{key: User{ID: 13, Name: "testUser"}, want: 9197180728666222400},
	}))
}
