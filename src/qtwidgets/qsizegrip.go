//  header block begin
// /usr/include/qt/QtWidgets/qsizegrip.h
// #include <qsizegrip.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 19
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
type QSizeGrip struct {
	*QWidget
}

func (this *QSizeGrip) GetCthis() unsafe.Pointer {
	return this.QWidget.GetCthis()
}
func NewQSizeGripFromPointer(cthis unsafe.Pointer) *QSizeGrip {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QSizeGrip{bcthis0}
}

// /usr/include/qt/QtWidgets/qsizegrip.h:53
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QSizeGrip) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSizeGrip10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qsizegrip.h:55
// index:0
// Public
// void QSizeGrip(class QWidget *)
func NewQSizeGrip(parent unsafe.Pointer) *QSizeGrip {
	cthis := qtrt.Calloc(1, 256) // 48
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSizeGripC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQSizeGripFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qsizegrip.h:56
// index:0
// Public virtual
// void ~QSizeGrip()
func DeleteQSizeGrip(*QSizeGrip) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSizeGripD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsizegrip.h:58
// index:0
// Public virtual
// QSize sizeHint()
func (this *QSizeGrip) SizeHint() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSizeGrip8sizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qsizegrip.h:59
// index:0
// Public virtual
// void setVisible(_Bool)
func (this *QSizeGrip) SetVisible(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSizeGrip10setVisibleEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsizegrip.h:62
// index:0
// Protected virtual
// void paintEvent(class QPaintEvent *)
func (this *QSizeGrip) PaintEvent(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSizeGrip10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsizegrip.h:63
// index:0
// Protected virtual
// void mousePressEvent(class QMouseEvent *)
func (this *QSizeGrip) MousePressEvent(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSizeGrip15mousePressEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsizegrip.h:64
// index:0
// Protected virtual
// void mouseMoveEvent(class QMouseEvent *)
func (this *QSizeGrip) MouseMoveEvent(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSizeGrip14mouseMoveEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsizegrip.h:65
// index:0
// Protected virtual
// void mouseReleaseEvent(class QMouseEvent *)
func (this *QSizeGrip) MouseReleaseEvent(mouseEvent unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSizeGrip17mouseReleaseEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), mouseEvent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsizegrip.h:66
// index:0
// Protected virtual
// void moveEvent(class QMoveEvent *)
func (this *QSizeGrip) MoveEvent(moveEvent unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSizeGrip9moveEventEP10QMoveEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), moveEvent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsizegrip.h:67
// index:0
// Protected virtual
// void showEvent(class QShowEvent *)
func (this *QSizeGrip) ShowEvent(showEvent unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSizeGrip9showEventEP10QShowEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), showEvent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsizegrip.h:68
// index:0
// Protected virtual
// void hideEvent(class QHideEvent *)
func (this *QSizeGrip) HideEvent(hideEvent unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSizeGrip9hideEventEP10QHideEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), hideEvent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsizegrip.h:69
// index:0
// Protected virtual
// bool eventFilter(class QObject *, class QEvent *)
func (this *QSizeGrip) EventFilter(arg0 unsafe.Pointer, arg1 unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSizeGrip11eventFilterEP7QObjectP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0, arg1)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qsizegrip.h:70
// index:0
// Protected virtual
// bool event(class QEvent *)
func (this *QSizeGrip) Event(arg0 unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSizeGrip5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
