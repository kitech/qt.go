package qtmultimedia

// /usr/include/qt/QtMultimedia/qmediaobject.h
// #include <qmediaobject.h>
// #include <QtMultimedia>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 15
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

// void addPropertyWatch(const QByteArray &)
func (this *QMediaObject) InheritAddPropertyWatch(f func(name *qtcore.QByteArray) /*void*/) {
	qtrt.SetAllInheritCallback(this, "addPropertyWatch", f)
}

// void removePropertyWatch(const QByteArray &)
func (this *QMediaObject) InheritRemovePropertyWatch(f func(name *qtcore.QByteArray) /*void*/) {
	qtrt.SetAllInheritCallback(this, "removePropertyWatch", f)
}

/*

 */
type QMediaObject struct {
	*qtcore.QObject
}
type QMediaObject_ITF interface {
	qtcore.QObject_ITF
	QMediaObject_PTR() *QMediaObject
}

func (ptr *QMediaObject) QMediaObject_PTR() *QMediaObject { return ptr }

func (this *QMediaObject) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QMediaObject) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQMediaObjectFromPointer(cthis unsafe.Pointer) *QMediaObject {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QMediaObject{bcthis0}
}
func (*QMediaObject) NewFromPointer(cthis unsafe.Pointer) *QMediaObject {
	return NewQMediaObjectFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qmediaobject.h:58
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QMediaObject) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QMediaObject10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qmediaobject.h:61
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QMediaObject()

/*

 */
func DeleteQMediaObject(this *QMediaObject) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QMediaObjectD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qmediaobject.h:63
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool isAvailable() const

/*
Returns true if the service is available for use.
*/
func (this *QMediaObject) IsAvailable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QMediaObject11isAvailableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qmediaobject.h:64
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] QMultimedia::AvailabilityStatus availability() const

/*
Returns the availability of the functionality offered by this object.

In some cases the functionality may not be available (for example, if the current operating system or platform does not provide the required functionality), or it may be temporarily unavailable (for example, audio playback during a phone call or similar).
*/
func (this *QMediaObject) Availability() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QMediaObject12availabilityEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qmediaobject.h:66
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QMediaService * service() const

/*
Returns the media service that provides the functionality of this multimedia object.
*/
func (this *QMediaObject) Service() *QMediaService /*777 QMediaService **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QMediaObject7serviceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMediaServiceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qmediaobject.h:68
// index:0
// Public Visibility=Default Availability=Available
// [4] int notifyInterval() const

/*

 */
func (this *QMediaObject) NotifyInterval() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QMediaObject14notifyIntervalEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qmediaobject.h:69
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setNotifyInterval(int)

/*

 */
func (this *QMediaObject) SetNotifyInterval(milliSeconds int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QMediaObject17setNotifyIntervalEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), milliSeconds)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaobject.h:71
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool bind(QObject *)

