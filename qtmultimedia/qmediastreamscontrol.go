package qtmultimedia

// /usr/include/qt/QtMultimedia/qmediastreamscontrol.h
// #include <qmediastreamscontrol.h>
// #include <QtMultimedia>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 3
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
type QMediaStreamsControl struct {
	*QMediaControl
}
type QMediaStreamsControl_ITF interface {
	QMediaControl_ITF
	QMediaStreamsControl_PTR() *QMediaStreamsControl
}

func (ptr *QMediaStreamsControl) QMediaStreamsControl_PTR() *QMediaStreamsControl { return ptr }

func (this *QMediaStreamsControl) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QMediaControl.GetCthis()
	}
}
func (this *QMediaStreamsControl) SetCthis(cthis unsafe.Pointer) {
	this.QMediaControl = NewQMediaControlFromPointer(cthis)
}
func NewQMediaStreamsControlFromPointer(cthis unsafe.Pointer) *QMediaStreamsControl {
	bcthis0 := NewQMediaControlFromPointer(cthis)
	return &QMediaStreamsControl{bcthis0}
}
func (*QMediaStreamsControl) NewFromPointer(cthis unsafe.Pointer) *QMediaStreamsControl {
	return NewQMediaStreamsControlFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qmediastreamscontrol.h:56
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QMediaStreamsControl) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QMediaStreamsControl10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qmediastreamscontrol.h:61
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QMediaStreamsControl()

/*

 */
func DeleteQMediaStreamsControl(this *QMediaStreamsControl) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QMediaStreamsControlD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qmediastreamscontrol.h:63
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] int streamCount()

/*
Returns the number of media streams.
*/
func (this *QMediaStreamsControl) StreamCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QMediaStreamsControl11streamCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qmediastreamscontrol.h:64
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] QMediaStreamsControl::StreamType streamType(int)

/*
Return the type of a media stream.
*/
func (this *QMediaStreamsControl) StreamType(streamNumber int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QMediaStreamsControl10streamTypeEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), streamNumber)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qmediastreamscontrol.h:66
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [16] QVariant metaData(int, const QString &)

/*
Returns the meta-data value of key for a given stream.

Useful metadata keys are QMediaMetaData::Title, QMediaMetaData::Description and QMediaMetaData::Language.
*/
func (this *QMediaStreamsControl) MetaData(streamNumber int, key string) *qtcore.QVariant /*123*/ {
	var tmpArg1 = qtcore.NewQString_5(key)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN20QMediaStreamsControl8metaDataEiRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), streamNumber, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtMultimedia/qmediastreamscontrol.h:68
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool isActive(int)

/*
Returns true if the media stream is active.
*/
func (this *QMediaStreamsControl) IsActive(streamNumber int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QMediaStreamsControl8isActiveEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), streamNumber)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qmediastreamscontrol.h:69
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setActive(int, bool)

/*
Sets the active state of a media stream.

Setting the active state of a media stream to true will activate it. If any other stream of the same type was previously active it will be deactivated. Setting the active state fo a media stream to false will deactivate it.

See also isActive().
*/
func (this *QMediaStreamsControl) SetActive(streamNumber int, state bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QMediaStreamsControl9setActiveEib", qtrt.FFI_TYPE_POINTER, this.GetCthis(), streamNumber, state)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediastreamscontrol.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void streamsChanged()

/*
The signal is emitted when the available streams list is changed.
*/
func (this *QMediaStreamsControl) StreamsChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QMediaStreamsControl14streamsChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediastreamscontrol.h:73
// index:0
// Public Visibility=Default Availability=Available
// [-2] void activeStreamsChanged()

/*
The signal is emitted when the active streams list is changed.
*/
func (this *QMediaStreamsControl) ActiveStreamsChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QMediaStreamsControl20activeStreamsChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediastreamscontrol.h:76
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QMediaStreamsControl(QObject *)

/*
Constructs a new media streams control with the given parent.
*/
func NewQMediaStreamsControl(parent qtcore.QObject_ITF /*777 QObject **/) *QMediaStreamsControl {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QMediaStreamsControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMediaStreamsControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QMediaStreamsControl")
	return gothis
}

// /usr/include/qt/QtMultimedia/qmediastreamscontrol.h:76
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QMediaStreamsControl(QObject *)

/*
Constructs a new media streams control with the given parent.
*/
func NewQMediaStreamsControl__() *QMediaStreamsControl {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN20QMediaStreamsControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMediaStreamsControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QMediaStreamsControl")
	return gothis
}

/*
Media stream type.

QMediaStreamsControl::DataStream4

*/
type QMediaStreamsControl__StreamType = int

// The stream type is unknown.
const QMediaStreamsControl__UnknownStream QMediaStreamsControl__StreamType = 0

// Video stream.
const QMediaStreamsControl__VideoStream QMediaStreamsControl__StreamType = 1

// Audio stream.
const QMediaStreamsControl__AudioStream QMediaStreamsControl__StreamType = 2

// Subpicture or teletext stream.
const QMediaStreamsControl__SubPictureStream QMediaStreamsControl__StreamType = 3

//
const QMediaStreamsControl__DataStream QMediaStreamsControl__StreamType = 4

func (this *QMediaStreamsControl) StreamTypeItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QMediaStreamsControl_StreamTypeItemName(val int) string {
	var nilthis *QMediaStreamsControl
	return nilthis.StreamTypeItemName(val)
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
