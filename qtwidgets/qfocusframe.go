package qtwidgets

// /usr/include/qt/QtWidgets/qfocusframe.h
// #include <qfocusframe.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 51
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"

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
// bool event(class QEvent *)
func (this *QFocusFrame) InheritEvent(f func(e *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// bool eventFilter(class QObject *, class QEvent *)
func (this *QFocusFrame) InheritEventFilter(f func(arg0 *qtcore.QObject /*777 QObject **/, arg1 *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "eventFilter", f)
}

// void paintEvent(class QPaintEvent *)
func (this *QFocusFrame) InheritPaintEvent(f func(arg0 *qtgui.QPaintEvent /*777 QPaintEvent **/)) {
	qtrt.SetAllInheritCallback(this, "paintEvent", f)
}

// void initStyleOption(class QStyleOption *)
func (this *QFocusFrame) InheritInitStyleOption(f func(option *QStyleOption /*777 QStyleOption **/)) {
	qtrt.SetAllInheritCallback(this, "initStyleOption", f)
}

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
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QFocusFrame) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFocusFrame10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qfocusframe.h:56
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QFocusFrame(QWidget *)
func NewQFocusFrame(parent *QWidget /*777 QWidget **/) *QFocusFrame {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFocusFrameC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQFocusFrameFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qfocusframe.h:57
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QFocusFrame()
func DeleteQFocusFrame(this *QFocusFrame) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFocusFrameD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qfocusframe.h:59
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWidget(QWidget *)
func (this *QFocusFrame) SetWidget(widget *QWidget /*777 QWidget **/) {
	var convArg0 = widget.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFocusFrame9setWidgetEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfocusframe.h:60
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * widget()
func (this *QFocusFrame) Widget() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFocusFrame6widgetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qfocusframe.h:63
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QFocusFrame) Event(e *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = e.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFocusFrame5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qfocusframe.h:65
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool eventFilter(QObject *, QEvent *)
func (this *QFocusFrame) EventFilter(arg0 *qtcore.QObject /*777 QObject **/, arg1 *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = arg0.GetCthis()
	var convArg1 = arg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFocusFrame11eventFilterEP7QObjectP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qfocusframe.h:66
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)
func (this *QFocusFrame) PaintEvent(arg0 *qtgui.QPaintEvent /*777 QPaintEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFocusFrame10paintEventEP11QPaintEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfocusframe.h:67
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void initStyleOption(QStyleOption *)
func (this *QFocusFrame) InitStyleOption(option *QStyleOption /*777 QStyleOption **/) {
	var convArg0 = option.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFocusFrame15initStyleOptionEP12QStyleOption", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
