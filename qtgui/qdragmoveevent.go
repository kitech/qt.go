package qtgui

// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 13
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

/*

 */
type QDragMoveEvent struct {
	*QDropEvent
}
type QDragMoveEvent_ITF interface {
	QDropEvent_ITF
	QDragMoveEvent_PTR() *QDragMoveEvent
}

func (ptr *QDragMoveEvent) QDragMoveEvent_PTR() *QDragMoveEvent { return ptr }

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

//  body block end

//  keep block begin

func init_unused_10687() {
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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
