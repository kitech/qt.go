package qtmultimedia

// /usr/include/qt/QtMultimedia/qcamera.h
// #include <qcamera.h>
// #include <QtMultimedia>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 18
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
type QCamera struct {
	*QMediaObject
}
type QCamera_ITF interface {
	QMediaObject_ITF
	QCamera_PTR() *QCamera
}

func (ptr *QCamera) QCamera_PTR() *QCamera { return ptr }

func (this *QCamera) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QMediaObject.GetCthis()
	}
}
func (this *QCamera) SetCthis(cthis unsafe.Pointer) {
	this.QMediaObject = NewQMediaObjectFromPointer(cthis)
}
func NewQCameraFromPointer(cthis unsafe.Pointer) *QCamera {
	bcthis0 := NewQMediaObjectFromPointer(cthis)
	return &QCamera{bcthis0}
}
func (*QCamera) NewFromPointer(cthis unsafe.Pointer) *QCamera {
	return NewQCameraFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qcamera.h:71
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QCamera) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QCamera10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qcamera.h:168
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QCamera(QObject *)

/*
Construct a QCamera with a parent.
*/
func (*QCamera) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QCamera {
	return NewQCamera(parent)
}
func NewQCamera(parent qtcore.QObject_ITF /*777 QObject **/) *QCamera {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QCameraC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCameraFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QCamera")
	return gothis
}

// /usr/include/qt/QtMultimedia/qcamera.h:168
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QCamera(QObject *)

/*
Construct a QCamera with a parent.
*/
func (*QCamera) NewForInheritp() *QCamera {
	return NewQCamerap()
}
func NewQCamerap() *QCamera {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN7QCameraC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCameraFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QCamera")
	return gothis
}

// /usr/include/qt/QtMultimedia/qcamera.h:169
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QCamera(const QByteArray &, QObject *)

