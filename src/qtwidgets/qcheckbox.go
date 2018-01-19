//  header block begin
// /usr/include/qt/QtWidgets/qcheckbox.h
// #include <qcheckbox.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 48
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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qcheckbox.h:56
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QCheckBox) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QCheckBox10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcheckbox.h:61
// index:0
// void QCheckBox(class QWidget *)
func NewQCheckBox(parent unsafe.Pointer) *QCheckBox {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QCheckBoxC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QCheckBox{cthis}
}

// /usr/include/qt/QtWidgets/qcheckbox.h:62
// index:1
// void QCheckBox(const class QString &, class QWidget *)
func NewQCheckBox_1(text unsafe.Pointer, parent unsafe.Pointer) *QCheckBox {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QCheckBoxC2ERK7QStringP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, text, parent)
	gopp.ErrPrint(err, rv)
	return &QCheckBox{cthis}
}

// /usr/include/qt/QtWidgets/qcheckbox.h:63
// index:0
// virtual
// void ~QCheckBox()
func DeleteQCheckBox(*QCheckBox) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QCheckBoxD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcheckbox.h:65
// index:0
// virtual
// QSize sizeHint()
func (this *QCheckBox) SizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QCheckBox8sizeHintEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcheckbox.h:66
// index:0
// virtual
// QSize minimumSizeHint()
func (this *QCheckBox) MinimumSizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QCheckBox15minimumSizeHintEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcheckbox.h:68
// index:0
// void setTristate(_Bool)
func (this *QCheckBox) SetTristate(y bool) {
	// 0: (, bool y), (&y)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QCheckBox11setTristateEb", ffiqt.FFI_TYPE_VOID, this.cthis, &y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcheckbox.h:69
// index:0
// bool isTristate()
func (this *QCheckBox) IsTristate() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QCheckBox10isTristateEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcheckbox.h:71
// index:0
// Qt::CheckState checkState()
func (this *QCheckBox) CheckState() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QCheckBox10checkStateEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcheckbox.h:72
// index:0
// void setCheckState(Qt::CheckState)
func (this *QCheckBox) SetCheckState(state int) {
	// 0: (, Qt::CheckState state), (&state)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QCheckBox13setCheckStateEN2Qt10CheckStateE", ffiqt.FFI_TYPE_VOID, this.cthis, &state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcheckbox.h:75
// index:0
// void stateChanged(int)
func (this *QCheckBox) StateChanged(arg0 int) {
	// 0: (, int arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QCheckBox12stateChangedEi", ffiqt.FFI_TYPE_VOID, this.cthis, &arg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
