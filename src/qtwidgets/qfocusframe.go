//  header block begin
// /usr/include/qt/QtWidgets/qfocusframe.h
// #include <qfocusframe.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 50
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
type QFocusFrame struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qfocusframe.h:54
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QFocusFrame) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFocusFrame10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfocusframe.h:56
// index:0
// void QFocusFrame(class QWidget *)
func NewQFocusFrame(parent unsafe.Pointer) *QFocusFrame {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFocusFrameC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QFocusFrame{cthis}
}

// /usr/include/qt/QtWidgets/qfocusframe.h:57
// index:0
// virtual
// void ~QFocusFrame()
func DeleteQFocusFrame(*QFocusFrame) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFocusFrameD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfocusframe.h:59
// index:0
// void setWidget(class QWidget *)
func (this *QFocusFrame) SetWidget(widget unsafe.Pointer) {
	// 0: (, QWidget * widget), (widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFocusFrame9setWidgetEP7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfocusframe.h:60
// index:0
// QWidget * widget()
func (this *QFocusFrame) Widget() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFocusFrame6widgetEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
