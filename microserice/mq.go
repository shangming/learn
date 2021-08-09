package microservice

import "sync"

const (
	ACK = 1
)

type Mq struct {
	mu sync.Mutex

	msgList   map[string][]interface{}
	msgListMu map[string]sync.Mutex
}

func (mq *Mq) send(key string, msg interface{}, ackChan <-chan int) {

}
