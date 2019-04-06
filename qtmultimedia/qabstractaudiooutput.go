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
// extern C begin: 5
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
type QAbstractAudioOutput struct {
	*qtcore.QObject
}
type QAbstractAudioOutput_ITF interface {
	qtcore.QObject_ITF
	QAbstractAudioOutput_PTR() *QAbstractAudioOutput
}

func (ptr *QAbstractAudioOutput) QAbstractAudioOutput_PTR() *QAbstractAudioOutput { return ptr }

func (this *QAbstractAudioOutput) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QAbstractAudioOutput) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQAbstractAudioOutputFromPointer(cthis unsafe.Pointer) *QAbstractAudioOutput {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QAbstractAudioOutput{bcthis0}
}
func (*QAbstractAudioOutput) NewFromPointer(cthis unsafe.Pointer) *QAbstractAudioOutput {
	return NewQAbstractAudioOutputFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qaudiosystem.h:72
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QAbstractAudioOutput) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QAbstractAudioOutput10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qaudiosystem.h:75
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void start(QIODevice *)

/*

 */
func (this *QAbstractAudioOutput) Start(device qtcore.QIODevice_ITF /*777 QIODevice **/) {
	var convArg0 unsafe.Pointer
	if device != nil && device.QIODevice_PTR() != nil {
		convArg0 = device.QIODevice_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QAbstractAudioOutput5startEP9QIODevice", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudiosystem.h:76
// index:1
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QIODevice * start()

/*

 */
func (this *QAbstractAudioOutput) Start1() *qtcore.QIODevice /*777 QIODevice **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QAbstractAudioOutput5startEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQIODeviceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qaudiosystem.h:77
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void stop()

/*

 */
func (this *QAbstractAudioOutput) Stop() {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QAbstractAudioOutput4stopEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudiosystem.h:78
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void reset()

/*

 */
func (this *QAbstractAudioOutput) Reset() {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QAbstractAudioOutput5resetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudiosystem.h:79
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void suspend()

/*

 */
func (this *QAbstractAudioOutput) Suspend() {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QAbstractAudioOutput7suspendEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudiosystem.h:80
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void resume()

/*

 */
func (this *QAbstractAudioOutput) Resume() {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QAbstractAudioOutput6resumeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudiosystem.h:81
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] int bytesFree() const

/*

 */
func (this *QAbstractAudioOutput) BytesFree() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QAbstractAudioOutput9bytesFreeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qaudiosystem.h:82
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] int periodSize() const

/*

 */
func (this *QAbstractAudioOutput) PeriodSize() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QAbstractAudioOutput10periodSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qaudiosystem.h:83
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setBufferSize(int)

/*

 */
func (this *QAbstractAudioOutput) SetBufferSize(value int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QAbstractAudioOutput13setBufferSizeEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), value)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudiosystem.h:84
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] int bufferSize() const

/*

 */
func (this *QAbstractAudioOutput) BufferSize() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QAbstractAudioOutput10bufferSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qaudiosystem.h:85
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setNotifyInterval(int)

/*

 */
func (this *QAbstractAudioOutput) SetNotifyInterval(milliSeconds int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QAbstractAudioOutput17setNotifyIntervalEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), milliSeconds)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudiosystem.h:86
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] int notifyInterval() const

/*

 */
func (this *QAbstractAudioOutput) NotifyInterval() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QAbstractAudioOutput14notifyIntervalEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qaudiosystem.h:87
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] qint64 processedUSecs() const

/*

 */
func (this *QAbstractAudioOutput) ProcessedUSecs() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QAbstractAudioOutput14processedUSecsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtMultimedia/qaudiosystem.h:88
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] qint64 elapsedUSecs() const

/*

 */
func (this *QAbstractAudioOutput) ElapsedUSecs() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QAbstractAudioOutput12elapsedUSecsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtMultimedia/qaudiosystem.h:89
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] QAudio::Error error() const

/*

 */
func (this *QAbstractAudioOutput) Error() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QAbstractAudioOutput5errorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qaudiosystem.h:90
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] QAudio::State state() const

/*

 */
func (this *QAbstractAudioOutput) State() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QAbstractAudioOutput5stateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qaudiosystem.h:91
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setFormat(const QAudioFormat &)

/*

 */
func (this *QAbstractAudioOutput) SetFormat(fmt_ QAudioFormat_ITF) {
	var convArg0 unsafe.Pointer
	if fmt_ != nil && fmt_.QAudioFormat_PTR() != nil {
		convArg0 = fmt_.QAudioFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QAbstractAudioOutput9setFormatERK12QAudioFormat", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudiosystem.h:92
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QAudioFormat format() const

/*

 */
func (this *QAbstractAudioOutput) Format() *QAudioFormat /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QAbstractAudioOutput6formatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQAudioFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQAudioFormat)
	return rv2
}

// /usr/include/qt/QtMultimedia/qaudiosystem.h:93
// index:0
// Public inline virtual Visibility=Default Availability=Available
// [-2] void setVolume(qreal)

/*

 */
func (this *QAbstractAudioOutput) SetVolume(arg0 float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QAbstractAudioOutput9setVolumeEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudiosystem.h:94
// index:0
// Public inline virtual Visibility=Default Availability=Available
// [8] qreal volume() const

/*

 */
func (this *QAbstractAudioOutput) Volume() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QAbstractAudioOutput6volumeEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtMultimedia/qaudiosystem.h:95
// index:0
// Public inline virtual Visibility=Default Availability=Available
// [8] QString category() const

/*

 */
func (this *QAbstractAudioOutput) Category() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QAbstractAudioOutput8categoryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtMultimedia/qaudiosystem.h:96
// index:0
// Public inline virtual Visibility=Default Availability=Available
// [-2] void setCategory(const QString &)

/*

 */
func (this *QAbstractAudioOutput) SetCategory(arg0 string) {
	var tmpArg0 = qtcore.NewQString5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN20QAbstractAudioOutput11setCategoryERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudiosystem.h:99
// index:0
// Public Visibility=Default Availability=Available
// [-2] void errorChanged(QAudio::Error)

/*

 */
func (this *QAbstractAudioOutput) ErrorChanged(error int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QAbstractAudioOutput12errorChangedEN6QAudio5ErrorE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), error)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudiosystem.h:100
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stateChanged(QAudio::State)

/*

 */
func (this *QAbstractAudioOutput) StateChanged(state int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QAbstractAudioOutput12stateChangedEN6QAudio5StateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), state)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudiosystem.h:101
// index:0
// Public Visibility=Default Availability=Available
// [-2] void notify()

/*

 */
func (this *QAbstractAudioOutput) Notify() {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QAbstractAudioOutput6notifyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

func DeleteQAbstractAudioOutput(this *QAbstractAudioOutput) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QAbstractAudioOutputD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_11789() {
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
