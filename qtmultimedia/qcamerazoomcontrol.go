package qtmultimedia

// /usr/include/qt/QtMultimedia/qcamerazoomcontrol.h
// #include <qcamerazoomcontrol.h>
// #include <QtMultimedia>

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
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtnetwork"

//  ext block end

//  body block begin

/*

 */
type QCameraZoomControl struct {
	*QMediaControl
}
type QCameraZoomControl_ITF interface {
	QMediaControl_ITF
	QCameraZoomControl_PTR() *QCameraZoomControl
}

func (ptr *QCameraZoomControl) QCameraZoomControl_PTR() *QCameraZoomControl { return ptr }

func (this *QCameraZoomControl) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QMediaControl.GetCthis()
	}
}
func (this *QCameraZoomControl) SetCthis(cthis unsafe.Pointer) {
	this.QMediaControl = NewQMediaControlFromPointer(cthis)
}
func NewQCameraZoomControlFromPointer(cthis unsafe.Pointer) *QCameraZoomControl {
	bcthis0 := NewQMediaControlFromPointer(cthis)
	return &QCameraZoomControl{bcthis0}
}
func (*QCameraZoomControl) NewFromPointer(cthis unsafe.Pointer) *QCameraZoomControl {
	return NewQCameraZoomControlFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qcamerazoomcontrol.h:53
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QCameraZoomControl) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QCameraZoomControl10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qcamerazoomcontrol.h:56
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QCameraZoomControl()

/*

 */
func DeleteQCameraZoomControl(this *QCameraZoomControl) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCameraZoomControlD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qcamerazoomcontrol.h:58
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] qreal maximumOpticalZoom() const

/*
Returns the maximum optical zoom value, or 1.0 if optical zoom is not supported.
*/
func (this *QCameraZoomControl) MaximumOpticalZoom() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QCameraZoomControl18maximumOpticalZoomEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtMultimedia/qcamerazoomcontrol.h:59
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] qreal maximumDigitalZoom() const

/*
Returns the maximum digital zoom value, or 1.0 if digital zoom is not supported.
*/
func (this *QCameraZoomControl) MaximumDigitalZoom() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QCameraZoomControl18maximumDigitalZoomEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtMultimedia/qcamerazoomcontrol.h:61
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] qreal requestedOpticalZoom() const

/*
Return the requested optical zoom value.
*/
func (this *QCameraZoomControl) RequestedOpticalZoom() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QCameraZoomControl20requestedOpticalZoomEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtMultimedia/qcamerazoomcontrol.h:62
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] qreal requestedDigitalZoom() const

/*
Return the requested digital zoom value.
*/
func (this *QCameraZoomControl) RequestedDigitalZoom() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QCameraZoomControl20requestedDigitalZoomEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtMultimedia/qcamerazoomcontrol.h:63
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] qreal currentOpticalZoom() const

/*
Return the current optical zoom value.
*/
func (this *QCameraZoomControl) CurrentOpticalZoom() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QCameraZoomControl18currentOpticalZoomEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtMultimedia/qcamerazoomcontrol.h:64
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] qreal currentDigitalZoom() const

/*
Return the current digital zoom value.
*/
func (this *QCameraZoomControl) CurrentDigitalZoom() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QCameraZoomControl18currentDigitalZoomEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtMultimedia/qcamerazoomcontrol.h:66
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void zoomTo(qreal, qreal)

/*
Sets optical and digital zoom values.

Zooming can be asynchronous with value changes reported with currentDigitalZoomChanged() and currentOpticalZoomChanged() signals.

The backend should expect and correctly handle frequent zoomTo() calls during zoom animations or slider movements.
*/
func (this *QCameraZoomControl) ZoomTo(optical float64, digital float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCameraZoomControl6zoomToEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), optical, digital)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcamerazoomcontrol.h:69
// index:0
// Public Visibility=Default Availability=Available
// [-2] void maximumOpticalZoomChanged(qreal)

/*
Signal emitted when the maximum supported optical zoom value changed.

The maximum supported zoom value can depend on other camera settings, like focusing mode.
*/
func (this *QCameraZoomControl) MaximumOpticalZoomChanged(arg0 float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCameraZoomControl25maximumOpticalZoomChangedEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcamerazoomcontrol.h:70
// index:0
// Public Visibility=Default Availability=Available
// [-2] void maximumDigitalZoomChanged(qreal)

/*
Signal emitted when the maximum supported digital zoom value changed.

The maximum supported zoom value can depend on other camera settings, like capture mode or resolution.
*/
func (this *QCameraZoomControl) MaximumDigitalZoomChanged(arg0 float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCameraZoomControl25maximumDigitalZoomChangedEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcamerazoomcontrol.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void requestedOpticalZoomChanged(qreal)

/*
Signal emitted when the requested optical zoom value changed.
*/
func (this *QCameraZoomControl) RequestedOpticalZoomChanged(opticalZoom float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCameraZoomControl27requestedOpticalZoomChangedEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), opticalZoom)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcamerazoomcontrol.h:73
// index:0
// Public Visibility=Default Availability=Available
// [-2] void requestedDigitalZoomChanged(qreal)

/*
Signal emitted when the requested digital zoom value changed.
*/
func (this *QCameraZoomControl) RequestedDigitalZoomChanged(digitalZoom float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCameraZoomControl27requestedDigitalZoomChangedEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), digitalZoom)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcamerazoomcontrol.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void currentOpticalZoomChanged(qreal)

/*
Signal emitted when the current optical zoom value changed.
*/
func (this *QCameraZoomControl) CurrentOpticalZoomChanged(opticalZoom float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCameraZoomControl25currentOpticalZoomChangedEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), opticalZoom)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcamerazoomcontrol.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void currentDigitalZoomChanged(qreal)

/*
Signal emitted when the current digital zoom value changed.
*/
func (this *QCameraZoomControl) CurrentDigitalZoomChanged(digitalZoom float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCameraZoomControl25currentDigitalZoomChangedEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), digitalZoom)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcamerazoomcontrol.h:78
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QCameraZoomControl(QObject *)

/*
Constructs a camera zoom control object with parent.
*/
func (*QCameraZoomControl) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QCameraZoomControl {
	return NewQCameraZoomControl(parent)
}
func NewQCameraZoomControl(parent qtcore.QObject_ITF /*777 QObject **/) *QCameraZoomControl {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCameraZoomControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCameraZoomControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QCameraZoomControl")
	return gothis
}

// /usr/include/qt/QtMultimedia/qcamerazoomcontrol.h:78
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QCameraZoomControl(QObject *)

/*
Constructs a camera zoom control object with parent.
*/
func (*QCameraZoomControl) NewForInherit__() *QCameraZoomControl {
	return NewQCameraZoomControl__()
}
func NewQCameraZoomControl__() *QCameraZoomControl {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCameraZoomControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCameraZoomControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QCameraZoomControl")
	return gothis
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
	if false {
		qtgui.KeepMe()
	}
	if false {
		qtnetwork.KeepMe()
	}
}

//  keep block end
