package qtmultimedia

// /usr/include/qt/QtMultimedia/qcameraviewfindersettingscontrol.h
// #include <qcameraviewfindersettingscontrol.h>
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
type QCameraViewfinderSettingsControl struct {
	*QMediaControl
}
type QCameraViewfinderSettingsControl_ITF interface {
	QMediaControl_ITF
	QCameraViewfinderSettingsControl_PTR() *QCameraViewfinderSettingsControl
}

func (ptr *QCameraViewfinderSettingsControl) QCameraViewfinderSettingsControl_PTR() *QCameraViewfinderSettingsControl {
	return ptr
}

func (this *QCameraViewfinderSettingsControl) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QMediaControl.GetCthis()
	}
}
func (this *QCameraViewfinderSettingsControl) SetCthis(cthis unsafe.Pointer) {
	this.QMediaControl = NewQMediaControlFromPointer(cthis)
}
func NewQCameraViewfinderSettingsControlFromPointer(cthis unsafe.Pointer) *QCameraViewfinderSettingsControl {
	bcthis0 := NewQMediaControlFromPointer(cthis)
	return &QCameraViewfinderSettingsControl{bcthis0}
}
func (*QCameraViewfinderSettingsControl) NewFromPointer(cthis unsafe.Pointer) *QCameraViewfinderSettingsControl {
	return NewQCameraViewfinderSettingsControlFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qcameraviewfindersettingscontrol.h:55
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QCameraViewfinderSettingsControl) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK32QCameraViewfinderSettingsControl10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qcameraviewfindersettingscontrol.h:66
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QCameraViewfinderSettingsControl()

/*

 */
func DeleteQCameraViewfinderSettingsControl(this *QCameraViewfinderSettingsControl) {
	rv, err := qtrt.InvokeQtFunc6("_ZN32QCameraViewfinderSettingsControlD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qcameraviewfindersettingscontrol.h:68
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool isViewfinderParameterSupported(QCameraViewfinderSettingsControl::ViewfinderParameter) const

/*
Returns true if configuration of viewfinder parameter is supported by camera backend.
*/
func (this *QCameraViewfinderSettingsControl) IsViewfinderParameterSupported(parameter int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK32QCameraViewfinderSettingsControl30isViewfinderParameterSupportedENS_19ViewfinderParameterE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), parameter)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qcameraviewfindersettingscontrol.h:69
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [16] QVariant viewfinderParameter(QCameraViewfinderSettingsControl::ViewfinderParameter) const

/*
Returns the value of viewfinder parameter.

See also setViewfinderParameter().
*/
func (this *QCameraViewfinderSettingsControl) ViewfinderParameter(parameter int) *qtcore.QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK32QCameraViewfinderSettingsControl19viewfinderParameterENS_19ViewfinderParameterE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), parameter)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtMultimedia/qcameraviewfindersettingscontrol.h:70
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setViewfinderParameter(QCameraViewfinderSettingsControl::ViewfinderParameter, const QVariant &)

/*
Set the prefferred value of viewfinder parameter.

Calling this while the camera is active may result in the camera being stopped and reloaded. If video recording is in progress, this call may be ignored.

If an unsupported parameter is specified the camera may fail to load, or the setting may be ignored.

Viewfinder parameters may also depend on other camera settings, especially in video capture mode. If camera configuration conflicts with viewfinder settings, the camara configuration is usually preferred.

See also viewfinderParameter().
*/
func (this *QCameraViewfinderSettingsControl) SetViewfinderParameter(parameter int, value qtcore.QVariant_ITF) {
	var convArg1 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg1 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN32QCameraViewfinderSettingsControl22setViewfinderParameterENS_19ViewfinderParameterERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), parameter, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameraviewfindersettingscontrol.h:73
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QCameraViewfinderSettingsControl(QObject *)

/*
Constructs a camera viewfinder control object with parent.
*/
func NewQCameraViewfinderSettingsControl(parent qtcore.QObject_ITF /*777 QObject **/) *QCameraViewfinderSettingsControl {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN32QCameraViewfinderSettingsControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCameraViewfinderSettingsControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QCameraViewfinderSettingsControl")
	return gothis
}

// /usr/include/qt/QtMultimedia/qcameraviewfindersettingscontrol.h:73
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QCameraViewfinderSettingsControl(QObject *)

/*
Constructs a camera viewfinder control object with parent.
*/
func NewQCameraViewfinderSettingsControl__() *QCameraViewfinderSettingsControl {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN32QCameraViewfinderSettingsControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCameraViewfinderSettingsControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QCameraViewfinderSettingsControl")
	return gothis
}

/*

 */
type QCameraViewfinderSettingsControl__ViewfinderParameter = int

// Viewfinder resolution, QSize.
const QCameraViewfinderSettingsControl__Resolution QCameraViewfinderSettingsControl__ViewfinderParameter = 0

// Pixel aspect ratio, QSize as in QVideoSurfaceFormat::pixelAspectRatio
const QCameraViewfinderSettingsControl__PixelAspectRatio QCameraViewfinderSettingsControl__ViewfinderParameter = 1

// Minimum viewfinder frame rate, qreal
const QCameraViewfinderSettingsControl__MinimumFrameRate QCameraViewfinderSettingsControl__ViewfinderParameter = 2

// Maximum viewfinder frame rate, qreal
const QCameraViewfinderSettingsControl__MaximumFrameRate QCameraViewfinderSettingsControl__ViewfinderParameter = 3

// Viewfinder pixel format, QVideoFrame::PixelFormat
const QCameraViewfinderSettingsControl__PixelFormat QCameraViewfinderSettingsControl__ViewfinderParameter = 4

//
const QCameraViewfinderSettingsControl__UserParameter QCameraViewfinderSettingsControl__ViewfinderParameter = 1000

func (this *QCameraViewfinderSettingsControl) ViewfinderParameterItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QCameraViewfinderSettingsControl_ViewfinderParameterItemName(val int) string {
	var nilthis *QCameraViewfinderSettingsControl
	return nilthis.ViewfinderParameterItemName(val)
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
