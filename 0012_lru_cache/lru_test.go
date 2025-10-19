package main

import (
	"iter"
	"testing"

	"github.com/stretchr/testify/require"
)

var _ LRUCache[int, int] = (*cacheImpl[int, int])(nil)

func TestLruCache(t *testing.T) {
	t.Parallel()

	const notFoundKey int = -1
	cache := New[int, int](2, notFoundKey)

	cache.Put(1, 1)

	cache.Put(2, 2)

	require.Equal(t, 1, cache.Get(1))

	cache.Put(3, 3)

	require.Equal(t, notFoundKey, cache.Get(2))
	require.Equal(t, 1, cache.Get(1))
	require.Equal(t, 3, cache.Get(3))
}

func TestLruCacheIterator(t *testing.T) {
	t.Parallel()

	const notFoundKey int = -1
	cache := New[int, int](2, notFoundKey)

	cache.Put(1, 1)

	cache.Put(2, 2)

	require.Equal(t, 1, cache.Get(1))

	cache.Put(3, 3)

	require.Equal(t, notFoundKey, cache.Get(2))
	require.Equal(t, 1, cache.Get(1))
	require.Equal(t, 3, cache.Get(3))

	keys, values := collect(cache.All())
	require.Equal(t, []int{3, 1}, keys)
	require.Equal(t, []int{3, 1}, values)

}

func collect[K comparable, V any](iterator iter.Seq2[K, V]) ([]K, []V) {
	keys := make([]K, 0)
	values := make([]V, 0)

	for k, v := range iterator {
		keys = append(keys, k)
		values = append(values, v)
	}

	return keys, values
}
