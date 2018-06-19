package qtpositioning

// /usr/include/qt/QtPositioning/qgeocoordinate.h
// #include <qgeocoordinate.h>
// #include <QtPositioning>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 26
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
type QGeoCoordinate struct {
	*qtrt.CObject
}
type QGeoCoordinate_ITF interface {
	QGeoCoordinate_PTR() *QGeoCoordinate
}

func (ptr *QGeoCoordinate) QGeoCoordinate_PTR() *QGeoCoordinate { return ptr }

func (this *QGeoCoordinate) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QGeoCoordinate) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQGeoCoordinateFromPointer(cthis unsafe.Pointer) *QGeoCoordinate {
	return &QGeoCoordinate{&qtrt.CObject{cthis}}
}
func (*QGeoCoordinate) NewFromPointer(cthis unsafe.Pointer) *QGeoCoordinate {
	return NewQGeoCoordinateFromPointer(cthis)
}

// /usr/include/qt/QtPositioning/qgeocoordinate.h:80
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGeoCoordinate()

/*
Constructs a coordinate. The coordinate will be invalid until setLatitude() and setLongitude() have been called.
*/
func NewQGeoCoordinate() *QGeoCoordinate {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGeoCoordinateC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGeoCoordinateFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGeoCoordinate)
	return gothis
}

// /usr/include/qt/QtPositioning/qgeocoordinate.h:81
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QGeoCoordinate(double, double)

/*
Constructs a coordinate. The coordinate will be invalid until setLatitude() and setLongitude() have been called.
*/
func NewQGeoCoordinate_1(latitude float64, longitude float64) *QGeoCoordinate {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGeoCoordinateC2Edd", qtrt.FFI_TYPE_POINTER, latitude, longitude)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGeoCoordinateFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGeoCoordinate)
	return gothis
}

// /usr/include/qt/QtPositioning/qgeocoordinate.h:82
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QGeoCoordinate(double, double, double)

/*
Constructs a coordinate. The coordinate will be invalid until setLatitude() and setLongitude() have been called.
*/
func NewQGeoCoordinate_2(latitude float64, longitude float64, altitude float64) *QGeoCoordinate {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGeoCoordinateC2Eddd", qtrt.FFI_TYPE_POINTER, latitude, longitude, altitude)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGeoCoordinateFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGeoCoordinate)
	return gothis
}

// /usr/include/qt/QtPositioning/qgeocoordinate.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QGeoCoordinate()

/*

 */
func DeleteQGeoCoordinate(this *QGeoCoordinate) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGeoCoordinateD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtPositioning/qgeocoordinate.h:86
// index:0
// Public Visibility=Default Availability=Available
// [8] QGeoCoordinate & operator=(const QGeoCoordinate &)

/*

 */
func (this *QGeoCoordinate) Operator_equal(other QGeoCoordinate_ITF) *QGeoCoordinate {
	var convArg0 unsafe.Pointer
	if other != nil && other.QGeoCoordinate_PTR() != nil {
		convArg0 = other.QGeoCoordinate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGeoCoordinateaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQGeoCoordinateFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQGeoCoordinate)
	return rv2
}

// /usr/include/qt/QtPositioning/qgeocoordinate.h:88
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QGeoCoordinate &) const

/*

 */
