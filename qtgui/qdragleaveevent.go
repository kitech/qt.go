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
// extern C begin: 2
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
type QDragLeaveEvent struct {
	*qtcore.QEvent
}
type QDragLeaveEvent_ITF interface {
	qtcore.QEvent_ITF
	QDragLeaveEvent_PTR() *QDragLeaveEvent
}

func (ptr *QDragLeaveEvent) QDragLeaveEvent_PTR() *QDragLeaveEvent { return ptr }

func (this *QDragLeaveEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QEvent.GetCthis()
	}
}
func (this *QDragLeaveEvent) SetCthis(cthis unsafe.Pointer) {
	this.QEvent = qtcore.NewQEventFromPointer(cthis)
}
func NewQDragLeaveEventFromPointer(cthis unsafe.Pointer) *QDragLeaveEvent {
	bcthis0 := qtcore.NewQEventFromPointer(cthis)
	return &QDragLeaveEvent{bcthis0}
}
func (*QDragLeaveEvent) NewFromPointer(cthis unsafe.Pointer) *QDragLeaveEvent {
	return NewQDragLeaveEventFromPointer(cthis)
}

//  body block end

//  keep block begin

func init_unused_10691() {
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
