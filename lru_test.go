package learn

import "testing"

func TestLru(t *testing.T) {
	// t.Fatal("not implemented")
	lru := LRU(5)
	lru.Set(1, 2)
	t.Log("key:", 1, "val:", lru.Get(1))
	t.Log("len:", lru.Len())
}
