package qtpositioning

// /usr/include/qt/QtPositioning/qgeopositioninfo.h
// #include <qgeopositioninfo.h>
// #include <QtPositioning>

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

//  ext block end

//  body block begin

/*

 */
type QGeoPositionInfo struct {
	*qtrt.CObject
}
type QGeoPositionInfo_ITF interface {
	QGeoPositionInfo_PTR() *QGeoPositionInfo
}

func (ptr *QGeoPositionInfo) QGeoPositionInfo_PTR() *QGeoPositionInfo { return ptr }

func (this *QGeoPositionInfo) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QGeoPositionInfo) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQGeoPositionInfoFromPointer(cthis unsafe.Pointer) *QGeoPositionInfo {
	return &QGeoPositionInfo{&qtrt.CObject{cthis}}
}
func (*QGeoPositionInfo) NewFromPointer(cthis unsafe.Pointer) *QGeoPositionInfo {
	return NewQGeoPositionInfoFromPointer(cthis)
}

// /usr/include/qt/QtPositioning/qgeopositioninfo.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGeoPositionInfo()

/*
Creates an invalid QGeoPositionInfo object.

See also isValid().
*/
func NewQGeoPositionInfo() *QGeoPositionInfo {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QGeoPositionInfoC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGeoPositionInfoFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGeoPositionInfo)
	return gothis
}

// /usr/include/qt/QtPositioning/qgeopositioninfo.h:65
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QGeoPositionInfo(const QGeoCoordinate &, const QDateTime &)

