package qtmock

import (
	"unsafe"

	"github.com/kitech/qt.go/qtcore"
	"github.com/kitech/qt.go/qtrt"
)

func KeepMe() {}

func QApplication_Translate(a string, b string, c string, d int) *qtcore.QString {
	// return qtcore.NewQString_5("mockhe")
	return qtcore.NewQString5(b)
}

func QCoreApplication_Translate(a string, b string, c string, d int) *qtcore.QString {
	// return qtcore.NewQString_5("mockhe")
	return qtcore.NewQString5(b)
}

// bool qRegisterResourceData(int, unsigned char const*, unsigned char const*, unsigned char const*)
func QRegisterResourceData(version int, struct_ unsafe.Pointer, name_ unsafe.Pointer, data_ unsafe.Pointer) bool {
	rv, err := qtrt.InvokeQtFunc6("_Z21qRegisterResourceDataiPKhS0_S0_", qtrt.FFI_TYPE_POINTER, version, struct_, name_, data_)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// bool qUnregisterResourceData(int, unsigned char const*, unsigned char const*, unsigned char const*)
func QUnregisterResourceData(version int, struct_ unsafe.Pointer, name_ unsafe.Pointer, data_ unsafe.Pointer) bool {
	rv, err := qtrt.InvokeQtFunc6("_Z23qUnregisterResourceDataiPKhS0_S0_", qtrt.FFI_TYPE_POINTER, version, struct_, name_, data_)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