/*
Construct a QCamera with a parent.
*/
func (*QCamera) NewForInherit1(deviceName qtcore.QByteArray_ITF, parent qtcore.QObject_ITF /*777 QObject **/) *QCamera {
	return NewQCamera1(deviceName, parent)
}
func NewQCamera1(deviceName qtcore.QByteArray_ITF, parent qtcore.QObject_ITF /*777 QObject **/) *QCamera {
	var convArg0 unsafe.Pointer
	if deviceName != nil && deviceName.QByteArray_PTR() != nil {
		convArg0 = deviceName.QByteArray_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg1 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QCameraC2ERK10QByteArrayP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCameraFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QCamera")
	return gothis
}

// /usr/include/qt/QtMultimedia/qcamera.h:169
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QCamera(const QByteArray &, QObject *)

/*
Construct a QCamera with a parent.
*/
func (*QCamera) NewForInherit1p(deviceName qtcore.QByteArray_ITF) *QCamera {
	return NewQCamera1p(deviceName)
}
func NewQCamera1p(deviceName qtcore.QByteArray_ITF) *QCamera {
	var convArg0 unsafe.Pointer
	if deviceName != nil && deviceName.QByteArray_PTR() != nil {
		convArg0 = deviceName.QByteArray_PTR().GetCthis()
	}
	// arg: 1, QObject *=Pointer, QObject=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN7QCameraC2ERK10QByteArrayP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCameraFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QCamera")
	return gothis
}

// /usr/include/qt/QtMultimedia/qcamera.h:170
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QCamera(const QCameraInfo &, QObject *)

/*
Construct a QCamera with a parent.
*/
func (*QCamera) NewForInherit2(cameraInfo QCameraInfo_ITF, parent qtcore.QObject_ITF /*777 QObject **/) *QCamera {
	return NewQCamera2(cameraInfo, parent)
}
func NewQCamera2(cameraInfo QCameraInfo_ITF, parent qtcore.QObject_ITF /*777 QObject **/) *QCamera {
	var convArg0 unsafe.Pointer
	if cameraInfo != nil && cameraInfo.QCameraInfo_PTR() != nil {
		convArg0 = cameraInfo.QCameraInfo_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg1 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QCameraC2ERK11QCameraInfoP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCameraFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QCamera")
	return gothis
}

// /usr/include/qt/QtMultimedia/qcamera.h:170
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QCamera(const QCameraInfo &, QObject *)

/*
Construct a QCamera with a parent.
*/
func (*QCamera) NewForInherit2p(cameraInfo QCameraInfo_ITF) *QCamera {
	return NewQCamera2p(cameraInfo)
}
func NewQCamera2p(cameraInfo QCameraInfo_ITF) *QCamera {
	var convArg0 unsafe.Pointer
	if cameraInfo != nil && cameraInfo.QCameraInfo_PTR() != nil {
		convArg0 = cameraInfo.QCameraInfo_PTR().GetCthis()
	}
	// arg: 1, QObject *=Pointer, QObject=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN7QCameraC2ERK11QCameraInfoP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCameraFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QCamera")
	return gothis
}

// /usr/include/qt/QtMultimedia/qcamera.h:171
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QCamera(QCamera::Position, QObject *)

/*
Construct a QCamera with a parent.
*/
func (*QCamera) NewForInherit3(position int, parent qtcore.QObject_ITF /*777 QObject **/) *QCamera {
	return NewQCamera3(position, parent)
}
func NewQCamera3(position int, parent qtcore.QObject_ITF /*777 QObject **/) *QCamera {
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg1 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QCameraC2ENS_8PositionEP7QObject", qtrt.FFI_TYPE_POINTER, position, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCameraFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QCamera")
	return gothis
}

// /usr/include/qt/QtMultimedia/qcamera.h:171
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QCamera(QCamera::Position, QObject *)

/*
Construct a QCamera with a parent.
*/
func (*QCamera) NewForInherit3p(position int) *QCamera {
	return NewQCamera3p(position)
}
func NewQCamera3p(position int) *QCamera {
	// arg: 1, QObject *=Pointer, QObject=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN7QCameraC2ENS_8PositionEP7QObject", qtrt.FFI_TYPE_POINTER, position, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCameraFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QCamera")
	return gothis
}

// /usr/include/qt/QtMultimedia/qcamera.h:172
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QCamera()

/*

 */
func DeleteQCamera(this *QCamera) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QCameraD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qcamera.h:176
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString deviceDescription(const QByteArray &)

/*

 */
func (this *QCamera) DeviceDescription(device qtcore.QByteArray_ITF) string {
	var convArg0 unsafe.Pointer
	if device != nil && device.QByteArray_PTR() != nil {
		convArg0 = device.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QCamera17deviceDescriptionERK10QByteArray", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}
func QCamera_DeviceDescription(device qtcore.QByteArray_ITF) string {
	var nilthis *QCamera
	rv := nilthis.DeviceDescription(device)
	return rv
}

// /usr/include/qt/QtMultimedia/qcamera.h:179
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] QMultimedia::AvailabilityStatus availability() const

/*
Reimplemented from QMediaObject::availability().

Returns the availability state of the camera service.
*/
func (this *QCamera) Availability() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QCamera12availabilityEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qcamera.h:181
// index:0
// Public Visibility=Default Availability=Available
// [4] QCamera::State state() const

/*

 */
func (this *QCamera) State() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QCamera5stateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qcamera.h:182
// index:0
// Public Visibility=Default Availability=Available
// [4] QCamera::Status status() const

/*

 */
func (this *QCamera) Status() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QCamera6statusEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qcamera.h:184
// index:0
// Public Visibility=Default Availability=Available
// [4] QCamera::CaptureModes captureMode() const

/*

 */
func (this *QCamera) CaptureMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QCamera11captureModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qcamera.h:185
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isCaptureModeSupported(QCamera::CaptureModes) const

/*
Returns true if the capture mode is suported.
*/
func (this *QCamera) IsCaptureModeSupported(mode int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QCamera22isCaptureModeSupportedE6QFlagsINS_11CaptureModeEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qcamera.h:187
// index:0
// Public Visibility=Default Availability=Available
// [8] QCameraExposure * exposure() const

/*
Returns the camera exposure control object.
*/
func (this *QCamera) Exposure() *QCameraExposure /*777 QCameraExposure **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QCamera8exposureEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQCameraExposureFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qcamera.h:188
// index:0
// Public Visibility=Default Availability=Available
// [8] QCameraFocus * focus() const

/*
Returns the camera focus control object.
*/
func (this *QCamera) Focus() *QCameraFocus /*777 QCameraFocus **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QCamera5focusEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQCameraFocusFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qcamera.h:189
// index:0
// Public Visibility=Default Availability=Available
// [8] QCameraImageProcessing * imageProcessing() const

/*
Returns the camera image processing control object.
*/
func (this *QCamera) ImageProcessing() *QCameraImageProcessing /*777 QCameraImageProcessing **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QCamera15imageProcessingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQCameraImageProcessingFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qcamera.h:193
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setViewfinder(QAbstractVideoSurface *)

/*
Sets the QVideoWidget based camera viewfinder. The previously set viewfinder is detached.
*/
func (this *QCamera) SetViewfinder(surface QAbstractVideoSurface_ITF /*777 QAbstractVideoSurface **/) {
	var convArg0 unsafe.Pointer
	if surface != nil && surface.QAbstractVideoSurface_PTR() != nil {
		convArg0 = surface.QAbstractVideoSurface_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QCamera13setViewfinderEP21QAbstractVideoSurface", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcamera.h:195
// index:0
// Public Visibility=Default Availability=Available
// [8] QCameraViewfinderSettings viewfinderSettings() const

/*
Returns the viewfinder settings being used by the camera.

Settings may change when the camera is started, for example if the viewfinder settings are undefined or if unsupported values are set.

If viewfinder settings are not supported by the camera, it always returns a null QCameraViewfinderSettings object.

This function was introduced in  Qt 5.5.

See also setViewfinderSettings().
*/
func (this *QCamera) ViewfinderSettings() *QCameraViewfinderSettings /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QCamera18viewfinderSettingsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCameraViewfinderSettingsFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCameraViewfinderSettings)
	return rv2
}

// /usr/include/qt/QtMultimedia/qcamera.h:196
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setViewfinderSettings(const QCameraViewfinderSettings &)

/*
Sets the viewfinder settings.

If some parameters are not specified, or null settings are passed, the camera will choose default values.

If the camera is used to capture videos or images, the viewfinder settings might be ignored if they conflict with the capture settings. You can check the actual viewfinder settings once the camera is in the QCamera::ActiveStatus status.

Changing the viewfinder settings while the camera is in the QCamera::ActiveState state may cause the camera to be restarted.

This function was introduced in  Qt 5.5.

See also viewfinderSettings(), supportedViewfinderResolutions(), supportedViewfinderFrameRateRanges(), and supportedViewfinderPixelFormats().
*/
func (this *QCamera) SetViewfinderSettings(settings QCameraViewfinderSettings_ITF) {
	var convArg0 unsafe.Pointer
	if settings != nil && settings.QCameraViewfinderSettings_PTR() != nil {
		convArg0 = settings.QCameraViewfinderSettings_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QCamera21setViewfinderSettingsERK25QCameraViewfinderSettings", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcamera.h:201
// index:0
// Public Visibility=Default Availability=Available
// [8] QList<QSize> supportedViewfinderResolutions(const QCameraViewfinderSettings &) const

/*
Returns a list of supported viewfinder resolutions.

This is a convenience function which retrieves unique resolutions from the supported settings.

If non null viewfinder settings are passed, the returned list is reduced to resolutions supported with partial settings applied.

The camera must be loaded before calling this function, otherwise the returned list is empty.

This function was introduced in  Qt 5.5.

See also QCameraViewfinderSettings::resolution() and setViewfinderSettings().
*/
func (this *QCamera) SupportedViewfinderResolutions(settings QCameraViewfinderSettings_ITF) *qtcore.QSizeList /*lll*/ {
	var convArg0 unsafe.Pointer
	if settings != nil && settings.QCameraViewfinderSettings_PTR() != nil {
		convArg0 = settings.QCameraViewfinderSettings_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QCamera30supportedViewfinderResolutionsERK25QCameraViewfinderSettings", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeListFromPointer(unsafe.Pointer(uintptr(rv))) //5552
	return rv2
}

// /usr/include/qt/QtMultimedia/qcamera.h:201
// index:0
// Public Visibility=Default Availability=Available
// [8] QList<QSize> supportedViewfinderResolutions(const QCameraViewfinderSettings &) const

/*
Returns a list of supported viewfinder resolutions.

This is a convenience function which retrieves unique resolutions from the supported settings.

If non null viewfinder settings are passed, the returned list is reduced to resolutions supported with partial settings applied.

The camera must be loaded before calling this function, otherwise the returned list is empty.

This function was introduced in  Qt 5.5.

See also QCameraViewfinderSettings::resolution() and setViewfinderSettings().
*/
func (this *QCamera) SupportedViewfinderResolutionsp() *qtcore.QSizeList /*lll*/ {
	// arg: 0, const QCameraViewfinderSettings &=LValueReference, QCameraViewfinderSettings=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QCamera30supportedViewfinderResolutionsERK25QCameraViewfinderSettings", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeListFromPointer(unsafe.Pointer(uintptr(rv))) //5552
	return rv2
}

// /usr/include/qt/QtMultimedia/qcamera.h:210
// index:0
// Public Visibility=Default Availability=Available
// [4] QCamera::Error error() const

/*
Returns the error state of the object.
*/
func (this *QCamera) Error() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QCamera5errorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qcamera.h:245
// index:1
// Public Visibility=Default Availability=Available
// [-2] void error(QCamera::Error)

/*
Returns the error state of the object.
*/
func (this *QCamera) Error1(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QCamera5errorENS_5ErrorE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcamera.h:211
// index:0
// Public Visibility=Default Availability=Available
// [8] QString errorString() const

/*
Returns a string describing a camera's error state.
*/
func (this *QCamera) ErrorString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QCamera11errorStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtMultimedia/qcamera.h:213
// index:0
// Public Visibility=Default Availability=Available
// [4] QCamera::LockTypes supportedLocks() const

/*
Returns the lock types, camera supports.
*/
func (this *QCamera) SupportedLocks() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QCamera14supportedLocksEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qcamera.h:214
// index:0
// Public Visibility=Default Availability=Available
// [4] QCamera::LockTypes requestedLocks() const

/*
Returns the requested lock types.
*/
func (this *QCamera) RequestedLocks() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QCamera14requestedLocksEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qcamera.h:216
// index:0
// Public Visibility=Default Availability=Available
// [4] QCamera::LockStatus lockStatus() const

/*
Returns the status of requested camera settings locks.

Note: Getter function for property lockStatus.
*/
func (this *QCamera) LockStatus() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QCamera10lockStatusEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qcamera.h:217
// index:1
// Public Visibility=Default Availability=Available
// [4] QCamera::LockStatus lockStatus(QCamera::LockType) const

/*
Returns the status of requested camera settings locks.

Note: Getter function for property lockStatus.
*/
func (this *QCamera) LockStatus1(lock int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QCamera10lockStatusENS_8LockTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), lock)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qcamera.h:220
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCaptureMode(QCamera::CaptureModes)

/*

 */
func (this *QCamera) SetCaptureMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QCamera14setCaptureModeE6QFlagsINS_11CaptureModeEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcamera.h:222
// index:0
// Public Visibility=Default Availability=Available
// [-2] void load()

/*
Open the camera device. The camera state is changed to QCamera::LoadedStatus.

It's not necessary to explcitly load the camera, unless unless the application have to read the supported camera settings and change the default depending on the camera capabilities.

In all the other cases it's possible to start the camera directly from unloaded state.
*/
func (this *QCamera) Load() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QCamera4loadEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcamera.h:223
// index:0
// Public Visibility=Default Availability=Available
// [-2] void unload()

/*
Close the camera device and deallocate the related resources. The camera state is changed to QCamera::UnloadedStatus.
*/
func (this *QCamera) Unload() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QCamera6unloadEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcamera.h:225
// index:0
// Public Visibility=Default Availability=Available
// [-2] void start()

/*
Starts the camera.

State is changed to QCamera::ActiveState if camera is started successfully, otherwise error() signal is emitted.

While the camera state is changed to QCamera::ActiveState, starting the camera service can be asynchronous with the actual status reported with QCamera::status property.
*/
func (this *QCamera) Start() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QCamera5startEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcamera.h:226
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stop()

/*
Stops the camera. The camera state is changed from QCamera::ActiveState to QCamera::LoadedState.
*/
func (this *QCamera) Stop() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QCamera4stopEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcamera.h:228
// index:0
// Public Visibility=Default Availability=Available
// [-2] void searchAndLock()

/*
Lock all the supported camera settings.
*/
func (this *QCamera) SearchAndLock() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QCamera13searchAndLockEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcamera.h:231
// index:1
// Public Visibility=Default Availability=Available
// [-2] void searchAndLock(QCamera::LockTypes)

/*
Lock all the supported camera settings.
*/
func (this *QCamera) SearchAndLock1(locks int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QCamera13searchAndLockE6QFlagsINS_8LockTypeEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), locks)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcamera.h:229
// index:0
// Public Visibility=Default Availability=Available
// [-2] void unlock()

/*
Unlock all the requested camera locks.
*/
func (this *QCamera) Unlock() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QCamera6unlockEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcamera.h:232
// index:1
// Public Visibility=Default Availability=Available
// [-2] void unlock(QCamera::LockTypes)

/*
Unlock all the requested camera locks.
*/
func (this *QCamera) Unlock1(locks int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QCamera6unlockE6QFlagsINS_8LockTypeEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), locks)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcamera.h:235
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stateChanged(QCamera::State)

/*
Signals the camera state has changed.

Usually the state changes is caused by calling load(), unload(), start() and stop(), but the state can also be changed change as a result of camera error.

Note: Notifier signal for property state.
*/
func (this *QCamera) StateChanged(state int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QCamera12stateChangedENS_5StateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), state)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcamera.h:236
// index:0
// Public Visibility=Default Availability=Available
// [-2] void captureModeChanged(QCamera::CaptureModes)

/*
Signals the capture mode has changed.

Note: Notifier signal for property captureMode.
*/
func (this *QCamera) CaptureModeChanged(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QCamera18captureModeChangedE6QFlagsINS_11CaptureModeEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcamera.h:237
// index:0
// Public Visibility=Default Availability=Available
// [-2] void statusChanged(QCamera::Status)

/*
Signals the camera status has changed.

Note: Notifier signal for property status.
*/
func (this *QCamera) StatusChanged(status int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QCamera13statusChangedENS_6StatusE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), status)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcamera.h:239
// index:0
// Public Visibility=Default Availability=Available
// [-2] void locked()

/*
Signals all the requested camera settings are locked.
*/
func (this *QCamera) Locked() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QCamera6lockedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcamera.h:240
// index:0
// Public Visibility=Default Availability=Available
// [-2] void lockFailed()

/*
Signals locking of at least one requested camera settings failed.
*/
func (this *QCamera) LockFailed() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QCamera10lockFailedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcamera.h:242
// index:0
// Public Visibility=Default Availability=Available
// [-2] void lockStatusChanged(QCamera::LockStatus, QCamera::LockChangeReason)

/*
Signals the overall status for all the requested camera locks was changed with specified reason.

Note: Signal lockStatusChanged is overloaded in this class. To connect to this signal by using the function pointer syntax, Qt provides a convenient helper for obtaining the function pointer as shown in this example:


  connect(camera, QOverload<QCamera::LockStatus, QCamera::LockChangeReason>::of(&QCamera::lockStatusChanged),
      [=](QCamera::LockStatus status, QCamera::LockChangeReason reason){ /-* ... *-/ });



Note: Notifier signal for property lockStatus.
*/
func (this *QCamera) LockStatusChanged(status int, reason int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QCamera17lockStatusChangedENS_10LockStatusENS_16LockChangeReasonE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), status, reason)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcamera.h:243
// index:1
// Public Visibility=Default Availability=Available
// [-2] void lockStatusChanged(QCamera::LockType, QCamera::LockStatus, QCamera::LockChangeReason)

/*
Signals the overall status for all the requested camera locks was changed with specified reason.

Note: Signal lockStatusChanged is overloaded in this class. To connect to this signal by using the function pointer syntax, Qt provides a convenient helper for obtaining the function pointer as shown in this example:


  connect(camera, QOverload<QCamera::LockStatus, QCamera::LockChangeReason>::of(&QCamera::lockStatusChanged),
      [=](QCamera::LockStatus status, QCamera::LockChangeReason reason){ /-* ... *-/ });



Note: Notifier signal for property lockStatus.
*/
func (this *QCamera) LockStatusChanged1(lock int, status int, reason int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QCamera17lockStatusChangedENS_8LockTypeENS_10LockStatusENS_16LockChangeReasonE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), lock, status, reason)
	qtrt.ErrPrint(err, rv)
}

