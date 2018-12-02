package qtmultimedia

// /usr/include/qt/QtMultimedia/qvideoencodersettingscontrol.h
// #include <qvideoencodersettingscontrol.h>
// #include <QtMultimedia>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 12
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
type QVideoEncoderSettingsControl struct {
	*QMediaControl
}
type QVideoEncoderSettingsControl_ITF interface {
	QMediaControl_ITF
	QVideoEncoderSettingsControl_PTR() *QVideoEncoderSettingsControl
}

func (ptr *QVideoEncoderSettingsControl) QVideoEncoderSettingsControl_PTR() *QVideoEncoderSettingsControl {
	return ptr
}

func (this *QVideoEncoderSettingsControl) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QMediaControl.GetCthis()
	}
}
func (this *QVideoEncoderSettingsControl) SetCthis(cthis unsafe.Pointer) {
	this.QMediaControl = NewQMediaControlFromPointer(cthis)
}
func NewQVideoEncoderSettingsControlFromPointer(cthis unsafe.Pointer) *QVideoEncoderSettingsControl {
	bcthis0 := NewQMediaControlFromPointer(cthis)
	return &QVideoEncoderSettingsControl{bcthis0}
}
func (*QVideoEncoderSettingsControl) NewFromPointer(cthis unsafe.Pointer) *QVideoEncoderSettingsControl {
	return NewQVideoEncoderSettingsControlFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qvideoencodersettingscontrol.h:62
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QVideoEncoderSettingsControl) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK28QVideoEncoderSettingsControl10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qvideoencodersettingscontrol.h:65
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QVideoEncoderSettingsControl()

/*

 */
func DeleteQVideoEncoderSettingsControl(this *QVideoEncoderSettingsControl) {
	rv, err := qtrt.InvokeQtFunc6("_ZN28QVideoEncoderSettingsControlD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qvideoencodersettingscontrol.h:67
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QList<QSize> supportedResolutions(const QVideoEncoderSettings &, bool *) const

/*
Returns a list of supported resolutions.

If non null video settings parameter is passed, the returned list is reduced to resolution supported with partial settings like video codec or frame rate applied.

If the encoder supports arbitrary resolutions within the supported resolutions range, *continuous is set to true, otherwise *continuous is set to false.

See also QVideoEncoderSettings::resolution().
*/
func (this *QVideoEncoderSettingsControl) SupportedResolutions(settings QVideoEncoderSettings_ITF, continuous *bool) *qtcore.QSizeList /*lll*/ {
	var convArg0 unsafe.Pointer
	if settings != nil && settings.QVideoEncoderSettings_PTR() != nil {
		convArg0 = settings.QVideoEncoderSettings_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK28QVideoEncoderSettingsControl20supportedResolutionsERK21QVideoEncoderSettingsPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, continuous)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeListFromPointer(unsafe.Pointer(uintptr(rv))) //5552
	return rv2
}

// /usr/include/qt/QtMultimedia/qvideoencodersettingscontrol.h:67
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QList<QSize> supportedResolutions(const QVideoEncoderSettings &, bool *) const

/*
Returns a list of supported resolutions.

If non null video settings parameter is passed, the returned list is reduced to resolution supported with partial settings like video codec or frame rate applied.

If the encoder supports arbitrary resolutions within the supported resolutions range, *continuous is set to true, otherwise *continuous is set to false.

See also QVideoEncoderSettings::resolution().
*/
func (this *QVideoEncoderSettingsControl) SupportedResolutions__(settings QVideoEncoderSettings_ITF) *qtcore.QSizeList /*lll*/ {
	var convArg0 unsafe.Pointer
	if settings != nil && settings.QVideoEncoderSettings_PTR() != nil {
		convArg0 = settings.QVideoEncoderSettings_PTR().GetCthis()
	}
	// arg: 1, bool *=Pointer, =Invalid, , Invalid
	var continuous unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK28QVideoEncoderSettingsControl20supportedResolutionsERK21QVideoEncoderSettingsPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, continuous)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeListFromPointer(unsafe.Pointer(uintptr(rv))) //5552
	return rv2
}

// /usr/include/qt/QtMultimedia/qvideoencodersettingscontrol.h:73
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QStringList supportedVideoCodecs() const

/*
Returns the list of supported video codecs.
*/
func (this *QVideoEncoderSettingsControl) SupportedVideoCodecs() *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK28QVideoEncoderSettingsControl20supportedVideoCodecsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtMultimedia/qvideoencodersettingscontrol.h:74
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QString videoCodecDescription(const QString &) const

/*
Returns a description of a video codec.
*/
func (this *QVideoEncoderSettingsControl) VideoCodecDescription(codecName string) string {
	var tmpArg0 = qtcore.NewQString_5(codecName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK28QVideoEncoderSettingsControl21videoCodecDescriptionERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtMultimedia/qvideoencodersettingscontrol.h:76
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QVideoEncoderSettings videoSettings() const

/*
Returns the video encoder settings.

The returned value may be different tha passed to QVideoEncoderSettingsControl::setVideoSettings() if the settings contains the default or undefined parameters. In this case if the undefined parameters are already resolved, they should be returned.

See also setVideoSettings().
*/
func (this *QVideoEncoderSettingsControl) VideoSettings() *QVideoEncoderSettings /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK28QVideoEncoderSettingsControl13videoSettingsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVideoEncoderSettingsFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVideoEncoderSettings)
	return rv2
}

// /usr/include/qt/QtMultimedia/qvideoencodersettingscontrol.h:77
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setVideoSettings(const QVideoEncoderSettings &)

/*
Sets the selected video encoder settings.

See also videoSettings().
*/
func (this *QVideoEncoderSettingsControl) SetVideoSettings(settings QVideoEncoderSettings_ITF) {
	var convArg0 unsafe.Pointer
	if settings != nil && settings.QVideoEncoderSettings_PTR() != nil {
		convArg0 = settings.QVideoEncoderSettings_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN28QVideoEncoderSettingsControl16setVideoSettingsERK21QVideoEncoderSettings", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qvideoencodersettingscontrol.h:80
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QVideoEncoderSettingsControl(QObject *)

/*
Create a new video encoder settings control object with the given parent.
*/
func (*QVideoEncoderSettingsControl) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QVideoEncoderSettingsControl {
	return NewQVideoEncoderSettingsControl(parent)
}
func NewQVideoEncoderSettingsControl(parent qtcore.QObject_ITF /*777 QObject **/) *QVideoEncoderSettingsControl {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN28QVideoEncoderSettingsControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVideoEncoderSettingsControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QVideoEncoderSettingsControl")
	return gothis
}

// /usr/include/qt/QtMultimedia/qvideoencodersettingscontrol.h:80
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QVideoEncoderSettingsControl(QObject *)

/*
Create a new video encoder settings control object with the given parent.
*/
func (*QVideoEncoderSettingsControl) NewForInherit__() *QVideoEncoderSettingsControl {
	return NewQVideoEncoderSettingsControl__()
}
func NewQVideoEncoderSettingsControl__() *QVideoEncoderSettingsControl {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN28QVideoEncoderSettingsControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVideoEncoderSettingsControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QVideoEncoderSettingsControl")
	return gothis
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
