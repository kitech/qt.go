package qtmultimedia

// /usr/include/qt/QtMultimedia/qaudioencodersettingscontrol.h
// #include <qaudioencodersettingscontrol.h>
// #include <QtMultimedia>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 54
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
type QAudioEncoderSettingsControl struct {
	*QMediaControl
}
type QAudioEncoderSettingsControl_ITF interface {
	QMediaControl_ITF
	QAudioEncoderSettingsControl_PTR() *QAudioEncoderSettingsControl
}

func (ptr *QAudioEncoderSettingsControl) QAudioEncoderSettingsControl_PTR() *QAudioEncoderSettingsControl {
	return ptr
}

func (this *QAudioEncoderSettingsControl) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QMediaControl.GetCthis()
	}
}
func (this *QAudioEncoderSettingsControl) SetCthis(cthis unsafe.Pointer) {
	this.QMediaControl = NewQMediaControlFromPointer(cthis)
}
func NewQAudioEncoderSettingsControlFromPointer(cthis unsafe.Pointer) *QAudioEncoderSettingsControl {
	bcthis0 := NewQMediaControlFromPointer(cthis)
	return &QAudioEncoderSettingsControl{bcthis0}
}
func (*QAudioEncoderSettingsControl) NewFromPointer(cthis unsafe.Pointer) *QAudioEncoderSettingsControl {
	return NewQAudioEncoderSettingsControlFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qaudioencodersettingscontrol.h:61
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QAudioEncoderSettingsControl) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK28QAudioEncoderSettingsControl10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qaudioencodersettingscontrol.h:64
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAudioEncoderSettingsControl()

/*

 */
func DeleteQAudioEncoderSettingsControl(this *QAudioEncoderSettingsControl) {
	rv, err := qtrt.InvokeQtFunc6("_ZN28QAudioEncoderSettingsControlD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qaudioencodersettingscontrol.h:66
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QStringList supportedAudioCodecs() const

/*
Returns the list of supported audio codec names.
*/
func (this *QAudioEncoderSettingsControl) SupportedAudioCodecs() *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK28QAudioEncoderSettingsControl20supportedAudioCodecsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtMultimedia/qaudioencodersettingscontrol.h:67
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QString codecDescription(const QString &) const

/*
Returns description of audio codec.
*/
func (this *QAudioEncoderSettingsControl) CodecDescription(codecName string) string {
	var tmpArg0 = qtcore.NewQString_5(codecName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK28QAudioEncoderSettingsControl16codecDescriptionERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtMultimedia/qaudioencodersettingscontrol.h:72
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QAudioEncoderSettings audioSettings() const

/*
Returns the audio encoder settings.

The returned value may be different tha passed to QAudioEncoderSettingsControl::setAudioSettings() if the settings contains the default or undefined parameters. In this case if the undefined parameters are already resolved, they should be returned.

See also setAudioSettings().
*/
func (this *QAudioEncoderSettingsControl) AudioSettings() *QAudioEncoderSettings /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK28QAudioEncoderSettingsControl13audioSettingsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQAudioEncoderSettingsFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQAudioEncoderSettings)
	return rv2
}

// /usr/include/qt/QtMultimedia/qaudioencodersettingscontrol.h:73
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setAudioSettings(const QAudioEncoderSettings &)

/*
Sets the selected audio settings.

See also audioSettings().
*/
func (this *QAudioEncoderSettingsControl) SetAudioSettings(arg0 QAudioEncoderSettings_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QAudioEncoderSettings_PTR() != nil {
		convArg0 = arg0.QAudioEncoderSettings_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN28QAudioEncoderSettingsControl16setAudioSettingsERK21QAudioEncoderSettings", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudioencodersettingscontrol.h:76
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QAudioEncoderSettingsControl(QObject *)

/*
Create a new audio encoder settings control object with the given parent.
*/
func (*QAudioEncoderSettingsControl) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QAudioEncoderSettingsControl {
	return NewQAudioEncoderSettingsControl(parent)
}
func NewQAudioEncoderSettingsControl(parent qtcore.QObject_ITF /*777 QObject **/) *QAudioEncoderSettingsControl {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN28QAudioEncoderSettingsControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAudioEncoderSettingsControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAudioEncoderSettingsControl")
	return gothis
}

// /usr/include/qt/QtMultimedia/qaudioencodersettingscontrol.h:76
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QAudioEncoderSettingsControl(QObject *)

/*
Create a new audio encoder settings control object with the given parent.
*/
func (*QAudioEncoderSettingsControl) NewForInherit__() *QAudioEncoderSettingsControl {
	return NewQAudioEncoderSettingsControl__()
}
func NewQAudioEncoderSettingsControl__() *QAudioEncoderSettingsControl {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN28QAudioEncoderSettingsControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAudioEncoderSettingsControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAudioEncoderSettingsControl")
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