/*


Depending on backend, changing some camera settings like capture mode, codecs or resolution in ActiveState may lead to changing the camera status to LoadedStatus and StartingStatus while the settings are applied and back to ActiveStatus when the camera is ready.


*/
type QCamera__Status = int

// The camera or camera backend is not available.
const QCamera__UnavailableStatus QCamera__Status = 0

// The initial camera status, with camera not loaded. The camera capabilities including supported capture settings may be unknown.
const QCamera__UnloadedStatus QCamera__Status = 1

// The camera device loading in result of state transition from QCamera::UnloadedState to QCamera::LoadedState or QCamera::ActiveState.
const QCamera__LoadingStatus QCamera__Status = 2

// The camera device is unloading in result of state transition from QCamera::LoadedState or QCamera::ActiveState to QCamera::UnloadedState.
const QCamera__UnloadingStatus QCamera__Status = 3

// The camera is loaded and ready to be configured. This status indicates the camera device is opened and it's possible to query for supported image and video capture settings, like resolution, framerate and codecs.
const QCamera__LoadedStatus QCamera__Status = 4

// The camera is in the power saving standby mode. The camera may come to the standby mode after some time of inactivity in the QCamera::LoadedState state.
const QCamera__StandbyStatus QCamera__Status = 5

// The camera is starting in result of state transition to QCamera::ActiveState. The camera service is not ready to capture yet.
const QCamera__StartingStatus QCamera__Status = 6

