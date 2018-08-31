package qtmultimedia

// /usr/include/qt/QtMultimedia/qcameraimagecapture.h
// #include <qcameraimagecapture.h>
// #include <QtMultimedia>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 42
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
func (this *QCameraImageCapture) InheritSetMediaObject(f func(arg0 *QMediaObject /*777 QMediaObject **/) bool) {
	qtrt.SetAllInheritCallback(this, "setMediaObject", f)
}

/*

 */
type QCameraImageCapture struct {
	*qtcore.QObject
	*QMediaBindableInterface
}
type QCameraImageCapture_ITF interface {
	qtcore.QObject_ITF
	QMediaBindableInterface_ITF
	QCameraImageCapture_PTR() *QCameraImageCapture
}

func (ptr *QCameraImageCapture) QCameraImageCapture_PTR() *QCameraImageCapture { return ptr }

func (this *QCameraImageCapture) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QCameraImageCapture) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
	this.QMediaBindableInterface = NewQMediaBindableInterfaceFromPointer(cthis)
}
func NewQCameraImageCaptureFromPointer(cthis unsafe.Pointer) *QCameraImageCapture {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	bcthis1 := NewQMediaBindableInterfaceFromPointer(cthis)
	return &QCameraImageCapture{bcthis0, bcthis1}
}
func (*QCameraImageCapture) NewFromPointer(cthis unsafe.Pointer) *QCameraImageCapture {
	return NewQCameraImageCaptureFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qcameraimagecapture.h:62
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QCameraImageCapture) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QCameraImageCapture10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qcameraimagecapture.h:90
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QCameraImageCapture(QMediaObject *, QObject *)

/*
Constructs a media recorder which records the media produced by mediaObject.

The parent is passed to QMediaObject.
*/
func (*QCameraImageCapture) NewForInherit(mediaObject QMediaObject_ITF /*777 QMediaObject **/, parent qtcore.QObject_ITF /*777 QObject **/) *QCameraImageCapture {
	return NewQCameraImageCapture(mediaObject, parent)
}
func NewQCameraImageCapture(mediaObject QMediaObject_ITF /*777 QMediaObject **/, parent qtcore.QObject_ITF /*777 QObject **/) *QCameraImageCapture {
	var convArg0 unsafe.Pointer
	if mediaObject != nil && mediaObject.QMediaObject_PTR() != nil {
		convArg0 = mediaObject.QMediaObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg1 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QCameraImageCaptureC2EP12QMediaObjectP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCameraImageCaptureFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QCameraImageCapture")
	return gothis
}

// /usr/include/qt/QtMultimedia/qcameraimagecapture.h:90
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QCameraImageCapture(QMediaObject *, QObject *)

/*
Constructs a media recorder which records the media produced by mediaObject.

The parent is passed to QMediaObject.
*/
func (*QCameraImageCapture) NewForInherit__(mediaObject QMediaObject_ITF /*777 QMediaObject **/) *QCameraImageCapture {
	return NewQCameraImageCapture__(mediaObject)
}
func NewQCameraImageCapture__(mediaObject QMediaObject_ITF /*777 QMediaObject **/) *QCameraImageCapture {
	var convArg0 unsafe.Pointer
	if mediaObject != nil && mediaObject.QMediaObject_PTR() != nil {
		convArg0 = mediaObject.QMediaObject_PTR().GetCthis()
	}
	// arg: 1, QObject *=Pointer, QObject=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN19QCameraImageCaptureC2EP12QMediaObjectP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCameraImageCaptureFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QCameraImageCapture")
	return gothis
}

// /usr/include/qt/QtMultimedia/qcameraimagecapture.h:91
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QCameraImageCapture()

/*

 */
func DeleteQCameraImageCapture(this *QCameraImageCapture) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QCameraImageCaptureD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 32)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qcameraimagecapture.h:93
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isAvailable() const

/*
Returns true if the images capture service ready to use.
*/
func (this *QCameraImageCapture) IsAvailable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QCameraImageCapture11isAvailableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qcameraimagecapture.h:94
// index:0
// Public Visibility=Default Availability=Available
// [4] QMultimedia::AvailabilityStatus availability() const

/*
Returns the availability of this functionality.
*/
func (this *QCameraImageCapture) Availability() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QCameraImageCapture12availabilityEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qcameraimagecapture.h:96
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QMediaObject * mediaObject() const

