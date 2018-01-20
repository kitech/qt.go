//  header block begin
// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 11
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
type QInputMethodEvent struct {
	*qtcore.QEvent
}

func (this *QInputMethodEvent) GetCthis() unsafe.Pointer {
	return this.QEvent.GetCthis()
}

// /usr/include/qt/QtGui/qevent.h:555
// index:0
// void QInputMethodEvent()
func NewQInputMethodEvent() *QInputMethodEvent {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QInputMethodEventC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQInputMethodEventFromPointer(cthis)
	return gothis
}
func NewQInputMethodEventFromPointer(cthis unsafe.Pointer) *QInputMethodEvent {
	bcthis0 := qtcore.NewQEventFromPointer(cthis)
	return &QInputMethodEvent{bcthis0}
}

// /usr/include/qt/QtGui/qevent.h:557
// index:0
// virtual
// void ~QInputMethodEvent()
func DeleteQInputMethodEvent(*QInputMethodEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QInputMethodEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:559
// index:0
// void setCommitString(const class QString &, int, int)
func (this *QInputMethodEvent) SetCommitString(commitString unsafe.Pointer, replaceFrom int, replaceLength int) {
	// 0: (, commitString const QString &, replaceFrom int, replaceLength int), (commitString, &replaceFrom, &replaceLength)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QInputMethodEvent15setCommitStringERK7QStringii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), commitString, &replaceFrom, &replaceLength)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:560
// index:0
// inline
// const QList<QInputMethodEvent::Attribute> & attributes()
func (this *QInputMethodEvent) Attributes() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QInputMethodEvent10attributesEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:561
// index:0
// inline
// const QString & preeditString()
func (this *QInputMethodEvent) PreeditString() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QInputMethodEvent13preeditStringEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:563
// index:0
// inline
// const QString & commitString()
func (this *QInputMethodEvent) CommitString() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QInputMethodEvent12commitStringEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:564
// index:0
// inline
// int replacementStart()
func (this *QInputMethodEvent) ReplacementStart() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QInputMethodEvent16replacementStartEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:565
// index:0
// inline
// int replacementLength()
func (this *QInputMethodEvent) ReplacementLength() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QInputMethodEvent17replacementLengthEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
