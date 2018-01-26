package qtmock

import (
	"mkuse/cffiqt"
	"qtcore"
	"unsafe"
)

func KeepMe() {}

func QApplication_Translate(a string, b string, c string, d int) *qtcore.QString {
	// return qtcore.NewQString_5("mockhe")
	return qtcore.NewQString_5(b)
}

func QCoreApplication_Translate(a string, b string, c string, d int) *qtcore.QString {
	// return qtcore.NewQString_5("mockhe")
	return qtcore.NewQString_5(b)
}

// bool qRegisterResourceData(int, unsigned char const*, unsigned char const*, unsigned char const*)
func QRegisterResourceData(version int, struct_ unsafe.Pointer, name_ unsafe.Pointer, data_ unsafe.Pointer) {
	ffiqt.InvokeQtFunc6("_Z21qRegisterResourceDataiPKhS0_S0_", ffiqt.FFI_TYPE_POINTER, version, struct_, name_, data_)
}

// bool qUnregisterResourceData(int, unsigned char const*, unsigned char const*, unsigned char const*)
func QUnregisterResourceData(version int, struct_ unsafe.Pointer, name_ unsafe.Pointer, data_ unsafe.Pointer) {
	ffiqt.InvokeQtFunc6("_Z23qUnregisterResourceDataiPKhS0_S0_", ffiqt.FFI_TYPE_POINTER, version, struct_, name_, data_)
}
