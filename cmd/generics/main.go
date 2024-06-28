package main

import (
	"cache/pkg/cache"
	"cache/pkg/cachetypes"
	"fmt"
	"log"
)

func main() {
	cache1 := cache.New[cachetypes.String, any](cache.BucketsCount)
	cache1.Set("7", "test")

	got, err := cache1.Get("7")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("got=%v, want=%v\n", got, "test")

	user := cachetypes.User{ID: 1, Name: "user", Roles: [4]string{"admin", "root"}}
	fmt.Println("struct User hash: ", user.Hash())

	cache2 := cache.New[cachetypes.User, any](cache.BucketsCount)
	cache2.Set(user, "store this in cache")

	got, err = cache2.Get(user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("got=%v, want=%v\n", got, "store this in cache")

	cache3 := cache.New[cachetypes.Int, string](cache.BucketsCount)
	cache3.Set(-2, "string for cache3")
	got2, err := cache3.Get(-2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("got=%v, want=%v\n", got2, "string for cache3")

}
