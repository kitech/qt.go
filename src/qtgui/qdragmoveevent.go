//  header block begin
// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 12
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
type QDragMoveEvent struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtGui/qevent.h:644
// index:0
// virtual
// void ~QDragMoveEvent()
func DeleteQDragMoveEvent(*QDragMoveEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDragMoveEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:646
// index:0
// inline
// QRect answerRect()
func (this *QDragMoveEvent) AnswerRect() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDragMoveEvent10answerRectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:648
// index:0
// inline
// void accept()
func (this *QDragMoveEvent) Accept() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDragMoveEvent6acceptEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:651
// index:1
// inline
// void accept(const class QRect &)
func (this *QDragMoveEvent) Accept_1(r unsafe.Pointer) {
	// 1: (, const QRect & r), (r)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDragMoveEvent6acceptERK5QRect", ffiqt.FFI_TYPE_VOID, this.cthis, r)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:649
// index:0
// inline
// void ignore()
func (this *QDragMoveEvent) Ignore() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDragMoveEvent6ignoreEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:652
// index:1
// inline
// void ignore(const class QRect &)
func (this *QDragMoveEvent) Ignore_1(r unsafe.Pointer) {
	// 1: (, const QRect & r), (r)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDragMoveEvent6ignoreERK5QRect", ffiqt.FFI_TYPE_VOID, this.cthis, r)
	gopp.ErrPrint(err, rv)
}

//  body block end
