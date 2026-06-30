package _4_sharded_lock

import (
	"hash"
	"hash/fnv"
	"sync"
)

//type Map[K comparable, V any] struct {
//	data map[K]V
//	mx sync.Mutex
//}

type SyncMap[K comparable, V any] struct {
	//shards []shard[K,V]

	// mx data mx data

	mx   []sync.Mutex
	data []map[K]V
}

//type shard[K comparable, V any] struct {
//	mx   sync.Mutex
//	data map[K]V
//}

type ShardedMap struct {
	locks []sync.Mutex
	data  []map[string]int
	hash  hash.Hash32
}

func NewShardedMap(shards int) *ShardedMap {
	data := make([]map[string]int, shards)
	locks := make([]sync.Mutex, shards)

	for i := 0; i < shards; i++ {
		data[i] = make(map[string]int) // default
	}

	return &ShardedMap{locks: locks, data: data}
}

func fnv32(key string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(key))
	return h.Sum32()
}

func (m *ShardedMap) getShard(key string) int {
	return int(fnv32(key)) % len(m.data)
}

func (m *ShardedMap) Set(key string, value int) {
	shard := m.getShard(key)

	m.locks[shard].Lock()
	m.data[shard][key] = value
	m.locks[shard].Unlock()
}

func (m *ShardedMap) Get(key string) (int, bool) {
	shard := m.getShard(key)
	m.locks[shard].Lock()
	val, ok := m.data[shard][key]
	m.locks[shard].Unlock()
	return val, ok
}
