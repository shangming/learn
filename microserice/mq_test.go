package microservice

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
)

type tt struct {
}

func TestMq(t *testing.T) {
	mq := Mq{}

	var wg1 sync.WaitGroup
	var wg2 sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg1.Add(1)
		go func(i int) {
			for j := 0; j < 300; j++ {
				msg := fmt.Sprint(i, "-", j)
				future := mq.Send("test", msg)
				ack := <-future
				t.Log("send:", i, msg, ack)
			}
			wg1.Done()
		}(i)
	}

	for i := 0; i < 2; i++ {
		wg2.Add(1)
		go func(i int) {
			for true {
				records := mq.Poll(i, "test", rand.Intn(10))
				t.Log("poll:", i, records)
				if len(records) > 0 {
					offset := records[len(records)-1].Offset
					future := mq.Confirm(i, "test", offset)
					ack := <-future
					t.Log("confirm:", i, offset, ack)
					if records[len(records)-1].Offset >= 30000-1 {
						break
					}
				}
			}
			wg2.Done()
		}(i)
	}
	wg1.Wait()
	wg2.Wait()
}
