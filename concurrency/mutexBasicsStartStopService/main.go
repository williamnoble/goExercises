package main

import (
	"sync"
	"time"
)

type Service struct {
	started bool
	stpCh   chan struct{}
	mutex   sync.Mutex
	cache   map[int]string
}

func (s *Service) Start() {
	s.stpCh = make(chan struct{})
	go func() {
		s.mutex.Lock()
		s.started = true
		s.cache[1] = "CacheObject1"
		s.mutex.Unlock()
	}()
}

func (s *Service) Stop() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.started {
		s.started = false
		close(s.stpCh)
	}
}
func main() {
	s := &Service{}
	s.Start()
	time.Sleep(1 * time.Second)
	s.Stop()
}
