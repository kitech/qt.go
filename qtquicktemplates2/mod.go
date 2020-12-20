package qtquicktemplates2

import (
	"unsafe"

	"github.com/kitech/qt.go/qtrt"
)

func KeepMe() {}
func init()   { qtrt.RegisterSubPackage("Quick") }

type Voidptr = unsafe.Pointer
