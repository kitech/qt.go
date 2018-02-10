package qtquick

import "unsafe"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtnetwork"
import "qt.go/qtgui"
import "qt.go/qtqml"

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
	if false {
		qtgui.KeepMe()
	}
	if false {
		qtqml.KeepMe()
	}
}

//  header block end

//  body block begin
// /usr/include/qt/QtQuick/qsgnode.h:193
// index:0
// Invalid Visibility=Default Availability=Available
// [-2] void qsgnode_set_description(QSGNode *, const QString &)
func Qsgnode_set_description(node *QSGNode /*777 QSGNode **/, description string) {
	var convArg0 = node.GetCthis()
	var tmpArg1 = qtcore.NewQString_5(description)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_Z23qsgnode_set_descriptionP7QSGNodeRK7QString", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

//  body block end
