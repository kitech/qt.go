package qtmultimedia

// /usr/include/qt/QtMultimedia/qsoundeffect.h
// #include <qsoundeffect.h>
// #include <QtMultimedia>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 11
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
type QSoundEffect struct {
	*qtcore.QObject
}
type QSoundEffect_ITF interface {
	qtcore.QObject_ITF
	QSoundEffect_PTR() *QSoundEffect
}

func (ptr *QSoundEffect) QSoundEffect_PTR() *QSoundEffect { return ptr }

func (this *QSoundEffect) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QSoundEffect) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQSoundEffectFromPointer(cthis unsafe.Pointer) *QSoundEffect {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QSoundEffect{bcthis0}
}
func (*QSoundEffect) NewFromPointer(cthis unsafe.Pointer) *QSoundEffect {
	return NewQSoundEffectFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qsoundeffect.h:56
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QSoundEffect) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QSoundEffect10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qsoundeffect.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSoundEffect(QObject *)

/*
Creates a QSoundEffect with the given parent.
*/
func (*QSoundEffect) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QSoundEffect {
	return NewQSoundEffect(parent)
}
func NewQSoundEffect(parent qtcore.QObject_ITF /*777 QObject **/) *QSoundEffect {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QSoundEffectC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSoundEffectFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSoundEffect")
	return gothis
}

// /usr/include/qt/QtMultimedia/qsoundeffect.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSoundEffect(QObject *)

/*
Creates a QSoundEffect with the given parent.
*/
func (*QSoundEffect) NewForInheritp() *QSoundEffect {
	return NewQSoundEffectp()
}
func NewQSoundEffectp() *QSoundEffect {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN12QSoundEffectC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSoundEffectFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSoundEffect")
	return gothis
}

// /usr/include/qt/QtMultimedia/qsoundeffect.h:84
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QSoundEffect()

/*

 */
func DeleteQSoundEffect(this *QSoundEffect) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QSoundEffectD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qsoundeffect.h:86
// index:0
// Public static Visibility=Default Availability=Available
// [8] QStringList supportedMimeTypes()

/*
Returns a list of the supported mime types for this platform.
*/
func (this *QSoundEffect) SupportedMimeTypes() *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QSoundEffect18supportedMimeTypesEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}
func QSoundEffect_SupportedMimeTypes() *qtcore.QStringList /*123*/ {
	var nilthis *QSoundEffect
	rv := nilthis.SupportedMimeTypes()
	return rv
}

// /usr/include/qt/QtMultimedia/qsoundeffect.h:88
// index:0
// Public Visibility=Default Availability=Available
// [8] QUrl source() const

/*
Returns the URL of the current source to play

Note: Getter function for property source.

See also setSource().
*/
func (this *QSoundEffect) Source() *qtcore.QUrl /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QSoundEffect6sourceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtMultimedia/qsoundeffect.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSource(const QUrl &)

