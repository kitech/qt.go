package qtwidgets

/*
#cgo CFLAGS: -std=c11
// #cgo LDFLAGS: -lQt5Inline
//  -lQt5Core -lQt5Gui -lQt5Widgets
*/
import "C"
import "qtcore"
import "qtgui"

func init() {
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

func KeepMe() {}
