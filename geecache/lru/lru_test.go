package lru

import (
	"fmt"
	"testing"
)

type String string

func (d String) Len() int {
	return len(d)
}

func TestCache_Get(t *testing.T) {
	lru := New(int64(0), nil)
	lru.Add("key1", String("1234"))
	if v, ok := lru.Get("key1"); !ok || string(v.(String)) != "1234" {
		t.Fatalf("cache hit key1=1234 failed")
	}else {
		fmt.Println(v.(String))
	}
	if _, ok := lru.Get("key2"); ok {
		t.Fatalf("cache miss key2 failed")
	}else {
		fmt.Println("missing")
	}
}

func TestCache_Add(t *testing.T) {
	lru := New(int64(12), func(key string, value Value){
		t.Logf("cache key: %s value = %s", key, value.(String))
	})
	lru.Add("123", String("1221321"))
	lru.Add("123", String("12"))
	lru.Add("456", String("abcd"))
	lru.Add("123", String("127"))
	t.Log(lru.nBytes)
	if v, ok := lru.Get("123"); !ok {
		t.Fatalf("%s should not get key 123", v.(String))
	}
	if v, ok := lru.Get("456"); ok {
		t.Logf("%s", v.(String))
	}
}
