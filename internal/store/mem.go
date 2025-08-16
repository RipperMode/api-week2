package store

import (
	"sync"

	"api-test/internal/model"
)

type MemStore struct {
	mu    sync.RWMutex
	tasks []model.Task
	next  int64
}

func NewMemStore() *MemStore {
	return &MemStore{tasks: []model.Task{}, next: 1}
}

func (s *MemStore) Create(title string) model.Task {
	s.mu.Lock()
	defer s.mu.Unlock()
	task := model.Task{ID: s.next, Title: title, Done: false}
	s.next++
	s.tasks = append(s.tasks, task)
	return task
}

func (s *MemStore) List() []model.Task {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return append([]model.Task(nil), s.tasks...) // copy
}

func (s *MemStore) Get(id int64) (model.Task, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, t := range s.tasks {
		if t.ID == id {
			return t, true
		}
	}
	return model.Task{}, false
}

func (s *MemStore) Update(id int64, title string, done bool) (model.Task, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, t := range s.tasks {
		if t.ID == id {
			s.tasks[i].Title = title
			s.tasks[i].Done = done
			return s.tasks[i], true
		}
	}
	return model.Task{}, false
}

func (s *MemStore) Delete(id int64) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, t := range s.tasks {
		if t.ID == id {
			s.tasks = append(s.tasks[:i], s.tasks[i+1:]...)
			return true
		}
	}
	return false
}
