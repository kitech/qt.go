package qtpositioning

// /usr/include/qt/QtPositioning/qgeosatelliteinfo.h
// #include <qgeosatelliteinfo.h>
// #include <QtPositioning>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 23
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
type QGeoSatelliteInfo struct {
	*qtrt.CObject
}
type QGeoSatelliteInfo_ITF interface {
	QGeoSatelliteInfo_PTR() *QGeoSatelliteInfo
}

func (ptr *QGeoSatelliteInfo) QGeoSatelliteInfo_PTR() *QGeoSatelliteInfo { return ptr }

func (this *QGeoSatelliteInfo) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QGeoSatelliteInfo) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQGeoSatelliteInfoFromPointer(cthis unsafe.Pointer) *QGeoSatelliteInfo {
	return &QGeoSatelliteInfo{&qtrt.CObject{cthis}}
}
func (*QGeoSatelliteInfo) NewFromPointer(cthis unsafe.Pointer) *QGeoSatelliteInfo {
	return NewQGeoSatelliteInfoFromPointer(cthis)
}

// /usr/include/qt/QtPositioning/qgeosatelliteinfo.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGeoSatelliteInfo()

/*
Creates a satellite information object.
*/
func (*QGeoSatelliteInfo) NewForInherit() *QGeoSatelliteInfo {
	return NewQGeoSatelliteInfo()
}
func NewQGeoSatelliteInfo() *QGeoSatelliteInfo {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGeoSatelliteInfoC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGeoSatelliteInfoFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGeoSatelliteInfo)
	return gothis
}

// /usr/include/qt/QtPositioning/qgeosatelliteinfo.h:66
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QGeoSatelliteInfo()

/*

 */
func DeleteQGeoSatelliteInfo(this *QGeoSatelliteInfo) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGeoSatelliteInfoD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtPositioning/qgeosatelliteinfo.h:68
// index:0
// Public Visibility=Default Availability=Available
// [8] QGeoSatelliteInfo & operator=(const QGeoSatelliteInfo &)

/*

 */
func (this *QGeoSatelliteInfo) Operator_equal(other QGeoSatelliteInfo_ITF) *QGeoSatelliteInfo {
	var convArg0 unsafe.Pointer
	if other != nil && other.QGeoSatelliteInfo_PTR() != nil {
		convArg0 = other.QGeoSatelliteInfo_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGeoSatelliteInfoaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQGeoSatelliteInfoFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQGeoSatelliteInfo)
	return rv2
}

// /usr/include/qt/QtPositioning/qgeosatelliteinfo.h:70
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QGeoSatelliteInfo &) const

/*

 */
func (this *QGeoSatelliteInfo) Operator_equal_equal(other QGeoSatelliteInfo_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QGeoSatelliteInfo_PTR() != nil {
		convArg0 = other.QGeoSatelliteInfo_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGeoSatelliteInfoeqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtPositioning/qgeosatelliteinfo.h:71
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QGeoSatelliteInfo &) const

/*

 */
func (this *QGeoSatelliteInfo) Operator_not_equal(other QGeoSatelliteInfo_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QGeoSatelliteInfo_PTR() != nil {
		convArg0 = other.QGeoSatelliteInfo_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGeoSatelliteInfoneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtPositioning/qgeosatelliteinfo.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSatelliteSystem(QGeoSatelliteInfo::SatelliteSystem)

/*
Sets the Satellite System (GPS, GLONASS, ...) to system.

See also satelliteSystem().
*/
func (this *QGeoSatelliteInfo) SetSatelliteSystem(system int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGeoSatelliteInfo18setSatelliteSystemENS_15SatelliteSystemE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), system)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeosatelliteinfo.h:76
// index:0
// Public Visibility=Default Availability=Available
// [4] QGeoSatelliteInfo::SatelliteSystem satelliteSystem() const

/*
Returns the Satellite System (GPS, GLONASS, ...)

See also setSatelliteSystem().
*/
func (this *QGeoSatelliteInfo) SatelliteSystem() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGeoSatelliteInfo15satelliteSystemEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtPositioning/qgeosatelliteinfo.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSatelliteIdentifier(int)

/*
Sets the satellite identifier number to satId.

The satellite identifier number can be used to identify a satellite inside the satellite system. For satellite system GPS the satellite identifier number represents the PRN (Pseudo-random noise) number. For satellite system GLONASS the satellite identifier number represents the slot number.

See also satelliteIdentifier().
*/
func (this *QGeoSatelliteInfo) SetSatelliteIdentifier(satId int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGeoSatelliteInfo22setSatelliteIdentifierEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), satId)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeosatelliteinfo.h:79
// index:0
// Public Visibility=Default Availability=Available
// [4] int satelliteIdentifier() const

/*
Returns the satellite identifier number.

The satellite identifier number can be used to identify a satellite inside the satellite system. For satellite system GPS the satellite identifier number represents the PRN (Pseudo-random noise) number. For satellite system GLONASS the satellite identifier number represents the slot number.

See also setSatelliteIdentifier().
*/
func (this *QGeoSatelliteInfo) SatelliteIdentifier() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGeoSatelliteInfo19satelliteIdentifierEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtPositioning/qgeosatelliteinfo.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSignalStrength(int)

/*
Sets the signal strength to signalStrength, in decibels.

See also signalStrength().
*/
func (this *QGeoSatelliteInfo) SetSignalStrength(signalStrength int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGeoSatelliteInfo17setSignalStrengthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), signalStrength)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeosatelliteinfo.h:82
// index:0
// Public Visibility=Default Availability=Available
// [4] int signalStrength() const

