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
// extern C begin: 20
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
type QVideoEncoderSettings struct {
	*qtrt.CObject
}
type QVideoEncoderSettings_ITF interface {
	QVideoEncoderSettings_PTR() *QVideoEncoderSettings
}

func (ptr *QVideoEncoderSettings) QVideoEncoderSettings_PTR() *QVideoEncoderSettings { return ptr }

func (this *QVideoEncoderSettings) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QVideoEncoderSettings) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQVideoEncoderSettingsFromPointer(cthis unsafe.Pointer) *QVideoEncoderSettings {
	return &QVideoEncoderSettings{&qtrt.CObject{cthis}}
}
func (*QVideoEncoderSettings) NewFromPointer(cthis unsafe.Pointer) *QVideoEncoderSettings {
	return NewQVideoEncoderSettingsFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qmediaencodersettings.h:100
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QVideoEncoderSettings()

/*

 */
func (*QVideoEncoderSettings) NewForInherit() *QVideoEncoderSettings {
	return NewQVideoEncoderSettings()
}
func NewQVideoEncoderSettings() *QVideoEncoderSettings {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QVideoEncoderSettingsC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVideoEncoderSettingsFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVideoEncoderSettings)
	return gothis
}

// /usr/include/qt/QtMultimedia/qmediaencodersettings.h:103
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QVideoEncoderSettings()

/*

 */
func DeleteQVideoEncoderSettings(this *QVideoEncoderSettings) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QVideoEncoderSettingsD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qmediaencodersettings.h:105
// index:0
// Public Visibility=Default Availability=Available
// [8] QVideoEncoderSettings & operator=(const QVideoEncoderSettings &)

/*

 */
func (this *QVideoEncoderSettings) Operator_equal(other QVideoEncoderSettings_ITF) *QVideoEncoderSettings {
	var convArg0 unsafe.Pointer
	if other != nil && other.QVideoEncoderSettings_PTR() != nil {
		convArg0 = other.QVideoEncoderSettings_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QVideoEncoderSettingsaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVideoEncoderSettingsFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVideoEncoderSettings)
	return rv2
}

// /usr/include/qt/QtMultimedia/qmediaencodersettings.h:106
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QVideoEncoderSettings &) const

/*

 */
func (this *QVideoEncoderSettings) Operator_equal_equal(other QVideoEncoderSettings_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QVideoEncoderSettings_PTR() != nil {
		convArg0 = other.QVideoEncoderSettings_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QVideoEncoderSettingseqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qmediaencodersettings.h:107
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator!=(const QVideoEncoderSettings &) const

/*

 */
func (this *QVideoEncoderSettings) Operator_not_equal(other QVideoEncoderSettings_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QVideoEncoderSettings_PTR() != nil {
		convArg0 = other.QVideoEncoderSettings_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QVideoEncoderSettingsneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qmediaencodersettings.h:109
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isNull() const

/*

 */
func (this *QVideoEncoderSettings) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QVideoEncoderSettings6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qmediaencodersettings.h:111
// index:0
// Public Visibility=Default Availability=Available
// [4] QMultimedia::EncodingMode encodingMode() const

/*

 */
func (this *QVideoEncoderSettings) EncodingMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QVideoEncoderSettings12encodingModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qmediaencodersettings.h:112
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setEncodingMode(QMultimedia::EncodingMode)

/*

 */
func (this *QVideoEncoderSettings) SetEncodingMode(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QVideoEncoderSettings15setEncodingModeEN11QMultimedia12EncodingModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaencodersettings.h:114
// index:0
// Public Visibility=Default Availability=Available
// [8] QString codec() const

/*

 */
func (this *QVideoEncoderSettings) Codec() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QVideoEncoderSettings5codecEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtMultimedia/qmediaencodersettings.h:115
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCodec(const QString &)

/*

 */
func (this *QVideoEncoderSettings) SetCodec(arg0 string) {
	var tmpArg0 = qtcore.NewQString5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QVideoEncoderSettings8setCodecERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaencodersettings.h:117
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize resolution() const

/*

 */
func (this *QVideoEncoderSettings) Resolution() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QVideoEncoderSettings10resolutionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtMultimedia/qmediaencodersettings.h:118
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setResolution(const QSize &)

/*

 */
func (this *QVideoEncoderSettings) SetResolution(arg0 qtcore.QSize_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QSize_PTR() != nil {
		convArg0 = arg0.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QVideoEncoderSettings13setResolutionERK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaencodersettings.h:119
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setResolution(int, int)

/*

 */
func (this *QVideoEncoderSettings) SetResolution1(width int, height int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QVideoEncoderSettings13setResolutionEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), width, height)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaencodersettings.h:121
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal frameRate() const

/*

 */
func (this *QVideoEncoderSettings) FrameRate() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QVideoEncoderSettings9frameRateEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtMultimedia/qmediaencodersettings.h:122
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFrameRate(qreal)

/*

 */
func (this *QVideoEncoderSettings) SetFrameRate(rate float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QVideoEncoderSettings12setFrameRateEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), rate)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaencodersettings.h:124
// index:0
// Public Visibility=Default Availability=Available
// [4] int bitRate() const

/*

 */
func (this *QVideoEncoderSettings) BitRate() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QVideoEncoderSettings7bitRateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qmediaencodersettings.h:125
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBitRate(int)

/*

 */
func (this *QVideoEncoderSettings) SetBitRate(bitrate int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QVideoEncoderSettings10setBitRateEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), bitrate)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaencodersettings.h:127
// index:0
// Public Visibility=Default Availability=Available
// [4] QMultimedia::EncodingQuality quality() const

/*

 */
func (this *QVideoEncoderSettings) Quality() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QVideoEncoderSettings7qualityEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qmediaencodersettings.h:128
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setQuality(QMultimedia::EncodingQuality)

/*

 */
func (this *QVideoEncoderSettings) SetQuality(quality int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QVideoEncoderSettings10setQualityEN11QMultimedia15EncodingQualityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), quality)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaencodersettings.h:130
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant encodingOption(const QString &) const

/*

 */
func (this *QVideoEncoderSettings) EncodingOption(option string) *qtcore.QVariant /*123*/ {
	var tmpArg0 = qtcore.NewQString5(option)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QVideoEncoderSettings14encodingOptionERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtMultimedia/qmediaencodersettings.h:132
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setEncodingOption(const QString &, const QVariant &)

/*

 */
func (this *QVideoEncoderSettings) SetEncodingOption(option string, value qtcore.QVariant_ITF) {
	var tmpArg0 = qtcore.NewQString5(option)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg1 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QVideoEncoderSettings17setEncodingOptionERK7QStringRK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

//  body block end

//  keep block begin

func init_unused_11763() {
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
