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
// extern C begin: 7
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
type QDragEnterEvent struct {
	*QDragMoveEvent
}
type QDragEnterEvent_ITF interface {
	QDragMoveEvent_ITF
	QDragEnterEvent_PTR() *QDragEnterEvent
}

func (ptr *QDragEnterEvent) QDragEnterEvent_PTR() *QDragEnterEvent { return ptr }

func (this *QDragEnterEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QDragMoveEvent.GetCthis()
	}
}
func (this *QDragEnterEvent) SetCthis(cthis unsafe.Pointer) {
	this.QDragMoveEvent = NewQDragMoveEventFromPointer(cthis)
}
func NewQDragEnterEventFromPointer(cthis unsafe.Pointer) *QDragEnterEvent {
	bcthis0 := NewQDragMoveEventFromPointer(cthis)
	return &QDragEnterEvent{bcthis0}
}
func (*QDragEnterEvent) NewFromPointer(cthis unsafe.Pointer) *QDragEnterEvent {
	return NewQDragEnterEventFromPointer(cthis)
}

//  body block end

//  keep block begin

func init_unused_10689() {
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
