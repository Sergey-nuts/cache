package cache_test

import (
	"cache/pkg/cache"
	"testing"
)

func TestHashStr(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			name: "emty string",
			args: args{key: ""},
			want: 0,
		},
		{
			name: "first string",
			args: args{key: "2"},
			want: 7147015373226909384,
		},
		{
			name: "second string",
			args: args{key: "string"},
			want: 2323256545573297332,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cache.HashStr(tt.args.key); got != tt.want {
				t.Errorf("HashStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashInt64(t *testing.T) {
	type args struct {
		key int64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			name: "first int case",
			args: args{key: 1},
			want: 1,
		},
		{
			name: "second int case",
			args: args{key: 123123123},
			want: 123123123,
		},
		{
			name: "third int case",
			args: args{key: -121},
			want: 121,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cache.HashInt64(tt.args.key); got != tt.want {
				t.Errorf("HashInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}
