package main

import (
	"testing"

	"qt.go/qtcore"
	"qt.go/qtrt"
)

func TestGetSignature0(t *testing.T) {
	tmer := qtcore.NewQTimer(nil)
	sig, err := qtrt.QObjectGetSignatureByName(tmer, "timeout")
	if err != nil {
		t.Fail()
	}

	sig, err := qtrt.QObjectGetSignatureByName(tmer, "timeout123")
	if err == nil {
		t.Fail()
	}
}
