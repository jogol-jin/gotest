package main

import (
	"fmt"
	"log"
	"time"
)

const (
	PUSH = "push"
	POP  = "pop"
	EXIT = "exit"
)

type Stack struct {
	len     int
	front   Data
	backend Data
	c       chan *Operation
	closed  bool
}

type Operation struct {
	operate string
	value   interface{}
	c       chan *Result
}

type Result struct {
	value interface{}
	err   error
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
		c:   make(chan *Operation),
	}
	s.front.next = &s.backend
	s.front.prev = &s.front
	s.backend.prev = &s.front
	s.backend.next = &s.backend
	go s.handle()
	return &s
}

func (s *Stack) Close() error {
	if s == nil {
		return fmt.Errorf("%s", "stack is nil")
	}
	operation := newOperation(EXIT, nil)
	s.c <- operation
	result := <-operation.c
	if result == nil {
		return fmt.Errorf("%s", "close error")
	}
	s = nil

	return result.err
}

func (s *Stack) handle() {
	if s == nil {
		return
	}
	for {
		operation := <-s.c
		if operation == nil {
			continue
		}

		//log.Printf("进行操作:%s, 值:%v", operation.operate, operation.value)
		switch operation.operate {
		case PUSH:
			if s.closed {
				operation.c <- &Result{
					err: fmt.Errorf("%s", "stack is closed"),
				}
			} else {
				err := s.add(operation.value)
				operation.c <- &Result{
					err: err,
				}
			}
		case POP:
			if s.closed {
				operation.c <- &Result{
					err: fmt.Errorf("%s", "stack is closed"),
				}
			} else {
				value, err := s.getAndRemove()
				operation.c <- &Result{
					value: value,
					err:   err,
				}
			}
		case EXIT:
			s.closed = true
			operation.c <- &Result{}
			//return
		}
		//log.Printf("目前长度:%d, 首位:%v, 末尾:%v", s.len, s.front.next.Value, s.backend.prev.Value)

	}
	return
}

func (s *Stack) isEmpty() bool {
	if s == nil {
		return false
	}
	return s.len <= 0
}

func (s *Stack) getAndRemove() (interface{}, error) {
	if s == nil {
		return nil, fmt.Errorf("%s", "stack is nil")
	}
	if s.isEmpty() {
		return nil, fmt.Errorf("%s", "stack is empty")
	}
	bp := s.backend.prev

	bpp := bp.prev
	bpp.next = &s.backend
	s.backend.prev = bpp
	return bp, nil
}

func (s *Stack) add(value interface{}) error {
	if s == nil {
		return fmt.Errorf("%s", "stack is nil")
	}
	d := Data{
		Value: value,
		stack: s,
	}
	bp := s.backend.prev
	bp.next = &d
	d.prev = bp
	d.next = &s.backend
	s.backend.prev = &d
	s.len++
	return nil
}
func newOperation(operate string, value interface{}) *Operation {
	return &Operation{
		operate: operate,
		c:       make(chan *Result, 1),
		value:   value,
	}
}

func (s *Stack) Pop() (interface{}, error) {
	if s == nil {
		return nil, fmt.Errorf("%s", "stack is nil")
	}
	if s.c == nil {
		return nil, fmt.Errorf("%s", "stack chan si nil")
	}
	operation := newOperation(POP, nil)
	s.c <- operation
	result := <-operation.c
	if result == nil {
		return nil, fmt.Errorf("%s", "pop error")
	}

	return result.value, result.err
}

func (s *Stack) Push(value interface{}) error {
	if s == nil {
		return fmt.Errorf("%s", "stack is nil")
	}
	if s.c == nil {
		return fmt.Errorf("%s", "stack chan si nil")
	}
	operation := newOperation(PUSH, value)
	s.c <- operation
	result := <-operation.c
	if result == nil {
		return fmt.Errorf("%s", "push error")
	}
	return result.err
}

func main() {
	s := NewStack()

	count := 0
	for i := 0; i <= 100; i++ {
		count++
		go addData(s)
	}
	popData(s)
	log.Println("hahaah")
	time.Sleep(5 * time.Second)

}

func addData(s *Stack) {
	for i := 0; i < 100; i++ {
		s.Push(time.Now())
	}
}

func popData(s *Stack) {
	for i := 0; i < 100; i++ {
		d, _ := s.Pop()
		log.Printf("data:%v", d)
	}
}
