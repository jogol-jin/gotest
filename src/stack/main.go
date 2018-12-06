package main

import (
	"sync"
	"fmt"
)

const (
	PUSH = "push"
	POP  = "pop"
)

type Stack struct {
	mux     sync.RWMutex
	len     int
	front   Data
	backend Data
}

type Data struct {
	prev  *Data
	next  *Data
	Value interface{}
	stack *Stack
}

func NewStack() *Stack {
	s := Stack{
		len: 0,
	}
	s.front.next = &s.backend
	s.front.prev = &s.front
	s.backend.prev = &s.front
	s.backend.next = &s.backend
	return &s
}

func (s *Stack) init() {

}

func (s *Stack) handle() error {
	if s == nil {
		return fmt.Errorf("%s", "stack is nil")
	}
	for {
		select {
		case operate := <-s.c:
			switch operate {
			case PUSH:

			}
		}
	}
	return nil
}

func (s *Stack) Push(data interface{}) {
	s.mux.Lock()
	defer s.mux.Unlock()
	d := Data{
		Value: data,
		stack: s,
	}
	bp := s.backend.prev
	bp.next = &d
	d.prev = bp
	d.next = &s.backend
	s.backend.prev = &d
	s.len++
	return
}

func (s *Stack) Pop() interface{} {
	s.mux.Lock()
	defer s.mux.Unlock()

	if s.len <= 0 {
		return nil
	}
	bp := s.backend.prev

	bpp := bp.prev
	bpp.next = &s.backend
	s.backend.prev = bpp
	return bp
}

func main() {
	s := NewStack()
	fmt.Printf("%#v\n", s.front.next.Value)
	fmt.Printf("%#v\n", s.backend.prev.Value)
	s.Push(1)
	fmt.Printf("%#v\n", s.front.next.Value)
	fmt.Printf("%#v\n", s.backend.prev.Value)
	s.Push(2)
	fmt.Printf("%#v\n", s.front.next.Value)
	fmt.Printf("%#v\n", s.backend.prev.Value)
	s.Push(3)
	fmt.Printf("%#v\n", s.front.next.Value)
	fmt.Printf("%#v\n", s.backend.prev.Value)
}
