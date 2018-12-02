package qtmultimedia

// /usr/include/qt/QtMultimedia/qmediarecorder.h
// #include <qmediarecorder.h>
// #include <QtMultimedia>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 3
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

// bool setMediaObject(QMediaObject *)
func (this *QMediaRecorder) InheritSetMediaObject(f func(object *QMediaObject /*777 QMediaObject **/) bool) {
	qtrt.SetAllInheritCallback(this, "setMediaObject", f)
}

/*

 */
type QMediaRecorder struct {
	*qtcore.QObject
	*QMediaBindableInterface
}
type QMediaRecorder_ITF interface {
	qtcore.QObject_ITF
	QMediaBindableInterface_ITF
	QMediaRecorder_PTR() *QMediaRecorder
}

func (ptr *QMediaRecorder) QMediaRecorder_PTR() *QMediaRecorder { return ptr }

func (this *QMediaRecorder) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QMediaRecorder) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
	this.QMediaBindableInterface = NewQMediaBindableInterfaceFromPointer(cthis)
}
func NewQMediaRecorderFromPointer(cthis unsafe.Pointer) *QMediaRecorder {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	bcthis1 := NewQMediaBindableInterfaceFromPointer(cthis)
	return &QMediaRecorder{bcthis0, bcthis1}
}
func (*QMediaRecorder) NewFromPointer(cthis unsafe.Pointer) *QMediaRecorder {
	return NewQMediaRecorderFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qmediarecorder.h:67
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QMediaRecorder) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QMediaRecorder10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qmediarecorder.h:109
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QMediaRecorder(QMediaObject *, QObject *)

/*
Constructs a media recorder which records the media produced by mediaObject.

The parent is passed to QMediaObject.
*/
func (*QMediaRecorder) NewForInherit(mediaObject QMediaObject_ITF /*777 QMediaObject **/, parent qtcore.QObject_ITF /*777 QObject **/) *QMediaRecorder {
	return NewQMediaRecorder(mediaObject, parent)
}
func NewQMediaRecorder(mediaObject QMediaObject_ITF /*777 QMediaObject **/, parent qtcore.QObject_ITF /*777 QObject **/) *QMediaRecorder {
	var convArg0 unsafe.Pointer
	if mediaObject != nil && mediaObject.QMediaObject_PTR() != nil {
		convArg0 = mediaObject.QMediaObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg1 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaRecorderC2EP12QMediaObjectP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMediaRecorderFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QMediaRecorder")
	return gothis
}

// /usr/include/qt/QtMultimedia/qmediarecorder.h:109
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QMediaRecorder(QMediaObject *, QObject *)

/*
Constructs a media recorder which records the media produced by mediaObject.

The parent is passed to QMediaObject.
*/
func (*QMediaRecorder) NewForInherit__(mediaObject QMediaObject_ITF /*777 QMediaObject **/) *QMediaRecorder {
	return NewQMediaRecorder__(mediaObject)
}
func NewQMediaRecorder__(mediaObject QMediaObject_ITF /*777 QMediaObject **/) *QMediaRecorder {
	var convArg0 unsafe.Pointer
	if mediaObject != nil && mediaObject.QMediaObject_PTR() != nil {
		convArg0 = mediaObject.QMediaObject_PTR().GetCthis()
	}
	// arg: 1, QObject *=Pointer, QObject=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaRecorderC2EP12QMediaObjectP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMediaRecorderFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QMediaRecorder")
	return gothis
}

// /usr/include/qt/QtMultimedia/qmediarecorder.h:110
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QMediaRecorder()

/*

 */
func DeleteQMediaRecorder(this *QMediaRecorder) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaRecorderD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 32)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qmediarecorder.h:112
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QMediaObject * mediaObject() const

