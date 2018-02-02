package qtgui

// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>

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
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"

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
	*QDropEvent
}

func (this *QDragMoveEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QDropEvent.GetCthis()
	}
}
func (this *QDragMoveEvent) SetCthis(cthis unsafe.Pointer) {
	this.QDropEvent = NewQDropEventFromPointer(cthis)
}
func NewQDragMoveEventFromPointer(cthis unsafe.Pointer) *QDragMoveEvent {
	bcthis0 := NewQDropEventFromPointer(cthis)
	return &QDragMoveEvent{bcthis0}
}
func (*QDragMoveEvent) NewFromPointer(cthis unsafe.Pointer) *QDragMoveEvent {
	return NewQDragMoveEventFromPointer(cthis)
}

// /usr/include/qt/QtGui/qevent.h:644
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QDragMoveEvent()
func DeleteQDragMoveEvent(this *QDragMoveEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDragMoveEventD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qevent.h:646
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QRect answerRect()
func (this *QDragMoveEvent) AnswerRect() *qtcore.QRect /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDragMoveEvent10answerRectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:648
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void accept()
func (this *QDragMoveEvent) Accept() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDragMoveEvent6acceptEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:651
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void accept(const QRect &)
func (this *QDragMoveEvent) Accept_1(r *qtcore.QRect) {
	var convArg0 = r.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDragMoveEvent6acceptERK5QRect", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:649
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void ignore()
func (this *QDragMoveEvent) Ignore() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDragMoveEvent6ignoreEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:652
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void ignore(const QRect &)
func (this *QDragMoveEvent) Ignore_1(r *qtcore.QRect) {
	var convArg0 = r.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDragMoveEvent6ignoreERK5QRect", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
