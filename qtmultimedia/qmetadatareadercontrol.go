package qtmultimedia

// /usr/include/qt/QtMultimedia/qmetadatareadercontrol.h
// #include <qmetadatareadercontrol.h>
// #include <QtMultimedia>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 5
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
type QMetaDataReaderControl struct {
	*QMediaControl
}
type QMetaDataReaderControl_ITF interface {
	QMediaControl_ITF
	QMetaDataReaderControl_PTR() *QMetaDataReaderControl
}

func (ptr *QMetaDataReaderControl) QMetaDataReaderControl_PTR() *QMetaDataReaderControl { return ptr }

func (this *QMetaDataReaderControl) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QMediaControl.GetCthis()
	}
}
func (this *QMetaDataReaderControl) SetCthis(cthis unsafe.Pointer) {
	this.QMediaControl = NewQMediaControlFromPointer(cthis)
}
func NewQMetaDataReaderControlFromPointer(cthis unsafe.Pointer) *QMetaDataReaderControl {
	bcthis0 := NewQMediaControlFromPointer(cthis)
	return &QMetaDataReaderControl{bcthis0}
}
func (*QMetaDataReaderControl) NewFromPointer(cthis unsafe.Pointer) *QMetaDataReaderControl {
	return NewQMetaDataReaderControlFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qmetadatareadercontrol.h:58
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QMetaDataReaderControl) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QMetaDataReaderControl10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qmetadatareadercontrol.h:60
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QMetaDataReaderControl()

/*

 */
func DeleteQMetaDataReaderControl(this *QMetaDataReaderControl) {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QMetaDataReaderControlD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qmetadatareadercontrol.h:62
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool isMetaDataAvailable() const

/*
Identifies if meta-data is available from a media service.

Returns true if the meta-data is available and false otherwise.
*/
func (this *QMetaDataReaderControl) IsMetaDataAvailable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QMetaDataReaderControl19isMetaDataAvailableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qmetadatareadercontrol.h:64
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [16] QVariant metaData(const QString &) const

/*
Returns the meta-data for the given key.
*/
func (this *QMetaDataReaderControl) MetaData(key string) *qtcore.QVariant /*123*/ {
	var tmpArg0 = qtcore.NewQString5(key)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QMetaDataReaderControl8metaDataERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtMultimedia/qmetadatareadercontrol.h:65
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QStringList availableMetaData() const

/*
Returns a list of keys there is meta-data available for.
*/
func (this *QMetaDataReaderControl) AvailableMetaData() *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QMetaDataReaderControl17availableMetaDataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtMultimedia/qmetadatareadercontrol.h:68
// index:0
// Public Visibility=Default Availability=Available
// [-2] void metaDataChanged()

/*
Signal the changes of meta-data.

If multiple meta-data elements are changed, metaDataChanged(const QString &key, const QVariant &value) signal is emitted for each of them with metaDataChanged() changed emitted once.

Note: Signal metaDataChanged is overloaded in this class. To connect to this one using the function pointer syntax, you must specify the signal type in a static cast, as shown in this example:


  connect(metaDataReaderControl, static_cast<void(QMetaDataReaderControl::*)()>(&QMetaDataReaderControl::metaDataChanged),
      [=](){ /-* ... *-/ });
*/
func (this *QMetaDataReaderControl) MetaDataChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QMetaDataReaderControl15metaDataChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmetadatareadercontrol.h:69
// index:1
// Public Visibility=Default Availability=Available
// [-2] void metaDataChanged(const QString &, const QVariant &)

/*
Signal the changes of meta-data.

If multiple meta-data elements are changed, metaDataChanged(const QString &key, const QVariant &value) signal is emitted for each of them with metaDataChanged() changed emitted once.

Note: Signal metaDataChanged is overloaded in this class. To connect to this one using the function pointer syntax, you must specify the signal type in a static cast, as shown in this example:


  connect(metaDataReaderControl, static_cast<void(QMetaDataReaderControl::*)()>(&QMetaDataReaderControl::metaDataChanged),
      [=](){ /-* ... *-/ });
*/
func (this *QMetaDataReaderControl) MetaDataChanged1(key string, value qtcore.QVariant_ITF) {
	var tmpArg0 = qtcore.NewQString5(key)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg1 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN22QMetaDataReaderControl15metaDataChangedERK7QStringRK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmetadatareadercontrol.h:71
// index:0
// Public Visibility=Default Availability=Available
// [-2] void metaDataAvailableChanged(bool)

/*
Signal the availability of meta-data has changed, available will be true if the multimedia object has meta-data.
*/
func (this *QMetaDataReaderControl) MetaDataAvailableChanged(available bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QMetaDataReaderControl24metaDataAvailableChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), available)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmetadatareadercontrol.h:74
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QMetaDataReaderControl(QObject *)

/*
Construct a QMetaDataReaderControl with parent. This class is meant as a base class for service specific meta data providers so this constructor is protected.
*/
func (*QMetaDataReaderControl) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QMetaDataReaderControl {
	return NewQMetaDataReaderControl(parent)
}
func NewQMetaDataReaderControl(parent qtcore.QObject_ITF /*777 QObject **/) *QMetaDataReaderControl {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN22QMetaDataReaderControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMetaDataReaderControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QMetaDataReaderControl")
	return gothis
}

// /usr/include/qt/QtMultimedia/qmetadatareadercontrol.h:74
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QMetaDataReaderControl(QObject *)

/*
Construct a QMetaDataReaderControl with parent. This class is meant as a base class for service specific meta data providers so this constructor is protected.
*/
func (*QMetaDataReaderControl) NewForInheritp() *QMetaDataReaderControl {
	return NewQMetaDataReaderControlp()
}
func NewQMetaDataReaderControlp() *QMetaDataReaderControl {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN22QMetaDataReaderControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMetaDataReaderControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QMetaDataReaderControl")
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
