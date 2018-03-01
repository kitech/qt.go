package main

import (
	"os"
	"testing"
	"time"

	"github.com/kitech/qt.go/qtcore"
	"github.com/kitech/qt.go/qtrt"
)

// using go test tests/qxxx_test.go
func TestHello(t *testing.T) {
	qtrt.SetDebugFFICall(true)
	args := []string{"./hello"}
	app := qtcore.NewQCoreApplication(len(args), args, 0)
	go func() {
		time.Sleep(3 * time.Second)
		os.Exit(0)
	}()
	app.Exec()
}
