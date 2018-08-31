package qtmultimedia

// /usr/include/qt/QtMultimedia/qaudioprobe.h
// #include <qaudioprobe.h>
// #include <QtMultimedia>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 9
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
type QAudioProbe struct {
	*qtcore.QObject
}
type QAudioProbe_ITF interface {
	qtcore.QObject_ITF
	QAudioProbe_PTR() *QAudioProbe
}

func (ptr *QAudioProbe) QAudioProbe_PTR() *QAudioProbe { return ptr }

func (this *QAudioProbe) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QAudioProbe) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQAudioProbeFromPointer(cthis unsafe.Pointer) *QAudioProbe {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QAudioProbe{bcthis0}
}
func (*QAudioProbe) NewFromPointer(cthis unsafe.Pointer) *QAudioProbe {
	return NewQAudioProbeFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qaudioprobe.h:54
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QAudioProbe) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QAudioProbe10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qaudioprobe.h:56
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAudioProbe(QObject *)

/*
Creates a new QAudioProbe class with a parent. After setting the source to monitor with setSource(), the audioBufferProbed() signal will be emitted when audio buffers are flowing in the source media object.
*/
func (*QAudioProbe) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QAudioProbe {
	return NewQAudioProbe(parent)
}
func NewQAudioProbe(parent qtcore.QObject_ITF /*777 QObject **/) *QAudioProbe {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QAudioProbeC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAudioProbeFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAudioProbe")
	return gothis
}

// /usr/include/qt/QtMultimedia/qaudioprobe.h:56
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAudioProbe(QObject *)

/*
Creates a new QAudioProbe class with a parent. After setting the source to monitor with setSource(), the audioBufferProbed() signal will be emitted when audio buffers are flowing in the source media object.
*/
func (*QAudioProbe) NewForInherit__() *QAudioProbe {
	return NewQAudioProbe__()
}
func NewQAudioProbe__() *QAudioProbe {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QAudioProbeC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAudioProbeFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAudioProbe")
	return gothis
}

// /usr/include/qt/QtMultimedia/qaudioprobe.h:57
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAudioProbe()

/*

 */
func DeleteQAudioProbe(this *QAudioProbe) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QAudioProbeD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qaudioprobe.h:59
// index:0
// Public Visibility=Default Availability=Available
// [1] bool setSource(QMediaObject *)

/*
Sets the media object to monitor to source.

If source is zero, this probe will be deactivated and this function wil return true.

If the media object does not support monitoring audio, this function will return false.

The previous object will no longer be monitored. Passing in the same object will be ignored, but monitoring will continue.
*/
func (this *QAudioProbe) SetSource(source QMediaObject_ITF /*777 QMediaObject **/) bool {
	var convArg0 unsafe.Pointer
	if source != nil && source.QMediaObject_PTR() != nil {
		convArg0 = source.QMediaObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QAudioProbe9setSourceEP12QMediaObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qaudioprobe.h:60
// index:1
// Public Visibility=Default Availability=Available
// [1] bool setSource(QMediaRecorder *)

/*
Sets the media object to monitor to source.

If source is zero, this probe will be deactivated and this function wil return true.

If the media object does not support monitoring audio, this function will return false.

The previous object will no longer be monitored. Passing in the same object will be ignored, but monitoring will continue.
*/
func (this *QAudioProbe) SetSource_1(source QMediaRecorder_ITF /*777 QMediaRecorder **/) bool {
	var convArg0 unsafe.Pointer
	if source != nil && source.QMediaRecorder_PTR() != nil {
		convArg0 = source.QMediaRecorder_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QAudioProbe9setSourceEP14QMediaRecorder", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qaudioprobe.h:62
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isActive() const

/*
Returns true if this probe is monitoring something, or false otherwise.

The source being monitored does not need to be active.
*/
func (this *QAudioProbe) IsActive() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QAudioProbe8isActiveEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qaudioprobe.h:65
// index:0
// Public Visibility=Default Availability=Available
// [-2] void audioBufferProbed(const QAudioBuffer &)

/*
This signal should be emitted when an audio buffer is processed in the media service.
*/
func (this *QAudioProbe) AudioBufferProbed(audioBuffer QAudioBuffer_ITF) {
	var convArg0 unsafe.Pointer
	if audioBuffer != nil && audioBuffer.QAudioBuffer_PTR() != nil {
		convArg0 = audioBuffer.QAudioBuffer_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QAudioProbe17audioBufferProbedERK12QAudioBuffer", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudioprobe.h:66
// index:0
// Public Visibility=Default Availability=Available
// [-2] void flush()

/*
This signal should be emitted when it is required to release all buffers. Application must release all outstanding references to audio buffers.
*/
func (this *QAudioProbe) Flush() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QAudioProbe5flushEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
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
