//  header block begin
// /usr/include/qt/QtWidgets/qsizegrip.h
// #include <qsizegrip.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 18
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
type QSizeGrip struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qsizegrip.h:53
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QSizeGrip) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSizeGrip10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsizegrip.h:55
// index:0
// void QSizeGrip(class QWidget *)
func NewQSizeGrip(parent unsafe.Pointer) *QSizeGrip {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSizeGripC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QSizeGrip{cthis}
}

// /usr/include/qt/QtWidgets/qsizegrip.h:56
// index:0
// virtual
// void ~QSizeGrip()
func DeleteQSizeGrip(*QSizeGrip) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSizeGripD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsizegrip.h:58
// index:0
// virtual
// QSize sizeHint()
func (this *QSizeGrip) SizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSizeGrip8sizeHintEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsizegrip.h:59
// index:0
// virtual
// void setVisible(_Bool)
func (this *QSizeGrip) SetVisible(arg0 bool) {
	// 0: (, bool arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSizeGrip10setVisibleEb", ffiqt.FFI_TYPE_VOID, this.cthis, &arg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
