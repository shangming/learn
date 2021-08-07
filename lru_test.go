package learn

import (
	"testing"
)

type val struct {
	a int
}

func TestLru(t *testing.T) {
	// t.Fatal("not implemented")
	lru := LRU(5)

	for i := 0; i < 6; i++ {
		lru.Set(i, val{a: i})
		t.Log(lru.GetLinkKey())
	}

	t.Log("get key:", 3, "val:", lru.Get(3))
	t.Log(lru.GetLinkKey())

	t.Log("get key:", 1, "val:", lru.Get(1))
	t.Log(lru.GetLinkKey())

	t.Log("get key:", 0, "val:", lru.Get(0))
	t.Log(lru.GetLinkKey())
}