/*
Creates an invalid QGeoPositionInfo object.

See also isValid().
*/
func NewQGeoPositionInfo_1(coordinate QGeoCoordinate_ITF, updateTime qtcore.QDateTime_ITF) *QGeoPositionInfo {
	var convArg0 unsafe.Pointer
	if coordinate != nil && coordinate.QGeoCoordinate_PTR() != nil {
		convArg0 = coordinate.QGeoCoordinate_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if updateTime != nil && updateTime.QDateTime_PTR() != nil {
		convArg1 = updateTime.QDateTime_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QGeoPositionInfoC2ERK14QGeoCoordinateRK9QDateTime", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGeoPositionInfoFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGeoPositionInfo)
	return gothis
}

// /usr/include/qt/QtPositioning/qgeopositioninfo.h:68
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QGeoPositionInfo()

/*

 */
func DeleteQGeoPositionInfo(this *QGeoPositionInfo) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QGeoPositionInfoD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtPositioning/qgeopositioninfo.h:70
// index:0
// Public Visibility=Default Availability=Available
// [8] QGeoPositionInfo & operator=(const QGeoPositionInfo &)

/*

 */
func (this *QGeoPositionInfo) Operator_equal(other QGeoPositionInfo_ITF) *QGeoPositionInfo {
	var convArg0 unsafe.Pointer
	if other != nil && other.QGeoPositionInfo_PTR() != nil {
		convArg0 = other.QGeoPositionInfo_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QGeoPositionInfoaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQGeoPositionInfoFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQGeoPositionInfo)
	return rv2
}

// /usr/include/qt/QtPositioning/qgeopositioninfo.h:72
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QGeoPositionInfo &) const

/*

 */
func (this *QGeoPositionInfo) Operator_equal_equal(other QGeoPositionInfo_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QGeoPositionInfo_PTR() != nil {
		convArg0 = other.QGeoPositionInfo_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QGeoPositionInfoeqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtPositioning/qgeopositioninfo.h:73
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QGeoPositionInfo &) const

/*

 */
func (this *QGeoPositionInfo) Operator_not_equal(other QGeoPositionInfo_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QGeoPositionInfo_PTR() != nil {
		convArg0 = other.QGeoPositionInfo_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QGeoPositionInfoneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtPositioning/qgeopositioninfo.h:77
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValid() const

/*
Returns true if the timestamp() and coordinate() values are both valid.

See also QGeoCoordinate::isValid() and QDateTime::isValid().
*/
func (this *QGeoPositionInfo) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QGeoPositionInfo7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtPositioning/qgeopositioninfo.h:79
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTimestamp(const QDateTime &)

/*
Sets the date and time at which this position was reported to timestamp.

The timestamp must be in UTC time.

See also timestamp().
*/
func (this *QGeoPositionInfo) SetTimestamp(timestamp qtcore.QDateTime_ITF) {
	var convArg0 unsafe.Pointer
	if timestamp != nil && timestamp.QDateTime_PTR() != nil {
		convArg0 = timestamp.QDateTime_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QGeoPositionInfo12setTimestampERK9QDateTime", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeopositioninfo.h:80
// index:0
// Public Visibility=Default Availability=Available
// [8] QDateTime timestamp() const

/*
Returns the date and time at which this position was reported, in UTC time.

Returns an invalid QDateTime if no date/time value has been set.

See also setTimestamp().
*/
func (this *QGeoPositionInfo) Timestamp() *qtcore.QDateTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QGeoPositionInfo9timestampEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQDateTime)
	return rv2
}

// /usr/include/qt/QtPositioning/qgeopositioninfo.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCoordinate(const QGeoCoordinate &)

/*
Sets the coordinate for this position to coordinate.

See also coordinate().
*/
func (this *QGeoPositionInfo) SetCoordinate(coordinate QGeoCoordinate_ITF) {
	var convArg0 unsafe.Pointer
	if coordinate != nil && coordinate.QGeoCoordinate_PTR() != nil {
		convArg0 = coordinate.QGeoCoordinate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QGeoPositionInfo13setCoordinateERK14QGeoCoordinate", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeopositioninfo.h:83
// index:0
// Public Visibility=Default Availability=Available
// [8] QGeoCoordinate coordinate() const

/*
Returns the coordinate for this position.

Returns an invalid coordinate if no coordinate has been set.

See also setCoordinate().
*/
func (this *QGeoPositionInfo) Coordinate() *QGeoCoordinate /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QGeoPositionInfo10coordinateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQGeoCoordinateFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQGeoCoordinate)
	return rv2
}

// /usr/include/qt/QtPositioning/qgeopositioninfo.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAttribute(QGeoPositionInfo::Attribute, qreal)

/*
Sets the value for attribute to value.

See also attribute().
*/
func (this *QGeoPositionInfo) SetAttribute(attribute int, value float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QGeoPositionInfo12setAttributeENS_9AttributeEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), attribute, value)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeopositioninfo.h:86
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal attribute(QGeoPositionInfo::Attribute) const

/*
Returns the value of the specified attribute as a qreal value.

Returns NaN if the value has not been set.

The function hasAttribute() should be used to determine whether or not a value has been set for an attribute.

See also hasAttribute() and setAttribute().
*/
func (this *QGeoPositionInfo) Attribute(attribute int) float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QGeoPositionInfo9attributeENS_9AttributeE", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), attribute)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtPositioning/qgeopositioninfo.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeAttribute(QGeoPositionInfo::Attribute)

/*
Removes the specified attribute and its value.
*/
func (this *QGeoPositionInfo) RemoveAttribute(attribute int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QGeoPositionInfo15removeAttributeENS_9AttributeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), attribute)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeopositioninfo.h:88
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasAttribute(QGeoPositionInfo::Attribute) const

/*
Returns true if the specified attribute is present for this QGeoPositionInfo object.
*/
func (this *QGeoPositionInfo) HasAttribute(attribute int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QGeoPositionInfo12hasAttributeENS_9AttributeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), attribute)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

/*
Defines the attributes for positional information.


*/
type QGeoPositionInfo__Attribute = int

// The bearing measured in degrees clockwise from true north to the direction of travel.
const QGeoPositionInfo__Direction QGeoPositionInfo__Attribute = 0

// The ground speed, in meters/sec.
const QGeoPositionInfo__GroundSpeed QGeoPositionInfo__Attribute = 1

// The vertical speed, in meters/sec.
const QGeoPositionInfo__VerticalSpeed QGeoPositionInfo__Attribute = 2

// The angle between the horizontal component of the magnetic field and true north, in degrees. Also known as magnetic declination. A positive value indicates a clockwise direction from true north and a negative value indicates a counter-clockwise direction.
const QGeoPositionInfo__MagneticVariation QGeoPositionInfo__Attribute = 3

// The accuracy of the provided latitude-longitude value, in meters.
const QGeoPositionInfo__HorizontalAccuracy QGeoPositionInfo__Attribute = 4

// The accuracy of the provided altitude value, in meters.
const QGeoPositionInfo__VerticalAccuracy QGeoPositionInfo__Attribute = 5

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
