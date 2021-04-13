package gosmtp

import (
	"fmt"
)

type Message struct {
	From    string
	To      string
	Subject string
	Body    string
}

func NewMessage(f, t, s, b string) (Message, error) {
	m := Message{From: f, To: t, Subject: s, Body: b}
	return m, m.verify()
}

func (m Message) verify() error {
	if len(m.From) == 0 {
		return fmt.Errorf("No sender specified")
	}
	if len(m.To) == 0 {
		return fmt.Errorf("No recipient specified")
	}
	if len(m.Subject) == 0 {
		return fmt.Errorf("No subject specified")
	}
	if len(m.Body) == 0 {
		return fmt.Errorf("No body specified")
	}
	return nil
}
