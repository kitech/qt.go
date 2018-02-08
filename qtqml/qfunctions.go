package qtqml

import "unsafe"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtnetwork"

func init() {
	if false {
		_ = unsafe.Pointer(uintptr(0))
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtnetwork.KeepMe()
	}
}

//  header block end

//  body block begin
// /usr/include/qt/QtQml/qqmlproperty.h:130
// index:42
// Invalid inline Visibility=Default Availability=Available
// [4] uint qHash(const QQmlProperty &)
func QHash_42(key *QQmlProperty) uint {
	var convArg0 = key.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_Z5qHashRK12QQmlProperty", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	return uint(rv) // 222
}

//  body block end
