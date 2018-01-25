package qtwidgets

// /usr/include/qt/QtWidgets/qfocusframe.h
// #include <qfocusframe.h>
// #include <QtWidgets>

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
	*QWidget
}

func (this *QFocusFrame) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QWidget.GetCthis()
	}
}
func (this *QFocusFrame) SetCthis(cthis unsafe.Pointer) {
	this.QWidget = NewQWidgetFromPointer(cthis)
}
func NewQFocusFrameFromPointer(cthis unsafe.Pointer) *QFocusFrame {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QFocusFrame{bcthis0}
}
func (*QFocusFrame) NewFromPointer(cthis unsafe.Pointer) *QFocusFrame {
	return NewQFocusFrameFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qfocusframe.h:54
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QFocusFrame) MetaObject() *qtcore.QMetaObject /*444 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFocusFrame10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qfocusframe.h:56
// index:0
// Public
// void QFocusFrame(class QWidget *)
func NewQFocusFrame(parent *QWidget /*444 QWidget **/) *QFocusFrame {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFocusFrameC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQFocusFrameFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qfocusframe.h:57
// index:0
// Public virtual
// void ~QFocusFrame()
func DeleteQFocusFrame(*QFocusFrame) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFocusFrameD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfocusframe.h:59
// index:0
// Public
// void setWidget(class QWidget *)
func (this *QFocusFrame) SetWidget(widget *QWidget /*444 QWidget **/) {
	var convArg0 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFocusFrame9setWidgetEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfocusframe.h:60
// index:0
// Public
// QWidget * widget()
func (this *QFocusFrame) Widget() *QWidget /*444 QWidget **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFocusFrame6widgetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qfocusframe.h:63
// index:0
// Protected virtual
// bool event(class QEvent *)
func (this *QFocusFrame) Event(e *qtcore.QEvent /*444 QEvent **/) bool {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFocusFrame5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qfocusframe.h:65
// index:0
// Protected virtual
// bool eventFilter(class QObject *, class QEvent *)
func (this *QFocusFrame) EventFilter(arg0 *qtcore.QObject /*444 QObject **/, arg1 *qtcore.QEvent /*444 QEvent **/) bool {
	var convArg0 = arg0.GetCthis()
	var convArg1 = arg1.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFocusFrame11eventFilterEP7QObjectP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qfocusframe.h:66
// index:0
// Protected virtual
// void paintEvent(class QPaintEvent *)
func (this *QFocusFrame) PaintEvent(arg0 *qtgui.QPaintEvent /*444 QPaintEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFocusFrame10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfocusframe.h:67
// index:0
// Protected
// void initStyleOption(class QStyleOption *)
func (this *QFocusFrame) InitStyleOption(option *QStyleOption /*444 QStyleOption **/) {
	var convArg0 = option.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFocusFrame15initStyleOptionEP12QStyleOption", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
