package qtmultimedia

// /usr/include/qt/QtMultimedia/qcamerainfo.h
// #include <qcamerainfo.h>
// #include <QtMultimedia>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
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
type QCameraInfo struct {
	*qtrt.CObject
}
type QCameraInfo_ITF interface {
	QCameraInfo_PTR() *QCameraInfo
}

func (ptr *QCameraInfo) QCameraInfo_PTR() *QCameraInfo { return ptr }

func (this *QCameraInfo) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QCameraInfo) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQCameraInfoFromPointer(cthis unsafe.Pointer) *QCameraInfo {
	return &QCameraInfo{&qtrt.CObject{cthis}}
}
func (*QCameraInfo) NewFromPointer(cthis unsafe.Pointer) *QCameraInfo {
	return NewQCameraInfoFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qcamerainfo.h:53
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QCameraInfo(const QByteArray &)

/*
Constructs a camera info object from a camera device name.

If no such device exists, the QCameraInfo object will be invalid and isNull() will return true.
*/
func (*QCameraInfo) NewForInherit(name qtcore.QByteArray_ITF) *QCameraInfo {
	return NewQCameraInfo(name)
}
func NewQCameraInfo(name qtcore.QByteArray_ITF) *QCameraInfo {
	var convArg0 unsafe.Pointer
	if name != nil && name.QByteArray_PTR() != nil {
		convArg0 = name.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QCameraInfoC2ERK10QByteArray", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCameraInfoFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQCameraInfo)
	return gothis
}

// /usr/include/qt/QtMultimedia/qcamerainfo.h:53
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QCameraInfo(const QByteArray &)

/*
Constructs a camera info object from a camera device name.

If no such device exists, the QCameraInfo object will be invalid and isNull() will return true.
*/
func (*QCameraInfo) NewForInherit__() *QCameraInfo {
	return NewQCameraInfo__()
}
func NewQCameraInfo__() *QCameraInfo {
	// arg: 0, const QByteArray &=LValueReference, QByteArray=Record, , Invalid
	var convArg0 = qtcore.NewQByteArray()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QCameraInfoC2ERK10QByteArray", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCameraInfoFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQCameraInfo)
	return gothis
}

// /usr/include/qt/QtMultimedia/qcamerainfo.h:54
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QCameraInfo(const QCamera &)

/*
Constructs a camera info object from a camera device name.

If no such device exists, the QCameraInfo object will be invalid and isNull() will return true.
*/
func (*QCameraInfo) NewForInherit_1(camera QCamera_ITF) *QCameraInfo {
	return NewQCameraInfo_1(camera)
}
func NewQCameraInfo_1(camera QCamera_ITF) *QCameraInfo {
	var convArg0 unsafe.Pointer
	if camera != nil && camera.QCamera_PTR() != nil {
		convArg0 = camera.QCamera_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QCameraInfoC2ERK7QCamera", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCameraInfoFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQCameraInfo)
	return gothis
}

// /usr/include/qt/QtMultimedia/qcamerainfo.h:56
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QCameraInfo()

/*

 */
func DeleteQCameraInfo(this *QCameraInfo) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QCameraInfoD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qcamerainfo.h:58
// index:0
// Public Visibility=Default Availability=Available
// [16] QCameraInfo & operator=(const QCameraInfo &)

/*

 */
func (this *QCameraInfo) Operator_equal(other QCameraInfo_ITF) *QCameraInfo {
	var convArg0 unsafe.Pointer
	if other != nil && other.QCameraInfo_PTR() != nil {
		convArg0 = other.QCameraInfo_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QCameraInfoaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCameraInfoFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCameraInfo)
	return rv2
}

// /usr/include/qt/QtMultimedia/qcamerainfo.h:59
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QCameraInfo &) const

/*

 */
func (this *QCameraInfo) Operator_equal_equal(other QCameraInfo_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QCameraInfo_PTR() != nil {
		convArg0 = other.QCameraInfo_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QCameraInfoeqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qcamerainfo.h:60
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QCameraInfo &) const

/*

 */
func (this *QCameraInfo) Operator_not_equal(other QCameraInfo_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QCameraInfo_PTR() != nil {
		convArg0 = other.QCameraInfo_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QCameraInfoneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qcamerainfo.h:62
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isNull() const

/*
Returns true if this QCameraInfo is null or invalid.
*/
func (this *QCameraInfo) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QCameraInfo6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qcamerainfo.h:64
// index:0
// Public Visibility=Default Availability=Available
// [8] QString deviceName() const

/*
Returns the device name of the camera

This is a unique ID to identify the camera and may not be human-readable.
*/
func (this *QCameraInfo) DeviceName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QCameraInfo10deviceNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtMultimedia/qcamerainfo.h:65
// index:0
// Public Visibility=Default Availability=Available
// [8] QString description() const

/*
Returns the human-readable description of the camera.
*/
func (this *QCameraInfo) Description() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QCameraInfo11descriptionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtMultimedia/qcamerainfo.h:66
// index:0
// Public Visibility=Default Availability=Available
// [4] QCamera::Position position() const

/*
Returns the physical position of the camera on the hardware system.
*/
func (this *QCameraInfo) Position() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QCameraInfo8positionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qcamerainfo.h:67
// index:0
// Public Visibility=Default Availability=Available
// [4] int orientation() const

/*
Returns the physical orientation of the camera sensor.

The value is the orientation angle (clockwise, in steps of 90 degrees) of the camera sensor in relation to the display in its natural orientation.

You can show the camera image in the correct orientation by rotating it by this value in the anti-clockwise direction.

For example, suppose a mobile device which is naturally in portrait orientation. The back-facing camera is mounted in landscape. If the top side of the camera sensor is aligned with the right edge of the screen in natural orientation, the value should be 270. If the top side of a front-facing camera sensor is aligned with the right of the screen, the value should be 90.
*/
func (this *QCameraInfo) Orientation() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QCameraInfo11orientationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qcamerainfo.h:69
// index:0
// Public static Visibility=Default Availability=Available
// [16] QCameraInfo defaultCamera()

/*
Returns the default camera on the system.

The returned object should be checked using isNull() before being used, in case there is no default camera or no cameras at all.

See also availableCameras().
*/
func (this *QCameraInfo) DefaultCamera() *QCameraInfo /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QCameraInfo13defaultCameraEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCameraInfoFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCameraInfo)
	return rv2
}
func QCameraInfo_DefaultCamera() *QCameraInfo /*123*/ {
	var nilthis *QCameraInfo
	rv := nilthis.DefaultCamera()
	return rv
}

// /usr/include/qt/QtMultimedia/qcamerainfo.h:70
// index:0
// Public static Visibility=Default Availability=Available
// [-2] QList<QCameraInfo> availableCameras(QCamera::Position)

/*
Returns a list of available cameras on the system which are located at position.

If position is not specified or if the value is QCamera::UnspecifiedPosition, a list of all available cameras will be returned.
*/
func (this *QCameraInfo) AvailableCameras(position int) *QCameraInfoList /*lll*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QCameraInfo16availableCamerasEN7QCamera8PositionE", qtrt.FFI_TYPE_POINTER, position)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCameraInfoListFromPointer(unsafe.Pointer(uintptr(rv))) //5552
	return rv2
}
func QCameraInfo_AvailableCameras(position int) *QCameraInfoList /*lll*/ {
	var nilthis *QCameraInfo
	rv := nilthis.AvailableCameras(position)
	return rv
}

// /usr/include/qt/QtMultimedia/qcamerainfo.h:70
// index:0
// Public static Visibility=Default Availability=Available
// [-2] QList<QCameraInfo> availableCameras(QCamera::Position)

/*
Returns a list of available cameras on the system which are located at position.

If position is not specified or if the value is QCamera::UnspecifiedPosition, a list of all available cameras will be returned.
*/
func (this *QCameraInfo) AvailableCameras__() *QCameraInfoList /*lll*/ {
	// arg: 0, QCamera::Position=Elaborated, QCamera::Position=Enum, , Invalid
	position := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN11QCameraInfo16availableCamerasEN7QCamera8PositionE", qtrt.FFI_TYPE_POINTER, position)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCameraInfoListFromPointer(unsafe.Pointer(uintptr(rv))) //5552
	return rv2
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
