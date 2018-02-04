package qtwidgets

// /usr/include/qt/QtWidgets/qradiobutton.h
// #include <qradiobutton.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 30
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
func (this *QRadioButton) InheritEvent(f func(e *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// bool hitButton(const class QPoint &)
func (this *QRadioButton) InheritHitButton(f func(arg0 *qtcore.QPoint) bool) {
	qtrt.SetAllInheritCallback(this, "hitButton", f)
}

// void paintEvent(class QPaintEvent *)
func (this *QRadioButton) InheritPaintEvent(f func(arg0 *qtgui.QPaintEvent /*777 QPaintEvent **/)) {
	qtrt.SetAllInheritCallback(this, "paintEvent", f)
}

// void mouseMoveEvent(class QMouseEvent *)
func (this *QRadioButton) InheritMouseMoveEvent(f func(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/)) {
	qtrt.SetAllInheritCallback(this, "mouseMoveEvent", f)
}

// void initStyleOption(class QStyleOptionButton *)
func (this *QRadioButton) InheritInitStyleOption(f func(button *QStyleOptionButton /*777 QStyleOptionButton **/)) {
	qtrt.SetAllInheritCallback(this, "initStyleOption", f)
}

type QRadioButton struct {
	*QAbstractButton
}

func (this *QRadioButton) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAbstractButton.GetCthis()
	}
}
func (this *QRadioButton) SetCthis(cthis unsafe.Pointer) {
	this.QAbstractButton = NewQAbstractButtonFromPointer(cthis)
}
func NewQRadioButtonFromPointer(cthis unsafe.Pointer) *QRadioButton {
	bcthis0 := NewQAbstractButtonFromPointer(cthis)
	return &QRadioButton{bcthis0}
}
func (*QRadioButton) NewFromPointer(cthis unsafe.Pointer) *QRadioButton {
	return NewQRadioButtonFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qradiobutton.h:56
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QRadioButton) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QRadioButton10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qradiobutton.h:59
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QRadioButton(QWidget *)
func NewQRadioButton(parent *QWidget /*777 QWidget **/) *QRadioButton {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QRadioButtonC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQRadioButtonFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qradiobutton.h:60
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QRadioButton(const QString &, QWidget *)
func NewQRadioButton_1(text *qtcore.QString, parent *QWidget /*777 QWidget **/) *QRadioButton {
	var convArg0 = text.GetCthis()
	var convArg1 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QRadioButtonC2ERK7QStringP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQRadioButtonFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qradiobutton.h:61
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QRadioButton()
func DeleteQRadioButton(this *QRadioButton) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QRadioButtonD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qradiobutton.h:63
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint()
func (this *QRadioButton) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QRadioButton8sizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qradiobutton.h:64
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize minimumSizeHint()
func (this *QRadioButton) MinimumSizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QRadioButton15minimumSizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qradiobutton.h:67
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QRadioButton) Event(e *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = e.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QRadioButton5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qradiobutton.h:68
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool hitButton(const QPoint &)
func (this *QRadioButton) HitButton(arg0 *qtcore.QPoint) bool {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QRadioButton9hitButtonERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qradiobutton.h:69
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)
func (this *QRadioButton) PaintEvent(arg0 *qtgui.QPaintEvent /*777 QPaintEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QRadioButton10paintEventEP11QPaintEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qradiobutton.h:70
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseMoveEvent(QMouseEvent *)
func (this *QRadioButton) MouseMoveEvent(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QRadioButton14mouseMoveEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qradiobutton.h:71
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void initStyleOption(QStyleOptionButton *)
func (this *QRadioButton) InitStyleOption(button *QStyleOptionButton /*777 QStyleOptionButton **/) {
	var convArg0 = button.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QRadioButton15initStyleOptionEP18QStyleOptionButton", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
