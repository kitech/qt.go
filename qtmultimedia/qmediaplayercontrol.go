package qtmultimedia

// /usr/include/qt/QtMultimedia/qmediaplayercontrol.h
// #include <qmediaplayercontrol.h>
// #include <QtMultimedia>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 22
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
type QMediaPlayerControl struct {
	*QMediaControl
}
type QMediaPlayerControl_ITF interface {
	QMediaControl_ITF
	QMediaPlayerControl_PTR() *QMediaPlayerControl
}

func (ptr *QMediaPlayerControl) QMediaPlayerControl_PTR() *QMediaPlayerControl { return ptr }

func (this *QMediaPlayerControl) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QMediaControl.GetCthis()
	}
}
func (this *QMediaPlayerControl) SetCthis(cthis unsafe.Pointer) {
	this.QMediaControl = NewQMediaControlFromPointer(cthis)
}
func NewQMediaPlayerControlFromPointer(cthis unsafe.Pointer) *QMediaPlayerControl {
	bcthis0 := NewQMediaControlFromPointer(cthis)
	return &QMediaPlayerControl{bcthis0}
}
func (*QMediaPlayerControl) NewFromPointer(cthis unsafe.Pointer) *QMediaPlayerControl {
	return NewQMediaPlayerControlFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qmediaplayercontrol.h:56
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QMediaPlayerControl) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QMediaPlayerControl10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qmediaplayercontrol.h:59
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QMediaPlayerControl()

/*

 */
func DeleteQMediaPlayerControl(this *QMediaPlayerControl) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QMediaPlayerControlD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qmediaplayercontrol.h:61
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] QMediaPlayer::State state() const

/*
Returns the state of a player control.
*/
func (this *QMediaPlayerControl) State() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QMediaPlayerControl5stateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qmediaplayercontrol.h:63
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] QMediaPlayer::MediaStatus mediaStatus() const

/*
Returns the status of the current media.
*/
func (this *QMediaPlayerControl) MediaStatus() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QMediaPlayerControl11mediaStatusEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qmediaplayercontrol.h:65
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] qint64 duration() const

/*
Returns the duration of the current media in milliseconds.
*/
func (this *QMediaPlayerControl) Duration() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QMediaPlayerControl8durationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtMultimedia/qmediaplayercontrol.h:67
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] qint64 position() const

/*
Returns the current playback position in milliseconds.

See also setPosition().
*/
func (this *QMediaPlayerControl) Position() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QMediaPlayerControl8positionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtMultimedia/qmediaplayercontrol.h:68
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setPosition(qint64)

/*
Sets the playback position of the current media. This will initiate a seek and it may take some time for playback to reach the position set.

See also position().
*/
func (this *QMediaPlayerControl) SetPosition(position int64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QMediaPlayerControl11setPositionEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), position)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaplayercontrol.h:70
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] int volume() const

/*
Returns the audio volume of a player control.

See also setVolume().
*/
func (this *QMediaPlayerControl) Volume() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QMediaPlayerControl6volumeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qmediaplayercontrol.h:71
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setVolume(int)

/*
Sets the audio volume of a player control.

The volume is scaled linearly, ranging from 0 (silence) to 100 (full volume).

See also volume().
*/
func (this *QMediaPlayerControl) SetVolume(volume int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QMediaPlayerControl9setVolumeEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), volume)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaplayercontrol.h:73
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool isMuted() const

/*
Returns the mute state of a player control.
*/
func (this *QMediaPlayerControl) IsMuted() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QMediaPlayerControl7isMutedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qmediaplayercontrol.h:74
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setMuted(bool)

/*
Sets the mute state of a player control.

See also isMuted().
*/
func (this *QMediaPlayerControl) SetMuted(mute bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QMediaPlayerControl8setMutedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mute)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaplayercontrol.h:76
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] int bufferStatus() const

/*
Returns the buffering progress of the current media. Progress is measured in the percentage of the buffer filled.
*/
func (this *QMediaPlayerControl) BufferStatus() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QMediaPlayerControl12bufferStatusEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qmediaplayercontrol.h:78
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool isAudioAvailable() const

/*
Identifies if there is audio output available for the current media.

Returns true if audio output is available and false otherwise.
*/
func (this *QMediaPlayerControl) IsAudioAvailable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QMediaPlayerControl16isAudioAvailableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qmediaplayercontrol.h:79
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool isVideoAvailable() const

/*
Identifies if there is video output available for the current media.

Returns true if video output is available and false otherwise.
*/
func (this *QMediaPlayerControl) IsVideoAvailable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QMediaPlayerControl16isVideoAvailableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qmediaplayercontrol.h:81
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool isSeekable() const

/*
Identifies if the current media is seekable.

Returns true if it possible to seek within the current media, and false otherwise.
*/
func (this *QMediaPlayerControl) IsSeekable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QMediaPlayerControl10isSeekableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qmediaplayercontrol.h:83
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QMediaTimeRange availablePlaybackRanges() const