/*
Reimplemented from QMediaBindableInterface::mediaObject().

See also setMediaObject().
*/
func (this *QCameraImageCapture) MediaObject() *QMediaObject /*777 QMediaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QCameraImageCapture11mediaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMediaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qcameraimagecapture.h:98
// index:0
// Public Visibility=Default Availability=Available
// [4] QCameraImageCapture::Error error() const

/*
Returns the current error state.

See also errorString().
*/
func (this *QCameraImageCapture) Error() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QCameraImageCapture5errorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qcameraimagecapture.h:125
// index:1
// Public Visibility=Default Availability=Available
// [-2] void error(int, QCameraImageCapture::Error, const QString &)

/*
Returns the current error state.

See also errorString().
*/
func (this *QCameraImageCapture) Error_1(id int, error int, errorString string) {
	var tmpArg2 = qtcore.NewQString_5(errorString)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN19QCameraImageCapture5errorEiNS_5ErrorERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), id, error, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameraimagecapture.h:99
// index:0
// Public Visibility=Default Availability=Available
// [8] QString errorString() const

/*
Returns a string describing the current error state.

See also error().
*/
func (this *QCameraImageCapture) ErrorString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QCameraImageCapture11errorStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtMultimedia/qcameraimagecapture.h:101
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isReadyForCapture() const

/*

 */
func (this *QCameraImageCapture) IsReadyForCapture() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QCameraImageCapture17isReadyForCaptureEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qcameraimagecapture.h:103
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList supportedImageCodecs() const

/*
Returns a list of supported image codecs.
*/
func (this *QCameraImageCapture) SupportedImageCodecs() *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QCameraImageCapture20supportedImageCodecsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtMultimedia/qcameraimagecapture.h:104
// index:0
// Public Visibility=Default Availability=Available
// [8] QString imageCodecDescription(const QString &) const

/*
Returns a description of an image codec.
*/
func (this *QCameraImageCapture) ImageCodecDescription(codecName string) string {
	var tmpArg0 = qtcore.NewQString_5(codecName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QCameraImageCapture21imageCodecDescriptionERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtMultimedia/qcameraimagecapture.h:106
// index:0
// Public Visibility=Default Availability=Available
// [8] QList<QSize> supportedResolutions(const QImageEncoderSettings &, bool *) const

/*
Returns a list of resolutions images can be encoded at.

If non null image settings parameter is passed, the returned list is reduced to resolution supported with partial settings like image codec or quality applied.

If the encoder supports arbitrary resolutions within the supported range, *continuous is set to true, otherwise *continuous is set to false.

See also QImageEncoderSettings::resolution().
*/
func (this *QCameraImageCapture) SupportedResolutions(settings QImageEncoderSettings_ITF, continuous *bool) *qtcore.QSizeList /*lll*/ {
	var convArg0 unsafe.Pointer
	if settings != nil && settings.QImageEncoderSettings_PTR() != nil {
		convArg0 = settings.QImageEncoderSettings_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QCameraImageCapture20supportedResolutionsERK21QImageEncoderSettingsPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, continuous)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeListFromPointer(unsafe.Pointer(uintptr(rv))) //5552
	return rv2
}

// /usr/include/qt/QtMultimedia/qcameraimagecapture.h:106
// index:0
// Public Visibility=Default Availability=Available
// [8] QList<QSize> supportedResolutions(const QImageEncoderSettings &, bool *) const

/*
Returns a list of resolutions images can be encoded at.

If non null image settings parameter is passed, the returned list is reduced to resolution supported with partial settings like image codec or quality applied.

If the encoder supports arbitrary resolutions within the supported range, *continuous is set to true, otherwise *continuous is set to false.

See also QImageEncoderSettings::resolution().
*/
func (this *QCameraImageCapture) SupportedResolutions__() *qtcore.QSizeList /*lll*/ {
	// arg: 0, const QImageEncoderSettings &=LValueReference, QImageEncoderSettings=Record, , Invalid
	var convArg0 unsafe.Pointer
	// arg: 1, bool *=Pointer, =Invalid, , Invalid
	var continuous unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QCameraImageCapture20supportedResolutionsERK21QImageEncoderSettingsPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, continuous)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeListFromPointer(unsafe.Pointer(uintptr(rv))) //5552
	return rv2
}

// /usr/include/qt/QtMultimedia/qcameraimagecapture.h:106
// index:0
// Public Visibility=Default Availability=Available
// [8] QList<QSize> supportedResolutions(const QImageEncoderSettings &, bool *) const

/*
Returns a list of resolutions images can be encoded at.

If non null image settings parameter is passed, the returned list is reduced to resolution supported with partial settings like image codec or quality applied.

If the encoder supports arbitrary resolutions within the supported range, *continuous is set to true, otherwise *continuous is set to false.

See also QImageEncoderSettings::resolution().
*/
func (this *QCameraImageCapture) SupportedResolutions__1(settings QImageEncoderSettings_ITF) *qtcore.QSizeList /*lll*/ {
	var convArg0 unsafe.Pointer
	if settings != nil && settings.QImageEncoderSettings_PTR() != nil {
		convArg0 = settings.QImageEncoderSettings_PTR().GetCthis()
	}
	// arg: 1, bool *=Pointer, =Invalid, , Invalid
	var continuous unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QCameraImageCapture20supportedResolutionsERK21QImageEncoderSettingsPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, continuous)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeListFromPointer(unsafe.Pointer(uintptr(rv))) //5552
	return rv2
}