// The camera is stopping in result of state transition from QCamera::ActiveState to QCamera::LoadedState or QCamera::UnloadedState.
const QCamera__StoppingStatus QCamera__Status = 7

// The camera has been started and can produce data. The viewfinder displays video frames in active state.
const QCamera__ActiveStatus QCamera__Status = 8

func (this *QCamera) StatusItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QCamera_StatusItemName(val int) string {
	var nilthis *QCamera
	return nilthis.StatusItemName(val)
}

/*


While the supported settings are unknown in this state, it's allowed to set the camera capture settings like codec, resolution, or frame rate.



In the Idle state it's allowed to query camera capabilities, set capture resolution, codecs, etc.

The viewfinder is not active in the loaded state.


*/
type QCamera__State = int

// The initial camera state, with camera not loaded, the camera capabilities except of supported capture modes are unknown.
const QCamera__UnloadedState QCamera__State = 0

// The camera is loaded and ready to be configured.
const QCamera__LoadedState QCamera__State = 1

// In the active state as soon as camera is started the viewfinder displays video frames and the camera is ready for capture.
const QCamera__ActiveState QCamera__State = 2

func (this *QCamera) StateItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QCamera_StateItemName(val int) string {
	var nilthis *QCamera
	return nilthis.StateItemName(val)
}

