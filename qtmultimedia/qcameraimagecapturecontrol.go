package qtmultimedia

// /usr/include/qt/QtMultimedia/qcameraimagecapturecontrol.h
// #include <qcameraimagecapturecontrol.h>
// #include <QtMultimedia>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 16
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
type QCameraImageCaptureControl struct {
	*QMediaControl
}
type QCameraImageCaptureControl_ITF interface {
	QMediaControl_ITF
	QCameraImageCaptureControl_PTR() *QCameraImageCaptureControl
}

func (ptr *QCameraImageCaptureControl) QCameraImageCaptureControl_PTR() *QCameraImageCaptureControl {
	return ptr
}

func (this *QCameraImageCaptureControl) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QMediaControl.GetCthis()
	}
}
func (this *QCameraImageCaptureControl) SetCthis(cthis unsafe.Pointer) {
	this.QMediaControl = NewQMediaControlFromPointer(cthis)
}
func NewQCameraImageCaptureControlFromPointer(cthis unsafe.Pointer) *QCameraImageCaptureControl {
	bcthis0 := NewQMediaControlFromPointer(cthis)
	return &QCameraImageCaptureControl{bcthis0}
}
func (*QCameraImageCaptureControl) NewFromPointer(cthis unsafe.Pointer) *QCameraImageCaptureControl {
	return NewQCameraImageCaptureControlFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qcameraimagecapturecontrol.h:58
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QCameraImageCaptureControl) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK26QCameraImageCaptureControl10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qcameraimagecapturecontrol.h:61
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QCameraImageCaptureControl()

/*

 */
func DeleteQCameraImageCaptureControl(this *QCameraImageCaptureControl) {
	rv, err := qtrt.InvokeQtFunc6("_ZN26QCameraImageCaptureControlD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qcameraimagecapturecontrol.h:63
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool isReadyForCapture() const

/*
Identifies if a capture control is ready to perform a capture immediately (all the resources necessary for image capture are allocated, hardware initialized, flash is charged, etc).

Returns true if the camera is ready for capture; and false if it is not.

It's permissible to call capture() while the camera status is QCamera::ActiveStatus regardless of isReadyForCapture property value. If camera is not ready to capture image immediately, the capture request is queued with all the related camera settings to be executed as soon as possible.
*/
func (this *QCameraImageCaptureControl) IsReadyForCapture() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK26QCameraImageCaptureControl17isReadyForCaptureEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qcameraimagecapturecontrol.h:65
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] QCameraImageCapture::DriveMode driveMode() const

/*
Returns the current camera drive mode.

See also setDriveMode().
*/
func (this *QCameraImageCaptureControl) DriveMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK26QCameraImageCaptureControl9driveModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qcameraimagecapturecontrol.h:66
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setDriveMode(QCameraImageCapture::DriveMode)

/*
Sets the current camera drive mode.

See also driveMode().
*/
func (this *QCameraImageCaptureControl) SetDriveMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN26QCameraImageCaptureControl12setDriveModeEN19QCameraImageCapture9DriveModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameraimagecapturecontrol.h:68
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] int capture(const QString &)

