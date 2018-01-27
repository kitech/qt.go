package qtwidgets

// /usr/include/qt/QtWidgets/qcheckbox.h
// #include <qcheckbox.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 55
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
// Public virtual
// const QMetaObject * metaObject()
func (this *QCheckBox) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QCheckBox10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qcheckbox.h:61
// index:0
// Public
// void QCheckBox(QWidget *)
func NewQCheckBox(parent *QWidget /*777 QWidget **/) *QCheckBox {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QCheckBoxC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQCheckBoxFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qcheckbox.h:62
// index:1
// Public
// void QCheckBox(const QString &, QWidget *)
func NewQCheckBox_1(text *qtcore.QString, parent *QWidget /*777 QWidget **/) *QCheckBox {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg0 = text.GetCthis()
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QCheckBoxC2ERK7QStringP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQCheckBoxFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qcheckbox.h:63
// index:0
// Public virtual
// void ~QCheckBox()
func DeleteQCheckBox(*QCheckBox) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QCheckBoxD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcheckbox.h:65
// index:0
// Public virtual
// QSize sizeHint()
func (this *QCheckBox) SizeHint() *qtcore.QSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QCheckBox8sizeHintEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qcheckbox.h:66
// index:0
// Public virtual
// QSize minimumSizeHint()
func (this *QCheckBox) MinimumSizeHint() *qtcore.QSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QCheckBox15minimumSizeHintEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qcheckbox.h:68
// index:0
// Public
// void setTristate(bool)
func (this *QCheckBox) SetTristate(y bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QCheckBox11setTristateEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcheckbox.h:69
// index:0
// Public
// bool isTristate()
func (this *QCheckBox) IsTristate() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QCheckBox10isTristateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qcheckbox.h:71
// index:0
// Public
// Qt::CheckState checkState()
func (this *QCheckBox) CheckState() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QCheckBox10checkStateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qcheckbox.h:72
// index:0
// Public
// void setCheckState(Qt::CheckState)
func (this *QCheckBox) SetCheckState(state int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QCheckBox13setCheckStateEN2Qt10CheckStateE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcheckbox.h:75
// index:0
// Public
// void stateChanged(int)
func (this *QCheckBox) StateChanged(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QCheckBox12stateChangedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcheckbox.h:78
// index:0
// Protected virtual
// bool event(QEvent *)
func (this *QCheckBox) Event(e *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QCheckBox5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qcheckbox.h:79
// index:0
// Protected virtual
// bool hitButton(const QPoint &)
func (this *QCheckBox) HitButton(pos *qtcore.QPoint) bool {
	var convArg0 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QCheckBox9hitButtonERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qcheckbox.h:80
// index:0
// Protected virtual
// void checkStateSet()
func (this *QCheckBox) CheckStateSet() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QCheckBox13checkStateSetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcheckbox.h:81
// index:0
// Protected virtual
// void nextCheckState()
func (this *QCheckBox) NextCheckState() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QCheckBox14nextCheckStateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcheckbox.h:82
// index:0
// Protected virtual
// void paintEvent(QPaintEvent *)
func (this *QCheckBox) PaintEvent(arg0 *qtgui.QPaintEvent /*777 QPaintEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QCheckBox10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcheckbox.h:83
// index:0
// Protected virtual
// void mouseMoveEvent(QMouseEvent *)
func (this *QCheckBox) MouseMoveEvent(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QCheckBox14mouseMoveEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcheckbox.h:84
// index:0
// Protected
// void initStyleOption(QStyleOptionButton *)
func (this *QCheckBox) InitStyleOption(option *QStyleOptionButton /*777 QStyleOptionButton **/) {
	var convArg0 = option.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QCheckBox15initStyleOptionEP18QStyleOptionButton", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