/*
Returns a range of times in milliseconds that can be played back.

Usually for local files this is a continuous interval equal to [0..duration()] or an empty time range if seeking is not supported, but for network sources it refers to the buffered parts of the media.
*/
func (this *QMediaPlayerControl) AvailablePlaybackRanges() *QMediaTimeRange /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QMediaPlayerControl23availablePlaybackRangesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMediaTimeRangeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMediaTimeRange)
	return rv2
}

// /usr/include/qt/QtMultimedia/qmediaplayercontrol.h:85
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] qreal playbackRate() const

/*
Returns the rate of playback.

See also setPlaybackRate().
*/
func (this *QMediaPlayerControl) PlaybackRate() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QMediaPlayerControl12playbackRateEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtMultimedia/qmediaplayercontrol.h:86
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setPlaybackRate(qreal)

/*
Sets the rate of playback.

See also playbackRate().
*/
func (this *QMediaPlayerControl) SetPlaybackRate(rate float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QMediaPlayerControl15setPlaybackRateEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), rate)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaplayercontrol.h:88
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QMediaContent media() const

/*
Returns the current media source.

See also setMedia().
*/
func (this *QMediaPlayerControl) Media() *QMediaContent /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QMediaPlayerControl5mediaEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMediaContentFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMediaContent)
	return rv2
}

// /usr/include/qt/QtMultimedia/qmediaplayercontrol.h:89
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] const QIODevice * mediaStream() const

