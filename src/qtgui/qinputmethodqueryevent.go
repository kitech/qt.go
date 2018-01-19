//  header block begin
// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 8
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"

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
}

//  ext block end

//  body block begin
type QInputMethodQueryEvent struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtGui/qevent.h:582
// index:0
// virtual
// void ~QInputMethodQueryEvent()
func DeleteQInputMethodQueryEvent(*QInputMethodQueryEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN22QInputMethodQueryEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:584
// index:0
// inline
// Qt::InputMethodQueries queries()
func (this *QInputMethodQueryEvent) Queries() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QInputMethodQueryEvent7queriesEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:586
// index:0
// void setValue(Qt::InputMethodQuery, const class QVariant &)
func (this *QInputMethodQueryEvent) SetValue(query int, value unsafe.Pointer) {
	// 0: (, Qt::InputMethodQuery query, const QVariant & value), (&query, value)
	rv, err := ffiqt.InvokeQtFunc6("_ZN22QInputMethodQueryEvent8setValueEN2Qt16InputMethodQueryERK8QVariant", ffiqt.FFI_TYPE_VOID, this.cthis, &query, value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:587
// index:0
// QVariant value(Qt::InputMethodQuery)
func (this *QInputMethodQueryEvent) Value(query int) {
	// 0: (, Qt::InputMethodQuery query), (&query)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QInputMethodQueryEvent5valueEN2Qt16InputMethodQueryE", ffiqt.FFI_TYPE_VOID, this.cthis, &query)
	gopp.ErrPrint(err, rv)
}

//  body block end
