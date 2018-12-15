package qtmultimedia

// /usr/include/qt/QtMultimedia/qvideoprobe.h
// #include <qvideoprobe.h>
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
type QVideoProbe struct {
	*qtcore.QObject
}
type QVideoProbe_ITF interface {
	qtcore.QObject_ITF
	QVideoProbe_PTR() *QVideoProbe
}

func (ptr *QVideoProbe) QVideoProbe_PTR() *QVideoProbe { return ptr }

func (this *QVideoProbe) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QVideoProbe) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQVideoProbeFromPointer(cthis unsafe.Pointer) *QVideoProbe {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QVideoProbe{bcthis0}
}
func (*QVideoProbe) NewFromPointer(cthis unsafe.Pointer) *QVideoProbe {
	return NewQVideoProbeFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qvideoprobe.h:54
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QVideoProbe) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QVideoProbe10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qvideoprobe.h:56
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QVideoProbe(QObject *)

/*
Creates a new QVideoProbe class with parent. After setting the source to monitor with setSource(), the videoFrameProbed() signal will be emitted when video frames are flowing in the source media object.
*/
func (*QVideoProbe) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QVideoProbe {
	return NewQVideoProbe(parent)
}
func NewQVideoProbe(parent qtcore.QObject_ITF /*777 QObject **/) *QVideoProbe {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QVideoProbeC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVideoProbeFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QVideoProbe")
	return gothis
}

// /usr/include/qt/QtMultimedia/qvideoprobe.h:56
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QVideoProbe(QObject *)

/*
Creates a new QVideoProbe class with parent. After setting the source to monitor with setSource(), the videoFrameProbed() signal will be emitted when video frames are flowing in the source media object.
*/
func (*QVideoProbe) NewForInheritp() *QVideoProbe {
	return NewQVideoProbep()
}
func NewQVideoProbep() *QVideoProbe {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QVideoProbeC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVideoProbeFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QVideoProbe")
	return gothis
}

// /usr/include/qt/QtMultimedia/qvideoprobe.h:57
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QVideoProbe()

/*

 */
func DeleteQVideoProbe(this *QVideoProbe) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QVideoProbeD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qvideoprobe.h:59
// index:0
// Public Visibility=Default Availability=Available
// [1] bool setSource(QMediaObject *)

/*
Sets the media object to monitor to source.

If source is zero, this probe will be deactivated and this function wil return true.

If the media object does not support monitoring video, this function will return false.

Any previously monitored objects will no longer be monitored. Passing in the same object will be ignored, but monitoring will continue.
*/
func (this *QVideoProbe) SetSource(source QMediaObject_ITF /*777 QMediaObject **/) bool {
	var convArg0 unsafe.Pointer
	if source != nil && source.QMediaObject_PTR() != nil {
		convArg0 = source.QMediaObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QVideoProbe9setSourceEP12QMediaObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qvideoprobe.h:60
// index:1
// Public Visibility=Default Availability=Available
// [1] bool setSource(QMediaRecorder *)

/*
Sets the media object to monitor to source.

If source is zero, this probe will be deactivated and this function wil return true.

If the media object does not support monitoring video, this function will return false.

Any previously monitored objects will no longer be monitored. Passing in the same object will be ignored, but monitoring will continue.
*/
func (this *QVideoProbe) SetSource1(source QMediaRecorder_ITF /*777 QMediaRecorder **/) bool {
	var convArg0 unsafe.Pointer
	if source != nil && source.QMediaRecorder_PTR() != nil {
		convArg0 = source.QMediaRecorder_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QVideoProbe9setSourceEP14QMediaRecorder", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qvideoprobe.h:62
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isActive() const

/*
Returns true if this probe is monitoring something, or false otherwise.

The source being monitored does not need to be active.
*/
func (this *QVideoProbe) IsActive() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QVideoProbe8isActiveEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qvideoprobe.h:65
// index:0
// Public Visibility=Default Availability=Available
// [-2] void videoFrameProbed(const QVideoFrame &)

/*
This signal should be emitted when a video frame is processed in the media service.
*/
func (this *QVideoProbe) VideoFrameProbed(frame QVideoFrame_ITF) {
	var convArg0 unsafe.Pointer
	if frame != nil && frame.QVideoFrame_PTR() != nil {
		convArg0 = frame.QVideoFrame_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QVideoProbe16videoFrameProbedERK11QVideoFrame", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qvideoprobe.h:66
// index:0
// Public Visibility=Default Availability=Available
// [-2] void flush()

/*
This signal should be emitted when it is required to release all frames. Application must release all outstanding references to video frames.
*/
func (this *QVideoProbe) Flush() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QVideoProbe5flushEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
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
