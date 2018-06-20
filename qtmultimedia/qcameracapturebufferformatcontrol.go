package qtmultimedia

// /usr/include/qt/QtMultimedia/qcameracapturebufferformatcontrol.h
// #include <qcameracapturebufferformatcontrol.h>
// #include <QtMultimedia>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 31
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
type QCameraCaptureBufferFormatControl struct {
	*QMediaControl
}
type QCameraCaptureBufferFormatControl_ITF interface {
	QMediaControl_ITF
	QCameraCaptureBufferFormatControl_PTR() *QCameraCaptureBufferFormatControl
}

func (ptr *QCameraCaptureBufferFormatControl) QCameraCaptureBufferFormatControl_PTR() *QCameraCaptureBufferFormatControl {
	return ptr
}

func (this *QCameraCaptureBufferFormatControl) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QMediaControl.GetCthis()
	}
}
func (this *QCameraCaptureBufferFormatControl) SetCthis(cthis unsafe.Pointer) {
	this.QMediaControl = NewQMediaControlFromPointer(cthis)
}
func NewQCameraCaptureBufferFormatControlFromPointer(cthis unsafe.Pointer) *QCameraCaptureBufferFormatControl {
	bcthis0 := NewQMediaControlFromPointer(cthis)
	return &QCameraCaptureBufferFormatControl{bcthis0}
}
func (*QCameraCaptureBufferFormatControl) NewFromPointer(cthis unsafe.Pointer) *QCameraCaptureBufferFormatControl {
	return NewQCameraCaptureBufferFormatControlFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qcameracapturebufferformatcontrol.h:53
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QCameraCaptureBufferFormatControl) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK33QCameraCaptureBufferFormatControl10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qcameracapturebufferformatcontrol.h:55
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QCameraCaptureBufferFormatControl()

/*

 */
func DeleteQCameraCaptureBufferFormatControl(this *QCameraCaptureBufferFormatControl) {
	rv, err := qtrt.InvokeQtFunc6("_ZN33QCameraCaptureBufferFormatControlD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qcameracapturebufferformatcontrol.h:58
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] QVideoFrame::PixelFormat bufferFormat() const

/*
Returns the current buffer capture format.

See also setBufferFormat().
*/
func (this *QCameraCaptureBufferFormatControl) BufferFormat() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK33QCameraCaptureBufferFormatControl12bufferFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qcameracapturebufferformatcontrol.h:59
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setBufferFormat(QVideoFrame::PixelFormat)

/*
Sets the buffer capture format.

See also bufferFormat().
*/
func (this *QCameraCaptureBufferFormatControl) SetBufferFormat(format int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN33QCameraCaptureBufferFormatControl15setBufferFormatEN11QVideoFrame11PixelFormatE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), format)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameracapturebufferformatcontrol.h:62
// index:0
// Public Visibility=Default Availability=Available
// [-2] void bufferFormatChanged(QVideoFrame::PixelFormat)

/*
Signals the buffer image capture format changed to format.
*/
func (this *QCameraCaptureBufferFormatControl) BufferFormatChanged(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN33QCameraCaptureBufferFormatControl19bufferFormatChangedEN11QVideoFrame11PixelFormatE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameracapturebufferformatcontrol.h:65
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QCameraCaptureBufferFormatControl(QObject *)

/*
Constructs a new image buffer capture format control object with the given parent
*/
func NewQCameraCaptureBufferFormatControl(parent qtcore.QObject_ITF /*777 QObject **/) *QCameraCaptureBufferFormatControl {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN33QCameraCaptureBufferFormatControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCameraCaptureBufferFormatControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QCameraCaptureBufferFormatControl")
	return gothis
}

// /usr/include/qt/QtMultimedia/qcameracapturebufferformatcontrol.h:65
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QCameraCaptureBufferFormatControl(QObject *)

/*
Constructs a new image buffer capture format control object with the given parent
*/
func NewQCameraCaptureBufferFormatControl__() *QCameraCaptureBufferFormatControl {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN33QCameraCaptureBufferFormatControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCameraCaptureBufferFormatControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QCameraCaptureBufferFormatControl")
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
