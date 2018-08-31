package qtmultimedia

// /usr/include/qt/QtMultimedia/qaudiorecorder.h
// #include <qaudiorecorder.h>
// #include <QtMultimedia>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 8
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
type QAudioRecorder struct {
	*QMediaRecorder
}
type QAudioRecorder_ITF interface {
	QMediaRecorder_ITF
	QAudioRecorder_PTR() *QAudioRecorder
}

func (ptr *QAudioRecorder) QAudioRecorder_PTR() *QAudioRecorder { return ptr }

func (this *QAudioRecorder) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QMediaRecorder.GetCthis()
	}
}
func (this *QAudioRecorder) SetCthis(cthis unsafe.Pointer) {
	this.QMediaRecorder = NewQMediaRecorderFromPointer(cthis)
}
func NewQAudioRecorderFromPointer(cthis unsafe.Pointer) *QAudioRecorder {
	bcthis0 := NewQMediaRecorderFromPointer(cthis)
	return &QAudioRecorder{bcthis0}
}
func (*QAudioRecorder) NewFromPointer(cthis unsafe.Pointer) *QAudioRecorder {
	return NewQAudioRecorderFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qaudiorecorder.h:62
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QAudioRecorder) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QAudioRecorder10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qaudiorecorder.h:65
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAudioRecorder(QObject *)

/*
Constructs an audio recorder. The parent is passed to QMediaObject.
*/
func (*QAudioRecorder) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QAudioRecorder {
	return NewQAudioRecorder(parent)
}
func NewQAudioRecorder(parent qtcore.QObject_ITF /*777 QObject **/) *QAudioRecorder {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QAudioRecorderC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAudioRecorderFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAudioRecorder")
	return gothis
}

// /usr/include/qt/QtMultimedia/qaudiorecorder.h:65
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAudioRecorder(QObject *)

/*
Constructs an audio recorder. The parent is passed to QMediaObject.
*/
func (*QAudioRecorder) NewForInherit__() *QAudioRecorder {
	return NewQAudioRecorder__()
}
func NewQAudioRecorder__() *QAudioRecorder {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN14QAudioRecorderC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAudioRecorderFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAudioRecorder")
	return gothis
}

// /usr/include/qt/QtMultimedia/qaudiorecorder.h:66
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAudioRecorder()

/*

 */
func DeleteQAudioRecorder(this *QAudioRecorder) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QAudioRecorderD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 32)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qaudiorecorder.h:68
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList audioInputs() const

/*
Returns a list of available audio inputs
*/
func (this *QAudioRecorder) AudioInputs() *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QAudioRecorder11audioInputsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtMultimedia/qaudiorecorder.h:69
// index:0
// Public Visibility=Default Availability=Available
// [8] QString defaultAudioInput() const

/*
Returns the default audio input name.
*/
func (this *QAudioRecorder) DefaultAudioInput() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QAudioRecorder17defaultAudioInputEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtMultimedia/qaudiorecorder.h:70
// index:0
// Public Visibility=Default Availability=Available
// [8] QString audioInputDescription(const QString &) const

/*
Returns the readable translated description of the audio input device with name.
*/
func (this *QAudioRecorder) AudioInputDescription(name string) string {
	var tmpArg0 = qtcore.NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QAudioRecorder21audioInputDescriptionERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtMultimedia/qaudiorecorder.h:72
// index:0
// Public Visibility=Default Availability=Available
// [8] QString audioInput() const

/*
Returns the active audio input name.

Note: Getter function for property audioInput.

See also setAudioInput().
*/
func (this *QAudioRecorder) AudioInput() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QAudioRecorder10audioInputEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtMultimedia/qaudiorecorder.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAudioInput(const QString &)

/*
Set the active audio input to name.

Note: Setter function for property audioInput.

See also audioInput().
*/
func (this *QAudioRecorder) SetAudioInput(name string) {
	var tmpArg0 = qtcore.NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QAudioRecorder13setAudioInputERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudiorecorder.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void audioInputChanged(const QString &)

/*
Signal emitted when active audio input changes to name.

Note: Notifier signal for property audioInput.
*/
func (this *QAudioRecorder) AudioInputChanged(name string) {
	var tmpArg0 = qtcore.NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QAudioRecorder17audioInputChangedERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudiorecorder.h:79
// index:0
// Public Visibility=Default Availability=Available
// [-2] void availableAudioInputsChanged()

/*
Signal is emitted when the available audio inputs change.
*/
func (this *QAudioRecorder) AvailableAudioInputsChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QAudioRecorder27availableAudioInputsChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
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
