package qtpositioning

// /usr/include/qt/QtPositioning/qgeolocation.h
// #include <qgeolocation.h>
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
type QGeoLocation struct {
	*qtrt.CObject
}
type QGeoLocation_ITF interface {
	QGeoLocation_PTR() *QGeoLocation
}

func (ptr *QGeoLocation) QGeoLocation_PTR() *QGeoLocation { return ptr }

func (this *QGeoLocation) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QGeoLocation) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQGeoLocationFromPointer(cthis unsafe.Pointer) *QGeoLocation {
	return &QGeoLocation{&qtrt.CObject{cthis}}
}
func (*QGeoLocation) NewFromPointer(cthis unsafe.Pointer) *QGeoLocation {
	return NewQGeoLocationFromPointer(cthis)
}

// /usr/include/qt/QtPositioning/qgeolocation.h:57
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGeoLocation()

/*
Constructs an new location object.
*/
func (*QGeoLocation) NewForInherit() *QGeoLocation {
	return NewQGeoLocation()
}
func NewQGeoLocation() *QGeoLocation {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QGeoLocationC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGeoLocationFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGeoLocation)
	return gothis
}

// /usr/include/qt/QtPositioning/qgeolocation.h:60
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QGeoLocation()

/*

 */
func DeleteQGeoLocation(this *QGeoLocation) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QGeoLocationD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtPositioning/qgeolocation.h:62
// index:0
// Public Visibility=Default Availability=Available
// [8] QGeoLocation & operator=(const QGeoLocation &)

/*

 */
func (this *QGeoLocation) Operator_equal(other QGeoLocation_ITF) *QGeoLocation {
	var convArg0 unsafe.Pointer
	if other != nil && other.QGeoLocation_PTR() != nil {
		convArg0 = other.QGeoLocation_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QGeoLocationaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQGeoLocationFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQGeoLocation)
	return rv2
}

// /usr/include/qt/QtPositioning/qgeolocation.h:64
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QGeoLocation &) const

/*

 */
func (this *QGeoLocation) Operator_equal_equal(other QGeoLocation_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QGeoLocation_PTR() != nil {
		convArg0 = other.QGeoLocation_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QGeoLocationeqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtPositioning/qgeolocation.h:65
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QGeoLocation &) const

/*

 */
func (this *QGeoLocation) Operator_not_equal(other QGeoLocation_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QGeoLocation_PTR() != nil {
		convArg0 = other.QGeoLocation_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QGeoLocationneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtPositioning/qgeolocation.h:69
// index:0
// Public Visibility=Default Availability=Available
// [8] QGeoAddress address() const

/*
Returns the address of the location.

See also setAddress().
*/
func (this *QGeoLocation) Address() *QGeoAddress /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QGeoLocation7addressEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQGeoAddressFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQGeoAddress)
	return rv2
}

// /usr/include/qt/QtPositioning/qgeolocation.h:70
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAddress(const QGeoAddress &)

/*
Sets the address of the location.

See also address().
*/
func (this *QGeoLocation) SetAddress(address QGeoAddress_ITF) {
	var convArg0 unsafe.Pointer
	if address != nil && address.QGeoAddress_PTR() != nil {
		convArg0 = address.QGeoAddress_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QGeoLocation10setAddressERK11QGeoAddress", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeolocation.h:71
// index:0
// Public Visibility=Default Availability=Available
// [8] QGeoCoordinate coordinate() const

/*
Returns the coordinate of the location.

See also setCoordinate().
*/
func (this *QGeoLocation) Coordinate() *QGeoCoordinate /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QGeoLocation10coordinateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQGeoCoordinateFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQGeoCoordinate)
	return rv2
}

// /usr/include/qt/QtPositioning/qgeolocation.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCoordinate(const QGeoCoordinate &)

/*
Sets the coordinate of the location.

See also coordinate().
*/
func (this *QGeoLocation) SetCoordinate(position QGeoCoordinate_ITF) {
	var convArg0 unsafe.Pointer
	if position != nil && position.QGeoCoordinate_PTR() != nil {
		convArg0 = position.QGeoCoordinate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QGeoLocation13setCoordinateERK14QGeoCoordinate", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeolocation.h:73
// index:0
// Public Visibility=Default Availability=Available
// [8] QGeoRectangle boundingBox() const

/*
Returns a bounding box which represents the recommended region to display when viewing this location.

For example, a building's location may have a region centered around the building, but the region is large enough to show it's immediate surrounding geographical context.

See also setBoundingBox().
*/
func (this *QGeoLocation) BoundingBox() *QGeoRectangle /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QGeoLocation11boundingBoxEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQGeoRectangleFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQGeoRectangle)
	return rv2
}

// /usr/include/qt/QtPositioning/qgeolocation.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBoundingBox(const QGeoRectangle &)

/*
Sets the boundingBox of the location.

See also boundingBox().
*/
func (this *QGeoLocation) SetBoundingBox(box QGeoRectangle_ITF) {
	var convArg0 unsafe.Pointer
	if box != nil && box.QGeoRectangle_PTR() != nil {
		convArg0 = box.QGeoRectangle_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QGeoLocation14setBoundingBoxERK13QGeoRectangle", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeolocation.h:76
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEmpty() const

/*
Returns true if all fields of the location are 0; otherwise returns false.
*/
func (this *QGeoLocation) IsEmpty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QGeoLocation7isEmptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
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
