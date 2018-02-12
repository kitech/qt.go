package main

import (
	"log"
	"testing"

	"github.com/kitech/qt.go/qtcore"
)

func Test0(t *testing.T) {
	so := qtcore.QCoreApplication_Translate("MainWindow", "...", "dummy123", 0)
	log.Println(so, so.Length())
	if so.Length() != 3 {
		t.Fail()
	}
}
