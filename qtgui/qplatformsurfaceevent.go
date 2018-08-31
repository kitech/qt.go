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
type QPlatformSurfaceEvent struct {
	*qtcore.QEvent
}
type QPlatformSurfaceEvent_ITF interface {
	qtcore.QEvent_ITF
	QPlatformSurfaceEvent_PTR() *QPlatformSurfaceEvent
}

func (ptr *QPlatformSurfaceEvent) QPlatformSurfaceEvent_PTR() *QPlatformSurfaceEvent { return ptr }

func (this *QPlatformSurfaceEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QEvent.GetCthis()
	}
}
func (this *QPlatformSurfaceEvent) SetCthis(cthis unsafe.Pointer) {
	this.QEvent = qtcore.NewQEventFromPointer(cthis)
}
func NewQPlatformSurfaceEventFromPointer(cthis unsafe.Pointer) *QPlatformSurfaceEvent {
	bcthis0 := qtcore.NewQEventFromPointer(cthis)
	return &QPlatformSurfaceEvent{bcthis0}
}
func (*QPlatformSurfaceEvent) NewFromPointer(cthis unsafe.Pointer) *QPlatformSurfaceEvent {
	return NewQPlatformSurfaceEventFromPointer(cthis)
}

// /usr/include/qt/QtGui/qevent.h:451
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QPlatformSurfaceEvent(QPlatformSurfaceEvent::SurfaceEventType)

/*

 */
func (*QPlatformSurfaceEvent) NewForInherit(surfaceEventType int) *QPlatformSurfaceEvent {
	return NewQPlatformSurfaceEvent(surfaceEventType)
}
func NewQPlatformSurfaceEvent(surfaceEventType int) *QPlatformSurfaceEvent {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QPlatformSurfaceEventC2ENS_16SurfaceEventTypeE", qtrt.FFI_TYPE_POINTER, surfaceEventType)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPlatformSurfaceEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPlatformSurfaceEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:452
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QPlatformSurfaceEvent()

/*

 */
func DeleteQPlatformSurfaceEvent(this *QPlatformSurfaceEvent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QPlatformSurfaceEventD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qevent.h:454
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QPlatformSurfaceEvent::SurfaceEventType surfaceEventType() const

/*

 */
func (this *QPlatformSurfaceEvent) SurfaceEventType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QPlatformSurfaceEvent16surfaceEventTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
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
