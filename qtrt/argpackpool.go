package qtrt

import (
	"reflect"

	"github.com/gammazero/deque"
)

type argPack struct {
	argtys   []byte
	argvals  []uint64
	argvals2 []Voidptr
	argrefps []*reflect.Value
}

func newArgPack() *argPack {
	const arrcnt = 20
	this := &argPack{}
	this.argtys = make([]byte, arrcnt)
	this.argvals = make([]uint64, arrcnt)
	this.argvals2 = make([]Voidptr, arrcnt)
	this.argrefps = make([]*reflect.Value, arrcnt)

	return this
}

var argpp = newArgPackPool(8)

type argPackPool struct {
	q deque.Deque
}

func newArgPackPool(initn int) *argPackPool {
	this := &argPackPool{}
	for i := 0; i < initn; i++ {
		item := newArgPack()
		this.q.PushBack(item)
	}
	return this
}

func (this *argPackPool) Put(item *argPack) {
	this.q.PushBack(item)
}

func (this *argPackPool) Get() (item *argPack) {
	if this.q.Len() != 0 {
		item = this.q.PopFront().(*argPack)
	}
	return
}