func (this *QGeoCoordinate) Operator_equal_equal(other QGeoCoordinate_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QGeoCoordinate_PTR() != nil {
		convArg0 = other.QGeoCoordinate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QGeoCoordinateeqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtPositioning/qgeocoordinate.h:89
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QGeoCoordinate &) const

/*

 */
func (this *QGeoCoordinate) Operator_not_equal(other QGeoCoordinate_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QGeoCoordinate_PTR() != nil {
		convArg0 = other.QGeoCoordinate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QGeoCoordinateneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtPositioning/qgeocoordinate.h:93
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValid() const

/*
Returns true if the longitude and latitude are valid.

Note: Getter function for property isValid.
*/
func (this *QGeoCoordinate) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QGeoCoordinate7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtPositioning/qgeocoordinate.h:94
// index:0
// Public Visibility=Default Availability=Available
// [4] QGeoCoordinate::CoordinateType type() const

/*
Returns the type of this coordinate.
*/
func (this *QGeoCoordinate) Type() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QGeoCoordinate4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtPositioning/qgeocoordinate.h:96
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLatitude(double)

/*
Sets the latitude (in decimal degrees) to latitude. The value should be in the WGS84 datum.

To be valid, the latitude must be between -90 to 90 inclusive.

Note: Setter function for property latitude.

See also latitude().
*/
func (this *QGeoCoordinate) SetLatitude(latitude float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGeoCoordinate11setLatitudeEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), latitude)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeocoordinate.h:97
// index:0
// Public Visibility=Default Availability=Available
// [8] double latitude() const

/*
Returns the latitude, in decimal degrees. The return value is undefined if the latitude has not been set.

A positive latitude indicates the Northern Hemisphere, and a negative latitude indicates the Southern Hemisphere.

Note: Getter function for property latitude.

See also setLatitude() and type().
*/
func (this *QGeoCoordinate) Latitude() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QGeoCoordinate8latitudeEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtPositioning/qgeocoordinate.h:99
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLongitude(double)

/*
Sets the longitude (in decimal degrees) to longitude. The value should be in the WGS84 datum.

To be valid, the longitude must be between -180 to 180 inclusive.

Note: Setter function for property longitude.

See also longitude().
*/
func (this *QGeoCoordinate) SetLongitude(longitude float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGeoCoordinate12setLongitudeEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), longitude)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeocoordinate.h:100
// index:0
// Public Visibility=Default Availability=Available
// [8] double longitude() const

/*
Returns the longitude, in decimal degrees. The return value is undefined if the longitude has not been set.

A positive longitude indicates the Eastern Hemisphere, and a negative longitude indicates the Western Hemisphere.

Note: Getter function for property longitude.

See also setLongitude() and type().
*/
func (this *QGeoCoordinate) Longitude() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QGeoCoordinate9longitudeEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtPositioning/qgeocoordinate.h:102
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAltitude(double)

/*
Sets the altitude (meters above sea level) to altitude.

Note: Setter function for property altitude.

See also altitude().
*/
func (this *QGeoCoordinate) SetAltitude(altitude float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGeoCoordinate11setAltitudeEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), altitude)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeocoordinate.h:103
// index:0
// Public Visibility=Default Availability=Available
// [8] double altitude() const

/*
Returns the altitude (meters above sea level).

The return value is undefined if the altitude has not been set.

Note: Getter function for property altitude.

See also setAltitude() and type().
*/
func (this *QGeoCoordinate) Altitude() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QGeoCoordinate8altitudeEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtPositioning/qgeocoordinate.h:105
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal distanceTo(const QGeoCoordinate &) const

/*
Returns the distance (in meters) from this coordinate to the coordinate specified by other. Altitude is not used in the calculation.

This calculation returns the great-circle distance between the two coordinates, with an assumption that the Earth is spherical for the purpose of this calculation.

Returns 0 if the type of this coordinate or the type of other is QGeoCoordinate::InvalidCoordinate.
*/
func (this *QGeoCoordinate) DistanceTo(other QGeoCoordinate_ITF) float64 {
	var convArg0 unsafe.Pointer
	if other != nil && other.QGeoCoordinate_PTR() != nil {
		convArg0 = other.QGeoCoordinate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QGeoCoordinate10distanceToERKS_", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtPositioning/qgeocoordinate.h:106
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal azimuthTo(const QGeoCoordinate &) const

/*
Returns the azimuth (or bearing) in degrees from this coordinate to the coordinate specified by other. Altitude is not used in the calculation.

The bearing returned is the bearing from the origin to other along the great-circle between the two coordinates. There is an assumption that the Earth is spherical for the purpose of this calculation.

Returns 0 if the type of this coordinate or the type of other is QGeoCoordinate::InvalidCoordinate.
*/
func (this *QGeoCoordinate) AzimuthTo(other QGeoCoordinate_ITF) float64 {
	var convArg0 unsafe.Pointer
	if other != nil && other.QGeoCoordinate_PTR() != nil {
		convArg0 = other.QGeoCoordinate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QGeoCoordinate9azimuthToERKS_", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtPositioning/qgeocoordinate.h:108
// index:0
// Public Visibility=Default Availability=Available
// [8] QGeoCoordinate atDistanceAndAzimuth(qreal, qreal, qreal) const

/*
Returns the coordinate that is reached by traveling distance meters from the current coordinate at azimuth (or bearing) along a great-circle. There is an assumption that the Earth is spherical for the purpose of this calculation.

The altitude will have distanceUp added to it.

Returns an invalid coordinate if this coordinate is invalid.
*/
func (this *QGeoCoordinate) AtDistanceAndAzimuth(distance float64, azimuth float64, distanceUp float64) *QGeoCoordinate /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QGeoCoordinate20atDistanceAndAzimuthEddd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), distance, azimuth, distanceUp)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQGeoCoordinateFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQGeoCoordinate)
	return rv2
}

// /usr/include/qt/QtPositioning/qgeocoordinate.h:108
// index:0
// Public Visibility=Default Availability=Available
// [8] QGeoCoordinate atDistanceAndAzimuth(qreal, qreal, qreal) const

