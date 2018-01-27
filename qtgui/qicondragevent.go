package qtgui

// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 2
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
type QIconDragEvent struct {
	*qtcore.QEvent
}

func (this *QIconDragEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QEvent.GetCthis()
	}
}
func (this *QIconDragEvent) SetCthis(cthis unsafe.Pointer) {
	this.QEvent = qtcore.NewQEventFromPointer(cthis)
}
func NewQIconDragEventFromPointer(cthis unsafe.Pointer) *QIconDragEvent {
	bcthis0 := qtcore.NewQEventFromPointer(cthis)
	return &QIconDragEvent{bcthis0}
}
func (*QIconDragEvent) NewFromPointer(cthis unsafe.Pointer) *QIconDragEvent {
	return NewQIconDragEventFromPointer(cthis)
}

// /usr/include/qt/QtGui/qevent.h:485
// index:0
// Public
// void QIconDragEvent()
func NewQIconDragEvent() *QIconDragEvent {
	cthis := qtrt.Calloc(1, 256) // 24
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QIconDragEventC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQIconDragEventFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:486
// index:0
// Public virtual
// void ~QIconDragEvent()
func DeleteQIconDragEvent(*QIconDragEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QIconDragEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

//  body block end
