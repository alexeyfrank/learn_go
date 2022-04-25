package server

import "sync"

type Storage struct {
	lock     sync.Mutex
	counters map[string]int64
}

func NewStorage() *Storage {
	return &Storage{
		counters: make(map[string]int64),
	}
}

func (s *Storage) SetCounter(key string, val int64) (int64, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.counters[key] = val

	return val, nil
}

func (s *Storage) GetCounter(key string) (int64, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.counters[key], nil
}

func (s *Storage) IncrementCounter(key string, val int64) (int64, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.counters[key] += val

	return s.counters[key], nil
}

func (s *Storage) DecrementCounter(key string, val int64) (int64, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.counters[key] -= val

	return s.counters[key], nil
}
