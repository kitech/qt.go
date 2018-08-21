package main

import (
	"log"
	"testing"

	"github.com/kitech/qt.go/qtcore"
)

func Test0(t *testing.T) {
	var optname string
	optname = qtcore.QProcess_ExitStatusItemName(1)
	log.Println(optname)
	if optname != "CrashExit" {
		t.Fail()
	}

	optname = qtcore.QUrl_ComponentFormattingOptionItemName(qtcore.QUrl__EncodeDelimiters)
	log.Println(optname)
	if optname != "EncodeDelimiters" {
		t.Fail()
	}

	optname = qtcore.KeyItemName(qtcore.Qt__Key_5)
	log.Println(optname)
	if optname != "Key_5" {
		t.Fail()
	}
}