/*


 */
type QCamera__CaptureMode = int

//
const QCamera__CaptureViewfinder QCamera__CaptureMode = 0

//
const QCamera__CaptureStillImage QCamera__CaptureMode = 1

//
const QCamera__CaptureVideo QCamera__CaptureMode = 2

func (this *QCamera) CaptureModeItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QCamera_CaptureModeItemName(val int) string {
	var nilthis *QCamera
	return nilthis.CaptureModeItemName(val)
}

/*

 */
type QCamera__Error = int

// No errors have occurred.
const QCamera__NoError QCamera__Error = 0

// An error has occurred.
const QCamera__CameraError QCamera__Error = 1

// System resource doesn't support requested functionality.
const QCamera__InvalidRequestError QCamera__Error = 2

// No camera service available.
const QCamera__ServiceMissingError QCamera__Error = 3

// The feature is not supported.
const QCamera__NotSupportedFeatureError QCamera__Error = 4

func (this *QCamera) ErrorItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QCamera_ErrorItemName(val int) string {
	var nilthis *QCamera
	return nilthis.ErrorItemName(val)
}

/*


The locked state usually means the requested parameter stays the same, except of the cases when the parameter is requested to be constantly updated. For example in continuous focusing mode, the focus is considered locked as long and the object is in focus, even while the actual focusing distance may be constantly changing.

*/
type QCamera__LockStatus = int

