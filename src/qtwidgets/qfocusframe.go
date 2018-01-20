//  header block begin
// /usr/include/qt/QtWidgets/qfocusframe.h
// #include <qfocusframe.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 54
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
	*QWidget
}

func (this *QFocusFrame) GetCthis() unsafe.Pointer {
	return this.QWidget.GetCthis()
}

// /usr/include/qt/QtWidgets/qfocusframe.h:54
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QFocusFrame) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFocusFrame10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfocusframe.h:56
// index:0
// void QFocusFrame(class QWidget *)
func NewQFocusFrame(parent unsafe.Pointer) *QFocusFrame {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFocusFrameC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQFocusFrameFromPointer(cthis)
	return gothis
}
func NewQFocusFrameFromPointer(cthis unsafe.Pointer) *QFocusFrame {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QFocusFrame{bcthis0}
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
	// 0: (, widget QWidget *), (widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFocusFrame9setWidgetEP7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfocusframe.h:60
// index:0
// QWidget * widget()
func (this *QFocusFrame) Widget() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFocusFrame6widgetEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfocusframe.h:63
// index:0
// virtual
// bool event(class QEvent *)
func (this *QFocusFrame) Event(e unsafe.Pointer) {
	// 0: (, e QEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFocusFrame5eventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfocusframe.h:65
// index:0
// virtual
// bool eventFilter(class QObject *, class QEvent *)
func (this *QFocusFrame) EventFilter(arg0 unsafe.Pointer, arg1 unsafe.Pointer) {
	// 0: (, QObject * arg0, QEvent * arg1), (arg0, arg1)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFocusFrame11eventFilterEP7QObjectP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0, arg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfocusframe.h:66
// index:0
// virtual
// void paintEvent(class QPaintEvent *)
func (this *QFocusFrame) PaintEvent(arg0 unsafe.Pointer) {
	// 0: (, QPaintEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFocusFrame10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfocusframe.h:67
// index:0
// void initStyleOption(class QStyleOption *)
func (this *QFocusFrame) InitStyleOption(option unsafe.Pointer) {
	// 0: (, option QStyleOption *), (option)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFocusFrame15initStyleOptionEP12QStyleOption", ffiqt.FFI_TYPE_VOID, this.GetCthis(), option)
	gopp.ErrPrint(err, rv)
}

//  body block end
