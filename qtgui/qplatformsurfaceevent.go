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
// extern C begin: 0
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
// size 24
type QPlatformSurfaceEvent struct {
	*qtcore.QEvent
}
type QPlatformSurfaceEvent_ITF interface {
	qtcore.QEvent_ITF
	QPlatformSurfaceEvent_PTR() *QPlatformSurfaceEvent
}

func (ptr *QPlatformSurfaceEvent) QPlatformSurfaceEvent_PTR() *QPlatformSurfaceEvent { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
// ignore GetCthis for 1 base
func QPlatformSurfaceEventFromptr(cthis Voidptr) *QPlatformSurfaceEvent {
	bcthis0 := qtcore.QEventFromptr(cthis)
	return &QPlatformSurfaceEvent{bcthis0}
}
func (*QPlatformSurfaceEvent) Fromptr(cthis Voidptr) *QPlatformSurfaceEvent {
	return QPlatformSurfaceEventFromptr(cthis)
}

func DeleteQPlatformSurfaceEvent(this *QPlatformSurfaceEvent) {
	rv, err := qtrt.Qtcc1(0, "_ZN21QPlatformSurfaceEventD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint2(err, rv)
	this.SetCthis(nil)
}

/*


 */
type QPlatformSurfaceEvent__SurfaceEventType = int

//
const QPlatformSurfaceEvent__SurfaceCreated QPlatformSurfaceEvent__SurfaceEventType = 0

//
const QPlatformSurfaceEvent__SurfaceAboutToBeDestroyed QPlatformSurfaceEvent__SurfaceEventType = 1

func (this *QPlatformSurfaceEvent) SurfaceEventTypeItemName(val int) string {
	switch val {
	case QPlatformSurfaceEvent__SurfaceCreated: // 0
		return "SurfaceCreated"
	case QPlatformSurfaceEvent__SurfaceAboutToBeDestroyed: // 1
		return "SurfaceAboutToBeDestroyed"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QPlatformSurfaceEvent_SurfaceEventTypeItemName(val int) string {
	var nilthis *QPlatformSurfaceEvent
	return nilthis.SurfaceEventTypeItemName(val)
}

//  body block end

//  keep block begin

func init_unused_10083() {
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