/*
Reimplemented from QMediaBindableInterface::mediaObject().

Returns the QMediaObject instance that this QMediaRecorder is bound too, or 0 otherwise.
*/
func (this *QMediaRecorder) MediaObject() *QMediaObject /*777 QMediaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QMediaRecorder11mediaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMediaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qmediarecorder.h:114
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isAvailable() const

/*
Returns true if media recorder service ready to use.

See also availabilityChanged().
*/
func (this *QMediaRecorder) IsAvailable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QMediaRecorder11isAvailableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qmediarecorder.h:115
// index:0
// Public Visibility=Default Availability=Available
// [4] QMultimedia::AvailabilityStatus availability() const

/*
Returns the availability of this functionality.

See also availabilityChanged().
*/
func (this *QMediaRecorder) Availability() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QMediaRecorder12availabilityEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qmediarecorder.h:117
// index:0
// Public Visibility=Default Availability=Available
// [8] QUrl outputLocation() const

/*

 */
func (this *QMediaRecorder) OutputLocation() *qtcore.QUrl /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QMediaRecorder14outputLocationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtMultimedia/qmediarecorder.h:118
// index:0
// Public Visibility=Default Availability=Available
// [1] bool setOutputLocation(const QUrl &)

/*

 */
func (this *QMediaRecorder) SetOutputLocation(location qtcore.QUrl_ITF) bool {
	var convArg0 unsafe.Pointer
	if location != nil && location.QUrl_PTR() != nil {
		convArg0 = location.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaRecorder17setOutputLocationERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qmediarecorder.h:120
// index:0
// Public Visibility=Default Availability=Available
// [8] QUrl actualLocation() const

/*

 */
func (this *QMediaRecorder) ActualLocation() *qtcore.QUrl /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QMediaRecorder14actualLocationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtMultimedia/qmediarecorder.h:122
// index:0
// Public Visibility=Default Availability=Available
// [4] QMediaRecorder::State state() const

/*
Returns the current media recorder state.

Note: Getter function for property state.

See also QMediaRecorder::State.
*/
func (this *QMediaRecorder) State() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QMediaRecorder5stateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qmediarecorder.h:123
// index:0
// Public Visibility=Default Availability=Available
// [4] QMediaRecorder::Status status() const

/*
Returns the current media recorder status.

Note: Getter function for property status.

See also QMediaRecorder::Status.
*/
func (this *QMediaRecorder) Status() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QMediaRecorder6statusEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qmediarecorder.h:125
// index:0
// Public Visibility=Default Availability=Available
// [4] QMediaRecorder::Error error() const

/*
Returns the current error state.

See also errorString().
*/
func (this *QMediaRecorder) Error() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QMediaRecorder5errorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qmediarecorder.h:185
// index:1
// Public Visibility=Default Availability=Available
// [-2] void error(QMediaRecorder::Error)

/*
Returns the current error state.

See also errorString().
*/
func (this *QMediaRecorder) Error_1(error int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaRecorder5errorENS_5ErrorE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), error)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediarecorder.h:126
// index:0
// Public Visibility=Default Availability=Available
// [8] QString errorString() const

/*
Returns a string describing the current error state.

See also error().
*/
func (this *QMediaRecorder) ErrorString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QMediaRecorder11errorStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtMultimedia/qmediarecorder.h:128
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 duration() const

/*

 */
func (this *QMediaRecorder) Duration() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QMediaRecorder8durationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtMultimedia/qmediarecorder.h:130
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isMuted() const

/*

 */
func (this *QMediaRecorder) IsMuted() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QMediaRecorder7isMutedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qmediarecorder.h:131
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal volume() const

/*

 */
func (this *QMediaRecorder) Volume() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QMediaRecorder6volumeEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtMultimedia/qmediarecorder.h:133
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList supportedContainers() const

/*
Returns a list of supported container formats.
*/
func (this *QMediaRecorder) SupportedContainers() *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QMediaRecorder19supportedContainersEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtMultimedia/qmediarecorder.h:134
// index:0
// Public Visibility=Default Availability=Available
// [8] QString containerDescription(const QString &) const

/*
Returns a description of a container format.
*/
func (this *QMediaRecorder) ContainerDescription(format string) string {
	var tmpArg0 = qtcore.NewQString_5(format)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QMediaRecorder20containerDescriptionERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtMultimedia/qmediarecorder.h:136
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList supportedAudioCodecs() const

/*
Returns a list of supported audio codecs.
*/
func (this *QMediaRecorder) SupportedAudioCodecs() *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QMediaRecorder20supportedAudioCodecsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtMultimedia/qmediarecorder.h:137
// index:0
// Public Visibility=Default Availability=Available
// [8] QString audioCodecDescription(const QString &) const

/*
Returns a description of an audio codec.
*/
func (this *QMediaRecorder) AudioCodecDescription(codecName string) string {
	var tmpArg0 = qtcore.NewQString_5(codecName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QMediaRecorder21audioCodecDescriptionERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtMultimedia/qmediarecorder.h:142
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList supportedVideoCodecs() const

/*
Returns a list of supported video codecs.
*/
func (this *QMediaRecorder) SupportedVideoCodecs() *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QMediaRecorder20supportedVideoCodecsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtMultimedia/qmediarecorder.h:143
// index:0
// Public Visibility=Default Availability=Available
// [8] QString videoCodecDescription(const QString &) const

/*
Returns a description of a video codec.

See also setEncodingSettings().
*/
func (this *QMediaRecorder) VideoCodecDescription(codecName string) string {
	var tmpArg0 = qtcore.NewQString_5(codecName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QMediaRecorder21videoCodecDescriptionERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtMultimedia/qmediarecorder.h:145
// index:0
// Public Visibility=Default Availability=Available
// [8] QList<QSize> supportedResolutions(const QVideoEncoderSettings &, bool *) const

/*
Returns a list of resolutions video can be encoded at.

If non null video settings parameter is passed, the returned list is reduced to resolution supported with partial settings like video codec or framerate applied.

If the encoder supports arbitrary resolutions within the supported range, *continuous is set to true, otherwise *continuous is set to false.

See also QVideoEncoderSettings::resolution().
*/
func (this *QMediaRecorder) SupportedResolutions(settings QVideoEncoderSettings_ITF, continuous *bool) *qtcore.QSizeList /*lll*/ {
	var convArg0 unsafe.Pointer
	if settings != nil && settings.QVideoEncoderSettings_PTR() != nil {
		convArg0 = settings.QVideoEncoderSettings_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QMediaRecorder20supportedResolutionsERK21QVideoEncoderSettingsPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, continuous)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeListFromPointer(unsafe.Pointer(uintptr(rv))) //5552
	return rv2
}

// /usr/include/qt/QtMultimedia/qmediarecorder.h:145
// index:0
// Public Visibility=Default Availability=Available
// [8] QList<QSize> supportedResolutions(const QVideoEncoderSettings &, bool *) const

/*
Returns a list of resolutions video can be encoded at.

If non null video settings parameter is passed, the returned list is reduced to resolution supported with partial settings like video codec or framerate applied.

If the encoder supports arbitrary resolutions within the supported range, *continuous is set to true, otherwise *continuous is set to false.

See also QVideoEncoderSettings::resolution().
*/
func (this *QMediaRecorder) SupportedResolutions__() *qtcore.QSizeList /*lll*/ {
	// arg: 0, const QVideoEncoderSettings &=LValueReference, QVideoEncoderSettings=Record, , Invalid
	var convArg0 unsafe.Pointer
	// arg: 1, bool *=Pointer, =Invalid, , Invalid
	var continuous unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QMediaRecorder20supportedResolutionsERK21QVideoEncoderSettingsPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, continuous)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeListFromPointer(unsafe.Pointer(uintptr(rv))) //5552
	return rv2
}

// /usr/include/qt/QtMultimedia/qmediarecorder.h:145
// index:0
// Public Visibility=Default Availability=Available
// [8] QList<QSize> supportedResolutions(const QVideoEncoderSettings &, bool *) const

/*
Returns a list of resolutions video can be encoded at.

If non null video settings parameter is passed, the returned list is reduced to resolution supported with partial settings like video codec or framerate applied.

If the encoder supports arbitrary resolutions within the supported range, *continuous is set to true, otherwise *continuous is set to false.

See also QVideoEncoderSettings::resolution().
*/
func (this *QMediaRecorder) SupportedResolutions__1(settings QVideoEncoderSettings_ITF) *qtcore.QSizeList /*lll*/ {
	var convArg0 unsafe.Pointer
	if settings != nil && settings.QVideoEncoderSettings_PTR() != nil {
		convArg0 = settings.QVideoEncoderSettings_PTR().GetCthis()
	}
	// arg: 1, bool *=Pointer, =Invalid, , Invalid
	var continuous unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QMediaRecorder20supportedResolutionsERK21QVideoEncoderSettingsPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, continuous)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeListFromPointer(unsafe.Pointer(uintptr(rv))) //5552
	return rv2
}

// /usr/include/qt/QtMultimedia/qmediarecorder.h:151
// index:0
// Public Visibility=Default Availability=Available
// [8] QAudioEncoderSettings audioSettings() const

/*
Returns the audio encoder settings being used.

See also setAudioSettings() and setEncodingSettings().
*/
func (this *QMediaRecorder) AudioSettings() *QAudioEncoderSettings /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QMediaRecorder13audioSettingsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQAudioEncoderSettingsFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQAudioEncoderSettings)
	return rv2
}

// /usr/include/qt/QtMultimedia/qmediarecorder.h:152
// index:0
// Public Visibility=Default Availability=Available
// [8] QVideoEncoderSettings videoSettings() const

/*
Returns the video encoder settings being used.

See also setVideoSettings() and setEncodingSettings().
*/
func (this *QMediaRecorder) VideoSettings() *QVideoEncoderSettings /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QMediaRecorder13videoSettingsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVideoEncoderSettingsFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVideoEncoderSettings)
	return rv2
}