/*
Set the current URL to play to url.

Note: Setter function for property source.

See also source().
*/
func (this *QSoundEffect) SetSource(url qtcore.QUrl_ITF) {
	var convArg0 unsafe.Pointer
	if url != nil && url.QUrl_PTR() != nil {
		convArg0 = url.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QSoundEffect9setSourceERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qsoundeffect.h:91
// index:0
// Public Visibility=Default Availability=Available
// [4] int loopCount() const

/*
Returns the total number of times that this sound effect will be played before stopping.

See the loopsRemaining() method for the number of loops currently remaining.

Note: Getter function for property loops.

See also setLoopCount().
*/
func (this *QSoundEffect) LoopCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QSoundEffect9loopCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qsoundeffect.h:92
// index:0
// Public Visibility=Default Availability=Available
// [4] int loopsRemaining() const

/*

 */
func (this *QSoundEffect) LoopsRemaining() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QSoundEffect14loopsRemainingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qsoundeffect.h:93
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLoopCount(int)

/*
Set the total number of times to play this sound effect to loopCount.

Setting the loop count to 0 or 1 means the sound effect will be played only once; pass QSoundEffect::Infinite to repeat indefinitely. The loop count can be changed while the sound effect is playing, in which case it will update the remaining loops to the new loopCount.

Note: Setter function for property loops.

See also loopCount() and loopsRemaining().
*/
func (this *QSoundEffect) SetLoopCount(loopCount int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QSoundEffect12setLoopCountEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), loopCount)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qsoundeffect.h:95
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal volume() const

/*
Returns the current volume of this sound effect, from 0.0 (silent) to 1.0 (maximum volume).

Note: Getter function for property volume.

See also setVolume().
*/
func (this *QSoundEffect) Volume() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QSoundEffect6volumeEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtMultimedia/qsoundeffect.h:96
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setVolume(qreal)

/*
Sets the sound effect volume to volume.

The volume is scaled linearly from 0.0 (silence) to 1.0 (full volume). Values outside this range will be clamped.

The default volume is 1.0.

UI volume controls should usually be scaled nonlinearly. For example, using a logarithmic scale will produce linear changes in perceived loudness, which is what a user would normally expect from a volume control. See QAudio::convertVolume() for more details.

Note: Setter function for property volume.

See also volume().
*/
func (this *QSoundEffect) SetVolume(volume float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QSoundEffect9setVolumeEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), volume)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qsoundeffect.h:98
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isMuted() const

/*
Returns whether this sound effect is muted

Note: Getter function for property muted.
*/
func (this *QSoundEffect) IsMuted() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QSoundEffect7isMutedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qsoundeffect.h:99
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMuted(bool)

/*
Sets whether to mute this sound effect's playback.

If muted is true, playback will be muted (silenced), and otherwise playback will occur with the currently specified volume().

Note: Setter function for property muted.

See also isMuted().
*/
func (this *QSoundEffect) SetMuted(muted bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QSoundEffect8setMutedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), muted)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qsoundeffect.h:101
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isLoaded() const

/*
Returns whether the sound effect has finished loading the source().
*/
func (this *QSoundEffect) IsLoaded() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QSoundEffect8isLoadedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qsoundeffect.h:103
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isPlaying() const

/*
Returns true if the sound effect is currently playing, or false otherwise

Note: Getter function for property playing.
*/
func (this *QSoundEffect) IsPlaying() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QSoundEffect9isPlayingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qsoundeffect.h:104
// index:0
// Public Visibility=Default Availability=Available
// [4] QSoundEffect::Status status() const

/*
Returns the current status of this sound effect.

Note: Getter function for property status.
*/
func (this *QSoundEffect) Status() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QSoundEffect6statusEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qsoundeffect.h:106
// index:0
// Public Visibility=Default Availability=Available
// [8] QString category() const

/*
Returns the current category for this sound effect.

Some platforms can perform different audio routing for different categories, or may allow the user to set different volume levels for different categories.

This setting will be ignored on platforms that do not support audio categories.

Note: Getter function for property category.

See also setCategory().
*/
func (this *QSoundEffect) Category() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QSoundEffect8categoryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtMultimedia/qsoundeffect.h:107
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCategory(const QString &)

