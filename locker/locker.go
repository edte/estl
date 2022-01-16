// @program:     tiny-stl
// @file:        locker.go
// @author:      edte
// @create:      2022-01-15 23:50
// @description:
package locker

import (
	"sync"
)

type Locker interface {
	Lock()
	Unlock()
	RLock()
	RUnlock()
}

type RWLock struct {
	rw sync.RWMutex
}

func NewRWLock() *RWLock {
	return &RWLock{
		rw: sync.RWMutex{},
	}
}

func (r *RWLock) Lock() {
	r.rw.Lock()
}

func (r *RWLock) Unlock() {
	r.rw.Unlock()
}

func (r *RWLock) RLock() {
	r.rw.RLock()
}

func (r *RWLock) RUnlock() {
	r.rw.RUnlock()
}

type Mutex struct {
	m sync.Mutex
}

func NewMutex() *Mutex {
	return &Mutex{
		m: sync.Mutex{},
	}
}

func (m *Mutex) Lock() {
	m.m.Lock()
}

func (m *Mutex) Unlock() {
	m.m.Unlock()
}

func (m *Mutex) RLock() {

}

func (m *Mutex) RUnlock() {

}
