package qtmultimedia

// /usr/include/qt/QtMultimedia/qcamerafocus.h
// #include <qcamerafocus.h>
// #include <QtMultimedia>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 10
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
type QCameraFocus struct {
	*qtcore.QObject
}
type QCameraFocus_ITF interface {
	qtcore.QObject_ITF
	QCameraFocus_PTR() *QCameraFocus
}

func (ptr *QCameraFocus) QCameraFocus_PTR() *QCameraFocus { return ptr }

func (this *QCameraFocus) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QCameraFocus) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQCameraFocusFromPointer(cthis unsafe.Pointer) *QCameraFocus {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QCameraFocus{bcthis0}
}
func (*QCameraFocus) NewFromPointer(cthis unsafe.Pointer) *QCameraFocus {
	return NewQCameraFocusFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qcamerafocus.h:96
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QCameraFocus) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QCameraFocus10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qcamerafocus.h:125
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isAvailable() const

/*
Returns true if focus related settings are supported by this camera.

You may need to also check if any specific features are supported.
*/
func (this *QCameraFocus) IsAvailable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QCameraFocus11isAvailableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qcamerafocus.h:127
// index:0
// Public Visibility=Default Availability=Available
// [4] QCameraFocus::FocusModes focusMode() const

/*

 */
func (this *QCameraFocus) FocusMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QCameraFocus9focusModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qcamerafocus.h:128
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFocusMode(QCameraFocus::FocusModes)

/*

 */
func (this *QCameraFocus) SetFocusMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QCameraFocus12setFocusModeE6QFlagsINS_9FocusModeEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcamerafocus.h:129
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isFocusModeSupported(QCameraFocus::FocusModes) const

