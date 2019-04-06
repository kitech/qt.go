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
// extern C begin: 3
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
type QWhatsThisClickedEvent struct {
	*qtcore.QEvent
}
type QWhatsThisClickedEvent_ITF interface {
	qtcore.QEvent_ITF
	QWhatsThisClickedEvent_PTR() *QWhatsThisClickedEvent
}

func (ptr *QWhatsThisClickedEvent) QWhatsThisClickedEvent_PTR() *QWhatsThisClickedEvent { return ptr }

func (this *QWhatsThisClickedEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QEvent.GetCthis()
	}
}
func (this *QWhatsThisClickedEvent) SetCthis(cthis unsafe.Pointer) {
	this.QEvent = qtcore.NewQEventFromPointer(cthis)
}
func NewQWhatsThisClickedEventFromPointer(cthis unsafe.Pointer) *QWhatsThisClickedEvent {
	bcthis0 := qtcore.NewQEventFromPointer(cthis)
	return &QWhatsThisClickedEvent{bcthis0}
}
func (*QWhatsThisClickedEvent) NewFromPointer(cthis unsafe.Pointer) *QWhatsThisClickedEvent {
	return NewQWhatsThisClickedEventFromPointer(cthis)
}

//  body block end

//  keep block begin

func init_unused_10697() {
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
