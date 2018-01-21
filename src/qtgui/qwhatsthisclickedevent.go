package qtgui

// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 3
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
type QWhatsThisClickedEvent struct {
	*qtcore.QEvent
}

func (this *QWhatsThisClickedEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QEvent.GetCthis()
	}
}
func NewQWhatsThisClickedEventFromPointer(cthis unsafe.Pointer) *QWhatsThisClickedEvent {
	bcthis0 := qtcore.NewQEventFromPointer(cthis)
	return &QWhatsThisClickedEvent{bcthis0}
}

// /usr/include/qt/QtGui/qevent.h:713
// index:0
// Public
// void QWhatsThisClickedEvent(const class QString &)
func NewQWhatsThisClickedEvent(href *qtcore.QString) *QWhatsThisClickedEvent {
	cthis := qtrt.Calloc(1, 256) // 32
	var convArg0 = href.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN22QWhatsThisClickedEventC2ERK7QString", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQWhatsThisClickedEventFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:714
// index:0
// Public virtual
// void ~QWhatsThisClickedEvent()
func DeleteQWhatsThisClickedEvent(*QWhatsThisClickedEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN22QWhatsThisClickedEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:716
// index:0
// Public inline
// QString href()
func (this *QWhatsThisClickedEvent) Href() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QWhatsThisClickedEvent4hrefEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

//  body block end
