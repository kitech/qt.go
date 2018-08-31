package qtmultimedia

// /usr/include/qt/QtMultimedia/qmediaencodersettings.h
// #include <qmediaencodersettings.h>
// #include <QtMultimedia>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 13
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
type QAudioEncoderSettings struct {
	*qtrt.CObject
}
type QAudioEncoderSettings_ITF interface {
	QAudioEncoderSettings_PTR() *QAudioEncoderSettings
}

func (ptr *QAudioEncoderSettings) QAudioEncoderSettings_PTR() *QAudioEncoderSettings { return ptr }

func (this *QAudioEncoderSettings) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QAudioEncoderSettings) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQAudioEncoderSettingsFromPointer(cthis unsafe.Pointer) *QAudioEncoderSettings {
	return &QAudioEncoderSettings{&qtrt.CObject{cthis}}
}
func (*QAudioEncoderSettings) NewFromPointer(cthis unsafe.Pointer) *QAudioEncoderSettings {
	return NewQAudioEncoderSettingsFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qmediaencodersettings.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAudioEncoderSettings()

/*

 */
func (*QAudioEncoderSettings) NewForInherit() *QAudioEncoderSettings {
	return NewQAudioEncoderSettings()
}
func NewQAudioEncoderSettings() *QAudioEncoderSettings {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QAudioEncoderSettingsC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAudioEncoderSettingsFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAudioEncoderSettings)
	return gothis
}

// /usr/include/qt/QtMultimedia/qmediaencodersettings.h:61
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QAudioEncoderSettings()

/*

 */
func DeleteQAudioEncoderSettings(this *QAudioEncoderSettings) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QAudioEncoderSettingsD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qmediaencodersettings.h:63
// index:0
// Public Visibility=Default Availability=Available
// [8] QAudioEncoderSettings & operator=(const QAudioEncoderSettings &)

/*

 */
func (this *QAudioEncoderSettings) Operator_equal(other QAudioEncoderSettings_ITF) *QAudioEncoderSettings {
	var convArg0 unsafe.Pointer
	if other != nil && other.QAudioEncoderSettings_PTR() != nil {
		convArg0 = other.QAudioEncoderSettings_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QAudioEncoderSettingsaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQAudioEncoderSettingsFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQAudioEncoderSettings)
	return rv2
}

// /usr/include/qt/QtMultimedia/qmediaencodersettings.h:64
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QAudioEncoderSettings &) const

/*

 */
func (this *QAudioEncoderSettings) Operator_equal_equal(other QAudioEncoderSettings_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QAudioEncoderSettings_PTR() != nil {
		convArg0 = other.QAudioEncoderSettings_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QAudioEncoderSettingseqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qmediaencodersettings.h:65
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator!=(const QAudioEncoderSettings &) const

/*

 */
func (this *QAudioEncoderSettings) Operator_not_equal(other QAudioEncoderSettings_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QAudioEncoderSettings_PTR() != nil {
		convArg0 = other.QAudioEncoderSettings_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QAudioEncoderSettingsneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qmediaencodersettings.h:67
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isNull() const

/*

 */
func (this *QAudioEncoderSettings) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QAudioEncoderSettings6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qmediaencodersettings.h:69
// index:0
// Public Visibility=Default Availability=Available
// [4] QMultimedia::EncodingMode encodingMode() const

/*

 */
func (this *QAudioEncoderSettings) EncodingMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QAudioEncoderSettings12encodingModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qmediaencodersettings.h:70
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setEncodingMode(QMultimedia::EncodingMode)

/*

 */
func (this *QAudioEncoderSettings) SetEncodingMode(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QAudioEncoderSettings15setEncodingModeEN11QMultimedia12EncodingModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaencodersettings.h:72
// index:0
// Public Visibility=Default Availability=Available
// [8] QString codec() const

/*

 */
func (this *QAudioEncoderSettings) Codec() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QAudioEncoderSettings5codecEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtMultimedia/qmediaencodersettings.h:73
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCodec(const QString &)

/*

 */
func (this *QAudioEncoderSettings) SetCodec(codec string) {
	var tmpArg0 = qtcore.NewQString_5(codec)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QAudioEncoderSettings8setCodecERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaencodersettings.h:75
// index:0
// Public Visibility=Default Availability=Available
// [4] int bitRate() const

/*

 */
func (this *QAudioEncoderSettings) BitRate() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QAudioEncoderSettings7bitRateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qmediaencodersettings.h:76
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBitRate(int)

/*

 */
func (this *QAudioEncoderSettings) SetBitRate(bitrate int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QAudioEncoderSettings10setBitRateEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), bitrate)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaencodersettings.h:78
// index:0
// Public Visibility=Default Availability=Available
// [4] int channelCount() const

/*

 */
func (this *QAudioEncoderSettings) ChannelCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QAudioEncoderSettings12channelCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qmediaencodersettings.h:79
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setChannelCount(int)

/*

 */
func (this *QAudioEncoderSettings) SetChannelCount(channels int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QAudioEncoderSettings15setChannelCountEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), channels)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaencodersettings.h:81
// index:0
// Public Visibility=Default Availability=Available
// [4] int sampleRate() const

/*

 */
func (this *QAudioEncoderSettings) SampleRate() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QAudioEncoderSettings10sampleRateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qmediaencodersettings.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSampleRate(int)

/*

 */
func (this *QAudioEncoderSettings) SetSampleRate(rate int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QAudioEncoderSettings13setSampleRateEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), rate)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaencodersettings.h:84
// index:0
// Public Visibility=Default Availability=Available
// [4] QMultimedia::EncodingQuality quality() const

/*

 */
func (this *QAudioEncoderSettings) Quality() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QAudioEncoderSettings7qualityEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qmediaencodersettings.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setQuality(QMultimedia::EncodingQuality)

/*

 */
func (this *QAudioEncoderSettings) SetQuality(quality int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QAudioEncoderSettings10setQualityEN11QMultimedia15EncodingQualityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), quality)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaencodersettings.h:87
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant encodingOption(const QString &) const

/*

 */
func (this *QAudioEncoderSettings) EncodingOption(option string) *qtcore.QVariant /*123*/ {
	var tmpArg0 = qtcore.NewQString_5(option)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QAudioEncoderSettings14encodingOptionERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtMultimedia/qmediaencodersettings.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setEncodingOption(const QString &, const QVariant &)

/*

 */
func (this *QAudioEncoderSettings) SetEncodingOption(option string, value qtcore.QVariant_ITF) {
	var tmpArg0 = qtcore.NewQString_5(option)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg1 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QAudioEncoderSettings17setEncodingOptionERK7QStringRK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
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