/*
Returns the current media stream. This is only a valid if a stream was passed to setMedia().

See also setMedia().
*/
func (this *QMediaPlayerControl) MediaStream() *qtcore.QIODevice /*777 const QIODevice **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QMediaPlayerControl11mediaStreamEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQIODeviceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qmediaplayercontrol.h:90
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setMedia(const QMediaContent &, QIODevice *)

/*
Sets the current media source. If a stream is supplied; data will be read from that instead of attempting to resolve the media source. The media source may still be used to supply media information such as mime type.

Setting the media to a null QMediaContent will cause the control to discard all information relating to the current media source and to cease all I/O operations related to that media.

Qt resource files are never passed as is. If the service supports QMediaServiceProviderHint::StreamPlayback, a stream is supplied, pointing to an opened QFile. Otherwise, the resource is copied into a temporary file and media contains the url to that file.

See also media().
*/
func (this *QMediaPlayerControl) SetMedia(media QMediaContent_ITF, stream qtcore.QIODevice_ITF /*777 QIODevice **/) {
	var convArg0 unsafe.Pointer
	if media != nil && media.QMediaContent_PTR() != nil {
		convArg0 = media.QMediaContent_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if stream != nil && stream.QIODevice_PTR() != nil {
		convArg1 = stream.QIODevice_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QMediaPlayerControl8setMediaERK13QMediaContentP9QIODevice", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaplayercontrol.h:92
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void play()

/*
Starts playback of the current media.

If successful the player control will immediately enter the playing state.

See also state().
*/
func (this *QMediaPlayerControl) Play() {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QMediaPlayerControl4playEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaplayercontrol.h:93
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void pause()

/*
Pauses playback of the current media.

If successful the player control will immediately enter the paused state.

See also state(), play(), and stop().
*/
func (this *QMediaPlayerControl) Pause() {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QMediaPlayerControl5pauseEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaplayercontrol.h:94
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void stop()

/*
Stops playback of the current media.

If successful the player control will immediately enter the stopped state.
*/
func (this *QMediaPlayerControl) Stop() {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QMediaPlayerControl4stopEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaplayercontrol.h:97
// index:0
// Public Visibility=Default Availability=Available
// [-2] void mediaChanged(const QMediaContent &)

/*
Signals that the current media content has changed.
*/
func (this *QMediaPlayerControl) MediaChanged(content QMediaContent_ITF) {
	var convArg0 unsafe.Pointer
	if content != nil && content.QMediaContent_PTR() != nil {
		convArg0 = content.QMediaContent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QMediaPlayerControl12mediaChangedERK13QMediaContent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaplayercontrol.h:98
// index:0
// Public Visibility=Default Availability=Available
// [-2] void durationChanged(qint64)

/*
Signals that the duration of the current media has changed.

See also duration().
*/
func (this *QMediaPlayerControl) DurationChanged(duration int64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QMediaPlayerControl15durationChangedEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), duration)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaplayercontrol.h:99
// index:0
// Public Visibility=Default Availability=Available
// [-2] void positionChanged(qint64)

/*
Signals the playback position has changed.

This is only emitted in when there has been a discontinous change in the playback postion, such as a seek or the position being reset.

See also position().
*/
func (this *QMediaPlayerControl) PositionChanged(position int64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QMediaPlayerControl15positionChangedEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), position)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaplayercontrol.h:100
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stateChanged(QMediaPlayer::State)

/*
Signals that the state of a player control has changed to newState.

See also state().
*/
func (this *QMediaPlayerControl) StateChanged(newState int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QMediaPlayerControl12stateChangedEN12QMediaPlayer5StateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), newState)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaplayercontrol.h:101
// index:0
// Public Visibility=Default Availability=Available
// [-2] void mediaStatusChanged(QMediaPlayer::MediaStatus)

/*
Signals that the status of the current media has changed.

See also mediaStatus().
*/
func (this *QMediaPlayerControl) MediaStatusChanged(status int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QMediaPlayerControl18mediaStatusChangedEN12QMediaPlayer11MediaStatusE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), status)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaplayercontrol.h:102
// index:0
// Public Visibility=Default Availability=Available
// [-2] void volumeChanged(int)

/*
Signals the audio volume of a player control has changed.

See also volume().
*/
func (this *QMediaPlayerControl) VolumeChanged(volume int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QMediaPlayerControl13volumeChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), volume)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaplayercontrol.h:103
// index:0
// Public Visibility=Default Availability=Available
// [-2] void mutedChanged(bool)

/*
Signals a change in the mute status of a player control.

See also isMuted().
*/
func (this *QMediaPlayerControl) MutedChanged(mute bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QMediaPlayerControl12mutedChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mute)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaplayercontrol.h:104
// index:0
// Public Visibility=Default Availability=Available
// [-2] void audioAvailableChanged(bool)

/*
Signals that there has been a change in the availability of audio output audioAvailable.

See also isAudioAvailable().
*/
func (this *QMediaPlayerControl) AudioAvailableChanged(audioAvailable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QMediaPlayerControl21audioAvailableChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), audioAvailable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaplayercontrol.h:105
// index:0
// Public Visibility=Default Availability=Available
// [-2] void videoAvailableChanged(bool)

/*
Signal that the availability of visual content has changed to videoAvailable.

See also isVideoAvailable().
*/
func (this *QMediaPlayerControl) VideoAvailableChanged(videoAvailable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QMediaPlayerControl21videoAvailableChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), videoAvailable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaplayercontrol.h:106
// index:0
// Public Visibility=Default Availability=Available
// [-2] void bufferStatusChanged(int)

/*
Signal the amount of the local buffer filled as a percentage by percentFilled.

See also bufferStatus().
*/
func (this *QMediaPlayerControl) BufferStatusChanged(percentFilled int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QMediaPlayerControl19bufferStatusChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), percentFilled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaplayercontrol.h:107
// index:0
// Public Visibility=Default Availability=Available
// [-2] void seekableChanged(bool)

/*
Signals that the seekable state of a player control has changed.

See also isSeekable().
*/
func (this *QMediaPlayerControl) SeekableChanged(seekable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QMediaPlayerControl15seekableChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), seekable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaplayercontrol.h:108
// index:0
// Public Visibility=Default Availability=Available
// [-2] void availablePlaybackRangesChanged(const QMediaTimeRange &)

/*
Signals that the available media playback ranges have changed.

See also QMediaPlayerControl::availablePlaybackRanges().
*/
func (this *QMediaPlayerControl) AvailablePlaybackRangesChanged(ranges QMediaTimeRange_ITF) {
	var convArg0 unsafe.Pointer
	if ranges != nil && ranges.QMediaTimeRange_PTR() != nil {
		convArg0 = ranges.QMediaTimeRange_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QMediaPlayerControl30availablePlaybackRangesChangedERK15QMediaTimeRange", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaplayercontrol.h:109
// index:0
// Public Visibility=Default Availability=Available
// [-2] void playbackRateChanged(qreal)

/*
Signal emitted when playback rate changes to rate.
*/
func (this *QMediaPlayerControl) PlaybackRateChanged(rate float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QMediaPlayerControl19playbackRateChangedEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), rate)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaplayercontrol.h:110
// index:0
// Public Visibility=Default Availability=Available
// [-2] void error(int, const QString &)

/*
Signals that an error has occurred. The errorString provides a more detailed explanation.
*/
func (this *QMediaPlayerControl) Error(error int, errorString string) {
	var tmpArg1 = qtcore.NewQString5(errorString)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN19QMediaPlayerControl5errorEiRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), error, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaplayercontrol.h:113
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QMediaPlayerControl(QObject *)

/*
Constructs a new media player control with the given parent.
*/
func (*QMediaPlayerControl) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QMediaPlayerControl {
	return NewQMediaPlayerControl(parent)
}
func NewQMediaPlayerControl(parent qtcore.QObject_ITF /*777 QObject **/) *QMediaPlayerControl {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QMediaPlayerControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMediaPlayerControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QMediaPlayerControl")
	return gothis
}

// /usr/include/qt/QtMultimedia/qmediaplayercontrol.h:113
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QMediaPlayerControl(QObject *)

/*
Constructs a new media player control with the given parent.
*/
func (*QMediaPlayerControl) NewForInheritp() *QMediaPlayerControl {
	return NewQMediaPlayerControlp()
}
func NewQMediaPlayerControlp() *QMediaPlayerControl {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN19QMediaPlayerControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMediaPlayerControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QMediaPlayerControl")
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
