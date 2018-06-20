package qtmultimedia

// /usr/include/qt/QtMultimedia/qcameraviewfindersettings.h
// #include <qcameraviewfindersettings.h>
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
type QCameraViewfinderSettings struct {
	*qtrt.CObject
}
type QCameraViewfinderSettings_ITF interface {
	QCameraViewfinderSettings_PTR() *QCameraViewfinderSettings
}

func (ptr *QCameraViewfinderSettings) QCameraViewfinderSettings_PTR() *QCameraViewfinderSettings {
	return ptr
}

func (this *QCameraViewfinderSettings) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QCameraViewfinderSettings) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQCameraViewfinderSettingsFromPointer(cthis unsafe.Pointer) *QCameraViewfinderSettings {
	return &QCameraViewfinderSettings{&qtrt.CObject{cthis}}
}
func (*QCameraViewfinderSettings) NewFromPointer(cthis unsafe.Pointer) *QCameraViewfinderSettings {
	return NewQCameraViewfinderSettingsFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qcameraviewfindersettings.h:56
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QCameraViewfinderSettings()

/*
Constructs a null viewfinder settings object.
*/
func NewQCameraViewfinderSettings() *QCameraViewfinderSettings {
	rv, err := qtrt.InvokeQtFunc6("_ZN25QCameraViewfinderSettingsC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCameraViewfinderSettingsFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQCameraViewfinderSettings)
	return gothis
}

// /usr/include/qt/QtMultimedia/qcameraviewfindersettings.h:59
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QCameraViewfinderSettings()

/*

 */
func DeleteQCameraViewfinderSettings(this *QCameraViewfinderSettings) {
	rv, err := qtrt.InvokeQtFunc6("_ZN25QCameraViewfinderSettingsD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qcameraviewfindersettings.h:61
// index:0
// Public Visibility=Default Availability=Available
// [8] QCameraViewfinderSettings & operator=(const QCameraViewfinderSettings &)

/*

 */
func (this *QCameraViewfinderSettings) Operator_equal(other QCameraViewfinderSettings_ITF) *QCameraViewfinderSettings {
	var convArg0 unsafe.Pointer
	if other != nil && other.QCameraViewfinderSettings_PTR() != nil {
		convArg0 = other.QCameraViewfinderSettings_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN25QCameraViewfinderSettingsaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCameraViewfinderSettingsFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCameraViewfinderSettings)
	return rv2
}

// /usr/include/qt/QtMultimedia/qcameraviewfindersettings.h:63
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QCameraViewfinderSettings & operator=(QCameraViewfinderSettings &&)

/*

 */
func (this *QCameraViewfinderSettings) Operator_equal_1(other unsafe.Pointer /*333*/) *QCameraViewfinderSettings {
	rv, err := qtrt.InvokeQtFunc6("_ZN25QCameraViewfinderSettingsaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCameraViewfinderSettingsFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCameraViewfinderSettings)
	return rv2
}

// /usr/include/qt/QtMultimedia/qcameraviewfindersettings.h:67
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QCameraViewfinderSettings &)

/*
Swaps this viewfinder settings object with other. This function is very fast and never fails.
*/
func (this *QCameraViewfinderSettings) Swap(other QCameraViewfinderSettings_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QCameraViewfinderSettings_PTR() != nil {
		convArg0 = other.QCameraViewfinderSettings_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN25QCameraViewfinderSettings4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameraviewfindersettings.h:70
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isNull() const

/*
Identifies if a viewfinder settings object is uninitalized.

Returns true if the settings are null, and false if they are not.
*/
func (this *QCameraViewfinderSettings) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK25QCameraViewfinderSettings6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qcameraviewfindersettings.h:72
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize resolution() const

/*
Returns the viewfinder resolution.

See also setResolution().
*/
func (this *QCameraViewfinderSettings) Resolution() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK25QCameraViewfinderSettings10resolutionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtMultimedia/qcameraviewfindersettings.h:73
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setResolution(const QSize &)

/*
Sets the viewfinder resolution.

If the given resolution is empty, the backend makes an optimal choice based on the supported resolutions and the other viewfinder settings.

If the camera is used to capture videos or images, the viewfinder resolution might be ignored if it conflicts with the capture resolution.

See also resolution(), QVideoEncoderSettings::setResolution(), QImageEncoderSettings::setResolution(), and QCamera::supportedViewfinderResolutions().
*/
func (this *QCameraViewfinderSettings) SetResolution(arg0 qtcore.QSize_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QSize_PTR() != nil {
		convArg0 = arg0.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN25QCameraViewfinderSettings13setResolutionERK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameraviewfindersettings.h:74
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void setResolution(int, int)

/*
Sets the viewfinder resolution.

If the given resolution is empty, the backend makes an optimal choice based on the supported resolutions and the other viewfinder settings.

If the camera is used to capture videos or images, the viewfinder resolution might be ignored if it conflicts with the capture resolution.

See also resolution(), QVideoEncoderSettings::setResolution(), QImageEncoderSettings::setResolution(), and QCamera::supportedViewfinderResolutions().
*/
func (this *QCameraViewfinderSettings) SetResolution_1(width int, height int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN25QCameraViewfinderSettings13setResolutionEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), width, height)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameraviewfindersettings.h:77
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal minimumFrameRate() const

/*
Returns the viewfinder minimum frame rate in frames per second.

See also setMinimumFrameRate() and maximumFrameRate().
*/
func (this *QCameraViewfinderSettings) MinimumFrameRate() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK25QCameraViewfinderSettings16minimumFrameRateEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtMultimedia/qcameraviewfindersettings.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMinimumFrameRate(qreal)

/*
Sets the viewfinder minimum frame rate in frames per second.

If the minimum frame rate is equal to the maximum frame rate, the frame rate is fixed. If not, the actual frame rate fluctuates between the minimum and the maximum.

If the given rate equals to 0, the backend makes an optimal choice based on the supported frame rates and the other viewfinder settings.

See also minimumFrameRate(), setMaximumFrameRate(), and QCamera::supportedViewfinderFrameRateRanges().
*/
func (this *QCameraViewfinderSettings) SetMinimumFrameRate(rate float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN25QCameraViewfinderSettings19setMinimumFrameRateEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), rate)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameraviewfindersettings.h:80
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal maximumFrameRate() const

/*
Returns the viewfinder maximum frame rate in frames per second.

See also setMaximumFrameRate() and minimumFrameRate().
*/
func (this *QCameraViewfinderSettings) MaximumFrameRate() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK25QCameraViewfinderSettings16maximumFrameRateEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtMultimedia/qcameraviewfindersettings.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMaximumFrameRate(qreal)

/*
Sets the viewfinder maximum frame rate in frames per second.

If the maximum frame rate is equal to the minimum frame rate, the frame rate is fixed. If not, the actual frame rate fluctuates between the minimum and the maximum.

If the given rate equals to 0, the backend makes an optimal choice based on the supported frame rates and the other viewfinder settings.

See also maximumFrameRate(), setMinimumFrameRate(), and QCamera::supportedViewfinderFrameRateRanges().
*/
func (this *QCameraViewfinderSettings) SetMaximumFrameRate(rate float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN25QCameraViewfinderSettings19setMaximumFrameRateEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), rate)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameraviewfindersettings.h:83
// index:0
// Public Visibility=Default Availability=Available
// [4] QVideoFrame::PixelFormat pixelFormat() const

/*
Returns the viewfinder pixel format.

See also setPixelFormat().
*/
func (this *QCameraViewfinderSettings) PixelFormat() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK25QCameraViewfinderSettings11pixelFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qcameraviewfindersettings.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPixelFormat(QVideoFrame::PixelFormat)

/*
Sets the viewfinder pixel format.

If the given format is equal to QVideoFrame::Format_Invalid, the backend uses the default format.

See also pixelFormat() and QCamera::supportedViewfinderPixelFormats().
*/
func (this *QCameraViewfinderSettings) SetPixelFormat(format int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN25QCameraViewfinderSettings14setPixelFormatEN11QVideoFrame11PixelFormatE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), format)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameraviewfindersettings.h:86
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize pixelAspectRatio() const

/*
Returns the viewfinder pixel aspect ratio.

See also setPixelAspectRatio().
*/
func (this *QCameraViewfinderSettings) PixelAspectRatio() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK25QCameraViewfinderSettings16pixelAspectRatioEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtMultimedia/qcameraviewfindersettings.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPixelAspectRatio(const QSize &)

/*
Sets the viewfinder pixel aspect ratio.

See also pixelAspectRatio().
*/
func (this *QCameraViewfinderSettings) SetPixelAspectRatio(ratio qtcore.QSize_ITF) {
	var convArg0 unsafe.Pointer
	if ratio != nil && ratio.QSize_PTR() != nil {
		convArg0 = ratio.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN25QCameraViewfinderSettings19setPixelAspectRatioERK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameraviewfindersettings.h:88
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void setPixelAspectRatio(int, int)

/*
Sets the viewfinder pixel aspect ratio.

See also pixelAspectRatio().
*/
func (this *QCameraViewfinderSettings) SetPixelAspectRatio_1(horizontal int, vertical int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN25QCameraViewfinderSettings19setPixelAspectRatioEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), horizontal, vertical)
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
