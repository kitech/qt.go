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
// extern C begin: 5
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
type QDropEvent struct {
	*qtcore.QEvent
}
type QDropEvent_ITF interface {
	qtcore.QEvent_ITF
	QDropEvent_PTR() *QDropEvent
}

func (ptr *QDropEvent) QDropEvent_PTR() *QDropEvent { return ptr }

func (this *QDropEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QEvent.GetCthis()
	}
}
func (this *QDropEvent) SetCthis(cthis unsafe.Pointer) {
	this.QEvent = qtcore.NewQEventFromPointer(cthis)
}
func NewQDropEventFromPointer(cthis unsafe.Pointer) *QDropEvent {
	bcthis0 := qtcore.NewQEventFromPointer(cthis)
	return &QDropEvent{bcthis0}
}
func (*QDropEvent) NewFromPointer(cthis unsafe.Pointer) *QDropEvent {
	return NewQDropEventFromPointer(cthis)
}

//  body block end

//  keep block begin

func init_unused_10685() {
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
