package qtmultimedia

// /usr/include/qt/QtMultimedia/qaudiodeviceinfo.h
// #include <qaudiodeviceinfo.h>
// #include <QtMultimedia>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 25
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
type QAudioDeviceInfo struct {
	*qtrt.CObject
}
type QAudioDeviceInfo_ITF interface {
	QAudioDeviceInfo_PTR() *QAudioDeviceInfo
}

func (ptr *QAudioDeviceInfo) QAudioDeviceInfo_PTR() *QAudioDeviceInfo { return ptr }

func (this *QAudioDeviceInfo) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QAudioDeviceInfo) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQAudioDeviceInfoFromPointer(cthis unsafe.Pointer) *QAudioDeviceInfo {
	return &QAudioDeviceInfo{&qtrt.CObject{cthis}}
}
func (*QAudioDeviceInfo) NewFromPointer(cthis unsafe.Pointer) *QAudioDeviceInfo {
	return NewQAudioDeviceInfoFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qaudiodeviceinfo.h:67
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAudioDeviceInfo()

/*
Constructs an empty QAudioDeviceInfo object.
*/
func NewQAudioDeviceInfo() *QAudioDeviceInfo {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAudioDeviceInfoC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAudioDeviceInfoFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAudioDeviceInfo)
	return gothis
}

// /usr/include/qt/QtMultimedia/qaudiodeviceinfo.h:69
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QAudioDeviceInfo()

/*

 */
func DeleteQAudioDeviceInfo(this *QAudioDeviceInfo) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAudioDeviceInfoD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qaudiodeviceinfo.h:71
// index:0
// Public Visibility=Default Availability=Available
// [8] QAudioDeviceInfo & operator=(const QAudioDeviceInfo &)

/*

 */
func (this *QAudioDeviceInfo) Operator_equal(other QAudioDeviceInfo_ITF) *QAudioDeviceInfo {
	var convArg0 unsafe.Pointer
	if other != nil && other.QAudioDeviceInfo_PTR() != nil {
		convArg0 = other.QAudioDeviceInfo_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAudioDeviceInfoaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQAudioDeviceInfoFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQAudioDeviceInfo)
	return rv2
}

// /usr/include/qt/QtMultimedia/qaudiodeviceinfo.h:73
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QAudioDeviceInfo &) const

/*

 */
func (this *QAudioDeviceInfo) Operator_equal_equal(other QAudioDeviceInfo_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QAudioDeviceInfo_PTR() != nil {
		convArg0 = other.QAudioDeviceInfo_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QAudioDeviceInfoeqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qaudiodeviceinfo.h:74
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator!=(const QAudioDeviceInfo &) const

/*

 */
func (this *QAudioDeviceInfo) Operator_not_equal(other QAudioDeviceInfo_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QAudioDeviceInfo_PTR() != nil {
		convArg0 = other.QAudioDeviceInfo_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QAudioDeviceInfoneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qaudiodeviceinfo.h:76
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isNull() const

/*
Returns whether this QAudioDeviceInfo object holds a valid device definition.
*/
func (this *QAudioDeviceInfo) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QAudioDeviceInfo6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qaudiodeviceinfo.h:78
// index:0
// Public Visibility=Default Availability=Available
// [8] QString deviceName() const

/*
Returns the human readable name of the audio device.

Device names vary depending on the platform/audio plugin being used.

They are a unique string identifier for the audio device.

eg. default, Intel, U0x46d0x9a4
*/
func (this *QAudioDeviceInfo) DeviceName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QAudioDeviceInfo10deviceNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtMultimedia/qaudiodeviceinfo.h:80
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isFormatSupported(const QAudioFormat &) const

/*
Returns true if the supplied settings are supported by the audio device described by this QAudioDeviceInfo.
*/
func (this *QAudioDeviceInfo) IsFormatSupported(format QAudioFormat_ITF) bool {
	var convArg0 unsafe.Pointer
	if format != nil && format.QAudioFormat_PTR() != nil {
		convArg0 = format.QAudioFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QAudioDeviceInfo17isFormatSupportedERK12QAudioFormat", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qaudiodeviceinfo.h:81
// index:0
// Public Visibility=Default Availability=Available
// [8] QAudioFormat preferredFormat() const

/*
Returns the default audio format settings for this device.

These settings are provided by the platform/audio plugin being used.

They are also dependent on the QAudio::Mode being used.

A typical audio system would provide something like:


Input settings: 8000Hz mono 8 bit.
Output settings: 44100Hz stereo 16 bit little endian.
*/
func (this *QAudioDeviceInfo) PreferredFormat() *QAudioFormat /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QAudioDeviceInfo15preferredFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQAudioFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQAudioFormat)
	return rv2
}

// /usr/include/qt/QtMultimedia/qaudiodeviceinfo.h:82
// index:0
// Public Visibility=Default Availability=Available
// [8] QAudioFormat nearestFormat(const QAudioFormat &) const

/*
Returns the closest QAudioFormat to the supplied settings that the system supports.

These settings are provided by the platform/audio plugin being used.

They are also dependent on the QAudio::Mode being used.
*/
func (this *QAudioDeviceInfo) NearestFormat(format QAudioFormat_ITF) *QAudioFormat /*123*/ {
	var convArg0 unsafe.Pointer
	if format != nil && format.QAudioFormat_PTR() != nil {
		convArg0 = format.QAudioFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QAudioDeviceInfo13nearestFormatERK12QAudioFormat", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQAudioFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQAudioFormat)
	return rv2
}

// /usr/include/qt/QtMultimedia/qaudiodeviceinfo.h:84
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList supportedCodecs() const

/*
Returns a list of supported codecs.

All platform and plugin implementations should provide support for:

"audio/pcm" - Linear PCM

For writing plugins to support additional codecs refer to:

http://www.iana.org/assignments/media-types/audio/
*/
func (this *QAudioDeviceInfo) SupportedCodecs() *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QAudioDeviceInfo15supportedCodecsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtMultimedia/qaudiodeviceinfo.h:91
// index:0
// Public static Visibility=Default Availability=Available
// [8] QAudioDeviceInfo defaultInputDevice()

/*
Returns the information for the default input audio device. All platform and audio plugin implementations provide a default audio device to use.
*/
func (this *QAudioDeviceInfo) DefaultInputDevice() *QAudioDeviceInfo /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAudioDeviceInfo18defaultInputDeviceEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQAudioDeviceInfoFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQAudioDeviceInfo)
	return rv2
}
func QAudioDeviceInfo_DefaultInputDevice() *QAudioDeviceInfo /*123*/ {
	var nilthis *QAudioDeviceInfo
	rv := nilthis.DefaultInputDevice()
	return rv
}

// /usr/include/qt/QtMultimedia/qaudiodeviceinfo.h:92
// index:0
// Public static Visibility=Default Availability=Available
// [8] QAudioDeviceInfo defaultOutputDevice()

/*
Returns the information for the default output audio device. All platform and audio plugin implementations provide a default audio device to use.
*/
func (this *QAudioDeviceInfo) DefaultOutputDevice() *QAudioDeviceInfo /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAudioDeviceInfo19defaultOutputDeviceEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQAudioDeviceInfoFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQAudioDeviceInfo)
	return rv2
}
func QAudioDeviceInfo_DefaultOutputDevice() *QAudioDeviceInfo /*123*/ {
	var nilthis *QAudioDeviceInfo
	rv := nilthis.DefaultOutputDevice()
	return rv
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
