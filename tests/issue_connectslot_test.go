package main

import (
	"log"
	"os"
	"testing"

	"github.com/kitech/qt.go/qtcore"
	"github.com/kitech/qt.go/qtrt"
)

// usage: go test -v (-run Connqq) tests/qxxx_test.go
func TestConnqq(t *testing.T) {
	log.SetFlags(log.Flags() | log.Lshortfile)
	qtrt.SetDebugFFICall(true)
	qtrt.SetDebugDynSlot(true)

	app := qtcore.NewQCoreApplication(len(os.Args), os.Args, 0)
	log.Println(app)

	tmer := qtcore.NewQTimer(nil)

	qtrt.ConnectSlot(tmer, "timeout()", app, "quit()")
	tmer.Start(2000)

	app.Exec()
}
