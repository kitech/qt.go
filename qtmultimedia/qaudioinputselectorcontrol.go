package qtmultimedia

// /usr/include/qt/QtMultimedia/qaudioinputselectorcontrol.h
// #include <qaudioinputselectorcontrol.h>
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
type QAudioInputSelectorControl struct {
	*QMediaControl
}
type QAudioInputSelectorControl_ITF interface {
	QMediaControl_ITF
	QAudioInputSelectorControl_PTR() *QAudioInputSelectorControl
}

func (ptr *QAudioInputSelectorControl) QAudioInputSelectorControl_PTR() *QAudioInputSelectorControl {
	return ptr
}

func (this *QAudioInputSelectorControl) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QMediaControl.GetCthis()
	}
}
func (this *QAudioInputSelectorControl) SetCthis(cthis unsafe.Pointer) {
	this.QMediaControl = NewQMediaControlFromPointer(cthis)
}
func NewQAudioInputSelectorControlFromPointer(cthis unsafe.Pointer) *QAudioInputSelectorControl {
	bcthis0 := NewQMediaControlFromPointer(cthis)
	return &QAudioInputSelectorControl{bcthis0}
}
func (*QAudioInputSelectorControl) NewFromPointer(cthis unsafe.Pointer) *QAudioInputSelectorControl {
	return NewQAudioInputSelectorControlFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qaudioinputselectorcontrol.h:53
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QAudioInputSelectorControl) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK26QAudioInputSelectorControl10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qaudioinputselectorcontrol.h:56
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAudioInputSelectorControl()

/*

 */
func DeleteQAudioInputSelectorControl(this *QAudioInputSelectorControl) {
	rv, err := qtrt.InvokeQtFunc6("_ZN26QAudioInputSelectorControlD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qaudioinputselectorcontrol.h:59
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QString inputDescription(const QString &) const

/*
Returns the description of the input name.
*/
func (this *QAudioInputSelectorControl) InputDescription(name string) string {
	var tmpArg0 = qtcore.NewQString5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK26QAudioInputSelectorControl16inputDescriptionERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtMultimedia/qaudioinputselectorcontrol.h:60
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QString defaultInput() const

/*
Returns the name of the default audio input.
*/
func (this *QAudioInputSelectorControl) DefaultInput() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK26QAudioInputSelectorControl12defaultInputEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtMultimedia/qaudioinputselectorcontrol.h:61
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QString activeInput() const

/*
Returns the name of the currently selected audio input.

See also setActiveInput().
*/
func (this *QAudioInputSelectorControl) ActiveInput() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK26QAudioInputSelectorControl11activeInputEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtMultimedia/qaudioinputselectorcontrol.h:64
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setActiveInput(const QString &)

/*
Set the active audio input to name.

See also activeInput().
*/
func (this *QAudioInputSelectorControl) SetActiveInput(name string) {
	var tmpArg0 = qtcore.NewQString5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN26QAudioInputSelectorControl14setActiveInputERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudioinputselectorcontrol.h:67
// index:0
// Public Visibility=Default Availability=Available
// [-2] void activeInputChanged(const QString &)

/*
Signals that the audio input has changed to name.
*/
func (this *QAudioInputSelectorControl) ActiveInputChanged(name string) {
	var tmpArg0 = qtcore.NewQString5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN26QAudioInputSelectorControl18activeInputChangedERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudioinputselectorcontrol.h:68
// index:0
// Public Visibility=Default Availability=Available
// [-2] void availableInputsChanged()

/*
Signals that list of available inputs has changed.
*/
func (this *QAudioInputSelectorControl) AvailableInputsChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN26QAudioInputSelectorControl22availableInputsChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudioinputselectorcontrol.h:71
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QAudioInputSelectorControl(QObject *)

/*
Constructs a new audio input selector control with the given parent.
*/
func (*QAudioInputSelectorControl) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QAudioInputSelectorControl {
	return NewQAudioInputSelectorControl(parent)
}
func NewQAudioInputSelectorControl(parent qtcore.QObject_ITF /*777 QObject **/) *QAudioInputSelectorControl {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN26QAudioInputSelectorControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAudioInputSelectorControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAudioInputSelectorControl")
	return gothis
}

// /usr/include/qt/QtMultimedia/qaudioinputselectorcontrol.h:71
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QAudioInputSelectorControl(QObject *)

/*
Constructs a new audio input selector control with the given parent.
*/
func (*QAudioInputSelectorControl) NewForInheritp() *QAudioInputSelectorControl {
	return NewQAudioInputSelectorControlp()
}
func NewQAudioInputSelectorControlp() *QAudioInputSelectorControl {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN26QAudioInputSelectorControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAudioInputSelectorControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAudioInputSelectorControl")
	return gothis
}

//  body block end

//  keep block begin

func init_unused_11775() {
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
