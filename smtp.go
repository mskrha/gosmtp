package gosmtp

import (
	"fmt"
	"net/smtp"
	"os"
	"strings"
	"time"
)

const (
	RFC2822 = "Mon, _2 Jan 2006 15:04:05 -0700"
)

type SMTP struct {
	server string
	agent  string
}

func NewServer(h, a string) (*SMTP, error) {
	var s SMTP

	if len(h) == 0 {
		return nil, fmt.Errorf("No SMTP relay specified")
	}

	switch len(strings.Split(h, ":")) {
	case 1:
		s.server = fmt.Sprintf("%s:25", h)
	case 2:
		s.server = h
	default:
		return nil, fmt.Errorf("SMTP relay invalid")
	}

	if len(a) == 0 {
		return nil, fmt.Errorf("No user-agent specified")
	}
	s.agent = a

	return &s, nil
}

func (s *SMTP) Send(m Message) (err error) {
	host, err := os.Hostname()
	if err != nil {
		return
	}

	err = m.verify()
	if err != nil {
		return
	}

	var x string
	t := time.Now()
	x += fmt.Sprintf("From: %s\r\n", m.From)
	x += fmt.Sprintf("To: %s\r\n", m.To)
	x += fmt.Sprintf("Subject: %s\r\n", m.Subject)
	x += fmt.Sprintf("Date: %s\r\n", t.Format(RFC2822))
	x += fmt.Sprintf("Message-ID: %d@%s\r\n", t.UnixNano(), host)
	x += fmt.Sprintf("User-Agent: %s\r\n", s.agent)
	x += "\r\n"
	x += m.Body
	x += "\r\n"

	c, err := smtp.Dial(s.server)
	if err != nil {
		return
	}

	err = c.Hello(host)
	if err != nil {
		return
	}

	err = c.Mail(m.From)
	if err != nil {
		return
	}

	err = c.Rcpt(m.To)
	if err != nil {
		return
	}

	d, err := c.Data()
	if err != nil {
		return
	}

	_, err = d.Write([]byte(x))
	if err != nil {
		return
	}

	return c.Quit()
}
