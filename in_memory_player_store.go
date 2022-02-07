package main

import "sync"

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{
		store: map[string]int{},
	}
}

type InMemoryPlayerStore struct {
	mu    sync.Mutex
	store map[string]int
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.mu.Lock()
	defer i.mu.Unlock()

	i.store[name]++
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}

func (i *InMemoryPlayerStore) GetLeague() League {
	league := make([]Player, 0)
	for name, wins := range i.store {
		league = append(league, Player{name, wins})
	}
	return league
}
