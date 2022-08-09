package main

import "sync"

type RWMap struct {
	sync.RWMutex
	m map[int]int
}

func (rw *RWMap) NewMap(n int) *RWMap {
	return &RWMap{
		m: make(map[int]int, n),
	}
}
func (rw *RWMap) Set(k, v int) {
	rw.Lock()
	defer rw.Unlock()
	rw.m[k] = v
}
func (rw *RWMap) Get(k int) (int, bool) {
	rw.RLock()
	defer rw.RUnlock()
	v, exit := rw.m[k]
	return v, exit

}
func (rw *RWMap) Delete(k int) {
	rw.RLock()
	defer rw.RUnlock()
	delete(rw.m, k)
}
func (rw *RWMap) Each(f func(k, v int) bool) {
	rw.RLock()
	defer rw.RUnlock()
	for k, v := range rw.m {
		if !f(k, v) {
			return
		}
	}
}
