package cmemory

import (
	"log"
	"math"
	"runtime"
	"sync"
	"sync/atomic"
)

// only support one call with a time add
type CBMap struct {
	mu sync.RWMutex
	m  map[uint64]interface{}
	no uint64
}

func NewCBMap() *CBMap {
	this := &CBMap{}
	this.m = make(map[uint64]interface{})
	return this
}

func (this *CBMap) Add(v interface{}) (cbno uint64) {
	var n uint64
	this.mu.RLock()
	for {
		if uint64(len(this.m)) == math.MaxUint64 {
			log.Fatalln("CBMap is full")
		}
		if runtime.GOOS == "android" {
			if this.no == math.MaxUint64 {
				this.no = 0
			}
			n = this.no + 1
		} else {
			if atomic.CompareAndSwapUint64(&this.no, math.MaxUint64, 0) {
			}
			n = atomic.AddUint64(&this.no, 1)
		}
		if _, ok := this.m[n]; ok {
			continue
		} else {
			break
		}
	}
	this.mu.RUnlock()
	this.mu.Lock()
	this.m[n] = v
	this.mu.Unlock()
	return n
}

func (this *CBMap) Get(cbno uint64) (interface{}, bool) {
	this.mu.RLock()
	defer this.mu.RUnlock()
	if v, ok := this.m[cbno]; ok {
		return v, ok
	}
	return nil, false
}

func (this *CBMap) Del(cbno uint64) {
	this.mu.Lock()
	defer this.mu.Unlock()
	if _, ok := this.m[cbno]; ok {
		delete(this.m, cbno)
	}
}