/*
Returns the signal strength, or -1 if the value has not been set.

See also setSignalStrength().
*/
func (this *QGeoSatelliteInfo) SignalStrength() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGeoSatelliteInfo14signalStrengthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtPositioning/qgeosatelliteinfo.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAttribute(QGeoSatelliteInfo::Attribute, qreal)

/*
Sets the value for attribute to value.

See also attribute().
*/
func (this *QGeoSatelliteInfo) SetAttribute(attribute int, value float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGeoSatelliteInfo12setAttributeENS_9AttributeEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), attribute, value)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeosatelliteinfo.h:85
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal attribute(QGeoSatelliteInfo::Attribute) const

/*
Returns the value of the specified attribute as a qreal value.

Returns -1 if the value has not been set.

See also hasAttribute() and setAttribute().
*/
func (this *QGeoSatelliteInfo) Attribute(attribute int) float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGeoSatelliteInfo9attributeENS_9AttributeE", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), attribute)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtPositioning/qgeosatelliteinfo.h:86
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeAttribute(QGeoSatelliteInfo::Attribute)

/*
Removes the specified attribute and its value.
*/
func (this *QGeoSatelliteInfo) RemoveAttribute(attribute int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGeoSatelliteInfo15removeAttributeENS_9AttributeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), attribute)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeosatelliteinfo.h:88
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasAttribute(QGeoSatelliteInfo::Attribute) const

/*
Returns true if the specified attribute is present in this update.
*/
func (this *QGeoSatelliteInfo) HasAttribute(attribute int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGeoSatelliteInfo12hasAttributeENS_9AttributeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), attribute)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

/*
Defines the attributes for the satellite information.


*/
type QGeoSatelliteInfo__Attribute = int

// The elevation of the satellite, in degrees.
const QGeoSatelliteInfo__Elevation QGeoSatelliteInfo__Attribute = 0

// The azimuth to true north, in degrees.
const QGeoSatelliteInfo__Azimuth QGeoSatelliteInfo__Attribute = 1

func (this *QGeoSatelliteInfo) AttributeItemName(val int) string {
	switch val {
	case QGeoSatelliteInfo__Elevation: // 0
		return "Elevation"
	case QGeoSatelliteInfo__Azimuth: // 1
		return "Azimuth"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QGeoSatelliteInfo_AttributeItemName(val int) string {
	var nilthis *QGeoSatelliteInfo
	return nilthis.AttributeItemName(val)
}

/*
Defines the GNSS system of the satellite.


*/
type QGeoSatelliteInfo__SatelliteSystem = int

//
const QGeoSatelliteInfo__Undefined QGeoSatelliteInfo__SatelliteSystem = 0

//
const QGeoSatelliteInfo__GPS QGeoSatelliteInfo__SatelliteSystem = 1

//
const QGeoSatelliteInfo__GLONASS QGeoSatelliteInfo__SatelliteSystem = 2

func (this *QGeoSatelliteInfo) SatelliteSystemItemName(val int) string {
	switch val {
	case QGeoSatelliteInfo__Undefined: // 0
		return "Undefined"
	case QGeoSatelliteInfo__GPS: // 1
		return "GPS"
	case QGeoSatelliteInfo__GLONASS: // 2
		return "GLONASS"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QGeoSatelliteInfo_SatelliteSystemItemName(val int) string {
	var nilthis *QGeoSatelliteInfo
	return nilthis.SatelliteSystemItemName(val)
}

//  body block end

//  keep block begin

func init_unused_11655() {
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
