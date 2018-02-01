package main

import (
	"gopp"
	"testing"
	"time"

	"qt.go/qtcore"
	"qt.go/qtrt"
)

func Test0(t *testing.T) {
	time.Sleep(5 * time.Second)
	qtrt.DebugFinal = false
	if true {
		for i := 0; i < 100; i++ {
			sz := qtcore.NewQString_5(gopp.RandomStringAlphaMixed(50000))
			_ = sz
			sz = nil
		}
	}

	time.Sleep(50 * time.Second)
}
