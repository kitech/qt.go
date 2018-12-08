package qtmultimedia

// /usr/include/qt/QtMultimedia/qmediaplayer.h
// #include <qmediaplayer.h>
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
type QMediaPlayer struct {
	*QMediaObject
}
type QMediaPlayer_ITF interface {
	QMediaObject_ITF
	QMediaPlayer_PTR() *QMediaPlayer
}

func (ptr *QMediaPlayer) QMediaPlayer_PTR() *QMediaPlayer { return ptr }

func (this *QMediaPlayer) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QMediaObject.GetCthis()
	}
}
func (this *QMediaPlayer) SetCthis(cthis unsafe.Pointer) {
	this.QMediaObject = NewQMediaObjectFromPointer(cthis)
}
func NewQMediaPlayerFromPointer(cthis unsafe.Pointer) *QMediaPlayer {
	bcthis0 := NewQMediaObjectFromPointer(cthis)
	return &QMediaPlayer{bcthis0}
}
func (*QMediaPlayer) NewFromPointer(cthis unsafe.Pointer) *QMediaPlayer {
	return NewQMediaPlayerFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qmediaplayer.h:61
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QMediaPlayer) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QMediaPlayer10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qmediaplayer.h:122
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QMediaPlayer(QObject *, QMediaPlayer::Flags)

/*
Construct a QMediaPlayer instance parented to parent and with flags.
*/
func (*QMediaPlayer) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/, flags int) *QMediaPlayer {
	return NewQMediaPlayer(parent, flags)
}
func NewQMediaPlayer(parent qtcore.QObject_ITF /*777 QObject **/, flags int) *QMediaPlayer {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QMediaPlayerC2EP7QObject6QFlagsINS_4FlagEE", qtrt.FFI_TYPE_POINTER, convArg0, flags)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMediaPlayerFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QMediaPlayer")
	return gothis
}

// /usr/include/qt/QtMultimedia/qmediaplayer.h:122
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QMediaPlayer(QObject *, QMediaPlayer::Flags)

/*
Construct a QMediaPlayer instance parented to parent and with flags.
*/
func (*QMediaPlayer) NewForInheritp() *QMediaPlayer {
	return NewQMediaPlayerp()
}
func NewQMediaPlayerp() *QMediaPlayer {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	// arg: 1, QMediaPlayer::Flags=Typedef, QMediaPlayer::Flags=Typedef, QFlags<QMediaPlayer::Flag>, Unexposed
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN12QMediaPlayerC2EP7QObject6QFlagsINS_4FlagEE", qtrt.FFI_TYPE_POINTER, convArg0, flags)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMediaPlayerFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QMediaPlayer")
	return gothis
}

// /usr/include/qt/QtMultimedia/qmediaplayer.h:122
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QMediaPlayer(QObject *, QMediaPlayer::Flags)

/*
Construct a QMediaPlayer instance parented to parent and with flags.
*/
func (*QMediaPlayer) NewForInheritp1(parent qtcore.QObject_ITF /*777 QObject **/) *QMediaPlayer {
	return NewQMediaPlayerp1(parent)
}
func NewQMediaPlayerp1(parent qtcore.QObject_ITF /*777 QObject **/) *QMediaPlayer {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	// arg: 1, QMediaPlayer::Flags=Typedef, QMediaPlayer::Flags=Typedef, QFlags<QMediaPlayer::Flag>, Unexposed
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN12QMediaPlayerC2EP7QObject6QFlagsINS_4FlagEE", qtrt.FFI_TYPE_POINTER, convArg0, flags)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMediaPlayerFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QMediaPlayer")
	return gothis
}

// /usr/include/qt/QtMultimedia/qmediaplayer.h:123
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QMediaPlayer()

/*

 */
func DeleteQMediaPlayer(this *QMediaPlayer) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QMediaPlayerD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qmediaplayer.h:125
// index:0
// Public static Visibility=Default Availability=Available
// [4] QMultimedia::SupportEstimate hasSupport(const QString &, const QStringList &, QMediaPlayer::Flags)

