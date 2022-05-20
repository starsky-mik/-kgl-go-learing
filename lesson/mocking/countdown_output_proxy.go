package mocking

import (
	"bytes"
	"fmt"
	"io"
	"time"
)

type CountdownOutputProxy interface {
	Write(s any) (n int, err error)
	Sleep()
}

type DefaultCountdownOutputProxy struct {
	Out      io.Writer
	Interval time.Duration
}

func (o DefaultCountdownOutputProxy) Write(s any) (n int, err error) {
	return fmt.Fprintln(o.Out, s)
}

func (o DefaultCountdownOutputProxy) Sleep() {
	time.Sleep(o.Interval)
}

type SpyCountdownOutputProxy struct {
	out      bytes.Buffer
	Interval time.Duration
	log      []string
}

func (p *SpyCountdownOutputProxy) Write(s any) (n int, err error) {
	p.log = append(p.log, fmt.Sprintf("Write line: \"%v\"", s))
	return fmt.Fprintln(&p.out, s)
}

func (p *SpyCountdownOutputProxy) Sleep() {
	p.log = append(p.log, fmt.Sprintf("Wait %v", p.Interval))
}

func (p SpyCountdownOutputProxy) Out() string {
	return p.out.String()
}

func (p SpyCountdownOutputProxy) Log() []string {
	return p.log
}
