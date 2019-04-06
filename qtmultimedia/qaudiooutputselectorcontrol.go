package qtmultimedia

// /usr/include/qt/QtMultimedia/qaudiooutputselectorcontrol.h
// #include <qaudiooutputselectorcontrol.h>
// #include <QtMultimedia>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 27
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
type QAudioOutputSelectorControl struct {
	*QMediaControl
}
type QAudioOutputSelectorControl_ITF interface {
	QMediaControl_ITF
	QAudioOutputSelectorControl_PTR() *QAudioOutputSelectorControl
}

func (ptr *QAudioOutputSelectorControl) QAudioOutputSelectorControl_PTR() *QAudioOutputSelectorControl {
	return ptr
}

func (this *QAudioOutputSelectorControl) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QMediaControl.GetCthis()
	}
}
func (this *QAudioOutputSelectorControl) SetCthis(cthis unsafe.Pointer) {
	this.QMediaControl = NewQMediaControlFromPointer(cthis)
}
func NewQAudioOutputSelectorControlFromPointer(cthis unsafe.Pointer) *QAudioOutputSelectorControl {
	bcthis0 := NewQMediaControlFromPointer(cthis)
	return &QAudioOutputSelectorControl{bcthis0}
}
func (*QAudioOutputSelectorControl) NewFromPointer(cthis unsafe.Pointer) *QAudioOutputSelectorControl {
	return NewQAudioOutputSelectorControlFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qaudiooutputselectorcontrol.h:53
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QAudioOutputSelectorControl) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK27QAudioOutputSelectorControl10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qaudiooutputselectorcontrol.h:56
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAudioOutputSelectorControl()

/*

 */
func DeleteQAudioOutputSelectorControl(this *QAudioOutputSelectorControl) {
	rv, err := qtrt.InvokeQtFunc6("_ZN27QAudioOutputSelectorControlD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qaudiooutputselectorcontrol.h:59
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QString outputDescription(const QString &) const

/*
Returns the description of the output name.
*/
func (this *QAudioOutputSelectorControl) OutputDescription(name string) string {
	var tmpArg0 = qtcore.NewQString5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK27QAudioOutputSelectorControl17outputDescriptionERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtMultimedia/qaudiooutputselectorcontrol.h:60
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QString defaultOutput() const

/*
Returns the name of the default audio output.
*/
func (this *QAudioOutputSelectorControl) DefaultOutput() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK27QAudioOutputSelectorControl13defaultOutputEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtMultimedia/qaudiooutputselectorcontrol.h:61
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QString activeOutput() const

/*
Returns the name of the currently selected audio output.

See also setActiveOutput().
*/
func (this *QAudioOutputSelectorControl) ActiveOutput() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK27QAudioOutputSelectorControl12activeOutputEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtMultimedia/qaudiooutputselectorcontrol.h:64
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setActiveOutput(const QString &)

/*
Set the active audio output to name.

See also activeOutput().
*/
func (this *QAudioOutputSelectorControl) SetActiveOutput(name string) {
	var tmpArg0 = qtcore.NewQString5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN27QAudioOutputSelectorControl15setActiveOutputERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudiooutputselectorcontrol.h:67
// index:0
// Public Visibility=Default Availability=Available
// [-2] void activeOutputChanged(const QString &)

/*
Signals that the audio output has changed to name.
*/
func (this *QAudioOutputSelectorControl) ActiveOutputChanged(name string) {
	var tmpArg0 = qtcore.NewQString5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN27QAudioOutputSelectorControl19activeOutputChangedERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudiooutputselectorcontrol.h:68
// index:0
// Public Visibility=Default Availability=Available
// [-2] void availableOutputsChanged()

/*
Signals that list of available outputs has changed.
*/
func (this *QAudioOutputSelectorControl) AvailableOutputsChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN27QAudioOutputSelectorControl23availableOutputsChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudiooutputselectorcontrol.h:71
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QAudioOutputSelectorControl(QObject *)

/*
Constructs a new audio output selector control with the given parent.
*/
func (*QAudioOutputSelectorControl) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QAudioOutputSelectorControl {
	return NewQAudioOutputSelectorControl(parent)
}
func NewQAudioOutputSelectorControl(parent qtcore.QObject_ITF /*777 QObject **/) *QAudioOutputSelectorControl {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN27QAudioOutputSelectorControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAudioOutputSelectorControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAudioOutputSelectorControl")
	return gothis
}

// /usr/include/qt/QtMultimedia/qaudiooutputselectorcontrol.h:71
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QAudioOutputSelectorControl(QObject *)

/*
Constructs a new audio output selector control with the given parent.
*/
func (*QAudioOutputSelectorControl) NewForInheritp() *QAudioOutputSelectorControl {
	return NewQAudioOutputSelectorControlp()
}
func NewQAudioOutputSelectorControlp() *QAudioOutputSelectorControl {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN27QAudioOutputSelectorControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAudioOutputSelectorControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAudioOutputSelectorControl")
	return gothis
}

//  body block end

//  keep block begin

func init_unused_11779() {
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
