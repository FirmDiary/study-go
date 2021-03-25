package main

import "sync"

type URLStore struct {
	urls map[string]string
	mu   sync.RWMutex
}

func NewUrlStore() *URLStore {
	return &URLStore{urls: make(map[string]string)}
}

func (s *URLStore) Get(key string) string {
	s.mu.RLock()
	defer s.mu.Unlock()
	return s.urls[key]
}

func (s *URLStore) Set(key string, url string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if _, present := s.urls[key]; present {
		return false
	}

	s.urls[key] = url
	return true
}

func (s *URLStore) Count() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.urls)
}

func (s *URLStore) Put(url string) string {

	for {
		//key := getKey
	}
}
