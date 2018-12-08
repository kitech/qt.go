package qtmultimedia

// /usr/include/qt/QtMultimedia/qcameraexposurecontrol.h
// #include <qcameraexposurecontrol.h>
// #include <QtMultimedia>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 14
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
type QCameraExposureControl struct {
	*QMediaControl
}
type QCameraExposureControl_ITF interface {
	QMediaControl_ITF
	QCameraExposureControl_PTR() *QCameraExposureControl
}

func (ptr *QCameraExposureControl) QCameraExposureControl_PTR() *QCameraExposureControl { return ptr }

func (this *QCameraExposureControl) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QMediaControl.GetCthis()
	}
}
func (this *QCameraExposureControl) SetCthis(cthis unsafe.Pointer) {
	this.QMediaControl = NewQMediaControlFromPointer(cthis)
}
func NewQCameraExposureControlFromPointer(cthis unsafe.Pointer) *QCameraExposureControl {
	bcthis0 := NewQMediaControlFromPointer(cthis)
	return &QCameraExposureControl{bcthis0}
}
func (*QCameraExposureControl) NewFromPointer(cthis unsafe.Pointer) *QCameraExposureControl {
	return NewQCameraExposureControlFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qcameraexposurecontrol.h:57
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QCameraExposureControl) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QCameraExposureControl10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qcameraexposurecontrol.h:61
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QCameraExposureControl()

/*

 */
func DeleteQCameraExposureControl(this *QCameraExposureControl) {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QCameraExposureControlD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qcameraexposurecontrol.h:77
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool isParameterSupported(QCameraExposureControl::ExposureParameter) const

/*
Returns true is exposure parameter is supported by backend.

This function was introduced in  Qt 5.0.
*/
func (this *QCameraExposureControl) IsParameterSupported(parameter int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QCameraExposureControl20isParameterSupportedENS_17ExposureParameterE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), parameter)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qcameraexposurecontrol.h:78
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QVariantList supportedParameterRange(QCameraExposureControl::ExposureParameter, bool *) const

/*
Returns the list of supported parameter values;

If the camera supports arbitrary exposure parameter value within the supported range, *continuous is set to true, otherwise *continuous is set to false.

This function was introduced in  Qt 5.0.
*/
func (this *QCameraExposureControl) SupportedParameterRange(parameter int, continuous *bool) *qtcore.QVariantList /*667*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QCameraExposureControl23supportedParameterRangeENS_17ExposureParameterEPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), parameter, continuous)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantListFromPointer(unsafe.Pointer(uintptr(rv))) //5551
	return rv2
}

// /usr/include/qt/QtMultimedia/qcameraexposurecontrol.h:80
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [16] QVariant requestedValue(QCameraExposureControl::ExposureParameter) const

/*
Returns the requested exposure parameter value.

This function was introduced in  Qt 5.0.
*/
func (this *QCameraExposureControl) RequestedValue(parameter int) *qtcore.QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QCameraExposureControl14requestedValueENS_17ExposureParameterE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), parameter)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtMultimedia/qcameraexposurecontrol.h:81
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [16] QVariant actualValue(QCameraExposureControl::ExposureParameter) const

/*
Returns the actual exposure parameter value, or invalid QVariant() if the value is unknown or not supported.

The actual parameter value may differ for the requested one if automatic mode is selected or camera supports only limited set of values within the supported range.

This function was introduced in  Qt 5.0.
*/
func (this *QCameraExposureControl) ActualValue(parameter int) *qtcore.QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QCameraExposureControl11actualValueENS_17ExposureParameterE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), parameter)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtMultimedia/qcameraexposurecontrol.h:82
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool setValue(QCameraExposureControl::ExposureParameter, const QVariant &)

/*
Set the exposure parameter to value. If a null or invalid QVariant is passed, backend should choose the value automatically, and if possible report the actual value to user with QCameraExposureControl::actualValue().

Returns true if parameter is supported and value is correct.

This function was introduced in  Qt 5.0.
*/
func (this *QCameraExposureControl) SetValue(parameter int, value qtcore.QVariant_ITF) bool {
	var convArg1 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg1 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN22QCameraExposureControl8setValueENS_17ExposureParameterERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), parameter, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qcameraexposurecontrol.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void requestedValueChanged(int)

