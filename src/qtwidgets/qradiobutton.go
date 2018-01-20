//  header block begin
// /usr/include/qt/QtWidgets/qradiobutton.h
// #include <qradiobutton.h>
// #include <QtWidgets>
package qtwidgets

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
type QRadioButton struct {
	*QAbstractButton
}

func (this *QRadioButton) GetCthis() unsafe.Pointer {
	return this.QAbstractButton.GetCthis()
}

// /usr/include/qt/QtWidgets/qradiobutton.h:56
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QRadioButton) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QRadioButton10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qradiobutton.h:59
// index:0
// void QRadioButton(class QWidget *)
func NewQRadioButton(parent unsafe.Pointer) *QRadioButton {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QRadioButtonC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQRadioButtonFromPointer(cthis)
	return gothis
}
func NewQRadioButtonFromPointer(cthis unsafe.Pointer) *QRadioButton {
	bcthis0 := NewQAbstractButtonFromPointer(cthis)
	return &QRadioButton{bcthis0}
}

// /usr/include/qt/QtWidgets/qradiobutton.h:60
// index:1
// void QRadioButton(const class QString &, class QWidget *)
func NewQRadioButton_1(text unsafe.Pointer, parent unsafe.Pointer) *QRadioButton {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QRadioButtonC2ERK7QStringP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, text, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQRadioButtonFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qradiobutton.h:61
// index:0
// virtual
// void ~QRadioButton()
func DeleteQRadioButton(*QRadioButton) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QRadioButtonD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qradiobutton.h:63
// index:0
// virtual
// QSize sizeHint()
func (this *QRadioButton) SizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QRadioButton8sizeHintEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qradiobutton.h:64
// index:0
// virtual
// QSize minimumSizeHint()
func (this *QRadioButton) MinimumSizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QRadioButton15minimumSizeHintEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qradiobutton.h:67
// index:0
// virtual
// bool event(class QEvent *)
func (this *QRadioButton) Event(e unsafe.Pointer) {
	// 0: (, e QEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QRadioButton5eventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qradiobutton.h:68
// index:0
// virtual
// bool hitButton(const class QPoint &)
func (this *QRadioButton) HitButton(arg0 unsafe.Pointer) {
	// 0: (, const QPoint & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QRadioButton9hitButtonERK6QPoint", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qradiobutton.h:69
// index:0
// virtual
// void paintEvent(class QPaintEvent *)
func (this *QRadioButton) PaintEvent(arg0 unsafe.Pointer) {
	// 0: (, QPaintEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QRadioButton10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qradiobutton.h:70
// index:0
// virtual
// void mouseMoveEvent(class QMouseEvent *)
func (this *QRadioButton) MouseMoveEvent(arg0 unsafe.Pointer) {
	// 0: (, QMouseEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QRadioButton14mouseMoveEventEP11QMouseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qradiobutton.h:71
// index:0
// void initStyleOption(class QStyleOptionButton *)
func (this *QRadioButton) InitStyleOption(button unsafe.Pointer) {
	// 0: (, button QStyleOptionButton *), (button)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QRadioButton15initStyleOptionEP18QStyleOptionButton", ffiqt.FFI_TYPE_VOID, this.GetCthis(), button)
	gopp.ErrPrint(err, rv)
}

//  body block end