/*
Sets the category of this sound effect to category.

Some platforms can perform different audio routing for different categories, or may allow the user to set different volume levels for different categories.

This setting will be ignored on platforms that do not support audio categories.

If this setting is changed while a sound effect is playing it will only take effect when the sound effect has stopped playing.

Note: Setter function for property category.

See also category().
*/
func (this *QSoundEffect) SetCategory(category string) {
	var tmpArg0 = qtcore.NewQString5(category)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QSoundEffect11setCategoryERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qsoundeffect.h:110
// index:0
// Public Visibility=Default Availability=Available
// [-2] void sourceChanged()

/*
The sourceChanged signal is emitted when the source has been changed.

Note: Notifier signal for property source.
*/
func (this *QSoundEffect) SourceChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QSoundEffect13sourceChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qsoundeffect.h:111
// index:0
// Public Visibility=Default Availability=Available
// [-2] void loopCountChanged()

/*
The loopCountChanged signal is emitted when the initial number of loops has changed.

Note: Notifier signal for property loops.
*/
func (this *QSoundEffect) LoopCountChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QSoundEffect16loopCountChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qsoundeffect.h:112
// index:0
// Public Visibility=Default Availability=Available
// [-2] void loopsRemainingChanged()

/*
The loopsRemainingChanged signal is emitted when the remaining number of loops has changed.

Note: Notifier signal for property loopsRemaining.
*/
func (this *QSoundEffect) LoopsRemainingChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QSoundEffect21loopsRemainingChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qsoundeffect.h:113
// index:0
// Public Visibility=Default Availability=Available
// [-2] void volumeChanged()

/*
The volumeChanged signal is emitted when the volume has changed.

Note: Notifier signal for property volume.
*/
func (this *QSoundEffect) VolumeChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QSoundEffect13volumeChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qsoundeffect.h:114
// index:0
// Public Visibility=Default Availability=Available
// [-2] void mutedChanged()

/*
The mutedChanged signal is emitted when the mute state has changed.

Note: Notifier signal for property muted.
*/
func (this *QSoundEffect) MutedChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QSoundEffect12mutedChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qsoundeffect.h:115
// index:0
// Public Visibility=Default Availability=Available
// [-2] void loadedChanged()

/*
The loadedChanged signal is emitted when the loading state has changed.
*/
func (this *QSoundEffect) LoadedChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QSoundEffect13loadedChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qsoundeffect.h:116
// index:0
// Public Visibility=Default Availability=Available
// [-2] void playingChanged()

/*
The playingChanged signal is emitted when the playing property has changed.

Note: Notifier signal for property playing.
*/
func (this *QSoundEffect) PlayingChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QSoundEffect14playingChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qsoundeffect.h:117
// index:0
// Public Visibility=Default Availability=Available
// [-2] void statusChanged()

/*
The statusChanged signal is emitted when the status property has changed.

Note: Notifier signal for property status.
*/
func (this *QSoundEffect) StatusChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QSoundEffect13statusChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qsoundeffect.h:118
// index:0
// Public Visibility=Default Availability=Available
// [-2] void categoryChanged()

/*
The categoryChanged signal is emitted when the category property has changed.

Note: Notifier signal for property category.
*/
func (this *QSoundEffect) CategoryChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QSoundEffect15categoryChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qsoundeffect.h:121
// index:0
// Public Visibility=Default Availability=Available
// [-2] void play()

/*
Start playback of the sound effect, looping the effect for the number of times as specified in the loops property.
*/
func (this *QSoundEffect) Play() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QSoundEffect4playEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qsoundeffect.h:122
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stop()

/*
Stop current playback.
*/
func (this *QSoundEffect) Stop() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QSoundEffect4stopEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

/*

 */
type QSoundEffect__Loop = int

//
const QSoundEffect__Infinite QSoundEffect__Loop = -2

func (this *QSoundEffect) LoopItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QSoundEffect_LoopItemName(val int) string {
	var nilthis *QSoundEffect
	return nilthis.LoopItemName(val)
}

/*

 */
type QSoundEffect__Status = int

// No source has been set or the source is null.
const QSoundEffect__Null QSoundEffect__Status = 0

// The SoundEffect is trying to load the source.
const QSoundEffect__Loading QSoundEffect__Status = 1

// The source is loaded and ready for play.
const QSoundEffect__Ready QSoundEffect__Status = 2

// An error occurred during operation, such as failure of loading the source.
const QSoundEffect__Error QSoundEffect__Status = 3

func (this *QSoundEffect) StatusItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QSoundEffect_StatusItemName(val int) string {
	var nilthis *QSoundEffect
	return nilthis.StatusItemName(val)
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
