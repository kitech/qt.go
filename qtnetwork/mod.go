package qtnetwork

/*
#cgo CFLAGS: -std=c11
// #cgo LDFLAGS: -lQt5Inline
//  -lQt5Core -lQt5Gui -lQt5Widgets
*/
import "C"
import (
	"unsafe"

	"github.com/kitech/qt.go/qtrt"
)

func KeepMe() {}
func init()   { qtrt.RegisterSubPackage("Network") }

type Voidptr = unsafe.Pointer
