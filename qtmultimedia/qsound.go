package qtmultimedia

// /usr/include/qt/QtMultimedia/qsound.h
// #include <qsound.h>
// #include <QtMultimedia>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 39
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
type QSound struct {
	*qtcore.QObject
}
type QSound_ITF interface {
	qtcore.QObject_ITF
	QSound_PTR() *QSound
}

func (ptr *QSound) QSound_PTR() *QSound { return ptr }

func (this *QSound) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QSound) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQSoundFromPointer(cthis unsafe.Pointer) *QSound {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QSound{bcthis0}
}
func (*QSound) NewFromPointer(cthis unsafe.Pointer) *QSound {
	return NewQSoundFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qsound.h:52
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QSound) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QSound10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qsound.h:59
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void play(const QString &)

/*
Plays the sound stored in the file specified by the given filename.

The file can either be a local file or in a resource.

See also stop(), loopsRemaining(), and isFinished().
*/
func (this *QSound) Play(filename string) {
	var tmpArg0 = qtcore.NewQString5(filename)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN6QSound4playERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}
func QSound_Play(filename string) {
	var nilthis *QSound
	nilthis.Play(filename)
}

// /usr/include/qt/QtMultimedia/qsound.h:72
// index:1
// Public Visibility=Default Availability=Available
// [-2] void play()

/*
Plays the sound stored in the file specified by the given filename.

The file can either be a local file or in a resource.

See also stop(), loopsRemaining(), and isFinished().
*/
func (this *QSound) Play1() {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QSound4playEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qsound.h:61
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSound(const QString &, QObject *)

/*
Constructs a QSound object from the file specified by the given filename and with the given parent.

The file can either be a local file or in a resource.

See also play().
*/
func (*QSound) NewForInherit(filename string, parent qtcore.QObject_ITF /*777 QObject **/) *QSound {
	return NewQSound(filename, parent)
}
func NewQSound(filename string, parent qtcore.QObject_ITF /*777 QObject **/) *QSound {
	var tmpArg0 = qtcore.NewQString5(filename)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg1 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QSoundC2ERK7QStringP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSoundFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSound")
	return gothis
}

// /usr/include/qt/QtMultimedia/qsound.h:61
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSound(const QString &, QObject *)

/*
Constructs a QSound object from the file specified by the given filename and with the given parent.

The file can either be a local file or in a resource.

See also play().
*/
func (*QSound) NewForInheritp(filename string) *QSound {
	return NewQSoundp(filename)
}
func NewQSoundp(filename string) *QSound {
	var tmpArg0 = qtcore.NewQString5(filename)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QObject *=Pointer, QObject=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN6QSoundC2ERK7QStringP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSoundFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSound")
	return gothis
}

// /usr/include/qt/QtMultimedia/qsound.h:62
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QSound()

/*

 */
func DeleteQSound(this *QSound) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QSoundD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qsound.h:64
// index:0
// Public Visibility=Default Availability=Available
// [4] int loops() const

/*
Returns the number of times the sound will play. Return value of QSound::Infinite indicates infinite number of loops

See also loopsRemaining() and setLoops().
*/
func (this *QSound) Loops() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QSound5loopsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qsound.h:65
// index:0
// Public Visibility=Default Availability=Available
// [4] int loopsRemaining() const

/*
Returns the remaining number of times the sound will loop (for all positive values this value decreases each time the sound is played). Return value of QSound::Infinite indicates infinite number of loops

See also loops() and isFinished().
*/
func (this *QSound) LoopsRemaining() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QSound14loopsRemainingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qsound.h:66
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLoops(int)

/*
Sets the sound to repeat the given number of times when it is played.

Note that passing the value QSound::Infinite will cause the sound to loop indefinitely.

See also loops().
*/
func (this *QSound) SetLoops(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QSound8setLoopsEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qsound.h:67
// index:0
// Public Visibility=Default Availability=Available
// [8] QString fileName() const

/*
Returns the filename associated with this QSound object.

See also QSound().
*/
func (this *QSound) FileName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QSound8fileNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtMultimedia/qsound.h:69
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isFinished() const

/*
Returns true if the sound has finished playing; otherwise returns false.
*/
func (this *QSound) IsFinished() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QSound10isFinishedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qsound.h:73
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stop()

/*
Stops the sound playing.

See also play().
*/
func (this *QSound) Stop() {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QSound4stopEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

/*

 */
type QSound__Loop = int

//
const QSound__Infinite QSound__Loop = -1

func (this *QSound) LoopItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QSound_LoopItemName(val int) string {
	var nilthis *QSound
	return nilthis.LoopItemName(val)
}

//  body block end

//  keep block begin

func init_unused_11905() {
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
