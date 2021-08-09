package microservice

import (
	"testing"
)

type tt struct {
}

func TestMq(t *testing.T) {
	v := []interface{}{1, 2, 3, 4}
	t.Log(v[1:3])
}
