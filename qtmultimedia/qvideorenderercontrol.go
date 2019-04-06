package qtmultimedia

// /usr/include/qt/QtMultimedia/qvideorenderercontrol.h
// #include <qvideorenderercontrol.h>
// #include <QtMultimedia>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 8
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtnetwork"

//  ext block end

//  body block begin

/*

 */
type QVideoRendererControl struct {
	*QMediaControl
}
type QVideoRendererControl_ITF interface {
	QMediaControl_ITF
	QVideoRendererControl_PTR() *QVideoRendererControl
}

func (ptr *QVideoRendererControl) QVideoRendererControl_PTR() *QVideoRendererControl { return ptr }

func (this *QVideoRendererControl) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QMediaControl.GetCthis()
	}
}
func (this *QVideoRendererControl) SetCthis(cthis unsafe.Pointer) {
	this.QMediaControl = NewQMediaControlFromPointer(cthis)
}
func NewQVideoRendererControlFromPointer(cthis unsafe.Pointer) *QVideoRendererControl {
	bcthis0 := NewQMediaControlFromPointer(cthis)
	return &QVideoRendererControl{bcthis0}
}
func (*QVideoRendererControl) NewFromPointer(cthis unsafe.Pointer) *QVideoRendererControl {
	return NewQVideoRendererControlFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qvideorenderercontrol.h:50
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QVideoRendererControl) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QVideoRendererControl10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qvideorenderercontrol.h:53
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QVideoRendererControl()

/*

 */
func DeleteQVideoRendererControl(this *QVideoRendererControl) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QVideoRendererControlD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qvideorenderercontrol.h:55
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QAbstractVideoSurface * surface() const

/*
Returns the surface a video producer renders to.

See also setSurface().
*/
func (this *QVideoRendererControl) Surface() *QAbstractVideoSurface /*777 QAbstractVideoSurface **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QVideoRendererControl7surfaceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQAbstractVideoSurfaceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qvideorenderercontrol.h:56
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setSurface(QAbstractVideoSurface *)

/*
Sets the surface a video producer renders to.

See also surface().
*/
func (this *QVideoRendererControl) SetSurface(surface QAbstractVideoSurface_ITF /*777 QAbstractVideoSurface **/) {
	var convArg0 unsafe.Pointer
	if surface != nil && surface.QAbstractVideoSurface_PTR() != nil {
		convArg0 = surface.QAbstractVideoSurface_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QVideoRendererControl10setSurfaceEP21QAbstractVideoSurface", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qvideorenderercontrol.h:59
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QVideoRendererControl(QObject *)

/*
Constructs a new video renderer media end point with the given parent.
*/
func (*QVideoRendererControl) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QVideoRendererControl {
	return NewQVideoRendererControl(parent)
}
func NewQVideoRendererControl(parent qtcore.QObject_ITF /*777 QObject **/) *QVideoRendererControl {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QVideoRendererControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVideoRendererControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QVideoRendererControl")
	return gothis
}

// /usr/include/qt/QtMultimedia/qvideorenderercontrol.h:59
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QVideoRendererControl(QObject *)

/*
Constructs a new video renderer media end point with the given parent.
*/
func (*QVideoRendererControl) NewForInheritp() *QVideoRendererControl {
	return NewQVideoRendererControlp()
}
func NewQVideoRendererControlp() *QVideoRendererControl {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN21QVideoRendererControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVideoRendererControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QVideoRendererControl")
	return gothis
}

//  body block end

//  keep block begin

func init_unused_11915() {
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
	if false {
		qtgui.KeepMe()
	}
	if false {
		qtnetwork.KeepMe()
	}
}

//  keep block end