// /usr/include/qt/QtMultimedia/qmediarecorder.h:153
// index:0
// Public Visibility=Default Availability=Available
// [8] QString containerFormat() const

/*
Returns the selected container format.

See also setContainerFormat().
*/
func (this *QMediaRecorder) ContainerFormat() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QMediaRecorder15containerFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtMultimedia/qmediarecorder.h:155
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAudioSettings(const QAudioEncoderSettings &)

/*
Sets the audio encoder settings.

If some parameters are not specified, or null settings are passed, the encoder will choose default encoding parameters, depending on media source properties.

It's only possible to change settings when the encoder is in the QMediaEncoder::StoppedState state.

See also audioSettings(), videoSettings(), and containerFormat().
*/
func (this *QMediaRecorder) SetAudioSettings(audioSettings QAudioEncoderSettings_ITF) {
	var convArg0 unsafe.Pointer
	if audioSettings != nil && audioSettings.QAudioEncoderSettings_PTR() != nil {
		convArg0 = audioSettings.QAudioEncoderSettings_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaRecorder16setAudioSettingsERK21QAudioEncoderSettings", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediarecorder.h:156
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setVideoSettings(const QVideoEncoderSettings &)

/*
Sets the video encoder settings.

If some parameters are not specified, or null settings are passed, the encoder will choose default encoding parameters, depending on media source properties.

It's only possible to change settings when the encoder is in the QMediaEncoder::StoppedState state.

See also audioSettings(), videoSettings(), and containerFormat().
*/
func (this *QMediaRecorder) SetVideoSettings(videoSettings QVideoEncoderSettings_ITF) {
	var convArg0 unsafe.Pointer
	if videoSettings != nil && videoSettings.QVideoEncoderSettings_PTR() != nil {
		convArg0 = videoSettings.QVideoEncoderSettings_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaRecorder16setVideoSettingsERK21QVideoEncoderSettings", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediarecorder.h:157
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setContainerFormat(const QString &)

/*
Sets the media container format.

If the container format is not specified, the encoder will choose format, depending on media source properties and encoding settings selected.

It's only possible to change settings when the encoder is in the QMediaEncoder::StoppedState state.

See also audioSettings(), videoSettings(), and containerFormat().
*/
func (this *QMediaRecorder) SetContainerFormat(container string) {
	var tmpArg0 = qtcore.NewQString_5(container)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaRecorder18setContainerFormatERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediarecorder.h:159
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setEncodingSettings(const QAudioEncoderSettings &, const QVideoEncoderSettings &, const QString &)

/*
Sets the audio and video encoder settings and container format.

If some parameters are not specified, or null settings are passed, the encoder will choose default encoding parameters, depending on media source properties.

It's only possible to change settings when the encoder is in the QMediaEncoder::StoppedState state.

See also audioSettings(), videoSettings(), and containerFormat().
*/
func (this *QMediaRecorder) SetEncodingSettings(audioSettings QAudioEncoderSettings_ITF, videoSettings QVideoEncoderSettings_ITF, containerMimeType string) {
	var convArg0 unsafe.Pointer
	if audioSettings != nil && audioSettings.QAudioEncoderSettings_PTR() != nil {
		convArg0 = audioSettings.QAudioEncoderSettings_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if videoSettings != nil && videoSettings.QVideoEncoderSettings_PTR() != nil {
		convArg1 = videoSettings.QVideoEncoderSettings_PTR().GetCthis()
	}
	var tmpArg2 = qtcore.NewQString_5(containerMimeType)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaRecorder19setEncodingSettingsERK21QAudioEncoderSettingsRK21QVideoEncoderSettingsRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediarecorder.h:159
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setEncodingSettings(const QAudioEncoderSettings &, const QVideoEncoderSettings &, const QString &)

/*
Sets the audio and video encoder settings and container format.

If some parameters are not specified, or null settings are passed, the encoder will choose default encoding parameters, depending on media source properties.

It's only possible to change settings when the encoder is in the QMediaEncoder::StoppedState state.

See also audioSettings(), videoSettings(), and containerFormat().
*/
func (this *QMediaRecorder) SetEncodingSettings__(audioSettings QAudioEncoderSettings_ITF) {
	var convArg0 unsafe.Pointer
	if audioSettings != nil && audioSettings.QAudioEncoderSettings_PTR() != nil {
		convArg0 = audioSettings.QAudioEncoderSettings_PTR().GetCthis()
	}
	// arg: 1, const QVideoEncoderSettings &=LValueReference, QVideoEncoderSettings=Record, , Invalid
	var convArg1 unsafe.Pointer
	// arg: 2, const QString &=LValueReference, QString=Record, , Invalid
	var convArg2 = qtcore.NewQString()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaRecorder19setEncodingSettingsERK21QAudioEncoderSettingsRK21QVideoEncoderSettingsRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediarecorder.h:159
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setEncodingSettings(const QAudioEncoderSettings &, const QVideoEncoderSettings &, const QString &)

/*
Sets the audio and video encoder settings and container format.

If some parameters are not specified, or null settings are passed, the encoder will choose default encoding parameters, depending on media source properties.

It's only possible to change settings when the encoder is in the QMediaEncoder::StoppedState state.

See also audioSettings(), videoSettings(), and containerFormat().
*/
func (this *QMediaRecorder) SetEncodingSettings__1(audioSettings QAudioEncoderSettings_ITF, videoSettings QVideoEncoderSettings_ITF) {
	var convArg0 unsafe.Pointer
	if audioSettings != nil && audioSettings.QAudioEncoderSettings_PTR() != nil {
		convArg0 = audioSettings.QAudioEncoderSettings_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if videoSettings != nil && videoSettings.QVideoEncoderSettings_PTR() != nil {
		convArg1 = videoSettings.QVideoEncoderSettings_PTR().GetCthis()
	}
	// arg: 2, const QString &=LValueReference, QString=Record, , Invalid
	var convArg2 = qtcore.NewQString()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaRecorder19setEncodingSettingsERK21QAudioEncoderSettingsRK21QVideoEncoderSettingsRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediarecorder.h:163
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isMetaDataAvailable() const

/*

 */
func (this *QMediaRecorder) IsMetaDataAvailable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QMediaRecorder19isMetaDataAvailableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qmediarecorder.h:164
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isMetaDataWritable() const

/*

 */
func (this *QMediaRecorder) IsMetaDataWritable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QMediaRecorder18isMetaDataWritableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qmediarecorder.h:166
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant metaData(const QString &) const

/*
Returns the value associated with a meta-data key.

See also setMetaData().
*/
func (this *QMediaRecorder) MetaData(key string) *qtcore.QVariant /*123*/ {
	var tmpArg0 = qtcore.NewQString_5(key)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QMediaRecorder8metaDataERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtMultimedia/qmediarecorder.h:167
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMetaData(const QString &, const QVariant &)

/*
Sets a value for a meta-data key.

Note: To ensure that meta data is set corretly, it should be set before starting the recording. Once the recording is stopped, any meta data set will be attached to the next recording.

See also metaData().
*/
func (this *QMediaRecorder) SetMetaData(key string, value qtcore.QVariant_ITF) {
	var tmpArg0 = qtcore.NewQString_5(key)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg1 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaRecorder11setMetaDataERK7QStringRK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediarecorder.h:168
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList availableMetaData() const

/*
Returns a list of keys there is meta-data available for.
*/
func (this *QMediaRecorder) AvailableMetaData() *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QMediaRecorder17availableMetaDataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtMultimedia/qmediarecorder.h:171
// index:0
// Public Visibility=Default Availability=Available
// [-2] void record()

/*
Start recording.

While the recorder state is changed immediately to QMediaRecorder::RecordingState, recording may start asynchronously, with statusChanged(QMediaRecorder::RecordingStatus) signal emitted when recording starts.

If recording fails error() signal is emitted with recorder state being reset back to QMediaRecorder::StoppedState.
*/
func (this *QMediaRecorder) Record() {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaRecorder6recordEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediarecorder.h:172
// index:0
// Public Visibility=Default Availability=Available
// [-2] void pause()

/*
Pause recording.

The recorder state is changed to QMediaRecorder::PausedState.

Depending on platform recording pause may be not supported, in this case the recorder state stays unchanged.
*/
func (this *QMediaRecorder) Pause() {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaRecorder5pauseEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediarecorder.h:173
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stop()

/*
Stop recording.

The recorder state is changed to QMediaRecorder::StoppedState.
*/
func (this *QMediaRecorder) Stop() {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaRecorder4stopEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediarecorder.h:174
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMuted(bool)

/*

 */
func (this *QMediaRecorder) SetMuted(muted bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaRecorder8setMutedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), muted)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediarecorder.h:175
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setVolume(qreal)

/*

 */
func (this *QMediaRecorder) SetVolume(volume float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaRecorder9setVolumeEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), volume)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediarecorder.h:178
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stateChanged(QMediaRecorder::State)

/*
Signals that a media recorder's state has changed.

Note: Notifier signal for property state.
*/
func (this *QMediaRecorder) StateChanged(state int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaRecorder12stateChangedENS_5StateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), state)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediarecorder.h:179
// index:0
// Public Visibility=Default Availability=Available
// [-2] void statusChanged(QMediaRecorder::Status)

/*

 */
func (this *QMediaRecorder) StatusChanged(status int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaRecorder13statusChangedENS_6StatusE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), status)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediarecorder.h:180
// index:0
// Public Visibility=Default Availability=Available
// [-2] void durationChanged(qint64)

/*
Signals that the duration of the recorded media has changed.

Note: Notifier signal for property duration.
*/
func (this *QMediaRecorder) DurationChanged(duration int64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaRecorder15durationChangedEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), duration)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediarecorder.h:181
// index:0
// Public Visibility=Default Availability=Available
// [-2] void mutedChanged(bool)

/*
Signals that the muted state has changed. If true the recording is being muted.

Note: Notifier signal for property muted.
*/
func (this *QMediaRecorder) MutedChanged(muted bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaRecorder12mutedChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), muted)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediarecorder.h:182
// index:0
// Public Visibility=Default Availability=Available
// [-2] void volumeChanged(qreal)

/*

 */
func (this *QMediaRecorder) VolumeChanged(volume float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaRecorder13volumeChangedEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), volume)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediarecorder.h:183
// index:0
// Public Visibility=Default Availability=Available
// [-2] void actualLocationChanged(const QUrl &)

/*
Signals that the actual location of the recorded media has changed. This signal is usually emitted when recording starts.

Note: Notifier signal for property actualLocation.
*/
func (this *QMediaRecorder) ActualLocationChanged(location qtcore.QUrl_ITF) {
	var convArg0 unsafe.Pointer
	if location != nil && location.QUrl_PTR() != nil {
		convArg0 = location.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaRecorder21actualLocationChangedERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediarecorder.h:187
// index:0
// Public Visibility=Default Availability=Available
// [-2] void metaDataAvailableChanged(bool)

/*
Signals that the available state of a media object's meta-data has changed.

Note: Notifier signal for property metaDataAvailable.
*/
func (this *QMediaRecorder) MetaDataAvailableChanged(available bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaRecorder24metaDataAvailableChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), available)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediarecorder.h:188
// index:0
// Public Visibility=Default Availability=Available
// [-2] void metaDataWritableChanged(bool)

/*
Signals that the writable state of a media object's meta-data has changed.

Note: Notifier signal for property metaDataWritable.
*/
func (this *QMediaRecorder) MetaDataWritableChanged(writable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaRecorder23metaDataWritableChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), writable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediarecorder.h:189
// index:0
// Public Visibility=Default Availability=Available
// [-2] void metaDataChanged()

/*
Signals that a media object's meta-data has changed.

If multiple meta-data elements are changed, metaDataChanged(const QString &key, const QVariant &value) signal is emitted for each of them with metaDataChanged() changed emitted once.

Note: Signal metaDataChanged is overloaded in this class. To connect to this signal by using the function pointer syntax, Qt provides a convenient helper for obtaining the function pointer as shown in this example:


  connect(mediaRecorder, QOverload<>::of(&QMediaRecorder::metaDataChanged),
      [=](){ /-* ... *-/ });
*/
func (this *QMediaRecorder) MetaDataChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaRecorder15metaDataChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediarecorder.h:190
// index:1
// Public Visibility=Default Availability=Available
// [-2] void metaDataChanged(const QString &, const QVariant &)

/*
Signals that a media object's meta-data has changed.

If multiple meta-data elements are changed, metaDataChanged(const QString &key, const QVariant &value) signal is emitted for each of them with metaDataChanged() changed emitted once.

Note: Signal metaDataChanged is overloaded in this class. To connect to this signal by using the function pointer syntax, Qt provides a convenient helper for obtaining the function pointer as shown in this example:


  connect(mediaRecorder, QOverload<>::of(&QMediaRecorder::metaDataChanged),
      [=](){ /-* ... *-/ });
*/
func (this *QMediaRecorder) MetaDataChanged_1(key string, value qtcore.QVariant_ITF) {
	var tmpArg0 = qtcore.NewQString_5(key)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg1 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaRecorder15metaDataChangedERK7QStringRK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediarecorder.h:192
// index:0
// Public Visibility=Default Availability=Available
// [-2] void availabilityChanged(bool)

/*
Signals that the media recorder is now available (if available is true), or not.

Note: Signal availabilityChanged is overloaded in this class. To connect to this signal by using the function pointer syntax, Qt provides a convenient helper for obtaining the function pointer as shown in this example:


  connect(mediaRecorder, QOverload<bool>::of(&QMediaRecorder::availabilityChanged),
      [=](bool available){ /-* ... *-/ });
*/
func (this *QMediaRecorder) AvailabilityChanged(available bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaRecorder19availabilityChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), available)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediarecorder.h:193
// index:1
// Public Visibility=Default Availability=Available
// [-2] void availabilityChanged(QMultimedia::AvailabilityStatus)

/*
Signals that the media recorder is now available (if available is true), or not.

Note: Signal availabilityChanged is overloaded in this class. To connect to this signal by using the function pointer syntax, Qt provides a convenient helper for obtaining the function pointer as shown in this example:


  connect(mediaRecorder, QOverload<bool>::of(&QMediaRecorder::availabilityChanged),
      [=](bool available){ /-* ... *-/ });
*/
func (this *QMediaRecorder) AvailabilityChanged_1(availability int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaRecorder19availabilityChangedEN11QMultimedia18AvailabilityStatusE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), availability)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediarecorder.h:197
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool setMediaObject(QMediaObject *)

/*

 */
func (this *QMediaRecorder) SetMediaObject(object QMediaObject_ITF /*777 QMediaObject **/) bool {
	var convArg0 unsafe.Pointer
	if object != nil && object.QMediaObject_PTR() != nil {
		convArg0 = object.QMediaObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaRecorder14setMediaObjectEP12QMediaObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

/*

 */
type QMediaRecorder__State = int

// The recorder is not active. If this is the state after recording then the actual created recording has finished being written to the final location and is ready on all platforms except on Android. On Android, due to platform limitations, there is no way to be certain that the recording has finished writing to the final location.
const QMediaRecorder__StoppedState QMediaRecorder__State = 0

// The recording is requested.
const QMediaRecorder__RecordingState QMediaRecorder__State = 1

// The recorder is paused.
const QMediaRecorder__PausedState QMediaRecorder__State = 2

func (this *QMediaRecorder) StateItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QMediaRecorder_StateItemName(val int) string {
	var nilthis *QMediaRecorder
	return nilthis.StateItemName(val)
}

/*

 */
type QMediaRecorder__Status = int

// The recorder is not available or not supported by connected media object.
const QMediaRecorder__UnavailableStatus QMediaRecorder__Status = 0

// The recorder is avilable but not loaded.
const QMediaRecorder__UnloadedStatus QMediaRecorder__Status = 1

// The recorder is initializing.
const QMediaRecorder__LoadingStatus QMediaRecorder__Status = 2

// The recorder is initialized and ready to record media.
const QMediaRecorder__LoadedStatus QMediaRecorder__Status = 3

// Recording is requested but not active yet.
const QMediaRecorder__StartingStatus QMediaRecorder__Status = 4

// Recording is active.
const QMediaRecorder__RecordingStatus QMediaRecorder__Status = 5

// Recording is paused.
const QMediaRecorder__PausedStatus QMediaRecorder__Status = 6

// Recording is stopped with media being finalized.
const QMediaRecorder__FinalizingStatus QMediaRecorder__Status = 7

func (this *QMediaRecorder) StatusItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QMediaRecorder_StatusItemName(val int) string {
	var nilthis *QMediaRecorder
	return nilthis.StatusItemName(val)
}

/*

 */
type QMediaRecorder__Error = int

// No Errors.
const QMediaRecorder__NoError QMediaRecorder__Error = 0

// Device is not ready or not available.
const QMediaRecorder__ResourceError QMediaRecorder__Error = 1

// Current format is not supported.
const QMediaRecorder__FormatError QMediaRecorder__Error = 2

// No space left on device.
const QMediaRecorder__OutOfSpaceError QMediaRecorder__Error = 3

func (this *QMediaRecorder) ErrorItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QMediaRecorder_ErrorItemName(val int) string {
	var nilthis *QMediaRecorder
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