/*
Returns the level of support a media player has for a mimeType and a set of codecs.

The flags argument allows additional requirements such as performance indicators to be specified.
*/
func (this *QMediaPlayer) HasSupport(mimeType string, codecs qtcore.QStringList_ITF, flags int) int {
	var tmpArg0 = qtcore.NewQString5(mimeType)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if codecs != nil && codecs.QStringList_PTR() != nil {
		convArg1 = codecs.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QMediaPlayer10hasSupportERK7QStringRK11QStringList6QFlagsINS_4FlagEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, flags)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}
func QMediaPlayer_HasSupport(mimeType string, codecs qtcore.QStringList_ITF, flags int) int {
	var nilthis *QMediaPlayer
	rv := nilthis.HasSupport(mimeType, codecs, flags)
	return rv
}

// /usr/include/qt/QtMultimedia/qmediaplayer.h:125
// index:0
// Public static Visibility=Default Availability=Available
// [4] QMultimedia::SupportEstimate hasSupport(const QString &, const QStringList &, QMediaPlayer::Flags)

/*
Returns the level of support a media player has for a mimeType and a set of codecs.

The flags argument allows additional requirements such as performance indicators to be specified.
*/
func (this *QMediaPlayer) HasSupportp(mimeType string) int {
	var tmpArg0 = qtcore.NewQString5(mimeType)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, const QStringList &=LValueReference, QStringList=Record, , Invalid
	var convArg1 unsafe.Pointer
	// arg: 2, QMediaPlayer::Flags=Typedef, QMediaPlayer::Flags=Typedef, QFlags<QMediaPlayer::Flag>, Unexposed
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN12QMediaPlayer10hasSupportERK7QStringRK11QStringList6QFlagsINS_4FlagEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, flags)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qmediaplayer.h:125
// index:0
// Public static Visibility=Default Availability=Available
// [4] QMultimedia::SupportEstimate hasSupport(const QString &, const QStringList &, QMediaPlayer::Flags)

/*
Returns the level of support a media player has for a mimeType and a set of codecs.

The flags argument allows additional requirements such as performance indicators to be specified.
*/
func (this *QMediaPlayer) HasSupportp1(mimeType string, codecs qtcore.QStringList_ITF) int {
	var tmpArg0 = qtcore.NewQString5(mimeType)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if codecs != nil && codecs.QStringList_PTR() != nil {
		convArg1 = codecs.QStringList_PTR().GetCthis()
	}
	// arg: 2, QMediaPlayer::Flags=Typedef, QMediaPlayer::Flags=Typedef, QFlags<QMediaPlayer::Flag>, Unexposed
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN12QMediaPlayer10hasSupportERK7QStringRK11QStringList6QFlagsINS_4FlagEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, flags)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qmediaplayer.h:128
// index:0
// Public static Visibility=Default Availability=Available
// [8] QStringList supportedMimeTypes(QMediaPlayer::Flags)

/*

 */
func (this *QMediaPlayer) SupportedMimeTypes(flags int) *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QMediaPlayer18supportedMimeTypesE6QFlagsINS_4FlagEE", qtrt.FFI_TYPE_POINTER, flags)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}
func QMediaPlayer_SupportedMimeTypes(flags int) *qtcore.QStringList /*123*/ {
	var nilthis *QMediaPlayer
	rv := nilthis.SupportedMimeTypes(flags)
	return rv
}

// /usr/include/qt/QtMultimedia/qmediaplayer.h:128
// index:0
// Public static Visibility=Default Availability=Available
// [8] QStringList supportedMimeTypes(QMediaPlayer::Flags)

/*

 */
