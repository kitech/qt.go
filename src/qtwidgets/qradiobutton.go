//  header block begin
// /usr/include/qt/QtWidgets/qradiobutton.h
// #include <qradiobutton.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 29
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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qradiobutton.h:56
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QRadioButton) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QRadioButton10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qradiobutton.h:59
// index:0
// void QRadioButton(class QWidget *)
func NewQRadioButton(parent unsafe.Pointer) *QRadioButton {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QRadioButtonC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QRadioButton{cthis}
}

// /usr/include/qt/QtWidgets/qradiobutton.h:60
// index:1
// void QRadioButton(const class QString &, class QWidget *)
func NewQRadioButton_1(text unsafe.Pointer, parent unsafe.Pointer) *QRadioButton {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QRadioButtonC2ERK7QStringP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, text, parent)
	gopp.ErrPrint(err, rv)
	return &QRadioButton{cthis}
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
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QRadioButton8sizeHintEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qradiobutton.h:64
// index:0
// virtual
// QSize minimumSizeHint()
func (this *QRadioButton) MinimumSizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QRadioButton15minimumSizeHintEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
