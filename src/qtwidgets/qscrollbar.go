//  header block begin
// /usr/include/qt/QtWidgets/qscrollbar.h
// #include <qscrollbar.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 6
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
type QScrollBar struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qscrollbar.h:57
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QScrollBar) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QScrollBar10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollbar.h:59
// index:0
// void QScrollBar(class QWidget *)
func NewQScrollBar(parent unsafe.Pointer) *QScrollBar {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QScrollBarC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QScrollBar{cthis}
}

// /usr/include/qt/QtWidgets/qscrollbar.h:60
// index:1
// void QScrollBar(Qt::Orientation, class QWidget *)
func NewQScrollBar_1(arg0 int, parent unsafe.Pointer) *QScrollBar {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QScrollBarC2EN2Qt11OrientationEP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, &arg0, parent)
	gopp.ErrPrint(err, rv)
	return &QScrollBar{cthis}
}

// /usr/include/qt/QtWidgets/qscrollbar.h:61
// index:0
// virtual
// void ~QScrollBar()
func DeleteQScrollBar(*QScrollBar) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QScrollBarD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollbar.h:63
// index:0
// virtual
// QSize sizeHint()
func (this *QScrollBar) SizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QScrollBar8sizeHintEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollbar.h:64
// index:0
// virtual
// bool event(class QEvent *)
func (this *QScrollBar) Event(event unsafe.Pointer) {
	// 0: (, QEvent * event), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QScrollBar5eventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.cthis, event)
	gopp.ErrPrint(err, rv)
}

//  body block end
