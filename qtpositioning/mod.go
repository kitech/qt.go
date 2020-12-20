package qtpositioning

/*
#cgo CFLAGS: -std=c11
// #cgo LDFLAGS: -lQt5Inline
//  -lQt5Core -lQt5Gui -lQt5Widgets
*/
// import "C"
import (
	"unsafe"

	"github.com/kitech/qt.go/qtcore"
	"github.com/kitech/qt.go/qtrt"
)

func init() {
	if false {
		qtcore.KeepMe()
	}
}

func KeepMe() {}
func init()   { qtrt.RegisterSubPackage("Positioning") }

type Voidptr = unsafe.Pointer
