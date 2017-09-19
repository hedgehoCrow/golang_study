package main

import (
	"fmt"
	"sync"
)

// KeyValue のための型。内部にmapを保持している
type KeyValue struct {
	store map[string]string // key-valueを格納するためのmap
	mu    sync.RWMutex      // 排他制御のためのmutex
}

func NewKeyValue() *KeyValue {
	return &KeyValue{store: make(map[string]string)}
}

func (kv *KeyValue) Set(key, val string) {
	kv.mu.Lock()         // まずLock
	defer kv.mu.Unlock() // メソッドを抜けた際にUnlock
	kv.store[key] = val
}

func (kv *KeyValue) Get(key string) (string, bool) {
	kv.mu.RLock()         // 参照用のRLock
	defer kv.mu.RUnlock() // メソッドを抜けた際にRUnlock
	val, ok := kv.store[key]
	return val, ok
}

func main() {
	kv := NewKeyValue()
	kv.Set("key", "value")
	value, ok := kv.Get("key")
	if ok {
		fmt.Println(value)
	}
}
