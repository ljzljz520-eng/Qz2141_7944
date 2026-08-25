package notify

import (
	"fmt"
	"sync"
	"training29/model"
)

type Message struct{ To, Subject, Body string }
type Sender struct {
	mu   sync.Mutex
	Sent []Message
}

func New() *Sender { return &Sender{Sent: []Message{}} }
func (n *Sender) Send(u model.User, r model.Record) Message {
	m := Message{To: u.Email, Subject: "培训29状态更新", Body: fmt.Sprintf("记录 %s 当前状态 %s", r.ID, r.Status)}
	n.mu.Lock()
	n.Sent = append(n.Sent, m)
	n.mu.Unlock()
	return m
}
func (n *Sender) Count() int { n.mu.Lock(); defer n.mu.Unlock(); return len(n.Sent) }
func (n *Sender) Last() Message {
	n.mu.Lock()
	defer n.mu.Unlock()
	if len(n.Sent) == 0 {
		return Message{}
	}
	return n.Sent[len(n.Sent)-1]
}
