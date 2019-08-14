
/* Sync Mutex

Mutex locks allow serial access of shared resources by causing goroutines to block and wait until locks are released.


The following sample illustrates a typical code scenario with the
Service type, which must be started before it is ready to be used. After the service has
started, the code updates an internal bool variable, started, to store its current state

*/


package main

import (
	"sync"
	"time"
)

type Service struct {
	started bool
	stpCh   chan struct{}
	mutex   sync.Mutex
}

func (s *Service) Start() {
	s.stpCh = make(chan struct{})
	go func() {
		s.mutex.Lock()
		s.started = true
		s.mutex.Unlock()
		<-s.stpCh // wait to be closed.
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
	time.Sleep(time.Second) // do some work
	s.Stop()
}

