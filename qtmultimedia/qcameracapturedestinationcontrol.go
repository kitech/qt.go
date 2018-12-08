package qtmultimedia

// /usr/include/qt/QtMultimedia/qcameracapturedestinationcontrol.h
// #include <qcameracapturedestinationcontrol.h>
// #include <QtMultimedia>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 6
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
type QCameraCaptureDestinationControl struct {
	*QMediaControl
}
type QCameraCaptureDestinationControl_ITF interface {
	QMediaControl_ITF
	QCameraCaptureDestinationControl_PTR() *QCameraCaptureDestinationControl
}

func (ptr *QCameraCaptureDestinationControl) QCameraCaptureDestinationControl_PTR() *QCameraCaptureDestinationControl {
	return ptr
}

func (this *QCameraCaptureDestinationControl) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QMediaControl.GetCthis()
	}
}
func (this *QCameraCaptureDestinationControl) SetCthis(cthis unsafe.Pointer) {
	this.QMediaControl = NewQMediaControlFromPointer(cthis)
}
func NewQCameraCaptureDestinationControlFromPointer(cthis unsafe.Pointer) *QCameraCaptureDestinationControl {
	bcthis0 := NewQMediaControlFromPointer(cthis)
	return &QCameraCaptureDestinationControl{bcthis0}
}
func (*QCameraCaptureDestinationControl) NewFromPointer(cthis unsafe.Pointer) *QCameraCaptureDestinationControl {
	return NewQCameraCaptureDestinationControlFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qcameracapturedestinationcontrol.h:53
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QCameraCaptureDestinationControl) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK32QCameraCaptureDestinationControl10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qcameracapturedestinationcontrol.h:55
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QCameraCaptureDestinationControl()

/*

 */
func DeleteQCameraCaptureDestinationControl(this *QCameraCaptureDestinationControl) {
	rv, err := qtrt.InvokeQtFunc6("_ZN32QCameraCaptureDestinationControlD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qcameracapturedestinationcontrol.h:57
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool isCaptureDestinationSupported(QCameraImageCapture::CaptureDestinations) const

/*
Returns true if the capture destination is supported; and false if it is not.
*/
func (this *QCameraCaptureDestinationControl) IsCaptureDestinationSupported(destination int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK32QCameraCaptureDestinationControl29isCaptureDestinationSupportedE6QFlagsIN19QCameraImageCapture18CaptureDestinationEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), destination)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qcameracapturedestinationcontrol.h:58
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] QCameraImageCapture::CaptureDestinations captureDestination() const

/*
Returns the current capture destination. The default destination is QCameraImageCapture::CaptureToFile.

See also setCaptureDestination().
*/
func (this *QCameraCaptureDestinationControl) CaptureDestination() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK32QCameraCaptureDestinationControl18captureDestinationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qcameracapturedestinationcontrol.h:59
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setCaptureDestination(QCameraImageCapture::CaptureDestinations)

/*
Sets the capture destination.

See also captureDestination().
*/
func (this *QCameraCaptureDestinationControl) SetCaptureDestination(destination int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN32QCameraCaptureDestinationControl21setCaptureDestinationE6QFlagsIN19QCameraImageCapture18CaptureDestinationEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), destination)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameracapturedestinationcontrol.h:62
// index:0
// Public Visibility=Default Availability=Available
// [-2] void captureDestinationChanged(QCameraImageCapture::CaptureDestinations)

/*
Signals the image capture destination changed.
*/
func (this *QCameraCaptureDestinationControl) CaptureDestinationChanged(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN32QCameraCaptureDestinationControl25captureDestinationChangedE6QFlagsIN19QCameraImageCapture18CaptureDestinationEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameracapturedestinationcontrol.h:65
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QCameraCaptureDestinationControl(QObject *)

/*
Constructs a new image capture destination control object with the given parent
*/
func (*QCameraCaptureDestinationControl) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QCameraCaptureDestinationControl {
	return NewQCameraCaptureDestinationControl(parent)
}
func NewQCameraCaptureDestinationControl(parent qtcore.QObject_ITF /*777 QObject **/) *QCameraCaptureDestinationControl {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN32QCameraCaptureDestinationControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCameraCaptureDestinationControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QCameraCaptureDestinationControl")
	return gothis
}

// /usr/include/qt/QtMultimedia/qcameracapturedestinationcontrol.h:65
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QCameraCaptureDestinationControl(QObject *)

/*
Constructs a new image capture destination control object with the given parent
*/
func (*QCameraCaptureDestinationControl) NewForInheritp() *QCameraCaptureDestinationControl {
	return NewQCameraCaptureDestinationControlp()
}
func NewQCameraCaptureDestinationControlp() *QCameraCaptureDestinationControl {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN32QCameraCaptureDestinationControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCameraCaptureDestinationControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QCameraCaptureDestinationControl")
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
