package qtmultimedia

// /usr/include/qt/QtMultimedia/qmediagaplessplaybackcontrol.h
// #include <qmediagaplessplaybackcontrol.h>
// #include <QtMultimedia>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 15
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
type QMediaGaplessPlaybackControl struct {
	*QMediaControl
}
type QMediaGaplessPlaybackControl_ITF interface {
	QMediaControl_ITF
	QMediaGaplessPlaybackControl_PTR() *QMediaGaplessPlaybackControl
}

func (ptr *QMediaGaplessPlaybackControl) QMediaGaplessPlaybackControl_PTR() *QMediaGaplessPlaybackControl {
	return ptr
}

func (this *QMediaGaplessPlaybackControl) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QMediaControl.GetCthis()
	}
}
func (this *QMediaGaplessPlaybackControl) SetCthis(cthis unsafe.Pointer) {
	this.QMediaControl = NewQMediaControlFromPointer(cthis)
}
func NewQMediaGaplessPlaybackControlFromPointer(cthis unsafe.Pointer) *QMediaGaplessPlaybackControl {
	bcthis0 := NewQMediaControlFromPointer(cthis)
	return &QMediaGaplessPlaybackControl{bcthis0}
}
func (*QMediaGaplessPlaybackControl) NewFromPointer(cthis unsafe.Pointer) *QMediaGaplessPlaybackControl {
	return NewQMediaGaplessPlaybackControlFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qmediagaplessplaybackcontrol.h:52
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QMediaGaplessPlaybackControl) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK28QMediaGaplessPlaybackControl10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qmediagaplessplaybackcontrol.h:54
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QMediaGaplessPlaybackControl()

/*

 */
func DeleteQMediaGaplessPlaybackControl(this *QMediaGaplessPlaybackControl) {
	rv, err := qtrt.InvokeQtFunc6("_ZN28QMediaGaplessPlaybackControlD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qmediagaplessplaybackcontrol.h:56
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QMediaContent nextMedia() const

/*
Returns the content of the next media

See also setNextMedia().
*/
func (this *QMediaGaplessPlaybackControl) NextMedia() *QMediaContent /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK28QMediaGaplessPlaybackControl9nextMediaEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMediaContentFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMediaContent)
	return rv2
}

// /usr/include/qt/QtMultimedia/qmediagaplessplaybackcontrol.h:57
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setNextMedia(const QMediaContent &)

/*
Sets the next media for smooth transition.

See also nextMedia().
*/
func (this *QMediaGaplessPlaybackControl) SetNextMedia(media QMediaContent_ITF) {
	var convArg0 unsafe.Pointer
	if media != nil && media.QMediaContent_PTR() != nil {
		convArg0 = media.QMediaContent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN28QMediaGaplessPlaybackControl12setNextMediaERK13QMediaContent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediagaplessplaybackcontrol.h:59
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool isCrossfadeSupported() const

/*
Indicates whether crossfading is supported or not. If crossfading is not supported, setCrossfadeTime() will be ignored and crossfadeTime() will always return 0.
*/
func (this *QMediaGaplessPlaybackControl) IsCrossfadeSupported() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK28QMediaGaplessPlaybackControl20isCrossfadeSupportedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qmediagaplessplaybackcontrol.h:60
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] qreal crossfadeTime() const

/*
Returns current crossfade time in seconds.

See also setCrossfadeTime().
*/
func (this *QMediaGaplessPlaybackControl) CrossfadeTime() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK28QMediaGaplessPlaybackControl13crossfadeTimeEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtMultimedia/qmediagaplessplaybackcontrol.h:61
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setCrossfadeTime(qreal)

/*
Sets the crossfadeTime in seconds for smooth transition.

Positive value means how much time it will take for the next media to transit from silent to full volume and vice versa for current one. So both current and the next one will be playing during this period of time.

A crossfade time of zero or negative will result in gapless playback (suitable for some continuous media).

See also crossfadeTime().
*/
func (this *QMediaGaplessPlaybackControl) SetCrossfadeTime(crossfadeTime float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN28QMediaGaplessPlaybackControl16setCrossfadeTimeEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), crossfadeTime)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediagaplessplaybackcontrol.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void crossfadeTimeChanged(qreal)

/*
Signals that the crossfadeTime has changed.

See also crossfadeTime().
*/
func (this *QMediaGaplessPlaybackControl) CrossfadeTimeChanged(crossfadeTime float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN28QMediaGaplessPlaybackControl20crossfadeTimeChangedEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), crossfadeTime)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediagaplessplaybackcontrol.h:65
// index:0
// Public Visibility=Default Availability=Available
// [-2] void nextMediaChanged(const QMediaContent &)

/*
Signals that the next media has changed (either explicitly via setNextMedia() or when the player clears the next media while advancing to it).

See also nextMedia().
*/
func (this *QMediaGaplessPlaybackControl) NextMediaChanged(media QMediaContent_ITF) {
	var convArg0 unsafe.Pointer
	if media != nil && media.QMediaContent_PTR() != nil {
		convArg0 = media.QMediaContent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN28QMediaGaplessPlaybackControl16nextMediaChangedERK13QMediaContent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediagaplessplaybackcontrol.h:66
// index:0
// Public Visibility=Default Availability=Available
// [-2] void advancedToNextMedia()

/*
Signals when the player advances to the next media (the content of next media will be cleared).

See also nextMedia().
*/
func (this *QMediaGaplessPlaybackControl) AdvancedToNextMedia() {
	rv, err := qtrt.InvokeQtFunc6("_ZN28QMediaGaplessPlaybackControl19advancedToNextMediaEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediagaplessplaybackcontrol.h:69
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QMediaGaplessPlaybackControl(QObject *)

/*
Constructs a new gapless playback control with the given parent.
*/
func (*QMediaGaplessPlaybackControl) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QMediaGaplessPlaybackControl {
	return NewQMediaGaplessPlaybackControl(parent)
}
func NewQMediaGaplessPlaybackControl(parent qtcore.QObject_ITF /*777 QObject **/) *QMediaGaplessPlaybackControl {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN28QMediaGaplessPlaybackControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMediaGaplessPlaybackControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QMediaGaplessPlaybackControl")
	return gothis
}

// /usr/include/qt/QtMultimedia/qmediagaplessplaybackcontrol.h:69
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QMediaGaplessPlaybackControl(QObject *)

/*
Constructs a new gapless playback control with the given parent.
*/
func (*QMediaGaplessPlaybackControl) NewForInheritp() *QMediaGaplessPlaybackControl {
	return NewQMediaGaplessPlaybackControlp()
}
func NewQMediaGaplessPlaybackControlp() *QMediaGaplessPlaybackControl {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN28QMediaGaplessPlaybackControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMediaGaplessPlaybackControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QMediaGaplessPlaybackControl")
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