/*
Returns the coordinate that is reached by traveling distance meters from the current coordinate at azimuth (or bearing) along a great-circle. There is an assumption that the Earth is spherical for the purpose of this calculation.

The altitude will have distanceUp added to it.

Returns an invalid coordinate if this coordinate is invalid.
*/
func (this *QGeoCoordinate) AtDistanceAndAzimuth__(distance float64, azimuth float64) *QGeoCoordinate /*123*/ {
	// arg: 2, qreal=Typedef, qreal=Typedef, double, Double
	distanceUp := float64(0.0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QGeoCoordinate20atDistanceAndAzimuthEddd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), distance, azimuth, distanceUp)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQGeoCoordinateFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQGeoCoordinate)
	return rv2
}

// /usr/include/qt/QtPositioning/qgeocoordinate.h:110
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toString(QGeoCoordinate::CoordinateFormat) const

/*
Returns this coordinate as a string in the specified format.

For example, if this coordinate has a latitude of -27.46758, a longitude of 153.027892 and an altitude of 28.1, these are the strings returned depending on format:


 format valueReturned string
Degrees-27.46758°, 153.02789°, 28.1m
DegreesWithHemisphere27.46758° S, 153.02789° E, 28.1m
DegreesMinutes-27° 28.054', 153° 1.673', 28.1m
DegreesMinutesWithHemisphere27° 28.054 S', 153° 1.673' E, 28.1m
DegreesMinutesSeconds-27° 28' 3.2", 153° 1' 40.4", 28.1m
DegreesMinutesSecondsWithHemisphere27° 28' 3.2" S, 153° 1' 40.4" E, 28.1m


The altitude field is omitted if no altitude is set.

If the coordinate is invalid, an empty string is returned.
*/
func (this *QGeoCoordinate) ToString(format int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QGeoCoordinate8toStringENS_16CoordinateFormatE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), format)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtPositioning/qgeocoordinate.h:110
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toString(QGeoCoordinate::CoordinateFormat) const

/*
Returns this coordinate as a string in the specified format.

For example, if this coordinate has a latitude of -27.46758, a longitude of 153.027892 and an altitude of 28.1, these are the strings returned depending on format:


 format valueReturned string
Degrees-27.46758°, 153.02789°, 28.1m
DegreesWithHemisphere27.46758° S, 153.02789° E, 28.1m
DegreesMinutes-27° 28.054', 153° 1.673', 28.1m
DegreesMinutesWithHemisphere27° 28.054 S', 153° 1.673' E, 28.1m
DegreesMinutesSeconds-27° 28' 3.2", 153° 1' 40.4", 28.1m
DegreesMinutesSecondsWithHemisphere27° 28' 3.2" S, 153° 1' 40.4" E, 28.1m


The altitude field is omitted if no altitude is set.

If the coordinate is invalid, an empty string is returned.
*/
func (this *QGeoCoordinate) ToString__() string {
	// arg: 0, QGeoCoordinate::CoordinateFormat=Enum, QGeoCoordinate::CoordinateFormat=Enum, , Invalid
	format := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QGeoCoordinate8toStringENS_16CoordinateFormatE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), format)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

/*
Defines the types of a coordinate.


*/
type QGeoCoordinate__CoordinateType = int

// An invalid coordinate. A coordinate is invalid if its latitude or longitude values are invalid.
const QGeoCoordinate__InvalidCoordinate QGeoCoordinate__CoordinateType = 0

// A coordinate with valid latitude and longitude values.
const QGeoCoordinate__Coordinate2D QGeoCoordinate__CoordinateType = 1

// A coordinate with valid latitude and longitude values, and also an altitude value.
const QGeoCoordinate__Coordinate3D QGeoCoordinate__CoordinateType = 2

/*
Defines the possible formatting options for toString().



See also toString().

*/
type QGeoCoordinate__CoordinateFormat = int

// Returns a string representation of the coordinates in decimal degrees format.
const QGeoCoordinate__Degrees QGeoCoordinate__CoordinateFormat = 0

// Returns a string representation of the coordinates in decimal degrees format, using 'N', 'S', 'E' or 'W' to indicate the hemispheres of the coordinates.
const QGeoCoordinate__DegreesWithHemisphere QGeoCoordinate__CoordinateFormat = 1

// Returns a string representation of the coordinates in degrees-minutes format.
const QGeoCoordinate__DegreesMinutes QGeoCoordinate__CoordinateFormat = 2

// Returns a string representation of the coordinates in degrees-minutes format, using 'N', 'S', 'E' or 'W' to indicate the hemispheres of the coordinates.
const QGeoCoordinate__DegreesMinutesWithHemisphere QGeoCoordinate__CoordinateFormat = 3

// Returns a string representation of the coordinates in degrees-minutes-seconds format.
const QGeoCoordinate__DegreesMinutesSeconds QGeoCoordinate__CoordinateFormat = 4

// Returns a string representation of the coordinates in degrees-minutes-seconds format, using 'N', 'S', 'E' or 'W' to indicate the hemispheres of the coordinates.
const QGeoCoordinate__DegreesMinutesSecondsWithHemisphere QGeoCoordinate__CoordinateFormat = 5

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