// /usr/include/qt/QtMultimedia/qcameraimagecapture.h:109
// index:0
// Public Visibility=Default Availability=Available
// [8] QImageEncoderSettings encodingSettings() const

/*
Returns the image encoder settings being used.

See also setEncodingSettings().
*/
func (this *QCameraImageCapture) EncodingSettings() *QImageEncoderSettings /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QCameraImageCapture16encodingSettingsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageEncoderSettingsFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImageEncoderSettings)
	return rv2
}

// /usr/include/qt/QtMultimedia/qcameraimagecapture.h:110
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setEncodingSettings(const QImageEncoderSettings &)

/*
Sets the image encoding settings.

If some parameters are not specified, or null settings are passed, the encoder choose the default encoding parameters.

See also encodingSettings().
*/
func (this *QCameraImageCapture) SetEncodingSettings(settings QImageEncoderSettings_ITF) {
	var convArg0 unsafe.Pointer
	if settings != nil && settings.QImageEncoderSettings_PTR() != nil {
		convArg0 = settings.QImageEncoderSettings_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QCameraImageCapture19setEncodingSettingsERK21QImageEncoderSettings", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameraimagecapture.h:113
// index:0
// Public Visibility=Default Availability=Available
// [4] QVideoFrame::PixelFormat bufferFormat() const

/*
Returns the buffer image capture format being used.

See also supportedBufferFormats() and setBufferFormat().
*/
func (this *QCameraImageCapture) BufferFormat() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QCameraImageCapture12bufferFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qcameraimagecapture.h:114
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBufferFormat(const QVideoFrame::PixelFormat)

/*
Sets the buffer image capture format to be used.

See also bufferFormat(), supportedBufferFormats(), and captureDestination().
*/
func (this *QCameraImageCapture) SetBufferFormat(format int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QCameraImageCapture15setBufferFormatEN11QVideoFrame11PixelFormatE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), format)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameraimagecapture.h:116
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isCaptureDestinationSupported(QCameraImageCapture::CaptureDestinations) const

/*
Returns true if the image capture destination is supported; otherwise returns false.

See also captureDestination() and setCaptureDestination().
*/
func (this *QCameraImageCapture) IsCaptureDestinationSupported(destination int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QCameraImageCapture29isCaptureDestinationSupportedE6QFlagsINS_18CaptureDestinationEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), destination)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qcameraimagecapture.h:117
// index:0
// Public Visibility=Default Availability=Available
// [4] QCameraImageCapture::CaptureDestinations captureDestination() const

/*
Returns the image capture destination being used.

See also isCaptureDestinationSupported() and setCaptureDestination().
*/
func (this *QCameraImageCapture) CaptureDestination() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QCameraImageCapture18captureDestinationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qcameraimagecapture.h:118
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCaptureDestination(QCameraImageCapture::CaptureDestinations)

/*
Sets the capture destination to be used.

See also isCaptureDestinationSupported() and captureDestination().
*/
func (this *QCameraImageCapture) SetCaptureDestination(destination int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QCameraImageCapture21setCaptureDestinationE6QFlagsINS_18CaptureDestinationEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), destination)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameraimagecapture.h:121
// index:0
// Public Visibility=Default Availability=Available
// [4] int capture(const QString &)