/*
Initiates the capture of an image to fileName. The fileName can be relative or empty, in this case the service should use the system specific place and file naming scheme.

The Camera service should save all the capture parameters like exposure settings or image processing parameters, so changes to camera parameters after capture() is called do not affect previous capture requests.

Returns the capture request id number, which is used later with imageExposed(), imageCaptured() and imageSaved() signals.
*/
func (this *QCameraImageCaptureControl) Capture(fileName string) int {
	var tmpArg0 = qtcore.NewQString5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN26QCameraImageCaptureControl7captureERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qcameraimagecapturecontrol.h:69
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void cancelCapture()

/*
Cancel pending capture requests.
*/
func (this *QCameraImageCaptureControl) CancelCapture() {
	rv, err := qtrt.InvokeQtFunc6("_ZN26QCameraImageCaptureControl13cancelCaptureEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameraimagecapturecontrol.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void readyForCaptureChanged(bool)

/*
Signals that a capture control's ready state has changed.
*/
func (this *QCameraImageCaptureControl) ReadyForCaptureChanged(ready bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN26QCameraImageCaptureControl22readyForCaptureChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ready)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameraimagecapturecontrol.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void imageExposed(int)

/*
Signals that an image with it requestId has just been exposed. This signal can be used for the shutter sound or other indicaton.
*/
func (this *QCameraImageCaptureControl) ImageExposed(requestId int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN26QCameraImageCaptureControl12imageExposedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), requestId)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameraimagecapturecontrol.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void imageCaptured(int, const QImage &)

/*
Signals that an image with it requestId has been captured and a preview is available.
*/
func (this *QCameraImageCaptureControl) ImageCaptured(requestId int, preview qtgui.QImage_ITF) {
	var convArg1 unsafe.Pointer
	if preview != nil && preview.QImage_PTR() != nil {
		convArg1 = preview.QImage_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN26QCameraImageCaptureControl13imageCapturedEiRK6QImage", qtrt.FFI_TYPE_POINTER, this.GetCthis(), requestId, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameraimagecapturecontrol.h:76
// index:0
// Public Visibility=Default Availability=Available
// [-2] void imageMetadataAvailable(int, const QString &, const QVariant &)

/*
Signals that a metadata for an image with request id is available. Signal also contains the key and value of the metadata.

This signal should be emitted between imageExposed and imageSaved signals.
*/
func (this *QCameraImageCaptureControl) ImageMetadataAvailable(id int, key string, value qtcore.QVariant_ITF) {
	var tmpArg1 = qtcore.NewQString5(key)
	var convArg1 = tmpArg1.GetCthis()
	var convArg2 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg2 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN26QCameraImageCaptureControl22imageMetadataAvailableEiRK7QStringRK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), id, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameraimagecapturecontrol.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void imageAvailable(int, const QVideoFrame &)

/*
Signals that a captured buffer with a requestId is available.
*/
func (this *QCameraImageCaptureControl) ImageAvailable(requestId int, buffer QVideoFrame_ITF) {
	var convArg1 unsafe.Pointer
	if buffer != nil && buffer.QVideoFrame_PTR() != nil {
		convArg1 = buffer.QVideoFrame_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN26QCameraImageCaptureControl14imageAvailableEiRK11QVideoFrame", qtrt.FFI_TYPE_POINTER, this.GetCthis(), requestId, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameraimagecapturecontrol.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void imageSaved(int, const QString &)

/*
Signals that a captured image with a requestId has been saved to fileName.
*/
func (this *QCameraImageCaptureControl) ImageSaved(requestId int, fileName string) {
	var tmpArg1 = qtcore.NewQString5(fileName)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN26QCameraImageCaptureControl10imageSavedEiRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), requestId, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameraimagecapturecontrol.h:80
// index:0
// Public Visibility=Default Availability=Available
// [-2] void error(int, int, const QString &)

/*
Signals the capture request id failed with error code and message errorString.

See also QCameraImageCapture::Error.
*/
func (this *QCameraImageCaptureControl) Error(id int, error int, errorString string) {
	var tmpArg2 = qtcore.NewQString5(errorString)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN26QCameraImageCaptureControl5errorEiiRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), id, error, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameraimagecapturecontrol.h:83
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QCameraImageCaptureControl(QObject *)

/*
Constructs a new image capture control object with the given parent
*/
func (*QCameraImageCaptureControl) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QCameraImageCaptureControl {
	return NewQCameraImageCaptureControl(parent)
}
func NewQCameraImageCaptureControl(parent qtcore.QObject_ITF /*777 QObject **/) *QCameraImageCaptureControl {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN26QCameraImageCaptureControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCameraImageCaptureControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QCameraImageCaptureControl")
	return gothis
}

// /usr/include/qt/QtMultimedia/qcameraimagecapturecontrol.h:83
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QCameraImageCaptureControl(QObject *)

/*
Constructs a new image capture control object with the given parent
*/
func (*QCameraImageCaptureControl) NewForInheritp() *QCameraImageCaptureControl {
	return NewQCameraImageCaptureControlp()
}
func NewQCameraImageCaptureControlp() *QCameraImageCaptureControl {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN26QCameraImageCaptureControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCameraImageCaptureControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QCameraImageCaptureControl")
	return gothis
}

//  body block end

//  keep block begin

func init_unused_11827() {
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
