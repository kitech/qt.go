// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qradiobutton.h
// #include <qradiobutton.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 30
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

// bool event(QEvent *)
func (this *QRadioButton) InheritEvent(f func(e *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// bool hitButton(const QPoint &)
func (this *QRadioButton) InheritHitButton(f func(arg0 *qtcore.QPoint) bool) {
	qtrt.SetAllInheritCallback(this, "hitButton", f)
}

// void paintEvent(QPaintEvent *)
func (this *QRadioButton) InheritPaintEvent(f func(arg0 *qtgui.QPaintEvent /*777 QPaintEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "paintEvent", f)
}

// void mouseMoveEvent(QMouseEvent *)
func (this *QRadioButton) InheritMouseMoveEvent(f func(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseMoveEvent", f)
}

// void initStyleOption(QStyleOptionButton *)
func (this *QRadioButton) InheritInitStyleOption(f func(button *QStyleOptionButton /*777 QStyleOptionButton **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "initStyleOption", f)
}

/*

 */
type QRadioButton struct {
	*QAbstractButton
}
type QRadioButton_ITF interface {
	QAbstractButton_ITF
	QRadioButton_PTR() *QRadioButton
}

func (ptr *QRadioButton) QRadioButton_PTR() *QRadioButton { return ptr }

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
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QRadioButton) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QRadioButton10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qradiobutton.h:59
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QRadioButton(QWidget *)

/*
Constructs a radio button with the given parent, but with no text or pixmap.

The parent argument is passed on to the QAbstractButton constructor.
*/
func (*QRadioButton) NewForInherit(parent QWidget_ITF /*777 QWidget **/) *QRadioButton {
	return NewQRadioButton(parent)
}
func NewQRadioButton(parent QWidget_ITF /*777 QWidget **/) *QRadioButton {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QRadioButtonC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRadioButtonFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QRadioButton")
	return gothis
}

// /usr/include/qt/QtWidgets/qradiobutton.h:59
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QRadioButton(QWidget *)

/*
Constructs a radio button with the given parent, but with no text or pixmap.

The parent argument is passed on to the QAbstractButton constructor.
*/
func (*QRadioButton) NewForInheritp() *QRadioButton {
	return NewQRadioButtonp()
}
func NewQRadioButtonp() *QRadioButton {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN12QRadioButtonC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRadioButtonFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QRadioButton")
	return gothis
}

// /usr/include/qt/QtWidgets/qradiobutton.h:60
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QRadioButton(const QString &, QWidget *)

/*
Constructs a radio button with the given parent, but with no text or pixmap.

The parent argument is passed on to the QAbstractButton constructor.
*/
func (*QRadioButton) NewForInherit1(text string, parent QWidget_ITF /*777 QWidget **/) *QRadioButton {
	return NewQRadioButton1(text, parent)
}
func NewQRadioButton1(text string, parent QWidget_ITF /*777 QWidget **/) *QRadioButton {
	var tmpArg0 = qtcore.NewQString5(text)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg1 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QRadioButtonC2ERK7QStringP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRadioButtonFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QRadioButton")
	return gothis
}

// /usr/include/qt/QtWidgets/qradiobutton.h:60
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QRadioButton(const QString &, QWidget *)

/*
Constructs a radio button with the given parent, but with no text or pixmap.

The parent argument is passed on to the QAbstractButton constructor.
*/
func (*QRadioButton) NewForInherit1p(text string) *QRadioButton {
	return NewQRadioButton1p(text)
}
func NewQRadioButton1p(text string) *QRadioButton {
	var tmpArg0 = qtcore.NewQString5(text)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN12QRadioButtonC2ERK7QStringP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRadioButtonFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QRadioButton")
	return gothis
}

// /usr/include/qt/QtWidgets/qradiobutton.h:61
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QRadioButton()

/*

 */
func DeleteQRadioButton(this *QRadioButton) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QRadioButtonD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qradiobutton.h:63
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint() const

/*
Reimplemented from QWidget::sizeHint().
*/
func (this *QRadioButton) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QRadioButton8sizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qradiobutton.h:64
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize minimumSizeHint() const

/*
Reimplemented from QWidget::minimumSizeHint().
*/
func (this *QRadioButton) MinimumSizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QRadioButton15minimumSizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qradiobutton.h:67
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QAbstractButton::event().
*/
func (this *QRadioButton) Event(e qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if e != nil && e.QEvent_PTR() != nil {
		convArg0 = e.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QRadioButton5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qradiobutton.h:68
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool hitButton(const QPoint &) const

/*
Reimplemented from QAbstractButton::hitButton().
*/
func (this *QRadioButton) HitButton(arg0 qtcore.QPoint_ITF) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QPoint_PTR() != nil {
		convArg0 = arg0.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QRadioButton9hitButtonERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qradiobutton.h:69
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)

/*
Reimplemented from QAbstractButton::paintEvent().
*/
func (this *QRadioButton) PaintEvent(arg0 qtgui.QPaintEvent_ITF /*777 QPaintEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QPaintEvent_PTR() != nil {
		convArg0 = arg0.QPaintEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QRadioButton10paintEventEP11QPaintEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qradiobutton.h:70
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseMoveEvent(QMouseEvent *)

/*
Reimplemented from QAbstractButton::mouseMoveEvent().
*/
func (this *QRadioButton) MouseMoveEvent(arg0 qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMouseEvent_PTR() != nil {
		convArg0 = arg0.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QRadioButton14mouseMoveEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qradiobutton.h:71
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void initStyleOption(QStyleOptionButton *) const

/*
Initialize option with the values from this QRadioButton. This method is useful for subclasses when they need a QStyleOptionButton, but don't want to fill in all the information themselves.

See also QStyleOption::initFrom().
*/
func (this *QRadioButton) InitStyleOption(button QStyleOptionButton_ITF /*777 QStyleOptionButton **/) {
	var convArg0 unsafe.Pointer
	if button != nil && button.QStyleOptionButton_PTR() != nil {
		convArg0 = button.QStyleOptionButton_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QRadioButton15initStyleOptionEP18QStyleOptionButton", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

//  body block end

//  keep block begin

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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  keep block end
