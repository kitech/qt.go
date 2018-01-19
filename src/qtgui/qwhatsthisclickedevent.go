//  header block begin
// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>
package qtgui

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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtGui/qevent.h:713
// index:0
// void QWhatsThisClickedEvent(const class QString &)
func NewQWhatsThisClickedEvent(href unsafe.Pointer) *QWhatsThisClickedEvent {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN22QWhatsThisClickedEventC2ERK7QString", ffiqt.FFI_TYPE_VOID, cthis, href)
	gopp.ErrPrint(err, rv)
	return &QWhatsThisClickedEvent{cthis}
}

// /usr/include/qt/QtGui/qevent.h:714
// index:0
// virtual
// void ~QWhatsThisClickedEvent()
func DeleteQWhatsThisClickedEvent(*QWhatsThisClickedEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN22QWhatsThisClickedEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:716
// index:0
// inline
// QString href()
func (this *QWhatsThisClickedEvent) Href() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QWhatsThisClickedEvent4hrefEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
