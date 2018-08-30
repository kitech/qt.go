package qtmultimedia

// /usr/include/qt/QtMultimedia/qmetadatawritercontrol.h
// #include <qmetadatawritercontrol.h>
// #include <QtMultimedia>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 9
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
type QMetaDataWriterControl struct {
	*QMediaControl
}
type QMetaDataWriterControl_ITF interface {
	QMediaControl_ITF
	QMetaDataWriterControl_PTR() *QMetaDataWriterControl
}

func (ptr *QMetaDataWriterControl) QMetaDataWriterControl_PTR() *QMetaDataWriterControl { return ptr }

func (this *QMetaDataWriterControl) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QMediaControl.GetCthis()
	}
}
func (this *QMetaDataWriterControl) SetCthis(cthis unsafe.Pointer) {
	this.QMediaControl = NewQMediaControlFromPointer(cthis)
}
func NewQMetaDataWriterControlFromPointer(cthis unsafe.Pointer) *QMetaDataWriterControl {
	bcthis0 := NewQMediaControlFromPointer(cthis)
	return &QMetaDataWriterControl{bcthis0}
}
func (*QMetaDataWriterControl) NewFromPointer(cthis unsafe.Pointer) *QMetaDataWriterControl {
	return NewQMetaDataWriterControlFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qmetadatawritercontrol.h:58
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QMetaDataWriterControl) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QMetaDataWriterControl10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qmetadatawritercontrol.h:60
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QMetaDataWriterControl()

/*

 */
func DeleteQMetaDataWriterControl(this *QMetaDataWriterControl) {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QMetaDataWriterControlD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qmetadatawritercontrol.h:62
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool isWritable() const

/*
Identifies if a media service's meta-data can be edited.

Returns true if the meta-data is writable and false otherwise.
*/
func (this *QMetaDataWriterControl) IsWritable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QMetaDataWriterControl10isWritableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qmetadatawritercontrol.h:63
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool isMetaDataAvailable() const

/*
Identifies if meta-data is available from a media service.

Returns true if the meta-data is available and false otherwise.
*/
func (this *QMetaDataWriterControl) IsMetaDataAvailable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QMetaDataWriterControl19isMetaDataAvailableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qmetadatawritercontrol.h:65
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [16] QVariant metaData(const QString &) const

/*
Returns the meta-data for the given key.

See also setMetaData().
*/
func (this *QMetaDataWriterControl) MetaData(key string) *qtcore.QVariant /*123*/ {
	var tmpArg0 = qtcore.NewQString_5(key)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QMetaDataWriterControl8metaDataERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtMultimedia/qmetadatawritercontrol.h:66
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setMetaData(const QString &, const QVariant &)

/*
Sets the value of the meta-data element with the given key.

See also metaData().
*/
func (this *QMetaDataWriterControl) SetMetaData(key string, value qtcore.QVariant_ITF) {
	var tmpArg0 = qtcore.NewQString_5(key)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg1 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN22QMetaDataWriterControl11setMetaDataERK7QStringRK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmetadatawritercontrol.h:67
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QStringList availableMetaData() const

/*
Returns a list of keys there is meta-data available for.
*/
func (this *QMetaDataWriterControl) AvailableMetaData() *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QMetaDataWriterControl17availableMetaDataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtMultimedia/qmetadatawritercontrol.h:70
// index:0
// Public Visibility=Default Availability=Available
// [-2] void metaDataChanged()

/*
Signal the changes of meta-data.

If multiple meta-data elements are changed, metaDataChanged(const QString &key, const QVariant &value) signal is emitted for each of them with metaDataChanged() changed emitted once.

Note: Signal metaDataChanged is overloaded in this class. To connect to this signal by using the function pointer syntax, Qt provides a convenient helper for obtaining the function pointer as shown in this example:


  connect(metaDataWriterControl, QOverload<>::of(&QMetaDataWriterControl::metaDataChanged),
      [=](){ /-* ... *-/ });
*/
func (this *QMetaDataWriterControl) MetaDataChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QMetaDataWriterControl15metaDataChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmetadatawritercontrol.h:71
// index:1
// Public Visibility=Default Availability=Available
// [-2] void metaDataChanged(const QString &, const QVariant &)

/*
Signal the changes of meta-data.

If multiple meta-data elements are changed, metaDataChanged(const QString &key, const QVariant &value) signal is emitted for each of them with metaDataChanged() changed emitted once.

Note: Signal metaDataChanged is overloaded in this class. To connect to this signal by using the function pointer syntax, Qt provides a convenient helper for obtaining the function pointer as shown in this example:


  connect(metaDataWriterControl, QOverload<>::of(&QMetaDataWriterControl::metaDataChanged),
      [=](){ /-* ... *-/ });
*/
func (this *QMetaDataWriterControl) MetaDataChanged_1(key string, value qtcore.QVariant_ITF) {
	var tmpArg0 = qtcore.NewQString_5(key)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg1 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN22QMetaDataWriterControl15metaDataChangedERK7QStringRK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmetadatawritercontrol.h:73
// index:0
// Public Visibility=Default Availability=Available
// [-2] void writableChanged(bool)

/*
Signal a change in the writable status of meta-data, writable will be true if meta-data elements can be added or adjusted.
*/
func (this *QMetaDataWriterControl) WritableChanged(writable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QMetaDataWriterControl15writableChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), writable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmetadatawritercontrol.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void metaDataAvailableChanged(bool)

/*
Signal the availability of meta-data has changed, available will be true if the multimedia object has meta-data.
*/
func (this *QMetaDataWriterControl) MetaDataAvailableChanged(available bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QMetaDataWriterControl24metaDataAvailableChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), available)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmetadatawritercontrol.h:77
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QMetaDataWriterControl(QObject *)

/*
Construct a QMetaDataWriterControl with parent. This class is meant as a base class for service specific meta data providers so this constructor is protected.
*/
func NewQMetaDataWriterControl(parent qtcore.QObject_ITF /*777 QObject **/) *QMetaDataWriterControl {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN22QMetaDataWriterControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMetaDataWriterControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QMetaDataWriterControl")
	return gothis
}

// /usr/include/qt/QtMultimedia/qmetadatawritercontrol.h:77
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QMetaDataWriterControl(QObject *)

/*
Construct a QMetaDataWriterControl with parent. This class is meant as a base class for service specific meta data providers so this constructor is protected.
*/
func NewQMetaDataWriterControl__() *QMetaDataWriterControl {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN22QMetaDataWriterControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMetaDataWriterControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QMetaDataWriterControl")
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
