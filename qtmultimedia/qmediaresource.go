package qtmultimedia

// /usr/include/qt/QtMultimedia/qmediaresource.h
// #include <qmediaresource.h>
// #include <QtMultimedia>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
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
type QMediaResource struct {
	*qtrt.CObject
}
type QMediaResource_ITF interface {
	QMediaResource_PTR() *QMediaResource
}

func (ptr *QMediaResource) QMediaResource_PTR() *QMediaResource { return ptr }

func (this *QMediaResource) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QMediaResource) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQMediaResourceFromPointer(cthis unsafe.Pointer) *QMediaResource {
	return &QMediaResource{&qtrt.CObject{cthis}}
}
func (*QMediaResource) NewFromPointer(cthis unsafe.Pointer) *QMediaResource {
	return NewQMediaResourceFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qmediaresource.h:56
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QMediaResource()

/*
Constructs a null media resource.
*/
func (*QMediaResource) NewForInherit() *QMediaResource {
	return NewQMediaResource()
}
func NewQMediaResource() *QMediaResource {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaResourceC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMediaResourceFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQMediaResource)
	return gothis
}

// /usr/include/qt/QtMultimedia/qmediaresource.h:57
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QMediaResource(const QUrl &, const QString &)

/*
Constructs a null media resource.
*/
func (*QMediaResource) NewForInherit1(url qtcore.QUrl_ITF, mimeType string) *QMediaResource {
	return NewQMediaResource1(url, mimeType)
}
func NewQMediaResource1(url qtcore.QUrl_ITF, mimeType string) *QMediaResource {
	var convArg0 unsafe.Pointer
	if url != nil && url.QUrl_PTR() != nil {
		convArg0 = url.QUrl_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString5(mimeType)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaResourceC2ERK4QUrlRK7QString", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMediaResourceFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQMediaResource)
	return gothis
}

// /usr/include/qt/QtMultimedia/qmediaresource.h:57
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QMediaResource(const QUrl &, const QString &)

/*
Constructs a null media resource.
*/
func (*QMediaResource) NewForInherit1p(url qtcore.QUrl_ITF) *QMediaResource {
	return NewQMediaResource1p(url)
}
func NewQMediaResource1p(url qtcore.QUrl_ITF) *QMediaResource {
	var convArg0 unsafe.Pointer
	if url != nil && url.QUrl_PTR() != nil {
		convArg0 = url.QUrl_PTR().GetCthis()
	}
	// arg: 1, const QString &=LValueReference, QString=Record, , Invalid
	var convArg1 = qtcore.NewQString()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaResourceC2ERK4QUrlRK7QString", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMediaResourceFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQMediaResource)
	return gothis
}

// /usr/include/qt/QtMultimedia/qmediaresource.h:58
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QMediaResource(const QNetworkRequest &, const QString &)

/*
Constructs a null media resource.
*/
func (*QMediaResource) NewForInherit2(request qtnetwork.QNetworkRequest_ITF, mimeType string) *QMediaResource {
	return NewQMediaResource2(request, mimeType)
}
func NewQMediaResource2(request qtnetwork.QNetworkRequest_ITF, mimeType string) *QMediaResource {
	var convArg0 unsafe.Pointer
	if request != nil && request.QNetworkRequest_PTR() != nil {
		convArg0 = request.QNetworkRequest_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString5(mimeType)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaResourceC2ERK15QNetworkRequestRK7QString", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMediaResourceFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQMediaResource)
	return gothis
}

// /usr/include/qt/QtMultimedia/qmediaresource.h:58
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QMediaResource(const QNetworkRequest &, const QString &)

/*
Constructs a null media resource.
*/
func (*QMediaResource) NewForInherit2p(request qtnetwork.QNetworkRequest_ITF) *QMediaResource {
	return NewQMediaResource2p(request)
}
func NewQMediaResource2p(request qtnetwork.QNetworkRequest_ITF) *QMediaResource {
	var convArg0 unsafe.Pointer
	if request != nil && request.QNetworkRequest_PTR() != nil {
		convArg0 = request.QNetworkRequest_PTR().GetCthis()
	}
	// arg: 1, const QString &=LValueReference, QString=Record, , Invalid
	var convArg1 = qtcore.NewQString()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaResourceC2ERK15QNetworkRequestRK7QString", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMediaResourceFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQMediaResource)
	return gothis
}

// /usr/include/qt/QtMultimedia/qmediaresource.h:60
// index:0
// Public Visibility=Default Availability=Available
// [8] QMediaResource & operator=(const QMediaResource &)

/*

 */
func (this *QMediaResource) Operator_equal(other QMediaResource_ITF) *QMediaResource {
	var convArg0 unsafe.Pointer
	if other != nil && other.QMediaResource_PTR() != nil {
		convArg0 = other.QMediaResource_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaResourceaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMediaResourceFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMediaResource)
	return rv2
}

// /usr/include/qt/QtMultimedia/qmediaresource.h:61
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QMediaResource()

/*

 */
func DeleteQMediaResource(this *QMediaResource) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaResourceD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qmediaresource.h:63
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isNull() const

/*
Identifies if a media resource is null.

Returns true if the resource is null, and false otherwise.
*/
func (this *QMediaResource) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QMediaResource6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qmediaresource.h:65
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QMediaResource &) const

/*

 */
func (this *QMediaResource) Operator_equal_equal(other QMediaResource_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QMediaResource_PTR() != nil {
		convArg0 = other.QMediaResource_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QMediaResourceeqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qmediaresource.h:66
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator!=(const QMediaResource &) const

/*

 */
func (this *QMediaResource) Operator_not_equal(other QMediaResource_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QMediaResource_PTR() != nil {
		convArg0 = other.QMediaResource_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QMediaResourceneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qmediaresource.h:68
// index:0
// Public Visibility=Default Availability=Available
// [8] QUrl url() const

/*
Returns the URL of a media resource.
*/
func (this *QMediaResource) Url() *qtcore.QUrl /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QMediaResource3urlEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtMultimedia/qmediaresource.h:69
// index:0
// Public Visibility=Default Availability=Available
// [8] QNetworkRequest request() const

/*
Returns the network request associated with this media resource.
*/
func (this *QMediaResource) Request() *qtnetwork.QNetworkRequest /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QMediaResource7requestEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtnetwork.NewQNetworkRequestFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtnetwork.DeleteQNetworkRequest)
	return rv2
}

// /usr/include/qt/QtMultimedia/qmediaresource.h:70
// index:0
// Public Visibility=Default Availability=Available
// [8] QString mimeType() const

/*
Returns the MIME type of a media resource.

This may be null if the MIME type is unknown.
*/
func (this *QMediaResource) MimeType() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QMediaResource8mimeTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtMultimedia/qmediaresource.h:72
// index:0
// Public Visibility=Default Availability=Available
// [8] QString language() const

/*
Returns the language of a media resource as an ISO 639-2 code.

This may be null if the language is unknown.

See also setLanguage().
*/
func (this *QMediaResource) Language() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QMediaResource8languageEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtMultimedia/qmediaresource.h:73
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLanguage(const QString &)

/*
Sets the language of a media resource.

See also language().
*/
func (this *QMediaResource) SetLanguage(language string) {
	var tmpArg0 = qtcore.NewQString5(language)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaResource11setLanguageERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaresource.h:75
// index:0
// Public Visibility=Default Availability=Available
// [8] QString audioCodec() const

/*
Returns the audio codec of a media resource.

This may be null if the media resource does not contain an audio stream, or the codec is unknown.

See also setAudioCodec().
*/
func (this *QMediaResource) AudioCodec() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QMediaResource10audioCodecEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtMultimedia/qmediaresource.h:76
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAudioCodec(const QString &)

/*
Sets the audio codec of a media resource.

See also audioCodec().
*/
func (this *QMediaResource) SetAudioCodec(codec string) {
	var tmpArg0 = qtcore.NewQString5(codec)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaResource13setAudioCodecERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaresource.h:78
// index:0
// Public Visibility=Default Availability=Available
// [8] QString videoCodec() const

/*
Returns the video codec of a media resource.

This may be null if the media resource does not contain a video stream, or the codec is unknonwn.

See also setVideoCodec().
*/
func (this *QMediaResource) VideoCodec() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QMediaResource10videoCodecEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtMultimedia/qmediaresource.h:79
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setVideoCodec(const QString &)

/*
Sets the video codec of media resource.

See also videoCodec().
*/
func (this *QMediaResource) SetVideoCodec(codec string) {
	var tmpArg0 = qtcore.NewQString5(codec)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaResource13setVideoCodecERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaresource.h:81
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 dataSize() const

/*
Returns the size in bytes of a media resource.

This may be zero if the size is unknown.

See also setDataSize().
*/
func (this *QMediaResource) DataSize() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QMediaResource8dataSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtMultimedia/qmediaresource.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDataSize(const qint64)

/*
Sets the size in bytes of a media resource.

See also dataSize().
*/
func (this *QMediaResource) SetDataSize(size int64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaResource11setDataSizeEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), size)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaresource.h:84
// index:0
// Public Visibility=Default Availability=Available
// [4] int audioBitRate() const

/*
Returns the bit rate in bits per second of a media resource's audio stream.

This may be zero if the bit rate is unknown, or the resource contains no audio stream.

See also setAudioBitRate().
*/
func (this *QMediaResource) AudioBitRate() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QMediaResource12audioBitRateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qmediaresource.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAudioBitRate(int)

/*
Sets the bit rate in bits per second of a media resource's video stream.

See also audioBitRate().
*/
func (this *QMediaResource) SetAudioBitRate(rate int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaResource15setAudioBitRateEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), rate)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaresource.h:87
// index:0
// Public Visibility=Default Availability=Available
// [4] int sampleRate() const

/*
Returns the audio sample rate of a media resource.

This may be zero if the sample size is unknown, or the resource contains no audio stream.

See also setSampleRate().
*/
func (this *QMediaResource) SampleRate() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QMediaResource10sampleRateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qmediaresource.h:88
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSampleRate(int)

/*
Sets the audio sampleRate of a media resource.

See also sampleRate().
*/
func (this *QMediaResource) SetSampleRate(frequency int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaResource13setSampleRateEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), frequency)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaresource.h:90
// index:0
// Public Visibility=Default Availability=Available
// [4] int channelCount() const

/*
Returns the number of audio channels in a media resource.

This may be zero if the sample size is unknown, or the resource contains no audio stream.

See also setChannelCount().
*/
func (this *QMediaResource) ChannelCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QMediaResource12channelCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qmediaresource.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setChannelCount(int)

/*
Sets the number of audio channels in a media resource.

See also channelCount().
*/
func (this *QMediaResource) SetChannelCount(channels int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaResource15setChannelCountEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), channels)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaresource.h:93
// index:0
// Public Visibility=Default Availability=Available
// [4] int videoBitRate() const

/*
Returns the bit rate in bits per second of a media resource's video stream.

This may be zero if the bit rate is unknown, or the resource contains no video stream.

See also setVideoBitRate().
*/
func (this *QMediaResource) VideoBitRate() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QMediaResource12videoBitRateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qmediaresource.h:94
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setVideoBitRate(int)

/*
Sets the bit rate in bits per second of a media resource's video stream.

See also videoBitRate().
*/
func (this *QMediaResource) SetVideoBitRate(rate int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaResource15setVideoBitRateEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), rate)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaresource.h:96
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize resolution() const

/*
Returns the resolution in pixels of a media resource.

This may be null is the resolution is unknown, or the resource contains no pixel data (i.e. the resource is an audio stream.

See also setResolution().
*/
func (this *QMediaResource) Resolution() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QMediaResource10resolutionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtMultimedia/qmediaresource.h:97
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setResolution(const QSize &)

/*
Sets the resolution in pixels of a media resource.

See also resolution().
*/
func (this *QMediaResource) SetResolution(resolution qtcore.QSize_ITF) {
	var convArg0 unsafe.Pointer
	if resolution != nil && resolution.QSize_PTR() != nil {
		convArg0 = resolution.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaResource13setResolutionERK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaresource.h:98
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setResolution(int, int)

/*
Sets the resolution in pixels of a media resource.

See also resolution().
*/
func (this *QMediaResource) SetResolution1(width int, height int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaResource13setResolutionEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), width, height)
	qtrt.ErrPrint(err, rv)
}

/*


 */
type QMediaResource__Property = int

//
const QMediaResource__Url QMediaResource__Property = 0

//
const QMediaResource__Request QMediaResource__Property = 1

//
const QMediaResource__MimeType QMediaResource__Property = 2

//
const QMediaResource__Language QMediaResource__Property = 3

//
const QMediaResource__AudioCodec QMediaResource__Property = 4

//
const QMediaResource__VideoCodec QMediaResource__Property = 5

//
const QMediaResource__DataSize QMediaResource__Property = 6

//
const QMediaResource__AudioBitRate QMediaResource__Property = 7

//
const QMediaResource__VideoBitRate QMediaResource__Property = 8

//
const QMediaResource__SampleRate QMediaResource__Property = 9

//
const QMediaResource__ChannelCount QMediaResource__Property = 10

//
const QMediaResource__Resolution QMediaResource__Property = 11

func (this *QMediaResource) PropertyItemName(val int) string {
	switch val {
	case QMediaResource__Url: // 0
		return "Url"
	case QMediaResource__Request: // 1
		return "Request"
	case QMediaResource__MimeType: // 2
		return "MimeType"
	case QMediaResource__Language: // 3
		return "Language"
	case QMediaResource__AudioCodec: // 4
		return "AudioCodec"
	case QMediaResource__VideoCodec: // 5
		return "VideoCodec"
	case QMediaResource__DataSize: // 6
		return "DataSize"
	case QMediaResource__AudioBitRate: // 7
		return "AudioBitRate"
	case QMediaResource__VideoBitRate: // 8
		return "VideoBitRate"
	case QMediaResource__SampleRate: // 9
		return "SampleRate"
	case QMediaResource__ChannelCount: // 10
		return "ChannelCount"
	case QMediaResource__Resolution: // 11
		return "Resolution"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QMediaResource_PropertyItemName(val int) string {
	var nilthis *QMediaResource
	return nilthis.PropertyItemName(val)
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
