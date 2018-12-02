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
// extern C begin: 21
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
type QImageEncoderSettings struct {
	*qtrt.CObject
}
type QImageEncoderSettings_ITF interface {
	QImageEncoderSettings_PTR() *QImageEncoderSettings
}

func (ptr *QImageEncoderSettings) QImageEncoderSettings_PTR() *QImageEncoderSettings { return ptr }

func (this *QImageEncoderSettings) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QImageEncoderSettings) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQImageEncoderSettingsFromPointer(cthis unsafe.Pointer) *QImageEncoderSettings {
	return &QImageEncoderSettings{&qtrt.CObject{cthis}}
}
func (*QImageEncoderSettings) NewFromPointer(cthis unsafe.Pointer) *QImageEncoderSettings {
	return NewQImageEncoderSettingsFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qmediaencodersettings.h:143
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QImageEncoderSettings()

/*

 */
func (*QImageEncoderSettings) NewForInherit() *QImageEncoderSettings {
	return NewQImageEncoderSettings()
}
func NewQImageEncoderSettings() *QImageEncoderSettings {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QImageEncoderSettingsC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQImageEncoderSettingsFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQImageEncoderSettings)
	return gothis
}

// /usr/include/qt/QtMultimedia/qmediaencodersettings.h:146
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QImageEncoderSettings()

/*

 */
func DeleteQImageEncoderSettings(this *QImageEncoderSettings) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QImageEncoderSettingsD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qmediaencodersettings.h:148
// index:0
// Public Visibility=Default Availability=Available
// [8] QImageEncoderSettings & operator=(const QImageEncoderSettings &)

/*

 */
func (this *QImageEncoderSettings) Operator_equal(other QImageEncoderSettings_ITF) *QImageEncoderSettings {
	var convArg0 unsafe.Pointer
	if other != nil && other.QImageEncoderSettings_PTR() != nil {
		convArg0 = other.QImageEncoderSettings_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QImageEncoderSettingsaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageEncoderSettingsFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImageEncoderSettings)
	return rv2
}

// /usr/include/qt/QtMultimedia/qmediaencodersettings.h:149
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QImageEncoderSettings &) const

/*

 */
func (this *QImageEncoderSettings) Operator_equal_equal(other QImageEncoderSettings_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QImageEncoderSettings_PTR() != nil {
		convArg0 = other.QImageEncoderSettings_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QImageEncoderSettingseqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qmediaencodersettings.h:150
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator!=(const QImageEncoderSettings &) const

/*

 */
func (this *QImageEncoderSettings) Operator_not_equal(other QImageEncoderSettings_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QImageEncoderSettings_PTR() != nil {
		convArg0 = other.QImageEncoderSettings_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QImageEncoderSettingsneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qmediaencodersettings.h:152
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isNull() const

/*

 */
func (this *QImageEncoderSettings) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QImageEncoderSettings6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qmediaencodersettings.h:154
// index:0
// Public Visibility=Default Availability=Available
// [8] QString codec() const

/*

 */
func (this *QImageEncoderSettings) Codec() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QImageEncoderSettings5codecEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtMultimedia/qmediaencodersettings.h:155
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCodec(const QString &)

/*

 */
func (this *QImageEncoderSettings) SetCodec(arg0 string) {
	var tmpArg0 = qtcore.NewQString_5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QImageEncoderSettings8setCodecERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaencodersettings.h:157
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize resolution() const

/*

 */
func (this *QImageEncoderSettings) Resolution() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QImageEncoderSettings10resolutionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtMultimedia/qmediaencodersettings.h:158
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setResolution(const QSize &)

/*

 */
func (this *QImageEncoderSettings) SetResolution(arg0 qtcore.QSize_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QSize_PTR() != nil {
		convArg0 = arg0.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QImageEncoderSettings13setResolutionERK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaencodersettings.h:159
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setResolution(int, int)

/*

 */
func (this *QImageEncoderSettings) SetResolution_1(width int, height int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QImageEncoderSettings13setResolutionEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), width, height)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaencodersettings.h:161
// index:0
// Public Visibility=Default Availability=Available
// [4] QMultimedia::EncodingQuality quality() const

/*

 */
func (this *QImageEncoderSettings) Quality() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QImageEncoderSettings7qualityEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qmediaencodersettings.h:162
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setQuality(QMultimedia::EncodingQuality)

/*

 */
func (this *QImageEncoderSettings) SetQuality(quality int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QImageEncoderSettings10setQualityEN11QMultimedia15EncodingQualityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), quality)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaencodersettings.h:164
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant encodingOption(const QString &) const

/*

 */
func (this *QImageEncoderSettings) EncodingOption(option string) *qtcore.QVariant /*123*/ {
	var tmpArg0 = qtcore.NewQString_5(option)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QImageEncoderSettings14encodingOptionERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtMultimedia/qmediaencodersettings.h:166
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setEncodingOption(const QString &, const QVariant &)

/*

 */
func (this *QImageEncoderSettings) SetEncodingOption(option string, value qtcore.QVariant_ITF) {
	var tmpArg0 = qtcore.NewQString_5(option)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg1 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QImageEncoderSettings17setEncodingOptionERK7QStringRK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
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
