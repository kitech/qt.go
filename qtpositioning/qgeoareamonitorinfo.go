package qtpositioning

// /usr/include/qt/QtPositioning/qgeoareamonitorinfo.h
// #include <qgeoareamonitorinfo.h>
// #include <QtPositioning>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 13
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

/*

 */
type QGeoAreaMonitorInfo struct {
	*qtrt.CObject
}
type QGeoAreaMonitorInfo_ITF interface {
	QGeoAreaMonitorInfo_PTR() *QGeoAreaMonitorInfo
}

func (ptr *QGeoAreaMonitorInfo) QGeoAreaMonitorInfo_PTR() *QGeoAreaMonitorInfo { return ptr }

func (this *QGeoAreaMonitorInfo) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QGeoAreaMonitorInfo) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQGeoAreaMonitorInfoFromPointer(cthis unsafe.Pointer) *QGeoAreaMonitorInfo {
	return &QGeoAreaMonitorInfo{&qtrt.CObject{cthis}}
}
func (*QGeoAreaMonitorInfo) NewFromPointer(cthis unsafe.Pointer) *QGeoAreaMonitorInfo {
	return NewQGeoAreaMonitorInfoFromPointer(cthis)
}

// /usr/include/qt/QtPositioning/qgeoareamonitorinfo.h:62
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGeoAreaMonitorInfo(const QString &)

/*
Constructs a QGeoAreaMonitorInfo object with the specified name.

See also name().
*/
func (*QGeoAreaMonitorInfo) NewForInherit(name string) *QGeoAreaMonitorInfo {
	return NewQGeoAreaMonitorInfo(name)
}
func NewQGeoAreaMonitorInfo(name string) *QGeoAreaMonitorInfo {
	var tmpArg0 = qtcore.NewQString5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGeoAreaMonitorInfoC2ERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGeoAreaMonitorInfoFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGeoAreaMonitorInfo)
	return gothis
}

// /usr/include/qt/QtPositioning/qgeoareamonitorinfo.h:62
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGeoAreaMonitorInfo(const QString &)

/*
Constructs a QGeoAreaMonitorInfo object with the specified name.

See also name().
*/
func (*QGeoAreaMonitorInfo) NewForInheritp() *QGeoAreaMonitorInfo {
	return NewQGeoAreaMonitorInfop()
}
func NewQGeoAreaMonitorInfop() *QGeoAreaMonitorInfo {
	// arg: 0, const QString &=LValueReference, QString=Record, , Invalid
	var convArg0 = qtcore.NewQString()
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGeoAreaMonitorInfoC2ERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGeoAreaMonitorInfoFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGeoAreaMonitorInfo)
	return gothis
}

// /usr/include/qt/QtPositioning/qgeoareamonitorinfo.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QGeoAreaMonitorInfo()

/*

 */
func DeleteQGeoAreaMonitorInfo(this *QGeoAreaMonitorInfo) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGeoAreaMonitorInfoD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtPositioning/qgeoareamonitorinfo.h:66
// index:0
// Public Visibility=Default Availability=Available
// [8] QGeoAreaMonitorInfo & operator=(const QGeoAreaMonitorInfo &)

/*

 */
func (this *QGeoAreaMonitorInfo) Operator_equal(other QGeoAreaMonitorInfo_ITF) *QGeoAreaMonitorInfo {
	var convArg0 unsafe.Pointer
	if other != nil && other.QGeoAreaMonitorInfo_PTR() != nil {
		convArg0 = other.QGeoAreaMonitorInfo_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGeoAreaMonitorInfoaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQGeoAreaMonitorInfoFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQGeoAreaMonitorInfo)
	return rv2
}

// /usr/include/qt/QtPositioning/qgeoareamonitorinfo.h:68
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QGeoAreaMonitorInfo &) const

/*

 */