/*
Bind object to this QMediaObject instance.

This method establishes a relationship between this media object and a helper object. The nature of the relationship depends on both parties. This methods returns true if the helper was successfully bound, false otherwise.

Most subclasses of QMediaObject provide more convenient functions that wrap this functionality, so this function rarely needs to be called directly.

The object passed must implement the QMediaBindableInterface interface.

See also QMediaBindableInterface.
*/
func (this *QMediaObject) Bind(arg0 qtcore.QObject_ITF /*777 QObject **/) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QObject_PTR() != nil {
		convArg0 = arg0.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QMediaObject4bindEP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qmediaobject.h:72
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void unbind(QObject *)

/*
Detach object from the QMediaObject instance.

Unbind the helper object from this media object. A warning will be generated if the object was not previously bound to this object.

See also QMediaBindableInterface.
*/
func (this *QMediaObject) Unbind(arg0 qtcore.QObject_ITF /*777 QObject **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QObject_PTR() != nil {
		convArg0 = arg0.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QMediaObject6unbindEP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaobject.h:74
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isMetaDataAvailable() const

/*
Returns true if there is meta-data associated with this media object, else false.
*/
func (this *QMediaObject) IsMetaDataAvailable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QMediaObject19isMetaDataAvailableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qmediaobject.h:76
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant metaData(const QString &) const

/*
Returns the value associated with a meta-data key.

See the list of predefined meta-data keys.
*/
func (this *QMediaObject) MetaData(key string) *qtcore.QVariant /*123*/ {
	var tmpArg0 = qtcore.NewQString5(key)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QMediaObject8metaDataERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtMultimedia/qmediaobject.h:77
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList availableMetaData() const

/*
Returns a list of keys there is meta-data available for.
*/
func (this *QMediaObject) AvailableMetaData() *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QMediaObject17availableMetaDataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtMultimedia/qmediaobject.h:80
// index:0
// Public Visibility=Default Availability=Available
// [-2] void notifyIntervalChanged(int)

/*
Signal a change in the notify interval period to milliseconds.

Note: Notifier signal for property notifyInterval.
*/
func (this *QMediaObject) NotifyIntervalChanged(milliSeconds int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QMediaObject21notifyIntervalChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), milliSeconds)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaobject.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void metaDataAvailableChanged(bool)

/*
Signals that the available state of a media object's meta-data has changed.
*/
func (this *QMediaObject) MetaDataAvailableChanged(available bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QMediaObject24metaDataAvailableChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), available)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaobject.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void metaDataChanged()

/*
Signals that this media object's meta-data has changed.

If multiple meta-data elements are changed, metaDataChanged(const QString &key, const QVariant &value) signal is emitted for each of them with metaDataChanged() changed emitted once.

Note: Signal metaDataChanged is overloaded in this class. To connect to this signal by using the function pointer syntax, Qt provides a convenient helper for obtaining the function pointer as shown in this example:


  connect(mediaObject, QOverload<>::of(&QMediaObject::metaDataChanged),
      [=](){ /-* ... *-/ });
*/
func (this *QMediaObject) MetaDataChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QMediaObject15metaDataChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaobject.h:84
// index:1
// Public Visibility=Default Availability=Available
// [-2] void metaDataChanged(const QString &, const QVariant &)

/*
Signals that this media object's meta-data has changed.

If multiple meta-data elements are changed, metaDataChanged(const QString &key, const QVariant &value) signal is emitted for each of them with metaDataChanged() changed emitted once.

Note: Signal metaDataChanged is overloaded in this class. To connect to this signal by using the function pointer syntax, Qt provides a convenient helper for obtaining the function pointer as shown in this example:


  connect(mediaObject, QOverload<>::of(&QMediaObject::metaDataChanged),
      [=](){ /-* ... *-/ });
*/
func (this *QMediaObject) MetaDataChanged1(key string, value qtcore.QVariant_ITF) {
	var tmpArg0 = qtcore.NewQString5(key)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg1 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QMediaObject15metaDataChangedERK7QStringRK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaobject.h:86
// index:0
// Public Visibility=Default Availability=Available
// [-2] void availabilityChanged(bool)

/*
Signal emitted when the availability state has changed to available.

Note: Signal availabilityChanged is overloaded in this class. To connect to this signal by using the function pointer syntax, Qt provides a convenient helper for obtaining the function pointer as shown in this example:


  connect(mediaObject, QOverload<bool>::of(&QMediaObject::availabilityChanged),
      [=](bool available){ /-* ... *-/ });
*/
func (this *QMediaObject) AvailabilityChanged(available bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QMediaObject19availabilityChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), available)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaobject.h:87
// index:1
// Public Visibility=Default Availability=Available
// [-2] void availabilityChanged(QMultimedia::AvailabilityStatus)

/*
Signal emitted when the availability state has changed to available.

Note: Signal availabilityChanged is overloaded in this class. To connect to this signal by using the function pointer syntax, Qt provides a convenient helper for obtaining the function pointer as shown in this example:


  connect(mediaObject, QOverload<bool>::of(&QMediaObject::availabilityChanged),
      [=](bool available){ /-* ... *-/ });
*/
func (this *QMediaObject) AvailabilityChanged1(availability int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QMediaObject19availabilityChangedEN11QMultimedia18AvailabilityStatusE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), availability)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaobject.h:90
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QMediaObject(QObject *, QMediaService *)

/*
Constructs a media object which uses the functionality provided by a media service.

The parent is passed to QObject.

This class is meant as a base class for multimedia objects so this constructor is protected.
*/
func (*QMediaObject) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/, service QMediaService_ITF /*777 QMediaService **/) *QMediaObject {
	return NewQMediaObject(parent, service)
}
func NewQMediaObject(parent qtcore.QObject_ITF /*777 QObject **/, service QMediaService_ITF /*777 QMediaService **/) *QMediaObject {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if service != nil && service.QMediaService_PTR() != nil {
		convArg1 = service.QMediaService_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QMediaObjectC2EP7QObjectP13QMediaService", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMediaObjectFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QMediaObject")
	return gothis
}

// /usr/include/qt/QtMultimedia/qmediaobject.h:93
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void addPropertyWatch(const QByteArray &)

/*
Watch the property name. The property's notify signal will be emitted once every notifyInterval milliseconds.

See also notifyInterval.
*/
func (this *QMediaObject) AddPropertyWatch(name qtcore.QByteArray_ITF) {
	var convArg0 unsafe.Pointer
	if name != nil && name.QByteArray_PTR() != nil {
		convArg0 = name.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QMediaObject16addPropertyWatchERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaobject.h:94
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void removePropertyWatch(const QByteArray &)

/*
Remove property name from the list of properties whose changes are regularly signaled.

See also notifyInterval.
*/
func (this *QMediaObject) RemovePropertyWatch(name qtcore.QByteArray_ITF) {
	var convArg0 unsafe.Pointer
	if name != nil && name.QByteArray_PTR() != nil {
		convArg0 = name.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QMediaObject19removePropertyWatchERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

//  body block end

//  keep block begin

func init_unused_11751() {
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
