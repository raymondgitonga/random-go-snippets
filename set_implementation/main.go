package main

import "sync"

func main() {

}

type IntSet struct {
	set  []int
	keys map[int]bool
	//make it thread safe
	lock sync.RWMutex
}

func (s *IntSet) Add(val int) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.keys == nil {
		s.keys = make(map[int]bool)
		s.set = []int{}
	}
	if !s.keys[val] {
		s.keys[val] = true
		s.set = append(s.set, val)
	}
}

func (s *IntSet) List() []int {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.set
}