func (this *QGeoAreaMonitorInfo) Operator_equal_equal(other QGeoAreaMonitorInfo_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QGeoAreaMonitorInfo_PTR() != nil {
		convArg0 = other.QGeoAreaMonitorInfo_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGeoAreaMonitorInfoeqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtPositioning/qgeoareamonitorinfo.h:69
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator!=(const QGeoAreaMonitorInfo &) const

/*

 */
func (this *QGeoAreaMonitorInfo) Operator_not_equal(other QGeoAreaMonitorInfo_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QGeoAreaMonitorInfo_PTR() != nil {
		convArg0 = other.QGeoAreaMonitorInfo_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGeoAreaMonitorInfoneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtPositioning/qgeoareamonitorinfo.h:71
// index:0
// Public Visibility=Default Availability=Available
// [8] QString name() const

/*
Returns the name of the QGeoAreaMonitorInfo object. The name should be used to for user-visibility purposes.

See also setName().
*/
func (this *QGeoAreaMonitorInfo) Name() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGeoAreaMonitorInfo4nameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtPositioning/qgeoareamonitorinfo.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setName(const QString &)

/*
Sets the user visibile name.

See also name().
*/
func (this *QGeoAreaMonitorInfo) SetName(name string) {
	var tmpArg0 = qtcore.NewQString5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGeoAreaMonitorInfo7setNameERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeoareamonitorinfo.h:74
// index:0
// Public Visibility=Default Availability=Available
// [8] QString identifier() const

/*
Returns the identifier of the QGeoAreaMonitorInfo object. The identifier is automatically generated upon construction of a new QGeoAreaMonitorInfo object.
*/
func (this *QGeoAreaMonitorInfo) Identifier() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGeoAreaMonitorInfo10identifierEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtPositioning/qgeoareamonitorinfo.h:75
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValid() const

/*
Returns true, if the monitor is valid. A valid QGeoAreaMonitorInfo has a non-empty name() and the monitored area is not empty(). Otherwise this function returns false.
*/
func (this *QGeoAreaMonitorInfo) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGeoAreaMonitorInfo7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtPositioning/qgeoareamonitorinfo.h:77
// index:0
// Public Visibility=Default Availability=Available
// [8] QGeoShape area() const

/*
Returns the boundaries of the to-be-monitored area. This area must not be empty.

See also setArea().
*/
func (this *QGeoAreaMonitorInfo) Area() *QGeoShape /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGeoAreaMonitorInfo4areaEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQGeoShapeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQGeoShape)
	return rv2
}

// /usr/include/qt/QtPositioning/qgeoareamonitorinfo.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setArea(const QGeoShape &)

/*
Sets the to-be-monitored area to newShape.

See also area().
*/
func (this *QGeoAreaMonitorInfo) SetArea(newShape QGeoShape_ITF) {
	var convArg0 unsafe.Pointer
	if newShape != nil && newShape.QGeoShape_PTR() != nil {
		convArg0 = newShape.QGeoShape_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGeoAreaMonitorInfo7setAreaERK9QGeoShape", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeoareamonitorinfo.h:80
// index:0
// Public Visibility=Default Availability=Available
// [8] QDateTime expiration() const

/*
Returns the expiry date.

After an active QGeoAreaMonitorInfo has expired the region is no longer monitored and the QGeoAreaMonitorInfo object is removed from the list of active monitors.

If the expiry QDateTime is invalid the QGeoAreaMonitorInfo object is treated as not having an expiry date. This implies an indefinite monitoring period if the object is persistent or until the current application closes if the object is non-persistent.

See also setExpiration() and QGeoAreaMonitorSource::activeMonitors().
*/
func (this *QGeoAreaMonitorInfo) Expiration() *qtcore.QDateTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGeoAreaMonitorInfo10expirationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQDateTime)
	return rv2
}

// /usr/include/qt/QtPositioning/qgeoareamonitorinfo.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setExpiration(const QDateTime &)

/*
Sets the expiry date and time to expiry.

See also expiration().
*/
func (this *QGeoAreaMonitorInfo) SetExpiration(expiry qtcore.QDateTime_ITF) {
	var convArg0 unsafe.Pointer
	if expiry != nil && expiry.QDateTime_PTR() != nil {
		convArg0 = expiry.QDateTime_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGeoAreaMonitorInfo13setExpirationERK9QDateTime", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeoareamonitorinfo.h:83
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isPersistent() const

/*
Returns true if the QGeoAreaMonitorInfo is persistent. The default value for this property is false.

A non-persistent QGeoAreaMonitorInfo will be removed by the system once the application owning the monitor object stops. Persistent objects remain active and can be retrieved once the application restarts.

If the system triggers an event associated to a persistent QGeoAreaMonitorInfo the relevant application will be re-started and the appropriate signal emitted.

See also setPersistent().
*/
func (this *QGeoAreaMonitorInfo) IsPersistent() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGeoAreaMonitorInfo12isPersistentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtPositioning/qgeoareamonitorinfo.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPersistent(bool)

/*
Sets the QGeoAreaMonitorInfo objects persistence to isPersistent.

Note that setting this flag does not imply that QGeoAreaMonitorInfoSource supports persistent monitoring. QGeoAreaMonitorSource::supportedAreaMonitorFeatures() can be used to check for this feature's availability.

See also isPersistent().
*/
func (this *QGeoAreaMonitorInfo) SetPersistent(isPersistent bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGeoAreaMonitorInfo13setPersistentEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), isPersistent)
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
}

//  keep block end