// The application is not interested in camera settings value. The camera may keep this parameter without changes, this is common with camera focus, or adjust exposure and white balance constantly to keep the viewfinder image nice.
const QCamera__Unlocked QCamera__LockStatus = 0

// The application has requested the camera focus, exposure or white balance lock with QCamera::searchAndLock(). This state indicates the camera is focusing or calculating exposure and white balance.
const QCamera__Searching QCamera__LockStatus = 1

// The camera focus, exposure or white balance is locked. The camera is ready to capture, application may check the exposure parameters.
const QCamera__Locked QCamera__LockStatus = 2

func (this *QCamera) LockStatusItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QCamera_LockStatusItemName(val int) string {
	var nilthis *QCamera
	return nilthis.LockStatusItemName(val)
}

/*

 */
type QCamera__LockChangeReason = int

// The lock status changed in result of user request, usually to unlock camera settings.
const QCamera__UserRequest QCamera__LockChangeReason = 0

// The lock status successfuly changed to QCamera::Locked.
const QCamera__LockAcquired QCamera__LockChangeReason = 1

// The camera failed to acquire the requested lock in result of autofocus failure, exposure out of supported range, etc.
const QCamera__LockFailed QCamera__LockChangeReason = 2

// The camera is not able to maintain the requested lock any more. Lock status is changed to QCamera::Unlocked.
const QCamera__LockLost QCamera__LockChangeReason = 3

