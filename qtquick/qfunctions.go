package qtquick

import "unsafe"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtnetwork"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtqml"

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

/*

 */
func Qsgnode_set_description(node QSGNode_ITF /*777 QSGNode **/, description string) {
	var convArg0 unsafe.Pointer
	if node != nil && node.QSGNode_PTR() != nil {
		convArg0 = node.QSGNode_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(description)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_Z23qsgnode_set_descriptionP7QSGNodeRK7QString", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

//  body block end