/*
Signal emitted when the requested exposure parameter value has changed, usually in result of setValue() call.

This function was introduced in  Qt 5.0.
*/
func (this *QCameraExposureControl) RequestedValueChanged(parameter int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QCameraExposureControl21requestedValueChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), parameter)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameraexposurecontrol.h:86
// index:0
// Public Visibility=Default Availability=Available
// [-2] void actualValueChanged(int)

/*
Signal emitted when the actual exposure parameter value has changed, usually in result of auto exposure algorithms or manual exposure parameter applied.

This function was introduced in  Qt 5.0.
*/
func (this *QCameraExposureControl) ActualValueChanged(parameter int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QCameraExposureControl18actualValueChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), parameter)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameraexposurecontrol.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void parameterRangeChanged(int)

/*
Signal emitted when the supported range of exposure parameter values has changed.

This function was introduced in  Qt 5.0.
*/
func (this *QCameraExposureControl) ParameterRangeChanged(parameter int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QCameraExposureControl21parameterRangeChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), parameter)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameraexposurecontrol.h:90
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QCameraExposureControl(QObject *)

/*
Constructs a camera exposure control object with parent.
*/
func (*QCameraExposureControl) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QCameraExposureControl {
	return NewQCameraExposureControl(parent)
}
func NewQCameraExposureControl(parent qtcore.QObject_ITF /*777 QObject **/) *QCameraExposureControl {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN22QCameraExposureControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCameraExposureControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QCameraExposureControl")
	return gothis
}

// /usr/include/qt/QtMultimedia/qcameraexposurecontrol.h:90
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QCameraExposureControl(QObject *)

/*
Constructs a camera exposure control object with parent.
*/
func (*QCameraExposureControl) NewForInheritp() *QCameraExposureControl {
	return NewQCameraExposureControlp()
}
func NewQCameraExposureControlp() *QCameraExposureControl {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN22QCameraExposureControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCameraExposureControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QCameraExposureControl")
	return gothis
}

/*


This value is only used in the manual flash mode.



This value is only used in the torch flash mode.


*/
type QCameraExposureControl__ExposureParameter = int

// Camera ISO sensitivity, specified as integer value.
const QCameraExposureControl__ISO QCameraExposureControl__ExposureParameter = 0

// Lens aperture is specified as an qreal F number. The supported apertures list can change depending on the focal length, in such a case the exposureParameterRangeChanged() signal is emitted.
const QCameraExposureControl__Aperture QCameraExposureControl__ExposureParameter = 1

// Shutter speed in seconds, specified as qreal.
const QCameraExposureControl__ShutterSpeed QCameraExposureControl__ExposureParameter = 2

// Exposure compensation, specified as qreal EV value.
const QCameraExposureControl__ExposureCompensation QCameraExposureControl__ExposureParameter = 3

//
const QCameraExposureControl__FlashPower QCameraExposureControl__ExposureParameter = 4

// Flash compensation, specified as qreal EV value.
const QCameraExposureControl__FlashCompensation QCameraExposureControl__ExposureParameter = 5

//
const QCameraExposureControl__TorchPower QCameraExposureControl__ExposureParameter = 6

// The relative frame coordinate of the point to use for exposure metering in spot metering mode, specified as a QPointF.
const QCameraExposureControl__SpotMeteringPoint QCameraExposureControl__ExposureParameter = 7

// Camera exposure mode.
const QCameraExposureControl__ExposureMode QCameraExposureControl__ExposureParameter = 8

// Camera metering mode.
const QCameraExposureControl__MeteringMode QCameraExposureControl__ExposureParameter = 9

//
const QCameraExposureControl__ExtendedExposureParameter QCameraExposureControl__ExposureParameter = 1000

func (this *QCameraExposureControl) ExposureParameterItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QCameraExposureControl_ExposureParameterItemName(val int) string {
	var nilthis *QCameraExposureControl
	return nilthis.ExposureParameterItemName(val)
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
