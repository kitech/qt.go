package qtqml

import "unsafe"
import "gopp"
import "qt.go/cffiqt"
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
		ffiqt.KeepMe()
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
// index:43
// Invalid inline
// uint qHash(const QQmlProperty &)
func QHash_43(key *QQmlProperty) uint {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_Z5qHashRK12QQmlProperty", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}

//  body block end
