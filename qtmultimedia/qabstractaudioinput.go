package qtmultimedia

// /usr/include/qt/QtMultimedia/qaudiosystem.h
// #include <qaudiosystem.h>
// #include <QtMultimedia>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 26
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
type QAbstractAudioInput struct {
	*qtcore.QObject
}
type QAbstractAudioInput_ITF interface {
	qtcore.QObject_ITF
	QAbstractAudioInput_PTR() *QAbstractAudioInput
}

func (ptr *QAbstractAudioInput) QAbstractAudioInput_PTR() *QAbstractAudioInput { return ptr }

func (this *QAbstractAudioInput) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QAbstractAudioInput) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQAbstractAudioInputFromPointer(cthis unsafe.Pointer) *QAbstractAudioInput {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QAbstractAudioInput{bcthis0}
}
func (*QAbstractAudioInput) NewFromPointer(cthis unsafe.Pointer) *QAbstractAudioInput {
	return NewQAbstractAudioInputFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qaudiosystem.h:106
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QAbstractAudioInput) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QAbstractAudioInput10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qaudiosystem.h:109
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void start(QIODevice *)

/*

 */
func (this *QAbstractAudioInput) Start(device qtcore.QIODevice_ITF /*777 QIODevice **/) {
	var convArg0 unsafe.Pointer
	if device != nil && device.QIODevice_PTR() != nil {
		convArg0 = device.QIODevice_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractAudioInput5startEP9QIODevice", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudiosystem.h:110
// index:1
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QIODevice * start()

/*

 */
func (this *QAbstractAudioInput) Start1() *qtcore.QIODevice /*777 QIODevice **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractAudioInput5startEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQIODeviceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qaudiosystem.h:111
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void stop()

/*

 */
func (this *QAbstractAudioInput) Stop() {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractAudioInput4stopEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudiosystem.h:112
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void reset()

/*

 */
func (this *QAbstractAudioInput) Reset() {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractAudioInput5resetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudiosystem.h:113
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void suspend()

/*

 */
func (this *QAbstractAudioInput) Suspend() {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractAudioInput7suspendEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudiosystem.h:114
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void resume()

/*

 */
func (this *QAbstractAudioInput) Resume() {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractAudioInput6resumeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudiosystem.h:115
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] int bytesReady() const

/*

 */
func (this *QAbstractAudioInput) BytesReady() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QAbstractAudioInput10bytesReadyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qaudiosystem.h:116
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] int periodSize() const

/*

 */
func (this *QAbstractAudioInput) PeriodSize() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QAbstractAudioInput10periodSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qaudiosystem.h:117
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setBufferSize(int)

/*

 */
func (this *QAbstractAudioInput) SetBufferSize(value int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractAudioInput13setBufferSizeEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), value)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudiosystem.h:118
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] int bufferSize() const

/*

 */
func (this *QAbstractAudioInput) BufferSize() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QAbstractAudioInput10bufferSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qaudiosystem.h:119
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setNotifyInterval(int)

/*

 */
func (this *QAbstractAudioInput) SetNotifyInterval(milliSeconds int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractAudioInput17setNotifyIntervalEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), milliSeconds)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudiosystem.h:120
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] int notifyInterval() const

/*

 */
func (this *QAbstractAudioInput) NotifyInterval() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QAbstractAudioInput14notifyIntervalEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qaudiosystem.h:121
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] qint64 processedUSecs() const

/*

 */
func (this *QAbstractAudioInput) ProcessedUSecs() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QAbstractAudioInput14processedUSecsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtMultimedia/qaudiosystem.h:122
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] qint64 elapsedUSecs() const

/*

 */
func (this *QAbstractAudioInput) ElapsedUSecs() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QAbstractAudioInput12elapsedUSecsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtMultimedia/qaudiosystem.h:123
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] QAudio::Error error() const

/*

 */
func (this *QAbstractAudioInput) Error() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QAbstractAudioInput5errorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qaudiosystem.h:124
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] QAudio::State state() const

/*

 */
func (this *QAbstractAudioInput) State() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QAbstractAudioInput5stateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qaudiosystem.h:125
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setFormat(const QAudioFormat &)

/*

 */
func (this *QAbstractAudioInput) SetFormat(fmt_ QAudioFormat_ITF) {
	var convArg0 unsafe.Pointer
	if fmt_ != nil && fmt_.QAudioFormat_PTR() != nil {
		convArg0 = fmt_.QAudioFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractAudioInput9setFormatERK12QAudioFormat", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudiosystem.h:126
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QAudioFormat format() const

/*

 */
func (this *QAbstractAudioInput) Format() *QAudioFormat /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QAbstractAudioInput6formatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQAudioFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQAudioFormat)
	return rv2
}

// /usr/include/qt/QtMultimedia/qaudiosystem.h:127
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setVolume(qreal)

/*

 */
func (this *QAbstractAudioInput) SetVolume(arg0 float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractAudioInput9setVolumeEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudiosystem.h:128
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] qreal volume() const

/*

 */
func (this *QAbstractAudioInput) Volume() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QAbstractAudioInput6volumeEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtMultimedia/qaudiosystem.h:131
// index:0
// Public Visibility=Default Availability=Available
// [-2] void errorChanged(QAudio::Error)

/*

 */
func (this *QAbstractAudioInput) ErrorChanged(error int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractAudioInput12errorChangedEN6QAudio5ErrorE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), error)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudiosystem.h:132
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stateChanged(QAudio::State)

/*

 */
func (this *QAbstractAudioInput) StateChanged(state int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractAudioInput12stateChangedEN6QAudio5StateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), state)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudiosystem.h:133
// index:0
// Public Visibility=Default Availability=Available
// [-2] void notify()

/*

 */
func (this *QAbstractAudioInput) Notify() {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractAudioInput6notifyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

func DeleteQAbstractAudioInput(this *QAbstractAudioInput) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractAudioInputD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_11791() {
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
