//  header block begin
// /usr/include/qt/QtCore/qfinalstate.h
// #include <qfinalstate.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 10
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"

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
}

//  ext block end

//  body block begin
type QFinalState struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/qfinalstate.h:52
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QFinalState) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFinalState10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfinalstate.h:54
// index:0
// void QFinalState(class QState *)
func NewQFinalState(parent unsafe.Pointer) *QFinalState {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFinalStateC2EP6QState", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QFinalState{cthis}
}

// /usr/include/qt/QtCore/qfinalstate.h:55
// index:0
// virtual
// void ~QFinalState()
func DeleteQFinalState(*QFinalState) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFinalStateD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

//  body block end
