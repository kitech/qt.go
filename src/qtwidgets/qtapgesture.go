//  header block begin
// /usr/include/qt/QtWidgets/qgesture.h
// #include <qgesture.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"
import "qtgui"

func init() {
	if false {
		reflect.TypeOf(123)
	}
	if false {
		reflect.TypeOf(unsafe.Sizeof(0))
	}
	if false {
		fmt.Println(123)
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
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin
type QTapGesture struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qgesture.h:236
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QTapGesture) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTapGesture10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:242
// index:0
// void QTapGesture(class QObject *)
func NewQTapGesture(parent unsafe.Pointer) *QTapGesture {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTapGestureC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QTapGesture{cthis}
}

// /usr/include/qt/QtWidgets/qgesture.h:243
// index:0
// virtual
// void ~QTapGesture()
func DeleteQTapGesture(*QTapGesture) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTapGestureD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:245
// index:0
// QPointF position()
func (this *QTapGesture) Position() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTapGesture8positionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:246
// index:0
// void setPosition(const class QPointF &)
func (this *QTapGesture) SetPosition(pos unsafe.Pointer) {
	// 0: (, const QPointF & pos), (pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTapGesture11setPositionERK7QPointF", ffiqt.FFI_TYPE_VOID, this.cthis, pos)
	gopp.ErrPrint(err, rv)
}

//  body block end