/*
Capture the image and save it to file. This operation is asynchronous in majority of cases, followed by signals QCameraImageCapture::imageExposed(), QCameraImageCapture::imageCaptured(), QCameraImageCapture::imageSaved() or QCameraImageCapture::error().

If an empty file is passed, the camera backend choses the default location and naming scheme for photos on the system, if only file name without full path is specified, the image will be saved to the default directory, with a full path reported with imageCaptured() and imageSaved() signals.

QCamera saves all the capture parameters like exposure settings or image processing parameters, so changes to camera parameters after capture() is called do not affect previous capture requests.

QCameraImageCapture::capture returns the capture Id parameter, used with imageExposed(), imageCaptured() and imageSaved() signals.

See also isReadyForCapture().
*/
func (this *QCameraImageCapture) Capture(location string) int {
	var tmpArg0 = qtcore.NewQString_5(location)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN19QCameraImageCapture7captureERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qcameraimagecapture.h:121
// index:0
// Public Visibility=Default Availability=Available
// [4] int capture(const QString &)

/*
Capture the image and save it to file. This operation is asynchronous in majority of cases, followed by signals QCameraImageCapture::imageExposed(), QCameraImageCapture::imageCaptured(), QCameraImageCapture::imageSaved() or QCameraImageCapture::error().

If an empty file is passed, the camera backend choses the default location and naming scheme for photos on the system, if only file name without full path is specified, the image will be saved to the default directory, with a full path reported with imageCaptured() and imageSaved() signals.

QCamera saves all the capture parameters like exposure settings or image processing parameters, so changes to camera parameters after capture() is called do not affect previous capture requests.

QCameraImageCapture::capture returns the capture Id parameter, used with imageExposed(), imageCaptured() and imageSaved() signals.

See also isReadyForCapture().
*/
func (this *QCameraImageCapture) Capture__() int {
	// arg: 0, const QString &=LValueReference, QString=Record, , Invalid
	var convArg0 = qtcore.NewQString()
	rv, err := qtrt.InvokeQtFunc6("_ZN19QCameraImageCapture7captureERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qcameraimagecapture.h:122
// index:0
// Public Visibility=Default Availability=Available
// [-2] void cancelCapture()

/*
Cancel incomplete capture requests. Already captured and queused for proicessing images may be discarded.
*/
func (this *QCameraImageCapture) CancelCapture() {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QCameraImageCapture13cancelCaptureEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameraimagecapture.h:127
// index:0
// Public Visibility=Default Availability=Available
// [-2] void readyForCaptureChanged(bool)

/*
Signals that a camera's ready for capture state has changed.

Note: Notifier signal for property readyForCapture.
*/
func (this *QCameraImageCapture) ReadyForCaptureChanged(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QCameraImageCapture22readyForCaptureChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameraimagecapture.h:128
// index:0
// Public Visibility=Default Availability=Available
// [-2] void bufferFormatChanged(QVideoFrame::PixelFormat)

/*
Signal emitted when the buffer format for the buffer image capture has changed.
*/
func (this *QCameraImageCapture) BufferFormatChanged(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QCameraImageCapture19bufferFormatChangedEN11QVideoFrame11PixelFormatE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameraimagecapture.h:129
// index:0
// Public Visibility=Default Availability=Available
// [-2] void captureDestinationChanged(QCameraImageCapture::CaptureDestinations)

/*
Signal emitted when the capture destination has changed.
*/
func (this *QCameraImageCapture) CaptureDestinationChanged(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QCameraImageCapture25captureDestinationChangedE6QFlagsINS_18CaptureDestinationEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameraimagecapture.h:131
// index:0
// Public Visibility=Default Availability=Available
// [-2] void imageExposed(int)

/*
Signal emitted when the frame with request id was exposed.
*/
func (this *QCameraImageCapture) ImageExposed(id int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QCameraImageCapture12imageExposedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), id)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameraimagecapture.h:132
// index:0
// Public Visibility=Default Availability=Available
// [-2] void imageCaptured(int, const QImage &)

/*
Signal emitted when the frame with request id was captured, but not processed and saved yet. Frame preview can be displayed to user.
*/
func (this *QCameraImageCapture) ImageCaptured(id int, preview qtgui.QImage_ITF) {
	var convArg1 unsafe.Pointer
	if preview != nil && preview.QImage_PTR() != nil {
		convArg1 = preview.QImage_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QCameraImageCapture13imageCapturedEiRK6QImage", qtrt.FFI_TYPE_POINTER, this.GetCthis(), id, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameraimagecapture.h:133
// index:0
// Public Visibility=Default Availability=Available
// [-2] void imageMetadataAvailable(int, const QString &, const QVariant &)

/*
Signals that a metadata for an image with request id is available. Also includes the key and value of the metadata.

This signal is emitted between imageExposed and imageSaved signals.
*/
func (this *QCameraImageCapture) ImageMetadataAvailable(id int, key string, value qtcore.QVariant_ITF) {
	var tmpArg1 = qtcore.NewQString_5(key)
	var convArg1 = tmpArg1.GetCthis()
	var convArg2 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg2 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QCameraImageCapture22imageMetadataAvailableEiRK7QStringRK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), id, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameraimagecapture.h:134
// index:0
// Public Visibility=Default Availability=Available
// [-2] void imageAvailable(int, const QVideoFrame &)

/*
Signal emitted when the frame with request id is available as buffer.
*/
func (this *QCameraImageCapture) ImageAvailable(id int, image QVideoFrame_ITF) {
	var convArg1 unsafe.Pointer
	if image != nil && image.QVideoFrame_PTR() != nil {
		convArg1 = image.QVideoFrame_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QCameraImageCapture14imageAvailableEiRK11QVideoFrame", qtrt.FFI_TYPE_POINTER, this.GetCthis(), id, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameraimagecapture.h:135
// index:0
// Public Visibility=Default Availability=Available
// [-2] void imageSaved(int, const QString &)

/*
Signal emitted when the frame with request id was saved to fileName.
*/
func (this *QCameraImageCapture) ImageSaved(id int, fileName string) {
	var tmpArg1 = qtcore.NewQString_5(fileName)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN19QCameraImageCapture10imageSavedEiRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), id, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameraimagecapture.h:138
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool setMediaObject(QMediaObject *)

/*
Reimplemented from QMediaBindableInterface::setMediaObject().

See also mediaObject().
*/
func (this *QCameraImageCapture) SetMediaObject(arg0 QMediaObject_ITF /*777 QMediaObject **/) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMediaObject_PTR() != nil {
		convArg0 = arg0.QMediaObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QCameraImageCapture14setMediaObjectEP12QMediaObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

/*

 */
type QCameraImageCapture__Error = int

// No Errors.
const QCameraImageCapture__NoError QCameraImageCapture__Error = 0

// The service is not ready for capture yet.
const QCameraImageCapture__NotReadyError QCameraImageCapture__Error = 1

// Device is not ready or not available.
const QCameraImageCapture__ResourceError QCameraImageCapture__Error = 2

// No space left on device.
const QCameraImageCapture__OutOfSpaceError QCameraImageCapture__Error = 3

// Device does not support stillimages capture.
const QCameraImageCapture__NotSupportedFeatureError QCameraImageCapture__Error = 4

// Current format is not supported.
const QCameraImageCapture__FormatError QCameraImageCapture__Error = 5

func (this *QCameraImageCapture) ErrorItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QCameraImageCapture_ErrorItemName(val int) string {
	var nilthis *QCameraImageCapture
	return nilthis.ErrorItemName(val)
}

/*

 */
type QCameraImageCapture__DriveMode = int

// Drive mode is capturing a single picture.
const QCameraImageCapture__SingleImageCapture QCameraImageCapture__DriveMode = 0

func (this *QCameraImageCapture) DriveModeItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QCameraImageCapture_DriveModeItemName(val int) string {
	var nilthis *QCameraImageCapture
	return nilthis.DriveModeItemName(val)
}

/*


 */
type QCameraImageCapture__CaptureDestination = int

//
const QCameraImageCapture__CaptureToFile QCameraImageCapture__CaptureDestination = 1

//
const QCameraImageCapture__CaptureToBuffer QCameraImageCapture__CaptureDestination = 2

func (this *QCameraImageCapture) CaptureDestinationItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QCameraImageCapture_CaptureDestinationItemName(val int) string {
	var nilthis *QCameraImageCapture
	return nilthis.CaptureDestinationItemName(val)
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
