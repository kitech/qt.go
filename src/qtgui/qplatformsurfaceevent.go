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
type QPlatformSurfaceEvent struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtGui/qevent.h:451
// index:0
// void QPlatformSurfaceEvent(enum QPlatformSurfaceEvent::SurfaceEventType)
func NewQPlatformSurfaceEvent(surfaceEventType int) *QPlatformSurfaceEvent {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QPlatformSurfaceEventC2ENS_16SurfaceEventTypeE", ffiqt.FFI_TYPE_VOID, cthis, &surfaceEventType)
	gopp.ErrPrint(err, rv)
	return &QPlatformSurfaceEvent{cthis}
}

// /usr/include/qt/QtGui/qevent.h:452
// index:0
// virtual
// void ~QPlatformSurfaceEvent()
func DeleteQPlatformSurfaceEvent(*QPlatformSurfaceEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QPlatformSurfaceEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:454
// index:0
// inline
// QPlatformSurfaceEvent::SurfaceEventType surfaceEventType()
func (this *QPlatformSurfaceEvent) SurfaceEventType() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QPlatformSurfaceEvent16surfaceEventTypeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
