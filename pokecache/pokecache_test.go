package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	cache := NewCache(time.Millisecond)
	if cache.cache == nil {
		t.Error("cache is nil")
	}
}

func TestAddGetCache(t *testing.T) {
	cache := NewCache(time.Millisecond)
	cases := []struct {
		inputKey string
		inputVal []byte
	}{
		{
			inputKey: "key",
			inputVal: []byte("value"),
		},
		{
			inputKey: "",
			inputVal: []byte("val"),
		},
	}
	for _, cs := range cases {
		cache.Add(cs.inputKey, []byte(cs.inputVal))
		actual, ok := cache.Get(cs.inputKey)
		if !ok {
			t.Errorf("%s not found", cs.inputKey)
		}
		if string(actual) != string(cs.inputVal) {
			t.Errorf("%s does not match %s", string(actual), string(cs.inputVal))
		}
	}
}

func TestReap(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)
	cache.Add("key", []byte("value"))
	time.Sleep(interval + time.Millisecond)
	_, ok := cache.Get("key")
	if ok {
		t.Error("key should have been reaped")
	}
}

func TestReapFail(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)
	cache.Add("key", []byte("value"))
	time.Sleep(interval / 2)
	_, ok := cache.Get("key")
	if !ok {
		t.Error("key should not have been reaped")
	}
}