// The lock is lost, but the camera is working hard to reacquire it. This value may be used in continuous focusing mode, when the camera loses the focus, the focus lock state is changed to Qcamera::Searching with LockTemporaryLost reason.
const QCamera__LockTemporaryLost QCamera__LockChangeReason = 4

func (this *QCamera) LockChangeReasonItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QCamera_LockChangeReasonItemName(val int) string {
	var nilthis *QCamera
	return nilthis.LockChangeReasonItemName(val)
}

/*


 */
type QCamera__LockType = int

//
const QCamera__NoLock QCamera__LockType = 0

//
const QCamera__LockExposure QCamera__LockType = 1

//
const QCamera__LockWhiteBalance QCamera__LockType = 2

//
const QCamera__LockFocus QCamera__LockType = 4

func (this *QCamera) LockTypeItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QCamera_LockTypeItemName(val int) string {
	var nilthis *QCamera
	return nilthis.LockTypeItemName(val)
}

/*
This enum specifies the physical position of the camera on the system hardware.



This enum was introduced or modified in  Qt 5.3.

See also QCameraInfo::position().

*/
type QCamera__Position = int

// The camera position is unspecified or unknown.
const QCamera__UnspecifiedPosition QCamera__Position = 0

// The camera is on the back face of the system hardware. For example on a mobile device, it means it is on the opposite side to that of the screen.
const QCamera__BackFace QCamera__Position = 1

// The camera is on the front face of the system hardware. For example on a mobile device, it means it is on the same side as that of the screen. Viewfinder frames of front-facing cameras are mirrored horizontally, so the users can see themselves as looking into a mirror. Captured images or videos are not mirrored.
const QCamera__FrontFace QCamera__Position = 2

func (this *QCamera) PositionItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QCamera_PositionItemName(val int) string {
	var nilthis *QCamera
	return nilthis.PositionItemName(val)
}

//  body block end

//  keep block begin

func init_unused_11809() {
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
