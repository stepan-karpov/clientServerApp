package utils

import (
 "sync"
)

type AtomicString struct {
 mu    sync.RWMutex
 value string
}

func (s *AtomicString) Load() string {
 s.mu.RLock()
 defer s.mu.RUnlock()
 return s.value
}

func (s *AtomicString) Store(val string) {
 s.mu.Lock()
 defer s.mu.Unlock()
 s.value = val
}
