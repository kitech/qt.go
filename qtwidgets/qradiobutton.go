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
import "qt.go/cffiqt"
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
// Public virtual
// const QMetaObject * metaObject()
func (this *QRadioButton) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QRadioButton10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qradiobutton.h:59
// index:0
// Public
// void QRadioButton(QWidget *)
func NewQRadioButton(parent *QWidget /*777 QWidget **/) *QRadioButton {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QRadioButtonC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQRadioButtonFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qradiobutton.h:60
// index:1
// Public
// void QRadioButton(const QString &, QWidget *)
func NewQRadioButton_1(text *qtcore.QString, parent *QWidget /*777 QWidget **/) *QRadioButton {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg0 = text.GetCthis()
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QRadioButtonC2ERK7QStringP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQRadioButtonFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qradiobutton.h:61
// index:0
// Public virtual
// void ~QRadioButton()
func DeleteQRadioButton(*QRadioButton) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QRadioButtonD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qradiobutton.h:63
// index:0
// Public virtual
// QSize sizeHint()
func (this *QRadioButton) SizeHint() *qtcore.QSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc7("_ZNK12QRadioButton8sizeHintEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qradiobutton.h:64
// index:0
// Public virtual
// QSize minimumSizeHint()
func (this *QRadioButton) MinimumSizeHint() *qtcore.QSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc7("_ZNK12QRadioButton15minimumSizeHintEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qradiobutton.h:67
// index:0
// Protected virtual
// bool event(QEvent *)
func (this *QRadioButton) Event(e *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QRadioButton5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qradiobutton.h:68
// index:0
// Protected virtual
// bool hitButton(const QPoint &)
func (this *QRadioButton) HitButton(arg0 *qtcore.QPoint) bool {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QRadioButton9hitButtonERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qradiobutton.h:69
// index:0
// Protected virtual
// void paintEvent(QPaintEvent *)
func (this *QRadioButton) PaintEvent(arg0 *qtgui.QPaintEvent /*777 QPaintEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QRadioButton10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qradiobutton.h:70
// index:0
// Protected virtual
// void mouseMoveEvent(QMouseEvent *)
func (this *QRadioButton) MouseMoveEvent(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QRadioButton14mouseMoveEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qradiobutton.h:71
// index:0
// Protected
// void initStyleOption(QStyleOptionButton *)
func (this *QRadioButton) InitStyleOption(button *QStyleOptionButton /*777 QStyleOptionButton **/) {
	var convArg0 = button.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QRadioButton15initStyleOptionEP18QStyleOptionButton", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
