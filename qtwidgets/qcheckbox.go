package qtwidgets

// /usr/include/qt/QtWidgets/qcheckbox.h
// #include <qcheckbox.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 55
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

// bool event(class QEvent *)
func (this *QCheckBox) InheritEvent(f func(e *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// bool hitButton(const class QPoint &)
func (this *QCheckBox) InheritHitButton(f func(pos *qtcore.QPoint) bool) {
	qtrt.SetAllInheritCallback(this, "hitButton", f)
}

// void checkStateSet()
func (this *QCheckBox) InheritCheckStateSet(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "checkStateSet", f)
}

// void nextCheckState()
func (this *QCheckBox) InheritNextCheckState(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "nextCheckState", f)
}

// void paintEvent(class QPaintEvent *)
func (this *QCheckBox) InheritPaintEvent(f func(arg0 *qtgui.QPaintEvent /*777 QPaintEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "paintEvent", f)
}

// void mouseMoveEvent(class QMouseEvent *)
func (this *QCheckBox) InheritMouseMoveEvent(f func(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseMoveEvent", f)
}

// void initStyleOption(class QStyleOptionButton *)
func (this *QCheckBox) InheritInitStyleOption(f func(option *QStyleOptionButton /*777 QStyleOptionButton **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "initStyleOption", f)
}

type QCheckBox struct {
	*QAbstractButton
}

func (this *QCheckBox) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAbstractButton.GetCthis()
	}
}
func (this *QCheckBox) SetCthis(cthis unsafe.Pointer) {
	this.QAbstractButton = NewQAbstractButtonFromPointer(cthis)
}
func NewQCheckBoxFromPointer(cthis unsafe.Pointer) *QCheckBox {
	bcthis0 := NewQAbstractButtonFromPointer(cthis)
	return &QCheckBox{bcthis0}
}
func (*QCheckBox) NewFromPointer(cthis unsafe.Pointer) *QCheckBox {
	return NewQCheckBoxFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qcheckbox.h:56
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QCheckBox) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QCheckBox10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qcheckbox.h:61
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QCheckBox(QWidget *)
func NewQCheckBox(parent *QWidget /*777 QWidget **/) *QCheckBox {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QCheckBoxC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCheckBoxFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qcheckbox.h:62
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QCheckBox(const QString &, QWidget *)
func NewQCheckBox_1(text string, parent *QWidget /*777 QWidget **/) *QCheckBox {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QCheckBoxC2ERK7QStringP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCheckBoxFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qcheckbox.h:63
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QCheckBox()
func DeleteQCheckBox(this *QCheckBox) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QCheckBoxD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qcheckbox.h:65
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint()
func (this *QCheckBox) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QCheckBox8sizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qcheckbox.h:66
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize minimumSizeHint()
func (this *QCheckBox) MinimumSizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QCheckBox15minimumSizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qcheckbox.h:68
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTristate(_Bool)
func (this *QCheckBox) SetTristate(y bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QCheckBox11setTristateEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), y)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcheckbox.h:69
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isTristate()
func (this *QCheckBox) IsTristate() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QCheckBox10isTristateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qcheckbox.h:71
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::CheckState checkState()
func (this *QCheckBox) CheckState() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QCheckBox10checkStateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qcheckbox.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCheckState(Qt::CheckState)
func (this *QCheckBox) SetCheckState(state int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QCheckBox13setCheckStateEN2Qt10CheckStateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), state)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcheckbox.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stateChanged(int)
func (this *QCheckBox) StateChanged(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QCheckBox12stateChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcheckbox.h:78
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QCheckBox) Event(e *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = e.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QCheckBox5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qcheckbox.h:79
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool hitButton(const QPoint &)
func (this *QCheckBox) HitButton(pos *qtcore.QPoint) bool {
	var convArg0 = pos.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QCheckBox9hitButtonERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qcheckbox.h:80
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void checkStateSet()
func (this *QCheckBox) CheckStateSet() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QCheckBox13checkStateSetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcheckbox.h:81
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void nextCheckState()
func (this *QCheckBox) NextCheckState() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QCheckBox14nextCheckStateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcheckbox.h:82
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)
func (this *QCheckBox) PaintEvent(arg0 *qtgui.QPaintEvent /*777 QPaintEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QCheckBox10paintEventEP11QPaintEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcheckbox.h:83
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseMoveEvent(QMouseEvent *)
func (this *QCheckBox) MouseMoveEvent(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QCheckBox14mouseMoveEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcheckbox.h:84
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void initStyleOption(QStyleOptionButton *)
func (this *QCheckBox) InitStyleOption(option *QStyleOptionButton /*777 QStyleOptionButton **/) {
	var convArg0 = option.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QCheckBox15initStyleOptionEP18QStyleOptionButton", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
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