/*
Returns true if the focus mode is supported by camera.
*/
func (this *QCameraFocus) IsFocusModeSupported(mode int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QCameraFocus20isFocusModeSupportedE6QFlagsINS_9FocusModeEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qcamerafocus.h:131
// index:0
// Public Visibility=Default Availability=Available
// [4] QCameraFocus::FocusPointMode focusPointMode() const

/*

 */
func (this *QCameraFocus) FocusPointMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QCameraFocus14focusPointModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qcamerafocus.h:132
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFocusPointMode(QCameraFocus::FocusPointMode)

/*

 */
func (this *QCameraFocus) SetFocusPointMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QCameraFocus17setFocusPointModeENS_14FocusPointModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcamerafocus.h:133
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isFocusPointModeSupported(QCameraFocus::FocusPointMode) const

/*
Returns true if focus point mode is supported.
*/
func (this *QCameraFocus) IsFocusPointModeSupported(arg0 int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QCameraFocus25isFocusPointModeSupportedENS_14FocusPointModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qcamerafocus.h:134
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF customFocusPoint() const

/*

 */
func (this *QCameraFocus) CustomFocusPoint() *qtcore.QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QCameraFocus16customFocusPointEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtMultimedia/qcamerafocus.h:135
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCustomFocusPoint(const QPointF &)

/*

 */
func (this *QCameraFocus) SetCustomFocusPoint(point qtcore.QPointF_ITF) {
	var convArg0 unsafe.Pointer
	if point != nil && point.QPointF_PTR() != nil {
		convArg0 = point.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QCameraFocus19setCustomFocusPointERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcamerafocus.h:137
// index:0
// Public Visibility=Default Availability=Available
// [-2] QCameraFocusZoneList focusZones() const

/*

 */
func (this *QCameraFocus) FocusZones() *QCameraFocusZoneList /*667*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QCameraFocus10focusZonesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCameraFocusZoneListFromPointer(unsafe.Pointer(uintptr(rv))) //5551
	return rv2
}

// /usr/include/qt/QtMultimedia/qcamerafocus.h:139
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal maximumOpticalZoom() const

/*
Returns the maximum optical zoom.

This will be 1.0 on cameras that do not support optical zoom.
*/
func (this *QCameraFocus) MaximumOpticalZoom() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QCameraFocus18maximumOpticalZoomEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtMultimedia/qcamerafocus.h:140
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal maximumDigitalZoom() const

/*
Returns the maximum digital zoom

This will be 1.0 on cameras that do not support digital zoom.
*/
func (this *QCameraFocus) MaximumDigitalZoom() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QCameraFocus18maximumDigitalZoomEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtMultimedia/qcamerafocus.h:141
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal opticalZoom() const

/*

 */
func (this *QCameraFocus) OpticalZoom() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QCameraFocus11opticalZoomEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtMultimedia/qcamerafocus.h:142
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal digitalZoom() const

/*

 */
func (this *QCameraFocus) DigitalZoom() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QCameraFocus11digitalZoomEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtMultimedia/qcamerafocus.h:144
// index:0
// Public Visibility=Default Availability=Available
// [-2] void zoomTo(qreal, qreal)

/*
Set the camera optical and digital zoom values.

Since there may be a physical component to move, the change in zoom value may not be instantaneous.
*/
func (this *QCameraFocus) ZoomTo(opticalZoom float64, digitalZoom float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QCameraFocus6zoomToEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), opticalZoom, digitalZoom)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcamerafocus.h:147
// index:0
// Public Visibility=Default Availability=Available
// [-2] void opticalZoomChanged(qreal)

/*
Signal emitted when optical zoom value changes to new value.

Note: Notifier signal for property opticalZoom.
*/
func (this *QCameraFocus) OpticalZoomChanged(arg0 float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QCameraFocus18opticalZoomChangedEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcamerafocus.h:148
// index:0
// Public Visibility=Default Availability=Available
// [-2] void digitalZoomChanged(qreal)

/*
Signal emitted when digital zoom value changes to new value.

Note: Notifier signal for property digitalZoom.
*/
func (this *QCameraFocus) DigitalZoomChanged(arg0 float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QCameraFocus18digitalZoomChangedEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcamerafocus.h:150
// index:0
// Public Visibility=Default Availability=Available
// [-2] void focusZonesChanged()

/*
This signal is emitted when the set of zones used in autofocusing is changed.

This can change when a zone is focused or loses focus, or new focus zones have been detected.

Note: Notifier signal for property focusZones.
*/
func (this *QCameraFocus) FocusZonesChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QCameraFocus17focusZonesChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcamerafocus.h:152
// index:0
// Public Visibility=Default Availability=Available
// [-2] void maximumOpticalZoomChanged(qreal)

/*
Signal emitted when the maximum supported optical zoom value changed.
*/
func (this *QCameraFocus) MaximumOpticalZoomChanged(arg0 float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QCameraFocus25maximumOpticalZoomChangedEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcamerafocus.h:153
// index:0
// Public Visibility=Default Availability=Available
// [-2] void maximumDigitalZoomChanged(qreal)

/*
Signal emitted when the maximum supported digital zoom value changed.

The maximum supported zoom value can depend on other camera settings, like capture mode or resolution.
*/
func (this *QCameraFocus) MaximumDigitalZoomChanged(arg0 float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QCameraFocus25maximumDigitalZoomChangedEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

func DeleteQCameraFocus(this *QCameraFocus) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QCameraFocusD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

/*


 */
type QCameraFocus__FocusMode = int

//
const QCameraFocus__ManualFocus QCameraFocus__FocusMode = 1

//
const QCameraFocus__HyperfocalFocus QCameraFocus__FocusMode = 2

//
const QCameraFocus__InfinityFocus QCameraFocus__FocusMode = 4

//
const QCameraFocus__AutoFocus QCameraFocus__FocusMode = 8

//
const QCameraFocus__ContinuousFocus QCameraFocus__FocusMode = 16

//
const QCameraFocus__MacroFocus QCameraFocus__FocusMode = 32

/*

 */
type QCameraFocus__FocusPointMode = int

// Automatically select one or multiple focus points.
const QCameraFocus__FocusPointAuto QCameraFocus__FocusPointMode = 0

// Focus to the frame center.
const QCameraFocus__FocusPointCenter QCameraFocus__FocusPointMode = 1

// Focus on faces in the frame.
const QCameraFocus__FocusPointFaceDetection QCameraFocus__FocusPointMode = 2

// Focus to the custom point, defined by QCameraFocus::customFocusPoint property.
const QCameraFocus__FocusPointCustom QCameraFocus__FocusPointMode = 3

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
