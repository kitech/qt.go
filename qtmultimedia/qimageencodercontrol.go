package qtmultimedia

// /usr/include/qt/QtMultimedia/qimageencodercontrol.h
// #include <qimageencodercontrol.h>
// #include <QtMultimedia>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 16
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
type QImageEncoderControl struct {
	*QMediaControl
}
type QImageEncoderControl_ITF interface {
	QMediaControl_ITF
	QImageEncoderControl_PTR() *QImageEncoderControl
}

func (ptr *QImageEncoderControl) QImageEncoderControl_PTR() *QImageEncoderControl { return ptr }

func (this *QImageEncoderControl) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QMediaControl.GetCthis()
	}
}
func (this *QImageEncoderControl) SetCthis(cthis unsafe.Pointer) {
	this.QMediaControl = NewQMediaControlFromPointer(cthis)
}
func NewQImageEncoderControlFromPointer(cthis unsafe.Pointer) *QImageEncoderControl {
	bcthis0 := NewQMediaControlFromPointer(cthis)
	return &QImageEncoderControl{bcthis0}
}
func (*QImageEncoderControl) NewFromPointer(cthis unsafe.Pointer) *QImageEncoderControl {
	return NewQImageEncoderControlFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qimageencodercontrol.h:62
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QImageEncoderControl) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QImageEncoderControl10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qimageencodercontrol.h:65
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QImageEncoderControl()

/*

 */
func DeleteQImageEncoderControl(this *QImageEncoderControl) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QImageEncoderControlD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qimageencodercontrol.h:67
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QStringList supportedImageCodecs() const

/*
Returns a list of supported image codecs.
*/
func (this *QImageEncoderControl) SupportedImageCodecs() *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QImageEncoderControl20supportedImageCodecsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtMultimedia/qimageencodercontrol.h:68
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QString imageCodecDescription(const QString &) const

/*
Returns a description of an image codec.
*/
func (this *QImageEncoderControl) ImageCodecDescription(codecName string) string {
	var tmpArg0 = qtcore.NewQString_5(codecName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QImageEncoderControl21imageCodecDescriptionERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtMultimedia/qimageencodercontrol.h:70
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QList<QSize> supportedResolutions(const QImageEncoderSettings &, bool *) const

/*
Returns a list of supported resolutions.

If non null image settings parameter is passed, the returned list is reduced to resolutions supported with partial settings applied. It can be used to query the list of resolutions, supported by specific image codec.

If the encoder supports arbitrary resolutions within the supported resolutions range, *continuous is set to true, otherwise *continuous is set to false.
*/
func (this *QImageEncoderControl) SupportedResolutions(settings QImageEncoderSettings_ITF, continuous *bool) *qtcore.QSizeList /*lll*/ {
	var convArg0 unsafe.Pointer
	if settings != nil && settings.QImageEncoderSettings_PTR() != nil {
		convArg0 = settings.QImageEncoderSettings_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QImageEncoderControl20supportedResolutionsERK21QImageEncoderSettingsPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, continuous)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeListFromPointer(unsafe.Pointer(uintptr(rv))) //5552
	return rv2
}

// /usr/include/qt/QtMultimedia/qimageencodercontrol.h:70
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QList<QSize> supportedResolutions(const QImageEncoderSettings &, bool *) const

/*
Returns a list of supported resolutions.

If non null image settings parameter is passed, the returned list is reduced to resolutions supported with partial settings applied. It can be used to query the list of resolutions, supported by specific image codec.

If the encoder supports arbitrary resolutions within the supported resolutions range, *continuous is set to true, otherwise *continuous is set to false.
*/
func (this *QImageEncoderControl) SupportedResolutions__(settings QImageEncoderSettings_ITF) *qtcore.QSizeList /*lll*/ {
	var convArg0 unsafe.Pointer
	if settings != nil && settings.QImageEncoderSettings_PTR() != nil {
		convArg0 = settings.QImageEncoderSettings_PTR().GetCthis()
	}
	// arg: 1, bool *=Pointer, =Invalid, , Invalid
	var continuous unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QImageEncoderControl20supportedResolutionsERK21QImageEncoderSettingsPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, continuous)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeListFromPointer(unsafe.Pointer(uintptr(rv))) //5552
	return rv2
}

// /usr/include/qt/QtMultimedia/qimageencodercontrol.h:73
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QImageEncoderSettings imageSettings() const

/*
Returns the currently used image encoder settings.

The returned value may be different tha passed to QImageEncoderControl::setImageSettings() if the settings contains the default or undefined parameters. In this case if the undefined parameters are already resolved, they should be returned.

See also setImageSettings().
*/
func (this *QImageEncoderControl) ImageSettings() *QImageEncoderSettings /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QImageEncoderControl13imageSettingsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageEncoderSettingsFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImageEncoderSettings)
	return rv2
}

// /usr/include/qt/QtMultimedia/qimageencodercontrol.h:74
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setImageSettings(const QImageEncoderSettings &)

/*
Sets the selected image encoder settings.

See also imageSettings().
*/
func (this *QImageEncoderControl) SetImageSettings(settings QImageEncoderSettings_ITF) {
	var convArg0 unsafe.Pointer
	if settings != nil && settings.QImageEncoderSettings_PTR() != nil {
		convArg0 = settings.QImageEncoderSettings_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QImageEncoderControl16setImageSettingsERK21QImageEncoderSettings", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qimageencodercontrol.h:77
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QImageEncoderControl(QObject *)

/*
Constructs a new image encoder control object with the given parent
*/
func (*QImageEncoderControl) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QImageEncoderControl {
	return NewQImageEncoderControl(parent)
}
func NewQImageEncoderControl(parent qtcore.QObject_ITF /*777 QObject **/) *QImageEncoderControl {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QImageEncoderControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQImageEncoderControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QImageEncoderControl")
	return gothis
}

// /usr/include/qt/QtMultimedia/qimageencodercontrol.h:77
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QImageEncoderControl(QObject *)

/*
Constructs a new image encoder control object with the given parent
*/
func (*QImageEncoderControl) NewForInherit__() *QImageEncoderControl {
	return NewQImageEncoderControl__()
}
func NewQImageEncoderControl__() *QImageEncoderControl {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN20QImageEncoderControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQImageEncoderControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QImageEncoderControl")
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
