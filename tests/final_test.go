package main

import (
	"testing"
	"time"

	gopp "github.com/kitech/goplusplus"
	"github.com/kitech/qt.go/qtcore"
	"github.com/kitech/qt.go/qtrt"
)

func Test0(t *testing.T) {
	time.Sleep(5 * time.Second)
	qtrt.DebugFinal = false
	if true {
		for i := 0; i < 100; i++ {
			sz := qtcore.NewQString5(gopp.RandomStringAlphaMixed(50000))
			_ = sz
			sz = nil
		}
	}

	time.Sleep(50 * time.Second)
}
