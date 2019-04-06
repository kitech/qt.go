package qtmultimedia

// /usr/include/qt/QtMultimedia/qmediaaudioprobecontrol.h
// #include <qmediaaudioprobecontrol.h>
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
type QMediaAudioProbeControl struct {
	*QMediaControl
}
type QMediaAudioProbeControl_ITF interface {
	QMediaControl_ITF
	QMediaAudioProbeControl_PTR() *QMediaAudioProbeControl
}

func (ptr *QMediaAudioProbeControl) QMediaAudioProbeControl_PTR() *QMediaAudioProbeControl { return ptr }

func (this *QMediaAudioProbeControl) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QMediaControl.GetCthis()
	}
}
func (this *QMediaAudioProbeControl) SetCthis(cthis unsafe.Pointer) {
	this.QMediaControl = NewQMediaControlFromPointer(cthis)
}
func NewQMediaAudioProbeControlFromPointer(cthis unsafe.Pointer) *QMediaAudioProbeControl {
	bcthis0 := NewQMediaControlFromPointer(cthis)
	return &QMediaAudioProbeControl{bcthis0}
}
func (*QMediaAudioProbeControl) NewFromPointer(cthis unsafe.Pointer) *QMediaAudioProbeControl {
	return NewQMediaAudioProbeControlFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qmediaaudioprobecontrol.h:50
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QMediaAudioProbeControl) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QMediaAudioProbeControl10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qmediaaudioprobecontrol.h:52
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QMediaAudioProbeControl()

/*

 */
func DeleteQMediaAudioProbeControl(this *QMediaAudioProbeControl) {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QMediaAudioProbeControlD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qmediaaudioprobecontrol.h:55
// index:0
// Public Visibility=Default Availability=Available
// [-2] void audioBufferProbed(const QAudioBuffer &)

/*
This signal should be emitted when an audio buffer is processed in the media service.
*/
func (this *QMediaAudioProbeControl) AudioBufferProbed(buffer QAudioBuffer_ITF) {
	var convArg0 unsafe.Pointer
	if buffer != nil && buffer.QAudioBuffer_PTR() != nil {
		convArg0 = buffer.QAudioBuffer_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN23QMediaAudioProbeControl17audioBufferProbedERK12QAudioBuffer", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaaudioprobecontrol.h:56
// index:0
// Public Visibility=Default Availability=Available
// [-2] void flush()

/*
This signal should be emitted when it is required to release all frames.
*/
func (this *QMediaAudioProbeControl) Flush() {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QMediaAudioProbeControl5flushEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaaudioprobecontrol.h:59
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QMediaAudioProbeControl(QObject *)

/*
Create a new media audio probe control object with the given parent.
*/
func (*QMediaAudioProbeControl) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QMediaAudioProbeControl {
	return NewQMediaAudioProbeControl(parent)
}
func NewQMediaAudioProbeControl(parent qtcore.QObject_ITF /*777 QObject **/) *QMediaAudioProbeControl {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN23QMediaAudioProbeControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMediaAudioProbeControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QMediaAudioProbeControl")
	return gothis
}

// /usr/include/qt/QtMultimedia/qmediaaudioprobecontrol.h:59
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QMediaAudioProbeControl(QObject *)

/*
Create a new media audio probe control object with the given parent.
*/
func (*QMediaAudioProbeControl) NewForInheritp() *QMediaAudioProbeControl {
	return NewQMediaAudioProbeControlp()
}
func NewQMediaAudioProbeControlp() *QMediaAudioProbeControl {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN23QMediaAudioProbeControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMediaAudioProbeControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QMediaAudioProbeControl")
	return gothis
}

//  body block end

//  keep block begin

func init_unused_11847() {
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