func (this *QMediaPlayer) SupportedMimeTypesp() *qtcore.QStringList /*123*/ {
	// arg: 0, QMediaPlayer::Flags=Typedef, QMediaPlayer::Flags=Typedef, QFlags<QMediaPlayer::Flag>, Unexposed
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN12QMediaPlayer18supportedMimeTypesE6QFlagsINS_4FlagEE", qtrt.FFI_TYPE_POINTER, flags)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtMultimedia/qmediaplayer.h:132
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setVideoOutput(QAbstractVideoSurface *)

/*
Attach a QVideoWidget video output to the media player.

If the media player has already video output attached, it will be replaced with a new one.
*/
func (this *QMediaPlayer) SetVideoOutput(surface QAbstractVideoSurface_ITF /*777 QAbstractVideoSurface **/) {
	var convArg0 unsafe.Pointer
	if surface != nil && surface.QAbstractVideoSurface_PTR() != nil {
		convArg0 = surface.QAbstractVideoSurface_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QMediaPlayer14setVideoOutputEP21QAbstractVideoSurface", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaplayer.h:134
// index:0
// Public Visibility=Default Availability=Available
// [8] QMediaContent media() const

/*

 */
func (this *QMediaPlayer) Media() *QMediaContent /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QMediaPlayer5mediaEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMediaContentFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMediaContent)
	return rv2
}

// /usr/include/qt/QtMultimedia/qmediaplayer.h:135
// index:0
// Public Visibility=Default Availability=Available
// [8] const QIODevice * mediaStream() const

/*
Returns the stream source of media data.

This is only valid if a stream was passed to setMedia().

See also setMedia().
*/
func (this *QMediaPlayer) MediaStream() *qtcore.QIODevice /*777 const QIODevice **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QMediaPlayer11mediaStreamEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQIODeviceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qmediaplayer.h:136
// index:0
// Public Visibility=Default Availability=Available
// [8] QMediaPlaylist * playlist() const

/*

 */
func (this *QMediaPlayer) Playlist() *QMediaPlaylist /*777 QMediaPlaylist **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QMediaPlayer8playlistEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMediaPlaylistFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qmediaplayer.h:137
// index:0
// Public Visibility=Default Availability=Available
// [8] QMediaContent currentMedia() const

/*

 */
func (this *QMediaPlayer) CurrentMedia() *QMediaContent /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QMediaPlayer12currentMediaEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMediaContentFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMediaContent)
	return rv2
}

// /usr/include/qt/QtMultimedia/qmediaplayer.h:139
// index:0
// Public Visibility=Default Availability=Available
// [4] QMediaPlayer::State state() const

/*

 */
func (this *QMediaPlayer) State() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QMediaPlayer5stateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qmediaplayer.h:140
// index:0
// Public Visibility=Default Availability=Available
// [4] QMediaPlayer::MediaStatus mediaStatus() const

/*

 */
func (this *QMediaPlayer) MediaStatus() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QMediaPlayer11mediaStatusEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qmediaplayer.h:142
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 duration() const

/*

 */
func (this *QMediaPlayer) Duration() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QMediaPlayer8durationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtMultimedia/qmediaplayer.h:143
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 position() const

/*

 */
func (this *QMediaPlayer) Position() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QMediaPlayer8positionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtMultimedia/qmediaplayer.h:145
// index:0
// Public Visibility=Default Availability=Available
// [4] int volume() const

/*

 */
func (this *QMediaPlayer) Volume() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QMediaPlayer6volumeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qmediaplayer.h:146
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isMuted() const

/*

 */
func (this *QMediaPlayer) IsMuted() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QMediaPlayer7isMutedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qmediaplayer.h:147
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isAudioAvailable() const

/*

 */
func (this *QMediaPlayer) IsAudioAvailable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QMediaPlayer16isAudioAvailableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qmediaplayer.h:148
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isVideoAvailable() const

/*

 */
func (this *QMediaPlayer) IsVideoAvailable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QMediaPlayer16isVideoAvailableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qmediaplayer.h:150
// index:0
// Public Visibility=Default Availability=Available
// [4] int bufferStatus() const

/*

 */
func (this *QMediaPlayer) BufferStatus() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QMediaPlayer12bufferStatusEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qmediaplayer.h:152
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isSeekable() const

/*

 */
func (this *QMediaPlayer) IsSeekable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QMediaPlayer10isSeekableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qmediaplayer.h:153
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal playbackRate() const

/*

 */
func (this *QMediaPlayer) PlaybackRate() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QMediaPlayer12playbackRateEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtMultimedia/qmediaplayer.h:155
// index:0
// Public Visibility=Default Availability=Available
// [4] QMediaPlayer::Error error() const

