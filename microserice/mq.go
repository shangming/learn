package microservice

import "sync"

type Ack struct {
	Id    int
	Topic string
}

type msgQueue struct {
	mu        sync.Mutex
	consumers map[int]int
	data      []interface{}
}

type Record struct {
	Offset int
	Msg    interface{}
}

type Mq struct {
	msgQueues sync.Map
}

func newMsgQueue() *msgQueue {
	q := &msgQueue{}
	q.consumers = make(map[int]int, 0)
	return q
}

func (mq *Mq) getMsgQueue(topic string) *msgQueue {
	x, _ := mq.msgQueues.LoadOrStore(topic, newMsgQueue())
	q, _ := x.(*msgQueue)
	return q
}

func (mq *Mq) Send(topic string, msg interface{}) <-chan Ack {

	c := make(chan Ack)

	go func() {
		q := mq.getMsgQueue(topic)
		q.mu.Lock()
		q.data = append(q.data, msg)
		ack := Ack{Id: len(q.data) - 1, Topic: topic}
		q.mu.Unlock()

		c <- ack
	}()

	return c
}

func (mq *Mq) Poll(consumerId int, topic string, num int) []Record {

	msgs := []Record{}

	q := mq.getMsgQueue(topic)
	q.mu.Lock()
	defer q.mu.Unlock()

	start, ok := q.consumers[consumerId]

	if !ok {
		q.consumers[consumerId] = 0
		start = 0
	}

	if num > len(q.data)-start {
		num = len(q.data) - start
	}

	for i := start; i < start+num; i++ {
		msgs = append(msgs, Record{Offset: i, Msg: q.data[i]})
	}

	return msgs
}

func (mq *Mq) Confirm(consumerId int, topic string, offset int) <-chan bool {

	c := make(chan bool)

	go func() {
		q := mq.getMsgQueue(topic)
		q.mu.Lock()
		defer q.mu.Unlock()
		q.consumers[consumerId] = offset + 1
		c <- true
	}()

	return c
}
