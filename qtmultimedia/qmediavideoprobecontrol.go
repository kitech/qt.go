package qtmultimedia

// /usr/include/qt/QtMultimedia/qmediavideoprobecontrol.h
// #include <qmediavideoprobecontrol.h>
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
type QMediaVideoProbeControl struct {
	*QMediaControl
}
type QMediaVideoProbeControl_ITF interface {
	QMediaControl_ITF
	QMediaVideoProbeControl_PTR() *QMediaVideoProbeControl
}

func (ptr *QMediaVideoProbeControl) QMediaVideoProbeControl_PTR() *QMediaVideoProbeControl { return ptr }

func (this *QMediaVideoProbeControl) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QMediaControl.GetCthis()
	}
}
func (this *QMediaVideoProbeControl) SetCthis(cthis unsafe.Pointer) {
	this.QMediaControl = NewQMediaControlFromPointer(cthis)
}
func NewQMediaVideoProbeControlFromPointer(cthis unsafe.Pointer) *QMediaVideoProbeControl {
	bcthis0 := NewQMediaControlFromPointer(cthis)
	return &QMediaVideoProbeControl{bcthis0}
}
func (*QMediaVideoProbeControl) NewFromPointer(cthis unsafe.Pointer) *QMediaVideoProbeControl {
	return NewQMediaVideoProbeControlFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qmediavideoprobecontrol.h:52
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QMediaVideoProbeControl) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QMediaVideoProbeControl10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qmediavideoprobecontrol.h:54
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QMediaVideoProbeControl()

/*

 */
func DeleteQMediaVideoProbeControl(this *QMediaVideoProbeControl) {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QMediaVideoProbeControlD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qmediavideoprobecontrol.h:57
// index:0
// Public Visibility=Default Availability=Available
// [-2] void videoFrameProbed(const QVideoFrame &)

/*
This signal should be emitted when a video frame is processed in the media service.
*/
func (this *QMediaVideoProbeControl) VideoFrameProbed(frame QVideoFrame_ITF) {
	var convArg0 unsafe.Pointer
	if frame != nil && frame.QVideoFrame_PTR() != nil {
		convArg0 = frame.QVideoFrame_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN23QMediaVideoProbeControl16videoFrameProbedERK11QVideoFrame", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediavideoprobecontrol.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void flush()

/*
This signal should be emitted when it is required to release all frames.
*/
func (this *QMediaVideoProbeControl) Flush() {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QMediaVideoProbeControl5flushEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediavideoprobecontrol.h:61
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QMediaVideoProbeControl(QObject *)

/*
Create a new media video probe control object with the given parent.
*/
func NewQMediaVideoProbeControl(parent qtcore.QObject_ITF /*777 QObject **/) *QMediaVideoProbeControl {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN23QMediaVideoProbeControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMediaVideoProbeControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QMediaVideoProbeControl")
	return gothis
}

// /usr/include/qt/QtMultimedia/qmediavideoprobecontrol.h:61
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QMediaVideoProbeControl(QObject *)

/*
Create a new media video probe control object with the given parent.
*/
func NewQMediaVideoProbeControl__() *QMediaVideoProbeControl {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN23QMediaVideoProbeControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMediaVideoProbeControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QMediaVideoProbeControl")
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