/*
Returns the current error state.
*/
func (this *QMediaPlayer) Error() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QMediaPlayer5errorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qmediaplayer.h:204
// index:1
// Public Visibility=Default Availability=Available
// [-2] void error(QMediaPlayer::Error)

/*
Returns the current error state.
*/
func (this *QMediaPlayer) Error1(error int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QMediaPlayer5errorENS_5ErrorE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), error)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaplayer.h:156
// index:0
// Public Visibility=Default Availability=Available
// [8] QString errorString() const

/*

 */
func (this *QMediaPlayer) ErrorString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QMediaPlayer11errorStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtMultimedia/qmediaplayer.h:158
// index:0
// Public Visibility=Default Availability=Available
// [8] QNetworkConfiguration currentNetworkConfiguration() const

/*
Returns the current network access point in use. If a default contructed QNetworkConfiguration is returned this feature is not available or that none of the current supplied configurations are in use.
*/
func (this *QMediaPlayer) CurrentNetworkConfiguration() *qtnetwork.QNetworkConfiguration /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QMediaPlayer27currentNetworkConfigurationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtnetwork.NewQNetworkConfigurationFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtnetwork.DeleteQNetworkConfiguration)
	return rv2
}

// /usr/include/qt/QtMultimedia/qmediaplayer.h:160
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] QMultimedia::AvailabilityStatus availability() const

/*
Reimplemented from QMediaObject::availability().
*/
func (this *QMediaPlayer) Availability() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QMediaPlayer12availabilityEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qmediaplayer.h:162
// index:0
// Public Visibility=Default Availability=Available
// [4] QAudio::Role audioRole() const

/*

 */
func (this *QMediaPlayer) AudioRole() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QMediaPlayer9audioRoleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qmediaplayer.h:163
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAudioRole(QAudio::Role)

/*

 */
func (this *QMediaPlayer) SetAudioRole(audioRole int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QMediaPlayer12setAudioRoleEN6QAudio4RoleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), audioRole)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaplayer.h:167
// index:0
// Public Visibility=Default Availability=Available
// [-2] void play()

/*
Start or resume playing the current source.
*/
func (this *QMediaPlayer) Play() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QMediaPlayer4playEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaplayer.h:168
// index:0
// Public Visibility=Default Availability=Available
// [-2] void pause()

/*
Pause playing the current source.
*/
func (this *QMediaPlayer) Pause() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QMediaPlayer5pauseEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaplayer.h:169
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stop()

/*
Stop playing, and reset the play position to the beginning.
*/
func (this *QMediaPlayer) Stop() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QMediaPlayer4stopEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaplayer.h:171
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPosition(qint64)

/*

 */
func (this *QMediaPlayer) SetPosition(position int64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QMediaPlayer11setPositionEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), position)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaplayer.h:172
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setVolume(int)

/*

 */
func (this *QMediaPlayer) SetVolume(volume int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QMediaPlayer9setVolumeEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), volume)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaplayer.h:173
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMuted(bool)

/*

 */
func (this *QMediaPlayer) SetMuted(muted bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QMediaPlayer8setMutedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), muted)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaplayer.h:175
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPlaybackRate(qreal)

/*

 */
func (this *QMediaPlayer) SetPlaybackRate(rate float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QMediaPlayer15setPlaybackRateEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), rate)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaplayer.h:177
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMedia(const QMediaContent &, QIODevice *)

