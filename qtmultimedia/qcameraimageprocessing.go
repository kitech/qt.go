package qtmultimedia

// /usr/include/qt/QtMultimedia/qcameraimageprocessing.h
// #include <qcameraimageprocessing.h>
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
type QCameraImageProcessing struct {
	*qtcore.QObject
}
type QCameraImageProcessing_ITF interface {
	qtcore.QObject_ITF
	QCameraImageProcessing_PTR() *QCameraImageProcessing
}

func (ptr *QCameraImageProcessing) QCameraImageProcessing_PTR() *QCameraImageProcessing { return ptr }

func (this *QCameraImageProcessing) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QCameraImageProcessing) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQCameraImageProcessingFromPointer(cthis unsafe.Pointer) *QCameraImageProcessing {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QCameraImageProcessing{bcthis0}
}
func (*QCameraImageProcessing) NewFromPointer(cthis unsafe.Pointer) *QCameraImageProcessing {
	return NewQCameraImageProcessingFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qcameraimageprocessing.h:62
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QCameraImageProcessing) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QCameraImageProcessing10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qcameraimageprocessing.h:91
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isAvailable() const

/*
Returns true if image processing related settings are supported by this camera.
*/
func (this *QCameraImageProcessing) IsAvailable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QCameraImageProcessing11isAvailableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qcameraimageprocessing.h:93
// index:0
// Public Visibility=Default Availability=Available
// [4] QCameraImageProcessing::WhiteBalanceMode whiteBalanceMode() const

/*
Returns the white balance mode being used.

See also setWhiteBalanceMode().
*/
func (this *QCameraImageProcessing) WhiteBalanceMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QCameraImageProcessing16whiteBalanceModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qcameraimageprocessing.h:94
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWhiteBalanceMode(QCameraImageProcessing::WhiteBalanceMode)

/*
Sets the white balance to mode.

See also whiteBalanceMode().
*/
func (this *QCameraImageProcessing) SetWhiteBalanceMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QCameraImageProcessing19setWhiteBalanceModeENS_16WhiteBalanceModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameraimageprocessing.h:95
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isWhiteBalanceModeSupported(QCameraImageProcessing::WhiteBalanceMode) const

/*
Returns true if the white balance mode is supported.
*/
func (this *QCameraImageProcessing) IsWhiteBalanceModeSupported(mode int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QCameraImageProcessing27isWhiteBalanceModeSupportedENS_16WhiteBalanceModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qcameraimageprocessing.h:97
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal manualWhiteBalance() const

/*
Returns the current color temperature if the current white balance mode is WhiteBalanceManual. For other modes the return value is undefined.

See also setManualWhiteBalance().
*/
func (this *QCameraImageProcessing) ManualWhiteBalance() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QCameraImageProcessing18manualWhiteBalanceEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtMultimedia/qcameraimageprocessing.h:98
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setManualWhiteBalance(qreal)

/*
Sets manual white balance to colorTemperature. This is used when whiteBalanceMode() is set to WhiteBalanceManual. The units are Kelvin.

See also manualWhiteBalance().
*/
func (this *QCameraImageProcessing) SetManualWhiteBalance(colorTemperature float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QCameraImageProcessing21setManualWhiteBalanceEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), colorTemperature)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameraimageprocessing.h:100
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal brightness() const

/*
Returns the brightness adjustment setting.

See also setBrightness().
*/
func (this *QCameraImageProcessing) Brightness() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QCameraImageProcessing10brightnessEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtMultimedia/qcameraimageprocessing.h:101
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBrightness(qreal)

/*
Set the brightness adjustment to value.

Valid brightness adjustment values range between -1.0 and 1.0, with a default of 0.

See also brightness().
*/
func (this *QCameraImageProcessing) SetBrightness(value float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QCameraImageProcessing13setBrightnessEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), value)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameraimageprocessing.h:103
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal contrast() const

/*
Returns the contrast adjustment setting.

See also setContrast().
*/
func (this *QCameraImageProcessing) Contrast() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QCameraImageProcessing8contrastEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtMultimedia/qcameraimageprocessing.h:104
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setContrast(qreal)

/*
Set the contrast adjustment to value.

Valid contrast adjustment values range between -1.0 and 1.0, with a default of 0.

See also contrast().
*/
func (this *QCameraImageProcessing) SetContrast(value float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QCameraImageProcessing11setContrastEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), value)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameraimageprocessing.h:106
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal saturation() const

/*
Returns the saturation adjustment value.

See also setSaturation().
*/
func (this *QCameraImageProcessing) Saturation() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QCameraImageProcessing10saturationEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtMultimedia/qcameraimageprocessing.h:107
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSaturation(qreal)

/*
Sets the saturation adjustment value to value.

Valid saturation values range between -1.0 and 1.0, with a default of 0.

See also saturation().
*/
func (this *QCameraImageProcessing) SetSaturation(value float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QCameraImageProcessing13setSaturationEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), value)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameraimageprocessing.h:109
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal sharpeningLevel() const

/*
Returns the sharpening adjustment level.

See also setSharpeningLevel().
*/
func (this *QCameraImageProcessing) SharpeningLevel() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QCameraImageProcessing15sharpeningLevelEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtMultimedia/qcameraimageprocessing.h:110
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSharpeningLevel(qreal)

/*
Sets the sharpening adjustment level.

Valid sharpening values range between -1.0 and 1.0, with a default of 0.

See also sharpeningLevel().
*/
func (this *QCameraImageProcessing) SetSharpeningLevel(value float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QCameraImageProcessing18setSharpeningLevelEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), value)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameraimageprocessing.h:112
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal denoisingLevel() const

