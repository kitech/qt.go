package qtmultimedia

// /usr/include/qt/QtMultimedia/qcameraimageprocessingcontrol.h
// #include <qcameraimageprocessingcontrol.h>
// #include <QtMultimedia>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 15
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
type QCameraImageProcessingControl struct {
	*QMediaControl
}
type QCameraImageProcessingControl_ITF interface {
	QMediaControl_ITF
	QCameraImageProcessingControl_PTR() *QCameraImageProcessingControl
}

func (ptr *QCameraImageProcessingControl) QCameraImageProcessingControl_PTR() *QCameraImageProcessingControl {
	return ptr
}

func (this *QCameraImageProcessingControl) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QMediaControl.GetCthis()
	}
}
func (this *QCameraImageProcessingControl) SetCthis(cthis unsafe.Pointer) {
	this.QMediaControl = NewQMediaControlFromPointer(cthis)
}
func NewQCameraImageProcessingControlFromPointer(cthis unsafe.Pointer) *QCameraImageProcessingControl {
	bcthis0 := NewQMediaControlFromPointer(cthis)
	return &QCameraImageProcessingControl{bcthis0}
}
func (*QCameraImageProcessingControl) NewFromPointer(cthis unsafe.Pointer) *QCameraImageProcessingControl {
	return NewQCameraImageProcessingControlFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qcameraimageprocessingcontrol.h:56
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QCameraImageProcessingControl) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK29QCameraImageProcessingControl10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qcameraimageprocessingcontrol.h:60
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QCameraImageProcessingControl()

/*

 */
func DeleteQCameraImageProcessingControl(this *QCameraImageProcessingControl) {
	rv, err := qtrt.InvokeQtFunc6("_ZN29QCameraImageProcessingControlD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qcameraimageprocessingcontrol.h:79
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool isParameterSupported(QCameraImageProcessingControl::ProcessingParameter) const

/*
Returns true if the camera supports adjusting image processing parameter.

Usually the supported setting is static, but some parameters may not be available depending on other camera settings, like presets. In such case the currently supported parameters should be returned.
*/
func (this *QCameraImageProcessingControl) IsParameterSupported(arg0 int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK29QCameraImageProcessingControl20isParameterSupportedENS_19ProcessingParameterE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qcameraimageprocessingcontrol.h:80
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool isParameterValueSupported(QCameraImageProcessingControl::ProcessingParameter, const QVariant &) const

/*
Returns true if the camera supports setting the image processing parameter value.

It's used only for parameters with a limited set of values, like WhiteBalancePreset.
*/
func (this *QCameraImageProcessingControl) IsParameterValueSupported(parameter int, value qtcore.QVariant_ITF) bool {
	var convArg1 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg1 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK29QCameraImageProcessingControl25isParameterValueSupportedENS_19ProcessingParameterERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), parameter, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qcameraimageprocessingcontrol.h:81
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [16] QVariant parameter(QCameraImageProcessingControl::ProcessingParameter) const

/*
Returns the image processing parameter value.

See also setParameter().
*/
func (this *QCameraImageProcessingControl) Parameter(parameter int) *qtcore.QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK29QCameraImageProcessingControl9parameterENS_19ProcessingParameterE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), parameter)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtMultimedia/qcameraimageprocessingcontrol.h:82
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setParameter(QCameraImageProcessingControl::ProcessingParameter, const QVariant &)

/*
Sets the image processing parameter value. Passing the null or invalid QVariant value allows backend to choose the suitable parameter value.

The valid values range depends on the parameter type. For WhiteBalancePreset the value should be one of QCameraImageProcessing::WhiteBalanceMode values; for Contrast, Saturation, Brightness, Sharpening and Denoising the value should be in [0..1.0] range with invalid QVariant value indicating the default parameter value; for ContrastAdjustment, SaturationAdjustment, BrightnessAdjustment, SharpeningAdjustment and DenoisingAdjustment the value should be in [-1.0..1.0] range with default 0.

See also parameter().
*/
func (this *QCameraImageProcessingControl) SetParameter(parameter int, value qtcore.QVariant_ITF) {
	var convArg1 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg1 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN29QCameraImageProcessingControl12setParameterENS_19ProcessingParameterERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), parameter, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameraimageprocessingcontrol.h:85
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QCameraImageProcessingControl(QObject *)

/*
Constructs an image processing control object with parent.
*/
func NewQCameraImageProcessingControl(parent qtcore.QObject_ITF /*777 QObject **/) *QCameraImageProcessingControl {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN29QCameraImageProcessingControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCameraImageProcessingControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QCameraImageProcessingControl")
	return gothis
}

// /usr/include/qt/QtMultimedia/qcameraimageprocessingcontrol.h:85
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QCameraImageProcessingControl(QObject *)

/*
Constructs an image processing control object with parent.
*/
func NewQCameraImageProcessingControl__() *QCameraImageProcessingControl {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN29QCameraImageProcessingControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCameraImageProcessingControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QCameraImageProcessingControl")
	return gothis
}

/*

 */
type QCameraImageProcessingControl__ProcessingParameter = int

// The white balance preset.
const QCameraImageProcessingControl__WhiteBalancePreset QCameraImageProcessingControl__ProcessingParameter = 0

// Color temperature in K. This value is used when the manual white balance mode is selected.
const QCameraImageProcessingControl__ColorTemperature QCameraImageProcessingControl__ProcessingParameter = 1

// Image contrast.
const QCameraImageProcessingControl__Contrast QCameraImageProcessingControl__ProcessingParameter = 2

// Image saturation.
const QCameraImageProcessingControl__Saturation QCameraImageProcessingControl__ProcessingParameter = 3

// Image brightness.
const QCameraImageProcessingControl__Brightness QCameraImageProcessingControl__ProcessingParameter = 4

// Amount of sharpening applied.
const QCameraImageProcessingControl__Sharpening QCameraImageProcessingControl__ProcessingParameter = 5

// Amount of denoising applied.
const QCameraImageProcessingControl__Denoising QCameraImageProcessingControl__ProcessingParameter = 6

// Image contrast adjustment.
const QCameraImageProcessingControl__ContrastAdjustment QCameraImageProcessingControl__ProcessingParameter = 7

// Image saturation adjustment.
const QCameraImageProcessingControl__SaturationAdjustment QCameraImageProcessingControl__ProcessingParameter = 8

// Image brightness adjustment.
const QCameraImageProcessingControl__BrightnessAdjustment QCameraImageProcessingControl__ProcessingParameter = 9

//
const QCameraImageProcessingControl__SharpeningAdjustment QCameraImageProcessingControl__ProcessingParameter = 10

//
const QCameraImageProcessingControl__DenoisingAdjustment QCameraImageProcessingControl__ProcessingParameter = 11

//
const QCameraImageProcessingControl__ColorFilter QCameraImageProcessingControl__ProcessingParameter = 12

//
const QCameraImageProcessingControl__ExtendedParameter QCameraImageProcessingControl__ProcessingParameter = 1000

func (this *QCameraImageProcessingControl) ProcessingParameterItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QCameraImageProcessingControl_ProcessingParameterItemName(val int) string {
	var nilthis *QCameraImageProcessingControl
	return nilthis.ProcessingParameterItemName(val)
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
