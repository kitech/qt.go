//  header block begin
// /usr/include/qt/QtWidgets/qcheckbox.h
// #include <qcheckbox.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 56
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
type QCheckBox struct {
	*QAbstractButton
}

func (this *QCheckBox) GetCthis() unsafe.Pointer {
	return this.QAbstractButton.GetCthis()
}
func NewQCheckBoxFromPointer(cthis unsafe.Pointer) *QCheckBox {
	bcthis0 := NewQAbstractButtonFromPointer(cthis)
	return &QCheckBox{bcthis0}
}

// /usr/include/qt/QtWidgets/qcheckbox.h:56
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QCheckBox) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QCheckBox10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qcheckbox.h:61
// index:0
// Public
// void QCheckBox(class QWidget *)
func NewQCheckBox(parent unsafe.Pointer) *QCheckBox {
	cthis := qtrt.Calloc(1, 256) // 48
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QCheckBoxC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQCheckBoxFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qcheckbox.h:62
// index:1
// Public
// void QCheckBox(const class QString &, class QWidget *)
func NewQCheckBox_1(text *qtcore.QString, parent unsafe.Pointer) *QCheckBox {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QCheckBoxC2ERK7QStringP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, convArg0, parent)
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
func (this *QCheckBox) SizeHint() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QCheckBox8sizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qcheckbox.h:66
// index:0
// Public virtual
// QSize minimumSizeHint()
func (this *QCheckBox) MinimumSizeHint() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QCheckBox15minimumSizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qcheckbox.h:68
// index:0
// Public
// void setTristate(_Bool)
func (this *QCheckBox) SetTristate(y bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QCheckBox11setTristateEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcheckbox.h:69
// index:0
// Public
// bool isTristate()
func (this *QCheckBox) IsTristate() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QCheckBox10isTristateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qcheckbox.h:71
// index:0
// Public
// Qt::CheckState checkState()
func (this *QCheckBox) CheckState() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QCheckBox10checkStateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qcheckbox.h:72
// index:0
// Public
// void setCheckState(Qt::CheckState)
func (this *QCheckBox) SetCheckState(state int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QCheckBox13setCheckStateEN2Qt10CheckStateE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcheckbox.h:75
// index:0
// Public
// void stateChanged(int)
func (this *QCheckBox) StateChanged(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QCheckBox12stateChangedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcheckbox.h:78
// index:0
// Protected virtual
// bool event(class QEvent *)
func (this *QCheckBox) Event(e unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QCheckBox5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qcheckbox.h:79
// index:0
// Protected virtual
// bool hitButton(const class QPoint &)
func (this *QCheckBox) HitButton(pos *qtcore.QPoint) interface{} {
	var convArg0 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QCheckBox9hitButtonERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
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
// void paintEvent(class QPaintEvent *)
func (this *QCheckBox) PaintEvent(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QCheckBox10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcheckbox.h:83
// index:0
// Protected virtual
// void mouseMoveEvent(class QMouseEvent *)
func (this *QCheckBox) MouseMoveEvent(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QCheckBox14mouseMoveEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcheckbox.h:84
// index:0
// Protected
// void initStyleOption(class QStyleOptionButton *)
func (this *QCheckBox) InitStyleOption(option unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QCheckBox15initStyleOptionEP18QStyleOptionButton", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), option)
	gopp.ErrPrint(err, rv)
}

//  body block end