/*
Returns the denoising adjustment level.

See also setDenoisingLevel().
*/
func (this *QCameraImageProcessing) DenoisingLevel() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QCameraImageProcessing14denoisingLevelEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtMultimedia/qcameraimageprocessing.h:113
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDenoisingLevel(qreal)

/*
Sets the denoising adjustment level.

Valid denoising values range between -1.0 and 1.0, with a default of 0.

If the parameter value is set to 0, the amount of denoising applied is selected by camera and depends on camera capabilities and settings. Changing value in -1.0..1.0 range adjusts the amount of denoising applied within the supported range.

See also denoisingLevel().
*/
func (this *QCameraImageProcessing) SetDenoisingLevel(value float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QCameraImageProcessing17setDenoisingLevelEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), value)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameraimageprocessing.h:115
// index:0
// Public Visibility=Default Availability=Available
// [4] QCameraImageProcessing::ColorFilter colorFilter() const

/*
Returns the color filter which will be applied to image data captured by the camera.

This function was introduced in  Qt 5.5.

See also setColorFilter().
*/
func (this *QCameraImageProcessing) ColorFilter() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QCameraImageProcessing11colorFilterEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qcameraimageprocessing.h:116
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setColorFilter(QCameraImageProcessing::ColorFilter)

/*
Sets the color filter which will be applied to image data captured by the camera.

This function was introduced in  Qt 5.5.

See also colorFilter().
*/
func (this *QCameraImageProcessing) SetColorFilter(filter int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QCameraImageProcessing14setColorFilterENS_11ColorFilterE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), filter)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameraimageprocessing.h:117
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isColorFilterSupported(QCameraImageProcessing::ColorFilter) const

/*
Returns true if a color filter is supported.

This function was introduced in  Qt 5.5.
*/
func (this *QCameraImageProcessing) IsColorFilterSupported(filter int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QCameraImageProcessing22isColorFilterSupportedENS_11ColorFilterE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), filter)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

func DeleteQCameraImageProcessing(this *QCameraImageProcessing) {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QCameraImageProcessingD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

/*

 */
type QCameraImageProcessing__WhiteBalanceMode = int

// Auto white balance mode.
const QCameraImageProcessing__WhiteBalanceAuto QCameraImageProcessing__WhiteBalanceMode = 0

// Manual white balance. In this mode the white balance should be set with setManualWhiteBalance()
const QCameraImageProcessing__WhiteBalanceManual QCameraImageProcessing__WhiteBalanceMode = 1

// Sunlight white balance mode.
const QCameraImageProcessing__WhiteBalanceSunlight QCameraImageProcessing__WhiteBalanceMode = 2

// Cloudy white balance mode.
const QCameraImageProcessing__WhiteBalanceCloudy QCameraImageProcessing__WhiteBalanceMode = 3

// Shade white balance mode.
const QCameraImageProcessing__WhiteBalanceShade QCameraImageProcessing__WhiteBalanceMode = 4

// Tungsten (incandescent) white balance mode.
const QCameraImageProcessing__WhiteBalanceTungsten QCameraImageProcessing__WhiteBalanceMode = 5

// Fluorescent white balance mode.
const QCameraImageProcessing__WhiteBalanceFluorescent QCameraImageProcessing__WhiteBalanceMode = 6

// Flash white balance mode.
const QCameraImageProcessing__WhiteBalanceFlash QCameraImageProcessing__WhiteBalanceMode = 7

// Sunset white balance mode.
const QCameraImageProcessing__WhiteBalanceSunset QCameraImageProcessing__WhiteBalanceMode = 8

//
const QCameraImageProcessing__WhiteBalanceVendor QCameraImageProcessing__WhiteBalanceMode = 1000

func (this *QCameraImageProcessing) WhiteBalanceModeItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QCameraImageProcessing_WhiteBalanceModeItemName(val int) string {
	var nilthis *QCameraImageProcessing
	return nilthis.WhiteBalanceModeItemName(val)
}

/*


This enum was introduced or modified in  Qt 5.5.

*/
type QCameraImageProcessing__ColorFilter = int

// No filter is applied to images.
const QCameraImageProcessing__ColorFilterNone QCameraImageProcessing__ColorFilter = 0

// A grayscale filter.
const QCameraImageProcessing__ColorFilterGrayscale QCameraImageProcessing__ColorFilter = 1

// A negative filter.
const QCameraImageProcessing__ColorFilterNegative QCameraImageProcessing__ColorFilter = 2

// A solarize filter.
const QCameraImageProcessing__ColorFilterSolarize QCameraImageProcessing__ColorFilter = 3

// A sepia filter.
const QCameraImageProcessing__ColorFilterSepia QCameraImageProcessing__ColorFilter = 4

// A posterize filter.
const QCameraImageProcessing__ColorFilterPosterize QCameraImageProcessing__ColorFilter = 5

// A whiteboard filter.
const QCameraImageProcessing__ColorFilterWhiteboard QCameraImageProcessing__ColorFilter = 6

// A blackboard filter.
const QCameraImageProcessing__ColorFilterBlackboard QCameraImageProcessing__ColorFilter = 7

// An aqua filter.
const QCameraImageProcessing__ColorFilterAqua QCameraImageProcessing__ColorFilter = 8

//
const QCameraImageProcessing__ColorFilterVendor QCameraImageProcessing__ColorFilter = 1000

func (this *QCameraImageProcessing) ColorFilterItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QCameraImageProcessing_ColorFilterItemName(val int) string {
	var nilthis *QCameraImageProcessing
	return nilthis.ColorFilterItemName(val)
}

//  body block end

//  keep block begin

func init_unused_11805() {
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
