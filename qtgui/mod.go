package qtgui

/*
#cgo CFLAGS: -std=c11
// #cgo LDFLAGS: -lQt5Inline
//  -lQt5Core -lQt5Gui -lQt5Widgets
*/
import "C"
import "qt.go/qtcore"

func init() {
	if false {
		qtcore.KeepMe()
	}
}

func KeepMe() {}