/*
Sets the current media source.

If a stream is supplied; media data will be read from it instead of resolving the media source. In this case the media source may still be used to resolve additional information about the media such as mime type. The stream must be open and readable.

Setting the media to a null QMediaContent will cause the player to discard all information relating to the current media source and to cease all I/O operations related to that media.

Note: This function returns immediately after recording the specified source of the media. It does not wait for the media to finish loading and does not check for errors. Listen for the mediaStatusChanged() and error() signals to be notified when the media is loaded and when an error occurs during loading.

Note: Setter function for property media.

See also media().
*/
func (this *QMediaPlayer) SetMedia(media QMediaContent_ITF, stream qtcore.QIODevice_ITF /*777 QIODevice **/) {
	var convArg0 unsafe.Pointer
	if media != nil && media.QMediaContent_PTR() != nil {
		convArg0 = media.QMediaContent_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if stream != nil && stream.QIODevice_PTR() != nil {
		convArg1 = stream.QIODevice_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QMediaPlayer8setMediaERK13QMediaContentP9QIODevice", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaplayer.h:177
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMedia(const QMediaContent &, QIODevice *)

/*
Sets the current media source.

If a stream is supplied; media data will be read from it instead of resolving the media source. In this case the media source may still be used to resolve additional information about the media such as mime type. The stream must be open and readable.

Setting the media to a null QMediaContent will cause the player to discard all information relating to the current media source and to cease all I/O operations related to that media.

Note: This function returns immediately after recording the specified source of the media. It does not wait for the media to finish loading and does not check for errors. Listen for the mediaStatusChanged() and error() signals to be notified when the media is loaded and when an error occurs during loading.

Note: Setter function for property media.

See also media().
*/
func (this *QMediaPlayer) SetMediap(media QMediaContent_ITF) {
	var convArg0 unsafe.Pointer
	if media != nil && media.QMediaContent_PTR() != nil {
		convArg0 = media.QMediaContent_PTR().GetCthis()
	}
	// arg: 1, QIODevice *=Pointer, QIODevice=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN12QMediaPlayer8setMediaERK13QMediaContentP9QIODevice", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaplayer.h:178
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPlaylist(QMediaPlaylist *)

/*

 */
func (this *QMediaPlayer) SetPlaylist(playlist QMediaPlaylist_ITF /*777 QMediaPlaylist **/) {
	var convArg0 unsafe.Pointer
	if playlist != nil && playlist.QMediaPlaylist_PTR() != nil {
		convArg0 = playlist.QMediaPlaylist_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QMediaPlayer11setPlaylistEP14QMediaPlaylist", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaplayer.h:183
// index:0
// Public Visibility=Default Availability=Available
// [-2] void mediaChanged(const QMediaContent &)

/*
Signals that the media source has been changed to media.

Note: Notifier signal for property media.

See also media() and currentMediaChanged().
*/
func (this *QMediaPlayer) MediaChanged(media QMediaContent_ITF) {
	var convArg0 unsafe.Pointer
	if media != nil && media.QMediaContent_PTR() != nil {
		convArg0 = media.QMediaContent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QMediaPlayer12mediaChangedERK13QMediaContent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaplayer.h:184
// index:0
// Public Visibility=Default Availability=Available
// [-2] void currentMediaChanged(const QMediaContent &)

/*
Signals that the current playing content has been changed to media.

Note: Notifier signal for property currentMedia.

See also currentMedia() and mediaChanged().
*/
func (this *QMediaPlayer) CurrentMediaChanged(media QMediaContent_ITF) {
	var convArg0 unsafe.Pointer
	if media != nil && media.QMediaContent_PTR() != nil {
		convArg0 = media.QMediaContent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QMediaPlayer19currentMediaChangedERK13QMediaContent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaplayer.h:186
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stateChanged(QMediaPlayer::State)

/*
Signal the state of the Player object has changed.

Note: Notifier signal for property state.
*/
func (this *QMediaPlayer) StateChanged(newState int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QMediaPlayer12stateChangedENS_5StateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), newState)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaplayer.h:187
// index:0
// Public Visibility=Default Availability=Available
// [-2] void mediaStatusChanged(QMediaPlayer::MediaStatus)

/*
Signals that the status of the current media has changed.

Note: Notifier signal for property mediaStatus.

See also mediaStatus().
*/
func (this *QMediaPlayer) MediaStatusChanged(status int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QMediaPlayer18mediaStatusChangedENS_11MediaStatusE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), status)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaplayer.h:189
// index:0
// Public Visibility=Default Availability=Available
// [-2] void durationChanged(qint64)

/*
Signal the duration of the content has changed to duration, expressed in milliseconds.

Note: Notifier signal for property duration.
*/
func (this *QMediaPlayer) DurationChanged(duration int64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QMediaPlayer15durationChangedEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), duration)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaplayer.h:190
// index:0
// Public Visibility=Default Availability=Available
// [-2] void positionChanged(qint64)

/*
Signal the position of the content has changed to position, expressed in milliseconds.

Note: Notifier signal for property position.
*/
func (this *QMediaPlayer) PositionChanged(position int64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QMediaPlayer15positionChangedEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), position)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaplayer.h:192
// index:0
// Public Visibility=Default Availability=Available
// [-2] void volumeChanged(int)

/*
Signal the playback volume has changed to volume.

Note: Notifier signal for property volume.
*/
func (this *QMediaPlayer) VolumeChanged(volume int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QMediaPlayer13volumeChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), volume)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaplayer.h:193
// index:0
// Public Visibility=Default Availability=Available
// [-2] void mutedChanged(bool)

/*
Signal the mute state has changed to muted.

Note: Notifier signal for property muted.
*/
func (this *QMediaPlayer) MutedChanged(muted bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QMediaPlayer12mutedChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), muted)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaplayer.h:194
// index:0
// Public Visibility=Default Availability=Available
// [-2] void audioAvailableChanged(bool)

/*
Signals the availability of audio content has changed to available.

Note: Notifier signal for property audioAvailable.
*/
func (this *QMediaPlayer) AudioAvailableChanged(available bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QMediaPlayer21audioAvailableChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), available)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaplayer.h:195
// index:0
// Public Visibility=Default Availability=Available
// [-2] void videoAvailableChanged(bool)

/*
Signal the availability of visual content has changed to videoAvailable.

Note: Notifier signal for property videoAvailable.
*/
func (this *QMediaPlayer) VideoAvailableChanged(videoAvailable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QMediaPlayer21videoAvailableChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), videoAvailable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaplayer.h:197
// index:0
// Public Visibility=Default Availability=Available
// [-2] void bufferStatusChanged(int)

/*
Signal the amount of the local buffer filled as a percentage by percentFilled.

Note: Notifier signal for property bufferStatus.
*/
func (this *QMediaPlayer) BufferStatusChanged(percentFilled int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QMediaPlayer19bufferStatusChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), percentFilled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaplayer.h:199
// index:0
// Public Visibility=Default Availability=Available
// [-2] void seekableChanged(bool)

/*
Signals the seekable status of the player object has changed.

Note: Notifier signal for property seekable.
*/
func (this *QMediaPlayer) SeekableChanged(seekable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QMediaPlayer15seekableChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), seekable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaplayer.h:200
// index:0
// Public Visibility=Default Availability=Available
// [-2] void playbackRateChanged(qreal)

/*
Signals the playbackRate has changed to rate.

Note: Notifier signal for property playbackRate.
*/
func (this *QMediaPlayer) PlaybackRateChanged(rate float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QMediaPlayer19playbackRateChangedEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), rate)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaplayer.h:202
// index:0
// Public Visibility=Default Availability=Available
// [-2] void audioRoleChanged(QAudio::Role)

/*
Signals that the audio role of the media player has changed.

This function was introduced in  Qt 5.6.
*/
func (this *QMediaPlayer) AudioRoleChanged(role int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QMediaPlayer16audioRoleChangedEN6QAudio4RoleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), role)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaplayer.h:206
// index:0
// Public Visibility=Default Availability=Available
// [-2] void networkConfigurationChanged(const QNetworkConfiguration &)

/*
Signal that the active in use network access point has been changed to configuration and all subsequent network access will use this configuration.
*/
func (this *QMediaPlayer) NetworkConfigurationChanged(configuration qtnetwork.QNetworkConfiguration_ITF) {
	var convArg0 unsafe.Pointer
	if configuration != nil && configuration.QNetworkConfiguration_PTR() != nil {
		convArg0 = configuration.QNetworkConfiguration_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QMediaPlayer27networkConfigurationChangedERK21QNetworkConfiguration", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaplayer.h:208
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool bind(QObject *)

/*

 */
func (this *QMediaPlayer) Bind(arg0 qtcore.QObject_ITF /*777 QObject **/) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QObject_PTR() != nil {
		convArg0 = arg0.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QMediaPlayer4bindEP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qmediaplayer.h:209
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void unbind(QObject *)

/*

 */
func (this *QMediaPlayer) Unbind(arg0 qtcore.QObject_ITF /*777 QObject **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QObject_PTR() != nil {
		convArg0 = arg0.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QMediaPlayer6unbindEP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

/*
Defines the current state of a media player.


*/
type QMediaPlayer__State = int

// The media player is not playing content, playback will begin from the start of the current track.
const QMediaPlayer__StoppedState QMediaPlayer__State = 0

// The media player is currently playing content.
const QMediaPlayer__PlayingState QMediaPlayer__State = 1

// The media player has paused playback, playback of the current track will resume from the position the player was paused at.
const QMediaPlayer__PausedState QMediaPlayer__State = 2

func (this *QMediaPlayer) StateItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QMediaPlayer_StateItemName(val int) string {
	var nilthis *QMediaPlayer
	return nilthis.StateItemName(val)
}

/*
Defines the status of a media player's current media.


*/
type QMediaPlayer__MediaStatus = int

// The status of the media cannot be determined.
const QMediaPlayer__UnknownMediaStatus QMediaPlayer__MediaStatus = 0

// The is no current media. The player is in the StoppedState.
const QMediaPlayer__NoMedia QMediaPlayer__MediaStatus = 1

// The current media is being loaded. The player may be in any state.
const QMediaPlayer__LoadingMedia QMediaPlayer__MediaStatus = 2

// The current media has been loaded. The player is in the StoppedState.
const QMediaPlayer__LoadedMedia QMediaPlayer__MediaStatus = 3

// Playback of the current media has stalled due to insufficient buffering or some other temporary interruption. The player is in the PlayingState or PausedState.
const QMediaPlayer__StalledMedia QMediaPlayer__MediaStatus = 4

// The player is buffering data but has enough data buffered for playback to continue for the immediate future. The player is in the PlayingState or PausedState.
const QMediaPlayer__BufferingMedia QMediaPlayer__MediaStatus = 5

// The player has fully buffered the current media. The player is in the PlayingState or PausedState.
const QMediaPlayer__BufferedMedia QMediaPlayer__MediaStatus = 6

// Playback has reached the end of the current media. The player is in the StoppedState.
const QMediaPlayer__EndOfMedia QMediaPlayer__MediaStatus = 7

// The current media cannot be played. The player is in the StoppedState.
const QMediaPlayer__InvalidMedia QMediaPlayer__MediaStatus = 8

func (this *QMediaPlayer) MediaStatusItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QMediaPlayer_MediaStatusItemName(val int) string {
	var nilthis *QMediaPlayer
	return nilthis.MediaStatusItemName(val)
}

/*


 */
type QMediaPlayer__Flag = int

//
const QMediaPlayer__LowLatency QMediaPlayer__Flag = 1

//
const QMediaPlayer__StreamPlayback QMediaPlayer__Flag = 2

//
const QMediaPlayer__VideoSurface QMediaPlayer__Flag = 4

func (this *QMediaPlayer) FlagItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QMediaPlayer_FlagItemName(val int) string {
	var nilthis *QMediaPlayer
	return nilthis.FlagItemName(val)
}

/*
Defines a media player error condition.


*/
type QMediaPlayer__Error = int

// No error has occurred.
const QMediaPlayer__NoError QMediaPlayer__Error = 0

// A media resource couldn't be resolved.
const QMediaPlayer__ResourceError QMediaPlayer__Error = 1

// The format of a media resource isn't (fully) supported. Playback may still be possible, but without an audio or video component.
const QMediaPlayer__FormatError QMediaPlayer__Error = 2

// A network error occurred.
const QMediaPlayer__NetworkError QMediaPlayer__Error = 3

// There are not the appropriate permissions to play a media resource.
const QMediaPlayer__AccessDeniedError QMediaPlayer__Error = 4

// A valid playback service was not found, playback cannot proceed.
const QMediaPlayer__ServiceMissingError QMediaPlayer__Error = 5

//
const QMediaPlayer__MediaIsPlaylist QMediaPlayer__Error = 6

func (this *QMediaPlayer) ErrorItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QMediaPlayer_ErrorItemName(val int) string {
	var nilthis *QMediaPlayer
	return nilthis.ErrorItemName(val)
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
